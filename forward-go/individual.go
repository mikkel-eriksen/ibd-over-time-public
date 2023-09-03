package main

import (
	"math/rand"
	"sort"

	"gonum.org/v1/gonum/stat/distuv"
)

type Individual struct {
	mom_id        int
	dad_id        int
	mom_break_pos []int
	dad_break_pos []int
	children      []int
}

func (ind *Individual) setParents(pop_size int) (int, int) {
	ind.mom_id = rand.Intn(pop_size)
	for {
		ind.dad_id = rand.Intn(pop_size)
		if ind.dad_id != ind.mom_id {
			break
		}
	}
	return ind.mom_id, ind.dad_id
}

func (ind *Individual) setBreakPos(chrom_break_pos []int, recomb_rate float64) {
	seq_len := chrom_break_pos[len(chrom_break_pos)-1]

	dist := distuv.Binomial{N: float64(seq_len), P: recomb_rate, Src: nil}

	for i := 0; i < 2; i++ {
		break_pos := []int{}
		for j := 0; j < len(chrom_break_pos)-1; j++ {
			if rand.Intn(2) == 1 {
				break_pos = append(break_pos, chrom_break_pos[j])
			}
		}

		num_recomb := int(dist.Rand())
		for j := 0; j < num_recomb; j++ {
			break_pos = append(break_pos, rand.Intn(seq_len-2)+1)
		}

		sort.Ints(break_pos)
		break_pos = append(break_pos, seq_len)

		if i == 0 {
			ind.mom_break_pos = break_pos
		} else {
			ind.dad_break_pos = break_pos
		}
	}
}
