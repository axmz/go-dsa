package main

import (
	"fmt"
)

func createAdjList(manager []int) map[int][]int {
	list := make(map[int][]int)

	for i, v := range manager {
		if v != -1 {
			list[v] = append(list[v], i)
		}
	}

	return list
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	adjList := createAdjList(manager)
	headSubordinates := adjList[headID]
	if len(headSubordinates) > 0 {
		return dfs(headID, &adjList, &informTime)
	} else {
		return 0
	}
}

func dfs(employeeId int, adjList *map[int][]int, informTime *[]int) int {
	timeToInform := (*informTime)[employeeId]
	maxSubordinatesInformTime := 0
	if subordinates, ok := (*adjList)[employeeId]; ok {
		for _, employee := range subordinates {
			subordinatesInformTime := dfs(employee, adjList, informTime)
			if subordinatesInformTime > maxSubordinatesInformTime {
				maxSubordinatesInformTime = subordinatesInformTime
			}
		}
	}

	return timeToInform + maxSubordinatesInformTime
}

func main() {

	n := 6
	headID := 2
	manager := []int{2, 2, -1, 2, 2, 2}
	informTime := []int{0, 0, 1, 0, 0, 0}

	// n := 8
	// headID := 4
	// manager := []int{2, 2, 4, 6, -1, 4, 4, 5}
	// informTime := []int{0, 0, 4, 0, 7, 3, 6, 0}

	fmt.Println(createAdjList(manager))
	res := numOfMinutes(n, headID, manager, informTime)
	fmt.Println(res)
}
