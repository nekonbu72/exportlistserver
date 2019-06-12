package main

import (
	"testing"
)

func TestNewPath(t *testing.T) {
	const url = "https://xxx.com/yyy?since=2019-06-11&before=2019-06-12"
	p := NewPath(url)
	if !p.HasQuery() {
		t.Error("HasQuery()")
	}

	if p.Query["since"][0] != "2019-06-11" {
		t.Error("since")
	}

	if p.Query["before"][0] != "2019-06-12" {
		t.Error("before")
	}

}
