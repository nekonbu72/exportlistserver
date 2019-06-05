package main

import (
	"testing"
)

func TestNewPath(t *testing.T) {
	const url string = "https://xxx.com?a=AAA&b=BBB&c=CCC&d=DDD&a=EEE"
	// const url string = "https://xxx.com"

	p := NewPath(url)
	if !p.HasQuery() {
		t.Error()
	}
	if p.Query["a"][0] != "AAA" {
		t.Error()
	}
	if p.Query["a"][1] != "EEE" {
		t.Error()
	}
}
