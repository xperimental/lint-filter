package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

const (
	FilterResultInclude = true
	FilterResultExclude = false
)

type FilterResult bool

type IssueMatcher struct {
	PathPrefix string
}

type FilterConfig struct {
	DefaultResult FilterResult   `json:"default"`
	Includes      []IssueMatcher `json:"includes"`
	Excludes      []IssueMatcher `json:"excludes"`
}

func (m *IssueMatcher) Matches(i *Issue) bool {
	for _, l := range i.Location {
		if strings.HasPrefix(l.File, m.PathPrefix) {
			return true
		}
	}
	return false
}

func (c *FilterConfig) Filter(i *Issue) FilterResult {
	for _, m := range c.Includes {
		if m.Matches(i) {
			return true
		}
	}
	for _, m := range c.Excludes {
		if m.Matches(i) {
			return false
		}
	}
	return c.DefaultResult
}

func (c *FilterConfig) Read(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, c)
	if err != nil {
		return err
	}
	return nil
}
