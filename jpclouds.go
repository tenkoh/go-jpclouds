package jpclouds

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type SpeechPart int

func (sp SpeechPart) String() string {
	return [...]string{"名詞", "動詞", "形容詞"}[sp]
}

const (
	Noun SpeechPart = iota
	Verb
	Adjective
)

func Preprocess(r io.Reader) (string, error) {
	// unicode的表記ゆれを除去する
	r = norm.NFKC.Reader(r)

	b, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("error occured in preprocessing")
	}

	// 英数字は半角、カタカナは全角で統一する
	b = width.Fold.Bytes(b)

	s := string(b)

	// 英数字を大文字で統一する
	s = strings.ToUpper(s)

	// remove spaces to detect multi-words english word as a one-word.
	s = strings.ReplaceAll(s, " ", "")

	return s, nil
}

func CollectWords(s string, targets ...SpeechPart) ([]string, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}
	var words []string
	targetMap := map[string]struct{}{}
	for _, target := range targets {
		targetMap[target.String()] = struct{}{}
	}
	tokens := t.Tokenize(s)
	for _, token := range tokens {
		sp, exist := token.FeatureAt(0)
		if !exist {
			continue
		}
		if _, exist = targetMap[sp]; exist {
			base, exist := token.BaseForm()
			if exist {
				words = append(words, base)
			} else {
				words = append(words, token.Surface)
			}
		}
	}
	return words, nil
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
