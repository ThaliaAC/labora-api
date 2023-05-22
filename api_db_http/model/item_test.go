package model

import (
	"testing"
)

type Test struct {
	item     Item
	expected int
}

var tests []Test = []Test{
	{
		item:     Item{Product: "licuadora", Quantity: 4, Price: 169},
		expected: 676,
	},
	{
		item:     Item{Product: "aspiradora", Quantity: 2, Price: 234},
		expected: 468,
	},
	{
		item:     Item{Product: "celular Xiomi", Quantity: 3, Price: 2100},
		expected: 6300,
	},
	{
		item:     Item{Product: "smartwatch", Quantity: 5, Price: 220},
		expected: 1100,
	},
	{
		item:     Item{Product: "horno", Quantity: 1, Price: 400},
		expected: 400,
	},
}

func TestTotalPriceWorksTDT(t *testing.T) {

	for _, test := range tests {
		generated := test.item.TotalPrice()
		if generated != test.expected {
			t.Errorf("Error! Generated: %.2v, Expected: %v", generated, test.expected)
		}
	}
}
