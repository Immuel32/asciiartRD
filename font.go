package main

func GenerateFont() map[rune][]string {
	font := make(map[rune][]string)

	for c := rune(32); c <= 126; c++ {
		lines := make([]string, 8)

		for i := 0; i < 8; i++ {
			row := []rune("........")

			switch {
			case c == ' ':
				row = []rune("        ")

			case c >= 'A' && c <= 'Z':
				if isVowel(c) {
					row[0], row[7], row[i] = '*', '*', '*'
				} else {
					for j := 0; j < 8; j++ {
						if (i+j)%2 == 0 {
							row[j] = '*'
						} else {
							row[j] = ' '
						}
					}
				}

			case c >= 'a' && c <= 'z':
				for j := 0; j < 8; j++ {
					if (i+j)%2 == 1 {
						row[j] = '*'
					} else {
						row[j] = ' '
					}
				}

			case c >= '0' && c <= '9':
				if i == 0 || i == 3 || i == 7 {
					row = []rune("********")
				} else {
					row[0], row[7] = '*', '*'
				}

			default:
				row[2], row[5] = '*', '*'
			}

			lines[i] = string(row)
		}

		font[c] = lines
	}

	return font
}

func isVowel(c rune) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
