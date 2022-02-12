# Golang implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life)

**This is my first non trivial Golang program, implementing 1d and 2d Conway's Game of Life.**

`X` indicates "Dead" cells, while `O` indicates "Live" ones.

## 1D

This "unofficial" version of GoF was implemented as a first step in this project.

### Rules
Neighbours (`N`) of a cell are their immediate neighbours (i.e. left and right cells).

Edge cells' neihbours include the "wrapped around" cell on its left or right.

### State diagram
<img src="states1d.svg" width="400">

### Run
``` bash
go run cmd/1d/main.go --cellCount 10 --stepCount 3
```

### Output example
```
Game of Life
=== Step 0 ===
O X O O X O X O X O

=== Step 1 ===
O X O O X X X X X O

=== Step 2 ===
O X O O X X X X X O

=== Step 3 ===
O X O O X X X X X O
Done
```