package main

import "testing"

func TestMax(t *testing.T) {
	tables := []struct {
		x int
		y int
		maxXY int
	}{
		{1, 2, 2},
		{3, 2, 3},
		{2, 2, 2},
		{-10, -1, -1},
	}

	for _, table := range tables {
		maxReturned := max(table.x, table.y)
		if maxReturned != table.maxXY {
			t.Errorf("Max of {%d, %d} was incorrect, got: %d, want: %d.",
				table.x, table.y, maxReturned, table.maxXY)
		}
	}
}


func TestMin(t *testing.T) {
	tables := []struct {
		x int
		y int
		minXY int
	}{
		{1, 2, 1},
		{3, 2, 2},
		{2, 2, 2},
		{-10, -1, -10},
	}

	for _, table := range tables {
		minReturned := min(table.x, table.y)
		if minReturned != table.minXY {
			t.Errorf("Min of {%d, %d} was incorrect, got: %d, want: %d.",
				table.x, table.y, minReturned, table.minXY)
		}
	}
}
