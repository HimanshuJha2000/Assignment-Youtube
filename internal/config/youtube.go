package config

type YoutubeConfig struct {
	Endpoint   string `toml:"endpoint"`
	Query      string `toml:"query"`
	TickerTime int64  `toml:"ticker_time"`
	MaxResults int    `toml:"max_results_count"`
}
