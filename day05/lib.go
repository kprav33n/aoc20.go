package day05

type Range struct {
	Min int
	Max int
}

func (r *Range) NarrowLower() {
	r.Max = (r.Max-r.Min)/2 + r.Min
}

func (r *Range) NarrowHigher() {
	r.Min = ((r.Max - r.Min) / 2) + r.Min + 1
}

func ParsePass(s string) int {
	row := Range{Min: 0, Max: 127}
	col := Range{Min: 0, Max: 7}

	for _, char := range s {
		switch char {
		case 'F':
			row.NarrowLower()
		case 'B':
			row.NarrowHigher()
		case 'L':
			col.NarrowLower()
		case 'R':
			col.NarrowHigher()
		}
	}

	return row.Min*8 + col.Min
}

func ParsePasses(passes []string) []int {
	var result []int
	for _, s := range passes {
		result = append(result, ParsePass(s))
	}

	return result
}
