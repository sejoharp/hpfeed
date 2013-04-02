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
	newsRepo := domain.CreateMongoDbNewsRepo(config.Dbhost, config.Dbname)
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
	helper.LogInfo("updateinterval: " + strconv.Itoa(config.Updateinterval))
	helper.LogInfo("listen port: " + strconv.Itoa(config.ListenPort))
	helper.LogInfo("listen path: " + config.ListenPath)
	helper.LogInfo("dbhost: " + config.Dbhost)
	helper.LogInfo("dbname: " + config.Dbname)
	helper.LogInfo("forum user: " + config.ForumUser)
}
