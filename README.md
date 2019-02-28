# crickExplorer
Block explorer for Crick blockchain

## Dependencies

```
go get github.com/gdamore/tcell
go get github.com/rivo/tview
go get github.com/fatih/color
```

## Launch 

Generate the blockchain.db file using [blockchain_go](https://github.com/delphicrypto/blockchain_go). For example, from the main folder,
```
export NODE_ID=3000
go run main.go
```
will generate blockchain_3000.db in the main folder. Point to that path to run the block explorer


```
go run *.go /path/to/blockchain.db
```

Note: you can do this in a second terminal while keeping open the blockchain, the blockexplorer will update in real time.

## TODO
