package rsh

import (
	b "github.com/zaramme/KifCloud-Logic/board"
	def "github.com/zaramme/KifCloud-Logic/define"
	//	mv "../move"
	"fmt"
	s "github.com/zaramme/KifCloud-Logic/structs"
	"testing"
)

func __rshM4() {
	fmt.Print()
}

func Test_getAndPutKIN_全パターン(t *testing.T) {
	//assert_getAndPutM2_初期盤面(def.KIN, t)
	assert_getAndPutM4_駒１つ全パターン(def.KIN, t)
}

func test_getAndPutGIN_全パターン(t *testing.T) {
	assert_getAndPutM4_駒１つ全パターン(def.GIN, t)
}

func test_getAndPutKEI_全パターン(t *testing.T) {
	assert_getAndPutM4_駒１つ全パターン(def.KEI, t)
}

func test_getAndPutKYO_全パターン(t *testing.T) {
	assert_getAndPutM4_駒１つ全パターン(def.KYO, t)
}

// func assert_getAndPutM2_初期盤面(kop def.KindOfPiece, t *testing.T) {

// 	var rsh *RshCode

// 	// テストパターン
// 	testCase := func(posList []s.Position, playerList []def.Player) {

// 		expectedBoard := b.NewBoardInit()

// 		// 盤面の設定
// 		// expectedBoard.Turn = def.BLACK

// 		// for index, pos := range posList {
// 		// 	expectedBoard.PositionMap[pos] = s.PieceStates{playerList[index], false, kop}
// 		// }

// 		rsh = NewRshCodeInit()
// 		rsh.Base_KIN = getM2fromBoard(expectedBoard)

// 		putPiecefromBase_KIN(rsh)

// 		// idの決定
// 		//		strByte := make([]byte, 0)

// 		// for index, pos := range posList {
// 		// 	strByte = append(strByte, []byte(pos.Output())...)
// 		// 	strByte = append(strByte, []byte(playerList[index].OutputInitial())...)
// 		// 	strByte = append(strByte, []byte(", ")...)
// 		// }

// 		//fmt.Printf("id=%s\n", string(strByte))
// 		assert_EachPositionOfBoard(expectedBoard, rsh.Board, "kintest", t)
// 	}

// 	// パターンの実行
// 	posList := []s.Position{s.Position{1, 1}, s.Position{1, 9}, s.Position{9, 1}, s.Position{9, 9}}
// 	playerList := make([]def.Player, 4) //{def.BLACK, def.WHITE, def.BLACK, def.WHITE}

// 	for i := 0; i < 10; i++ {
// 		for z := 1; z <= 16; z++ {
// 			tmp := z
// 			playerList[0] = (tmp%2 == 1)
// 			tmp = tmp / 2
// 			playerList[1] = (tmp%2 == 1)
// 			tmp = tmp / 2
// 			playerList[2] = (tmp%2 == 1)
// 			tmp = tmp / 2
// 			playerList[3] = (tmp%2 == 1)
// 			testCase(posList, playerList)

// 		}
// 	}

// }

func assert_getAndPutM4_駒１つ全パターン(kop def.KindOfPiece, t *testing.T) {

	var rsh *RshCode

	// テストパターン
	testCase := func(posList []s.Position, playerList []def.Player, capList map[def.Player]int) {

		expectedBoard := b.NewBoard()

		// 盤面の設定
		expectedBoard.Turn = def.BLACK

		for index, pos := range posList {
			expectedBoard.PositionMap[pos] = s.PieceStates{playerList[index], false, kop}
			//fmt.Printf("%s, \n", pos.Output())
		}

		for index, i := range capList {
			expectedBoard.CapturedMap[s.CapArea{index, kop}] = i
		}
		rsh = NewRshCodeInit()

		switch kop {
		case def.KIN:
			rsh.Base_KIN = getKINfromBoard(expectedBoard)
		case def.GIN:
			rsh.Base_GIN = getGINfromBoard(expectedBoard)
		case def.KEI:
			rsh.Base_KEI = getKEIfromBoard(expectedBoard)
		case def.KYO:
			rsh.Base_KYO = getKYOfromBoard(expectedBoard)
		default:
			return
		}

		putPieceFromM4(rsh, kop)

		// idの決定
		strByte := make([]byte, 0)

		for index, pos := range posList {
			strByte = append(strByte, []byte(pos.Output())...)
			strByte = append(strByte, []byte(playerList[index].OutputInitial())...)
			strByte = append(strByte, []byte(", ")...)
		}

		//fmt.Printf("id=%s\n", string(strByte))
		assert_EachPositionOfBoard(expectedBoard, rsh.Board, string(strByte), t)
	}

	// パターンの実行

	//	fmt.Printf("検証中・・・\n")
	posList := make([]s.Position, 1)
	playerList := make([]def.Player, 1)
	capList := make(map[def.Player]int)

	for i := 0; i < 162; i++ {
		pos, player := reverse164ToPiecePosition(i)
		posList[0] = pos
		playerList[0] = player
		testCase(posList, playerList, capList)
	}

	posList = make([]s.Position, 0)
	playerList = make([]def.Player, 0)
	capList = make(map[def.Player]int)
	capList[def.BLACK] = 1
	testCase(posList, playerList, capList)

	posList = make([]s.Position, 0)
	playerList = make([]def.Player, 0)
	capList = make(map[def.Player]int)
	capList[def.WHITE] = 1
	testCase(posList, playerList, capList)

}
