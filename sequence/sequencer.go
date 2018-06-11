package sequence

type sequencer struct {
	repo SequenceRepo
}

func (sqr *sequencer) GetPreview(sequenceID string) (next int, err error) {
	seq, err := sqr.repo.Find(sequenceID)
	if err != nil {
		return
	}

	next = seq.Next()
	return
}

func (sqr *sequencer) Increment(sequenceID string) (err error) {
	seq, err := sqr.repo.Find(sequenceID)
	if err != nil {
		return
	}
	seq.Increment()

	sqr.repo.Update(seq)

	return
}
