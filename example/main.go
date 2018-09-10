package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	bitbar "github.com/josa42/go-bitbar"
)

func main() {
	installFlag := flag.Bool("install", false, "Install")
	flag.Parse()

	if *installFlag {
		runInstall()
		return
	}

	runMain()
}

func runMain() {
	r := bitbar.New("Example")

	r.Menu.NewItem("Item")
	r.Menu.NewItem("Item with a Link").Link("http://example.org")

	r.Menu.NewDivider()

	item := r.Menu.NewItem("Item with a Submenu")
	item.Menu.NewItem("Item A")

	sub := item.Menu.NewItem("Item B")
	sub.Menu.NewItem("Item B.1")
	sub.Menu.NewItem("Item B.2")

	r.Menu.NewDivider()

	r.Menu.
		NewItem("Background Script").
		Command("/usr/bin/say", "Hello")

	r.Menu.
		NewItem("Terminal Script").
		Command("banner", "-w", "50", "Hello").
		OpenTerminal()

	r.Menu.NewDivider()

	r.Menu.NewItem(":smile: :cookie: :tada:")

	r.Menu.NewDivider()

	r.Menu.
		NewItem(":smile: (press <alt>)").
		NewAlternateItem(":disappointed:")

	r.Menu.NewDivider()

	r.Menu.NewItem("Color: Green").Color("green")
	r.Menu.NewItem("Color: Red").Color("red")
	r.Menu.NewItem("Color: Blue").Color("blue")
	r.Menu.NewItem("Font: Menlo").Font("Menlo").Link("https://en.wikipedia.org/wiki/Menlo_(typeface)")
	r.Menu.
		NewItem("Lorem ipsum dolor sit amet").MaxLength(10).
		NewAlternateItem("Lorem ipsum dolor sit amet")

	r.Menu.NewDivider()

	if bitbar.InDarkMode() {
		r.Menu.NewItem("Mode: Dark")
	} else {
		r.Menu.NewItem("Mode: Light")
	}
	r.Menu.NewDivider()

	r.Menu.NewImage("https://github.com/matryer/bitbar/raw/master/Docs/bitbar-32.png")

	r.Menu.NewDivider()

	r.Menu.
		NewItem("Refresh").
		Command("true").
		Refresh()

	r.Print()

}

func runInstall() {
	file, _ := filepath.Abs(os.Args[0])

	exec.
		Command("cp", file, "/tmp/example.2s.cgo").
		Run()

	exec.
		Command("open", fmt.Sprintf("bitbar://openPlugin?src=file://%s", "/tmp/example.2s.cgo")).
		Run()
}
