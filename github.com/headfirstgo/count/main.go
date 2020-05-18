// count tallies the number of times each line
// occur within a file

package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	/* var names []string  -- this program is with array and slices
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if line == name {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
		for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
	} */
	// simplified with map
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}

}
