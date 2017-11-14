package logger

import (
	"fmt"
)

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	Black Attribute = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// return a formatted color based on the chosen Attribute for specific text
func GetColorFormat(color, style Attribute, text string) string {
	return fmt.Sprintf("%+v[%+v;%+vm%s%s[%+vm", escape, color, style, text, escape, Reset)
}
