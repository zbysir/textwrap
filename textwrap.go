package textwrap

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"strings"
	"unicode"
)

// TextWrap Wrap text without breaking words as much as possible;
// Support Chinese characters and super long words
func TextWrap(s string, f font.Face, width float64) []string {
	return textWrap(s, func(s string) bool {
		return fontWidth(f, s) > width
	})
}

// ParseFontFace Parse font file to FontFace
func ParseFontFace(fontBody []byte, fontSize float64) (font.Face, error) {
	f, err := truetype.Parse(fontBody)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(f, &truetype.Options{
		Size: fontSize,
	})

	return face, nil
}

func splitOnSpace(x string) []string {
	var result []string
	pi := 0
	ps := false
	for i, c := range x {
		isHan := unicode.Is(unicode.Han, c)
		s := unicode.IsSpace(c) || isHan
		if (s != ps || isHan) && i > 0 {
			result = append(result, x[pi:i])
			pi = i
		}
		ps = s
	}
	result = append(result, x[pi:])
	return result
}

func breakWord(s string, over func(string) bool) ([]string, string) {
	if !over(s) {
		return nil, s
	}
	var result []string
	for {
		for i := range s {
			if !over(s[:i+1]) {
				continue
			}

			result = append(result, s[:i])
			s = s[i:]
			break
		}
		if s == "" || !over(s) {
			break
		}
	}

	return result, s
}

func textWrap(s string, over func(string) bool) []string {
	var result []string
	for _, l := range strings.Split(s, "\n") {
		words := splitOnSpace(l)
		line := ""
		for _, word := range words {
			testLine := line + word
			if over(testLine) {
				if strings.TrimSpace(line) != "" {
					i, last := breakWord(line, over)
					result = append(result, i...)

					// for case:
					//  |Unite|
					//  |d St |
					if !over(last + word) {
						line = last + word
					} else {
						if last != "" {
							result = append(result, last)
						}
						line = word
					}
				} else {
					line = word
				}
			} else {
				line = testLine
			}
		}

		if strings.TrimSpace(line) != "" {
			word, last := breakWord(line, over)
			result = append(result, word...)
			if last != "" {
				result = append(result, last)
			}
		}
	}

	for i, line := range result {
		result[i] = strings.TrimSpace(line)
	}
	return result
}

func fontWidth(f font.Face, s string) float64 {
	advance := font.MeasureString(f, s)
	return float64(advance >> 6)
}
