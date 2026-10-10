package anyway

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

// Command types understood by the script format.
const (
	CmdText  = "TEXT"
	CmdClear = "CLEAR"
	CmdBG    = "BG"
	CmdBGM   = "BGM"
	CmdSFX   = "SFX"
)

// Scene is one script: an ordered list of commands.
type Scene struct {
	ID       int       `yaml:"scene_id"`
	Name     string    `yaml:"scene_name"`
	Commands []Command `yaml:"commands"`
}

// Command is a single step of a scene. Which fields matter depends on Type.
type Command struct {
	Type  string `yaml:"type"`
	Text  string `yaml:"text"`
	Image string `yaml:"image"`
	Audio string `yaml:"audio"`
}

// Load reads and validates a scene from a YAML file.
func Load(path string) (*Scene, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(data)
}

// Parse decodes and validates a scene from YAML bytes.
func Parse(data []byte) (*Scene, error) {
	var s Scene
	if err := yaml.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("anyway: parse scene: %w", err)
	}
	if err := s.validate(); err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *Scene) validate() error {
	for i, c := range s.Commands {
		var missing string
		switch c.Type {
		case CmdText:
			if c.Text == "" {
				missing = "text"
			}
		case CmdBG:
			if c.Image == "" {
				missing = "image"
			}
		case CmdBGM, CmdSFX:
			if c.Audio == "" {
				missing = "audio"
			}
		case CmdClear:
		default:
			return fmt.Errorf("anyway: command %d: unknown type %q", i, c.Type)
		}
		if missing != "" {
			return fmt.Errorf("anyway: command %d (%s): missing %q", i, c.Type, missing)
		}
	}
	return nil
}
