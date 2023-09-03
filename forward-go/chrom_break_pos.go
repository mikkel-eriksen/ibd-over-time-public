package main

func getChromBreakPos(chrom_lengths []int) []int {
	chrom_break_pos := make([]int, len(chrom_lengths)+1)
	for i := 0; i < len(chrom_lengths); i++ {
		chrom_break_pos[i+1] = chrom_break_pos[i] + chrom_lengths[i]
	}
	return chrom_break_pos
}
