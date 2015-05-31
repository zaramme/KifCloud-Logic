package rsh

import (
	"fmt"
	b "github.com/zaramme/KifCloud-Logic/board"
	code "github.com/zaramme/KifCloud-Logic/code"
	def "github.com/zaramme/KifCloud-Logic/define"
	math "github.com/zaramme/KifCloud-Logic/math"
	s "github.com/zaramme/KifCloud-Logic/structs"
)

// RshコードからBoardデータを生成する

func ___() {
	fmt.Printf("this is a test")
}

func BuildBoardFromRshCode(rsh *RshCode) *b.Board {

	rsh.Board = b.NewBoard()

	putPieceFromTKandAdd(rsh)
	putPieceFromM2(rsh)
	putPiecefromBase_KIN(rsh)
	putPiecefromBase_GIN(rsh)
	putPiecefromBase_KEI(rsh)
	putPiecefromBase_KYO(rsh)

	putPiecefromP16(rsh)

	if len(rsh.Add_Prom) != 0 {
		applyPiecesPromoted(rsh)
	}
	return rsh.Board
}

func putPieceFromTKandAdd(rsh *RshCode) {
	tk_int := rsh.Base_TK.ToInt()
	add_int := rsh.Add_TK.ToInt()

	rsh.Board.Turn = def.BLACK
	if tk_int >= TK_CARRY {
		tk_int -= TK_CARRY
		rsh.Board.Turn = def.WHITE
	}

	//fmt.Printf("----(reading)tk__int＝{%d}\n", tk_int)

	// 先手駒の配置
	tk_black_int := tk_int % 45
	tk_black_add := add_int % 2

	//fmt.Printf("----tk_black_int＝{%d},tk_black_add =%d}\n", tk_black_int, tk_black_add)

	tk_black_x := tk_black_int/5 + 1
	tk_black_y := tk_black_int%5 + 1

	if tk_black_add != 1 {
		tk_black_y += 5
	}

	//fmt.Printf("----先手王の配置＝{%d,%d}\n", tk_black_x, tk_black_y)
	rsh.Board.PositionMap[s.Position{tk_black_x, tk_black_y}] = s.PieceStates{
		def.BLACK, false, def.OH}

	tk_white_int := tk_int / 45
	tk_white_add := add_int / 2

	//fmt.Printf("----tk_white_int＝{%d},tk_white_add =%d}\n", tk_white_int, tk_white_add)

	tk_white_x := tk_white_int/5 + 1
	tk_white_y := tk_white_int%5 + 1

	//fmt.Printf("----tk_white_x＝{%d},tk_white_y =%d}\n", tk_white_x, tk_white_y)

	if tk_white_add == 1 {
		tk_white_y += 5
	}

	//fmt.Printf("----後手王の配置＝{%d,%d}\n", tk_white_x, tk_white_y)
	rsh.Board.PositionMap[s.Position{tk_white_x, tk_white_y}] = s.PieceStates{
		def.WHITE, false, def.OH}

}

func putPieceFromM2(rsh *RshCode) {
	m2 := rsh.Base_M2
	n164Array := convertCode64To164Ary(m2)

	nHishaAry := n164Array[0:2]
	// fmt.Printf("nHishaAry=")
	// fmt.Print(nHishaAry)
	// fmt.Printf("\n")
	putPieceFromN164Ary(rsh.Board, nHishaAry, def.HISHA, false)

	nKakuAry := n164Array[2:4]
	// fmt.Printf("nKakuAry=")
	// fmt.Print(nKakuAry)
	// fmt.Printf("\n")
	putPieceFromN164Ary(rsh.Board, nKakuAry, def.KAKU, false)
}

func putPieceFromM4(rsh *RshCode, kop def.KindOfPiece) {

	var base code.Code64
	switch kop {
	case def.KIN:
		base = rsh.Base_KIN
	case def.GIN:
		base = rsh.Base_GIN
	case def.KEI:
		base = rsh.Base_KEI
	case def.KYO:
		base = rsh.Base_KYO
	default:
		return
	}

	if len(base) == 0 {
		return
	}

	n164Array := convertCode64To164Ary(base)
	putPieceFromN164Ary(rsh.Board, n164Array, kop, false)
}

func putPiecefromBase_KIN(rsh *RshCode) {
	putPieceFromM4(rsh, def.KIN)
}

func putPiecefromBase_GIN(rsh *RshCode) {
	putPieceFromM4(rsh, def.GIN)
}

func putPiecefromBase_KEI(rsh *RshCode) {
	putPieceFromM4(rsh, def.KEI)
}

func putPiecefromBase_KYO(rsh *RshCode) {
	putPieceFromM4(rsh, def.KYO)
}

func putPiecefromP16(rsh *RshCode) {
	base_Black := rsh.Base_P18Black
	putPiecefromP16_Position(base_Black, def.BLACK, rsh.Board)

	base_White := rsh.Base_P18White
	putPiecefromP16_Position(base_White, def.WHITE, rsh.Board)

	putPiecefromP16_Captured(rsh.Base_P18Cap, rsh.Add_P18ExCap, rsh.Board)

	putPiecefromP16_Promoted(rsh)
}

func putPiecefromP16_Position(code code.Code64, player def.Player, brd *b.Board) {
	code = code.Unpadding()
	digit := code.ToInt()

	//fmt.Printf("digit = %d\n", digit)
	for posX := 1; posX <= 9; posX++ {
		posY := digit % 10
		digit = digit / 10

		if posY == 0 {
			continue
		}

		pos := s.Position{posX, posY}
		//fmt.Printf("駒を配置します…%s,%s\n", pos.Output(), player.Output())
		brd.PositionMap[pos] = s.PieceStates{
			Player:      player,
			KindOfPiece: def.FU,
			IsPromoted:  false,
		}
	}
}

func putPiecefromP16_Promoted(rsh *RshCode) {
	addPromoted := rsh.Add_P18Prom

	length := len(addPromoted)
	//	fmt.Printf("prom = {%s}, length= %d\n", addPromoted.ToString(), length)

	n164arrayList := make([][]int, 0)
	switch {
	case length == 0:
		return
	case 1 <= length && length <= 7:
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted))
	case 7 < length && length <= 14:
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted[:7]))
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted[7:]))
	case 14 < length && length <= 21:
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted[:7]))
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted[7:14]))
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted[14:]))
	case 21 < length:
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted[:7]))
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted[7:14]))
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted[14:21]))
		n164arrayList = append(n164arrayList, convertCode64To164Ary(addPromoted[21:]))
	}

	n164Array := make([]int, 0)
	for i := 0; i < len(n164arrayList); i++ {
		n164Array = append(n164Array, n164arrayList[i]...)
	}

	putPieceFromN164Ary(rsh.Board, n164Array, def.FU, true)
}

func putPiecefromP16_Captured(base, add code.Code64, brd *b.Board) {
	baseDigit := base.ToInt()
	baseN8Array := math.GetNary(baseDigit, 8)
	base_black := baseN8Array[0]
	base_white := 0
	if len(baseN8Array) == 2 {
		base_white = baseN8Array[1]
	}
	addDigit := add.ToInt()

	var add_black, add_white int
	switch {
	case addDigit < 4:
		add_black = addDigit % 2
		add_white = addDigit / 2
	case addDigit == 4:
		add_black = 2
		add_white = 0
	case addDigit == 5:
		add_black = 0
		add_white = 2
	}

	brd.CapturedMap[s.CapArea{def.BLACK, def.FU}] = base_black + add_black*8
	brd.CapturedMap[s.CapArea{def.WHITE, def.FU}] = base_white + add_white*8
}

func applyPiecesPromoted(rsh *RshCode) {
	brd := rsh.Board

	prom := rsh.Add_Prom.ToInt()

	baseN16Array := math.GetNary(prom, 2)

	// fmt.Printf("applyPiecePromoted　= {")
	// for _, value := range baseN16Array {
	// 	fmt.Printf("%d, ", value)
	// }
	// fmt.Printf("}\n")

	currentCount := 0

	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			pos := s.Position{x, y}
			if _, ok := brd.PositionMap[pos]; !ok {
				continue // 探索座標に駒が無い場合はスキップ
			}

			kop := brd.PositionMap[pos].KindOfPiece
			if kop == def.FU || kop == def.OH || kop == def.KIN {
				continue // 歩・金・玉はスキップ
			}

			//			fmt.Printf("(applyPiecePromoted: 対象座標・・・{%d,%d}\n", x, y)

			if baseN16Array[currentCount] == 1 {
				//fmt.Printf("(applyPiecePromoted: {%d,%d}を成り駒とします)\n", x, y)
				brd.PositionMap[pos] = s.PieceStates{
					Player:      brd.PositionMap[pos].Player,
					IsPromoted:  true,
					KindOfPiece: brd.PositionMap[pos].KindOfPiece,
				}
			}
			currentCount++
			if currentCount == len(baseN16Array) {
				return
			}

		}
	}
}
