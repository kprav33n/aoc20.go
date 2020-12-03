package day02

import (
	"fmt"
	"strings"
)

type Entry struct {
	Min      int
	Max      int
	Char     byte
	Password []byte
}

func ParseLines(content string) []Entry {
	var entries []Entry
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			continue
		}

		entry, err := ParseEntry(line)
		if err != nil {
			panic(fmt.Sprintf("%s: %s", line, err))
		}
		entries = append(entries, entry)
	}
	return entries
}

func ParseEntry(s string) (Entry, error) {
	var entry Entry
	_, err := fmt.Sscanf(
		s, "%d-%d %c: %s",
		&entry.Min, &entry.Max, &entry.Char, &entry.Password,
	)
	return entry, err
}

func IsValid(entry Entry) bool {
	count := 0
	for _, ch := range entry.Password {
		if ch == entry.Char {
			count += 1
		}
	}

	return entry.Min <= count && count <= entry.Max
}

func IsValid2(entry Entry) bool {
	first := entry.Password[entry.Min-1] == entry.Char
	second := entry.Password[entry.Max-1] == entry.Char
	return first != second
}

func Count(predicate func(Entry) bool, entries []Entry) int {
	count := 0
	for _, entry := range entries {
		if predicate(entry) {
			count += 1
		}
	}
	return count
}
