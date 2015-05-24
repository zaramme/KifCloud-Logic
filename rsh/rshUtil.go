package rsh

import (
	b "KifuLibrary-Logic/board"
	code "KifuLibrary-Logic/code"
	def "KifuLibrary-Logic/define"
	math "KifuLibrary-Logic/math"
	s "KifuLibrary-Logic/structs"
	// "fmt"
)

// 164進数→PotitionMapに変換
func reverse164ToPiecePosition(num int) (pos s.Position, player def.Player) {
	if num == 162 {
		return s.Position{0, 0}, def.Player(def.BLACK)
	}

	if num == 163 {
		return s.Position{0, 0}, def.Player(def.WHITE)
	}

	if num >= 81 {
		player = def.WHITE
		num -= 81
	} else {
		player = def.Player(def.BLACK)
	}

	x := (num / 9) + 1
	y := (num % 9) + 1
	pos = s.Position{x, y}

	return pos, player
}

func convertCode64To164Ary(code64 code.Code64) math.N164ary {
	if len(code64) == 0 {
		return make(math.N164ary, 0)
	}
	digit := code64.ToInt()
	return math.Convert164Ary(digit)
}

func putPieceFromN164Ary(brd *b.Board, n164Ary math.N164ary, kop def.KindOfPiece, isPromoted def.IsPromoted) {
	if len(n164Ary) == 0 {
		return
	}
	for _, value := range n164Ary {
		pos, player := reverse164ToPiecePosition(value)

		if pos.IsCaptured() {
			brd.CapturedMap[s.CapArea{player, kop}]++
		} else {
			brd.PositionMap[pos] = s.PieceStates{
				Player:      player,
				KindOfPiece: kop,
				IsPromoted:  isPromoted,
			}
		}
	}
}

func convert164AryToBase64(n164Array math.N164ary) (code64 code.Code64) {
	digit := math.Reverse164Ary(n164Array)
	code64 = code.NewCode64FromInt(digit)
	return code64
}

func convert45AryToBase64(n45Array []int) (code64 code.Code64) {
	digit := math.ReverseNary(n45Array, 45)
	code64 = code.NewCode64FromInt(digit)
	return code64
}

// PositionMap→164進数に変換
func convertPiecePositionTo164(pos s.Position, brd *b.Board) int {
	var num int
	if _, ok := brd.PositionMap[pos]; !ok {
		return -1 // 駒が存在しない場合はエラー返却
	}

	ps := brd.PositionMap[pos]
	if ps.Player == def.Player(def.WHITE) {
		num += BLACK_CARRY // 後手番の場合は定数を付加
	}
	num += (pos.X-1)*9 + (pos.Y - 1) // 座標に応じて数値を決定
	return num
}
