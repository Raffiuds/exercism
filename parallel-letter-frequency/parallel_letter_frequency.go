package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

func (f FreqMap) merge(m FreqMap) {
	for k, v := range m {
		f[k] += v
	}
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(texts []string) FreqMap {

	var rwm sync.RWMutex
	var wg sync.WaitGroup

	freqMaps := []FreqMap{}

	for _, text := range texts {

		wg.Add(1)
		go func(t string) {

			rwm.Lock()
			freqMaps = append(freqMaps, Frequency(t))
			rwm.Unlock()

			wg.Done()
		}(text)

	}

	wg.Wait()

	return margeMaps(freqMaps)
}

func margeMaps(freqmaps []FreqMap) FreqMap {

	m := FreqMap{}

	for _, freqmap := range freqmaps {
		for k, value := range freqmap {
			if v, ok := m[k]; ok {
				value += v
			}
			m[k] = value
		}
	}
	return m

}
