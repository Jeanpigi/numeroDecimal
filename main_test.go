package main

import "testing"

func TestContadorDecimal(t *testing.T) {
	tables := []struct {
		number int
		result int
	} {
		{798775, 6},
		{2938103, 7},
		{1234, 4},
		{1920381, 7},
	}

	for _, item := range tables {
		total := ContadorDecimal(item.number)
		if total != item.result {
			t.Errorf("Sum was incorrect, got %v, expected %v", total, item.result)
		}
	}
}