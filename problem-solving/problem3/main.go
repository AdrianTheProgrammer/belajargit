package main

import (
	"fmt"
	"sort"
)

func main() {
	dragonknight([]int{5, 4}, []int{7, 8, 4})
	dragonknight([]int{5, 10}, []int{5})
	dragonknight([]int{7, 2}, []int{4, 3, 1, 2})
	dragonknight([]int{7, 2}, []int{2, 1, 8, 5})
}

func dragonknight(dragon, knight []int) {
	sort.Slice(dragon, func(i, j int) bool {
		return dragon[i] < dragon[j]
	})
	sort.Slice(knight, func(i, j int) bool {
		return knight[i] < knight[j]
	})

	var drag_pos, knig_pos, hasil int = 0, 0, 0
	var knig_surv []int

	for drag_pos < len(dragon) && knig_pos < len(knight) {
		if dragon[drag_pos] <= knight[knig_pos] {
			knig_surv = append(knig_surv, knight[knig_pos])
			knig_pos++
			drag_pos++
		} else {
			knig_pos++
		}
	}

	if len(knig_surv) == 0 || len(knig_surv) < len(dragon) {
		fmt.Println("Knight Fall")
	} else {
		for i := 0; i < len(knig_surv); i++ {
			hasil += knig_surv[i]
		}
		fmt.Println(hasil)
	}
}
