package config

import (
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	WorkSpace   string      `toml:"workspace"`
	File        string      `toml:"configfile"`
	Mode        string      `toml:"mode"`
	WebConfig   WebConfig   `toml:"WebConfig"`
	GRPCConfig  GRPCConfig  `toml:"GRPCConfig"`
	LogConfig   LogConfig   `toml:"LogConfig"`
	RedisConfig RedisConfig `toml:"RedisConfig"`
	MysqlConfig MysqlConfig `toml:"MysqlConfig"`
	//Redis
	//Mysql
	rwmtx *sync.RWMutex
}

func defaultConfig() Config {
	ret := Config{
		WorkSpace:   ".",
		File:        "./config.toml",
		Mode:        "all",
		WebConfig:   defaultWebConfig(),
		GRPCConfig:  defaultGRPCConfig(),
		LogConfig:   defaultLogConfig(),
		RedisConfig: defaultRedisConfig(),
		MysqlConfig: defaultMysqlConfig(),
		rwmtx:       &sync.RWMutex{},
	}
	return ret
}

func LoadFromFile(configFile string) (*Config, error) {
	cfg := defaultConfig()
	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		return &cfg, err
	}
	cfg.File = configFile
	return &cfg, nil
}

func (cfg *Config) Reload() error {
	rwmtx := cfg.rwmtx
	rwmtx.Lock()
	defer rwmtx.Unlock()
	configFile := cfg.File
	newConfig, err := LoadFromFile(configFile)
	if err != nil {
		return err
	}
	*cfg = *newConfig
	cfg.File = configFile
	cfg.rwmtx = rwmtx
	return nil
}
