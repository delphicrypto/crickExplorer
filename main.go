/*
A presentation of the tview package, implemented with tview.

Navigation

The presentation will advance to the next slide when the primitive demonstrated
in the current slide is left (usually by hitting Enter or Escape). Additionally,
the following shortcuts can be used:

  - Ctrl-N: Jump to next slide
  - Ctrl-P: Jump to previous slide
*/
package main

import (
	"fmt"
	"strconv"
	//"time"
	"os"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	cc "github.com/delphicrypto/blockchain_go"
	"github.com/fsnotify/fsnotify"
)

// Slide is a function which returns the slide's main primitive and its title.
// It receives a "nextSlide" function which can be called to advance the
// presentation to the next slide.
type Slide func(nextSlide func()) (title string, content tview.Primitive)

// The application.
var app = tview.NewApplication()
var chain *cc.Blockchain
var dbFile string
var watcher *fsnotify.Watcher

func main() {
	dbFile =  os.Args[1]

	// The presentation slides.
	slides := []Slide{
		Cover,
		Table,
		Grid,
	}

	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false)

	
	
	// Create the pages for all slides.
	currentSlide := 0
	info.Highlight(strconv.Itoa(currentSlide))
	pages := tview.NewPages()
	previousSlide := func() {
		currentSlide = (currentSlide - 1 + len(slides)) % len(slides)
		info.Highlight(strconv.Itoa(currentSlide)).
			ScrollToHighlight()
		pages.SwitchToPage(strconv.Itoa(currentSlide))
	}
	nextSlide := func() {
		currentSlide = (currentSlide + 1) % len(slides)
		info.Highlight(strconv.Itoa(currentSlide)).
			ScrollToHighlight()
		pages.SwitchToPage(strconv.Itoa(currentSlide))
	}
	for index, slide := range slides {
		title, primitive := slide(nextSlide)
		pages.AddPage(strconv.Itoa(index), primitive, true, index == currentSlide)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""]  `, index+1, index, title)
	}

	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
		}
		return event
	})

	watcher, _ = fsnotify.NewWatcher()
    
    defer watcher.Close()

    // out of the box fsnotify can watch a single file, or a single directory
    if err := watcher.Add(dbFile); err != nil {
        fmt.Println("ERROR", err)
    }
    // go func() {
    //     for {
    //         select {
    //         // watch for events
    //         case event := <-watcher.Events:
    //             updateMain(main)
    //             updateTable(table, blockDisplay)
    //             app.draw()

    //             // watch for errors
    //         case err := <-watcher.Errors:
    //             fmt.Println("ERROR", err)
    //         }
    //     }
    // }()

	// Start the application.
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
