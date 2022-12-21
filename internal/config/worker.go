package config

type worker struct {
	Name       string `toml:"app_name"`
	ListenPort int    `toml:"listen_port"`
	ListenIP   string `toml:"listen_ip"`
}
