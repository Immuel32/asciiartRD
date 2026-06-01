package main

import (
	"fmt"
	"strings"
)

type Animation struct {
	text   string
	frames int
	data   []string
}

func NewAnimation(text string, frames int) *Animation {
	return &Animation{
		text:   text,
		frames: frames,
		data:   make([]string, frames),
	}
}

func padToTenLines(lines []string) string {
	for len(lines) < 10 {
		lines = append(lines, strings.Repeat(" ", len(lines[0])))
	}
	return strings.Join(lines, "\n")
}

func (a *Animation) GenerateSpinFrames() {
	for i := 0; i < a.frames; i++ {

		shift := i % len(a.text)
		line := a.text[shift:] + a.text[:shift]
		lines := []string{}
		for j := 0; j < 10; j++ {
			lines = append(lines, line)
		}
		a.data[i] = padToTenLines(lines)
	}
}

func (a *Animation) GenerateWaveFrames() {
	for i := 0; i < a.frames; i++ {
		lines := []string{}
		for j := 0; j < 10; j++ {

			offset := (i + j) % (len(a.text) + 1)
			line := strings.Repeat(" ", offset) + a.text
			lines = append(lines, line)
		}
		a.data[i] = padToTenLines(lines)
	}
}
func (a *Animation) GenerateZoomFrames() {
	for i := 0; i < a.frames; i++ {
		scale := 1 + i%3 // zoom levels
		line := strings.Repeat(a.text+" ", scale)
		lines := []string{}
		for j := 0; j < 10; j++ {
			lines = append(lines, line)
		}
		a.data[i] = padToTenLines(lines)
	}
}

func (a *Animation) GetFrame(index int) string {
	if a.frames == 0 {
		return ""
	}
	return a.data[index%a.frames]
}

func (a *Animation) Play() string {
	var sb strings.Builder
	for i := 0; i < a.frames; i++ {
		sb.WriteString(fmt.Sprintf("=== Frame %d ===\n", i))
		sb.WriteString(a.GetFrame(i))
		sb.WriteString("\n\n")
	}
	return sb.String()
}

func main() {
	anim := NewAnimation("HELLO", 4)
	anim.GenerateSpinFrames()
	fmt.Println(anim.Play())

	anim2 := NewAnimation("WAVE", 6)
	anim2.GenerateWaveFrames()
	fmt.Println(anim2.Play())

	anim3 := NewAnimation("ZOOM", 5)
	anim3.GenerateZoomFrames()
	fmt.Println(anim3.Play())
}
