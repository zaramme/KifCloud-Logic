package rsh

import (
	b "github.com/zaramme/KifCloud-Logic/board"
	"github.com/zaramme/KifCloud-Logic/code"
	def "github.com/zaramme/KifCloud-Logic/define"
	//	move "github.com/zaramme/KifCloud-Logic/move"
	s "github.com/zaramme/KifCloud-Logic/structs"
	//"fmt"
	"testing"
)

// 初期盤面RSH（11／9日時点）
const init_rsh = "Icflk2hGLCEL2UtPMR4e0ohgUgpT9-gsFBnWE0"

// 初期盤面RSH（１０／２６日時点）
//const init_rsh = "YHflk2hGLCEL2UtPMR4e0ohgUgpT9-gsFBnWE0"

var DoDebugOutput = false

func Test_BoardToString(t *testing.T) {

	// brd := b.NewBoardInit()
	// rsh := ConvertRshFromBoard(brd)

	// fmt.Printf("初期盤面RSH … %s\n", rsh.ToString())
	// fmt.Printf("---- tk … %s\n", rsh.Base_TK.ToString())
	// fmt.Printf("---- M2 … %s\n", rsh.Base_M2.ToString())
	// fmt.Printf("---- KIN … %s\n", rsh.Base_KIN.ToString())
	// fmt.Printf("---- GIN … %s\n", rsh.Base_GIN.ToString())
	// fmt.Printf("---- KEI … %s\n", rsh.Base_KEI.ToString())
	// fmt.Printf("---- KYO … %s\n", rsh.Base_KYO.ToString())
	// fmt.Printf("---- P18black … %s\n", rsh.Base_P18Black.ToString())
	// fmt.Printf("---- P18white … %s\n", rsh.Base_P18White.ToString())
	// fmt.Printf("---- P18cap … %s\n", rsh.Base_P18Cap.ToString())

}

func Test_ComposeAndDevideTKandAddP16exCap(t *testing.T) {

	testCase := func(tk, addP16 int) {
		tkCode := code.NewCode64FromInt(tk)
		addP16Code := code.NewCode64FromInt(addP16)

		composite := composeAddTkAndAddP16ExCap(tkCode, addP16Code, false)

		tkCodeDiveded, addP16CodeDevided, _ := divideADDTkAndAddP16ExCap(composite)

		if tkCode.ToInt() != tkCodeDiveded.ToInt() || addP16Code.ToInt() != addP16CodeDevided.ToInt() {
			t.Errorf("TK-AddP16の合成・分割処理に失敗しました。Imput = [%d,%d] Composite = [%d] Output = [%d,%d]",
				tk, addP16, composite.ToInt(), tkCodeDiveded.ToInt(), addP16CodeDevided.ToInt())
		}
	}

	// 全パターン実行(tk = 0 ~ 3, addP16 = 0 ~ 9)
	for tk := 0; tk < 4; tk++ {
		for addP16 := 0; addP16 < 6; addP16++ {
			testCase(tk, addP16)
		}
	}

}

func Test_NewRshCodeFromString_初期盤面(t *testing.T) {

	str := "Icflk2hGLCEL2UtPMR4e0ohgUgpT9-gsFBnWE0"
	// Ic : flk2h : GLCEL : 2UtPM : R4e0o : hgUgp : T9-gs : FBnWE : 02
	rsh, err := NewRshCodeFromString(str)

	if err != nil {
		t.Error("変換に失敗しました")
	}

	assert := func(expect string, actual code.Code64, id string) {
		if expect != actual.ToString() {
			t.Errorf("[%s]期待値と実測値が異なっています(,%s but %s)", id, expect, actual.ToString())
		}
	}

	assert("Ic", rsh.Base_TK, "tk")
	assert("GLCEL", rsh.Base_KIN, "KIN")
	assert("2UtPM", rsh.Base_GIN, "GIN")
	assert("R4e0o", rsh.Base_KEI, "KEI")
	assert("hgUgp", rsh.Base_KYO, "KYO")
	assert("T9-gs", rsh.Base_P18Black, "P18Black")
	assert("FBnWE", rsh.Base_P18White, "P18White")
	assert("0", rsh.Base_P18Cap, "P18Cap")

}

// func Test_NewRshCodeFromString_ToString_tk(t *testing.T) {
// 	brd := b.NewBoardInit()
// 	// m := move.NewMoveFromMoveCode("w91KY_91!")
// 	// brd.AddMove(m)
// 	// m = move.NewMoveFromMoveCode("w11KY_11!")
// 	//brd.AddMove(m)
// 	// m := move.NewMoveFromMoveCode("w56OH_51")
// 	// brd.AddMove(m)
// 	m := move.NewMoveFromMoveCode("w93FU_93!")
// 	brd.AddMove(m)
// 	m = move.NewMoveFromMoveCode("b97FU_97!")
// 	brd.AddMove(m)
// 	// m = move.NewMoveFromMoveCode("b83FU_83!")
// 	// brd.AddMove(m)
// 	// m = move.NewMoveFromMoveCode("w87FU_87!")
// 	// brd.AddMove(m)

// 	rsh := ConvertRshFromBoard(brd)

// 	ary := convertCode64To164Ary(rsh.Add_P18Prom)

// }

func Test_NewRshCodeFromString_ToString_初期盤面(t *testing.T) {

	testCase := func(str string) {
		rsh, err := NewRshCodeFromString(str)

		if err != nil {
			t.Error("変換に失敗しました")
		}

		if str != rsh.ToString() {
			t.Errorf("rshの文字列変換に失敗しました。\n----変換前=%s\n----変換後=%s, ", str, rsh.ToString())
		}
	}

	testCase("4sflk2hGLCEL2UtPMR4e0ohgUgpT0AEsFBnWE0")
	testCase("Icflk2hGLCEL2UtPMR4e0ohgUgpT9-gsFBnWE0|r104")
	testCase("5sflk2hGLCEL2UtPMR4e0ohgUgpT9-gsFBnWE0|6")      // tk_add
	testCase("4sflk2hGLCEL2UtPMR4e0ohgUgpT9-gsFBnWE0|l001")   // add_prom
	testCase("jcflk2hGLCEL2UtPMR4e0ohgUgpT9-gsP8nWE0|6N1")    // tk_add + add_p16prom
	testCase("Icflk2hGLCEL2UtPMR4e0ohgUgpT9-gsF-75e0|l001e2") // add_prom + add16prom
	testCase("4sflk2hGLCEL2UtPMR4e0ohgUgpV__gsP8nWE0|r1045-")
	testCase("4sflk2hGLCEL2UtPMR4e0ohgUgpTRrp4F1a-10|0Ea3") // 16prom
	testCase("Icflk2hGLCEL2UtPMR4e0ohgUgpT9-gsF1a-10|bI2")  // tk + 16prom
}

//asflk2hGLCEL2UtPMR4e0ohgUgpLRzgsXtPhl0|o0

func assert_EachPositionOfBoard(expected, actual *b.Board, id string, t *testing.T) {

	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			// result := true

			pos := s.Position{x, y}
			if psExpected, ok := expected.PositionMap[pos]; ok {
				if psActual, ok2 := actual.PositionMap[pos]; ok2 {
					if psExpected != psActual {
						t.Errorf("[id = %s]期待値と実測値が異なります。座標＝[%d,%d]\n", id, x, y)
						t.Errorf("--------期待値＝", expected.PositionMap[pos])
						t.Errorf("--------実測値＝", actual.PositionMap[pos])
					}
				} else {
					t.Errorf("[id = %s]期待値と実測値が異なります。座標＝[%d,%d]に実際は存在しない", id, x, y)
				}
			} else if _, ok3 := actual.PositionMap[pos]; ok3 {
				t.Errorf("[id = %s]期待値と実測値が異なります。", id)
				t.Errorf("--------座標＝[%d,%d]に期待外の駒%!", x, y, actual.PositionMap[pos])
			}

			// if result == false {
			// 	t.Errorf("[id = %s]期待値と実測値が異なります。座標＝[%d,%d]", id, x, y)
			// }
		}
	}

	kList := []def.KindOfPiece{def.OH, def.HISHA, def.KAKU, def.KIN, def.GIN, def.KEI, def.KYO, def.FU}
	pList := []def.Player{def.BLACK, def.WHITE}

	for _, kop := range kList {
		for _, player := range pList {
			captured := s.CapArea{player, kop}
			expCount, expOk := expected.CapturedMap[captured]
			actCount, actOk := expected.CapturedMap[captured]

			switch {
			case expOk && actOk:
				if expCount != actCount {
					t.Errorf("[id = %s]期待値と実測値が異なります。座標＝[%s,%s] 期待値＝%d, 実測値=%d", id, player, kop, expCount, actCount)
				} else {
					// fmt.Printf("➡OK\n")
				}
				continue
			case expOk && !actOk:
				t.Errorf("[id = %s]期待値と実測値が異なります。座標＝[%s,%s] 期待値＝%d, 実測値=0", id, player, kop, expCount)

			case !expOk && actOk:
				t.Errorf("[id = %s]期待値と実測値が異なります。座標＝[%s,%s] 期待値＝0, 実測値=%d", id, player, kop, actCount)
			case !expOk && !actOk:
				continue
			}
		}
	}
}
