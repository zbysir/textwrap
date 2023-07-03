# Textwrap
Text wrap algorithms like css ['overflow-wrap: break-word;'](https://developer.mozilla.org/zh-CN/docs/Web/CSS/overflow-wrap)

![overflow-wrap: break-word](overflow-wrap-demo.png)

- Wrap text without breaking words as much as possible
- Support Chinese characters and super long words

## Usage
```
go get github.com/zbysir/textwrap
```
```go
textwrap.TextWrap("text", func(s string) bool {
		return len(s) > 10
})
```

## Known problems
- Does not follow the ["Unicode line breaking algorithm"](http://unicode.org/reports/tr14/)
