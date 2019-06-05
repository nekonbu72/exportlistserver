package main

import "net/url"

// Path represents the path of a request.
type Path struct {
	Path string
	// Query map[string][]string
	Query url.Values
}

// NewPath makes a new Path from the specified
// path string.
func NewPath(p string) *Path {
	u, _ := url.Parse(p)
	query := u.Query()
	return &Path{Path: p, Query: query}
}

// HasID gets whether this path has an ID
// or not.
func (p *Path) HasQuery() bool {
	return len(p.Query) > 0
}
