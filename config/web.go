package config

type WebConfig struct {
	HttpPort  string `toml:"HttpPort"`
	HttpsPort string `toml:"HttpsPort"`
	CertPath  string `toml:"CertPath"`
	KeyPath   string `toml:"KeyPath"`
	Host      string `toml:"Host"`
}

func defaultWebConfig() WebConfig {
	return WebConfig{
		HttpPort:  "80",
		HttpsPort: "443",
		CertPath:  "server.crt",
		KeyPath:   "server.key",
		Host:      "binacs.cn",
	}
}
