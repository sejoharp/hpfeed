package interfaces

import (
	"bitbucket.org/joscha/hpfeed/helper"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

type ForumReader struct {
	forumUser   string
	forumPasswd string
}

const forumPrefix = "http://kickern-hamburg.de/phpBB2"

func CreateNewForumReader(forumUser string, forumPasswd string) *ForumReader {
	return &ForumReader{forumUser: forumUser, forumPasswd: forumPasswd}
}

func (this *ForumReader) GetData() []byte {
	client := &http.Client{}
	client.Jar = createJar()

	this.login(client)
	rawData := this.getHTMLData(client)
	this.logout(client)
	return rawData
}

func (this *ForumReader) login(client *http.Client) {
	params := url.Values{"username": []string{this.forumUser}, "password": []string{this.forumPasswd}, "login": {"anmelden"}}
	resp, err := client.PostForm(forumPrefix+"/login.php", params)
	helper.HandleFatalError("login error: ", err)
	resp.Body.Close()
}

func (this *ForumReader) getHTMLData(client *http.Client) []byte {
	resp, err := client.Get(forumPrefix + "/viewforum.php?f=15")
	helper.HandleFatalError("reading data error: ", err)
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body
}

func (this *ForumReader) logout(client *http.Client) {
	sessionId := getSessionId(client.Jar)
	logoutUrl := forumPrefix + "/login.php?logout=true&sid=" + sessionId
	resp, err := client.Get(logoutUrl)
	helper.HandleFatalError("logout error: ", err)
	resp.Body.Close()
}

func (this *ForumReader) IsAvailable() bool {
	return isWebsiteAvailable("kickern-hamburg.de")
}

func isWebsiteAvailable(name string) bool {
	timeout := time.Duration(5) * time.Second
	conn, _ := net.DialTimeout("tcp", name+":80", timeout)
	if conn != nil {
		conn.Close()
		return true
	}
	return false
}

func createJar() *Jar {
	storage := make(map[string]*http.Cookie)
	jar := new(Jar)
	jar.storage = storage
	return jar
}

type Jar struct {
	storage map[string]*http.Cookie
}

func (this *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		this.storage[cookie.Name] = cookie
	}
}

func (this *Jar) Cookies(u *url.URL) []*http.Cookie {
	cookies := []*http.Cookie{}
	for _, cookie := range this.storage {
		cookies = append(cookies, cookie)
	}
	return cookies
}

func getSessionId(jar http.CookieJar) string {
	key := "phpbb2mysql_sid"
	url, _ := url.Parse("egal")
	for _, cookie := range jar.Cookies(url) {
		if cookie.Name == key {
			return cookie.Value
		}
	}
	panic("no sessionid found")
}
