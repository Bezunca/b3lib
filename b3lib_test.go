package b3lib

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {

	fetch := New(5*time.Second, http.DefaultClient)

	tickers := []string{"ITSA4", "HSML11", "VRTA11"}

	prices, errs := fetch(tickers)
	t.Logf("Prices: %+v\nErrors: %#v", prices, errs)

	t.Logf("Trying ITSA4 again with cache")
	prices2, errs2 := fetch([]string{"ITSA4", "CVCB3"})
	t.Logf("Prices: %+v\nErrors: %#v", prices2, errs2)

	t.Logf("Sleeping")
	<-time.After(6 * time.Second)
	t.Logf("Woke")

	t.Logf("Trying ITSA4 again without cache")
	prices3, errs3 := fetch([]string{"ITSA4"})
	t.Logf("Prices: %+v\nErrors: %#v", prices3, errs3)

	assert.Equal(t, len(tickers)+1, len(prices)+len(errs)+1)
}
