package anyway

import "strings"

// wrapText greedily breaks s into lines no wider than maxWidth, as reported
// by measure. A single word wider than maxWidth gets its own line.
func wrapText(s string, maxWidth float64, measure func(string) float64) []string {
	var lines []string
	cur := ""
	for _, word := range strings.Fields(s) {
		if cur == "" {
			cur = word
			continue
		}
		if candidate := cur + " " + word; measure(candidate) <= maxWidth {
			cur = candidate
			continue
		}
		lines = append(lines, cur)
		cur = word
	}
	if cur != "" {
		lines = append(lines, cur)
	}
	return lines
}

// textbox holds wrapped text and reveals it one rune per tick.
type textbox struct {
	lines    [][]rune
	total    int // runes across all lines
	revealed int
}

// append wraps s and adds it below the current text. Only the new text is
// animated; what was already visible stays visible.
func (t *textbox) append(s string, maxWidth float64, measure func(string) float64) {
	for _, l := range wrapText(s, maxWidth, measure) {
		r := []rune(l)
		t.lines = append(t.lines, r)
		t.total += len(r)
	}
}

func (t *textbox) tick() {
	if t.revealed < t.total {
		t.revealed++
	}
}

func (t *textbox) skip()      { t.revealed = t.total }
func (t *textbox) done() bool { return t.revealed >= t.total }

func (t *textbox) clear() {
	t.lines = nil
	t.total = 0
	t.revealed = 0
}

// visible returns the lines as currently revealed, omitting lines not yet reached.
func (t *textbox) visible() []string {
	var out []string
	left := t.revealed
	for _, l := range t.lines {
		if left <= 0 {
			break
		}
		n := min(left, len(l))
		out = append(out, string(l[:n]))
		left -= n
	}
	return out
}
