package facade

import "time"

type Deal struct {
	DialType int
	Price    float64
	Size     float64
	Balance  float64
	Date     time.Time
}
type TradeHistory struct {
	Deals []Deal
}

func NewTradeHistory() *TradeHistory {
	return &TradeHistory{
		Deals: []Deal{},
	}
}
func (t *TradeHistory) SetDeal(d Deal) {
	d.Date = time.Now()
	t.Deals = append(t.Deals, d)
}

func (t *TradeHistory) GetAll() []Deal {
	return t.Deals
}
