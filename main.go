package main

import (
	"bitbucket.org/joscha/hpfeed/domain"
	"bitbucket.org/joscha/hpfeed/helper"
	"bitbucket.org/joscha/hpfeed/interfaces"
	"bitbucket.org/joscha/hpfeed/usecases"
	"strconv"
)

func main() {
	config := interfaces.CreateNewConfigurator("hpfeed.conf").LoadConfig()
	newsRepo := domain.CreateCouchDbRepo(config.Dbhost, config.Dbport, config.Dbname)
	service := usecases.CreateNewMessageInteractor(newsRepo)
	forumReader := interfaces.CreateNewForumReader(config.ForumUser, config.ForumPasswd)
	feedUpdaterBatch := interfaces.CreateNewFeedUpdater(config.Updateinterval, service, forumReader)
	feedBuilder := interfaces.CreateNewFeedBuilder(config.Updateinterval)
	webservice := interfaces.CreateNewWebserviceHandler(service, feedBuilder, config.ListenPort, config.ListenPath)

	logConfig(config)
	feedUpdaterBatch.StartFeedUpdateCycle()
	helper.LogInfo("feed update batch started.")
	helper.LogInfo("rssfeed online.")
	webservice.StartHttpserver()
}

func logConfig(config *interfaces.Config) {
	helper.LogInfo("config for updateinterval: " + strconv.Itoa(config.Updateinterval))
	helper.LogInfo("config for listen port: " + strconv.Itoa(config.ListenPort))
	helper.LogInfo("config for listen path: " + config.ListenPath)
	helper.LogInfo("config for dbhost: " + config.Dbhost)
	helper.LogInfo("config for dbname: " + config.Dbname)
	helper.LogInfo("config for dbport: " + config.Dbport)
	helper.LogInfo("config for forum user: " + config.ForumUser)
}
