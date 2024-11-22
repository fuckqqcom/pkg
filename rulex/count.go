package rulex

import "regexp"

var (
	chineseRegex = regexp.MustCompile("[\u4e00-\u9fa5]")
	englishRegex = regexp.MustCompile(`[a-zA-Z]+`)
)

func WordCount(text string) (int, int) {
	chineseCount := len(chineseRegex.FindAllString(text, -1))
	englishCount := len(englishRegex.FindAllString(text, -1))
	return chineseCount, englishCount
}
