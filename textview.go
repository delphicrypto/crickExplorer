package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	cc "github.com/delphicrypto/blockchain_go"
)

// TextView1 demonstrates the basic text view.
func TextView1(nextSlide func()) (title string, content tview.Primitive) {
	textView := tview.NewTextView().
		SetTextColor(tcell.ColorYellow).
		SetScrollable(false).
		SetDoneFunc(func(key tcell.Key) {
			nextSlide()
		})
	textView.SetChangedFunc(func() {
		if textView.HasFocus() {
			app.Draw()
		}
	})
	go func() {
		height := 0
		for {
			chain = cc.NewBlockchain(dbFile)
			bci := chain.Iterator()
			block := bci.Next()
			if block.Height > height {
				height = block.Height
				textView.SetText(fmt.Sprintf("Block %d: %x \n", height, block.Hash))
			}
			
			chain.CloseDB()

			time.Sleep(200 * time.Millisecond)
		}
	}()
	textView.SetBorder(true).SetTitle("TextView implements io.Writer")
	return "List", textView
}