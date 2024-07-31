package item

import "testing"

func TestAppend(t *testing.T) {
	items := Items{
		&Item{Type: &Type{TypeId: 1}},
		&Item{Type: &Type{TypeId: 2}},
		&Item{Type: &Type{TypeId: 3}},
	}

	_, err := items.Append(&Item{})
	if err == nil {
		t.Errorf("items.Append() == nil, want error")
	}

	items, err = items.Append(&Item{Type: &Type{TypeId: 4}})
	if err != nil {
		t.Errorf("items.Append() == %v, want nil", err)
	}
	if items.Len() != 4 {
		t.Errorf("items.Append() == %d, want 4", items.Len())
	}

	// add a duplicate
	items, err = items.Append(&Item{Type: &Type{TypeId: 4}})
	if err != nil {
		t.Errorf("items.Append() == %v, want nil", err)
	}

	if items.Len() != 4 {
		t.Errorf("items.Append() == %d, want 4", items.Len())
	}

}
