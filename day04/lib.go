package day04

import (
	"strings"
)

type Passport map[string]string

func ParsePassports(s string) ([]Passport, error) {
	entries := strings.Split(s, "\n\n")
	var passports []Passport
	for _, entry := range entries {
		if entry == "" {
			continue
		}

		entry = strings.ReplaceAll(entry, "\n", " ")
		columns := strings.Split(entry, " ")
		p := make(map[string]string)
		for _, column := range columns {
			if column == "" {
				continue
			}

			pairs := strings.Split(column, ":")
			p[pairs[0]] = pairs[1]
		}
		passports = append(passports, p)
	}

	return passports, nil
}

func IsValid(p Passport) bool {
	_, cidPresent := p["cid"]
	return len(p) == 8 || (len(p) == 7 && !cidPresent)
}

func CountValid(ps []Passport) int {
	count := 0
	for _, p := range ps {
		if IsValid(p) {
			count++
		}
	}
	return count
}
