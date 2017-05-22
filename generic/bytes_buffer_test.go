package generic

import "testing"

func TestBytesBuffer(t *testing.T) {
	text := "This is text"
	seq := "This is sequence"
	allLen := len(text) + len(seq)

	buf := new(bytesBuffer)
	buf.WriteString(text)
	buf.WriteANSISequence(seq)

	if buf.Len() != allLen {
		t.Fatalf("bytes buffer length %d differs from text and sequence length %d", buf.Len(), allLen)
	}
}
