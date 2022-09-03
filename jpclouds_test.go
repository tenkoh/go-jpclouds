package jpclouds_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/tenkoh/go-jpclouds"
)

func TestNouns(t *testing.T) {
	r := strings.NewReader("吾輩は猫である")
	got, err := jpclouds.Nouns(r)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"吾輩", "猫"}
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
