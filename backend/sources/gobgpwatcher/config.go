package gobgpwatcher

type Config struct {
	Id   int
	Name string

	Api string `ini:"api"`
}
