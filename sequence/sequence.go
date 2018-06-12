package sequence

type Sequence struct {
	SequenceID string
	Current    int
	Start      int
	End        int
}

func (seq *Sequence) Next() int {
	if seq.Current == 0 {
		seq.Current = seq.Start - 1
	}

	next := seq.Current + 1
	if next > seq.End {
		return seq.Start
	}
	return next
}

func (seq *Sequence) Increment() {
	seq.Current = seq.Next()
}
