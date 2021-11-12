// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Piece is just a piece
type Piece struct {
	name  string
	color string
}

// Tile does stuff
type Tile struct {
	held *Piece
}

// Board holds tiles
type Board struct {
	board [8][8]Tile
}

func (b *Board) config() {

	// [row][column]
	b.board = [8][8]Tile{
		// Row 1 [0][0-7]
		{
			Tile{&Piece{"Rook", "White"}}, Tile{&Piece{"Bishop", "White"}}, Tile{&Piece{"Knight", "White"}}, Tile{&Piece{"King", "White"}}, Tile{&Piece{"Queen", "White"}}, Tile{&Piece{"Knight", "White"}}, Tile{&Piece{"Bishop", "White"}}, Tile{&Piece{"Rook", "White"}},
		},

		// Row 2 [1][0-7]
		{
			//Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}},
			Tile{}, Tile{&Piece{"Pawn", "White"}}, Tile{&Piece{"Pawn", "White"}}, Tile{&Piece{"Pawn", "White"}}, Tile{&Piece{"Pawn", "White"}}, Tile{&Piece{"Pawn", "White"}}, Tile{&Piece{"Pawn", "White"}}, Tile{&Piece{"Pawn", "White"}},
		},

		// Row 3 [2][0-7]
		{
			Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{},
		},

		// Row 4 [3][0-7]
		{
			Tile{}, Tile{}, Tile{}, Tile{&Piece{"Rook", "White"}}, Tile{}, Tile{}, Tile{}, Tile{},
		},

		// Row 5 [4][0-7]
		{
			Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{},
		},

		// Row 6 [5][0-7]
		{
			Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{},
		},

		// Row 7 [6][0-7]
		{
			Tile{&Piece{"Pawn", "Black"}}, Tile{&Piece{"Pawn", "Black"}}, Tile{&Piece{"Pawn", "Black"}}, Tile{&Piece{"Pawn", "Black"}}, Tile{&Piece{"Pawn", "Black"}}, Tile{&Piece{"Pawn", "Black"}}, Tile{&Piece{"Pawn", "Black"}}, Tile{&Piece{"Pawn", "Black"}},
		},

		// Row 8 [7][0-7]
		{
			Tile{&Piece{"Rook", "Black"}}, Tile{&Piece{"Bishop", "Black"}}, Tile{&Piece{"Knight", "Black"}}, Tile{&Piece{"King", "Black"}}, Tile{&Piece{"Queen", "Black"}}, Tile{&Piece{"Knight", "Black"}}, Tile{&Piece{"Bishop", "Black"}}, Tile{&Piece{"Rook", "Black"}},
		},
	}

}

type Vertex struct {
	x, y int
}

func (b *Board) getTile(x int, y int) *Tile {
	foo := &b.board[y][x]
	return foo
}

func (b *Board) getTileLocation(t *Tile) Vertex {
	var foo *Tile

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			foo = b.getTile(x, y)
			if foo == t {
				return Vertex{x, y}
			}
		}
	}
	return Vertex{-1, -1}
}

func (b *Board) swapPieces(t1 *Tile, t2 *Tile) {
	temp := t1.held
	t1.held = t2.held
	t2.held = temp
}

func (b *Board) hardMove(t1 *Tile, t2 *Tile) {
	fmt.Println("T1 Held: ", t1.held, "\nT2 Held: ", t2.held)
	newTile := Tile{t1.held}
	t1.held = &Piece{}
	t2.held = newTile.held
	fmt.Println("T1 Held: ", t1.held, "\nT2 Held: ", t2.held)
}

// Give tile of piece you want to move and tile you want to move to
func (b *Board) movePiece(t1 *Tile, t2 *Tile) {
	validMoves := getPieceMoves(t1.held, b)
	has := false
	for _, element := range validMoves {
		if element == t2 {
			has = true
		}
	}

	if has {
		fmt.Println("Valid Move")
		b.hardMove(t1, t2)
		//b.swapPieces(t1, t2)
	} else {
		fmt.Println("Invalid Move")
	}
}

func (p *Piece) getMoves() {
	//fmt.Println("Hello!")
}

// For later, add a check to prevent duplicates
func getPieceMoves(p *Piece, b *Board) []*Tile {
	x, y := getPieceLocation(p, b)

	x1, y1 := 0, 0

	var moves []*Tile
	var foo *Tile

	if p.name == "Rook" {

		// Tile Distance
		leftTiles := x
		rightTiles := 8 - x - 1
		upTiles := y
		downTiles := 8 - y - 1
		//fmt.Println(leftTiles, rightTiles, upTiles, downTiles)

		x1 = 1
		for leftTiles > 0 {
			foo = b.getTile(x-x1, y)
			if foo.held != nil {
				if foo.held.color != p.color {
					moves = append(moves, foo)
				}
				//fmt.Println("Left stopped at ", x-x1)
				leftTiles = 0
			} else {
				moves = append(moves, foo)
			}
			leftTiles--
			x1++
		}

		x1 = 1
		for rightTiles >= 0 {
			foo = b.getTile(x+x1, y)
			if foo.held != nil {
				if foo.held.color != p.color {
					moves = append(moves, foo)
				}
				//fmt.Println("Right stopped at ", x+x1)
				rightTiles = 0
			} else {
				moves = append(moves, foo)
			}
			rightTiles--
			x1++
		}

		y1 = 1
		for upTiles > 0 {
			foo = b.getTile(x, y-y1)
			if foo.held != nil {
				if foo.held.color != p.color {
					moves = append(moves, foo)
				}
				//fmt.Println("Up stopped at ", y-y1)
				upTiles = 0
			} else {
				moves = append(moves, foo)
			}
			upTiles--
			y1++
		}

		y1 = 1
		for downTiles >= 0 {
			foo = b.getTile(x, y+y1)
			if foo.held != nil {
				if foo.held.color != p.color {
					moves = append(moves, foo)
				}
				//fmt.Println("Down stopped at ", y+y1)
				downTiles = 0
			} else {
				moves = append(moves, foo)
			}
			downTiles--
			y1++
		}

		/*for i := 0; i < 8; i++ {
			foo = b.getTile(i, y)
			if b.getTile(i, y).held == nil {
				moves = append(moves, foo)
			}

			foo = b.getTile(x, i)
			if b.getTile(x, i).held == nil {
				moves = append(moves, foo)
			}
		}*/
	}

	return moves
}

func getPieceLocation(p *Piece, b *Board) (int, int) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if b.getTile(x, y).held == p {
				return x, y
			}
		}
	}
	return -1, -1
}

func main() {
	fmt.Println("Hello, World!")

	b := Board{}
	b.config()

	moves := getPieceMoves(b.getTile(0, 0).held, &b)

	for _, element := range moves {
		fmt.Println(b.getTileLocation(element))
	}

	fmt.Println(b.getTile(0, 0).held, b.getTile(0, 6).held)
	fmt.Println(b.getTileLocation(getPieceMoves(b.getTile(0, 0).held, &b)[0]))
	//fmt.Println("Moves", getPieceMoves(b.getTile(0, 0).held, &b))
	b.movePiece(b.getTile(0, 0), b.getTile(0, 6))
	fmt.Println(b.getTile(0, 0).held, b.getTile(0, 6).held)

	//fmt.Println(b.getTileLocation(b.getTile(7, 7)))
	//for x := 0; x < 10; x++ {
	//	fmt.Println(generatePiece() + generateMove())
	//}
}

func generateMove() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	columns := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	row := r1.Intn(8) + 1
	text := columns[r1.Intn(8)] + strconv.Itoa(row)
	return text
}

func generatePiece() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	pieces := [16]string{"", "", "", "", "", "", "", "", "R", "R", "B", "B", "N", "N", "Q", "K"}
	return pieces[r1.Intn(16)]
}
