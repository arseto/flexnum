package segments

import "testing"

func TestStringSegment(t *testing.T) {
	seg := NewStringSegment("EXP", "/")
	actual := seg.GetSegmentStr()
	expected := "EXP/"

	if actual != expected {
		t.Errorf("Expected: %s, actual: %s", expected, actual)
	}
}
