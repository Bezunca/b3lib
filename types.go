package b3lib

import (
	"time"
)

type (
	// FetchedPrice struct to group ticker and it's current price
	FetchedPrice struct {
		Ticker   string `json:"ticker"`
		IntPrice int    `json:"int_price"`
	}

	b3PriceColumns struct {
		Name            string `json:"name"`
		FriendlyName    string `json:"friendlyName"`
		FriendlyNamePt  string `json:"friendlyNamePt"`
		FriendlyNameEn  string `json:"friendlyNameEn"`
		Type            int    `json:"type"`
		Format          string `json:"format"`
		ColumnAlignment int    `json:"columnAlignment"`
		ValueAlignment  int    `json:"valueAlignment"`
	}

	b3PriceResponse struct {
		Name         string           `json:"name"`
		FriendlyName string           `json:"friendlyName"`
		Columns      []b3PriceColumns `json:"columns"`
		Values       [][6]interface{} `json:"values"`
		// Inner slices follows columns order
	}

	tickerEntry struct {
		Price     FetchedPrice
		Timestamp time.Time
	}
)
