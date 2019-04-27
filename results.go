package main

import (
	"regexp"
	"strings"
)

var (
	lineRegexp = regexp.MustCompile(`(il1|il2|dl1|dl2|ul1|ul2)\.(\w+)\s+([\d\.]+)\s`)
)

type BenchResults struct {
	Name          string `json:"name"`
	Accesses      string `json:"accesses"`
	Hits          string `json:"hits"`
	Misses        string `json:"misses"`
	Replacements  string `json:"replacements"`
	Writebacks    string `json:"writebacks"`
	Invalidations string `json:"invalidations"`
	MissRate      string `json:"miss_rate"`
	ReplRate      string `json:"repl_rate"`
	WbRate        string `json:"wb_rate"`
	InvRate       string `json:"inv_rate"`
}

func (b *BenchResults) setFieldValue(field, value string) {
	switch field {
	case "name":
		b.Name = value
	case "accesses":
		b.Accesses = value
	case "hits":
		b.Hits = value
	case "misses":
		b.Misses = value
	case "replacements":
		b.Replacements = value
	case "writebacks":
		b.Writebacks = value
	case "invalidations":
		b.Invalidations = value
	case "miss_rate":
		b.MissRate = value
	case "repl_rate":
		b.ReplRate = value
	case "wb_rate":
		b.WbRate = value
	case "inv_rate":
		b.InvRate = value
	}
}

func parseResults(results string) []*BenchResults {
	il1Results := BenchResults{Name: il1}
	il2Results := BenchResults{Name: il2}
	dl1Results := BenchResults{Name: dl1}
	dl2Results := BenchResults{Name: dl2}

	for _, line := range strings.Split(results, "\n") {
		cacheName, field, value := processLine(line)
		switch cacheName {
		case il1:
			il1Results.setFieldValue(field, value)
		case il2:
			il2Results.setFieldValue(field, value)
		case dl1:
			dl1Results.setFieldValue(field, value)
		case dl2:
			dl2Results.setFieldValue(field, value)

		// sim-cache unified caches points instructions to the data cache, so
		// results end up in the data cache
		case ul1:
			dl1Results.setFieldValue(field, value)
		case ul2:
			dl2Results.setFieldValue(field, value)
		}
	}

	return rejectResults(
		[]*BenchResults{&il1Results, &il2Results, &dl1Results, &dl2Results},
		func(result *BenchResults) bool {
			return result.Accesses == result.Hits && result.Hits == ""
		})
}

func processLine(line string) (string, string, string) {
	matched := lineRegexp.FindAllStringSubmatch(line, -1)
	if len(matched) == 0 {
		return "", "", ""
	}

	return matched[0][1], matched[0][2], matched[0][3]
}

func rejectResults(results []*BenchResults, rejectFn func(*BenchResults) bool) []*BenchResults {
	filteredResults := []*BenchResults{}
	for _, result := range results {
		if rejectFn(result) {
			continue
		}

		filteredResults = append(filteredResults, result)
	}
	return filteredResults
}
