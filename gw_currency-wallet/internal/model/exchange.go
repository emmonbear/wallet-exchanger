package model

type ExchangeRates struct {
	USD float64 `db:"usd" json:"USD"`
	RUB float64 `db:"rub" json:"RUB"`
	EUR float64 `db:"eur" json:"EUR"`
}

type ExchangeRequest struct {
	FromCurrency string  `json:"from_currency" binding:"required"`
	ToCurrency   string  `json:"to_currency" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
}
