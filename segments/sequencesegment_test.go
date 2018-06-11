package segments

import "testing"

func TestGetSequenceSegment(t *testing.T) {
	seg := NewSequenceSegment(12, 6, "/")

	expected := "000012/"
	t.Logf("Expected: %s", expected)

	actual := seg.GetSegmentStr()
	t.Logf("Actual: %s", actual)

	if actual != expected {
		t.Errorf("Expected: %s, actual: %s", expected, actual)
	}

	seg = NewSequenceSegment(13, 5, "")
	actual = seg.GetSegmentStr()

	expected = "00013"

	if actual != expected {
		t.Errorf("Expected: %s, actual: %s", expected, actual)
	}

}
