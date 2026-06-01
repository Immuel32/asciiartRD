package main

import (
	"fmt"
	"strings"
)

type ArtBuilder struct {
	texts  []string
	styles []string
}

func NewArtBuilder() *ArtBuilder { return &ArtBuilder{} }

func (ab *ArtBuilder) AddText(text string) *ArtBuilder {
	ab.texts = append(ab.texts, text)
	ab.styles = append(ab.styles, "normal")
	return ab
}

func (ab *ArtBuilder) SetStyle(style string) *ArtBuilder {
	if len(ab.styles) == 0 {
		panic("SetStyle called before AddText")
	}
	switch style {
	case "normal", "bold", "italic", "outline":
		ab.styles[len(ab.styles)-1] = style
	default:
		panic(fmt.Sprintf("Unsupported style: %s", style))
	}
	return ab
}

func (ab *ArtBuilder) Build() string {
	if len(ab.texts) == 0 {
		return ""
	}
	lines := make([]string, 8)
	for i, txt := range ab.texts {
		style := ab.styles[i]
		block := render(txt, style)
		for r := 0; r < 8; r++ {
			lines[r] += block[r]
		}
	}
	return strings.Join(lines, "\n") + "\n"
}

func render(text, style string) []string {
	out := make([]string, 8)
	for r := 0; r < 8; r++ {
		line := ""
		for _, ch := range text {
			base := baseChar(ch, r)
			switch style {
			case "bold":
				line += double(base)
			case "italic":
				line += strings.Repeat(" ", r/2) + base
			case "outline":
				line += "|" + base[:len(base)-1] + "|"
			default:
				line += base
			}
		}
		out[r] = line
	}
	return out
}

func baseChar(_ rune, row int) string {
	pattern := []string{
		"*    * ",
		"*    * ",
		"****** ",
		"*    * ",
		"*    * ",
		"*    * ",
		"*    * ",
		"       ",
	}
	return pattern[row]
}

func double(s string) string {
	res := ""
	for _, r := range s {
		res += string(r) + string(r)
	}
	return res
}
