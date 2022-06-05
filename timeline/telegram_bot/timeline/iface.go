package timeline

type Interface interface {
	Add(tm Timeline) (string, error)
	Update(tm Timeline) error
	Read(query string) (Timelines, error)
	LastTimeline(query string) (Timeline, error)
	Delete(id string) (string, error)
	ListTable() (string, error)
	DropTable(table string) (string, error)
}
