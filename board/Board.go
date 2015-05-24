package board

import (
	def "KifuLibrary-Logic/define"

	m "KifuLibrary-Logic/move"
	s "KifuLibrary-Logic/structs"
	"errors"
	//"fmt"
)

type Board struct {
	// RSH code.rsh
	PositionMap s.PositionMap
	CapturedMap s.CapturedMap
	Turn        def.Player
}

func NewBoardInit() *Board {
	board := NewBoard()
	board.init()
	return board
}

func NewBoard() *Board {
	board := new(Board)
	board.PositionMap = make(s.PositionMap)
	board.CapturedMap = make(s.CapturedMap)
	board.Turn = def.Player(def.BLACK)
	return board

}

// func newBoardFromRSH() *Board {
// 	return null
// }

//初期盤面の生成
func (board Board) init() {
	// 一段目
	board.PositionMap[s.Position{9, 1}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.KYO)}
	board.PositionMap[s.Position{8, 1}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.KEI)}
	board.PositionMap[s.Position{7, 1}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.GIN)}
	board.PositionMap[s.Position{6, 1}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.KIN)}
	board.PositionMap[s.Position{5, 1}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.OH)}
	board.PositionMap[s.Position{4, 1}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.KIN)}
	board.PositionMap[s.Position{3, 1}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.GIN)}
	board.PositionMap[s.Position{2, 1}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.KEI)}
	board.PositionMap[s.Position{1, 1}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.KYO)}

	// 二段目
	board.PositionMap[s.Position{8, 2}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.HISHA)}
	board.PositionMap[s.Position{2, 2}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.KAKU)}

	// 三段目
	board.PositionMap[s.Position{9, 3}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{8, 3}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{7, 3}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{5, 3}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{6, 3}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{4, 3}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{3, 3}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{2, 3}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{1, 3}] = s.PieceStates{def.WHITE, false, def.KindOfPiece(def.FU)}

	// 七段目
	board.PositionMap[s.Position{9, 7}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{8, 7}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{7, 7}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{5, 7}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{6, 7}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{4, 7}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{3, 7}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{2, 7}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.FU)}
	board.PositionMap[s.Position{1, 7}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.FU)}

	// 八段目
	board.PositionMap[s.Position{8, 8}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.KAKU)}
	board.PositionMap[s.Position{2, 8}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.HISHA)}

	// 九段目
	board.PositionMap[s.Position{1, 9}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.KYO)}
	board.PositionMap[s.Position{2, 9}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.KEI)}
	board.PositionMap[s.Position{3, 9}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.GIN)}
	board.PositionMap[s.Position{4, 9}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.KIN)}
	board.PositionMap[s.Position{5, 9}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.OH)}
	board.PositionMap[s.Position{6, 9}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.KIN)}
	board.PositionMap[s.Position{7, 9}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.GIN)}
	board.PositionMap[s.Position{8, 9}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.KEI)}
	board.PositionMap[s.Position{9, 9}] = s.PieceStates{def.BLACK, false, def.KindOfPiece(def.KYO)}
}

func (board Board) GetHashID() {

}

// 局面を進める
func (board *Board) AddMove(move *m.Move) error {

	//fmt.Printf("[AddMove開始]///%s\n", move.ToMoveCode())

	if move.IsResigned {
		return nil
	}

	// Prevの処理
	if move.Prev.IsCaptured() {
		//		fmt.Printf("---- 駒を打つ処理です。= %s\n", move.ToMoveCode())
		// 持ち駒から駒を消去
		CapAreaToRemove := m.NewCapAreafromMove(move)
		//		fmt.Printf("---- capArea = %s  / value = %d \n", CapAreaToRemove.Output(), board.CapturedMap[CapAreaToRemove])
		if board.CapturedMap[CapAreaToRemove] == 0 {
			return errors.New("AddMove：盤面の整合性エラーが発生しました(CODE01)")
		}
		board.CapturedMap[CapAreaToRemove] -= 1
		if board.CapturedMap[CapAreaToRemove] == 0 {
			delete(board.CapturedMap, CapAreaToRemove)
		}
	} else {
		// 盤上から駒を消去
		PositionToRemove := move.Prev
		if _, ok := board.PositionMap[PositionToRemove]; !ok {
			return errors.New("AddMove：盤面の整合性エラーが発生しました(CODE02)")
		} else {
			delete(board.PositionMap, PositionToRemove)
		}
	}

	// 駒を取る処理
	PositionToMove := move.Next
	if _, ok := board.PositionMap[PositionToMove]; ok {
		// 駒取りが発生する場合の処理
		//fmt.Printf("駒取りが発生しました\n")

		CapturedPiece := board.PositionMap[PositionToMove]
		//fmt.Printf("----CapturedPiece = %s \n", CapturedPiece)
		CapturedTo := s.NewCapAreafromPieceStates(CapturedPiece)
		CapturedTo.Player = !CapturedPiece.Player

		//fmt.Printf("----CapturedPiece(current) = %d \n", len(board.CapturedMap))
		board.CapturedMap[CapturedTo] += 1
		//fmt.Printf("----CapturedPiece(after) = %d \n", len(board.CapturedMap))

		delete(board.PositionMap, PositionToMove)
	}

	// Nextの処理
	board.PositionMap[PositionToMove] = s.PieceStates{
		Player:      move.Player,
		KindOfPiece: move.KindOfPiece,
		IsPromoted:  move.IsPromoted,
	}

	// 手番の処理
	board.Turn = !move.Player

	var sum int
	for _, i := range board.CapturedMap {
		sum += i
	}
	// エラー検出
	if len(board.PositionMap)+sum != 40 {
		return errors.New("AddMove：盤面の整合性エラーが発生しました(CODE99)")
	}

	return nil
}
