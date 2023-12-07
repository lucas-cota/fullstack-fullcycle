package dto

type TradeInput struct {
	OrderID       string `json:"oder_id"`
	InvestorID    string `json:"investor_id"`
	AssetID       string `json:"asset_id"`
	CurrentShares int    `json:"current_shares"`
	Shares        int    `json:"shares"`
	Price         float64    `json:"price"`
	OrderType     string `json:"order_type"`
}

type OrderOutPut struct {
	OrderID           string               `json:"oder_id"`
	InvestorID        string               `json:"investor_id"`
	AssetID           string               `json:"asset_id"`
	OrderType         int                  `json:"order_type"`
	Status            int                  `json:"status"`
	Partial           int                  `json:"partial"`
	Shares            string               `json:"shares"`
	TransactionOutput []*TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	TransactionID string `json:"transaction_id"`
	BuyerID       string `json:"buyer_id"`
	SellerID      string `json:"seller_id"`
	AssetID       string `json:"asset_id"`
	Price         float64    `json:"price"`
	Shares        string `json:"shares"`
}
