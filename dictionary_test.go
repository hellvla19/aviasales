package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func Test_updateDictionary(t *testing.T) {
	ao := assert.New(t)
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.SetRequestURI("localhost:8080/load")
	ctx.Request.SetBodyString(`["foobar", "aabb", "baba", "boofar", "test"]`)
	ctx.Request.Header.SetMethod("POST")
	expectedWords := map[string][]string{
		"aabb":   {"aabb", "baba"},
		"abfoor": {"foobar", "boofar"},
		"estt":   {"test"},
	}

	updateDictionary(ctx)
	ao.Equal(expectedWords, D.Words)
}

func Test_getAnagrams(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.SetRequestURI("localhost:8080/get?word=foobar")
	ctx.Request.Header.SetMethod("GET")

	getAnagrams(ctx)
}

func Test_searchAnagrams(t *testing.T) {
	ao := assert.New(t)
	testTable := []struct {
		tcase, word      string
		dictionaryWords  map[string][]string
		expectedAnagrams []string
	}{
		{
			tcase: "success",
			word:  "foobar",
			dictionaryWords: map[string][]string{
				"aabb":   {"aabb", "baba"},
				"abfoor": {"foobar", "boofar"},
				"estt":   {"test"},
			},
			expectedAnagrams: []string{"foobar", "boofar"},
		},
		{
			tcase:            "empty dictionary",
			word:             "foobar",
			dictionaryWords:  nil,
			expectedAnagrams: nil,
		},
		{
			tcase: "empty anagrams",
			word:  "fffbar",
			dictionaryWords: map[string][]string{
				"aabb":   {"aabb", "baba"},
				"abfoor": {"foobar", "boofar"},
				"estt":   {"test"},
			},
			expectedAnagrams: nil,
		}}

	for _, testUnit := range testTable {
		D.Words = testUnit.dictionaryWords
		anagrams := searchAnagrams(testUnit.word)
		ao.Equal(testUnit.expectedAnagrams, anagrams, testUnit.tcase)
	}
}
