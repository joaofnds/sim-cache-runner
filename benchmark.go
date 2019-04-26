package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

type Entry struct {
	Config  CacheConfig
	Results []*BenchResults
}

type Benchmark struct {
	Entries []*Entry
	Exec    string
	Args    string
}

func (b *Benchmark) Run() {
	var wg sync.WaitGroup
	wg.Add(len(b.Entries))
	for _, entry := range b.Entries {
		go (func(e *Entry) {
			out, err := b.runConfig(e.Config)
			if err != nil {
				log.Fatalf("failed to run bench: %v\n", err)
			}
			e.Results = parseResults(string(out))
			wg.Done()
		})(entry)
	}
	wg.Wait()
}

func (b *Benchmark) runConfig(config CacheConfig) ([]byte, error) {
	cmd := exec.Command(
		"./sim-cache",
		"-cache:il1", config.Il1.ParamStr,
		"-cache:il2", config.Il2.ParamStr,
		"-cache:dl1", config.Dl1.ParamStr,
		"-cache:dl2", config.Dl2.ParamStr,
		"-tlb:itlb", config.Itlb.ParamStr,
		"-tlb:dtlb", config.Dtlb.ParamStr,
		b.Exec, b.Args,
	)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, nil
}

func (b *Benchmark) ShowResults() {
	for _, entry := range b.Entries {
		fmt.Println("-------------- start entry ------------")
		for _, result := range entry.Results {
			cache := entry.Config.getCacheByName(result.Name)
			fmt.Printf("-------------- %s ------------\n", result.Name)
			fmt.Println(cache.Name)
			fmt.Println(cache.Sets)
			fmt.Println(cache.BlockSize)
			fmt.Println(cache.Assoc)
			fmt.Println(cache.Repl)
			fmt.Println(result.Accesses)
			fmt.Println(result.Hits)
			fmt.Println(result.Misses)
			fmt.Println(result.Replacements)
			fmt.Println(result.Writebacks)
			fmt.Println(result.Invalidations)
			fmt.Println(result.MissRate)
			fmt.Println(result.ReplRate)
			fmt.Println(result.WbRate)
			fmt.Println(result.InvRate)
			fmt.Println("-------------------------------")
		}
		fmt.Println("--------------- end entry -------------")
	}
}
