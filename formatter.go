package main

type Segment interface {
	GetSegmentStr() string
}

type Formatter interface {
	AddSegment(Segment)
	Format() string
}

func NewFormatter() Formatter {
	return &formatter{}
}

type formatter struct {
	segments []Segment
}

func (f *formatter) AddSegment(seg Segment) {
	f.segments = append(f.segments, seg)
}

func (f *formatter) Format() string {
	var val string
	for _, seg := range f.segments {
		val += seg.GetSegmentStr()
	}
	return val
}
