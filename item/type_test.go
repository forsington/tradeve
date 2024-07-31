package item

import (
	"testing"

	"github.com/forsington/tradeve/config"
)

func TestExcludeGroups(t *testing.T) {
	types := Types{
		&Type{TypeId: 1, TypeName: "Tritanium", GroupId: 1},
		&Type{TypeId: 2, TypeName: "Some skin", GroupId: 1953},
		&Type{TypeId: 3, TypeName: "A nice dress", GroupId: 53},
	}

	cases := []struct {
		name       string
		exclusions *config.ExcludeGroups
		expected   int
	}{
		{
			name:       "no exclusions",
			exclusions: &config.ExcludeGroups{},
			expected:   3,
		},
		{
			name:       "exclude skins",
			exclusions: &config.ExcludeGroups{Skins: true},
			expected:   2,
		},
		{
			name:       "exclude skins and Wearables",
			exclusions: &config.ExcludeGroups{Skins: true, Wearables: true},
			expected:   1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := types.ExcludeGroups(tc.exclusions)
			if len(actual) != tc.expected {
				t.Errorf("expected %d types, got %d", tc.expected, len(actual))
			}
		})
	}
}
