package main

import (
	"github.com/arseto/flexnum/segments"
	"testing"
	"time"
)

func TestFormatter(t *testing.T) {
	f := NewFormatter()

	strseg := segments.NewStringSegment("EXP", "/")
	f.AddSegment(strseg)

	dateseg := segments.NewDateSegment(
		time.Now(),
		"0106",
		"/",
	)
	f.AddSegment(dateseg)

	seqseg := segments.NewSequenceSegment(12, 6, "")
	f.AddSegment(seqseg)

	expected := "EXP/" + time.Now().Format("0106") + "/000012"
	t.Logf("Expected: %s", expected)

	actual := f.Format()
	t.Logf("Actual: %s", actual)

	if actual != expected {
		t.Errorf("Expected: %s, actual: %s", expected, actual)
	}
}
