package interfaces

import (
	"code.google.com/p/goconf/conf"
	"hpfeed/helper"
)

type Config struct {
	Updateinterval int
	ListenPort     int
	ListenPath     string
	Dbname         string
	Dbhost         string
	ForumPasswd    string
	ForumUser      string
}

type Configurator struct {
	configFilename string
}

func CreateNewConfigurator(configFilename string) *Configurator {
	return &Configurator{configFilename: configFilename}
}

func (this *Configurator) LoadConfig() *Config {
	config := Config{}
	configFile, err := conf.ReadConfigFile(this.configFilename)
	helper.HandleFatalError("loading config file failed:", err)

	config.Updateinterval, err = configFile.GetInt("", "updateinterval")
	helper.HandleFatalError("updateinterval", err)
	config.ListenPort, err = configFile.GetInt("", "listenPort")
	helper.HandleFatalError("listenPort", err)
	config.ListenPath, err = configFile.GetString("", "listenPath")
	helper.HandleFatalError("listenPath", err)
	config.Dbhost, err = configFile.GetString("", "dbhost")
	helper.HandleFatalError("dbhost", err)
	config.Dbname, err = configFile.GetString("", "dbname")
	helper.HandleFatalError("dbname", err)
	config.ForumUser, err = configFile.GetString("", "forumUser")
	helper.HandleFatalError("forumUser", err)
	config.ForumPasswd, err = configFile.GetString("", "forumPasswd")
	helper.HandleFatalError("forumPasswd", err)
	return &config
}
