package order

import (
	"testing"
	"time"
)

func TestFlippers(t *testing.T) {
	orders := Orders{
		&Order{Issued: time.Now().Add(-time.Hour * 36)},
		&Order{Issued: time.Now().Add(-time.Hour)},
		&Order{Issued: time.Now()},
	}

	flippers := orders.Flippers()

	if flippers != 2 {
		t.Errorf("orders.Flippers() == %d, want 2", flippers)
	}
}

func TestToday(t *testing.T) {
	orders := Orders{
		&Order{FetchDate: time.Now().Truncate(24 * time.Hour).Add(-time.Hour)},
		&Order{FetchDate: time.Now().Truncate(24 * time.Hour).Add(-time.Hour * 36)},
		&Order{FetchDate: time.Now().Truncate(24 * time.Hour).Add(time.Hour)},
		&Order{FetchDate: time.Now().Truncate(24 * time.Hour).Add(time.Hour * 2)},
	}

	today := orders.Today()
	if len(today) != 2 {
		t.Errorf("orders.Today() == %d, want 2", len(today))
	}
}
