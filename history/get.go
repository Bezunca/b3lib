package history

import (
	"fmt"

	zip "github.com/Bezunca/zip_in_memory"
)

// Get returns a list of infos for assets in B3
func GetByYear(year uint) ([]AssetInfo, error) {
	responseData, err := download(fmt.Sprintf("http://bvmf.bmfbovespa.com.br/InstDados/SerHist/COTAHIST_A%v.ZIP", year))
	if err != nil {
		return nil, err
	}
	encodedContent, err := zip.ExtractFirstFile(responseData)
	if err != nil {
		return nil, err
	}
	data, err := parseHistoricDataFromBytes(encodedContent, int(year))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetSpecificDay(day, month, year uint) ([]AssetInfo, error) {
	if day > 31 || day < 1 {
		return nil, &DayOutOfRangeError{Day: day}
	}
	if month > 12 || month < 1 {
		return nil, &MonthOutOfRangeError{Month: month}
	}
	if year < 2000 {
		return nil, &YearOutOfRangeError{Year: year}
	}

	responseData, err := download(fmt.Sprintf("http://bvmf.bmfbovespa.com.br/InstDados/SerHist/COTAHIST_D%02d%02d%d.ZIP", day, month, year))
	if err != nil {
		return nil, err
	}
	encodedContent, err := zip.ExtractFirstFile(responseData)
	if err != nil {
		return nil, err
	}
	data, err := parseHistoricDataFromBytes(encodedContent, int(year))
	if err != nil {
		return nil, err
	}

	return data, nil
}
