// Baseline file for 1brc

package main

import (
	"bufio"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

func baseline(inputFilepath string) error {
	type stats struct {
		min, max, sum float64
		count         int64
	}

	f, err := os.Open(inputFilepath)
	if err != nil {
		slog.Error("Cannot open input data file!", "ERROR", err)
		return err
	}
	defer f.Close()

	stationStats := make(map[string]stats)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// before, after, found := strings.Cut(s, sep)
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}

		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			return err
		}

		s, ok := stationStats[station]
		if !ok {
			s.min = temp
			s.max = temp
			s.sum = temp
			s.count = 1
		} else {
			s.min = min(s.min, temp)
			s.max = max(s.max, temp)
			s.sum += temp
			s.count++
		}
		stationStats[station] = s
	}

	stations := make([]string, 0, len(stationStats))
	for station := range stationStats {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	// fmt.Fprint(output, "{")
	for _, station := range stations {
		s := stationStats[station]
		mean := s.sum / float64(s.count)
		slog.Debug("Station Stats:", "Station", station, "min", s.min, "mean", mean, "max", s.max)
	}

	// fmt.Fprint(output, "}\n")
	return nil

}
