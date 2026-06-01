package main

import (
	"strings"
)

func StringToArt(input string) string {
	if input == "" {
		return ""
	}
	digits := map[rune][]string{
		'0': {
			" ___ ",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},

		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},

		'2': {
			"|___ ",
			"    |",
			" ___|",
			"|    ",
			"|___ ",
		},
		'3': {
			"____ ",
			"    |",
			"____|",
			"    |",
			"____|",
		},
		'4': {
			"|   |",
			"|   |",
			"|___|",
			"    |",
			"    |",
		},
		'5': {
			" ___ ",
			"|",
			"|___ ",
			"    |",
			"___ |",
		},
		'6': {
			" ___ ",
			"|    ",
			"|___ ",
			"|   |",
			"|___|",
		},
		'7': {
			"____ ",
			"    |",
			"   | ",
			"  |  ",
			"|    ",
		},
		'8': {
			" ___ ",
			"|   |",
			"|___|",
			"|   |",
			"|___|",
		},
		'9': {
			" ___ ",
			"|   |",
			"|___|",
			"    |",
			" ___|",
		},
	}
	var result strings.Builder
	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		for _, char := range line {
			if _, ok := digits[char]; !ok {
				return ""
			}
		}
		for row := 0; row < 5; row++ {
			for i, char := range line {
				part := digits[char][row]
				if len(line) > 1 && i < len(line)-1 {
					part = strings.TrimRight(part, " ")
				}
				result.WriteString(part)
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
