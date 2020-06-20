package b3lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// New function to setup B3 fetch and cache prices
func New(cacheTimeout time.Duration, client *http.Client) func(tickers []string) ([]FetchedPrice, []error) {

	cache := make(map[string]tickerEntry)

	return func(tickers []string) (prices []FetchedPrice, err []error) {

		var newTickers []string
		for _, ticker := range tickers {
			if entry, ok := cache[ticker]; !ok || time.Since(entry.Timestamp) > cacheTimeout {
				newTickers = append(newTickers, ticker)
			} else {
				prices = append(prices, entry.Price)
			}
		}

		var fetched []FetchedPrice
		if len(newTickers) > 0 {
			fetched, err = fetchNewPrices(client, newTickers)
			for _, price := range fetched {
				cache[price.Ticker] = tickerEntry{
					Price:     price,
					Timestamp: time.Now(),
				}
			}
		}

		return append(prices, fetched...), err
	}
}

func fetchNewPrices(client *http.Client, tickers []string) ([]FetchedPrice, []error) {
	today := time.Now()
	todayWeekday := today.Weekday()
	if todayWeekday == time.Sunday {
		today = today.Add(-48 * time.Hour)
	} else if todayWeekday == time.Saturday || today.Hour() < 6 {
		// B3 only got prices values after market opening, so before 6AM, we get yesterday values
		// If it's after 6AM, we zero the prices values, waiting for market opening
		today = today.Add(-24 * time.Hour)
	}
	date := today.Format("2006-01-02")

	fetchedPricesChan := make(chan FetchedPrice, len(tickers))
	errChan := make(chan error)
	defer func() {
		close(fetchedPricesChan)
		close(errChan)
	}()

	for _, ticker := range tickers {
		go getCurrentPrice(client, date, ticker, fetchedPricesChan, errChan)
	}

	var fetchedPrices []FetchedPrice
	var errors []error
	for i := 0; i < len(tickers); i++ {
		select {
		case fetchedPrice := <-fetchedPricesChan:
			fetchedPrices = append(fetchedPrices, fetchedPrice)
		case err := <-errChan:
			errors = append(errors, err)
		}
	}

	if len(errors) != 0 {
		return nil, errors
	}

	return fetchedPrices, nil
}

func getCurrentPrice(client *http.Client, date string, ticker string, prices chan<- FetchedPrice, errChan chan<- error) {
	url := fmt.Sprintf("https://arquivos.b3.com.br/apinegocios/ticker/%v/%v", ticker, date)
	response, err := client.Get(url)
	if err != nil {
		errChan <- &FetchError{ticker, date, err}
		return
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			errChan <- &CloseBodyError{ticker, date, err}
		}
	}()

	b3Response := new(b3PriceResponse)
	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(b3Response); err != nil {
		errChan <- &JSONDecodeError{ticker, date, err}
		return
	}

	price := 0
	if len(b3Response.Values) != 0 {
		price, _ = strconv.Atoi(
			strings.ReplaceAll(fmt.Sprint(b3Response.Values[0][2].(float64)), ".", ""),
		)
	}

	prices <- FetchedPrice{
		Ticker:   b3Response.Name,
		IntPrice: price,
	}
}
