package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	white = iota
	gray
	black
)

type group []int

func (left group) greater(right group) bool {
	if len(left) != len(right) {
		return len(left) > len(right)
	}

	minValue := math.MaxInt
	isLeftGreater := false
	for i := 0; i < len(left); i++ {
		if left[i] < minValue {
			minValue = left[i]
			isLeftGreater = true
		}

		if right[i] < minValue {
			minValue = right[i]
			isLeftGreater = false
		}
	}

	return isLeftGreater
}

func main() {
	in := bufio.NewReader(os.Stdin)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var clientsCount, attributesCount, eventsCount, collisionsCount int
	fmt.Fscan(in, &clientsCount, &attributesCount, &eventsCount, &collisionsCount)

	attributesByClient := getAttributesByClients(
		in,
		clientsCount, attributesCount, eventsCount,
	)

	graph := buildGraph(
		attributesByClient,
		clientsCount, attributesCount, collisionsCount,
	)

	groups := getGroups(graph)

	result := selectGroup(groups)

	for i, client := range result {
		if i != 0 {
			writer.WriteString(" ")
		}

		writer.WriteString(strconv.Itoa(client))
	}
	writer.WriteString("\n")

}

func getAttributesByClients(
	in *bufio.Reader,
	clientsCount, attributesCount, eventsCount int,
) []map[int]int {
	attributesByClient := make([]map[int]int, clientsCount)

	for eventIdx := 0; eventIdx < eventsCount; eventIdx++ {
		var clientId, attributeId, attributeValue int
		fmt.Fscan(in, &clientId, &attributeId, &attributeValue)

		if attributesByClient[clientId] == nil {
			attributesByClient[clientId] = make(map[int]int, attributesCount)
		}

		attributesByClient[clientId][attributeId] = attributeValue
	}

	return attributesByClient
}

func buildGraph(
	attributesByClient []map[int]int,
	clientsCount, attributesCount, collisionsCount int,
) [][]int {
	graph := make([][]int, clientsCount)

	for leftClientId, leftAttributes := range attributesByClient {
		for rightClientId, rightAttributes := range attributesByClient {
			if leftClientId == rightClientId {
				continue
			}

			matchingAttributes := 0
			for attribute, leftValue := range leftAttributes {
				if rightValue, ok := rightAttributes[attribute]; ok && leftValue == rightValue {
					matchingAttributes++
				}

				if matchingAttributes == collisionsCount {
					break
				}
			}

			if matchingAttributes == collisionsCount {
				graph[leftClientId] = append(graph[leftClientId], rightClientId)
			}
		}
	}

	return graph
}

func getGroups(graph [][]int) []group {

	color := make([]int, len(graph))

	groups := make([]group, 0, len(graph))
	groupIdx := 0
	groups = append(groups, group{})
	for start := range graph {
		if color[start] != white {
			continue
		}

		queue := []int{start}
		color[start] = gray

		for i := 0; i < len(queue); i++ {
			curNode := queue[i]

			for _, neighbour := range graph[curNode] {
				if color[neighbour] == white {
					queue = append(queue, neighbour)
					color[neighbour] = gray
				}
			}

			groups[groupIdx] = append(groups[groupIdx], curNode)
			color[curNode] = black
		}

		groupIdx++
		groups = append(groups, group{})
	}

	return groups
}

func selectGroup(groups []group) group {
	maxGroup := group{}

	for _, group := range groups {
		if group.greater(maxGroup) {
			maxGroup = group
		}
	}

	sort.Slice(maxGroup, func(i, j int) bool {
		return maxGroup[i] < maxGroup[j]
	})

	return maxGroup
}
