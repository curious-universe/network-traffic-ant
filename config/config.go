package config

import (
	"errors"
	"fmt"
	"github.com/curious-universe/network-traffic-ant/nerror"
	"github.com/curious-universe/network-traffic-ant/zaplog"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync/atomic"
)

var (
	globalConf atomic.Value
)

var defaultConf = &Config{
	Log{
		zaplog.Config{Level: "debug", Development: true},
		zaplog.FileLogConfig{Filename: ""},
	},
}

type Log struct {
	Config zaplog.Config        `yaml:"config" toml:"config" json:"config"`
	File   zaplog.FileLogConfig `yaml:"file" toml:"file" json:"file"`
}

type Config struct {
	Log `yaml:"log" toml:"log" json:"log"`
}

func init() {
	StoreGlobalConfig(defaultConf)
}

// GetGlobalConfig returns the global configuration for this server.
// It should store configuration from command line and configuration file.
// Other parts of the system can read the global configuration use this function.
func GetGlobalConfig() *Config {
	return globalConf.Load().(*Config)
}

// StoreGlobalConfig stores a new config to the globalConf. It mostly uses in the test to avoid some data races.
func StoreGlobalConfig(config *Config) {
	globalConf.Store(config)
}

// Load loads config options from a toml file.
func (c *Config) Load(confFile string) error {
	yml, err := ioutil.ReadFile(confFile)
	nerror.MustNil(err)
	err = yaml.Unmarshal(yml, &c)

	return err
}

// Valid checks if this config is valid.
func (c *Config) Valid() error {
	// test log level
	l := zap.NewAtomicLevel()
	return l.UnmarshalText([]byte(c.Log.Config.Level))
}

func InitializeConfig(confPath string, configCheck bool) {
	cfg := GetGlobalConfig()
	var err error
	if confPath != "" {
		err = cfg.Load(confPath)
		nerror.MustNil(err)
	} else {
		// configCheck should have the config file specified.
		if configCheck {
			_, _ = fmt.Fprintln(os.Stderr, "config check failed", errors.New("no config file specified for config-check"))
			os.Exit(1)
		}
	}

	if err := cfg.Valid(); err != nil {
		if !filepath.IsAbs(confPath) {
			if tmp, err := filepath.Abs(confPath); err == nil {
				confPath = tmp
			}
		}
		_, _ = fmt.Fprintln(os.Stderr, "load config file:", confPath)
		_, _ = fmt.Fprintln(os.Stderr, "invalid config", err)
		os.Exit(1)
	}
	if configCheck {
		fmt.Println("config check successful")
		os.Exit(0)
	}
	StoreGlobalConfig(cfg)
}
