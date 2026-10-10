package anyway

// player walks a scene's commands. It has no dependency on the renderer so
// the flow can be tested without a window.
type player struct {
	scene *Scene
	next  int
	box   textbox

	width   float64
	measure func(string) float64
}

// advance runs commands until the next TEXT is shown or the scene ends.
// It reports false when there is nothing left to run.
//
// BG, BGM and SFX are validated at load time but not executed yet.
func (p *player) advance() bool {
	for p.next < len(p.scene.Commands) {
		c := p.scene.Commands[p.next]
		p.next++
		switch c.Type {
		case CmdText:
			p.box.append(c.Text, p.width, p.measure)
			return true
		case CmdClear:
			p.box.clear()
		}
	}
	return false
}

// press is the player's "continue" input. It returns true when the scene is
// over and the game should end.
func (p *player) press() (finished bool) {
	if !p.box.done() {
		p.box.skip()
		return false
	}
	return !p.advance()
}
