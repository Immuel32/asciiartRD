package main

import (
	"strings"
)

type Animation struct {
	text       string
	frameCount int
	frames     []string
}

func NewAnimation(text string, frameCount int) *Animation {
	return &Animation{
		text:       text,
		frameCount: frameCount,
		frames:     make([]string, frameCount),
	}
}

func (a *Animation) GenerateSpinFrames() {
	chars := []string{"|", "/", "-", "\\"}

	for i := 0; i < a.frameCount; i++ {
		symbol := chars[i%len(chars)]
		a.frames[i] = makeFrame(symbol + " " + a.text + " " + symbol)
	}
}

func (a *Animation) GenerateWaveFrames() {
	for i := 0; i < a.frameCount; i++ {
		padding := strings.Repeat(" ", i)
		a.frames[i] = makeFrame(padding + a.text)
	}
}

func (a *Animation) GenerateZoomFrames() {
	for i := 0; i < a.frameCount; i++ {
		text := a.text

		for j := 0; j < i; j++ {
			text += "!"
		}

		a.frames[i] = makeFrame(text)
	}
}

func (a *Animation) GetFrame(index int) string {
	if len(a.frames) == 0 {
		return ""
	}

	return a.frames[index%len(a.frames)]
}

func (a *Animation) Play() string {
	var out strings.Builder

	for i, frame := range a.frames {
		out.WriteString("=== Frame ")
		out.WriteString(string(rune('0' + i)))
		out.WriteString(" ===\n")
		out.WriteString(frame)
		out.WriteString("\n")
	}

	return out.String()
}

func makeFrame(content string) string {
	const width = 20
	lines := make([]string, 10)

	line := content
	if len(line) < width {
		line += strings.Repeat(" ", width-len(line))
	}

	for i := 0; i < 10; i++ {
		lines[i] = line
	}

	return strings.Join(lines, "\n") + "\n"
}
