package main

import (
	"testing"
)

func TestNewPath(t *testing.T) {
	const url = "https://xxx.com?a=AAA&b=BBB&c=CCC&d=DDD&a=EEE"
	p := NewPath(url)
	if !p.HasQuery() {
		t.Error()
	}

	if p.Query["a"][0] != "AAA" {
		t.Error()
	}

	if p.Query["b"][0] != "BBB" {
		t.Error()
	}

	if p.Query["a"][1] != "EEE" {
		t.Error()
	}
}
