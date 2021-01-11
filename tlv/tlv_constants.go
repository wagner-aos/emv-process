package tlv

import "github.com/elliotchance/orderedmap"

type Bit struct {
	bit      int
	subField string
	value    string
	size     int
}

var bit48 = []tag{
	{name: "01", def: 2, size: 1},
	{name: "02", def: 2, size: 4},
	{name: "03", def: 2, size: 2},
	{name: "04", def: 2, size: 1},
	{name: "05", def: 2, size: 42},
	{name: "06", def: 2, size: 11},
	{name: "07", def: 2, size: 11},
	{name: "08", def: 2, size: 15},
	{name: "09", def: 2, size: 1},
	{name: "10", def: 2, size: 3},
	{name: "11", def: 2, size: 4},
	{name: "12", def: 2, size: 28},
	{name: "13", def: 2, size: 28},
	{name: "14", def: 2, size: 28},
	{name: "15", def: 2, size: 57},
	{name: "16", def: 2, size: 5},
	{name: "17", def: 2, size: 6},
	{name: "18", def: 2, size: 20},
	{name: "19", def: 2, size: 1},
	{name: "20", def: 2, size: 28},
	{name: "21", def: 2, size: 37},
	{name: "22", def: 2, size: 40},
	{name: "23", def: 2, size: 39},
	{name: "24", def: 2, size: 1},
}

var bit47 = []tag{
	{name: "01", def: 2, size: 3},
	{name: "02", def: 2, size: 4},
	{name: "03", def: 2, size: 12},
	{name: "04", def: 2, size: 15},
	{name: "05", def: 2, size: 4},
	{name: "06", def: 2, size: 12},
	{name: "07", def: 2, size: 12},
	{name: "08", def: 2, size: 1},
	{name: "09", def: 2, size: 1},
	{name: "10", def: 2, size: 1},
	{name: "11", def: 2, size: 1},
	{name: "12", def: 2, size: 14},
	{name: "13", def: 2, size: 1},
	{name: "14", def: 2, size: 1},
	{name: "15", def: 2, size: 6},
	{name: "16", def: 2, size: 1},
	{name: "17", def: 2, size: 15},
	{name: "18", def: 2, size: 30},
	{name: "19", def: 2, size: 2},
	{name: "20", def: 2, size: 2},
	{name: "21", def: 2, size: 3},
	{name: "22", def: 2, size: 10},
	{name: "23", def: 2, size: 2},
	{name: "24", def: 2, size: 5},
	{name: "25", def: 2, size: 2},
	{name: "26", def: 2, size: 1},
	{name: "27", def: 2, size: 1},
	{name: "28", def: 2, size: 1},
	{name: "27", def: 2, size: 50},
	{name: "30", def: 2, size: 6},
	{name: "31", def: 2, size: 1},
	{name: "32", def: 2, size: 1},
	{name: "33", def: 2, size: 2},
	{name: "35", def: 2, size: 4},
	{name: "36", def: 2, size: 6},
	{name: "37", def: 2, size: 2},
	{name: "38", def: 2, size: 1},
}

var bit61 = []tag{
	{name: "01", def: 2, size: 29},
	{name: "02", def: 2, size: 48},
	{name: "03", def: 2, size: 5}, // TODO: verificar size 6 documentacao
}

var bit112 = []tag{
	{name: "01", def: 3, size: 2},
	{name: "02", def: 3, size: 12},
	{name: "03", def: 3, size: 218},
	{name: "04", def: 3, size: 230},
	{name: "05", def: 3, size: 230},
	{name: "06", def: 3, size: 230},
}

var tagMap = orderedmap.NewOrderedMap()

//load Emv tags available
func loadTags(field int) []tag {
	switch field {
	case 47:
		get(bit47)
	case 48:
		get(bit48)
	case 61:
		get(bit61)
	case 112:
		get(bit112)
	}

	return []tag{}
}

func get(bits []tag) {
	for _, f := range bits {
		tagMap.Set(f.name, tag{name: f.name, def: f.def, size: f.size})
	}
}
