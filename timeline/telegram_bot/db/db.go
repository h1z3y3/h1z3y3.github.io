package db

type Record map[string]interface{}

type DB interface {
	ListTables() ([]string, error)
	DropTable(table string) error
	Create(table string, record Record) (string, error)
	Delete(table string, id string) error
	Count(table string) (uint64, error)
	Read(table, query string) ([]Record, error)
}
