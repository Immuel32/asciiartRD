package main

func GeneratePattern(c rune) []string {

	patterns := map[rune][]string{
		'A': {
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		},
	}
	pattern, ok := patterns[c]
	if !ok {
		return []string{}
	}
	return pattern
}
