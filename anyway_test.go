package anyway

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

// one unit of width per rune, so tests do not need a font.
func runeWidth(s string) float64 { return float64(utf8.RuneCountInString(s)) }

func TestWrapText(t *testing.T) {
	got := wrapText("one two three four", 10, runeWidth)
	want := []string{"one two", "three four"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestWrapTextLongWordKeepsOwnLine(t *testing.T) {
	got := wrapText("a supercalifragilistic b", 5, runeWidth)
	want := []string{"a", "supercalifragilistic", "b"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestTextboxNeverSplitsRunes(t *testing.T) {
	var b textbox
	b.append("“Knowledge…”", 100, runeWidth)
	for !b.done() {
		b.tick()
		for _, l := range b.visible() {
			if !utf8.ValidString(l) {
				t.Fatalf("invalid utf-8 at %d: %q", b.revealed, l)
			}
		}
	}
	if got := strings.Join(b.visible(), " "); got != "“Knowledge…”" {
		t.Fatalf("got %q", got)
	}
}

func TestTextboxAppendOnlyAnimatesNewText(t *testing.T) {
	var b textbox
	b.append("one", 100, runeWidth)
	b.skip()
	b.append("two", 100, runeWidth)
	if b.done() {
		t.Fatal("new text should not be revealed yet")
	}
	if got := b.visible(); !reflect.DeepEqual(got, []string{"one"}) {
		t.Fatalf("got %q", got)
	}
}

func TestParseValidation(t *testing.T) {
	cases := map[string]string{
		"unknown type": "commands:\n  - type: NOPE\n",
		"empty text":   "commands:\n  - type: TEXT\n",
		"empty image":  "commands:\n  - type: BG\n",
		"empty audio":  "commands:\n  - type: BGM\n",
	}
	for name, src := range cases {
		if _, err := Parse([]byte(src)); err == nil {
			t.Errorf("%s: expected error", name)
		}
	}
}

func newTestPlayer(t *testing.T, src string) *player {
	t.Helper()
	s, err := Parse([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	p := &player{scene: s, width: 100, measure: runeWidth}
	p.advance()
	return p
}

func TestPlayerFlow(t *testing.T) {
	p := newTestPlayer(t, `
commands:
  - type: TEXT
    text: one
  - type: BGM
    audio: a.mp3
  - type: TEXT
    text: two
  - type: CLEAR
  - type: TEXT
    text: three
`)
	if p.press() { // first press finishes the typewriter
		t.Fatal("ended too early")
	}
	if !p.box.done() {
		t.Fatal("press should skip to end of text")
	}
	if p.press() { // advances to "two", skipping BGM
		t.Fatal("ended too early")
	}
	if p.box.done() {
		t.Fatal("two should be animating")
	}
	p.press()
	p.press() // CLEAR then "three"
	if got := p.box.visible(); len(got) != 0 {
		t.Fatalf("clear should drop old text, got %q", got)
	}
	p.press()
	if got := p.box.visible(); !reflect.DeepEqual(got, []string{"three"}) {
		t.Fatalf("got %q", got)
	}
	if !p.press() {
		t.Fatal("scene should be finished after last text")
	}
}
