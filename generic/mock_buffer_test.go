package generic

import "testing"

func TestMockBuffer(t *testing.T) {
	text := "This is text"
	seq := "This is sequence"

	buf := new(mockBuffer)
	buf.WriteString(text)
	buf.WriteANSISequence(seq)

	if buf.Len() != len(text) {
		t.Fatalf("mock buffer length %d differs from text length %d", buf.Len(), len(text))
	}
}
