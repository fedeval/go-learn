package main

import (
	"fmt"
	"sort"
	"strings"
)

type SortByAge []Agent

func (a SortByAge) Len() int           { return len(a) }
func (a SortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type SortByName []Agent

func (a SortByName) Len() int           { return len(a) }
func (a SortByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByName) Less(i, j int) bool { return a[i].First < a[j].First }

func Exercise4() {
	sortableByAge := SortByAge(Agents)

	sort.Sort(sortableByAge)

	for i, a := range sortableByAge {
		fmt.Println(i+1, "-", a.First, a.Age)
	}

	sortableByName := SortByName(Agents)

	sort.Sort(sortableByName)

	for i, a := range sortableByName {
		fmt.Println(i+1, "-", a.First, a.Age)
	}

	for _, a := range Agents {
		sort.Strings(a.Sayings)
		fmt.Println(strings.ToUpper(a.First), strings.ToUpper(a.Last))
		for i, saying := range a.Sayings {
			fmt.Printf("\t %v - %s\n", i+1, saying)
		}
	}

}
