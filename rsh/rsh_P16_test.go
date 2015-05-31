package rsh

import (
	b "github.com/zaramme/KifCloud-Logic/board"
	def "github.com/zaramme/KifCloud-Logic/define"
	//	mv "../move"
	"fmt"
	s "github.com/zaramme/KifCloud-Logic/structs"
	"strconv"
	"testing"
)

func __rsh_p16test() {
	fmt.Print()
}

func Test_getAndPutP16Black_先手駒10000パターン(t *testing.T) {
	// テストパターン
	strToPlist := func(num int) []s.Position {
		pList := make([]s.Position, 0)
		for i := 1; i <= 9; i++ {
			if num < 1 {
				break
			}
			y := num % 10
			pos := s.Position{i, y}
			pList = append(pList, pos)
			num = num / 10
		}
		return pList
	}

	errCount := 0
	testCase := func(posList []s.Position) {
		brd := b.NewBoard()

		for _, pos := range posList {
			brd.PositionMap[pos] = s.PieceStates{true, false, def.FU}
		}

		rsh := NewRshCodeInit()
		var err error
		rsh.Base_P18Black, rsh.Base_P18White, _, _, _, err = getP18fromBoard(brd)
		if err != nil {
			t.Errorf("エラーを検出しました。error = %s", err.Error())
		}

		if len(rsh.Base_P18Black) != 5 {
			t.Errorf("Base_P18Blackの文字数が異常です。p18Black = %s\n posList = %+v", rsh.Base_P18Black.ToString(), posList)
			errCount++
			return
		}
	}

	for i := 0; i < 1000000000; i += 10101 {
		pList := strToPlist(i)
		testCase(pList)
		if errCount > 100 {
			break
		}
	}

	pList := strToPlist(999999999)
	testCase(pList)

}

func Test_getAndPutP16Black_後手駒10000パターン(t *testing.T) {
	// テストパターン
	strToPlist := func(num int) []s.Position {
		pList := make([]s.Position, 0)
		for i := 1; i <= 9; i++ {
			if num < 1 {
				break
			}
			y := num % 10
			pos := s.Position{i, y}
			pList = append(pList, pos)
			num = num / 10
		}
		return pList
	}

	errCount := 0
	testCase := func(posList []s.Position) {
		brd := b.NewBoard()

		for _, pos := range posList {
			brd.PositionMap[pos] = s.PieceStates{false, false, def.FU}
		}

		rsh := NewRshCodeInit()
		var err error
		_, rsh.Base_P18White, _, _, _, err = getP18fromBoard(brd)
		if err != nil {
			t.Errorf("エラーを検出しました。error = %s", err.Error())
		}

		if len(rsh.Base_P18White) != 5 {
			t.Errorf("Base_P18whiteの文字数が異常です。p18Black = %s\n posList = %+v", rsh.Base_P18White.ToString(), posList)
			errCount++
			return
		}

	}

	// パターンを走査する
	for i := 0; i < 1000000000; i += 10101 {
		pList := strToPlist(i)
		testCase(pList)
		if errCount > 100 {
			break
		}
	}

	pList := strToPlist(999999999)
	testCase(pList)

}

func Test_getAndputP16cap_exCap_全パターン(t *testing.T) {

	var rsh *RshCode

	// テストパターン
	testCase := func(capBlack, capWhite int) {

		expectedBoard := b.NewBoard()

		// 盤面の設定
		expectedBoard.Turn = def.BLACK

		expectedBoard.CapturedMap[s.CapArea{def.BLACK, def.FU}] = capBlack
		expectedBoard.CapturedMap[s.CapArea{def.WHITE, def.FU}] = capWhite

		rsh = NewRshCodeInit()
		var err error
		rsh.Base_P18Black, rsh.Base_P18White, rsh.Base_P18Cap, rsh.Add_P18ExCap, rsh.Add_P18Prom, err = getP18fromBoard(expectedBoard)
		if err != nil {
			t.Errorf("エラーが検出されました。error = %s", err.Error())
		}
		putPiecefromP16(rsh)

		//idの決定
		strByte := make([]byte, 0)

		appendString := func(str string) {
			strByte = append(strByte, []byte(str)...)
		}

		appendString(strconv.Itoa(capBlack))
		appendString(",")
		appendString(strconv.Itoa(capWhite))

		// fmt.Printf("id=%s, ", string(strByte))

		assert_EachPositionOfBoard(expectedBoard, rsh.Board, string(strByte), t)
	}

	for black := 0; black < 19; black++ {
		for white := 0; white < 19; white++ {
			if black+white < 19 {
				testCase(black, white)
			}
		}
	}
}

func Test_getAndputP16Prom_単一座標(t *testing.T) {

	var rsh *RshCode

	// テストパターン
	testCase := func(posList []s.Position, plyList []def.Player) {

		expectedBoard := b.NewBoard()

		// 盤面の設定
		expectedBoard.Turn = def.BLACK

		for index, pos := range posList {
			expectedBoard.PositionMap[pos] = s.PieceStates{plyList[index], true, def.FU}
		}

		rsh = NewRshCodeInit()
		var err error
		rsh.Base_P18Black, rsh.Base_P18White, rsh.Base_P18Cap, rsh.Add_P18ExCap, rsh.Add_P18Prom, err = getP18fromBoard(expectedBoard)
		if err != nil {
			t.Errorf("エラーが検出されました。error = %s", err.Error())
		}

		putPiecefromP16(rsh)

		//idの決定
		strByte := make([]byte, 0)

		appendString := func(str string) {
			strByte = append(strByte, []byte(str)...)
		}
		appendString("[")

		for _, pos := range posList {
			appendString(pos.Output())
		}
		appendString("]")

		// fmt.Printf("id=%s, ", string(strByte))

		assert_EachPositionOfBoard(expectedBoard, rsh.Board, string(strByte), t)
	}

	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			testCase([]s.Position{s.Position{x, y}}, []def.Player{def.BLACK})
		}
	}
	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			testCase([]s.Position{s.Position{x, y}}, []def.Player{def.WHITE})
		}
	}

}

func Test_getAndputP16Prom_1から10駒(t *testing.T) {

	var rsh *RshCode

	// テストパターン
	testCase := func(posList []s.Position, plyList []def.Player) {

		var err error
		expectedBoard := b.NewBoard()

		// 盤面の設定
		expectedBoard.Turn = def.BLACK
		for index, pos := range posList {
			expectedBoard.PositionMap[pos] = s.PieceStates{plyList[index], true, def.FU}
			//fmt.Printf("(init)と金を配置します。@%s\n", pos.Output())
		}
		//		fmt.Printf("(init)終了\n")
		// 盤面　→　コード　→　盤面
		rsh = NewRshCodeInit()
		//		fmt.Println("★Builder")
		rsh.Base_P18Black, rsh.Base_P18White, rsh.Base_P18Cap, rsh.Add_P18ExCap, rsh.Add_P18Prom, err = getP18fromBoard(expectedBoard)
		if err != nil {
			t.Errorf("エラーを検出しました。 error = [ %s ]", err)
		}
		//		fmt.Println("★Reader")
		putPiecefromP16(rsh)

		assert_EachPositionOfBoard(expectedBoard, rsh.Board, getIdFromPosList(posList), t)
	}

	for n := 1; n <= 16; n++ {
		//		fmt.Println("------------------------------------------------")
		//		fmt.Printf("[n=%d]", n)
		posList := make([]s.Position, n)
		plyList := make([]def.Player, n)

		for t := 0; t <= 146; t = t + 2 {
			for m := 0; m <= n-1; m++ {
				posList[m], plyList[m] = reverse164ToPiecePosition(m + t)
			}
			testCase(posList, plyList)
		}
	}
}

func Test_164aryToCode64_最小値の長さ(t *testing.T) {

	// min164Array := []int{0, 1, 2, 3, 4}

	// code64 := convert164AryToBase64(min164Array)

	// //fmt.Printf("{1,2,3,4,5} = %s\n", code64.ToString())

	// max164Array := []int{157, 158, 159, 160, 161}
}

func getIdFromPosList(posList []s.Position) string {
	//idの決定
	strByte := make([]byte, 0)

	appendString := func(str string) {
		strByte = append(strByte, []byte(str)...)
	}
	appendString("[")
	for _, pos := range posList {
		appendString(pos.Output())
	}
	appendString("]")

	return string(strByte)
}
