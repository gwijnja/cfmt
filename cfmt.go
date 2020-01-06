package cfmt

import (
	"fmt"
)

const (
	Black     = 30
	Red       = 31
	Green     = 32
	Yellow    = 33
	Blue      = 34
	Magenta   = 35
	Cyan      = 36
	LightGray = 37

	DarkGray     = 90
	LightRed     = 91
	LightGreen   = 92
	LightYellow  = 93
	LightBlue    = 94
	LightMagenta = 95
	LightCyan    = 96
	White        = 97
)

func Sprintf(fg int, bg int, format string, a ...interface{}) string {
	return fmt.Sprintf(esc(fg)+esc(10+bg)+format+esc(0), a...)
}

func Printf(fg int, bg int, format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(esc(fg)+esc(10+bg)+format+esc(0), a...)
}

func esc(v int) string {
	return fmt.Sprintf("\x1b[%dm", v)
}
