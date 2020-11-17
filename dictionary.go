package main

import (
	"aviasales/util"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"strings"
)

type Dictionary struct {
	Words []string
}

var D = Dictionary{}

// updateDictionary обновляет наш словарь.
func updateDictionary(ctx *fasthttp.RequestCtx) {
	var w []string
	if err := json.Unmarshal(ctx.PostBody(), &w); err != nil {
		ctx.WriteString("Failed to load dictionary!")
		log.Fatal(err)
	}
	D.Words = w
	if len(D.Words) != 0 {
		ctx.WriteString("Success update dictionary!")
	}
}

// getAnagrams получение анаграмм по заданному слову.
func getAnagrams(ctx *fasthttp.RequestCtx) {
	word := ctx.QueryArgs().Peek("word")
	anagrams := searchAnagrams(string(word))
	ctx.WriteString(fmt.Sprint(anagrams))
}

// searchAnagrams сопоставление переданного слова с имеющимися в словаре для поиска анаграмм.
// Два слова считаются анаграммами, если одно можно получить из другого перестановкой букв (без учета регистра).
func searchAnagrams(word string) []string {
	if D.Words == nil {
		return nil
	}
	var words []string
	preparedWord := util.SortString(strings.ToLower(strings.TrimSpace(word)))
	for _, w := range D.Words {
		preparedW := util.SortString(strings.ToLower(strings.TrimSpace(w)))
		if preparedWord == preparedW {
			words = append(words, w)
		}
	}
	return words
}
