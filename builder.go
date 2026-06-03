package main

import (
	"strings"
)

type segment struct {
	text  string
	style string
}

type ArtBuilder struct {
	parts []segment
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{}
}

func (ab *ArtBuilder) AddText(text string) *ArtBuilder {
	ab.parts = append(ab.parts, segment{
		text:  text,
		style: "normal",
	})
	return ab
}

func (ab *ArtBuilder) SetStyle(style string) *ArtBuilder {
	if len(ab.parts) == 0 {
		panic("SetStyle called before AddText")
	}

	switch style {
	case "normal", "bold", "italic", "outline":
		ab.parts[len(ab.parts)-1].style = style
	default:
		panic("unsupported style")
	}

	return ab
}

func (ab *ArtBuilder) Build() string {
	if len(ab.parts) == 0 {
		return ""
	}

	lines := make([]string, 8)

	for _, p := range ab.parts {
		for row := 0; row < 8; row++ {
			lines[row] += render(p.text, p.style, row)
		}
	}

	return strings.Join(lines, "\n") + "\n"
}

func render(text, style string, row int) string {
	switch style {
	case "bold":
		return strings.Repeat(text, 2)

	case "italic":
		return strings.Repeat(" ", row/2+1) + text

	case "outline":
		return "|" + text + "|"

	default: // normal
		return text
	}
}
