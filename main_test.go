package main_test

import (
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCliArgs(t *testing.T) {
	const binaryName = "aoc20"

	tests := []struct {
		name   string
		args   []string
		output string
	}{
		{"day01a", []string{"day01a"}, "567171"},
		{"day01b", []string{"day01b"}, "212428694"},
		{"day02a", []string{"day02a"}, "524"},
		{"day02b", []string{"day02b"}, "485"},
		{"day03a", []string{"day03a"}, "294"},
		{"day03b", []string{"day03b"}, "5774564250"},
		{"day04a", []string{"day04a"}, "222"},
		{"day05a", []string{"day05a"}, "955"},
		{"day06a", []string{"day06a"}, "6748"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command(path.Join(dir, binaryName), tt.args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}

			require.New(t).Equal(tt.output, strings.Trim(string(output), "\n\r"))
		})
	}
}
