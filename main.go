// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Piece struct {
	name  string
	color string
}

type Tile struct {
	held interface{}
}

type Board struct {
	board [8][8]Tile
}

var threedim [5][10][4]int

func (b *Board) config() {

	// [row][column]
	b.board = [8][8]Tile{
		// Row 1 [0][0-7]
		{
			Tile{Piece{"Rook", "White"}}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil},
		},

		// Row 2 [1][0-7]
		{
			Tile{nil}, Tile{Piece{"Pawn", "White"}}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil},
		},

		// Row 3 [2][0-7]
		{
			Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil},
		},

		// Row 4 [3][0-7]
		{
			Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil},
		},

		// Row 5 [4][0-7]
		{
			Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil},
		},

		// Row 6 [5][0-7]
		{
			Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil},
		},

		// Row 7 [6][0-7]
		{
			Tile{nil}, Tile{Piece{"Pawn", "Black"}}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil},
		},

		// Row 8 [7][0-7]
		{
			Tile{Piece{"Rook", "Black"}}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil}, Tile{nil},
		},
	}

	//b.board = make([8][8]struct{}, 1)
	//b.board = make([8][8]interface{})

	//b.board[0][0] = nil

	//b.board[0] = [8]Piece{nil, nil, nil, nil, nil, nil, nil, nil}
	//[8]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	// b.board = [8][8]interface{
	// 	{Piece{"Pawn", "White"}, nil, nil, nil, nil, nil, nil, nil},
	// 	{nil, nil, nil, nil, nil, nil, nil, nil},
	// 	{nil, nil, nil, nil, nil, nil, nil, nil},
	// 	{nil, nil, nil, nil, nil, nil, nil, nil},
	// 	{nil, nil, nil, nil, nil, nil, nil, nil},
	// 	{nil, nil, nil, nil, nil, nil, nil, nil},
	// 	{nil, nil, nil, nil, nil, nil, nil, nil},
	// 	{nil, nil, nil, nil, nil, nil, nil, nil},
	// }
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

	fmt.Println(b)
	fmt.Println(b.board[0][0])
	fmt.Println(b.board[1][1])
	fmt.Println(b.board[6][1])
	fmt.Println(b.board[7][0])

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
