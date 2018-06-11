package sequence

type Sequence struct {
	sequenceID string
	current    int
	start      int
	end        int
}

func (seq *Sequence) Next() int {
	if seq.current == 0 {
		seq.current = seq.start - 1
	}

	next := seq.current + 1
	if next > seq.end {
		return seq.start
	}
	return next
}

func (seq *Sequence) Increment() {
	seq.current = seq.Next()
}

type SequenceRepo interface {
	Find(sequenceID string) (*Sequence, error)
	Create(seq *Sequence) error
	Update(seq *Sequence) error
	Delete(sequenceID string) error
}
