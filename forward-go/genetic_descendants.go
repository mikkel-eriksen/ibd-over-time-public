package main

func getGeneticDescendants(pop [][]Individual, chrom_lengths []int) ([]float64, []float64, []int, []float64, []int) {
	num_gen := len(pop)
	pop_size := len(pop[0])
	seq_len := sum(chrom_lengths)

	mean_num_descendants := make([]float64, num_gen)
	mean_num_descendants[0] = 1

	mean_segment_count := make([]float64, num_gen)
	mean_segment_count[0] = float64(len(chrom_lengths))

	segment_len := make([]int, num_gen)
	segment_len[0] = int(mean(chrom_lengths))

	roh_freq := make([]float64, num_gen)
	roh_freq[0] = 0
	mean_roh_len := make([]int, num_gen)
	mean_roh_len[0] = 0

	// Run recombination simulation
	prev_gen := make([]ChromPair, pop_size)
	for j := 0; j < pop_size; j++ {
		for k := 0; k < len(chrom_lengths); k++ {
			prev_gen[j].chrom_pair[0].addSegment(j, chrom_lengths[k])
			prev_gen[j].chrom_pair[1].addSegment(j, chrom_lengths[k])
		}
	}

	for i := 1; i < num_gen; i++ {
		curr_gen := make([]ChromPair, pop_size)
		num_descendants := make([]int, pop_size)
		num_segments := 0
		roh_lengths := make([]int, pop_size)
		roh_counts := make([]int, pop_size)
		for j := 0; j < pop_size; j++ {
			curr_gen[j].chrom_pair[0] = prev_gen[pop[i][j].mom_id].getChrom(pop[i][j].mom_break_pos)
			curr_gen[j].chrom_pair[1] = prev_gen[pop[i][j].dad_id].getChrom(pop[i][j].dad_break_pos)

			// Get number of descendants
			ancestors := map[int]struct{}{}
			for k := 0; k < len(curr_gen[j].chrom_pair[0].ids); k++ {
				ancestors[curr_gen[j].chrom_pair[0].ids[k]] = struct{}{}
			}
			for k := 0; k < len(curr_gen[j].chrom_pair[1].ids); k++ {
				ancestors[curr_gen[j].chrom_pair[1].ids[k]] = struct{}{}
			}
			for k := range ancestors {
				num_descendants[k]++
			}

			// Get number of segments
			num_segments += len(curr_gen[j].chrom_pair[0].ids) + len(curr_gen[j].chrom_pair[1].ids)

			// Get lengths of ROH
			roh_lengths[j], roh_counts[j] = curr_gen[j].getROH(seq_len)
		}
		// Get average number of descendants
		mean_num_descendants[i] = mean_non_zero(num_descendants)

		// Get average segment count
		mean_segment_count[i] = float64(num_segments) / float64(pop_size)

		// Get average segment length
		segment_len[i] = 2 * seq_len * pop_size / num_segments

		// Get ROH frequency
		roh_freq[i] = float64(mean(roh_lengths)) / float64(seq_len)
		s := sum(roh_counts)
		if s > 0 {
			mean_roh_len[i] = sum(roh_lengths) / s
		} else {
			mean_roh_len[i] = 0
		}

		prev_gen = curr_gen
	}

	return mean_num_descendants, mean_segment_count, segment_len, roh_freq, mean_roh_len
}
