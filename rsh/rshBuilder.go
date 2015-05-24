package rsh

// 局面をRSHに変換するクラス

import (
	"fmt"
	b "github.com/zaramme/KifCloud-Logic/board"
	code "github.com/zaramme/KifCloud-Logic/code"
	def "github.com/zaramme/KifCloud-Logic/define"
	math "github.com/zaramme/KifCloud-Logic/math"
	s "github.com/zaramme/KifCloud-Logic/structs"
	m "math"
)

const TK_CARRY = 45 * 45 // TKの先手・後手情報の桁上げ数値

func __() {
	fmt.Printf("test")

}

func ConvertRshFromBoard(brd *b.Board) *RshCode {
	rsh := new(RshCode)
	rsh.Base_TK, rsh.Add_TK = getTKfromBoard(brd)
	rsh.Base_M2 = getM2fromBoard(brd)
	rsh.Base_KIN = getKINfromBoard(brd)
	rsh.Base_GIN = getGINfromBoard(brd)
	rsh.Base_KEI = getKEIfromBoard(brd)
	rsh.Base_KYO = getKYOfromBoard(brd)
	rsh.Base_P18Black, rsh.Base_P18White, rsh.Base_P18Cap, rsh.Add_P18ExCap, rsh.Add_P18Prom = getP18fromBoard(brd)
	rsh.Add_Prom = getPromfromBoard(brd)
	return rsh
}

func getTKfromBoard(brd *b.Board) (tk, add_tk code.Code64) {
	// その座標の駒が王かどうかを判定して、王ならPieceStatesを返す
	search := func(x, y int) (isFound bool, ps s.PieceStates) {

		if value, ok := brd.PositionMap[s.Position{x, y}]; ok {

			if value.KindOfPiece != def.OH {
				return false, ps
			}
			return true, value
		} else {
			return false, ps
		}
	}

	int_tk_black := 0
	int_tk_white := 0
	add_tk_black := 0
	add_tk_white := 0
	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {

			// 王以外はスキップ
			isFound, ps := search(x, y)
			if !isFound {
				continue
			}

			switch ps.Player {
			case def.BLACK:
				int_tk_black, add_tk_black = convertTKonIntAndAddFromPiece(s.Position{x, y}, ps.Player)
			case def.WHITE:
				int_tk_white, add_tk_white = convertTKonIntAndAddFromPiece(s.Position{x, y}, ps.Player)
			}

		}
	}

	shift := 0
	if brd.Turn == def.WHITE {
		shift = TK_CARRY
	}

	tk_int := int_tk_black + int_tk_white*45 + shift
	//fmt.Printf("----(building)tk__int＝{%d}\n", tk_int)

	tk = code.NewCode64FromInt(tk_int)

	add_tk = code.NewCode64FromInt(add_tk_black + add_tk_white*2)

	return tk, add_tk
}

// 玉の座標を４５進数＋シフトに変換
func convertTKonIntAndAddFromPiece(pos s.Position, player def.Player) (posCode int, add int) {

	// 座標が６以上の場合はシフト
	var shift = false
	if pos.Y >= 6 {
		shift = true
		pos.Y -= 5
	}

	// 45進数に変換
	posCode = (pos.X-1)*5 + (pos.Y - 1)

	switch {
	case player == def.BLACK && shift == false:
		add = 1
	case player == def.WHITE && shift == true:
		add = 1
	default:
		add = 0
	}

	return posCode, add
}

func getM2fromBoard(brd *b.Board) code.Code64 {

	n164Array := make(math.N164ary, 0)
	n164Array = append(n164Array, getPieceCodesByKindOfPiece(def.HISHA, brd)...)
	n164Array = append(n164Array, getPieceCodesByKindOfPiece(def.KAKU, brd)...)

	return convert164AryToBase64(n164Array)
}

func getKINfromBoard(brd *b.Board) code.Code64 {
	n164Array := getPieceCodesByKindOfPiece(def.KIN, brd)
	return convert164AryToBase64(n164Array)
}

func getGINfromBoard(brd *b.Board) code.Code64 {
	n164Array := getPieceCodesByKindOfPiece(def.GIN, brd)
	return convert164AryToBase64(n164Array)
}

func getKEIfromBoard(brd *b.Board) code.Code64 {
	n164Array := getPieceCodesByKindOfPiece(def.KEI, brd)
	return convert164AryToBase64(n164Array)
}

func getKYOfromBoard(brd *b.Board) code.Code64 {
	n164Array := getPieceCodesByKindOfPiece(def.KYO, brd)
	return convert164AryToBase64(n164Array)
}

func getP18fromBoard(brd *b.Board) (black, white, cap, add_Prom, add_ExCap code.Code64) {
	digitBlack := 0
	digitWhite := 0

	// 盤上の駒を走査(black、White、add_Prom)
	n164Array := make(math.N164ary, 0)
	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {

			// 座標に駒がない場合はスキップ
			if _, ok := brd.PositionMap[s.Position{x, y}]; !ok {
				continue
			}
			value := brd.PositionMap[s.Position{x, y}]
			// 歩以外はスキップ
			if value.KindOfPiece != def.FU {
				continue
			}
			// 成り駒として追加する
			if value.IsPromoted == true {
				digit := convertPiecePositionTo164(s.Position{x, y}, brd)
				n164Array = append(n164Array, digit)
				continue
			}
			// 盤上の駒として追加する
			switch value.Player {
			case def.BLACK:
				digitBlack += y * int(m.Pow(10, float64(x-1)))
			case def.WHITE:
				digitWhite += y * int(m.Pow(10, float64(x-1)))
			}
		}
	}

	//fmt.Printf("black = %d\n", digitBlack)
	//fmt.Printf("white = %d\n", digitWhite)
	black = code.NewCode64FromInt(digitBlack)
	white = code.NewCode64FromInt(digitWhite)

	promLength := len(n164Array)

	add_Prom = make(code.Code64, 0)
	appendCode64 := func(n164array []int) {
		add_Prom = append(add_Prom, convert164AryToBase64(n164array)...)
	}

	appendCode64seven := func(n164array []int) {
		code64seven := convert164AryToBase64(n164array)
		for len(code64seven) < 7 {
			code64seven = append(code64seven, code.NewCode64FromInt(0)...)
		}
		add_Prom = append(add_Prom, code64seven...)
	}

	// fmt.Printf("n164ary = ", n164Array)
	// fmt.Printf("length = [%d]\n", len(n164Array))
	switch {
	case promLength == 0:
		add_Prom = make(code.Code64, 0)
	case 1 <= promLength && promLength <= 5:
		appendCode64(n164Array)
	case 6 <= promLength && promLength <= 10:
		appendCode64seven(n164Array[:5])
		appendCode64(n164Array[5:])
	case 11 <= promLength && promLength <= 15:
		appendCode64seven(n164Array[:5])
		appendCode64seven(n164Array[5:10])
		appendCode64(n164Array[10:])
	case 16 <= promLength && promLength <= 18:
		appendCode64seven(n164Array[:5])
		appendCode64seven(n164Array[5:10])
		appendCode64seven(n164Array[10:15])
		appendCode64(n164Array[15:])
		//		fmt.Printf("(all)----%d", add_Prom.ToString())
	}

	// fmt.Printf("digitBlack = %d\n", digitBlack)
	// fmt.Printf("digitWhite = %d\n", digitWhite)

	// 持ち駒を格納（cap,add_ExCap)
	capBlack := s.CapArea{def.BLACK, def.FU}
	digitAddBlack := 0
	digitCapturedBlack := 0
	if value, ok := brd.CapturedMap[capBlack]; ok {
		switch {
		case value < 8:
			digitCapturedBlack = value
		case 8 <= value && value < 16:
			digitCapturedBlack = value - 8
			digitAddBlack = 1
		case 16 <= value:
			digitCapturedBlack = value - 16
			digitAddBlack = 2
		}
	}
	capWhite := s.CapArea{def.WHITE, def.FU}
	digitCapturedWhite := 0
	digitAddWhite := 0
	if value, ok := brd.CapturedMap[capWhite]; ok {
		switch {
		case value < 8:
			digitCapturedWhite = value
		case 8 <= value && value < 16:
			digitCapturedWhite = value - 8
			digitAddWhite = 1
		case 16 <= value:
			digitCapturedWhite = value - 16
			digitAddWhite = 2
		}
	}

	digitAddBlackAndWhite := 0

	switch {
	case digitAddBlack != 2 && digitAddWhite != 2:
		digitAddBlackAndWhite = digitAddBlack + digitAddWhite*2
	case digitAddBlack == 2 && digitAddWhite != 2:
		digitAddBlackAndWhite = 4
	case digitAddBlack != 2 && digitAddWhite == 2:
		digitAddBlackAndWhite = 5
	}

	cap = code.NewCode64FromInt(digitCapturedBlack + digitCapturedWhite*8)
	add_ExCap = code.NewCode64FromInt(digitAddBlackAndWhite)
	return black, white, cap, add_ExCap, add_Prom
}

func getPromfromBoard(brd *b.Board) code.Code64 {
	pieceCount := 0
	promMap := make([]int, 16)

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

			if !brd.PositionMap[pos].IsPromoted {
				promMap[pieceCount] = 0
				pieceCount++
				continue // 成り駒ではない場合はスキップ
			}

			promMap[pieceCount] = 1
			pieceCount++
		}
	}

	var promDigit int

	// promMapを１０進数にパース
	current := 1
	for i := 0; i < 16; i++ {
		promDigit += promMap[i] * current
		current *= 2
	}

	if promDigit == 0 {
		return code.NewCode64Nil()
	}

	// ２文字以下の場合は３文字まで０を埋める
	prom := code.NewCode64FromInt(promDigit)
	prom = prom.Padding(3)
	return prom
}

func getPieceCodesByKindOfPiece(kop def.KindOfPiece, brd *b.Board) (result math.N164ary) {
	tkArray := make(math.N164ary, 0)

	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			pos := s.Position{x, y}
			if _, ok := brd.PositionMap[pos]; !ok {
				continue
			}

			if brd.PositionMap[pos].KindOfPiece != kop {
				continue
			}
			digit := convertPiecePositionTo164(pos, brd)
			tkArray = append(tkArray, digit)
		}
	}

	capBlack := s.CapArea{def.Player(def.BLACK), def.KindOfPiece(kop)}
	capWhite := s.CapArea{def.Player(def.WHITE), def.KindOfPiece(kop)}

	if value, ok := brd.CapturedMap[capBlack]; ok {
		for i := 0; i < value; i++ {
			tkArray = append(tkArray, 162)
		}
	}

	if value, ok := brd.CapturedMap[capWhite]; ok {
		for i := 0; i < value; i++ {
			tkArray = append(tkArray, 163)
		}
	}

	sortedTkArray := tkArray.Sort()
	return sortedTkArray
}
