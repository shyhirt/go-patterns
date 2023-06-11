package facade

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewCryptoExchange(t *testing.T) {
	go Run()
	exchange := NewCryptoExchange("login", "pass", "btc-usdt")
	err := exchange.apiProcess.Login()
	exchange.Balance()
	log.Printf("Balance: %.2f", exchange.wallet.getBalance())
	assert.NoError(t, err)
	err = exchange.Buy(100, 5)
	assert.NoError(t, err)
	err = exchange.Sell(100, 6)
	assert.NoError(t, err)
	history := exchange.tradeHistory.GetAll()
	tw := table.NewWriter()
	log.Println(history)

	for rowIdx := 0; rowIdx < len(history); rowIdx++ {
		row := make(table.Row, 6)
		row[0] = rowIdx
		row[1] = exchange.GetTypeString(history[rowIdx].DialType)
		row[2] = fmt.Sprintf("%2.f", history[rowIdx].Price)
		row[3] = fmt.Sprintf("%2.f", history[rowIdx].Size)
		row[4] = fmt.Sprintf("%2.f", history[rowIdx].Balance)
		row[5] = fmt.Sprintf("%v", history[rowIdx].Date)

		tw.AppendRow(row)
	}
	tw.AppendHeader(table.Row{"#", "BUY/SELL", "PRICE", "SIZE", "BALANCE", "DATE"})
	tw.SetStyle(table.StyleLight)
	log.Println("")
	fmt.Println(tw.Render())

}
