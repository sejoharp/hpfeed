package main

import (
	"bitbucket.org/joscha/hpfeed/domain"
	"bitbucket.org/joscha/hpfeed/helper"
	"bitbucket.org/joscha/hpfeed/interfaces"
	"bitbucket.org/joscha/hpfeed/usecases"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := interfaces.CreateNewConfigurator().LoadConfig()
	newsRepo := domain.CreateCouchDbRepo(config.Dbhost, config.Dbport, config.Dbname, config.Dbuser, config.Dbpassword)
	interactor := usecases.CreateNewMessageInteractor(newsRepo)
	forumReader := interfaces.CreateNewForumReader(config.ForumUser, config.ForumPasswd)
	parser := interfaces.CreateNewParser()
	feedUpdaterBatch := interfaces.CreateNewFeedUpdater(config.Updateinterval, interactor, forumReader, parser)
	feedBuilder := interfaces.CreateNewFeedBuilder(config.Updateinterval)
	webservice := interfaces.CreateNewWebservice(interactor, feedBuilder, config.ListenPort, config.ListenPath)

	config.Log()
	feedUpdaterBatch.Start()
	helper.LogInfo("--> feed update batch started.")
	webservice.Start()
	helper.LogInfo("--> rssfeed online.")

	shutdownChannel := make(chan os.Signal)
	signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGTERM)

	<-shutdownChannel
	
	helper.LogInfo("--> shutting down feed.")
	webservice.Stop()
	helper.LogInfo("--> stopping webservice.")
	feedUpdaterBatch.Stop()
	helper.LogInfo("--> stopping update batch.")
	helper.LogInfo("--> rssfeed stopped.")
}
