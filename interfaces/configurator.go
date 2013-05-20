package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"code.google.com/p/goconf/conf"
	"flag"
	"strconv"
)

type Config struct {
	Updateinterval int
	ListenPort     int
	ListenPath     string
	Dbname         string
	Dbhost         string
	Dbport         string
	Dbuser         string
	Dbpassword     string
	ForumPasswd    string
	ForumUser      string
}

type Configurator struct {
	configFilename string
}

func (this *Config) Log() {
	helper.LogInfo("config for updateinterval: " + strconv.Itoa(this.Updateinterval))
	helper.LogInfo("config for listen port: " + strconv.Itoa(this.ListenPort))
	helper.LogInfo("config for listen path: " + this.ListenPath)
	helper.LogInfo("config for dbhost: " + this.Dbhost)
	helper.LogInfo("config for dbname: " + this.Dbname)
	helper.LogInfo("config for dbport: " + this.Dbport)
	helper.LogInfo("config for forum user: " + this.ForumUser)
}

func CreateNewConfigurator() *Configurator {
	return &Configurator{}
}

func (this *Configurator) LoadConfig() *Config {
	var configFileName string
	flag.StringVar(&configFileName, "config", "hpfeed.conf", "path to config file")
	flag.Parse()

	config := Config{}
	configFile, err := conf.ReadConfigFile(configFileName)
	helper.HandleFatalError("loading config file failed (-config= forgotten):", err)

	config.Updateinterval, err = configFile.GetInt("", "updateinterval")
	helper.HandleFatalError("updateinterval", err)
	config.ListenPort, err = configFile.GetInt("", "listenPort")
	helper.HandleFatalError("listenPort", err)
	config.ListenPath, err = configFile.GetString("", "listenPath")
	helper.HandleFatalError("listenPath", err)
	config.Dbhost, err = configFile.GetString("", "dbhost")
	helper.HandleFatalError("dbhost", err)
	config.Dbport, err = configFile.GetString("", "dbport")
	helper.HandleFatalError("dbport", err)
	config.Dbname, err = configFile.GetString("", "dbname")
	helper.HandleFatalError("dbname", err)
	config.Dbuser, err = configFile.GetString("", "dbuser")
	helper.HandleFatalError("dbuser", err)
	config.Dbpassword, err = configFile.GetString("", "dbpassword")
	helper.HandleFatalError("dbpassword", err)
	config.ForumUser, err = configFile.GetString("", "forumUser")
	helper.HandleFatalError("forumUser", err)
	config.ForumPasswd, err = configFile.GetString("", "forumPasswd")
	helper.HandleFatalError("forumPasswd", err)
	return &config
}
