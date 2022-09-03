package jpclouds

import (
	"io"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

func Nouns(r io.Reader) ([]string, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var nouns []string
	tokens := t.Tokenize(string(b))
	for _, token := range tokens {
		w, exist := token.FeatureAt(0)
		if !exist {
			continue
		}
		if w == "名詞" {
			nouns = append(nouns, token.Surface)
		}
	}
	return nouns, nil
}

func WordCount(words []string) (map[string]int, error) {
	counter := map[string]int{}
	for _, w := range words {
		i, exist := counter[w]
		if !exist {
			counter[w] = 1
			continue
		}
		counter[w] = i + 1
	}
	return counter, nil
}
