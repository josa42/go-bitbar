package bitbar

import (
	"fmt"
	"os"
	"strings"
)

type menuItem interface {
	print()
}

type Menu struct {
	items []menuItem
	level int
}

func (m *Menu) Add(item menuItem) {
	m.items = append(m.items, item)
}

func (m *Menu) NewItem(label string) *Item {

	item := NewItem(label)
	item.Menu.level = m.level + 1

	m.Add(item)

	return item
}

func (m *Menu) NewImage(path string) *Item {

	item := NewItem("")
	item.Menu.level = m.level + 1
	item.Image(path)

	m.Add(item)

	return item
}

func (m *Menu) NewDivider() *Divider {
	divider := &Divider{}
	m.Add(divider)

	return divider
}

func (m *Menu) NewUpdateItem(version, updateFlag string) *Item {

	// TODO right labels
	item := NewItem(fmt.Sprintf("Current Version: %s", version))
	doUpdate := item.NewAlternateItem("Check for Updates...")

	doUpdate.
		Command(os.Args[0], updateFlag).
		OpenTerminal()

	m.Add(item)

	return item
}

func (m Menu) print() {
	for _, i := range m.items {
		if m.level > 0 {
			fmt.Printf("%s", strings.Repeat("--", m.level))
		}
		i.print()
	}
}
