package move

import (
	def "KifuLibrary-Logic/define"
	s "KifuLibrary-Logic/structs"
	"testing"
)

func Test_move_construct(t *testing.T) {
	prev := s.Position{2, 7}
	next := s.Position{2, 6}
	var move Move = Move{
		Prev: prev,
		Next: next,
		PieceStates: s.PieceStates{
			IsPromoted:  false,
			KindOfPiece: def.FU,
			Player:      def.BLACK,
		},
	}

	if move.Prev.X != 2 || move.Prev.Y != 7 {
		t.Errorf("Moveの初期化に失敗しました(prev)")
	}
	if move.Next.X != 2 || move.Next.Y != 6 {
		t.Errorf("Moveの初期化に失敗しました(next)")
	}
	if move.Player != def.BLACK {
		t.Errorf("moveの初期化に失敗しました")
	}
	if move.IsPromoted != false {
		t.Errorf("moveの初期化に失敗しました")
	}
	if move.KindOfPiece != def.FU {
		t.Errorf("moveの初期化に失敗しました")
	}
}

func Test_NewMoveFromMoveCode_標準(t *testing.T) {

	actual := NewMoveFromMoveCode("w76FU_77")
	expect := &Move{
		Prev: s.Position{7, 7},
		Next: s.Position{7, 6},
		PieceStates: s.PieceStates{
			KindOfPiece: def.FU,
			Player:      def.WHITE,
			IsPromoted:  false,
		},
		IsResigned: false,
	}
	assert_compareMoves(actual, expect, t)
}

func Test_NewMoveFromMoveCode_成り(t *testing.T) {

	actual := NewMoveFromMoveCode("b52KA_88!")
	expect := &Move{
		Prev: s.Position{8, 8},
		Next: s.Position{5, 2},
		PieceStates: s.PieceStates{
			KindOfPiece: def.KAKU,
			Player:      def.BLACK,
			IsPromoted:  true,
		},
		IsResigned: false,
	}
	assert_compareMoves(actual, expect, t)
}

func Test_NewMoveFromMoveCode_打つ(t *testing.T) {

	actual := NewMoveFromMoveCode("b55KA_00")
	expect := &Move{
		Prev: s.Position{0, 0},
		Next: s.Position{5, 5},
		PieceStates: s.PieceStates{
			KindOfPiece: def.KAKU,
			Player:      def.BLACK,
			IsPromoted:  false,
		},
		IsResigned: false,
	}
	assert_compareMoves(actual, expect, t)
}

func Test_NewMoveFromMoveCode_投了(t *testing.T) {
	actual := NewMoveFromMoveCode("b_resigned")
	expect := &Move{
		Prev: s.Position{0, 0},
		Next: s.Position{0, 0},
		PieceStates: s.PieceStates{
			KindOfPiece: 0,
			Player:      def.BLACK,
			IsPromoted:  false,
		},
		IsResigned: true,
	}

	assert_compareMoves(actual, expect, t)

	actual = NewMoveFromMoveCode("w_resigned")
	expect = &Move{
		Prev: s.Position{0, 0},
		Next: s.Position{0, 0},
		PieceStates: s.PieceStates{
			KindOfPiece: 0,
			Player:      def.WHITE,
			IsPromoted:  false,
		},
		IsResigned: true,
	}

	assert_compareMoves(actual, expect, t)

}

func Test_convertMoveCodeToPlayer(t *testing.T) {

	asrt := func(expected def.Player, isValid bool, input string) {
		actual, err := convertMoveCodeToPlayer([]byte(input))
		if err != nil && isValid {
			t.Errorf("Playerの変換に失敗しました。(変換エラー…%s)", err)
			return
		}
		if err == nil && !isValid {
			t.Errorf("エラー検出に失敗しました。(入力値…%s)", input)
			return
		}
		if !isValid {
			return
		}
		if expected != actual {
			t.Errorf("Playerの変換に失敗しました。期待値＝%b、実測値%b", expected, actual)
		}
	}

	asrt(true, true, "b11KA_22+")
	asrt(false, true, "w95HI(35+")
	asrt(false, true, "w11KA(22")

	asrt(false, false, "k11KA(22")
	asrt(true, false, "12KA(22+")
}

func Test_convertMoveCodeIsResined(t *testing.T) {

	// test pattert
	asrt := func(expected bool, input string) {
		actual := convertMoveCodeIsResigned([]byte(input))
		if actual != expected {
			t.Errorf("isResignedの検出に失敗しました。input = %s", input)
		}
	}

	// test cases

	asrt(true, "b_resigned")
	asrt(true, "w_resigned")

	asrt(false, "b11KA_22+")
	asrt(false, "w95HI(35+")
	asrt(false, "w11KA(22")

	asrt(false, "k11KA(22")
	asrt(false, "12KA(22+")

}

func Test_convertMoveCodeToMoveNext(t *testing.T) {
	assert_convertMoveCodeToMoveNext(11, "b11KA(23)", t)
	assert_convertMoveCodeToMoveNext(55, "w55OH(23)", t)
	assert_convertMoveCodeToMoveNext(26, "b26ka(23)", t)
	assert_convertMoveCodeToMoveNext(99, "w99ka(23)", t)
	assert_convertMoveCodeToMoveNext(48, "b480ka(23)", t)

	assert_convertMoveCodeToMoveNext(-1, "w00ka(23)", t)
	assert_convertMoveCodeToMoveNext(-1, "bAAka(23)", t)
}

func Test_convertMoveCodeToKindOfPiece(t *testing.T) {

	assert_convertMoveCodeToKindOfPiece(def.OH, "w12OH+(34)", t)
	assert_convertMoveCodeToKindOfPiece(def.HISHA, "b12HI(34)", t)
	assert_convertMoveCodeToKindOfPiece(def.KAKU, "w12KA+(34)", t)
	assert_convertMoveCodeToKindOfPiece(def.KIN, "b12KI(34)", t)
	assert_convertMoveCodeToKindOfPiece(def.GIN, "w12GI+(34)", t)
	assert_convertMoveCodeToKindOfPiece(def.KEI, "b12KE(34)", t)
	assert_convertMoveCodeToKindOfPiece(def.KYO, "w12KY+(34)", t)
	assert_convertMoveCodeToKindOfPiece(def.FU, "b12FU(34)", t)

	assert_convertMoveCodeToKindOfPiece(-1, "w12KO+(34)", t)
	assert_convertMoveCodeToKindOfPiece(-1, "b12K12+(34)", t)
	assert_convertMoveCodeToKindOfPiece(-1, "w122K12+(34)", t)
}

func Test_convertMoveCodeToMovePrev(t *testing.T) {
	assert_convertMoveCodeToMovePrev(22, "b11KA(22)", t)
	assert_convertMoveCodeToMovePrev(44, "b55OH(44)", t)
	assert_convertMoveCodeToMovePrev(66, "b26ka(66)", t)
	assert_convertMoveCodeToMovePrev(88, "w99ka(88)", t)
	assert_convertMoveCodeToMovePrev(0, "w11ka(00)", t)

	assert_convertMoveCodeToMovePrev(-1, "w480ka(23)", t)
	assert_convertMoveCodeToMovePrev(-1, "w11ka(AA)", t)
}

func Test_convertMoveCodeToIsPromoted(t *testing.T) {
	assert_convertMoveCodeToIsPromoted(true, true, "w11KA_22!", t)
	assert_convertMoveCodeToIsPromoted(true, true, "w95HI_35!", t)
	assert_convertMoveCodeToIsPromoted(false, true, "b11KA_22", t)

	assert_convertMoveCodeToIsPromoted(true, false, "b11KA_22*", t)
	assert_convertMoveCodeToIsPromoted(true, false, "b122KA_22!", t)
}

func Test_ToMoveCode(t *testing.T) {
	asrt := func(input string) {
		move := NewMoveFromMoveCode(input)
		act := move.ToMoveCode()
		if input != act {
			t.Errorf("MoveCodeへのの変換に失敗しました。期待値＝%s、実測値 = %s", input, act)
		}
	}

	asrt("b76FU_77")
	asrt("b76FU_77!")
	asrt("w99KA_22")
	asrt("w99KA_22!")
	asrt("w22HI_82")
	asrt("w22HI_82!")
	asrt("b_resigned")
	asrt("w_resigned")

}

func Test_ToJpnCode(t *testing.T) {
	asrt := func(input string, exp string) {
		move := NewMoveFromMoveCode(input)
		act := move.ToJpnCode()

		if exp != act {
			t.Errorf("JPNCodeへのの変換に失敗しました。期待値＝%s、実測値 = %s", exp, act)
		}
	}
	asrt("b76FU_77", "▲７六歩(77)")
	asrt("b76FU_77!", "▲７六と(77)")
	asrt("w99KA_22", "▽９九角(22)")
	asrt("w99KA_22!", "▽９九馬(22)")
	asrt("w22HI_82", "▽２二飛(82)")
	asrt("w22HI_82!", "▽２二龍(82)")
	asrt("b_resigned", "投了")
	asrt("w_resigned!", "投了")

}

func Test_ToJsCode(t *testing.T) {
	asrt := func(moveCode string, exp string) {
		move := NewMoveFromMoveCode(moveCode)
		act := move.ToJsCode()
		if exp != act {
			t.Errorf("JSCodeへのの変換に失敗しました。期待値＝%d、実測値%d", exp, act)
		}
	}
	asrt("w22KA_11!", "11,22,KAKU,true")
	asrt("w86HI_88", "88,86,HISHA,false")
	asrt("b55FU_54", "54,55,FU,false")
	asrt("b_resigned", "")
	asrt("w_resigned", "")

}

func assert_convertMoveCodeToKindOfPiece(expected def.KindOfPiece, input string, t *testing.T) {
	actual, err := convertMoveCodeToKindOfPiece([]byte(input))

	if err != nil && expected != -1 {
		t.Errorf("エラーを検出しました。➡(%s)", err)
		return
	}

	if expected != actual {
		t.Errorf("駒種類の変換に失敗しました。期待値＝%d、実測値%d", expected, actual)
	}
}

func assert_convertMoveCodeToMoveNext(expectedInt int, input string, t *testing.T) {

	expected := s.Position{expectedInt / 10, expectedInt % 10}
	var actual s.Position
	var err error

	actual, err = convertMoveCodeToMoveNext([]byte(input))

	if err != nil && expectedInt != -1 {
		t.Errorf("nextの変換に失敗しました。(変換エラー…%s)", err)
		return
	}

	if expectedInt == -1 {
		return // エラー検出の場合はこれ以降の処理を行わない
	}

	if expected != actual {
		t.Errorf("nextの変換に失敗しました。期待値＝%d、実測値%d", expected, actual)
	}
}

func assert_convertMoveCodeToMovePrev(expectedInt int, input string, t *testing.T) {

	expected := s.Position{expectedInt / 10, expectedInt % 10}
	var actual s.Position
	var err error

	actual, err = convertMoveCodeToMovePrev([]byte(input))

	if err != nil && expectedInt != -1 {
		t.Errorf("Prevの変換に失敗しました。(変換エラー…%s)", err)
		return
	}

	if expectedInt == -1 {
		return // エラー検出の場合はこれ以降の処理を行わない
	}

	if expected != actual {
		t.Errorf("Prevの変換に失敗しました。期待値＝%d、実測値%d", expected, actual)
	}
}

func assert_convertMoveCodeToIsPromoted(expected bool, isValid bool, input string, t *testing.T) {

	actual, err := convertMoveCodeToIsPromoted([]byte(input))

	if err != nil && isValid {
		t.Errorf("isPromotedの変換に失敗しました。(変換エラー…%s)", err)
		return
	}

	if err == nil && !isValid {
		t.Errorf("エラー検出に失敗しました。(入力値…%s)", input)
		return
	}

	if !isValid {
		return
	}

	if expected != actual {
		t.Errorf("[code = %s]IsPromotedの変換に失敗しました。期待値＝%b、実測値%b", input, expected, actual)
	}

}

func assert_compareMoves(m1 *Move, m2 *Move, t *testing.T) {
	if m1.Prev != m2.Prev {
		t.Errorf("Moveの値が期待値と異なっています(Prev)")
	}
	if m1.Next != m2.Next {
		t.Errorf("Moveの値が期待値と異なっています(Next)")
	}
	if m1.IsPromoted != m2.IsPromoted {
		t.Errorf("Moveの値が期待値と異なっています(isPromoted)")
	}
	if m1.KindOfPiece != m2.KindOfPiece {
		t.Errorf("Moveの値が期待値と異なっています(KindOfPiece)")
	}
	if m1.Player != m2.Player {
		t.Errorf("Moveの値が期待値と異なっています(Player)")
	}
}
