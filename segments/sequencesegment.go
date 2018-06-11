package segments

import "fmt"

type SequenceSegment struct {
	value     int
	digits    int
	separator string
}

func NewSequenceSegment(value, digits int, separator string) *SequenceSegment {
	return &SequenceSegment{
		value,
		digits,
		separator,
	}
}

func (seg *SequenceSegment) GetSegmentStr() string {
	format := "%0" + fmt.Sprintf("%d", seg.digits) + "d"
	padded := fmt.Sprintf(format, seg.value)
	return padded + seg.separator
}
