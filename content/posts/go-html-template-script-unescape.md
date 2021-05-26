---
title: Golang unescape values in html templates
date: 2021-05-25 21:56:44
tags: [Golang]
---

While using values in a html template, my colleague faced a problem that the values in `<script>` tags are all escaped.
So the template developer can not use the values in javascript. It just looks like this:

```
// ... 

<script type="text/json" id="channel" ch="recommend" tk="">
   [&#34;news&#34;,&#34;comedy&#34;,&#34;cartoon&#34;,&#34;tech&#34;,&#34;travelling&#34;,&#34;fashion&#34;,&#34;photograph&#34;,&#34;household&#34;,&#34;movies&#34;,&#34;foods&#34;,&#34;military&#34;,&#34;health&#34;,&#34;test&#34;]
</script>

// ...
```

I tried to find out the reason. So I copied the template values in other HTML tags rather than in `<script>`.
It rendered correctly, and was not escaped: 

```
// ...

<div>
  ["news","comedy","cartoon","tech","travelling","fashion","photograph","household","movies","foods","military","health","test"]
<div>

// ...
```

First, I found some reasons like this, they said using `json.Encoder` rather than using `json.Marshal()` directly. 
So I can `SetEscapeHTML(false)`. This is the example code:

```go
func toRawJson(v interface{}) (string, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(&v)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(buf.String(), "\n"), nil
}
```

I tried, but it didn't work either.

Now, let's think about for a moment. Only `<script>` has the problem, but other html tags not. 
So if I could say the value is actually correct, but the template rendered in wrong ?

I started google the problems about template render but not the wrong value.

Then I found the `html/template` official document:

Let's see its introduction: `html/template` [introduction](https://golang.org/pkg/html/template/#hdr-Introduction):

> This package wraps package text/template so you can share its template API 
> to parse and execute HTML templates safely.
> 
> ```
> tmpl, err := template.New("name").Parse(...)
> // Error checking elided
> err = tmpl.Execute(out, data)
> ```
> 
> If successful, tmpl will now be injection-safe. Otherwise, 
> err is an error defined in the docs for ErrorCode.
> 
> HTML templates treat data values as plain text which should be encoded so they 
> can be safely embedded in an HTML document. 
> The escaping is contextual, so actions can appear within JavaScript, CSS, and URI contexts.
> The security model used by this package assumes that template authors are trusted, 
> while Execute's data parameter is not. More details are provided below.
>
> Example
> ```
> import "text/template"
> ...
> t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
> err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
> ```
>
> produces
> 
> ```
> Hello, <script>alert('you have been pwned')</script>!
> ```
> 
> but the contextual autoescaping in html/template
> 
> ```
> import "html/template"
> ...
> t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
> err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
> ```
> 
> produces safe, escaped HTML output
> 
> ```
> Hello, &lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;!
> ```


\\o//. Now we find out the reason, but how we resolve this problem ?

Maybe I can provide a template function to unescape the value.

```go 
	r := gin.Default()

	funcMap["unescapeHTML"] = func(s string) template.HTML {
		return template.HTML(s)
	}
	
	r.SetFuncMap(funcMap)
```

Then you can use it as below:

```
<script type="text/json" id="channel" ch="recommend" tk="">
    {{ .channel_data | unescapeHTML }}
</script>
```

Finally, it works.

```
// ... 

<script type="text/json" id="channel" ch="recommend" tk="">
   ["news","comedy","cartoon","tech","travelling","fashion","photograph","household","movies","foods","military","health","test"]
</script>

// ...
```




