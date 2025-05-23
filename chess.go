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

func printBoard(b Board) {
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", 8-i)
		for j := 0; j < 8; j++ {
			if b[i][j] == nil {
				fmt.Print(". ")
			} else {
				fmt.Print(b[i][j].Color + b[i][j].Type + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println(" a b c d e f g h")
}

func parseMove (input string) (fromX, fromY, toX, toY int, err error) {
	if len(input) != 5 {
		return 0, 0, 0, 0, errors.New("Invalid input format.")
	}

	fromY = int (input[0] - 'a')
	fromX = 8 - int (input[1] - '0')
	toY = int (input[3] - 'a')
	toX = 8 - int (input[4] - '0')
	return
}

func movePiece(b *Board, fx, fy, tx, ty int) error {
	piece := b[fx][fy]
	if piece == nil {
		return errors.New ("No piece at source.")
	}

	if b[tx][ty] != nil && b[tx][ty].Color == piece.Color {
		return errors.New ("Can't capture your own piece.")
	}

	b[tx][ty] = piece
	b[fx][fy] = nil
	return nil
}

// main function

func main() {
	board := initBoard()
	scanner := bufio.NewScanner(os.Stdin)
	currentPlayer := "w"

	for {
		printBoard(board)
		fmt.Printf("%s's move: ", currentPlayer)
		scanner.Scan()
		input := scanner.Text()

		fx, fy, tx, ty, err := parseMove(input)
		if err != nil {
			fmt.Println("Invalid move:", err)
			continue
		}

		piece := board[fx][fy]
		if piece == nil || piece.Color != currentPlayer {
			fmt.Println("Not you piece!")
			continue
		}

		err = movePiece (&board, fx, fy, tx, ty)
		if err != nil {
			fmt.Println ("Error:", err)
			continue
		}

		currentPlayer = togglePlayer (currentPlayer)
	}
}

func togglePlayer (p string) string {
	if p == "w" {
		return "b"
	}
	return "w"
}