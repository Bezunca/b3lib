package history

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseContentLine(rawLine string, year int) (*Asset, *Price, error) {
	tipReg, err := strconv.ParseInt(rawLine[:2], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	date, err := time.Parse("20060102 -0700", fmt.Sprintf("%s -0300", rawLine[2:2+8]))
	if err != nil {
		return nil, nil, err
	}
	bdiCode, err := strconv.ParseInt(rawLine[10:10+2], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	marketType, err := strconv.ParseInt(rawLine[24:24+3], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	priceOpen, err := strconv.ParseInt(rawLine[56:56+13], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	priceMax, err := strconv.ParseInt(rawLine[69:69+13], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	priceMin, err := strconv.ParseInt(rawLine[82:82+13], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	priceMean, err := strconv.ParseInt(rawLine[95:95+13], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	priceClose, err := strconv.ParseInt(rawLine[108:108+13], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	priceBid, err := strconv.ParseInt(rawLine[121:121+13], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	priceAsk, err := strconv.ParseInt(rawLine[134:134+13], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	totalTrades, err := strconv.ParseInt(rawLine[147:147+5], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	totalQuantity, err := strconv.ParseInt(rawLine[152:152+18], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	totalVolume, err := strconv.ParseInt(rawLine[170:170+18], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	preExe, err := strconv.ParseInt(rawLine[188:188+13], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	indOpc, err := strconv.ParseInt(rawLine[201:201+1], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	expirationDate, err := time.Parse("20060102 -0700", fmt.Sprintf("%s -0300", rawLine[202:202+8]))
	if err != nil {
		return nil, nil, err
	}
	fatCot, err := strconv.ParseInt(rawLine[210:210+7], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	ptoExe, err := strconv.ParseInt(rawLine[210:210+7], 10, 64)
	if err != nil {
		return nil, nil, err
	}
	distributionNumber, err := strconv.ParseInt(rawLine[242:242+3], 10, 64)
	if err != nil {
		return nil, nil, err
	}

	allData := AssetInfo{
		Year:                   year,
		TipReg:                 int(tipReg),
		DataCollectionDate:     date,
		BDICode:                int(bdiCode),
		Ticker:                 strings.TrimSpace(rawLine[12 : 12+12]),
		MarketType:             int(marketType),
		CompanyName:            strings.TrimSpace(rawLine[27 : 27+12]),
		SecurityType:           strings.TrimSpace(rawLine[39 : 39+10]),
		FutureMarketExpiration: strings.TrimSpace(rawLine[49 : 49+3]),
		Currency:               strings.TrimSpace(rawLine[52 : 52+4]),
		PriceOpen:              FixedPoint(priceOpen),
		PriceMax:               FixedPoint(priceMax),
		PriceMin:               FixedPoint(priceMin),
		PriceMean:              FixedPoint(priceMean),
		PriceClose:             FixedPoint(priceClose),
		PriceBid:               FixedPoint(priceBid),
		PriceAsk:               FixedPoint(priceAsk),
		TotalTrades:            int(totalTrades),
		TotalQuantity:          int(totalQuantity),
		TotalVolume:            FixedPoint(totalVolume),
		PreExe:                 FixedPoint(preExe),
		IndOpc:                 int(indOpc),
		ExpirationDate:         expirationDate,
		FatCot:                 int(fatCot),
		PtoExe:                 int(ptoExe),
		ISINCode:               strings.TrimSpace(rawLine[230 : 230+12]),
		DistributionNumber:     int(distributionNumber),
	}

	return &Asset{
		Ticker:       allData.Ticker,
		CompanyName:  allData.CompanyName,
		MarketType:   allData.MarketType,
		SecurityType: allData.SecurityType,
		Currency:     allData.Currency,
	}, &Price{
		Year:                   allData.Year,
		DataCollectionDate:     allData.DataCollectionDate,
		Ticker:                 allData.Ticker,
		FutureMarketExpiration: allData.FutureMarketExpiration,
		PriceClose:             allData.PriceClose,
		TotalVolume:            allData.TotalVolume,
		PreExe:                 allData.PreExe,
		ExpirationDate:         allData.ExpirationDate,
	},
	nil
}

func parseHistoricData(rawData []string, year int) (map[string]*Asset, []Price, error) {
	priceList := make([]Price, len(rawData)-3)
	assetMap := map[string]*Asset{}
	for i, rawLine := range rawData[1 : len(rawData)-2] {
		assetSlice, priceSlice, err := parseContentLine(rawLine, year)
		if err != nil {
			return nil, nil, err
		}
		priceList[i] = *priceSlice
		assetMap[assetSlice.Ticker] = assetSlice
	}
	return assetMap, priceList, nil
}

func parseHistoricDataFromBytes(data []byte, year int) (map[string]*Asset, []Price, error) {
	return parseHistoricData(strings.Split(string(data), "\n"), year)
}
