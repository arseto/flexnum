package segments

type StringSegment struct {
	value     string
	separator string
}

func NewStringSegment(value, separator string) *StringSegment {
	return &StringSegment{
		value,
		separator,
	}
}

func (seg *StringSegment) GetSegmentStr() string {
	return seg.value + seg.separator
}
