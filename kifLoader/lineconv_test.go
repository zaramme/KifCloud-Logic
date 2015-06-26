package kifLoader

import (
	"fmt"
	def "github.com/zaramme/KifCloud-Logic/define"
	m "github.com/zaramme/KifCloud-Logic/move"
	s "github.com/zaramme/KifCloud-Logic/structs"
	"testing"
)

const debug_convertKifCodeToMovecode = true

func __lineconv() {
	fmt.Print("")
}

func Test_convertKifCodeToMovecode(t *testing.T) {

	// テストケース
	caseExec := func(movecode string, expectPrev, expectNext s.Position, expectPieceStates s.PieceStates, isResigned bool) {
		desc, err := convertKifCodeToMoveCode(movecode, 1)

		if err != nil {
			t.Error(err.Error())
		}

		move, ok := desc.(*m.Move)

		if !ok {
			t.Errorf("move型に変換できません。", movecode)
			return
		}
		if move.Prev != (expectPrev) {
			t.Errorf("prevの値が期待値と異なっています。(%s)", movecode)
		}

		if move.Next != (expectNext) {
			t.Errorf("nextの値が期待値と異なっています。(%s)", movecode)
		}

		if move.PieceStates != (expectPieceStates) {
			t.Errorf("pieceStatesの値が期待値と異なっています,%s", move.PieceStates)
		}

		if move.IsResigned != isResigned {
			t.Errorf("isResignedの値が期待値と異なっています,%s", move.IsResigned)
		}
	}

	caseExec("２六歩(27)", s.Position{2, 7}, s.Position{2, 6}, s.PieceStates{true, false, def.FU}, false)
	caseExec("３四歩(33)", s.Position{3, 3}, s.Position{3, 4}, s.PieceStates{true, false, def.FU}, false)
	caseExec("３八銀(39)", s.Position{3, 9}, s.Position{3, 8}, s.PieceStates{true, false, def.GIN}, false)
	caseExec("６八銀(79)", s.Position{7, 9}, s.Position{6, 8}, s.PieceStates{true, false, def.GIN}, false)
	caseExec("３二成銀(31)", s.Position{3, 1}, s.Position{3, 2}, s.PieceStates{true, true, def.GIN}, false)
	caseExec("２七銀成(38)", s.Position{3, 8}, s.Position{2, 7}, s.PieceStates{true, true, def.GIN}, false)

	caseExec("1 ２六歩(27)", s.Position{2, 7}, s.Position{2, 6}, s.PieceStates{true, false, def.FU}, false)
	caseExec("12 ３四歩(33)", s.Position{3, 3}, s.Position{3, 4}, s.PieceStates{true, false, def.FU}, false)
	caseExec("150 ３八銀(39)", s.Position{3, 9}, s.Position{3, 8}, s.PieceStates{true, false, def.GIN}, false)
	caseExec("1309 ６八銀(79)", s.Position{7, 9}, s.Position{6, 8}, s.PieceStates{true, false, def.GIN}, false)

	// caseExec("３三角(22)", s.Potirion{3, 3}, s.Position{2, 2}, s.PieceStates{})
	// caseExec("４六歩(47)", s.Potirion{4, 6}, s.Position{4, 7}, s.PieceStates{})
	// caseExec("４三銀(32)", s.Potirion{4, 3}, s.Position{3, 2}, s.PieceStates{})
	// caseExec("１五歩(16)", s.Potirion{1, 5}, s.Position{1, 6}, s.PieceStates{})
	// caseExec("３二金(41)", s.Potirion{3, 2}, s.Position{4, 1}, s.PieceStates{})
	// caseExec("７七銀(68)", s.Potirion{7, 7}, s.Position{6, 8}, s.PieceStates{})
	// caseExec("５四銀(43)", s.Potirion{5, 4}, s.Position{4, 3}, s.PieceStates{})
	// caseExec("７八飛(28)", s.Potirion{7, 8}, s.Position{2, 8}, s.PieceStates{})
	// caseExec("２四角(33)", s.Potirion{2, 4}, s.Position{3, 3}, s.PieceStates{})
	// caseExec("４八飛(78)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("４二飛(82)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("３六銀(27)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("４三銀(54)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("５八金(49)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("７二銀(71)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("６八玉(59)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("６二玉(51)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("７八玉(68)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("７一玉(62)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("７九角(88)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("３五歩(34)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("４七銀(36)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("５四銀(43)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("５六歩(57)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("４三金(32)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("３六歩(37)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("３二飛(42)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("３八飛(48)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("３四金(43)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("８八玉(78)", s.Potirion{}, s.Position{}, s.PieceStates{})
	// caseExec("５一角(24)", s.Potirion{}, s.Position{}, s.PieceStates{})

}

func Test_convertKifCodeToMovecode_種別(t *testing.T) {
	caseExec := func(movecode string, objType string) {
		desc, err := convertKifCodeToMoveCode(movecode, 1)

		if err != nil {
			t.Errorf("[%s]error値が返されました。error = %s", movecode, err.Error())
			return
		}
		switch objType {
		case "move":
			if _, ok := desc.(*m.Move); ok {
			} else {
				t.Errorf("[%s]期待された型に変換されませんでした。 期待値 = %s", movecode, objType)
			}
			return
		case "repeatMove":
			if _, ok := desc.(*m.RepeatMove); ok {
			} else {
				t.Errorf("[%s]期待された型に変換されませんでした。 期待値 = %s", movecode, objType)
			}
			return
		case "endGame":
			if _, ok := desc.(*m.EndGame); ok {
			} else {
				t.Errorf("[%s]期待された型に変換されませんでした。 期待値 = %s", movecode, objType)
			}
			return
		}
		t.Errorf("[%s]不正な期待値です。 期待値 = %s", movecode, objType)
	}

	caseExec("３八銀(33)", "move")
	caseExec("同銀(33)", "repeatMove")
	caseExec("同成銀(33)", "repeatMove")
	caseExec("まで先手の勝ち", "endGame")
	caseExec("２六歩(27)", "move")
	caseExec("３四歩(33)", "move")
	caseExec("２五歩(26)", "move")
	caseExec("３三角(22)", "move")
	caseExec("５六歩(57)", "move")
	caseExec("３二金(41)", "move")
	caseExec("５五歩(56)", "move")
	caseExec("同　角(33)", "repeatMove")
}

func Test_convertKifCodeToEndGame(t *testing.T) {
	caseExec := func(movecode string, expectPrev, expectNext s.Position, expectPieceStates s.PieceStates, isResigned bool) {
		desc, err := convertKifCodeToMoveCode(movecode, 1)

		if err != nil {
			t.Error(err.Error())
		}

		_, ok := desc.(*m.EndGame)

		if !ok {
			t.Errorf("endGame型に変換できません。", movecode)
			return
		}
	}

	caseExec("投了", s.Position{0, 0}, s.Position{0, 0}, s.PieceStates{true, false, 0}, true)
	caseExec("中断", s.Position{0, 0}, s.Position{0, 0}, s.PieceStates{true, false, 0}, true)
	caseExec("反則", s.Position{0, 0}, s.Position{0, 0}, s.PieceStates{true, false, 0}, true)
	caseExec("引き分け", s.Position{0, 0}, s.Position{0, 0}, s.PieceStates{true, false, 0}, true)
	caseExec("まで先手の勝ち", s.Position{0, 0}, s.Position{0, 0}, s.PieceStates{true, false, 0}, true)
	caseExec("まで後手の勝ち", s.Position{0, 0}, s.Position{0, 0}, s.PieceStates{true, false, 0}, true)
	caseExec("13 投了", s.Position{0, 0}, s.Position{0, 0}, s.PieceStates{true, false, 0}, true)

}

func Test_convertKifCodeToRepeatMove(t *testing.T) {
	caseExec := func(kifCode string, expect s.PieceStates) {
		desc, err := convertKifCodeToMoveCode(kifCode, 1)

		if err != nil {
			t.Error(err.Error())
			return
		}

		rMove, ok := desc.(*m.RepeatMove)

		if !ok {
			t.Errorf("RepatMove型に変換できません。", kifCode)
			return
		}

		if rMove.PieceStates != expect {
			t.Errorf("RepeatMove.PieceStatesが期待値と異なります. expect = %#v, actual %#v", expect, rMove.PieceStates)
		}
	}

	caseExec("同銀(33)", s.PieceStates{def.BLACK, false, def.GIN})
	caseExec("同 銀(33)", s.PieceStates{def.BLACK, false, def.GIN})
	caseExec("同　銀(99)", s.PieceStates{def.BLACK, false, def.GIN})
	caseExec("同龍(11)", s.PieceStates{def.BLACK, true, def.HISHA})
	caseExec("同馬(33)", s.PieceStates{def.BLACK, true, def.KAKU})
	caseExec("同　　と(99)", s.PieceStates{def.BLACK, true, def.FU})
	caseExec("同成銀(58)", s.PieceStates{def.BLACK, true, def.GIN})

}

func Test_skipBlank(t *testing.T) {
	testCase := func(input string, expectSeq int) {
		rList := []rune(input)
		seq := 0
		skipBlank(rList, &seq)

		if expectSeq != seq {
			t.Errorf("[%s]seqの期待値が異なります。expect = %d, actual = %d", input, expectSeq, seq)
		}
	}
	testCase(" ２三銀", 1)
	testCase("　２三銀", 1)
	testCase("  ２三銀", 2)
	testCase(" 　２三銀", 2)
	testCase("   ２三銀", 3)
}

func Test_recognizePiece(t *testing.T) {
	testCase := func(input string, kop def.KindOfPiece, isP def.IsPromoted, expectSeq int) {
		rList := []rune(input)
		seq := 0
		actKop, actIsP := recognizePiece(rList, &seq, false)

		if actKop != kop {
			t.Errorf("[%s]kopの期待値が異なります。expect = %d, actual = %d", input, kop, actKop)
		}
		if actIsP != isP {
			t.Errorf("[%s]isPの期待値が異なります。expect = %d, actual = %d", input, isP, actIsP)
		}
		if seq != expectSeq {
			t.Errorf("[%s]seqの期待値が異なります。expect = %d, actual = %d", input, expectSeq, seq)
		}

	}

	testCase("銀", def.GIN, false, 1)
	testCase(" 銀", def.GIN, false, 2)
	testCase(" と", def.FU, true, 2)
	testCase(" さ", def.KindOfPiece(0), false, 1)
	testCase("銅", def.KindOfPiece(0), false, 0)

}

func Test_recognizePromotion(t *testing.T) {
	testCase := func(input string, currentP, isP def.IsPromoted, expectSeq int) {
		rList := []rune(input)
		seq := 0
		actIsP := recognizePromotion(rList, &seq, currentP)

		if actIsP != isP {
			t.Errorf("[%s]isPの期待値が異なります。expect = %d, actual = %d", input, isP, actIsP)
		}
		if seq != expectSeq {
			t.Errorf("[%s]seqの期待値が異なります。expect = %d, actual = %d", input, expectSeq, seq)
		}
	}

	testCase("成銀", false, true, 1)
	testCase(" 成銀", false, true, 2)
	testCase("　成銀", false, true, 2)
	testCase("銀成", false, false, 0)
}

func Test_recognizeNext(t *testing.T) {
	testCase := func(input string, expectPosition s.Position, expSeq int) {
		rList := []rune(input)
		seq := 0
		actualPosition, ok := recognizeNext(rList, &seq)

		if !ok {
			t.Errorf("[%s]Nextの検出に失敗しました。seq = %d", input, seq)
		}
		if actualPosition != expectPosition {
			t.Errorf("[%s]positionの期待値が異なります。expect = %s, actual = %s", input, expectPosition.Output(), actualPosition.Output())
		}

		if seq != expSeq {
			t.Errorf("[%s]seqの期待値が異なります。expect = %d, actual = %d", input, expSeq, seq)
		}
	}

	testCaseNotOk := func(input string, expSeq int) {
		rList := []rune(input)
		seq := 0
		_, ok := recognizeNext(rList, &seq)
		if ok {
			t.Errorf("[%s]を不正なＯＫフラグが検出されました。seq = %d", input, seq)
		}
		if expSeq != seq {
			t.Errorf("[%s]seqの期待値が異なります。expect = %d, actual = %d", input, expSeq, seq)
		}
	}

	testCase("２三成銀", s.Position{2, 3}, 2)
	testCase(" ５五飛車", s.Position{5, 5}, 3)
	testCase("９九成銀", s.Position{9, 9}, 2)
	testCaseNotOk("成銀", 0)
	testCaseNotOk(" 九９成銀", 1)

}

func Test_recognizePrev(t *testing.T) {
	testCase := func(input string, expectPosition s.Position, expSeq int) {
		rList := []rune(input)
		seq := 0
		actualPosition, ok := recognizePrev(rList, &seq)

		if !ok {
			t.Errorf("[%s]Prevの検出に失敗しました。seq = %d", input, seq)
		}
		if actualPosition != expectPosition {
			t.Errorf("[%s]positionの期待値が異なります。expect = %s, actual = %s", input, expectPosition.Output(), actualPosition.Output())
		}
		if seq != expSeq {
			t.Errorf("[%s]seqの期待値が異なります。expect = %d, actual = %d", input, expSeq, seq)
		}
	}

	testCaseNotOk := func(input string, expSeq int) {
		rList := []rune(input)
		seq := 0
		_, ok := recognizePrev(rList, &seq)
		if ok {
			t.Errorf("[%s]を不正なＯＫフラグが検出されました。seq = %d", input, seq)
		}
		if expSeq != seq {
			t.Errorf("[%s]seqの期待値が異なります。expect = %d, actual = %d", input, expSeq, seq)
		}
	}

	testCase("(99)", s.Position{9, 9}, 4)
	testCase("打", s.Position{0, 0}, 1)
	testCase(" (11)", s.Position{1, 1}, 5)
	testCaseNotOk("３二銀", 0)
	testCaseNotOk(" ３二銀", 1)
	testCaseNotOk("(３二銀", 1)
	testCaseNotOk("(32銀", 3)
	testCaseNotOk("(3)", 2)
}

func Test_skipSbNums(t *testing.T) {
	testCase := func(input string, expectSeq int) {
		rList := []rune(input)
		seq := 0
		skipSbNums(rList, &seq)

		if expectSeq != seq {
			t.Errorf("[%s]seqの期待値が異なります。expect = %d, actual = %d", input, expectSeq, seq)
		}
	}
	testCase("23 ", 2)
	testCase("130　２三銀", 3)
	testCase("5  ２三銀", 1)
	testCase(" 23 ", 3)

}
