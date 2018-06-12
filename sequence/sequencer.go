package sequence

import "errors"

type Sequencer interface {
	Next(sequenceID string) (next *Sequence, err error)
	Increment(sequenceID string) error
}

type SequenceRepo interface {
	Find(sequenceID string) (*Sequence, error)
	Exist(sequenceID string) (bool, error)
	Create(seq *Sequence) error
	Update(seq *Sequence) error
	Delete(sequenceID string) error
}

type sequencer struct {
	repo SequenceRepo
}

func (sqr *sequencer) CreateNew(seq *Sequence) (err error) {
	if seq.SequenceID == "" {
		return errors.New("sequenceID is required")
	}
	err = sqr.repo.Create(seq)
	return
}

func (sqr *sequencer) Next(sequenceID string) (next *Sequence, err error) {
	next, err = sqr.repo.Find(sequenceID)
	if err != nil {
		return
	}

	next.Increment()
	return
}

func (sqr *sequencer) Increment(sequenceID string) (err error) {
	seq, err := sqr.Next(sequenceID)
	if err != nil {
		return
	}
	err = sqr.repo.Update(seq)

	return
}
