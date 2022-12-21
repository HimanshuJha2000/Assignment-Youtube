package config

// application : struct to hold application level configs
type application struct {
	Name       string `toml:"app_name"`
	ListenPort int    `toml:"listen_port"`
	ListenIP   string `toml:"listen_ip"`
}
