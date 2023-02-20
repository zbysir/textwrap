package textwrap

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"strings"
	"unicode"
	"unicode/utf8"
)

// TextWrapFont measure and Wrap text
func TextWrapFont(s string, f font.Face, width float64) []string {
	return TextWrap(s, func(s string) bool {
		return measureWidth(f, s) > width
	})
}

// TextWrap Wrap text without breaking words as much as possible;
// Support Chinese characters and super long words
func TextWrap(s string, over func(string) bool) []string {
	var result []string
	for _, l := range strings.Split(s, "\n") {
		words := breakLine(l)
		line := ""
		for _, next := range words {
			testLine := line + next
			if !over(testLine) {
				line = testLine
				continue
			}

			if strings.TrimSpace(line) == "" {
				line = next
				continue
			}

			i, last := breakWord(line, over)
			result = append(result, i...)
			// 如果一个单词被分为了多行则需要处理最后一行，尝试拼接 next
			//  case:
			//  |Unite|
			//  |d St |
			if last != "" {
				if !over(last + next) {
					line = last + next
					continue
				} else {
					result = append(result, last)
				}
			}

			// 超出之后下一个单词就是下一行的开始
			line = next
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

// ParseFont Parse font file to FontFace
func ParseFont(fontBody []byte, fontSize float64) (font.Face, error) {
	f, err := truetype.Parse(fontBody)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(f, &truetype.Options{
		Size: fontSize,
	})

	return face, nil
}

// return "break opportunities" (http://unicode.org/reports/tr14/)
func breakLine(x string) []string {
	var result []string
	pi := 0
	prevKind := -1
	rs := []rune(x)
	for i, c := range rs {
		var kind = -1
		if !unicode.IsSpace(c) {
			kind = utf8.RuneLen(c)
		}

		isWide := kind >= 3
		// break on each Wide char and space
		if (kind != prevKind || isWide) && i != 0 {
			result = append(result, string(rs[pi:i]))
			pi = i
		}
		prevKind = kind
	}
	result = append(result, string(rs[pi:]))
	return result
}

func breakWord(s string, over func(string) bool) ([]string, string) {
	var result []string
	runes := []rune(s)
	for len(runes) != 0 && over(string(runes)) {
		max := 0
		for i := range runes {
			if over(string(runes[:i+1])) {
				max = i
				break
			}
		}

		if max == 0 {
			max = 1
		}

		result = append(result, string(runes[:max]))
		runes = runes[max:]
	}

	return result, string(runes)
}

func measureWidth(f font.Face, s string) float64 {
	advance := font.MeasureString(f, s)
	return float64(advance >> 6)
}
