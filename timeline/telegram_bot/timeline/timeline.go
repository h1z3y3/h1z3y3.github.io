package timeline

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	dbIface "github.com/h1z3y3/h1z3y3.github.io/timeline/telegram_bot/db"
	"github.com/pkg/errors"
)

type timeline struct {
	db    dbIface.DB
	table string
}

type Timeline struct {
	Id        string   `json:"id"`
	Content   string   `json:"content"`
	Timestamp int64    `json:"timestamp"`
	Images    []string `json:"images"`
}

func (t Timeline) String() string {
	bs, _ := json.MarshalIndent(t, "", "  ")
	return string(bs)
}

func (t Timeline) ToDBRecord() dbIface.Record {
	return dbIface.Record{
		"id":        t.Id,
		"content":   t.Content,
		"timestamp": time.Now().Unix(),
		"images":    t.Images,
	}
}

func parseDBRecord(record dbIface.Record) Timeline {
	t := Timeline{
		Id:        record["id"].(string),
		Content:   record["content"].(string),
		Timestamp: int64(record["timestamp"].(float64)),
	}

	if v, ok := record["images"]; v != nil && ok {
		images := v.([]interface{})
		for _, img := range images {
			t.Images = append(t.Images, img.(string))
		}
	}

	return t
}

type Timelines []Timeline

func (ts Timelines) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func (ts Timelines) Less(i, j int) bool {
	return ts[i].Timestamp < ts[j].Timestamp
}

func (ts Timelines) Len() int {
	return len(ts)
}

func (ts Timelines) String() string {
	var list []string
	for _, v := range ts {
		list = append(list, v.String())
	}
	return strings.Join(list, "\n")
}

func NewTimeline(db dbIface.DB, table string) *timeline {
	return &timeline{
		db:    db,
		table: table,
	}
}

func (t *timeline) Add(tm Timeline) (string, error) {
	if tm.Content == "" {
		return "", errors.New("content is required")
	}

	tm.Id = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", time.Now().UnixNano()))))
	id, err := t.db.Create(t.table, tm.ToDBRecord())

	if err != nil {
		return "", errors.Wrap(err, "create record error")
	}

	count, err := t.db.Count(t.table)
	if err != nil {
		return id, errors.Wrap(err, "get record count error")
	}

	return fmt.Sprintf("id: %s\ncontent: %s\ncount: %d", id, tm.Content, count), nil
}

func (t *timeline) Update(tm Timeline) error {
	if tm.Id == "" {
		return errors.New("id is required!")
	}

	err := t.db.Update(t.table, tm.ToDBRecord())
	if err != nil {
		return errors.Wrap(err, "update record error")
	}

	return nil
}

func (t *timeline) Delete(id string) (string, error) {
	if id == "" {
		return "", errors.New("id is required")
	}

	ids := strings.Split(id, " ")
	var errs []string
	wg := sync.WaitGroup{}
	for _, v := range ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			err := t.db.Delete(t.table, id)
			if err != nil {
				errs = append(errs, err.Error())
			}
		}(v)

	}
	wg.Wait()

	if len(errs) > 0 {
		return "", errors.New(strings.Join(errs, "\n"))
	}

	count, err := t.db.Count(t.table)
	if err != nil {
		return id, errors.Wrap(err, "get record count error")
	}

	return fmt.Sprintf("deleted id: %s\ncount: %d", id, count), nil
}

func (t *timeline) Read(query string) (Timelines, error) {
	records, err := t.db.Read(t.table, query)
	if err != nil {
		return []Timeline{}, errors.Wrap(err, "get records error")
	}

	result := []Timeline{}

	for _, v := range records {
		result = append(result, parseDBRecord(v))
	}

	sort.Sort(Timelines(result))

	return result, nil
}

func (t *timeline) LastTimeline(query string) (Timeline, error) {
	list, err := t.Read(query)
	if err != nil {
		return Timeline{}, err
	}
	if len(list) == 0 {
		return Timeline{}, errors.New("404")
	}

	return list[len(list)-1], nil
}

func (t *timeline) ListTable() (string, error) {
	tables, err := t.db.ListTables()
	if err != nil {
		return "", errors.Wrap(err, "get tables error")
	}

	return strings.Join(tables, "\n"), nil
}

func (t *timeline) DropTable(table string) (string, error) {
	if table == "" {
		return "", errors.New("table name is required")
	}

	err := t.db.DropTable(table)

	if err != nil {
		return "", errors.Wrap(err, "drop table error")
	}

	return "success", nil
}
