package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose messages")

type size struct {
	idx  int
	size int64
}

type result struct {
	dir            string
	nfiles, nbytes int64
}

func main() {
	// Determine the initial directories
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree
	fileSizes := make(chan size)
	var wg sync.WaitGroup
	for i, root := range roots {
		wg.Add(1)
		go walkDir(i, root, fileSizes, &wg)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// Print results periodically
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	results := make([]result, len(roots))
	for i, root := range roots {
		results[i].dir = root
	}
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			results[size.idx].nfiles++
			results[size.idx].nbytes += size.size
		case <-tick:
			printDistUsage(results)
		}
	}
	printDistUsage(results) // final totals
}

func walkDir(i int, dir string, fileSizes chan<- size, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(i, subdir, fileSizes, wg)
		} else {
			fileSizes <- size{i, entry.Size()}
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func printDistUsage(results []result) {
	for _, result := range results {
		fmt.Printf("%s: %d files, %.1f GB\n", result.dir, result.nfiles, float64(result.nbytes)/1e9)
	}
}
