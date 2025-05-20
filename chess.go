type Piece struct {
	Color string 
	Type string
}

type Board [8][8]*Piece

func initBoard() Board {
	var board Board

	for i := 0; i < 8; i++ {
		board[1][i] = &Piece{"b", "P"}
		board[6][i] = &Piece{"w", "P"}
	}

	backRank := []string {"R", "N", "B", "Q", "K", "B", "N", "R"}
	for i, p := range backRank {
		board[0][i] = &Piece{"b", p}
		board[7][i] = &Piece{"w", p}
	}

	return board
}