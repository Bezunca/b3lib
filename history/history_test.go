package history

import "testing"

func TestHistoryByYear(t *testing.T) {
	assets, err := GetByYear(2019)
	if err != nil {
		t.Fatal(err)
	}

	for _, asset := range assets {
		t.Logf("%+v", asset)
	}

}

func TestHistorySpecificDay(t *testing.T) {
	assets, err := GetSpecificDay(30, 9, 2019)
	if err != nil {
		t.Fatal(err)
	}

	for _, asset := range assets {
		t.Logf("%+v", asset)
	}

}
