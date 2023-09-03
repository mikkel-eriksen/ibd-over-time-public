package main

func getGenealogicalDescendants(pop [][]Individual) []float64 {
	num_gen := len(pop)
	pop_size := len(pop[0])

	num_descendants := new_matrix(num_gen, pop_size)

	for j := 0; j < pop_size; j++ {
		q := map[int]struct{}{}
		q[j] = struct{}{}

		for i := 0; i < num_gen; i++ {
			num_descendants[i][j] = len(q)
			new_q := map[int]struct{}{}
			for k := range q {
				for l := 0; l < len(pop[i][k].children); l++ {
					new_q[pop[i][k].children[l]] = struct{}{}
				}
			}
			q = new_q
		}
	}

	mean_num_descendants := make([]float64, num_gen)
	for i := 0; i < num_gen; i++ {
		mean_num_descendants[i] = mean_non_zero(num_descendants[i])
	}

	return mean_num_descendants
}
