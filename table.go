package main

import (
	"fmt"
	"strings"
    "time"
    //"log"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
    cc "github.com/delphicrypto/blockchain_go"

)

const headers = `[::bu]Block|Hash|Time|Difficulty|Problem|Solution (k)|Txs`
const maxBytes = 6

func printBlock(code *tview.TextView, b cc.Block, bc *cc.Blockchain) {
    fmt.Fprint(code, fmt.Sprintf("[yellow::b]=============================== Block %d ===============================\n", b.Height))
    fmt.Fprint(code, fmt.Sprintln(""))
    fmt.Fprint(code, fmt.Sprintf("[blue]Hash:   %x\n", b.Hash))
    fmt.Fprint(code, fmt.Sprintf("Prev:   %x\n", b.PrevBlockHash))
    fmt.Fprint(code, fmt.Sprintf("Target: %064x\n", b.Target))
    fmt.Fprint(code, fmt.Sprintf("Difficulty: %d\n", targetToDifficulty(b.Target)))
    time := int64(0)
    if len(b.PrevBlockHash) > 0 {
        prevBlock, _ := bc.GetBlockFromHash(b.PrevBlockHash)
        time = (b.Timestamp - prevBlock.Timestamp) / 1e9
    }
    fmt.Fprint(code, fmt.Sprintf("Time: %d\n", time))
   
    validBlock := b.Validate(bc)
    if validBlock {
        fmt.Fprint(code, fmt.Sprintf("[green]Valid\n"))
    } else {
        fmt.Fprint(code, fmt.Sprintf("[red]Not valid\n"))
    }
    fmt.Fprint(code, fmt.Sprintln("[yellow]========================================================================"))
    fmt.Fprint(code, fmt.Sprintln(""))
    if len(b.SolutionHash) > 0 {
        fmt.Fprint(code, fmt.Sprintf("[blue]Solution to %x: \n", b.SolutionHash))
        fmt.Fprint(code, fmt.Sprintln(b.Solution))
        
        validSol := b.HasValidSolution(bc)
        if validSol {
            fmt.Fprint(code, fmt.Sprintf("[green]Valid\n"))
        } else {
            fmt.Fprint(code, fmt.Sprintf("[red]Not valid\n"))
        }
    } else {
         fmt.Fprint(code, fmt.Sprintf("No solution\n"))
    }
    fmt.Fprint(code, fmt.Sprintln(""))
    if len(b.ProblemGraphHash) > 0 {
        fmt.Fprint(code, fmt.Sprintf("[blue]New Problem %x \n", b.ProblemGraphHash))
    } else {
        fmt.Fprint(code, fmt.Sprintf("No problem\n"))
    }
    fmt.Fprint(code, fmt.Sprintln(""))
    for _, tx := range b.Transactions {
        fmt.Fprint(code, fmt.Sprintln("[yellow]", tx))
    }
}

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

    table.SetSelectable(true, false).
            SetSeparator(' ')
    app.SetFocus(table)
	code := tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	code.SetBorderPadding(1, 1, 2, 0)

   
    go func() {
        height := -1
        for {
            chain = cc.NewBlockchain(dbFile)
            bci := chain.Iterator()
            if chain.GetBestHeight() > height {
                for {
                    block := bci.Next()
                    if block.Height > height {
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
                                    time := int64(0)
                                    if len(block.PrevBlockHash) > 0 {
                                        prevBlock, _ := chain.GetBlockFromHash(block.PrevBlockHash)
                                        time = (block.Timestamp - prevBlock.Timestamp) / 1e9
                                    }
                                    cell = fmt.Sprintf("%d", time)
                                case 3: 
                                    cell = fmt.Sprintf("%d", targetToDifficulty(block.Target))
                                case 4: 
                                    if len(block.ProblemGraphHash) >= maxBytes {
                                        cell = fmt.Sprintf("%x..", block.ProblemGraphHash[:maxBytes])
                                    } else {
                                            cell = fmt.Sprintf("%x", block.ProblemGraphHash)
                                    }
                                case 5: 
                                    cell = fmt.Sprintf("%d", len(block.Solution))
                                case 6: 
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
                    } else {
                        break
                    }
                    if len(block.PrevBlockHash) == 0 {
                        break
                    } 
                }
                height = chain.GetBestHeight()
            }
            chain.CloseDB()
            table.SetSelectedFunc(func(row int, column int) {
                chain = cc.NewBlockchain(dbFile)
                b, err := chain.GetBlockFromHeight(row - 1)
                
                if err == nil {
                    code.Clear()
                    printBlock(code, b, chain)
                }
                chain.CloseDB()
            })
            
            //update
            app.Draw()
            
            time.Sleep(500 * time.Millisecond)
        }
    }()

	return "Crick Chain", tview.NewFlex().
		AddItem(table, 0, 1, true).
		AddItem(code, codeWidth, 1, false)
}
