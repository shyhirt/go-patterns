package proxy

import (
	"log"
	"testing"
	"time"
)

func Test_proxy(t *testing.T) {
	cacheQuery := newMemProxy(3)
	code, body := cacheQuery.handleRequest("/api/get", "GET")
	log.Printf("%d %s", code, body)
	//It's cached query
	code, body = cacheQuery.handleRequest("/api/get", "GET")
	log.Printf("%d %s", code, body)
	//Cleaned cache after 3 second
	<-time.After(time.Second * 4)
	code, body = cacheQuery.handleRequest("/api/get", "GET")
	log.Printf("%d %s", code, body)
}
