package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type StationData struct {
	sum   float64
	count int
	min   float64
	max   float64
}

func main() {
	filename := "samples/measurements-rounding.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make(map[string]*StationData)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ";")

		station := split[0]
		temperature, err := strconv.ParseFloat(split[1], 32)
		if err != nil {
			panic(err)
		}

		stationData, ok := data[station]
		if !ok {
			data[station] = &StationData{temperature, 1, temperature, temperature}
			continue
		}

		stationData.sum += temperature
		stationData.count += 1
		stationData.min = math.Min(stationData.min, temperature)
		stationData.max = math.Max(stationData.max, temperature)
	}

	stations := make([]string, 0, len(data))
	for station := range data {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	print("{")
	for i, station := range stations {
		stationData := data[station]
		fmt.Printf("%s=%.1f/%.1f/%.1f", station, stationData.min, stationData.sum/float64(stationData.count), stationData.max)
		if i != len(stations)-1 {
			print(", ")
		}
	}
	print("}")
}
