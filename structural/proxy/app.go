package proxy

import "fmt"

type App struct {
}

func (a *App) handleRequest(url, method string) (int, string) {
	//
	return 200, fmt.Sprintf("%s data from %s", method, url)
}
