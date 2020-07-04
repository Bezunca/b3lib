package history

import "fmt"

type (
	DayOutOfRangeError struct {
		Day uint
	}
	MonthOutOfRangeError struct {
		Month uint
	}
	YearOutOfRangeError struct {
		Year uint
	}

	Not200StatusCode struct {
		StatusCode int
		Err        string
	}
)

func (e *DayOutOfRangeError) Error() string {
	return fmt.Sprintf("day out of range [1, 31]; Day: %d", e.Day)
}

func (e *MonthOutOfRangeError) Error() string {
	return fmt.Sprintf("month out of range [1, 12]; Month: %d", e.Month)
}

func (e *YearOutOfRangeError) Error() string {
	return fmt.Sprintf("year out of range [2000, ∞); Day: %d", e.Year)
}

func (e *Not200StatusCode) Error() string {
	return fmt.Sprintf("received status code %d from B3 api; Error: %s", e.StatusCode, e.Err)
}
