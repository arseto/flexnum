package segments

import "time"
import "testing"

func TestGetDateSegment(t *testing.T) {
	seg := NewDateSegment(
		time.Date(2018,
			3,
			30,
			23,
			0,
			0,
			0,
			time.UTC,
		),
		"0106",
		"/",
	)

	actual := seg.GetSegmentStr()
	expected := "0318/"

	if actual != expected {
		t.Errorf("Expected: %s, actual: %s", expected, actual)
	}
}
