package structs

import (
	def "github.com/zaramme/KifCloud-Logic/define"
	"testing"
)

func Test_typedef_load(t *testing.T) {
	var KindOfPiece def.KindOfPiece = 1
	if KindOfPiece != 1 {
		t.Errorf("KindOfPiece型の読み込みに失敗しました")
	}

	var player def.Player = true
	if !player {
		t.Errorf("IsBlack型の読み込みに失敗しました")
	}

	var isPromoted def.IsPromoted = true
	if !isPromoted {
		t.Errorf("IsPromoted型の読み込みに失敗しました")
	}
}

func Test_const_load(t *testing.T) {
	if def.BLACK != true {
		t.Errorf("定数値BLACKの読み込みに失敗しました")
	}

	if def.WHITE != false {
		t.Errorf("定数値WHITEの読み込みに失敗しました")
	}

	if def.OH != 1 {
		t.Errorf("定数値OHの読み込みに失敗しました")
	}

	if def.KIN != 2 {
		t.Errorf("定数値KINの読み込みに失敗しました")
	}

	if def.FU != 8 {
		t.Errorf("定数値FUの読み込みに失敗しました")
	}
}

func Test_piecestates_construct(t *testing.T) {
	var ps PieceStates = PieceStates{def.BLACK, true, def.OH}

	if ps.Player != def.BLACK {
		t.Errorf("PieceStatesの初期化に失敗しました(player）")
	}

	if ps.IsPromoted != true {
		t.Errorf("PieceStatesの初期化に失敗しました(isPromoted）")
	}

	if ps.KindOfPiece != def.OH {
		t.Errorf("PieceStatesの初期化に失敗しました(KindOfPiece）")
	}
}

func Test_getPiecename(t *testing.T) {
	var ps PieceStates

	for j := 0; j < 8; j++ {
		var KindOfPiece def.KindOfPiece = def.KindOfPiece(j)
		ps = PieceStates{def.BLACK, false, KindOfPiece}

		var expected string
		switch j {
		case def.OH:
			expected = def.PIECENAME_OH
		case def.KIN:
			expected = def.PIECENAME_KIN
		case def.GIN:
			expected = def.PIECENAME_GIN
		case def.KEI:
			expected = def.PIECENAME_KEI
		case def.KYO:
			expected = def.PIECENAME_KYO
		case def.KAKU:
			expected = def.PIECENAME_KAKU
		case def.HISHA:
			expected = def.PIECENAME_HISHA
		case def.FU:
			expected = def.PIECENAME_FU
		}
		if ps.GetPieceName() != expected {
			t.Errorf("getPieceNameが不正な値を出力しました", ps.GetPieceName())
		}
	}
	for j := 2; j < 8; j++ {
		var KindOfPiece def.KindOfPiece = def.KindOfPiece(j)
		ps = PieceStates{def.BLACK, true, KindOfPiece}

		var expected string
		switch j {
		case def.GIN:
			expected = def.PIECENAME_PROMOTED_GIN
		case def.KEI:
			expected = def.PIECENAME_PROMOTED_KEI
		case def.KYO:
			expected = def.PIECENAME_PROMOTED_KYO
		case def.KAKU:
			expected = def.PIECENAME_PROMOTED_KAKU
		case def.HISHA:
			expected = def.PIECENAME_PROMOTED_HISHA
		case def.FU:
			expected = def.PIECENAME_PROMOTED_FU
		}
		if ps.GetPieceName() != expected {
			t.Errorf("getPieceNameが不正な値を出力しました", ps.GetPieceName())
		}
	}
}
