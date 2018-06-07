package core

import (
	"fmt"

	"net/url"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config interface {
	Configurator() *viper.Viper
	DSN(instance string) string
	Database(instance string) *viper.Viper
	DatabaseInstances() []string
	Addr(instance string) string
	Server(instance string) *viper.Viper
	ServerInstances() []string
	Core() *viper.Viper
}

type config struct {
	conf              viper.Viper
	databaseInstances []string
	serverInstances   []string
}

func NewConfig() Config {
	c := viper.New()
	c.AddConfigPath(".")
	c.AddConfigPath("config")
	c.AddConfigPath("$HOME/config")
	c.SetConfigName(".env")
	c.SetConfigType("toml")

	setDefaults(c)

	err := c.ReadInConfig()
	if err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
	}

	// Watching and re-reading config files
	c.WatchConfig()
	c.OnConfigChange(func(e fsnotify.Event) {
		setDefaults(c)
	})

	databaseInstances := getInstances(c, "database")
	serverInstances := getInstances(c, "server")

	return &config{*c, databaseInstances, serverInstances}
}

func getInstances(v *viper.Viper, group string) []string {
	var instances []string
	for key := range v.GetStringMapStringSlice(group) {
		instances = append(instances, key)
	}
	return instances
}

func setDefaults(v *viper.Viper) {
	defaults := map[string]interface{}{
		"core": map[string]interface{}{
			"secure": map[string]interface{}{
				"key": "My secret",
			},
			"log": map[string]string{
				"level":   "debug",
				"storage": "storage/logs",
			},
			"request": map[string]interface{}{
				"log": map[string]interface{}{
					"level":   "debug",
					"storage": "storage/logs",
				},
			},
		},
		"server": map[string]interface{}{
			"main": map[string]interface{}{
				"host":  "127.0.0.1",
				"port":  8080,
				"name":  "mvc-app",
				"owner": "owner",
			},
		},
		"database": map[string]interface{}{},
	}

	for key, value := range defaults {
		v.SetDefault(key, value)
	}
}

func (c *config) Configurator() *viper.Viper {
	return &c.conf
}

func (c *config) DatabaseInstances() []string {
	return c.databaseInstances
}

func (c *config) ServerInstances() []string {
	return c.serverInstances
}

func (c *config) DSN(instance string) (dsn string) {
	section := "database." + instance
	if !c.conf.IsSet(section) {
		return ""
	}
	s := c.conf.Sub(section)

	if s.GetString("driver") == "mssql" {
		defaultParameters := map[string]interface{}{
			"app name":        "go",
			"keepAlive":       30,
			"failoverpartner": "",
			"failoverport":    1433,
			"packet size":     4096,
			"log":             0,
			"TrustServerCertificate": true,
			"certificate":            "",
			"hostNameInCertificate":  "",
			"ServerSPN":              "",
			"Workstation ID":         "",
			"ApplicationIntent":      "",
		}

		query := url.Values{}

		p := s.Sub("parameters")

		if param := s.GetString("database"); param != "" {
			query.Add("database", param)
		}

		for key, value := range defaultParameters {
			if p.Get(key) == nil {
				p.SetDefault(key, value)
			}
			query.Add(key, p.GetString(key))
		}

		u := &url.URL{
			Scheme:   "sqlserver",
			User:     url.UserPassword(s.GetString("username"), s.GetString("password")),
			Host:     fmt.Sprintf("%s:%d", s.GetString("host"), s.GetInt("port")),
			RawQuery: query.Encode(),
		}

		dsn = u.String()
	}

	if s.GetString("driver") == "postgres" {
		dsn = "postgresql://" + s.GetString("username") +
			":" + s.GetString("password") + "@" +
			s.GetString("host") + ":" + s.GetString("port") + "/" +
			s.GetString("database") + "?sslmode=" + s.GetString("sslmode")
	}

	if s.GetString("driver") == "sqlite3" {
		dsn = s.GetString("database")
	}

	return
}

func (c *config) Database(instance string) *viper.Viper {
	section := "database." + instance
	if !c.conf.IsSet(section) {
		return nil
	}
	s := c.conf.Sub(section)

	return s
}

func (c *config) Addr(instance string) string {
	section := "server." + instance
	if !c.conf.IsSet(section) {
		return ""
	}
	s := c.conf.Sub(section)

	port := s.GetString("port")

	// heroku
	if port == "0" || port == "" {
		s.BindEnv("PORT")
		port = s.GetString("PORT")
	}

	return s.GetString("host") + ":" + port
}

func (c *config) Server(instance string) *viper.Viper {
	section := "server." + instance
	if !c.conf.IsSet(section) {
		return nil
	}
	s := c.conf.Sub(section)

	return s
}

func (c *config) Core() *viper.Viper {
	section := "core"
	if !c.conf.IsSet(section) {
		return nil
	}
	s := c.conf.Sub(section)

	return s
}
