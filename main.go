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
			Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{}, Tile{},
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

func (p *Piece) getMoves() {
	//fmt.Println("Hello!")
}

// For later, add a check to prevent duplicates
func getPieceMoves(p *Piece, b *Board) []*Tile {
	x, y := getPieceLocation(p, b)

	var moves []*Tile
	var foo *Tile

	if p.name == "Rook" {
		for i := 0; i < 8; i++ {
			foo = b.getTile(i, y)
			if b.getTile(i, y).held == nil {
				moves = append(moves, foo)
			}

			foo = b.getTile(x, i)
			if b.getTile(x, i).held == nil {
				moves = append(moves, foo)
			}
		}
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

	fmt.Println(getPieceMoves(b.getTile(0, 0).held, &b))

	moves := getPieceMoves(b.getTile(0, 0).held, &b)

	for _, element := range moves {
		fmt.Println(b.getTileLocation(element))
	}

	t1 := b.getTile(5, 5)
	t2 := b.getTile(4, 4)
	fmt.Println(&t1)
	fmt.Println(&t2)
	// WOO HOO!
	if &t1 == &t2 {
		fmt.Println("Shit.")
	}

	foo := b.getTile(3, 4)
	fmt.Println(b.getTileLocation(foo))

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
