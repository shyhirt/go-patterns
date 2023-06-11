package proxy

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type MemProxy struct {
	application    *App
	cache          map[string]CacheData
	timeExpiredSec int
	mut            sync.Mutex
}

func newMemProxy(timeExpiredSec int) *MemProxy {
	return &MemProxy{
		application:    &App{},
		cache:          make(map[string]CacheData),
		timeExpiredSec: timeExpiredSec,
		mut:            sync.Mutex{},
	}
}

func (m *MemProxy) handleRequest(url, method string) (int, string) {
	key := fmt.Sprintf("%s-%s", url, method)
	data, exist := m.getCachedData(key)
	if !exist {
		data.Code, data.Data = m.application.handleRequest(url, method)
		data.Expired = time.Now().Add(time.Second * time.Duration(m.timeExpiredSec))
		m.mut.Lock()
		m.cache[key] = data
		m.mut.Unlock()
	} else {
		log.Printf("%s from cache\r\n", url)
	}
	return data.Code, data.Data
}

func (m *MemProxy) getCachedData(key string) (data CacheData, exist bool) {
	m.mut.Lock()
	data, exist = m.cache[key]
	if exist && time.Now().After(data.Expired) {
		delete(m.cache, key)
		exist = false
	}
	m.mut.Unlock()
	return
}
