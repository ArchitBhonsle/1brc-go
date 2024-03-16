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

func Naive(file *os.File) string {
	scanner := bufio.NewScanner(file)
	data := make(map[string]*StationData)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ";")

		station := split[0]
		temperature, err := strconv.ParseFloat(split[1], 64)
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

	result := "{"
	for i, station := range stations {
		stationData := data[station]

		average := stationData.sum / float64(stationData.count)
		average = math.Ceil(average*10.0) / 10.0
		result += fmt.Sprintf("%s=%.1f/%.1f/%.1f", station, stationData.min, average, stationData.max)
		if i != len(stations)-1 {
			result += ", "
		}
	}
	result += "}\n"

	return result
}

func StringBuilder(file *os.File) string {
	scanner := bufio.NewScanner(file)
	data := make(map[string]*StationData)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ";")

		station := split[0]
		temperature, err := strconv.ParseFloat(split[1], 64)
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

	result := strings.Builder{}
	result.WriteString("{")
	for i, station := range stations {
		stationData := data[station]

		average := stationData.sum / float64(stationData.count)
		average = math.Ceil(average*10.0) / 10.0
		result.WriteString(fmt.Sprintf("%s=%.1f/%.1f/%.1f", station, stationData.min, average, stationData.max))
		if i != len(stations)-1 {
			result.WriteString(", ")
		}
	}
	result.WriteString("}\n")

	return result.String()
}

func main() {
	filename := "samples/measurements-rounding.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	println(Naive(file))
}
