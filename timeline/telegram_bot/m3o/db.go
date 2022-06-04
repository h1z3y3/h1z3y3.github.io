package m3o

import (
	"time"

	dbIface "github.com/h1z3y3/h1z3y3.github.io/timeline/telegram_bot/db"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

const (
	m3oDBPath = "/v1/db"
)

type db struct {
	prefix string
	apiKey string
}

func NewDB(apiKey string) *db {
	return &db{
		apiKey: apiKey,
	}
}

func getDBApi(path string) string {
	return m3oDomain + m3oDBPath + path
}

func (d *db) tableName(tableName string) string {
	return d.prefix + "_" + tableName
}

func (d *db) WithTablePrefix(prefix string) *db {
	d.prefix = prefix
	return d
}

func (d *db) ListTables() ([]string, error) {
	resp, err := httplib.Get(getDBApi("/ListTables")).
		Header("Authorization", "Bearer "+d.apiKey).
		Header("Content-Type", "application/json").
		Retries(3).RetryDelay(1*time.Second).
		SetTimeout(2*time.Second, 5*time.Second).String()

	if err != nil {
		return nil, err
	}

	code := gjson.Get(resp, "code").Uint()
	if code > 0 {
		return nil, errors.New(resp)
	}

	tables := gjson.Get(resp, "tables").Array()
	if len(tables) == 0 {
		return []string{}, nil
	}

	result := []string{}

	for _, v := range tables {
		result = append(result, v.Str)
	}

	return result, nil
}

func (d *db) DropTable(table string) error {
	data := map[string]interface{}{
		"table": d.tableName(table),
	}

	req, err := httplib.Post(getDBApi("/DropTable")).
		Header("Authorization", "Bearer "+d.apiKey).
		Header("Content-Type", "application/json").
		SetTimeout(2*time.Second, 5*time.Second).
		Retries(3).RetryDelay(1 * time.Second).
		JSONBody(data)

	if err != nil {
		return err
	}

	_, err = req.String()

	if err != nil {
		return err
	}

	return nil
}

func (d *db) Create(table string, record dbIface.Record) (string, error) {
	data := map[string]interface{}{
		"table":  d.tableName(table),
		"record": record,
	}

	req, err := httplib.Post(getDBApi("/Create")).
		Header("Authorization", "Bearer "+d.apiKey).
		Header("Content-Type", "application/json").
		SetTimeout(2*time.Second, 5*time.Second).
		Retries(3).RetryDelay(1 * time.Second).
		JSONBody(data)

	if err != nil {
		return "", err
	}

	resp, err := req.String()

	if err != nil {
		return "", err
	}

	code := gjson.Get(resp, "code").Uint()
	if code > 0 {
		return "", errors.New(resp)
	}

	return gjson.Get(resp, "id").Str, nil
}

func (d *db) Count(table string) (uint64, error) {
	data := map[string]interface{}{
		"table": d.tableName(table),
	}

	req, err := httplib.Post(getDBApi("/Count")).
		Header("Authorization", "Bearer "+d.apiKey).
		Header("Content-Type", "application/json").
		SetTimeout(2*time.Second, 5*time.Second).
		Retries(3).RetryDelay(1 * time.Second).
		JSONBody(data)

	if err != nil {
		return 0, err
	}

	resp, err := req.String()

	if err != nil {
		return 0, err
	}

	code := gjson.Get(resp, "code").Uint()
	if code > 0 {
		return 0, errors.New(resp)
	}

	return gjson.Get(resp, "count").Uint(), nil
}

func (d *db) Delete(table string, id string) error {
	data := map[string]interface{}{
		"table": d.tableName(table),
		"id":    id,
	}

	req, err := httplib.Post(getDBApi("/Delete")).
		Header("Authorization", "Bearer "+d.apiKey).
		Header("Content-Type", "application/json").
		SetTimeout(2*time.Second, 5*time.Second).
		Retries(3).RetryDelay(1 * time.Second).
		JSONBody(data)

	if err != nil {
		return err
	}

	resp, err := req.String()

	if err != nil {
		return err
	}

	code := gjson.Get(resp, "code").Uint()
	if code > 0 {
		return errors.New(resp)
	}

	return nil
}

func (d *db) Read(table, query string) ([]dbIface.Record, error) {
	data := map[string]interface{}{
		"table": d.tableName(table),
		"query": query,
	}

	req, err := httplib.Post(getDBApi("/Read")).
		Header("Authorization", "Bearer "+d.apiKey).
		Header("Content-Type", "application/json").
		SetTimeout(2*time.Second, 5*time.Second).
		Retries(3).RetryDelay(1 * time.Second).
		JSONBody(data)

	if err != nil {
		return nil, err
	}

	resp, err := req.String()

	if err != nil {
		return nil, err
	}

	code := gjson.Get(resp, "code").Uint()
	if code > 0 {
		return nil, errors.New(resp)
	}

	records := gjson.Get(resp, "records").Array()
	if len(records) == 0 {
		return []dbIface.Record{}, nil
	}

	result := []dbIface.Record{}

	for _, r := range records {
		m := map[string]interface{}{}
		for k, v := range r.Map() {
			m[k] = v.Value()
		}
		result = append(result, m)
	}

	return result, nil
}
