package rsh

import (
	"fmt"
	b "github.com/zaramme/KifCloud-Logic/board"
	"github.com/zaramme/KifCloud-Logic/code"
	def "github.com/zaramme/KifCloud-Logic/define"
	"github.com/zaramme/KifCloud-Logic/math"
	mv "github.com/zaramme/KifCloud-Logic/move"
	s "github.com/zaramme/KifCloud-Logic/structs"
	plmath "math"
	"testing"
)

func Test_ConvertPiecePositionTo164(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	var actual int
	var expected int

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			switch {
			case (j == 1 || j == 3): // １段目、３段目はすべて後手駒
				expected = (i-1)*9 + (j - 1) + 81
			case (i == 2 && j == 2 || i == 8 && j == 2):
				expected = (i-1)*9 + (j - 1) + 81
			case (i == 2 && j == 8 || i == 8 && j == 8):
				expected = (i-1)*9 + (j - 1)
			case (j == 7 || j == 9): // 七段目、九段目はすべて先手駒
				expected = (i-1)*9 + (j - 1)
			default:
				expected = -1

			}
			actual = convertPiecePositionTo164(s.Position{i, j}, rsh.Board)
			if actual != expected {
				t.Errorf("値が期待値と異なっています", i, j, actual)
			}
		}
	}
}

func Test_getM2fromBoard_初期盤面の復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()

	m2 := getM2fromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(m2, "M2")

	assert_N164EqualPosition(n164ary[0], s.Position{2, 8}, t)
	assert_N164EqualPosition(n164ary[1], s.Position{8, 2}, t)
	assert_N164EqualPosition(n164ary[2], s.Position{8, 8}, t)
	assert_N164EqualPosition(n164ary[3], s.Position{2, 2}, t)
}

func Test_getM2fromBoard_持ち駒込みの復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	delete(rsh.Board.PositionMap, s.Position{2, 2})
	delete(rsh.Board.PositionMap, s.Position{2, 8})
	rsh.Board.CapturedMap[s.CapArea{def.BLACK, def.HISHA}] = 1
	rsh.Board.CapturedMap[s.CapArea{def.WHITE, def.KAKU}] = 1

	m2 := getM2fromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(m2, "M2")

	assert_N164EqualPosition(n164ary[0], s.Position{8, 2}, t)
	assert_IsN164Captured(n164ary[1], true, t)
	assert_N164EqualPosition(n164ary[2], s.Position{8, 8}, t)
	assert_IsN164Captured(n164ary[3], false, t)
}

func Test_getKINfromBoard_初期盤面の復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()

	m2 := getKINfromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(m2, "M4(KIN)")

	assert_N164EqualPosition(n164ary[0], s.Position{4, 9}, t)
	assert_N164EqualPosition(n164ary[1], s.Position{6, 9}, t)
	assert_N164EqualPosition(n164ary[2], s.Position{4, 1}, t)
	assert_N164EqualPosition(n164ary[3], s.Position{6, 1}, t)
}

func Test_getKINfromBoard_持ち駒込みの復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	delete(rsh.Board.PositionMap, s.Position{4, 1})
	delete(rsh.Board.PositionMap, s.Position{6, 1})
	delete(rsh.Board.PositionMap, s.Position{4, 9})
	rsh.Board.CapturedMap[s.CapArea{def.BLACK, def.KIN}] = 2
	rsh.Board.CapturedMap[s.CapArea{def.WHITE, def.KIN}] = 1

	base_kin := getKINfromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(base_kin, "M4(KIN)")

	assert_N164EqualPosition(n164ary[0], s.Position{6, 9}, t)
	assert_IsN164Captured(n164ary[1], def.BLACK, t)
	assert_IsN164Captured(n164ary[2], def.BLACK, t)
	assert_IsN164Captured(n164ary[3], def.WHITE, t)
}

func Test_getGINfromBoard_初期盤面の復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()

	base_gin := getGINfromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(base_gin, "M4(GIN)")

	assert_N164EqualPosition(n164ary[0], s.Position{3, 9}, t)
	assert_N164EqualPosition(n164ary[1], s.Position{7, 9}, t)
	assert_N164EqualPosition(n164ary[2], s.Position{3, 1}, t)
	assert_N164EqualPosition(n164ary[3], s.Position{7, 1}, t)
}

func Test_getGINfromBoard_持ち駒込みの復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	delete(rsh.Board.PositionMap, s.Position{7, 1})
	delete(rsh.Board.PositionMap, s.Position{7, 9})
	delete(rsh.Board.PositionMap, s.Position{3, 9})
	rsh.Board.CapturedMap[s.CapArea{def.WHITE, def.GIN}] = 3

	code := getGINfromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(code, "M4(GIN)")

	assert_N164EqualPosition(n164ary[0], s.Position{3, 1}, t)
	assert_IsN164Captured(n164ary[1], def.WHITE, t)
	assert_IsN164Captured(n164ary[2], def.WHITE, t)
	assert_IsN164Captured(n164ary[3], def.WHITE, t)
}

func Test_getKEIfromBoard_初期盤面の復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()

	code := getKEIfromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(code, "M4(KEI)")

	assert_N164EqualPosition(n164ary[0], s.Position{2, 9}, t)
	assert_N164EqualPosition(n164ary[1], s.Position{8, 9}, t)
	assert_N164EqualPosition(n164ary[2], s.Position{2, 1}, t)
	assert_N164EqualPosition(n164ary[3], s.Position{8, 1}, t)
}

func Test_getKEIfromBoard_持ち駒込みの復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	delete(rsh.Board.PositionMap, s.Position{2, 1})
	delete(rsh.Board.PositionMap, s.Position{8, 1})
	delete(rsh.Board.PositionMap, s.Position{2, 9})
	delete(rsh.Board.PositionMap, s.Position{8, 9})
	rsh.Board.CapturedMap[s.CapArea{def.BLACK, def.KEI}] = 2
	rsh.Board.CapturedMap[s.CapArea{def.WHITE, def.KEI}] = 2

	code := getKEIfromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(code, "M4(KEI)")

	assert_IsN164Captured(n164ary[0], def.BLACK, t)
	assert_IsN164Captured(n164ary[1], def.BLACK, t)
	assert_IsN164Captured(n164ary[2], def.WHITE, t)
	assert_IsN164Captured(n164ary[3], def.WHITE, t)
}

func Test_getKYOfromBoard_初期盤面の復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()

	code := getKYOfromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(code, "M4(KYO)")

	assert_N164EqualPosition(n164ary[0], s.Position{1, 9}, t)
	assert_N164EqualPosition(n164ary[1], s.Position{9, 9}, t)
	assert_N164EqualPosition(n164ary[2], s.Position{1, 1}, t)
	assert_N164EqualPosition(n164ary[3], s.Position{9, 1}, t)
}

func Test_getKYOfromBoard_持ち駒込みの復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	move := mv.NewMoveFromMoveCode("b91HI_28")
	rsh.Board.AddMove(move)

	code := getKYOfromBoard(rsh.Board)

	n164ary := convert64To164AndOutput64Stats(code, "M4(GIN)")

	assert_N164EqualPosition(n164ary[0], s.Position{1, 9}, t)
	assert_N164EqualPosition(n164ary[1], s.Position{9, 9}, t)
	assert_N164EqualPosition(n164ary[2], s.Position{1, 1}, t)
	assert_IsN164Captured(n164ary[3], def.BLACK, t)
}

func Test_getP18fromBoard_初期盤面の復元照合(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	pBlack, pWhite, pCap, _, _ := getP18fromBoard(rsh.Board)

	assert_IsP18Position(pBlack, 777777777, t)
	assert_IsP18Position(pWhite, 333333333, t)
	assert_IsP18Captured(pCap, []int{0, 0}, t)
}

// func Test_getP18fromBoard_ケース１(t *testing.T) {
// 	rsh := new(RshCode)
// 	rsh.Board = b.NewBoardInit()
// 	delete(rsh.Board.PositionMap, s.Position{1, 7})
// 	delete(rsh.Board.PositionMap, s.Position{3, 7})
// 	delete(rsh.Board.PositionMap, s.Position{5, 7})
// 	delete(rsh.Board.PositionMap, s.Position{7, 7})
// 	delete(rsh.Board.PositionMap, s.Position{9, 7})
// 	rsh.Board.CapturedMap[s.CapArea{def.BLACK, def.FU}] = 3
// 	rsh.Board.CapturedMap[s.CapArea{def.WHITE, def.FU}] = 2
// 	pBlack, pWhite, pCap, _, _ := getP18fromBoard(rsh.Board)

// 	assert_IsP18Position(pBlack, 70707070, t)
// 	assert_IsP18Position(pWhite, 333333333, t)
// 	assert_IsP18Captured(pCap, []int{3, 2}, t)
// }

// func Test_getP18fromBoard_持ち駒増減(t *testing.T) {
// 	rsh := new(RshCode)
// 	rsh.Board = b.NewBoardInit()
// 	for i := 1; i <= 9; i++ {
// 		delete(rsh.Board.PositionMap, s.Position{i, 3})
// 		delete(rsh.Board.PositionMap, s.Position{i, 7})
// 	}

// 	testExec := func(black, white int, capExpected, addExpected []int) {
// 		rsh.Board.CapturedMap[s.CapArea{def.BLACK, def.FU}] = black
// 		rsh.Board.CapturedMap[s.CapArea{def.WHITE, def.FU}] = white

// 		_, _, pCap, pAdd, _ := getP18fromBoard(rsh.Board)

// 		assert_IsP18Captured(pCap, capExpected, t)
// 		assert_IsP18Add_Captured(pAdd, addExpected, t)
// 	}

// 	testExec(18, 0, []int{2, 0}, []int{2, 0})
// 	testExec(17, 1, []int{1, 1}, []int{2, 0})
// 	testExec(16, 2, []int{0, 2}, []int{2, 0})
// 	testExec(15, 3, []int{7, 3}, []int{1, 0})
// 	testExec(14, 4, []int{6, 4}, []int{1, 0})
// 	testExec(13, 5, []int{5, 5}, []int{1, 0})
// 	testExec(12, 6, []int{4, 6}, []int{1, 0})
// 	testExec(11, 7, []int{3, 7}, []int{1, 0})
// 	testExec(10, 8, []int{2, 0}, []int{1, 1})
// 	testExec(9, 9, []int{1, 1}, []int{1, 1})
// 	testExec(8, 10, []int{0, 2}, []int{1, 1})
// 	testExec(7, 11, []int{7, 3}, []int{0, 1})
// 	testExec(6, 12, []int{6, 4}, []int{0, 1})
// 	testExec(5, 13, []int{5, 5}, []int{0, 1})
// 	testExec(4, 14, []int{4, 6}, []int{0, 1})
// 	testExec(3, 15, []int{3, 7}, []int{0, 1})
// 	testExec(2, 16, []int{2, 0}, []int{0, 2})
// 	testExec(1, 17, []int{1, 1}, []int{0, 2})
// 	testExec(0, 18, []int{0, 2}, []int{0, 2})

// }

// func Test_getP18fromBoard_持ち駒なし１(t *testing.T) {
// 	rsh := new(RshCode)
// 	rsh.Board = b.NewBoardInit()

// 	var move *mv.Move
// 	move = mv.NewMoveFromMoveCode("b26FU(27)")
// 	rsh.Board.AddMove(move)

// 	var pBlack, pWhite, pCap code.Code64

// 	pBlack, pWhite, pCap, _, _ = getP18fromBoard(rsh.Board)

// 	assert_IsP18Position(pBlack, 777777767, t)
// 	assert_IsP18Position(pWhite, 333333333, t)
// 	assert_IsP18Captured(pCap, []int{0, 0}, t)

// 	move = mv.NewMoveFromMoveCode("w34FU(33)")
// 	rsh.Board.AddMove(move)
// 	pBlack, pWhite, pCap, _, _ = getP18fromBoard(rsh.Board)

// 	assert_IsP18Position(pBlack, 777777767, t)
// 	assert_IsP18Position(pWhite, 333333433, t)
// 	assert_IsP18Captured(pCap, []int{0, 0}, t)

// }

func Test_getProPsfromBoard_最小値(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	move := mv.NewMoveFromMoveCode("b11HI_28!")
	rsh.Board.AddMove(move)

	actual := getPromfromBoard(rsh.Board)
	var expected code.Code64

	expected = code.NewCode64FromInt(1)
	expected = expected.Padding(3)
	assert_Code64Equals(expected, actual, "props_1", t)
}

func Test_getProPsfromBoard_最大値(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	move := mv.NewMoveFromMoveCode("b98KY_99")
	move2 := mv.NewMoveFromMoveCode("w99HI_28!")

	if move == nil || move2 == nil {
		t.Error("moveコードの変換に失敗しました。")
		return
	}
	rsh.Board.AddMove(move)
	rsh.Board.AddMove(move2)

	actual := getPromfromBoard(rsh.Board)
	var expected code.Code64

	expected = code.NewCode64FromInt(int(plmath.Pow(2, 15)))
	expected = expected.Padding(3)
	assert_Code64Equals(expected, actual, "props_1", t)
}

func Test_getProPsfromBoard_最小から最大(t *testing.T) {
	rsh := new(RshCode)
	rsh.Board = b.NewBoardInit()
	rsh.Board.Turn = def.WHITE

	testUnit := func(moveCode string, expected []int) {
		move := mv.NewMoveFromMoveCode(moveCode)
		rsh.Board.AddMove(move)
		actual := getPromfromBoard(rsh.Board)

		cur := 1
		expectedDigit := 0
		for i := 0; i < len(expected); i++ {
			expectedDigit += expected[i] * cur
			cur *= 2
		}
		//		fmt.Printf("現在のexpectedDigit = %d\n", expectedDigit)
		expectedCode64 := code.NewCode64FromInt(expectedDigit)
		expectedCode64 = expectedCode64.Padding(3)

		assert_Code64Equals(expectedCode64, actual, moveCode, t)
	}

	testUnit("w11KY_11!", []int{1})
	testUnit("b19KY_19!", []int{1, 1})
	testUnit("w21KE_21!", []int{1, 1, 1})
	testUnit("b22KA_22!", []int{1, 1, 1, 1})
	testUnit("w28HI_28!", []int{1, 1, 1, 1, 1})
	testUnit("b29KE_29!", []int{1, 1, 1, 1, 1, 1})
	testUnit("b31GI_31!", []int{1, 1, 1, 1, 1, 1, 1})
	testUnit("b39GI_39!", []int{1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("271GI_71!", []int{1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b79GI_79!", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("w81KE_81!", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b82HI_82!", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b88KA_88!", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b89KE_89!", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b91KY_91!", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b99KY_99!", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})

	testUnit("b19KY_19", []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b22KA_22", []int{1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b29KE_29", []int{1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b39GI_39", []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1})
	testUnit("b79GI_79", []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1})
	testUnit("b82HI_82", []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1})
	testUnit("b89KE_89", []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1})
	testUnit("b99KY_99", []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0})

}

func convert64To164AndOutput64Stats(code code.Code64, name string) math.N164ary {

	digit := code.ToInt()
	n164Ary := math.Convert164Ary(digit)

	if DoDebugOutput {
		fmt.Printf("%sを解析しています…。\n", name)
		fmt.Printf("---- %s = %s\n", name, code.ToString())
		fmt.Printf("---- %s(１０進数）＝%d\n", name, digit)
		fmt.Printf("---- %s[164] = ", name)
		fmt.Print(n164Ary)
		fmt.Print("\n")
	}
	return n164Ary
}

func convert64ToTKValueAndOutputStats(code code.Code64, name string) math.N164ary {

	digit := code.ToInt()
	n45Ary := math.GetNary(digit, 45)

	if DoDebugOutput {
		fmt.Printf("%sを解析しています…。\n", name)
		fmt.Printf("---- %s = %s\n", name, code.ToString())
		fmt.Printf("---- %s(１０進数）＝%d\n", name, digit)
		fmt.Printf("---- %s[45] = ", name)
		fmt.Print(n45Ary)
		fmt.Print("\n")
	}
	return n45Ary
}

func assert_N164EqualPosition(n164 int, expPos s.Position, t *testing.T) {
	if n164 > 80 {
		n164 -= 81
	}
	actPos := s.Position{(n164 / 9) + 1, (n164 % 9) + 1}
	if actPos != expPos {
		t.Errorf("値が期待値と異なっています。 期待値{%d,%d}➡実測値{%d,%d}", expPos.X, expPos.Y, actPos.X, actPos.Y)
	}
}

func assert_N45EqualPosition(n164 int, add int, isBlack bool, expPos s.Position, t *testing.T) {
	if isBlack && add%2 != 1 {
		n164 += 5
	}

	if !isBlack && add/2 == 1 {
		n164 += 5
	}

	actPos := s.Position{(n164 / 9) + 1, (n164 % 9) + 1}
	if actPos != expPos {
		t.Errorf("値が期待値と異なっています。 期待値{%d,%d}➡実測値{%d,%d}", expPos.X, expPos.Y, actPos.X, actPos.Y)
	}
}

func assert_IsN164Captured(n164 int, player def.Player, t *testing.T) {
	if n164 < 162 {
		t.Errorf("値が持ち駒ではありません。実測値{%d}", n164)
	}
	if n164 == 162 && player != def.BLACK {
		t.Errorf("値が持ち駒ではありません。期待値＝163、 実測値{%d}", n164)
	}
	if n164 == 163 && player != def.WHITE {
		t.Errorf("値が持ち駒ではありません。期待値＝162、 実測値{%d}", n164)
	}
}

func assert_IsP18Position(p18 code.Code64, expected int, t *testing.T) {
	actual := p18.ToInt()

	if actual != expected {
		t.Errorf("値が期待値と異なっています。期待値＝%d、 実測値=%d", expected, actual)
	}
}

func assert_IsP18Captured(p18cap code.Code64, expected []int, t *testing.T) {
	actual := make([]int, 2)
	digit := p18cap.ToInt()

	for i := 0; i < len(actual); i++ {
		actual[i] = digit % 8
		digit = digit / 8
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("p18の持ち駒の値が異なります。期待/実測　→", expected, actual)
			return
		}
	}
}

func assert_IsP18Add_Captured(p18AddCap code.Code64, expected []int, t *testing.T) {
	actual := make([]int, 2)
	digit := p18AddCap.ToInt()

	for i := 0; i < len(actual); i++ {
		actual[i] = digit % 3
		digit = digit / 3
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Errorf("p18の持ち駒シフトの値が異なります。期待/実測　→", expected, actual)
			return
		}
	}
}

func assert_Code64Equals(expected, actual code.Code64, name string, t *testing.T) {
	if DoDebugOutput {
		fmt.Println("propsを解析しています///")
		fmt.Printf(" ----props = %s\n", actual.ToString())
	}
	if len(expected) != len(actual) {
		t.Errorf("----%sが期待値と異なっています期待値(len)=%d, 実測値(len)=%d", name, len(expected), len(actual))
		return
	}

	var assertOK = true
	for i := 0; i < len(expected); i++ {
		if actual[i].Num != expected[i].Num {
			assertOK = false
		}
	}

	if !assertOK {
		t.Errorf("----%sが期待値と異なっています期待値=%s, 実測値=%s", name, expected.ToString(), actual.ToString())
	}
}
