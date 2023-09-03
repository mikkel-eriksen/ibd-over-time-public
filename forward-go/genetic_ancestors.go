package main

func getGeneticAncestors(gen_minus int, pop [][]Individual, chrom_lengths []int) float64 {
	pop_size := len(pop[0])

	// Run recombination simulation
	prev_gen := make([]ChromPair, pop_size)
	for j := 0; j < pop_size; j++ {
		for k := 0; k < len(chrom_lengths); k++ {
			prev_gen[j].chrom_pair[0].addSegment(j, chrom_lengths[k])
			prev_gen[j].chrom_pair[1].addSegment(j, chrom_lengths[k])
		}
	}

	for i := gen_minus - 1; i > -1; i-- {
		curr_gen := make([]ChromPair, pop_size)
		for j := 0; j < pop_size; j++ {
			curr_gen[j].chrom_pair[0] = prev_gen[pop[i][j].mom_id].getChrom(pop[i][j].mom_break_pos)
			curr_gen[j].chrom_pair[1] = prev_gen[pop[i][j].dad_id].getChrom(pop[i][j].dad_break_pos)
		}
		prev_gen = curr_gen
	}

	// Count number of ancestors
	num_acestors := make([]int, pop_size)
	for j := 0; j < pop_size; j++ {
		ancestors := map[int]struct{}{}
		for k := 0; k < len(prev_gen[j].chrom_pair[0].ids); k++ {
			ancestors[prev_gen[j].chrom_pair[0].ids[k]] = struct{}{}
		}
		for k := 0; k < len(prev_gen[j].chrom_pair[1].ids); k++ {
			ancestors[prev_gen[j].chrom_pair[1].ids[k]] = struct{}{}
		}
		num_acestors[j] = len(ancestors)
	}

	return mean(num_acestors)

}
