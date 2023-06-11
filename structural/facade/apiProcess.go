package facade

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	url             = "http://localhost:3000"
	loginEndpoint   = "/auth?login=%s&pass=%s"
	sellEndpoint    = "/sell?token=%s&price=%f&amount=%f&pair=%s"
	buyEndpoint     = "/buy?token=%s&price=%f&amount=%f&pair=%s"
	balanceEndpoint = "/balance?token=%s&pair=%s"
	maxRetry        = 5
)

type ApiProcess struct {
	account *Account
	token   string
}

func NewApiProcess(account *Account) *ApiProcess {
	return &ApiProcess{
		account: account,
	}
}
func (a *ApiProcess) Login() error {
	requestURL := fmt.Sprintf("%s"+loginEndpoint, url, a.account.login, a.account.password)
	log.Println("login to host", requestURL)
	body, err := a.request(requestURL)
	a.token = body
	return err
}
func (a *ApiProcess) Balance(pair string) (balance float64, err error) {
	requestURL := fmt.Sprintf("%s"+balanceEndpoint, url, a.token, pair)
	body, err := a.request(requestURL)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(body, 64)
}
func (a *ApiProcess) Sell(amount, price float64, pair string) (info string, err error) {
	return a.action(amount, price, pair, sellEndpoint)
}

func (a *ApiProcess) Buy(amount, price float64, pair string) (info string, err error) {
	return a.action(amount, price, pair, buyEndpoint)
}

func (a *ApiProcess) action(amount, price float64, pair string, endpoint string) (info string, err error) {
	requestURL := fmt.Sprintf("%s"+endpoint, url, a.token, price, amount, pair)
	body, err := a.request(requestURL)
	return body, err
}
func (a *ApiProcess) request(requestURL string, retries ...int) (string, error) {
	retry := 0
	if len(retries) > 0 {
		retry = retries[0]
	}
	res, err := http.Get(requestURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != fiber.StatusOK {
		if err != nil {
			return "", err
		}
		retry++
		if retry >= maxRetry {
			log.Fatalln("login or password invalid")
		}
		time.Sleep(time.Duration(retry) * time.Second)
		return a.request(requestURL, retry)
	}
	body, err := io.ReadAll(res.Body)
	return string(body), err
}
