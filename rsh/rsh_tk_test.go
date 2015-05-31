package rsh

import (
	b "github.com/zaramme/KifCloud-Logic/board"
	def "github.com/zaramme/KifCloud-Logic/define"
	//	mv "../move"
	s "github.com/zaramme/KifCloud-Logic/structs"
	//	"fmt"
	"testing"
)

func Test_getAndputTK_全パターン(t *testing.T) {

	var rsh *RshCode

	testCase := func(black_x, black_y, white_x, white_y int, turn def.Player) {

		expectedBoard := b.NewBoard()
		//盤面の設定
		expectedBoard.Turn = turn
		posBlack := s.Position{black_x, black_y}
		posWhite := s.Position{white_x, white_y}
		expectedBoard.PositionMap[posBlack] = s.PieceStates{def.BLACK, false, def.OH}
		expectedBoard.PositionMap[posWhite] = s.PieceStates{def.WHITE, false, def.OH}

		rsh = NewRshCodeInit()

		// 盤面をコードに
		var err error
		rsh.Base_TK, rsh.Add_TK, err = getTKfromBoard(expectedBoard)
		if err != nil {
			t.Errorf("エラーを検出しました。... %e", err)
		}

		if len(rsh.Base_TK) != 2 {
			t.Errorf("tkの値が不正です。turn{%s},black{%d,%d},white{%d,%d} => tk = %s", turn.Output(), black_x, black_y, white_x, white_y, rsh.Base_TK.ToString())
			return
		}

		// コードから盤面を再現
		putPieceFromTKandAdd(rsh)

		// idの決定
		strByte := make([]byte, 0)

		strByte = append(strByte, []byte(posBlack.Output())...)
		strByte = append(strByte, []byte("+")...)
		strByte = append(strByte, []byte(posWhite.Output())...)

		assert_EachPositionOfBoard(expectedBoard, rsh.Board, string(strByte), t)
		assert_TurnOfBoard(expectedBoard, rsh.Board, t)
	}

	for bx := 1; bx <= 9; bx++ {
		for by := 1; by <= 9; by++ {
			for wx := 1; wx <= 9; wx++ {
				for wy := 1; wy <= 9; wy++ {
					if bx == wx && by == wy {
						continue
					}
					testCase(bx, by, wx, wy, def.BLACK)
					testCase(bx, by, wx, wy, def.WHITE)
				}
			}
		}
	}
}

func Test_convertTK_最大値(t *testing.T) {

	MAX := 45

	testCase := func(x, y int, player def.Player) {

		actual_code, _ := convertTKonIntAndAddFromPiece(s.Position{x, y}, player)

		if actual_code > MAX {
			t.Errorf("値が最大値%dを超えました, 条件：[%d,%d], 実測値：{%d}",
				MAX, x, y, actual_code)
		}
	}

	//　全パターン実行(盤面上の全座標＊先手後手)
	for x := 0; x < 9; x++ {
		for y := 0; y < 0; y++ {
			testCase(x, y, def.BLACK)
			testCase(x, y, def.WHITE)
		}
	}
}

func Test_convertTK_全パターン(t *testing.T) {

	testCase := func(x, y int, player def.Player, expected_code, expeted_add int) {

		actual_code, actual_add := convertTKonIntAndAddFromPiece(s.Position{x, y}, player)

		if expected_code != actual_code || expeted_add != actual_add {
			t.Errorf("値が期待値と異なっています。, 条件：[%d,%d], 期待値：{%d, %d}, 実測値：{%d,%d}",
				x, y, expected_code, expeted_add, actual_code, actual_add)
		}
	}

	testCase(1, 1, def.BLACK, 0, 1)
	testCase(1, 2, def.BLACK, 1, 1)
	testCase(1, 3, def.BLACK, 2, 1)
	testCase(1, 4, def.BLACK, 3, 1)
	testCase(1, 5, def.BLACK, 4, 1)
	testCase(1, 6, def.BLACK, 0, 0)
	testCase(1, 7, def.BLACK, 1, 0)
	testCase(1, 8, def.BLACK, 2, 0)
	testCase(1, 9, def.BLACK, 3, 0)

	testCase(9, 1, def.WHITE, 8*5+0, 0)
	testCase(9, 2, def.WHITE, 8*5+1, 0)
	testCase(9, 3, def.WHITE, 8*5+2, 0)
	testCase(9, 4, def.WHITE, 8*5+3, 0)
	testCase(9, 5, def.WHITE, 8*5+4, 0)
	testCase(9, 6, def.WHITE, 8*5+0, 1)
	testCase(9, 7, def.WHITE, 8*5+1, 1)
	testCase(9, 8, def.WHITE, 8*5+2, 1)
	testCase(9, 9, def.WHITE, 8*5+3, 1)
}
