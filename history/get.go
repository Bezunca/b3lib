package history

import (
	"fmt"

	zip "github.com/Bezunca/zip_in_memory"
)

// Get returns a list of infos for assets in B3
func GetByYear(year uint) (map[string]*Asset, []Price, error) {
	responseData, err := download(fmt.Sprintf("http://bvmf.bmfbovespa.com.br/InstDados/SerHist/COTAHIST_A%v.ZIP", year))
	if err != nil {
		return nil, nil, err
	}
	encodedContent, err := zip.ExtractFirstFile(responseData)
	if err != nil {
		return nil, nil, err
	}
	assets, prices, err := parseHistoricDataFromBytes(encodedContent, int(year))
	if err != nil {
		return nil, nil, err
	}

	return assets, prices, nil
}

func GetSpecificDay(day, month, year uint) (map[string]*Asset, []Price, error) {
	if day > 31 || day < 1 {
		return nil, nil, &DayOutOfRangeError{Day: day}
	}
	if month > 12 || month < 1 {
		return nil, nil, &MonthOutOfRangeError{Month: month}
	}
	if year < 2000 {
		return nil, nil, &YearOutOfRangeError{Year: year}
	}

	responseData, err := download(fmt.Sprintf("http://bvmf.bmfbovespa.com.br/InstDados/SerHist/COTAHIST_D%02d%02d%d.ZIP", day, month, year))
	if err != nil {
		return nil, nil, err
	}
	encodedContent, err := zip.ExtractFirstFile(responseData)
	if err != nil {
		return nil, nil, err
	}
	assets, prices, err := parseHistoricDataFromBytes(encodedContent, int(year))
	if err != nil {
		return nil, nil, err
	}

	return assets, prices, nil
}
