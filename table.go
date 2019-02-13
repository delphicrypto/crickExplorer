package main

import (
	"fmt"
	"strings"
    "time"
    "log"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
    cc "github.com/delphicrypto/blockchain_go"

)

const headers = `Block|Hash|Difficulty|Problem|Solution (k)|Txs`
const maxBytes = 6



// Table demonstrates the Table.
func Table(nextSlide func()) (title string, content tview.Primitive) {
	table := tview.NewTable().
		SetFixed(1, 1)
    for column, cell := range strings.Split(headers, "|") {
        color := tcell.ColorYellow
        align := tview.AlignCenter
        
        tableCell := tview.NewTableCell(cell).
            SetTextColor(color).
            SetAlign(align).
            SetSelectable(false)
        table.SetCell(0, column, tableCell)
    }

	table.SetBorder(true).SetTitle("Crick Chain")
    table.SetBorders(false).
            SetSelectable(true, false).
            SetSeparator('|')

	code := tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	code.SetBorderPadding(1, 1, 2, 0)

    app.SetFocus(table)
    go func() {
        height := -1
        for {
            chain = cc.NewBlockchain(dbFile)
            bci := chain.Iterator()
            if chain.GetBestHeight() > height {
                height = chain.GetBestHeight()
                for {
                    block := bci.Next()
                                   
                    for column, _ := range strings.Split(headers, "|") {
                        color := tcell.ColorWhite
                        if column == 0 {
                            color = tcell.ColorDarkCyan
                        }
                        align := tview.AlignCenter
                        var cell string
                        switch column {
                            case 0: 
                                cell = fmt.Sprintf("%d", block.Height) 
                            case 1: 
                                cell = fmt.Sprintf("%x..", block.Hash[:maxBytes])
                            case 2: 
                                cell = fmt.Sprintf("%d", targetToDifficulty(block.Target))
                            case 3: 
                                if len(block.ProblemGraphHash) >= maxBytes {
                                    cell = fmt.Sprintf("%x..", block.ProblemGraphHash[:maxBytes])
                                } else {
                                        cell = fmt.Sprintf("%x", block.ProblemGraphHash)
                                }
                            case 4: 
                                cell = fmt.Sprintf("%d", len(block.Solution))
                            case 5: 
                                cell = fmt.Sprintf("%d", len(block.Transactions))
                        }
                        tableCell := tview.NewTableCell(cell).
                            SetTextColor(color).
                            SetAlign(align).
                            SetSelectable(column != 0)
                        // if (column == 2) || (column == 3) || (column == 4) {
                        //     tableCell.SetExpansion(1)
                        // }
                        table.SetCell(block.Height + 1, column, tableCell)
                    }
                    if len(block.PrevBlockHash) == 0 {
                        break
                    } 
                }
                //to print block on selection func (*Table) SetSelectedFunc
                //to print block on move func (*Table) SetSelectionChangedFunc
                //print it with
                table.SetSelectedFunc(func(row int, column int) {
                    b, err := chain.GetBlockFromHeight(row - 1)
                    if err != nil {
                        log.Panic(err)
                    }
                    fmt.Fprint(code, fmt.Sprintf("%x", b.Hash))
                })
                
                //update
                app.Draw()
            }
            chain.CloseDB()
            time.Sleep(500 * time.Millisecond)
        }
    }()

	return "Table", tview.NewFlex().
		AddItem(table, 0, 1, false).
		AddItem(code, codeWidth, 1, false)
}
