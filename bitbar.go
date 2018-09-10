package bitbar

import (
	"os/exec"
	"strings"
)

// BitBar :
type BitBar struct {
	Root *Item
	Menu Menu
}

// New :
func New(title string) *BitBar {
	r := &BitBar{}

	r.Root = r.Menu.NewItem(title)
	r.Menu.NewDivider()

	return r
}

func NewImage(path string) *BitBar {
	r := &BitBar{}

	r.Root = r.Menu.NewImage(path)
	r.Menu.NewDivider()

	return r
}

// Print :
func (r *BitBar) Print() {
	r.Menu.print()
}

// InDarkMode :
func InDarkMode() bool {
	out, _ := exec.Command("/usr/bin/defaults", "read", "-g", "AppleInterfaceStyle").Output()
	return strings.TrimSpace(string(out)) == "Dark"
}
