package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/h1z3y3/h1z3y3.github.io/timeline/telegram_bot/m3o"
	"github.com/h1z3y3/h1z3y3.github.io/timeline/telegram_bot/timeline"
	"gopkg.in/yaml.v2"
)

var (
	m3oApiKey = os.Getenv("TG_M3O_TOKEN")
)

func main() {
	table := time.Now().Format("200601")
	db := m3o.NewDB(m3oApiKey).WithTablePrefix("tg_tm")
	tm := timeline.NewTimeline(db, table)

	list, err := tm.Read("")

	if err != nil {
		log.Fatal(err)
	}

	if len(list) == 0 {
		log.Println("empty list")
		return
	}

	bs, err := yaml.Marshal(map[string]interface{}{
		"date": list[0].Timestamp,
		"list": list,
	})

	fmt.Println("yaml:")
	fmt.Println(string(bs))

	f, err := os.Create(fmt.Sprintf("../../data/emo/emo%s.yaml", table))
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(bs)

	if err != nil {
		log.Fatal(err)
	}

}
