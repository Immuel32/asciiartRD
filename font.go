package main

func GenerateFont() map[rune][]string {
	font := make(map[rune][]string)

	for c := rune(32); c <= rune(126); c++ {
		lines := make([]string, 8)

		if c == ' ' {
			for i := 0; i < 8; i++ {
				lines[i] = "        "
			}
			font[c] = lines
			continue
		}

		for i := 0; i < 8; i++ {
			line := []rune("........")
			lines[i] = string(line)
		}

		switch {
		case c >= 'A' && c <= 'Z':

			if isVowel(c) {
				for i := 0; i < 8; i++ {
					row := []rune(lines[i])
					row[0], row[7] = '*', '*'
					row[i] = '*'
					lines[i] = string(row)
				}
			} else {
				for i := 0; i < 8; i++ {
					row := []rune(lines[i])
					for j := 0; j < 8; j++ {
						if (i+j)%2 == 0 {
							row[j] = '*'
						} else {
							row[j] = ' '
						}
					}
					lines[i] = string(row)
				}
			}

		case c >= 'a' && c <= 'z':
			for i := 0; i < 8; i++ {
				row := []rune(lines[i])
				for j := 0; j < 8; j++ {
					if (i+j)%2 == 1 {
						row[j] = '*'
					} else {
						row[j] = ' '
					}
				}
				lines[i] = string(row)
			}

		case c >= '0' && c <= '9':

			for i := 0; i < 8; i++ {
				row := []rune(lines[i])
				if i == 0 || i == 7 || i == 3 {
					for j := 0; j < 8; j++ {
						row[j] = '*'
					}
				} else {
					row[0], row[7] = '*', '*'
				}
				lines[i] = string(row)
			}

		default:
			for i := 0; i < 8; i++ {
				row := []rune(lines[i])
				row[2], row[5] = '*', '*'
				lines[i] = string(row)
			}
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
