package textwrap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitOnSpace(t *testing.T) {
	lines := breakLine("本文实例，讲述了Go语言清除文件中空行的方法。This article ...")
	assert.Equal(t, []string{"本", "文", "实", "例", "，", "讲", "述", "了", "Go", "语", "言", "清", "除", "文", "件", "中", "空", "行", "的", "方", "法", "。", "This", " ", "article", " ", "..."}, lines)
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
