package main

import (
	"fmt"

	"github.com/headfirstgo/prose"
)

func main() {
	phrases := []string{"my pareteee", "a rodeo lloo"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my pareteee", "a rodeo lloo", "adfsdf"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
