package structs

import (
	def "KifuLibrary-Logic/define"
)

type PieceStates struct {
	Player      def.Player
	IsPromoted  def.IsPromoted
	KindOfPiece def.KindOfPiece
}

func (p PieceStates) GetPieceName() string {
	if !p.IsPromoted {
		switch p.KindOfPiece {
		case def.OH:
			return def.PIECENAME_OH
		case def.KIN:
			return def.PIECENAME_KIN
		case def.GIN:
			return def.PIECENAME_GIN
		case def.KYO:
			return def.PIECENAME_KYO
		case def.KEI:
			return def.PIECENAME_KEI
		case def.KAKU:
			return def.PIECENAME_KAKU
		case def.HISHA:
			return def.PIECENAME_HISHA
		case def.FU:
			return def.PIECENAME_FU
		}
	} else {
		switch p.KindOfPiece {
		case def.GIN:
			return def.PIECENAME_PROMOTED_GIN
		case def.KYO:
			return def.PIECENAME_PROMOTED_KYO
		case def.KEI:
			return def.PIECENAME_PROMOTED_KEI
		case def.KAKU:
			return def.PIECENAME_PROMOTED_KAKU
		case def.HISHA:
			return def.PIECENAME_PROMOTED_HISHA
		case def.FU:
			return def.PIECENAME_PROMOTED_FU
		}
	}
	return ""
}
