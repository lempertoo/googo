package main

import (
	"sort"
	"sync"
)

func Scan(host string, ports []int) []int {

	var wg sync.WaitGroup

	jobs := make(chan int)

	results := make(chan int)

	for i := 0; i < WorkerCount; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for port := range jobs {

				if ScanPort(host, port) {
					results <- port
				}

			}

		}()

	}

	go func() {

		for _, port := range ports {
			jobs <- port
		}

		close(jobs)

		wg.Wait()

		close(results)

	}()

	var open []int

	for port := range results {
		open = append(open, port)
	}

	sort.Ints(open)

	return open

}
