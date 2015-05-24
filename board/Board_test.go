package board

import (
	def "KifuLibrary-Logic/define"
	m "KifuLibrary-Logic/move"
	s "KifuLibrary-Logic/structs"
	"testing"
)

func Test_boardInitiate(t *testing.T) {
	actual := NewBoardInit()
	actual.Turn = def.Player(def.BLACK)

	// １一〜９九までの座標を全て走査する
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			exists := true
			var kop def.KindOfPiece
			switch {
			case (i == 1 || i == 9) && (j == 1 || j == 9):
				kop = def.KindOfPiece(def.KYO)
			case (i == 2 || i == 8) && (j == 1 || j == 9):
				kop = def.KindOfPiece(def.KEI)
			case (i == 3 || i == 7) && (j == 1 || j == 9):
				kop = def.KindOfPiece(def.GIN)
			case (i == 4 || i == 6) && (j == 1 || j == 9):
				kop = def.KindOfPiece(def.KIN)
			case (i == 5) && (j == 1 || j == 9):
				kop = def.KindOfPiece(def.OH)
			case j == 3 || j == 7:
				kop = def.KindOfPiece(def.FU)
			case i == 2 && j == 2 || i == 8 && j == 8:
				kop = def.KindOfPiece(def.KAKU)
			case i == 2 && j == 8 || i == 8 && j == 2:
				kop = def.KindOfPiece(def.HISHA)
			case i == 2 && j == 8 || i == 8 && j == 2:
				kop = def.KindOfPiece(def.HISHA)
			default:
				exists = false // 駒が無い場合
			}

			pos := s.Position{i, j}
			if exists {
				player := def.Player(def.BLACK)
				if j <= 3 {
					player = def.Player(def.WHITE)
				}
				expected := s.PieceStates{player, false, kop}
				assert_PieceStatesEquealonBoard(expected, actual, pos, t)
			} else {
				assert_notExistsOnBoardPosition(actual, pos, t)
			}

		}
	}

	// 駒台の駒が無い事を確認する
	if len(actual.CapturedMap) != 0 {
		t.Errorf("駒台の初期化が正しく行われていません")
	}
}

func assert_PieceStatesEquealonBoard(expected s.PieceStates, board *Board, pos s.Position, t *testing.T) {

	if _, ok := board.PositionMap[pos]; !ok {
		t.Errorf("盤上の状態が期待値と一致しません", pos.X, pos.Y)
		t.Errorf("ーーーーーーーーーーーーーーーー期待値＝", expected.GetPieceName(), expected.IsPromoted, expected.Player)
	}

	if board.PositionMap[pos] != expected {
		ps := board.PositionMap[pos]
		t.Errorf("盤上の状態が期待値と一致しません", pos.X, pos.Y)
		t.Errorf("ーーーーーーーーーーーーーーーー実際値＝", ps.GetPieceName(), ps.IsPromoted, ps.Player)
		t.Errorf("ーーーーーーーーーーーーーーーー期待値＝", expected.GetPieceName(), expected.IsPromoted, expected.Player)
	}
}

func assert_notExistsOnBoardPosition(board *Board, pos s.Position, t *testing.T) {
	if _, ok := board.PositionMap[pos]; ok {
		t.Errorf("盤上の状態が期待値と一致しません")
	}
}

func Test_AddMove(t *testing.T) {
	var board *Board
	var pos s.Position
	var cap s.CapArea
	var move *m.Move
	var expectedPs s.PieceStates

	board = NewBoardInit()
	move = m.NewMoveFromMoveCode("b26FU_27")
	board.AddMove(move)
	expectedPs = s.PieceStates{def.BLACK, false, def.FU}
	pos = s.Position{2, 6}
	assert_PieceExistsOnPosition(board, pos, expectedPs, t)
	pos = s.Position{2, 5}
	assert_PieceNotExists(board, pos, t)

	// 駒を取る
	board = NewBoardInit()
	move = m.NewMoveFromMoveCode("b22KA_88")
	board.AddMove(move)
	expectedPs = s.PieceStates{def.BLACK, false, def.KAKU}
	pos = s.Position{2, 2}
	assert_PieceExistsOnPosition(board, pos, expectedPs, t)
	pos = s.Position{8, 8}
	assert_PieceNotExists(board, pos, t)
	cap = s.CapArea{def.BLACK, def.KAKU}
	assert_CapturedPieceNum(board, cap, 1, t)

	// 駒を打つ
	move = m.NewMoveFromMoveCode("b55KA_00")
	board.AddMove(move)
	expectedPs = s.PieceStates{def.BLACK, false, def.KAKU}
	pos = s.Position{5, 5}
	assert_PieceExistsOnPosition(board, pos, expectedPs, t)
	cap = s.CapArea{def.BLACK, def.KAKU}
	assert_CapturedPieceNotExists(board, cap, t)

	// 駒を成る
	board = NewBoardInit()
	move = m.NewMoveFromMoveCode("b53HI_28!")
	board.AddMove(move)
	expectedPs = s.PieceStates{def.BLACK, true, def.HISHA}
	pos = s.Position{5, 3}
	assert_PieceExistsOnPosition(board, pos, expectedPs, t)
	pos = s.Position{2, 8}
	assert_PieceNotExists(board, pos, t)
	cap = s.CapArea{def.BLACK, def.FU}
	assert_CapturedPieceNum(board, cap, 1, t)

}

func assert_PieceExistsOnPosition(board *Board, pos s.Position, expectedPs s.PieceStates, t *testing.T) {

	if value, ok := board.PositionMap[pos]; ok {
		if value != expectedPs {
			t.Errorf("駒の状態が期待値と一致しません。")
		}
	} else {
		t.Errorf("座標が期待値と一致しません。（%sは空座標です）", pos.Output())
	}
}

func assert_PieceNotExists(board *Board, pos s.Position, t *testing.T) {
	if _, ok := board.PositionMap[pos]; ok {
		t.Errorf("座標が期待値と一致しません。（%sに駒が存在します）", pos.Output())
	}
}

func assert_CapturedPieceNum(board *Board, cap s.CapArea, expectedNum int, t *testing.T) {
	if value, ok := board.CapturedMap[cap]; ok {
		if value != expectedNum {
			t.Errorf("駒台の状態が期待値と一致しません。")
		}
	} else {
		t.Errorf("座標が期待値と一致しません。（%sは空座標です）", cap.Output())
	}
}

func assert_CapturedPieceNotExists(board *Board, cap s.CapArea, t *testing.T) {
	if _, ok := board.CapturedMap[cap]; ok {
		t.Errorf("駒台の状態が期待値と一致しません。(%sに駒が存在します）", cap.Output())
	}
}

func assert_Player(board *Board, expected def.Player, t *testing.T) {
	if board.Turn != expected {
		if expected == def.BLACK {
			t.Errorf("手番が期待値と異なります。期待値…先手", expected)
		} else {
			t.Errorf("手番が期待値と異なります。期待値…後手", expected)
		}
	}
}
