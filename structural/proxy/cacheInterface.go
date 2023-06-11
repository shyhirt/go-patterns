package proxy

import "time"

type CacheData struct {
	Code    int
	Data    string
	Expired time.Time
}
type CacheInterface interface {
	handleRequest(string, string) CacheData
}
