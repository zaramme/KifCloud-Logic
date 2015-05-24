package define

type KindOfPiece int

func (this KindOfPiece) ToString() string {
	switch this {
	case OH:
		return "OH"
	case KIN:
		return "KIN"
	case GIN:
		return "GIN"
	case KEI:
		return "KEI"
	case KYO:
		return "KYO"
	case KAKU:
		return "KAKU"
	case HISHA:
		return "HISHA"
	case FU:
		return "FU"
	default:
		return ""
	}
}

func (this KindOfPiece) IsPromPiece() bool {
	if this == FU || this == OH || this == KIN {
		return false
	}
	return true
}

type Player bool

type IsPromoted bool

func Isblack() bool {
	return BLACK
}

func (this Player) Output() string {
	if this == BLACK {
		return "(BLACK)"
	} else {
		return "(WHITE)"
	}
	return "(undefined)"
}

func (this Player) OutputInitial() string {
	if this == BLACK {
		return "B"
	} else {
		return "W"
	}
	return "!"
}
func (this Player) OutputMark() string {
	if this == BLACK {
		return "▲"
	}
	return "▽"
}

const (
	OH = iota
	KIN
	GIN
	KEI
	KYO
	KAKU
	HISHA
	FU
)

const (
	MB_1 = "１"
	MB_2 = "２"
	MB_3 = "３"
	MB_4 = "４"
	MB_5 = "５"
	MB_6 = "６"
	MB_7 = "７"
	MB_8 = "８"
	MB_9 = "９"
)

const (
	CC_1 = "一"
	CC_2 = "二"
	CC_3 = "三"
	CC_4 = "四"
	CC_5 = "五"
	CC_6 = "六"
	CC_7 = "七"
	CC_8 = "八"
	CC_9 = "九"
)

type CC int

func (this CC) ToString() string {
	switch this {
	case 1:
		return CC_1
	case 2:
		return CC_2
	case 3:
		return CC_3
	case 4:
		return CC_4
	case 5:
		return CC_5
	case 6:
		return CC_6
	case 7:
		return CC_7
	case 8:
		return CC_8
	case 9:
		return CC_9
	}
	return "#err"
}

type MB int

func (this MB) ToString() string {
	switch this {
	case 1:
		return MB_1
	case 2:
		return MB_2
	case 3:
		return MB_3
	case 4:
		return MB_4
	case 5:
		return MB_5
	case 6:
		return MB_6
	case 7:
		return MB_7
	case 8:
		return MB_8
	case 9:
		return MB_9
	}

	return "#err"
}

const PIECENAME_OH = "玉"
const PIECENAME_KIN = "金"
const PIECENAME_GIN = "銀"
const PIECENAME_KEI = "桂"
const PIECENAME_KYO = "香"
const PIECENAME_KAKU = "角"
const PIECENAME_HISHA = "飛"
const PIECENAME_FU = "歩"

const PIECENAME_PROMOTED_GIN = "成銀"
const PIECENAME_PROMOTED_KEI = "成桂"
const PIECENAME_PROMOTED_KYO = "成香"
const PIECENAME_PROMOTED_KAKU = "馬"
const PIECENAME_PROMOTED_HISHA = "龍"
const PIECENAME_PROMOTED_FU = "と"

const BLACK = true

const WHITE = false
