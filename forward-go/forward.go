package main

import (
	"flag"
	"math/rand"
	"time"
)

func main() {
	// Command-line parser
	pop_size := flag.Int("p", 100, "Population size")
	num_gen := flag.Int("g", 50, "Number of generations")
	recomb_rate := flag.Float64("r", 1e-8, "Recombination rate")
	chrom_len_file := flag.String("f", "", "Filepath of chromosome lengths file")
	seed := flag.Int64("s", 0, "seed")
	flag.Parse()

	// Set seed
	if *seed != 0 {
		rand.Seed(*seed)
	}

	// Read file with chromosome length
	chrom_lengths := []int{3000000000}
	if len(*chrom_len_file) > 0 {
		chrom_lengths = read_file(*chrom_len_file)
	}

	// Initialize data map
	data := map[string]interface{}{
		"pop_size":      *pop_size,
		"num_gen":       *num_gen,
		"recomb_rate":   *recomb_rate,
		"seed":          *seed,
		"chrom_lengths": chrom_lengths,
		"seq_len":       sum(chrom_lengths),
	}

	// Simulate backward-in-time population
	backward_pop := getBackwardPopulation(*pop_size, *num_gen, chrom_lengths, *recomb_rate)

	// Get number of genealogical ancestors, TMRCA, and IAP
	timestamp := time.Now()
	genealogical_ancestors, TMRCA, IAP := getGenealogicalAncestors(backward_pop)
	data["genealogical_ancestors"] = genealogical_ancestors
	data["TMRCA"] = TMRCA
	data["IAP"] = IAP
	data["runtime_genealogical_ancestors"] = time.Since(timestamp).Seconds()

	// Simulate recombinations to get number of genetic ancestors
	timestamp = time.Now()
	genetic_ancestors := make([]float64, *num_gen)
	for i := 0; i < *num_gen; i++ {
		genetic_ancestors[i] = getGeneticAncestors(i, backward_pop, chrom_lengths)
	}
	data["genetic_ancestors"] = genetic_ancestors
	data["runtime_genetic_ancestors"] = time.Since(timestamp).Seconds()

	// Simulate forward-in-time population
	forward_pop := getForwardPopulation(*pop_size, *num_gen, chrom_lengths, *recomb_rate)

	// Get number of genealogical descendants
	timestamp = time.Now()
	genealogical_descendants := getGenealogicalDescendants(forward_pop)
	data["genealogical_descendants"] = genealogical_descendants
	data["runtime_genealogical_descendants"] = time.Since(timestamp).Seconds()

	// Simulate recombinations and get number of genetic descendants, length of segments, ROH frequencies and lengths
	timestamp = time.Now()
	genetic_descendants, segment_count, segment_len, roh_freq, roh_len := getGeneticDescendants(forward_pop, chrom_lengths)
	data["genetic_descendants"] = genetic_descendants
	data["segment_count"] = segment_count
	data["segment_len"] = segment_len
	data["roh_freq"] = roh_freq
	data["roh_len"] = roh_len
	data["runtime_genetic_descendants"] = time.Since(timestamp).Seconds()

	// Print data in JSON format to stdout
	printJSON(data)
}
