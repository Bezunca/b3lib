package history

import "testing"

func TestHistoryByYear(t *testing.T) {
	assets, prices, err := GetByYear(2019)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", assets)
	for _, price := range prices {
		t.Logf("%+v", price)
	}
}

func TestHistorySpecificDay(t *testing.T) {
	assets, prices, err := GetSpecificDay(30, 9, 2019)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", assets)
	for _, price := range prices {
		t.Logf("%+v", price)
	}
}
