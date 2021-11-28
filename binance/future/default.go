package future

import (
	"github.com/workfoxes/tripwire/binance"
)

type OrderRequestOption struct {
	symbol string
	side   binance.SideType
	//positionSide     *PositionSideType
	orderType   binance.OrderType
	timeInForce *binance.TimeInForceType
	quantity    string
	reduceOnly  *bool
	price       *string
	newClientOrderID *string
	stopPrice        *string
	//workingType      *WorkingType
	activationPrice  *string
	callbackRate     *string
	priceProtect     *bool
	newOrderRespType binance.NewOrderRespType
	closePosition    *bool
}

// Order define order info
type Order struct {
	Symbol           string                  `json:"symbol"`
	OrderID          int64                   `json:"orderId"`
	ClientOrderID    string                  `json:"clientOrderId"`
	Price            string                  `json:"price"`
	ReduceOnly       bool                    `json:"reduceOnly"`
	OrigQuantity     string                  `json:"origQty"`
	ExecutedQuantity string                  `json:"executedQty"`
	CumQuantity      string                  `json:"cumQty"`
	CumQuote         string                  `json:"cumQuote"`
	StopPrice        string                  `json:"stopPrice"`
	Time             int64                   `json:"time"`
	UpdateTime       int64                   `json:"updateTime"`
	ActivatePrice    string                  `json:"activatePrice"`
	PriceRate        string                  `json:"priceRate"`
	AvgPrice         string                  `json:"avgPrice"`
	OrigType         string                  `json:"origType"`
	PriceProtect     bool                    `json:"priceProtect"`
	ClosePosition bool                    `json:"closePosition"`
	Status        binance.OrderStatusType `json:"status"`
	TimeInForce   binance.TimeInForceType `json:"timeInForce"`
	Type          binance.OrderType       `json:"type"`
	Side          binance.SideType        `json:"side"`
	//WorkingType      WorkingType      `json:"workingType"`
	//PositionSide     PositionSideType `json:"positionSide"`
}
