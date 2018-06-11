package sequence

import "github.com/arseto/flexnum/segments"

type segmentFactory struct {
	repo SequenceRepo
}

func (sf *segmentFactory) MakeSegment(sequenceID string, digits int, separator string) (
	seg *segments.SequenceSegment, err error) {
	seq, err := sf.repo.Find(sequenceID)
	if err != nil {
		return
	}
	seg = segments.NewSequenceSegment(seq.current, digits, separator)
	return
}
