package generic

import (
	"errors"
	"strings"
)

// Render renders the template given.
func (rndr *Renderer) renderTemplate(tpl *string, buf stringWriter) error {
	var openSeq = rndr.TagOpen.String()
	var closeSeq = rndr.TagClose.String()
	var doubleClose = strings.Repeat(closeSeq, 2)

	pos := 0
	for pos < len(*tpl) {
		// Search for tag open marker
		idx := strings.Index((*tpl)[pos:], openSeq)

		// Write the chunk from the position till the open marker or the end of line
		var endPos int
		if idx > 0 {
			endPos = pos + idx
		} else if idx == -1 {
			endPos = len(*tpl)
		}

		if endPos > 0 {
			// Convert all doubled tag close sequences
			for {
				doubleCloseIdx := strings.Index((*tpl)[pos:endPos], doubleClose)
				if doubleCloseIdx >= 0 {
					n, _ := buf.WriteString((*tpl)[pos : pos+doubleCloseIdx])
					pos += n

					n, _ = buf.WriteString(closeSeq)
					pos += n * 2
				} else {
					break
				}
			}

			if pos < endPos {
				n, _ := buf.WriteString((*tpl)[pos:endPos])
				pos += n
			}
		}

		// Convert tag if any into ANSI sequence
		if idx >= 0 {
			openIdx := pos

			// Check if open tag sequence is doubled
			secondOpenIdx := strings.Index((*tpl)[openIdx+1:], openSeq)
			if secondOpenIdx == 0 {
				buf.WriteString(openSeq)
				pos += len(openSeq) * 2
			} else {
				// Find the closing tag position
				closingIdx := strings.Index((*tpl)[openIdx:], closeSeq)
				if closingIdx >= 0 {
					pos += closingIdx + 1
				} else {
					return errors.New("attribute tag misses closing sequence")
				}

				// Get tag content and split to attribute list
				content := (*tpl)[openIdx+1 : openIdx+closingIdx]
				attributes := strings.Split(content, rndr.AttributeDelimiter.String())

				// Build ANSI sequence from attributes
				sequnce, err := rndr.ansiSequence(attributes)
				if err != nil {
					return err
				}

				buf.WriteANSISequence(sequnce)
			}
		}
	}

	if rndr.AutoReset && buf.Len() > 0 {
		buf.WriteANSISequence(ansiResetSequence)
	}

	return nil
}
