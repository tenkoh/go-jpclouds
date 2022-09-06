package jpclouds_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/tenkoh/go-jpclouds"
)

func TestPreprocess(t *testing.T) {
	r := strings.NewReader("吾輩は cat である。ﾅﾏｴﾊﾏﾀﾞﾅｲ。ａｂｃ。")
	got, err := jpclouds.Preprocess(r)
	if err != nil {
		t.Fatal(err)
	}
	want := "吾輩はCATである。ナマエハマダナイ。ABC。"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestCollectWords(t *testing.T) {
	s := "私はきれいな雪原に出かけた"
	got, err := jpclouds.CollectWords(s, jpclouds.Noun, jpclouds.Verb, jpclouds.Adjective)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"私", "きれい", "雪原", "出かける"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %+v, got %+v", want, got)
	}
}

func TestWordCount(t *testing.T) {
	words := []string{"golang", "gopher", "golang"}
	want := map[string]int{
		"golang": 2,
		"gopher": 1,
	}
	got, _ := jpclouds.WordCount(words)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %+v, got %+v", want, got)
	}
}
