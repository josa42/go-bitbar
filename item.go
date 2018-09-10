package bitbar

import (
	"fmt"
	"strconv"
	"strings"
)

// Divider :
type Divider struct{}

func (d Divider) print() {
	fmt.Println("---")
}

// Item :
type Item struct {
	label         string
	link          string
	color         string
	font          string
	size          int
	command       *Command
	dropdown      bool
	maxLength     int
	trim          bool
	alternateItem *Item
	alternate     bool
	emojize       bool
	Menu          Menu
	image         *Image
	templateImage *Image
}

// NewItem :
func NewItem(label string) *Item {
	return &Item{
		label:    label,
		trim:     true,
		dropdown: true,
	}
}

// AddAlternateItem :
func (i *Item) NewAlternateItem(label string) *Item {
	i.alternateItem = NewItem(label)
	i.alternateItem.alternate = true
	i.alternateItem.Menu.level = i.Menu.level
	return i.alternateItem
}

// Link :
// > href=.. to make the item clickable
func (i *Item) Link(link string) *Item {
	// validation?
	i.link = link
	return i
}

// Color :
// > color=.. to change their text color. eg. color=red or color=#ff0000
func (i *Item) Color(color string) *Item {
	// validation?
	i.color = color
	return i
}

// Font :
// > font=.. to change their text font. eg. font=UbuntuMono-Bold
func (i *Item) Font(font string) *Item {
	// validation?
	i.font = font
	return i
}

// Size :
// > size=.. to change their text size. eg. size=12
func (i *Item) Size(size int) *Item {
	// validation?
	i.size = size
	return i
}

// MaxLength :
// > length=.. to truncate the line to the specified number of characters. A … will be added to any truncated strings, as well as a tooltip displaying the full string. eg. length=10
func (i *Item) MaxLength(maxLength int) *Item {
	// validation?
	i.maxLength = maxLength
	return i
}

// Command :
func (i *Item) Command(name string, arg ...string) *Item {
	// validation?
	i.command = NewCommand(name, arg)
	return i
}

func (i *Item) OpenTerminal() *Item {
	i.command.terminal = true
	return i
}

func (i *Item) Refresh() *Item {
	i.command.refresh = true
	return i
}

func (i *Item) Image(path string) *Item {
	// validation?
	i.image = newImage(path)
	return i
}

func (i *Item) TemplateImage(path string) *Item {
	// validation?Item
	i.templateImage = newImage(path)
	return i
}

func (i *Item) Resize(width, height int) *Item {
	if i.image != nil {
		i.image.width = width
		i.image.height = height
	}

	if i.templateImage != nil {
		i.image.width = width
		i.image.height = height
	}
	return i
}

// Emojize :
// > emojize=false will disable parsing of github style :mushroom: into 🍄
func (i *Item) NoEmojis() *Item {
	i.emojize = false
	return i
}

func (i Item) print() {

	args := []string{}

	if i.link != "" {
		args = append(args, fmt.Sprintf("href=%s", i.link))
	}

	if i.color != "" {
		args = append(args, fmt.Sprintf("color=%s", i.color))
	}

	if i.font != "" {
		args = append(args, fmt.Sprintf("font=%s", i.font))
	}

	if i.size > 0 {
		args = append(args, fmt.Sprintf("size=%d", i.size))
	}

	if i.maxLength > 0 {
		args = append(args, fmt.Sprintf("length=%d", i.maxLength))
	}

	args = append(args, fmt.Sprintf("trim=%s", strconv.FormatBool(i.trim)))

	if i.emojize {
		args = append(args, "emojize=true")
	}

	if i.alternate {
		args = append(args, "alternate=true")
	}

	if !i.dropdown {
		args = append(args, "dropdown=false")
	}

	if i.command != nil {
		args = append(args, i.command.string())
	}

	if i.image != nil {
		args = append(args, i.image.string("image"))
	}

	if i.templateImage != nil {
		args = append(args, i.templateImage.string("templateImage"))
	}

	fmt.Printf("%s | %s\n", i.label, strings.Join(args, " "))

	i.Menu.print()

	if i.alternateItem != nil {
		i.alternateItem.print()
	}
}
