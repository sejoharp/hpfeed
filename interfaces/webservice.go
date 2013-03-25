package interfaces

import (
	"hpfeed/helper"
	"hpfeed/usecases"
	"net/http"
	"strconv"
)

type WebserviceHandler struct {
	service     usecases.MessageInteractorInterface
	feedBuilder FeedBuilderInterface
	listenPort  int
	listenPath  string
}

func CreateNewWebserviceHandler(service usecases.MessageInteractorInterface, feedBuilder FeedBuilderInterface, listenPort int, listenPath string) *WebserviceHandler {
	return &WebserviceHandler{service: service, feedBuilder: feedBuilder, listenPort: listenPort, listenPath: listenPath}
}

func (this *WebserviceHandler) StartHttpserver() {
	http.HandleFunc("/"+this.listenPath, func(res http.ResponseWriter, req *http.Request) {
		this.feedHandler(res, req)
	})
	err := http.ListenAndServe(":"+strconv.Itoa(this.listenPort), nil)
	helper.HandleFatalError("starting http server failed:", err)
}

func (this *WebserviceHandler) feedHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/rss+xml")
	responseWriter.Write(this.generateResponse())
}

func (this *WebserviceHandler) generateResponse() []byte {
	allMessages := this.service.GetAllMessages()
	return this.feedBuilder.Generate(allMessages)
}
