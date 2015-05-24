package kifLoader

import (
	def "KifuLibrary-Logic/define"
	m "KifuLibrary-Logic/move"
	s "KifuLibrary-Logic/structs"
	"fmt"
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
	// caseExec("９四歩(93)", s.Potirion{9, 4}, s.Position{9, 3}, s.PieceStates{})
	// caseExec("６八銀(79)", s.Potirion{6, 8}, s.Position{7, 9}, s.PieceStates{})
	// caseExec("３二銀(31)", s.Potirion{3, 2}, s.Position{3, 1}, s.PieceStates{})
	// caseExec("２七銀(38)", s.Potirion{2, 7}, s.Position{3, 8}, s.PieceStates{})
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
}
