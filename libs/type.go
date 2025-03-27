package libs

// Struct matching the TOML structure
type App struct {
	Name    string `toml:"name" yaml:"name"`
	Version string `toml:"version" yaml:"version"`
	Debug   bool   `toml:"debug" yaml:"debug"`
}

type Config struct {
	App App `toml:"app" yaml:"app"`
}

// Resty
// Struct untuk menyimpan data response
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// CSV
type Client struct {
	ID   int    `csv:"id"`
	Name string `csv:"name"`
	Age  int    `csv:"age"`
}
