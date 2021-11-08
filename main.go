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
			Tile{&Piece{"Rook", "White"}}, Tile{&Piece{"Bishop", "White"}}, Tile{Piece{"Knight", "White"}}, Tile{Piece{"King", "White"}}, Tile{Piece{"Queen", "White"}}, Tile{Piece{"Knight", "White"}}, Tile{Piece{"Bishop", "White"}}, Tile{Piece{"Rook", "White"}},
		},

		// Row 2 [1][0-7]
		{
			//Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}},
			Tile{}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}}, Tile{Piece{"Pawn", "White"}},
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
			Tile{Piece{"Pawn", "Black"}}, Tile{Piece{"Pawn", "Black"}}, Tile{Piece{"Pawn", "Black"}}, Tile{Piece{"Pawn", "Black"}}, Tile{Piece{"Pawn", "Black"}}, Tile{Piece{"Pawn", "Black"}}, Tile{Piece{"Pawn", "Black"}}, Tile{Piece{"Pawn", "Black"}},
		},

		// Row 8 [7][0-7]
		{
			Tile{Piece{"Rook", "Black"}}, Tile{Piece{"Bishop", "Black"}}, Tile{Piece{"Knight", "Black"}}, Tile{Piece{"King", "Black"}}, Tile{Piece{"Queen", "Black"}}, Tile{Piece{"Knight", "Black"}}, Tile{Piece{"Bishop", "Black"}}, Tile{Piece{"Rook", "Black"}},
		},
	}

}

func (b *Board) getTile(x int, y int) Tile {
	return b.board[y][x]
}

func (p *Piece) getMoves() {
	fmt.Println("Hello!")
}

func main() {
	fmt.Println("Hello, World!")

	// a 8x8 chessboard! we will store all pieces here.
	// pieces are saved like wK, wRa, bRh, etc
	// first lower case is the color piece
	// then the upper case is the actual peice type (none if pawn)
	//the last letter (lowercase) is the column it starts on.
	// wK is the white king, w = white, K = king
	// wRa is the white left rook, w = white, R = rook, a = leftmost column

	b := Board{}
	b.config()

	fmt.Println(b.getTile(0, 0))
	fmt.Println(b.getTile(0, 1))
	b.getTile(0, 0).held.getMoves()

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
