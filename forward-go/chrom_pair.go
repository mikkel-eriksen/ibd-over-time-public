package main

import "sort"

type ChromPair struct {
	chrom_pair [2]Sequence
}

func (cp ChromPair) getChrom(break_pos []int) Sequence {
	chrom_index := 0 // rand.Intn(2)
	seq := cp.chrom_pair[chrom_index].getSequence(0, break_pos[0])
	chrom_index = 1 - chrom_index

	for i := 0; i < len(break_pos)-1; i++ {
		if break_pos[i] != break_pos[i+1] {
			seq.extend(cp.chrom_pair[chrom_index].getSequence(break_pos[i], break_pos[i+1]))
			chrom_index = 1 - chrom_index
		}
	}

	return seq
}

func (cp ChromPair) getROH(seq_len int) (int, int) {
	// Merge break positions from both chromosome sequences
	break_pos := []int{0}
	break_pos = append(break_pos, cp.chrom_pair[0].lengths[0])
	for k := 1; k < len(cp.chrom_pair[0].lengths); k++ {
		break_pos = append(
			break_pos,
			break_pos[k-1]+cp.chrom_pair[0].lengths[k],
		)
	}
	break_pos = append(break_pos, cp.chrom_pair[1].lengths[0])
	for k := 1; k < len(cp.chrom_pair[1].lengths); k++ {
		break_pos = append(
			break_pos,
			break_pos[k-1]+cp.chrom_pair[1].lengths[k],
		)
	}
	sort.Ints(break_pos)
	break_pos = append(break_pos, seq_len)

	// Find ROH for each block
	roh_length := 0
	roh_count := 0
	for k := 0; k < len(break_pos)-1; k++ {
		if break_pos[k] != break_pos[k+1] {
			if cp.chrom_pair[0].getSequence(break_pos[k], break_pos[k+1]).ids[0] ==
				cp.chrom_pair[1].getSequence(break_pos[k], break_pos[k+1]).ids[0] {
				roh_length += break_pos[k+1] - break_pos[k]
				roh_count++
			}
		}
	}
	return roh_length, roh_count
}
