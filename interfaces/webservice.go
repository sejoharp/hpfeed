package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"bitbucket.org/joscha/hpfeed/usecases"
	"net"
	"net/http"
	"strconv"
)

type Handler struct {
	*http.ServeMux
}

type Webservice struct {
	interactor  usecases.MessageInteractorInterface
	feedBuilder FeedBuilderInterface
	listenPort  int
	listenPath  string
	listener    *net.TCPListener
}

func CreateNewWebservice(
	interactor usecases.MessageInteractorInterface,
	feedBuilder FeedBuilderInterface,
	listenPort int,
	listenPath string) *Webservice {
	return &Webservice{
		interactor:  interactor,
		feedBuilder: feedBuilder,
		listenPort:  listenPort,
		listenPath:  listenPath}
}

func (this *Webservice) feedHandler(responseWriter http.ResponseWriter, request *http.Request) {
	allMessages := this.interactor.GetAllMessages()
	response := this.feedBuilder.Generate(allMessages)
	responseWriter.Header().Set("Content-Type", "application/rss+xml")
	responseWriter.Write(response)
}

func (this *Webservice) Start() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(this.listenPort))
	helper.HandleFatalError("error resolving webservice address.", err)
	this.listener, err = net.ListenTCP("tcp", tcpAddr)
	helper.HandleFatalError("error opening webservice port.", err)

	handler := &Handler{ServeMux: http.NewServeMux()}

	handler.ServeMux.HandleFunc(this.listenPath,
		func(res http.ResponseWriter, req *http.Request) {
			this.feedHandler(res, req)
		})

	go http.Serve(this.listener, handler)
}

func (this *Webservice) Stop() error {
	return this.listener.Close()
}
