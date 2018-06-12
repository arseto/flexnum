package sequence

import "github.com/arseto/flexnum/segments"

type SegmentFactory interface {
	NextSequenceSegment(sequenceID string, digits int, separator string) (
		*segments.SequenceSegment, error)
}

type segmentFactory struct {
	sqr Sequencer
}

func (sf *segmentFactory) NextSequenceSegment(sequenceID string, digits int, separator string) (
	seg *segments.SequenceSegment, err error) {
	seq, err := sf.sqr.Next(sequenceID)
	if err != nil {
		return
	}
	seg = segments.NewSequenceSegment(seq.Current, digits, separator)
	return
}
