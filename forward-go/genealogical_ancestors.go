package main

func getGenealogicalAncestors(pop [][]Individual) ([]float64, int, int) {
	num_gen := len(pop)
	pop_size := len(pop[0])

	num_ancestors := new_matrix(num_gen, pop_size)
	num_desc_in_gen0 := new_matrix(num_gen, pop_size)

	// Run BFS from each individual in generation 0
	for j := 0; j < pop_size; j++ {
		q := map[int]struct{}{}
		q[j] = struct{}{}

		for i := 0; i < num_gen-1; i++ {
			num_ancestors[i][j] = len(q)
			new_q := map[int]struct{}{}
			for k := range q {
				new_q[pop[i][k].mom_id] = struct{}{}
				new_q[pop[i][k].dad_id] = struct{}{}
				num_desc_in_gen0[i][k] += 1
			}
			q = new_q
		}
		for k := range q {
			num_desc_in_gen0[num_gen-1][k] += 1
		}
		num_ancestors[num_gen-1][j] = len(q)
	}

	// Calculate mean number of ancestors and find TMRCA and IAP
	mean_num_ancestors := make([]float64, num_gen)
	TMRCA := -1
	IAP := -1
	for i := 0; i < num_gen; i++ {
		mean_num_ancestors[i] = mean(num_ancestors[i])
		if TMRCA == -1 && max(num_desc_in_gen0[i]) == pop_size {
			TMRCA = i
		}
		if IAP == -1 && int(mean_non_zero(num_desc_in_gen0[i])) == pop_size {
			IAP = i
		}
	}

	return mean_num_ancestors, TMRCA, IAP
}
