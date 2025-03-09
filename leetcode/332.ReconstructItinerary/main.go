package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	flightsCount := len(tickets)

	flights := make(map[string][][]string)
	for _, t := range tickets {
		from := t[0]
		flights[from] = append(flights[from], t)
	}

	for from := range flights {
		sort.Slice(flights[from], func(i, j int) bool {
			return flights[from][i][1] < flights[from][j][1]
		})
	}

	itinerary := []string{}
	var fly func(from string, flightsCount int)
	fly = func(from string, flightsCount int) {
		for len(flights[from]) > 0 {
			to := flights[from][0][1]
			flights[from] = flights[from][1:]
			fly(to, flightsCount-1)
		}
		itinerary = append([]string{from}, itinerary...)
	}

	fly("JFK", flightsCount)

	return itinerary
}

func main() {
	// tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	// tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	tickets := [][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}}
	// tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "JFK"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}, {"ATL", "AAA"}, {"AAA", "BBB"}, {"BBB", "ATL"}}
	// tickets := [][]string{ {"AXA", "EZE"}, {"EZE", "AUA"}, {"ADL", "JFK"}, {"ADL", "TIA"}, {"AUA", "AXA"}, {"EZE", "TIA"}, {"EZE", "TIA"}, {"AXA", "EZE"}, {"EZE", "ADL"}, {"ANU", "EZE"}, {"TIA", "EZE"}, {"JFK", "ADL"}, {"AUA", "JFK"}, {"JFK", "EZE"}, {"EZE", "ANU"}, {"ADL", "AUA"}, {"ANU", "AXA"}, {"AXA", "ADL"}, {"AUA", "JFK"}, {"EZE", "ADL"}, {"ANU", "TIA"}, {"AUA", "JFK"}, {"TIA", "JFK"}, {"EZE", "AUA"}, {"AXA", "EZE"}, {"AUA", "ANU"}, {"ADL", "AXA"}, {"EZE", "ADL"}, {"AUA", "ANU"}, {"AXA", "EZE"}, {"TIA", "AUA"}, {"AXA", "EZE"}, {"AUA", "SYD"}, {"ADL", "JFK"}, {"EZE", "AUA"}, {"ADL", "ANU"}, {"AUA", "TIA"}, {"ADL", "EZE"}, {"TIA", "JFK"}, {"AXA", "ANU"}, {"JFK", "AXA"}, {"JFK", "ADL"}, {"ADL", "EZE"}, {"AXA", "TIA"}, {"JFK", "AUA"}, {"ADL", "EZE"}, {"JFK", "ADL"}, {"ADL", "AXA"}, {"TIA", "AUA"}, {"AXA", "JFK"}, {"ADL", "AUA"}, {"TIA", "JFK"}, {"JFK", "ADL"}, {"JFK", "ADL"}, {"ANU", "AXA"}, {"TIA", "AXA"}, {"EZE", "JFK"}, {"EZE", "AXA"}, {"ADL", "TIA"}, {"JFK", "AUA"}, {"TIA", "EZE"}, {"EZE", "ADL"}, {"JFK", "ANU"}, {"TIA", "AUA"}, {"EZE", "ADL"}, {"ADL", "JFK"}, {"ANU", "AXA"}, {"AUA", "AXA"}, {"ANU", "EZE"}, {"ADL", "AXA"}, {"ANU", "AXA"}, {"TIA", "ADL"}, {"JFK", "ADL"}, {"JFK", "TIA"}, {"AUA", "ADL"}, {"AUA", "TIA"}, {"TIA", "JFK"}, {"EZE", "JFK"}, {"AUA", "ADL"}, {"ADL", "AUA"}, {"EZE", "ANU"}, {"ADL", "ANU"}, {"AUA", "AXA"}, {"AXA", "TIA"}, {"AXA", "TIA"}, {"ADL", "AXA"}, {"EZE", "AXA"}, {"AXA", "JFK"}, {"JFK", "AUA"}, {"ANU", "ADL"}, {"AXA", "TIA"}, {"ANU", "AUA"}, {"JFK", "EZE"}, {"AXA", "ADL"}, {"TIA", "EZE"}, {"JFK", "AXA"}, {"AXA", "ADL"}, {"EZE", "AUA"}, {"AXA", "ANU"}, {"ADL", "EZE"}, {"AUA", "EZE"}}

	fmt.Println(findItinerary(tickets))
}
