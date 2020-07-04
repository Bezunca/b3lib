package history

import "time"

// FixedPoint represents the actual value multiplied by 100
type FixedPoint int

// AssetInfo holds information of assets in B3
type AssetInfo struct {
	Year                   int        `json:"year" bson:"year"`
	TipReg                 int        `json:"-" bson:"-" `
	DataCollectionDate     time.Time  `json:"data" bson:"date"`
	BDICode                int        `json:"-" bson:"-"`
	Ticker                 string     `json:"ticker" bson:"ticker"`
	MarketType             int        `json:"market_type" bson:"market_type"`
	CompanyName            string     `json:"shot_company_name" bson:"shot_company_name"`
	SecurityType           string     `json:"security_type" bson:"security_type"`
	FutureMarketExpiration string     `json:" future_market_expiration" bson:" future_market_expiration"` // Need to check this against an example that actually has a value
	Currency               string     `json:"currency" bson:"currency"`
	PriceOpen              FixedPoint `json:"-" bson:"-"`
	PriceMax               FixedPoint `json:"-" bson:"-"`
	PriceMin               FixedPoint `json:"-" bson:"-"`
	PriceMean              FixedPoint `json:"-" bson:"-"`
	PriceClose             FixedPoint `json:"price_close" bson:"price_close"`
	PriceBid               FixedPoint `json:"-" bson:"-"`
	PriceAsk               FixedPoint `json:"-" bson:"-"`
	TotalTrades            int        `json:"-" bson:"-"`
	TotalQuantity          int        `json:"-" bson:"-"`
	TotalVolume            FixedPoint `json:"total_volume" bson:"total_volume"`
	PreExe                 FixedPoint `json:"execution_price" bson:"execution_price"` // Needs further investigation
	IndOpc                 int        `json:"-" bson:"-"`                             // Needs further investigation
	ExpirationDate         time.Time  `json:"expiration_date" bson:"expiration_date"`
	FatCot                 int        `json:"-" bson:"-"` // Needs further investigation
	PtoExe                 int        `json:"-" bson:"-"` // Needs further investigation
	ISINCode               string     `json:"-" bson:"-"`
	DistributionNumber     int        `json:"-" bson:"-"`
}
