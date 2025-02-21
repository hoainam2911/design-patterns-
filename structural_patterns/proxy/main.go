package main

import "fmt"


type server interface {
	handleRequest(string, string) (int, string)
}
type Application struct{}

func (a *Application) handleRequest(url, method string) (int, string) {

	if url == "/app/status" && method == "GET" {
		return 200, "Ok" 
	} else if url == "/create/user" && method == "POST" {
		return 201, "User Created" 
	}
	return 404, "Not Found" 
}

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{}, 
		maxAllowedRequest: 2,              
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed" 
	}
	return n.application.handleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	} else if n.rateLimiter[url] >= n.maxAllowedRequest {
		return false 
	} else {
		n.rateLimiter[url]++
	}
	return true
}


func main() {
	nginxServer := newNginxServer()

	appStatusURL := "/app/status"
	createUserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
}