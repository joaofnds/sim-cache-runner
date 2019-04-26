package main

import "regexp"

const (
	il1  = "il1"
	il2  = "il2"
	dl1  = "dl1"
	dl2  = "dl2"
	itlb = "itlb"
	dtlb = "dtlb"
)

var (
	configRegexp = regexp.MustCompile(`(il1|il2|dl1|dl2):(\d+):(\d+):(\d):([l|f|r])`)
)

// <name>:<nsets>:<bsize>:<assoc>:<repl>
type Cache struct {
	Name      string `json:"name"`
	Sets      string `json:"nsets"`
	BlockSize string `json:"bsize"`
	Assoc     string `json:"assoc"`
	Repl      string `json:"repl"`
	ParamStr  string
}

type CacheConfig struct {
	Il1  Cache `json:"il1"`
	Il2  Cache `json:"il2"`
	Dl1  Cache `json:"dl1"`
	Dl2  Cache `json:"dl2"`
	Itlb Cache `json:"itlb"`
	Dtlb Cache `json:"dtlb"`
}

func newCacheConfig(il1, il2, dl1, dl2, itlb, dtlb string) CacheConfig {
	newConfig := CacheConfig{}
	if il1 == "" {
		newConfig.Il1 = Cache{ParamStr: "none"}
	} else {
		newConfig.Il1 = buildCache(il1)
	}

	if il2 == "" {
		newConfig.Il2 = Cache{ParamStr: "none"}
	} else {
		newConfig.Il2 = buildCache(il2)
	}

	if dl1 == "" {
		newConfig.Dl1 = Cache{ParamStr: "none"}
	} else {
		newConfig.Dl1 = buildCache(dl1)
	}

	if dl2 == "" {
		newConfig.Dl2 = Cache{ParamStr: "none"}
	} else {
		newConfig.Dl2 = buildCache(dl2)
	}

	if itlb == "" {
		newConfig.Itlb = Cache{ParamStr: "none"}
	} else {
		newConfig.Itlb = buildCache(il1)
	}

	if dtlb == "" {
		newConfig.Dtlb = Cache{ParamStr: "none"}
	} else {
		newConfig.Dtlb = buildCache(dtlb)
	}

	return newConfig
}

func (c *CacheConfig) getCacheByName(name string) *Cache {
	switch name {
	case il1:
		return &c.Il1
	case il2:
		return &c.Il2
	case dl1:
		return &c.Dl1
	case dl2:
		return &c.Dl2
	case itlb:
		return &c.Itlb
	case dtlb:
		return &c.Dtlb
	default:
		return nil
	}
}

func buildCache(config string) Cache {
	name, sets, blockSize, assoc, repl := processConfigStr(config)
	return Cache{
		Name:      name,
		ParamStr:  config,
		Sets:      sets,
		BlockSize: blockSize,
		Assoc:     assoc,
		Repl:      repl,
	}
}

func processConfigStr(config string) (string, string, string, string, string) {
	matched := configRegexp.FindAllStringSubmatch(config, -1)
	if len(matched) == 0 {
		return "", "", "", "", ""
	}

	return matched[0][1], matched[0][2], matched[0][3], matched[0][4], matched[0][5]
}
