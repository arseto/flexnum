package segments

import "time"

type DateSegment struct {
	value     time.Time
	format    string
	separator string
}

func NewDateSegment(value time.Time, format, separator string) *DateSegment {
	return &DateSegment{
		value,
		format,
		separator,
	}
}

func (seg *DateSegment) GetSegmentStr() string {
	timeStr := seg.value.Format(seg.format)
	return timeStr + seg.separator
}
