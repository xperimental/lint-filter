package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

const (
	FilterResultInclude = true
	FilterResultExclude = false
)

type FilterResult bool

type LocationMatcher struct {
	PathPrefix string
}

type LocationFilters struct {
	Includes []LocationMatcher `json:"includes"`
	Excludes []LocationMatcher `json:"excludes"`
}

type FilterConfig struct {
	DefaultResult   FilterResult    `json:"default"`
	LocationFilters LocationFilters `json:"locationFilters"`
}

func (m *LocationMatcher) Matches(l *Location) bool {
	return strings.HasPrefix(l.File, m.PathPrefix)
}

func (c *FilterConfig) Filter(filtered Issue) (Issue, error) {
	c.filterLocations(&filtered)
	if len(filtered.Locations) == 0 {
		return filtered, errors.New("No locations left.")
	}
	return filtered, nil
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

func (c *FilterConfig) filterLocations(i *Issue) {
	filtered := make([]Location, 0, len(i.Locations))
	for _, l := range i.Locations {
		location := c.filterLocation(&l)
		if location != nil {
			filtered = append(filtered, *location)
		}
	}
	i.Locations = filtered
}

func (c *FilterConfig) filterLocation(l *Location) *Location {
	for _, f := range c.LocationFilters.Includes {
		if f.Matches(l) {
			return l
		}
	}
	for _, f := range c.LocationFilters.Excludes {
		if f.Matches(l) {
			return nil
		}
	}
	if c.DefaultResult {
		return l
	}
	return nil
}
