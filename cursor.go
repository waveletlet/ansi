package ansi

import (
	"strconv"
)

// Control Sequence Introducer
// ============================================================================
// | ESC [ 	0x9B 	CSI | Control Sequence Introducer | Most of the useful      |
// |                  |                             | sequences, see next     |
// |                  |                             | section.                |
// ============================================================================
// For CSI, or "Control Sequence Introducer" commands, the ESC [ is followed
// by any number (including none) of "parameter bytes" in the range 0x30–0x3F
// (ASCII 0–9:;<=>?), then by any number of "intermediate bytes" in the range
// 0x20–0x2F (ASCII space and !"#$%&'()*+,-./), then finally by a single
// "final byte" in the range 0x40–0x7E (ASCII @A–Z[\]^_`a–z{).

func ShowCursor() string { return prefix + "?25h" }
func HideCursor() string { return prefix + "?25l" }

// Absolute Movement
func MoveCursorTo(row, column int) string {
	return prefix + strconv.Itoa(row) + ";" + strconv.Itoa(column) + "H"
}

// Relative Movement
func MoveCursorUp(rows int) string         { return prefix + strconv.Itoa(rows) + "A" }
func MoveCursorDown(rows int) string       { return prefix + strconv.Itoa(rows) + "B" }
func MoveCursorRight(columns int) string   { return prefix + strconv.Itoa(columns) + "C" }
func MoveCursorLeft(columns int) string    { return prefix + strconv.Itoa(columns) + "D" }
func MoveCursorUpperLeft(count int) string { return prefix + strconv.Itoa(count) + "H" }
func MoveCursorToNextLine() string         { return escape + "E" }

func ClearLineRight() string    { return prefix + "0K" }
func ClearLineLeft() string     { return prefix + "1K" }
func ClearEntireLine() string   { return prefix + "2K" }
func ClearScreenDown() string   { return prefix + "0J" }
func ClearScreenUp() string     { return prefix + "1J" }
func ClearEntireScreen() string { return prefix + "2J" }

func SaveAttributes() string    { return escape + "7" }
func RestoreAttributes() string { return escape + "8" }
