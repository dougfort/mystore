package main

// catalogEntry describes a single item from an imaginary store
type catalogEntry struct {
	itemID      string
	description string
	price       int
	available   int
}

var catalog = []catalogEntry{
	catalogEntry{
		itemID:      "1001",
		description: "abacus",
		price:       1,
		available:   1000000000,
	},
	catalogEntry{
		itemID:      "1002",
		description: "pile of rocks",
		price:       2,
		available:   1000000000,
	},
	catalogEntry{
		itemID:      "1003",
		description: "sliderule",
		price:       3,
		available:   1000000000,
	},
	catalogEntry{
		itemID:      "1004",
		description: "Curta",
		price:       4,
		available:   1000000000,
	},
	catalogEntry{
		itemID:      "1005",
		description: "IBM 360",
		price:       1000,
		available:   1000000000,
	},
	catalogEntry{
		itemID:      "1006",
		description: "Qubit",
		price:       6,
		available:   1000000000,
	},
	catalogEntry{
		itemID:      "1007",
		description: "human being",
		price:       7,
		available:   1000000000,
	},
	catalogEntry{
		itemID:      "1008",
		description: "Clever Hans",
		price:       8,
		available:   1000000000,
	},
}
