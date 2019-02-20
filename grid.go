package main

import (
	"fmt"
	"math"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/guptarohit/asciigraph"
	cc "github.com/delphicrypto/blockchain_go"
)

func barText(tw *tview.TextView, h int) {
	for i := 0; i< h; i++ {
		if i ==0 {
			fmt.Fprint(tw, fmt.Sprintln("[blue]▅▅▅ "))
		} else {
			fmt.Fprint(tw, fmt.Sprintln("[blue]███[teal]▌"))
		}
	}
	fmt.Fprint(tw, fmt.Sprintln("[teal]▝▀▀▘"))
}

// Grid demonstrates the grid layout.
func Grid(nextSlide func()) (title string, content tview.Primitive) {
	modalShown := false
	pages := tview.NewPages()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text).
			SetDoneFunc(func(key tcell.Key) {
				if modalShown {
					nextSlide()
					modalShown = false
				} else {
					pages.ShowPage("modal")
					modalShown = true
				}
			})
	}

	menu := newPrimitive("Menu")
	main := tview.NewTextView().
		SetScrollable(false).
		SetDynamicColors(true)
	sideBar := newPrimitive("Side Bar")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(0, -4, 0).
		SetBorders(true).
		AddItem(newPrimitive("Statistics "), 0, 0, 1, 3, 0, 0, true).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
		AddItem(main, 1, 0, 1, 3, 0, 0, false).
		AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	// Layout for screens wider than 100 cells.
	grid.AddItem(menu, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(sideBar, 1, 2, 1, 1, 0, 100, false)

	modal := tview.NewModal().
		SetText("Resize the window to see how the grid layout adapts").
		AddButtons([]string{"Ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.HidePage("modal")
	})

	pages.AddPage("grid", grid, true, true).
		AddPage("modal", modal, false, false)

	//barText(main, 5)
	
    data := []float64{}
    chain = cc.NewBlockchain(dbFile)
    defer chain.CloseDB()
    bci := chain.Iterator()
    for {
    	block := bci.Next()
    	diff := math.Log(targetToDifficultyFloat64(block.Target))
    	data = append(data, diff)
    	data = append(data, diff)
    	if len(block.PrevBlockHash) == 0 {
			break
		}
    }
    
    graph := asciigraph.Plot(invertArray(data))

    fmt.Fprint(main, graph)
	return "Graphs", pages
}

//█▅
