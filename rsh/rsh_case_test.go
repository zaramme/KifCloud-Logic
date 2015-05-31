package rsh

import (
	"fmt"
	b "github.com/zaramme/KifCloud-Logic/board"
	def "github.com/zaramme/KifCloud-Logic/define"
	// mth "github.com/zaramme/KifCloud-Logic/math"
	mv "github.com/zaramme/KifCloud-Logic/move"
	s "github.com/zaramme/KifCloud-Logic/structs"
	"strconv"
	"testing"
)

func __rsh_case_test() {
	fmt.Print("this is dummy")
}
func Test_ケース１_成り駒のエンコーディング(t *testing.T) {
	//	fmt.Print("Test_ケース１_成り駒のエンコーディング-----------------------------\n")

	imput := "qEqZ-FWky_dcqUGC_Q7hmHpgUgGOrxpZh6ibg9"
	brd := strToBoard(imput, t)

	// ７九に後手の飛車がいる
	piece := assert_GetPiece(brd, s.Position{7, 9}, t)
	assert_PieceState(piece, def.HISHA, def.WHITE, false, t)

	// ▽69飛を着手
	move := mv.NewMoveFromMoveCode("w69HI_79!")
	brd.AddMove(move)

	// 成り駒の状態を確認
	assert_PieceNotExist(brd, s.Position{7, 9}, t)
	piece = assert_GetPiece(brd, s.Position{6, 9}, t)
	assert_PieceState(piece, def.HISHA, def.WHITE, true, t)

	// （１）Rshオブジェクト　→　Rsh文字列　→　boardオブジェクトの順で変換
	rsh, err := ConvertRshFromBoard(brd)
	if err != nil {
		t.Errorf("エラーを検出しました。 error = [ %s ]", err)
	}
	rshString, err := rsh.ToString()
	if err != nil {
		t.Errorf("エラーを検出しました。 error = [ %s ]", err)
	}
	brd2 := strToBoard(rshString, t)

	// 変換後、飛車の成り位置が変わらないことを確認
	piece = assert_GetPiece(brd2, s.Position{6, 9}, t)
	assert_PieceState(piece, def.HISHA, def.WHITE, true, t)

	// （１）の前後で全ての駒の位置に変化がないことを確認
	assert_EachPositionOfBoard(brd, brd2, "1", t)
	// // brd2 := rsh.ToString()
}

func Test_Case2_APIエラーパターン1(t *testing.T) {
}

func strToBoard(str string, t *testing.T) *b.Board {
	rsh, err := NewRshCodeFromString(str)
	if err != nil {
		t.Errorf(err.Error())
	}
	board := BuildBoardFromRshCode(rsh)
	return board
}

func assert_GetPiece(b *b.Board, pos s.Position, t *testing.T) s.PieceStates {
	piece, ok := b.PositionMap[pos]
	if !ok {
		t.Errorf("駒が所定の位置にありません。@%s", pos.Output())
	}
	return piece
}

func assert_PieceNotExist(b *b.Board, pos s.Position, t *testing.T) {
	_, ok := b.PositionMap[pos]
	if ok {
		t.Errorf("駒が意図しない位置にあります。@%s", pos.Output())
	}
}

func assert_PieceState(ps s.PieceStates, kop def.KindOfPiece, ply def.Player, isP def.IsPromoted, t *testing.T) {

	setError := func(exp string, act string) {
		t.Errorf("駒の状態が期待値と異なります。[exp]%v  | [act]%v", exp, act)
	}

	if ps.KindOfPiece != kop {
		setError(kop.ToString(), ps.KindOfPiece.ToString())
	}

	if ps.Player != ply {
		setError(string(ply.Output()), ps.Player.Output())
	}

	if ps.IsPromoted != isP {
		setError(strconv.FormatBool(bool(ps.IsPromoted)), strconv.FormatBool(bool(isP)))
	}

}
