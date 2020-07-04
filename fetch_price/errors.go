package fetch_price

import (
	"errors"
	"fmt"
)

// FetchError for when we had an error while fetching a quotation from B3
type FetchError struct {
	Ticker string
	Date   string
	Err    error
}

// Error method to comply with error interface
func (e *FetchError) Error() string {
	return fmt.Sprintf("failed to fetch B3 quotation for %v on %v: %v", e.Ticker, e.Date, e.Err)
}

// Unwrap method to comply with errors.Unwrap
func (e *FetchError) Unwrap() error {
	return e.Err
}

// Is method to comply with errors.Is
func (e *FetchError) Is(target error) bool {
	tar, ok := target.(*FetchError)
	if !ok {
		return false
	}

	if tar.Ticker == "" && tar.Date == "" && tar.Err == nil {
		return true
	}

	return e.Ticker == tar.Ticker && e.Date == tar.Date && errors.Is(e.Err, tar.Err)
}

// CloseBodyError for when we had an error while closing body from B3 response
type CloseBodyError struct {
	Ticker string
	Date   string
	Err    error
}

// Error method to comply with error interface
func (e *CloseBodyError) Error() string {
	return fmt.Sprintf("failed to close B3 response body for %v on %v: %v", e.Ticker, e.Date, e.Err)
}

// Unwrap method to comply with errors.Unwrap
func (e *CloseBodyError) Unwrap() error {
	return e.Err
}

// Is method to comply with errors.Is
func (e *CloseBodyError) Is(target error) bool {
	tar, ok := target.(*CloseBodyError)
	if !ok {
		return false
	}

	if tar.Ticker == "" && tar.Date == "" && tar.Err == nil {
		return true
	}

	return e.Ticker == tar.Ticker && e.Date == tar.Date && errors.Is(e.Err, tar.Err)
}

// JSONDecodeError for when we had an error while decoding json body from B3 response
type JSONDecodeError struct {
	Ticker string
	Date   string
	Err    error
}

// Error method to comply with error interface
func (e *JSONDecodeError) Error() string {
	return fmt.Sprintf("failed to decode B3 quotation for %v on %v: %v", e.Ticker, e.Date, e.Err)
}

// Unwrap method to comply with errors.Unwrap
func (e *JSONDecodeError) Unwrap() error {
	return e.Err
}

// Is method to comply with errors.Is
func (e *JSONDecodeError) Is(target error) bool {
	tar, ok := target.(*JSONDecodeError)
	if !ok {
		return false
	}

	if tar.Ticker == "" && tar.Date == "" && tar.Err == nil {
		return true
	}

	return e.Ticker == tar.Ticker && e.Date == tar.Date && errors.Is(e.Err, tar.Err)
}
