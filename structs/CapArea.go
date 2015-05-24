package structs

import (
	def "github.com/zaramme/KifCloud-Logic/define"
)

type CapArea struct {
	Player      def.Player
	KindOfPiece def.KindOfPiece
}

func NewCapAreafromPieceStates(ps PieceStates) CapArea {
	instance := CapArea{ps.Player, ps.KindOfPiece}
	return instance
}

func (this CapArea) Output() string {
	var byList []byte
	byList = append(byList, []byte("{")...)

	var player string
	if this.Player == def.BLACK {
		player = "BLACK"
	} else {
		player = "WHITE"
	}
	byList = append(byList, []byte(player)...)
	var kop string
	switch this.KindOfPiece {
	case def.HISHA:
		kop = "HISHA"
	case def.KAKU:
		kop = "KAKU"
	case def.KIN:
		kop = "KIN"
	case def.GIN:
		kop = "GIN"
	case def.KEI:
		kop = "KEI"
	case def.KYO:
		kop = "KYO"
	case def.FU:
		kop = "FU"
	}
	byList = append(byList, []byte(",")...)
	byList = append(byList, []byte(kop)...)
	byList = append(byList, []byte("}")...)
	return string(byList)
}

type CapturedMap map[CapArea]int
