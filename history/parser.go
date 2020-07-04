package history

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseContentLine(rawLine string, year int) (*AssetInfo, error) {
	tipReg, err := strconv.ParseInt(rawLine[:2], 10, 64)
	if err != nil {
		return nil, err
	}
	date, err := time.Parse("20060102 -0700", fmt.Sprintf("%s -0300", rawLine[2:2+8]))
	if err != nil {
		return nil, err
	}
	bdiCode, err := strconv.ParseInt(rawLine[10:10+2], 10, 64)
	if err != nil {
		return nil, err
	}
	marketType, err := strconv.ParseInt(rawLine[24:24+3], 10, 64)
	if err != nil {
		return nil, err
	}
	priceOpen, err := strconv.ParseInt(rawLine[56:56+13], 10, 64)
	if err != nil {
		return nil, err
	}
	priceMax, err := strconv.ParseInt(rawLine[69:69+13], 10, 64)
	if err != nil {
		return nil, err
	}
	priceMin, err := strconv.ParseInt(rawLine[82:82+13], 10, 64)
	if err != nil {
		return nil, err
	}
	priceMean, err := strconv.ParseInt(rawLine[95:95+13], 10, 64)
	if err != nil {
		return nil, err
	}
	priceClose, err := strconv.ParseInt(rawLine[108:108+13], 10, 64)
	if err != nil {
		return nil, err
	}
	priceBid, err := strconv.ParseInt(rawLine[121:121+13], 10, 64)
	if err != nil {
		return nil, err
	}
	priceAsk, err := strconv.ParseInt(rawLine[134:134+13], 10, 64)
	if err != nil {
		return nil, err
	}
	totalTrades, err := strconv.ParseInt(rawLine[147:147+5], 10, 64)
	if err != nil {
		return nil, err
	}
	totalQuantity, err := strconv.ParseInt(rawLine[152:152+18], 10, 64)
	if err != nil {
		return nil, err
	}
	totalVolume, err := strconv.ParseInt(rawLine[170:170+18], 10, 64)
	if err != nil {
		return nil, err
	}
	preExe, err := strconv.ParseInt(rawLine[188:188+13], 10, 64)
	if err != nil {
		return nil, err
	}
	indOpc, err := strconv.ParseInt(rawLine[201:201+1], 10, 64)
	if err != nil {
		return nil, err
	}
	expirationDate, err := time.Parse("20060102 -0700", fmt.Sprintf("%s -0300", rawLine[202:202+8]))
	if err != nil {
		return nil, err
	}
	fatCot, err := strconv.ParseInt(rawLine[210:210+7], 10, 64)
	if err != nil {
		return nil, err
	}
	ptoExe, err := strconv.ParseInt(rawLine[210:210+7], 10, 64)
	if err != nil {
		return nil, err
	}
	distributionNumber, err := strconv.ParseInt(rawLine[242:242+3], 10, 64)
	if err != nil {
		return nil, err
	}
	return &AssetInfo{
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
	}, nil
}

func parseHistoricData(rawData []string, year int) ([]AssetInfo, error) {
	contentList := make([]AssetInfo, len(rawData)-3)
	for i, rawLine := range rawData[1 : len(rawData)-2] {
		slice, err := parseContentLine(rawLine, year)
		if err != nil {
			return nil, err
		}
		contentList[i] = *slice
	}
	return contentList, nil
}

func parseHistoricDataFromBytes(data []byte, year int) ([]AssetInfo, error) {
	return parseHistoricData(strings.Split(string(data), "\n"), year)
}
