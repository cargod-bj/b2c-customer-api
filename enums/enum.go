package enum

import "sort"

const SmsTimeOut int = 10 //单位分钟
const (
	C2B = 1
)

var marketingSourceText = map[int]string{
	C2B: "C",
}

func MarketingSourceText(code int) string {
	return marketingSourceText[code]
}

func MarketingSourceList() []int {
	return getKeys(marketingSourceText)
}

func getKeys(maps map[int]string) []int {
	keys := make([]int, 0, len(maps))
	for mapKey, _ := range maps {
		keys = append(keys, mapKey)
	}
	sort.Ints(keys)
	return keys
}
