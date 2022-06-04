package timeline

import (
	"encoding/json"
	"fmt"
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
	Content string
}

func NewTimeline(db dbIface.DB, table string) *timeline {
	return &timeline{
		db:    db,
		table: table,
	}
}

func (t *timeline) Add(record Timeline) (string, error) {
	if record.Content == "" {
		return "", errors.New("content is required")
	}

	id, err := t.db.Create(t.table, dbIface.Record{
		"content":   record.Content,
		"timestamp": time.Now().Unix(),
	})

	if err != nil {
		return "", errors.Wrap(err, "create record error")
	}

	count, err := t.db.Count(t.table)
	if err != nil {
		return id, errors.Wrap(err, "get record count error")
	}

	return fmt.Sprintf("id: %s\ncontent: %s\ncount: %d", id, record.Content, count), nil
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

func (t *timeline) Read(query string) (string, error) {
	records, err := t.db.Read(t.table, query)
	if err != nil {
		return "", errors.Wrap(err, "get records error")
	}

	result := []string{}

	for _, v := range records {
		b, _ := json.MarshalIndent(v, "", "  ")
		result = append(result, string(b))
	}

	return strings.Join(result, "\n"), nil
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
