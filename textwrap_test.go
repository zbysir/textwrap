package textwrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreakLine(t *testing.T) {
	lines := breakLine("🍓🍓标准的具体规则在Unicode 换行算法 (Unicode Line Breaking Algorithm, UAX #14）中详细给出。")
	assert.Equal(t, []string{"🍓", "🍓", "标", "准", "的", "具", "体", "规", "则", "在", "Unicode", " ", "换", "行", "算", "法", " ", "(Unicode", " ", "Line", " ", "Breaking", " ", "Algorithm,", " ", "UAX", " ", "#14", "）", "中", "详", "细", "给", "出", "。"}, lines)
}

func TestTextWrap(t *testing.T) {
	cases := []struct {
		Text  string
		Width int
		Out   []string
	}{
		{
			Text:  "本文实例，讲述了Go语言清除文件中空行的方法。This article ...",
			Width: 9,
			Out:   []string{"本文实例，讲述了", "Go语言清除文件中", "空行的方法。", "This", "article", "..."},
		},
		{
			Text:  "United St",
			Width: 5,
			Out:   []string{"Unite", "d St"},
		},
	}
	for _, c := range cases {
		t.Run(c.Text, func(t *testing.T) {
			lines := TextWrap(c.Text, func(s string) bool {
				return len([]rune(s)) > c.Width
			})

			assert.Equal(t, c.Out, lines)
		})
	}
}

func TestBreakWord(t *testing.T) {
	lines, last := breakWord("中国字", func(s string) bool {
		return len([]rune(s)) >= 1
	})
	assert.Equal(t, []string{"中", "国", "字"}, lines)
	assert.Equal(t, "", last)
}
