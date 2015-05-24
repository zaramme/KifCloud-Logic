package rsh

import (
	b "KifuLibrary-Logic/board"
	"KifuLibrary-Logic/code"
	def "KifuLibrary-Logic/define"
	s "KifuLibrary-Logic/structs"
	"testing"
)

func Test_BuildBoardFromRshCode_初期盤面(t *testing.T) {

	testCase := func(expectedBoard *b.Board) {

		rsh := ConvertRshFromBoard(expectedBoard)
		actualboard := BuildBoardFromRshCode(rsh)

		assert_EachPositionOfBoard(expectedBoard, actualboard, "init", t)
	}

	expectedBoard := b.NewBoardInit()

	testCase(expectedBoard)
}

func Test_BuildBoardFromRshCode_初期盤面全部成り駒(t *testing.T) {

	testCase := func(expectedBoard *b.Board) {

		rsh := ConvertRshFromBoard(expectedBoard)
		actualboard := BuildBoardFromRshCode(rsh)

		assert_EachPositionOfBoard(expectedBoard, actualboard, "initProm", t)
	}

	expectedBoard := b.NewBoardInit()

	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			pos := s.Position{x, y}
			if ps, ok := expectedBoard.PositionMap[pos]; !ok {
				continue
			} else {
				if ps.KindOfPiece != def.OH && ps.KindOfPiece != def.KIN {
					expectedBoard.PositionMap[pos] = s.PieceStates{
						Player:      ps.Player,
						KindOfPiece: ps.KindOfPiece,
						IsPromoted:  true,
					}
				}
			}
		}
	}

	testCase(expectedBoard)
}

func Test_rsh_putPieceFromM2_初期盤面(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	rsh.Base_M2 = getM2fromBoard(rsh.Board)

	rsh.Board = b.NewBoard()

	putPieceFromM2(rsh)

	assert_PieceExistsOnPosition(rsh.Board, s.Position{2, 8}, s.PieceStates{def.BLACK, false, def.HISHA}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{8, 2}, s.PieceStates{def.WHITE, false, def.HISHA}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{2, 2}, s.PieceStates{def.WHITE, false, def.KAKU}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{8, 8}, s.PieceStates{def.BLACK, false, def.KAKU}, t)
}

func Test_rsh_putPieceFromKIN_初期盤面(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()

	rsh.Base_KIN = getKINfromBoard(rsh.Board)

	rsh.Board = b.NewBoard()

	putPiecefromBase_KIN(rsh)

	assert_PieceExistsOnPosition(rsh.Board, s.Position{4, 1}, s.PieceStates{def.WHITE, false, def.KIN}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{4, 9}, s.PieceStates{def.BLACK, false, def.KIN}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{6, 1}, s.PieceStates{def.WHITE, false, def.KIN}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{6, 9}, s.PieceStates{def.BLACK, false, def.KIN}, t)
}

func Test_rsh_putPieceFromKIN_持ち駒込み(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	delete(rsh.Board.PositionMap, s.Position{4, 1})
	rsh.Board.CapturedMap[s.CapArea{def.BLACK, def.KIN}]++

	rsh.Base_KIN = getKINfromBoard(rsh.Board)

	rsh.Board = b.NewBoard()

	putPiecefromBase_KIN(rsh)

	assert_PieceNotExists(rsh.Board, s.Position{4, 1}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{4, 9}, s.PieceStates{def.BLACK, false, def.KIN}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{6, 1}, s.PieceStates{def.WHITE, false, def.KIN}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{6, 9}, s.PieceStates{def.BLACK, false, def.KIN}, t)

	assert_CapturedPieceNum(rsh.Board, s.CapArea{def.BLACK, def.KIN}, 1, t)
}

func Test_rsh_putPieceFromGIN_初期盤面(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()

	rsh.Base_GIN = getGINfromBoard(rsh.Board)

	rsh.Board = b.NewBoard()

	putPiecefromBase_GIN(rsh)

	assert_PieceExistsOnPosition(rsh.Board, s.Position{3, 1}, s.PieceStates{def.WHITE, false, def.GIN}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{3, 9}, s.PieceStates{def.BLACK, false, def.GIN}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{7, 1}, s.PieceStates{def.WHITE, false, def.GIN}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{7, 9}, s.PieceStates{def.BLACK, false, def.GIN}, t)
}

func Test_rsh_putPieceFromKEI_初期盤面(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()

	rsh.Base_KEI = getKEIfromBoard(rsh.Board)

	rsh.Board = b.NewBoard()

	putPiecefromBase_KEI(rsh)

	assert_PieceExistsOnPosition(rsh.Board, s.Position{2, 1}, s.PieceStates{def.WHITE, false, def.KEI}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{2, 9}, s.PieceStates{def.BLACK, false, def.KEI}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{8, 1}, s.PieceStates{def.WHITE, false, def.KEI}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{8, 9}, s.PieceStates{def.BLACK, false, def.KEI}, t)
}

func Test_rsh_putPieceFromKYO_初期盤面(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()

	rsh.Base_KYO = getKYOfromBoard(rsh.Board)

	rsh.Board = b.NewBoard()

	putPiecefromBase_KYO(rsh)

	assert_PieceExistsOnPosition(rsh.Board, s.Position{1, 1}, s.PieceStates{def.WHITE, false, def.KYO}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{1, 9}, s.PieceStates{def.BLACK, false, def.KYO}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{9, 1}, s.PieceStates{def.WHITE, false, def.KYO}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{9, 9}, s.PieceStates{def.BLACK, false, def.KYO}, t)
}

func Test_rsh_putPieceFromP16_Position_1(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoard()

	input := code.NewCode64FromInt(123456789)
	putPiecefromP16_Position(input, def.BLACK, rsh.Board)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{1, 9}, s.PieceStates{def.BLACK, false, def.FU}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{2, 8}, s.PieceStates{def.BLACK, false, def.FU}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{3, 7}, s.PieceStates{def.BLACK, false, def.FU}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{4, 6}, s.PieceStates{def.BLACK, false, def.FU}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{5, 5}, s.PieceStates{def.BLACK, false, def.FU}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{6, 4}, s.PieceStates{def.BLACK, false, def.FU}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{7, 3}, s.PieceStates{def.BLACK, false, def.FU}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{8, 2}, s.PieceStates{def.BLACK, false, def.FU}, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{9, 1}, s.PieceStates{def.BLACK, false, def.FU}, t)
}

func Test_rsh_putPieceFromP16_Position_2(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoard()

	input := code.NewCode64FromInt(90909090)
	putPiecefromP16_Position(input, def.WHITE, rsh.Board)
	assert_PieceNotExistsOnVarticalLine(rsh.Board, 1, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{2, 9}, s.PieceStates{def.WHITE, false, def.FU}, t)
	assert_PieceNotExistsOnVarticalLine(rsh.Board, 3, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{4, 9}, s.PieceStates{def.WHITE, false, def.FU}, t)
	assert_PieceNotExistsOnVarticalLine(rsh.Board, 5, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{6, 9}, s.PieceStates{def.WHITE, false, def.FU}, t)
	assert_PieceNotExistsOnVarticalLine(rsh.Board, 7, t)
	assert_PieceExistsOnPosition(rsh.Board, s.Position{8, 9}, s.PieceStates{def.WHITE, false, def.FU}, t)
	assert_PieceNotExistsOnVarticalLine(rsh.Board, 9, t)
}

func Test_rsh_putPieceFromP16_Captured_baseの範囲(t *testing.T) {

	testCase := func(P18Cap, expBlack, expWhite int) {
		rsh := new(RshCode)
		rsh.Board = b.NewBoard()
		rsh.Base_P18Cap = code.NewCode64FromInt(P18Cap)
		putPiecefromP16_Captured(rsh.Base_P18Cap, code.NewCode64FromInt(0), rsh.Board)

		assert_CapturedPieceNum(rsh.Board, s.CapArea{def.BLACK, def.FU}, expBlack, t)
		assert_CapturedPieceNum(rsh.Board, s.CapArea{def.WHITE, def.FU}, expWhite, t)
	}

	testCase(1, 1, 0)
	testCase(2, 2, 0)
	testCase(3, 3, 0)
	testCase(4, 4, 0)
	testCase(5, 5, 0)
	testCase(6, 6, 0)
	testCase(7, 7, 0)
	testCase(8, 0, 1)
	testCase(9, 1, 1)
	testCase(10, 2, 1)
	testCase(11, 3, 1)
	testCase(12, 4, 1)
	testCase(13, 5, 1)
	testCase(14, 6, 1)
	testCase(15, 7, 1)
	testCase(16, 0, 2)
	testCase(17, 1, 2)
	testCase(18, 2, 2)
	testCase(19, 3, 2)
	testCase(50, 2, 6)
	testCase(51, 3, 6)
	testCase(52, 4, 6)
	testCase(53, 5, 6)
	testCase(54, 6, 6)
	testCase(55, 7, 6)
	testCase(56, 0, 7)
	testCase(57, 1, 7)
	testCase(58, 2, 7)
	testCase(59, 3, 7)
	testCase(60, 4, 7)
	testCase(61, 5, 7)
	testCase(62, 6, 7)
	testCase(63, 7, 7)

}

func Test_rsh_putPieceFromP16_Captured_add使用(t *testing.T) {

	testCase := func(P18Cap, P18Add, expBlack, expWhite int) {
		rsh := new(RshCode)
		rsh.Board = b.NewBoard()
		rsh.Base_P18Cap = code.NewCode64FromInt(P18Cap)
		rsh.Add_P18ExCap = code.NewCode64FromInt(P18Add)
		putPiecefromP16_Captured(rsh.Base_P18Cap, rsh.Add_P18ExCap, rsh.Board)

		assert_CapturedPieceNum(rsh.Board, s.CapArea{def.BLACK, def.FU}, expBlack, t)
		assert_CapturedPieceNum(rsh.Board, s.CapArea{def.WHITE, def.FU}, expWhite, t)
	}

	testCase(1, 1, 9, 0)
	testCase(1, 2, 1, 8)
	testCase(1, 3, 9, 8)
	testCase(1, 4, 17, 0)
	testCase(1, 5, 1, 16)

	testCase(10, 1, 10, 1)
	testCase(10, 2, 2, 9)
	testCase(10, 3, 10, 9)
	testCase(10, 4, 18, 1)
	testCase(10, 5, 2, 17)

	testCase(2, 4, 18, 0)
	testCase(16, 5, 0, 18)

}

func Test_reverse164ToPiecePosition(t *testing.T) {
	var actPos s.Position
	var actPlayer def.Player
	actPos, actPlayer = reverse164ToPiecePosition(30)

	if actPos.X != 4 || actPos.Y != 4 {
		t.Errorf("値が期待値と異なっています", actPos.X, actPos.Y)
	}

	if actPlayer != def.Player(def.BLACK) {
		t.Errorf("値が期待値と異なっています", actPos.X, actPos.Y)
	}

	actPos, actPlayer = reverse164ToPiecePosition(111)

	if actPos.X != 4 || actPos.Y != 4 {
		t.Errorf("値が期待値と異なっています", actPos.X, actPos.Y)
	}

	if actPlayer != def.Player(def.WHITE) {
		t.Errorf("値が期待値と異なっています", actPos.X, actPos.Y)
	}

	actPos, actPlayer = reverse164ToPiecePosition(162)

	if actPos.X != 0 || actPos.Y != 0 {
		t.Errorf("値が期待値と異なっています", actPos.X, actPos.Y)
	}

	if actPlayer != def.Player(def.BLACK) {
		t.Errorf("値が期待値と異なっています", actPos.X, actPos.Y)
	}

	actPos, actPlayer = reverse164ToPiecePosition(163)

	if actPos.X != 0 || actPos.Y != 0 {
		t.Errorf("値が期待値と異なっています", actPos.X, actPos.Y)
	}

	if actPlayer != def.Player(def.WHITE) {
		t.Errorf("値が期待値と異なっています", actPos.X, actPos.Y)
	}
}

func assert_PieceExistsOnPosition(board *b.Board, pos s.Position, expectedPs s.PieceStates, t *testing.T) {

	if value, ok := board.PositionMap[pos]; ok {
		if value != expectedPs {
			t.Errorf("駒の状態が期待値と一致しません。")
		}
	} else {
		t.Errorf("座標が期待値と一致しません。（%sは空座標です）", pos.Output())
	}
}

func assert_PieceNotExists(board *b.Board, pos s.Position, t *testing.T) {
	if _, ok := board.PositionMap[pos]; ok {
		t.Errorf("座標が期待値と一致しません。（%sに駒が存在します）", pos.Output())
	}
}

func assert_PieceNotExistsOnVarticalLine(board *b.Board, posX int, t *testing.T) {
	for posY := 1; posY <= 9; posY++ {
		assert_PieceNotExists(board, s.Position{posX, posY}, t)
	}
}

func assert_TurnOfBoard(exp *b.Board, act *b.Board, t *testing.T) {
	if exp.Turn != act.Turn {
		t.Errorf("boardの手番情報が期待値と一致しません。期待値= %d, 実測値 = %d", exp.Turn, act.Turn)
	}
}

func assert_CapturedPieceNum(board *b.Board, cap s.CapArea, expectedNum int, t *testing.T) {
	if value, ok := board.CapturedMap[cap]; ok {
		if value != expectedNum {
			t.Errorf("駒台の状態が期待値と一致しません。期待値 = %d, 実測値 = %d", expectedNum, value)
		}
	} else {
		t.Errorf("座標が期待値と一致しません。（%sは空座標です）", cap.Output())
	}
}

// func assert_CapturedPieceNotExists(board *b.Board, cap s.CapArea, t *testing.T) {
// 	if _, ok := board.CapturedMap[cap]; ok {
// 		t.Errorf("駒台の状態が期待値と一致しません。(%sに駒が存在します）", cap.Output())
// 	}
// }
