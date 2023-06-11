package facade

import "errors"

const (
	typeSell = iota
	typeBuy
)

var errTrade = errors.New("unknown action error")

type CryptoExchange struct {
	apiProcess   *ApiProcess
	wallet       *Wallet
	tradeHistory *TradeHistory
	pair         string
}

func NewCryptoExchange(login, pass, pair string) *CryptoExchange {
	account := NewAccount(login, pass)
	return &CryptoExchange{
		tradeHistory: NewTradeHistory(),
		wallet:       NewWallet(),
		apiProcess:   NewApiProcess(account),
		pair:         pair,
	}
}
func (c *CryptoExchange) GetTypeString(typ int) string {
	if typ == typeSell {
		return "sell"
	}
	if typ == typeBuy {
		return "buy"
	}
	return "n/a"
}
func (c *CryptoExchange) Buy(amount, price float64) error {
	return c.trade(amount, price, typeBuy)
}

func (c *CryptoExchange) Sell(amount, price float64) error {
	return c.trade(amount, price, typeSell)
}
func (c *CryptoExchange) Balance() error {
	balance, err := c.apiProcess.Balance(c.pair)
	if err != nil {
		return err
	}
	c.wallet.setBalance(balance)
	return nil
}
func (c *CryptoExchange) handleResponse(resp string, d Deal) error {
	_ = resp
	// Handle response from exchange and run some actions
	//....
	//....
	// If no errors put to trade history
	c.tradeHistory.SetDeal(d)
	return nil
}

func (c *CryptoExchange) trade(amount, price float64, tradeAction int) error {
	if c.apiProcess.token == "" {
		err := c.apiProcess.Login()
		if err != nil {
			return err
		}
	}
	err := c.Balance()
	if err != nil {
		return err
	}
	err = c.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	var resp string
	err = errTrade
	if tradeAction == typeBuy {
		resp, err = c.apiProcess.Buy(amount, price, c.pair)
	}
	if tradeAction == typeSell {
		resp, err = c.apiProcess.Sell(amount, price, c.pair)
	}
	if err != nil {
		return err
	}
	return c.handleResponse(resp, Deal{
		DialType: tradeAction,
		Price:    price,
		Size:     amount,
		Balance:  c.wallet.getBalance(),
	})
}
