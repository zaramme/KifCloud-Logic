package rsh

import (
	b "github.com/zaramme/KifCloud-Logic/board"
	def "github.com/zaramme/KifCloud-Logic/define"
	//	mv "../move"
	"fmt"
	s "github.com/zaramme/KifCloud-Logic/structs"
	"testing"
)

func __rshM2() {
	fmt.Print()
}
func Test_getAndPutM2_全パターン(t *testing.T) {

	var rsh *RshCode

	// テストパターン
	testCase := func(posList []s.Position, playerList []def.Player) {

		expectedBoard := b.NewBoard()

		// 盤面の設定
		expectedBoard.Turn = def.BLACK

		for index, pos := range posList {

			var kop def.KindOfPiece
			if index >= 2 {
				kop = def.HISHA
			} else {
				kop = def.KAKU
			}
			expectedBoard.PositionMap[pos] = s.PieceStates{playerList[index], false, kop}
		}

		rsh = NewRshCodeInit()
		rsh.Base_M2 = getM2fromBoard(expectedBoard)

		putPieceFromM2(rsh)

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
	posList := []s.Position{s.Position{2, 8}, s.Position{8, 2}, s.Position{2, 2}, s.Position{8, 8}}
	playerList := make([]def.Player, 4) //{def.BLACK, def.WHITE, def.BLACK, def.WHITE}

	for i := 0; i < 10; i++ {

		for z := 1; z <= 16; z++ {
			tmp := z
			playerList[0] = (tmp%2 == 1)
			tmp = tmp / 2
			playerList[1] = (tmp%2 == 1)
			tmp = tmp / 2
			playerList[2] = (tmp%2 == 1)
			tmp = tmp / 2
			playerList[3] = (tmp%2 == 1)
			testCase(posList, playerList)

		}
	}
}
