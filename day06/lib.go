package day06

func ParseAnswer(s string) int {
	m := make(map[rune]struct{})
	for _, char := range s {
		if char != '\n' {
			m[char] = struct{}{}
		}
	}
	return len(m)
}

func ParseAnswers(groups []string) []int {
	var counts []int
	for _, g := range groups {
		counts = append(counts, ParseAnswer(g))
	}

	return counts
}
