package main

type Sequence struct {
	ids     []int
	lengths []int
}

func (seq *Sequence) addSegment(id, length int) {
	if length > 0 {
		seq.ids = append(seq.ids, id)
		seq.lengths = append(seq.lengths, length)
	}
}

func (seq1 *Sequence) extend(seq2 Sequence) {
	seq1.ids = append(seq1.ids, seq2.ids...)
	seq1.lengths = append(seq1.lengths, seq2.lengths...)
}

func (in_seq Sequence) getSequence(start_pos, end_pos int) Sequence {
	i := 0
	pointer := 0
	var out_seq Sequence
	for {
		pointer += in_seq.lengths[i]
		if start_pos >= pointer {
			i += 1
		} else if end_pos > pointer {
			out_seq.addSegment(in_seq.ids[i], pointer-start_pos)
			i += 1
			start_pos = pointer
		} else {
			out_seq.addSegment(in_seq.ids[i], end_pos-start_pos)
			return out_seq
		}
	}
}
