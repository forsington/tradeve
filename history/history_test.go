package history

import (
	"testing"
	"time"
)

func TestTrimDays(t *testing.T) {
	dateFmt := "2006-01-02"
	histories := MarketHistories{
		{
			Date: time.Now().AddDate(0, 0, -5).Format(dateFmt),
		},
		{
			Date: time.Now().AddDate(0, 0, -4).Format(dateFmt),
		},
		{
			Date: time.Now().AddDate(0, 0, -3).Format(dateFmt),
		},
		{
			Date: time.Now().AddDate(0, 0, -2).Format(dateFmt),
		},
		{
			Date: time.Now().AddDate(0, 0, -1).Format(dateFmt),
		},
	}

	days := 2
	trimmed := histories.TrimDays(days)
	if len(trimmed) != days {
		t.Errorf("expected %d, got %d", days, len(trimmed))
	}

}
