package main

import (
	"bitbucket.org/joscha/hpfeed/domain"
	"bitbucket.org/joscha/hpfeed/helper"
	"bitbucket.org/joscha/hpfeed/interfaces"
	"bitbucket.org/joscha/hpfeed/usecases"
)

func main() {
	config := interfaces.CreateNewConfigurator("hpfeed.conf").LoadConfig()
	newsRepo := domain.CreateCouchDbRepo(config.Dbhost, config.Dbport, config.Dbname)
	service := usecases.CreateNewMessageInteractor(newsRepo)
	forumReader := interfaces.CreateNewForumReader(config.ForumUser, config.ForumPasswd)
	feedUpdaterBatch := interfaces.CreateNewFeedUpdater(config.Updateinterval, service, forumReader)
	feedBuilder := interfaces.CreateNewFeedBuilder(config.Updateinterval)
	webservice := interfaces.CreateNewWebserviceHandler(service, feedBuilder, config.ListenPort, config.ListenPath)

	config.Log()
	feedUpdaterBatch.StartFeedUpdateCycle()
	helper.LogInfo("feed update batch started.")
	helper.LogInfo("rssfeed online.")
	webservice.StartHttpserver()
}


