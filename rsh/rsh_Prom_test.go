package rsh

import (
	b "github.com/zaramme/KifCloud-Logic/board"
	"github.com/zaramme/KifCloud-Logic/code"
	def "github.com/zaramme/KifCloud-Logic/define"
	"github.com/zaramme/KifCloud-Logic/move"
	//	mv "../move"
	"fmt"
	s "github.com/zaramme/KifCloud-Logic/structs"
	"testing"
)

func __rsh_Prom_test() {
	ps := s.PieceStates{def.BLACK, true, def.FU}
	fmt.Printf("test", ps)
}

func Test_getAndPutProm_初期配置の全パターン(t *testing.T) {

	testCase := func(promList []bool) {
		expectedBoard := b.NewBoardInit()

		m := move.NewMoveFromMoveCode("b19KY_19!")
		expectedBoard.AddMove(m)

		m = move.NewMoveFromMoveCode("b89KE_89!")
		expectedBoard.AddMove(m)

		m = move.NewMoveFromMoveCode("w22KA_22!")
		expectedBoard.AddMove(m)

		rsh := NewRshCodeInit()
		rsh.Board = b.NewBoardInit()
		var err error
		rsh.Add_Prom, err = getPromfromBoard(expectedBoard)
		if err != nil {
			t.Errorf("エラーを検出しました。 error = [ %s ]", err)
		}

		//fmt.Printf("addprom = %d\n", rsh.Add_Prom.ToInt())
		applyPiecesPromoted(rsh)

		assert_EachPositionOfBoard(expectedBoard, rsh.Board, "t1", t)
	}

	testCase(make([]bool, 0))

}

func Test_getAndPutProm_Rsh経由変換(t *testing.T) {

	brd := b.NewBoardInit()
	rsh, err := ConvertRshFromBoard(brd)
	if err != nil {
		t.Errorf("エラーを検出しました。 error = [ %s ]", err)
	}
	max := 64 * 64
	for i := 0; i < max; i++ {
		rsh.Add_Prom, _ = code.NewCode64FromInt(i)
		prev, err := rsh.ToString()
		if err != nil {
			t.Errorf("エラーを検出しました。 error = [ %s ]", err)
			return
		}

		rsh2, err := NewRshCodeFromString(prev)
		if err != nil {
			t.Errorf("エラーを検出しました。 error = [ %s ]", err)
			return
		}

		curr, err := rsh2.ToString()
		if err != nil {
			t.Errorf("エラーを検出しました。 error = [ %s ]", err)
		}
		if prev != curr {
			t.Errorf("Add_Promの変換値が異なっています。\nprev= %s, \ncurr= %s", prev, curr)
			return
		}

		//fmt.Printf("%s\n", curr)
		brdprom := BuildBoardFromRshCode(rsh2)
		for _, ps := range brdprom.PositionMap {
			if ps.IsPromoted {
				//fmt.Printf("---- %s, %s\n", pos.Output(), ps.KindOfPiece.ToString())
			}
		}
	}

}
