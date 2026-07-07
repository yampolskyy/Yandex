package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var (
	mp     = make(map[string]int)
	wg     sync.WaitGroup
	mx     sync.Mutex
	sm     int
	client http.Client
)

func BestStudents(names []string) (string, error) {
	for _, name := range names {
		wg.Add(1)

		go func(name string) {
			defer wg.Done()

			url := fmt.Sprintf("http://localhost:8082/mark?name=%s", name)
			res, err := client.Get(url)
			if err != nil {
				mx.Lock()
				mp[name] = 0
				mx.Unlock()
				return
			}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				mx.Lock()
				mp[name] = 0
				mx.Unlock()
				return
			}
			defer res.Body.Close()

			val, err := strconv.Atoi(string(body))
			mx.Lock()
			mp[name] = val
			mx.Unlock()
		}(name)
	}

	wg.Wait()

	if len(mp) == 0 {
		return "", fmt.Errorf("no data")
	}

	for _, m := range mp {
		if m == 0 {
			return "", fmt.Errorf("no data for name")
		}
		sm += m
	}

	x := sm / len(mp)
	answer := make([]string, len(mp))

	for n, m := range mp {
		if m > x {
			answer = append(answer, n)
		}
	}

	sort.Slice(answer, func(i int, j int) bool {
		return answer[i] < answer[j]
	})

	return strings.Join(answer, ","), nil
}
