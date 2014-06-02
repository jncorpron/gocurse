package curses

// #define _Bool int
// #define NCURSES_OPAQUE 1
// #include <curses.h>
// #cgo LDFLAGS: -lncurses
import "C"

import (
	"fmt"
)

type Window C.WINDOW

// Cursor options.
type CursorState int

const (
	HideCursor CursorState = iota
	NormalCursor
	HighCursor
)

type WindowAttribute int

const (
	WA_ALTCHARSET WindowAttribute = C.WA_ALTCHARSET
	WA_BLINK                      = C.WA_BLINK
	WA_BOLD                       = C.WA_BOLD
	WA_DIM                        = C.WA_DIM
	WA_INVIS                      = C.WA_INVIS
	WA_LEFT                       = C.WA_LEFT
	WA_PROTECT                    = C.WA_PROTECT
	WA_REVERSE                    = C.WA_REVERSE
	WA_RIGHT                      = C.WA_RIGHT
	WA_STANDOUT                   = C.WA_STANDOUT
	WA_UNDERLINE                  = C.WA_UNDERLINE
)

const (
	A_ALTCHARSET = C.A_ALTCHARSET
	A_BLINK      = C.A_BLINK
	A_BOLD       = C.A_BOLD
	A_DIM        = C.A_DIM
	A_INVIS      = C.A_INVIS
	A_PROTECT    = C.A_PROTECT
	A_REVERSE    = C.A_REVERSE
	A_STANDOUT   = C.A_STANDOUT
	A_UNDERLINE  = C.A_UNDERLINE
	A_ATTRIBUTES = C.A_ATTRIBUTES
	A_CHARTEXT   = C.A_CHARTEXT
	A_COLOR      = C.A_COLOR

	COLOR_BLACK   = C.COLOR_BLACK
	COLOR_BLUE    = C.COLOR_BLUE
	COLOR_GREEN   = C.COLOR_GREEN
	COLOR_CYAN    = C.COLOR_CYAN
	COLOR_RED     = C.COLOR_RED
	COLOR_MAGENTA = C.COLOR_MAGENTA
	COLOR_YELLOW  = C.COLOR_YELLOW
	COLOR_WHITE   = C.COLOR_WHITE

	KEY_BREAK     = C.KEY_BREAK
	KEY_DOWN      = C.KEY_DOWN
	KEY_UP        = C.KEY_UP
	KEY_LEFT      = C.KEY_LEFT
	KEY_RIGHT     = C.KEY_RIGHT
	KEY_HOME      = C.KEY_HOME
	KEY_BACKSPACE = C.KEY_BACKSPACE
	KEY_F0        = C.KEY_F0
	KEY_DL        = C.KEY_DL
	KEY_IL        = C.KEY_IL
	KEY_DC        = C.KEY_DC
	KEY_IC        = C.KEY_IC
	KEY_EIC       = C.KEY_EIC
	KEY_CLEAR     = C.KEY_CLEAR
	KEY_EOS       = C.KEY_EOS
	KEY_EOL       = C.KEY_EOL
	KEY_SF        = C.KEY_SF
	KEY_SR        = C.KEY_SR
	KEY_NPAGE     = C.KEY_NPAGE
	KEY_PPAGE     = C.KEY_PPAGE
	KEY_STAB      = C.KEY_STAB
	KEY_CTAB      = C.KEY_CTAB
	KEY_CATAB     = C.KEY_CATAB
	KEY_ENTER     = C.KEY_ENTER
	KEY_SRESET    = C.KEY_SRESET
	KEY_RESET     = C.KEY_RESET
	KEY_PRINT     = C.KEY_PRINT
	KEY_LL        = C.KEY_LL
	KEY_A1        = C.KEY_A1
	KEY_A3        = C.KEY_A3
	KEY_B2        = C.KEY_B2
	KEY_C1        = C.KEY_C1
	KEY_C3        = C.KEY_C3
	KEY_BTAB      = C.KEY_BTAB
	KEY_BEG       = C.KEY_BEG
	KEY_CANCEL    = C.KEY_CANCEL
	KEY_CLOSE     = C.KEY_CLOSE
	KEY_COMMAND   = C.KEY_COMMAND
	KEY_COPY      = C.KEY_COPY
	KEY_CREATE    = C.KEY_CREATE
	KEY_END       = C.KEY_END
	KEY_EXIT      = C.KEY_EXIT
	KEY_FIND      = C.KEY_FIND
	KEY_HELP      = C.KEY_HELP
	KEY_MARK      = C.KEY_MARK
	KEY_MESSAGE   = C.KEY_MESSAGE
	KEY_MOVE      = C.KEY_MOVE
	KEY_NEXT      = C.KEY_NEXT
	KEY_OPEN      = C.KEY_OPEN
	KEY_OPTIONS   = C.KEY_OPTIONS
	KEY_PREVIOUS  = C.KEY_PREVIOUS
	KEY_REDO      = C.KEY_REDO
	KEY_REFERENCE = C.KEY_REFERENCE
	KEY_REFRESH   = C.KEY_REFRESH
	KEY_REPLACE   = C.KEY_REPLACE
	KEY_RESTART   = C.KEY_RESTART
	KEY_RESUME    = C.KEY_RESUME
	KEY_SAVE      = C.KEY_SAVE
	KEY_SBEG      = C.KEY_SBEG
	KEY_SCANCEL   = C.KEY_SCANCEL
	KEY_SCOMMAND  = C.KEY_SCOMMAND
	KEY_SCOPY     = C.KEY_SCOPY
	KEY_SCREATE   = C.KEY_SCREATE
	KEY_SDC       = C.KEY_SDC
	KEY_SDL       = C.KEY_SDL
	KEY_SELECT    = C.KEY_SELECT
	KEY_SEND      = C.KEY_SEND
	KEY_SEOL      = C.KEY_SEOL
	KEY_SEXIT     = C.KEY_SEXIT
	KEY_SFIND     = C.KEY_SFIND
	KEY_SHELP     = C.KEY_SHELP
	KEY_SHOME     = C.KEY_SHOME
	KEY_SIC       = C.KEY_SIC
	KEY_SLEFT     = C.KEY_SLEFT
	KEY_SMESSAGE  = C.KEY_SMESSAGE
	KEY_SMOVE     = C.KEY_SMOVE
	KEY_SNEXT     = C.KEY_SNEXT
	KEY_SOPTIONS  = C.KEY_SOPTIONS
	KEY_SPREVIOUS = C.KEY_SPREVIOUS
	KEY_SPRINT    = C.KEY_SPRINT
	KEY_SREDO     = C.KEY_SREDO
	KEY_SREPLACE  = C.KEY_SREPLACE
	KEY_SRIGHT    = C.KEY_SRIGHT
	KEY_SRSUME    = C.KEY_SRSUME
	KEY_SSAVE     = C.KEY_SSAVE
	KEY_SSUSPEND  = C.KEY_SSUSPEND
	KEY_SUNDO     = C.KEY_SUNDO
	KEY_SUSPEND   = C.KEY_SUSPEND
	KEY_UNDO      = C.KEY_UNDO
)

func Initialize() (*Window, error) {
	window := (*Window)(C.initscr())

	if window == nil {
		return nil, CursesError{"Initialize failed"}
	}

	return window, nil
}

func NewWindow(rows int, columns int, startY int, startX int) (*Window, error) {
	window := (*Window)(C.newwin(C.int(rows), C.int(columns), C.int(startY), C.int(startX)))

	if window == nil {
		return nil, CursesError{"Failed to create window"}
	}

	return window, nil
}

func (win *Window) Delete() error {
	if int(C.delwin((*C.WINDOW)(win))) == C.ERR {
		return CursesError{"delete failed"}
	}
	return nil
}

func (win *Window) SubWindow(rows int, columns int, startY int, startX int) (*Window, error) {
	window := (*Window)(C.subwin((*C.WINDOW)(win), C.int(rows), C.int(columns), C.int(startY), C.int(startX)))

	if window == nil {
		return nil, CursesError{"Failed to create window"}
	}

	return window, nil
}

func (win *Window) DerivedWindow(rows int, columns int, startY int, startX int) (*Window, error) {
	window := (*Window)(C.derwin((*C.WINDOW)(win), C.int(rows), C.int(columns), C.int(startY), C.int(startX)))

	if window == nil {
		return nil, CursesError{"Failed to create window"}
	}

	return window, nil
}

func HasColors() bool {
	return isOk(C.has_colors())
}

func EnableColor() error {
	if HasColors() {
		return CursesError{"terminal does not support color"}
	}
	C.start_color()

	return nil
}

func InitColorPair(pairID int, foreground int, background int) error {
	if C.init_pair(C.short(pairID), C.short(foreground), C.short(background)) == 0 {
		return CursesError{"Init_pair failed"}
	}
	return nil
}

func ColorPair(pairID int) int32 {
	return int32(C.COLOR_PAIR(C.int(pairID)))
}

func DisableEcho() error {
	if int(C.noecho()) == C.ERR {
		return CursesError{"NoEcho failed"}
	}
	return nil
}

func EnableEcho() error {
	if int(C.noecho()) == C.ERR {
		return CursesError{"Echo failed"}
	}
	return nil
}

func Update() error {
	if int(C.doupdate()) == C.ERR {
		return CursesError{"Doupdate failed"}
	}
	return nil
}

func SetCursorState(state CursorState) error {
	if C.curs_set(C.int(state)) == C.ERR {
		return CursesError{"Curs_set failed"}
	}
	return nil
}

func DisableCharBreak() error {
	if C.nocbreak() == C.ERR {
		return CursesError{"Nocbreak failed"}
	}
	return nil
}

func EnableCharBreak() error {
	if C.cbreak() == C.ERR {
		return CursesError{"Cbreak failed"}
	}
	return nil
}

func End() error {
	if C.endwin() == C.ERR {
		return CursesError{"End failed"}
	}
	return nil
}

func (win *Window) GetChar() int {
	return int(C.wgetch((*C.WINDOW)(win)))
}

func (win *Window) Timeout(timeout int) {
	C.wtimeout((*C.WINDOW)(win), C.int(timeout))
}

func (win *Window) AddGet(x, y int, c int32, flags int32) {
	C.mvwaddch((*C.WINDOW)(win), C.int(y), C.int(x), C.chtype(c)|C.chtype(flags))
}

// Since CGO currently can't handle varg C functions we'll mimic the
// ncurses addstr functions.
func (win *Window) AddString(x, y int, str string, flags int32, v ...interface{}) {
	var resolvedString string
	if v != nil {
		resolvedString = fmt.Sprintf(str, v)
	} else {
		resolvedString = str
	}

	win.Move(x, y)

	for i := 0; i < len(resolvedString); i++ {
		C.waddch((*C.WINDOW)(win), C.chtype(resolvedString[i])|C.chtype(flags))
	}
}

// Normally Y is the first parameter passed in curses.
func (win *Window) Move(x, y int) {
	C.wmove((*C.WINDOW)(win), C.int(y), C.int(x))
}

func (win *Window) Resize(rows, cols int) {
	C.wresize((*C.WINDOW)(win), C.int(rows), C.int(cols))
}

func (w *Window) EnableKeypad(enabled bool) error {
	var keypadEnabledCode int = 0
	if enabled {
		keypadEnabledCode = 1
	}
	if C.keypad((*C.WINDOW)(w), C.int(keypadEnabledCode)) == C.ERR {
		return CursesError{"Keypad failed"}
	}
	return nil
}

func (win *Window) Refresh() error {
	if C.wrefresh((*C.WINDOW)(win)) == C.ERR {
		return CursesError{"refresh failed"}
	}
	return nil
}

func (win *Window) RedrawLine(startLine, numLines int) {
	C.wredrawln((*C.WINDOW)(win), C.int(startLine), C.int(numLines))
}

func (win *Window) Redraw() {
	C.redrawwin((*C.WINDOW)(win))
}

func (win *Window) Clear() {
	C.wclear((*C.WINDOW)(win))
}

func (win *Window) Erase() {
	C.werase((*C.WINDOW)(win))
}

func (win *Window) ClearToBottom() {
	C.wclrtobot((*C.WINDOW)(win))
}

func (win *Window) ClearToEOL() {
	C.wclrtoeol((*C.WINDOW)(win))
}

func (win *Window) Box(verch, horch int) {
	C.box((*C.WINDOW)(win), C.chtype(verch), C.chtype(horch))
}

func (win *Window) Background(colour int) {
	C.wbkgd((*C.WINDOW)(win), C.chtype(colour))
}

func (win *Window) AttrOn(flags int32) {
	C.wattron((*C.WINDOW)(win), C.int(flags))
}

func (win *Window) AttrOff(flags int32) {
	C.wattroff((*C.WINDOW)(win), C.int(flags))
}

type CursesError struct {
	message string
}

func (ce CursesError) Error() string {
	return ce.message
}
