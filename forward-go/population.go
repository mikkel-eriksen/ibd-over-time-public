package main

func getBackwardPopulation(pop_size, num_gen int, chrom_lengths []int, recomb_rate float64) [][]Individual {
	chrom_break_pos := getChromBreakPos(chrom_lengths)
	pop := make([][]Individual, num_gen)

	for i := 0; i < num_gen-1; i++ {
		for j := 0; j < pop_size; j++ {
			pop[i] = append(pop[i], Individual{})
			pop[i][j].setParents(pop_size)
			pop[i][j].setBreakPos(chrom_break_pos, recomb_rate)
		}
	}

	for j := 0; j < pop_size; j++ {
		pop[num_gen-1] = append(pop[num_gen-1], Individual{})
	}

	return pop
}

func getForwardPopulation(pop_size, num_gen int, chrom_lengths []int, recomb_rate float64) [][]Individual {
	chrom_break_pos := getChromBreakPos(chrom_lengths)
	pop := make([][]Individual, num_gen)

	for j := 0; j < pop_size; j++ {
		pop[0] = append(pop[0], Individual{})
	}

	for i := 1; i < num_gen; i++ {
		for j := 0; j < pop_size; j++ {
			pop[i] = append(pop[i], Individual{})
			mom_id, dad_id := pop[i][j].setParents(pop_size)
			pop[i][j].setBreakPos(chrom_break_pos, recomb_rate)
			pop[i-1][mom_id].children = append(pop[i-1][mom_id].children, j)
			pop[i-1][dad_id].children = append(pop[i-1][dad_id].children, j)
		}
	}

	return pop
}
