package main

import (
	"fmt"
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
	db := m3o.NewDB(m3oApiKey).WithTablePrefix("tg_tm")
	tm := timeline.NewTimeline(db, time.Now().Format("200601"))

	list, err := tm.Read("")

	fmt.Println("-->>", list, err)

	bs, err := yaml.Marshal(list)

	fmt.Println("-->>", string(bs))

}
