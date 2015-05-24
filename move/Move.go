package move

import (
	"errors"
	def "github.com/zaramme/KifCloud-Logic/define"
	s "github.com/zaramme/KifCloud-Logic/structs"
	"strconv"
)

type Move struct {
	Prev s.Position
	Next s.Position
	s.PieceStates
	IsResigned bool
}

const MOVECODE_BLACK = "b"
const MOVECODE_WHITE = "w"
const MOVECODE_OH = "OH"
const MOVECODE_KAKU = "KA"
const MOVECODE_HISHA = "HI"
const MOVECODE_KIN = "KI"
const MOVECODE_GIN = "GI"
const MOVECODE_KEI = "KE"
const MOVECODE_KYO = "KY"
const MOVECODE_FU = "FU"

const MOVECODE_RESIGNED = "resigned"

const MOVECODE_PROMOTED = "!"

const JS_MOVECODE_SEPARATOR = ","

func NewMoveFromMoveCode(moveCode string) *Move {
	move := new(Move)
	byList := []byte(moveCode)

	var player def.Player
	if value, err := convertMoveCodeToPlayer(byList); err == nil {
		player = value
	}

	isResigned := convertMoveCodeIsResigned(byList)

	if isResigned {
		move.Player = player
		move.IsResigned = true
		return move
	}

	// コードからNextを取得
	var next s.Position
	if value, ok := convertMoveCodeToMoveNext(byList); ok == nil {
		next = value
	} else {
		return nil
	}
	move.Next = next

	// コードからPrevを取得
	var prev s.Position
	if value, ok := convertMoveCodeToMovePrev(byList); ok == nil {
		prev = value
	} else {
		return nil
	}

	move.Prev = prev

	var kop def.KindOfPiece
	if value, ok := convertMoveCodeToKindOfPiece(byList); ok == nil {
		kop = value
	} else {
		return nil
	}

	var isPromoted def.IsPromoted
	if value, ok := convertMoveCodeToIsPromoted(byList); ok == nil {
		isPromoted = def.IsPromoted(value)
	} else {
		return nil
	}

	move.Player = player
	move.KindOfPiece = kop
	move.IsPromoted = isPromoted
	move.IsResigned = false

	return move
}

func NewCapAreafromMove(move *Move) s.CapArea {
	instance := s.CapArea{move.Player, move.KindOfPiece}
	return instance
}

func convertMoveCodeToPlayer(byList []byte) (def.Player, error) {
	byListPrev := byList[0:1]
	strPlayer := string(byListPrev)

	errorValue := def.Player(def.WHITE)

	switch strPlayer {
	case MOVECODE_BLACK:
		return def.BLACK, nil
	case MOVECODE_WHITE:
		return def.WHITE, nil
	default:
		return errorValue, errors.New("playerの値が不正です。")
	}
}
func convertMoveCodeIsResigned(byList []byte) bool {
	if len(byList) < 10 {
		return false
	}
	byListResigned := byList[2:10]

	if string(byListResigned) == MOVECODE_RESIGNED {
		return true
	}
	return false
}

func convertMoveCodeToMoveNext(byList []byte) (s.Position, error) {
	byListPrev := byList[1:3]
	strPrev := string(byListPrev)

	errorValue := s.Position{0, 0}

	var intPrev int
	if value, err := strconv.Atoi(strPrev); err == nil {
		intPrev = value
	} else {
		return errorValue, err
	}

	if intPrev < 1 || 99 < intPrev {
		return errorValue, errors.New("nextの値が不正です。")
	}

	pos := s.Position{intPrev / 10, intPrev % 10}
	return pos, nil
}

func convertMoveCodeToMovePrev(byList []byte) (s.Position, error) {
	byListPrev := byList[6:8]
	strPrev := string(byListPrev)

	errorValue := s.Position{0, 0}

	var intPrev int
	if value, err := strconv.Atoi(strPrev); err == nil {
		intPrev = value
	} else {
		return errorValue, err
	}

	if intPrev < 0 || 99 < intPrev {
		return errorValue, errors.New("Prevの値が不正です。")
	}

	Prev := s.Position{intPrev / 10, intPrev % 10}

	return Prev, nil
}

func convertMoveCodeToKindOfPiece(byList []byte) (def.KindOfPiece, error) {
	byListKind := byList[3:5]
	strKind := string(byListKind)
	var kind def.KindOfPiece
	var err error = nil

	switch strKind {
	case MOVECODE_OH:
		kind = def.OH
	case MOVECODE_HISHA:
		kind = def.HISHA
	case MOVECODE_KAKU:
		kind = def.KAKU
	case MOVECODE_KIN:
		kind = def.KIN
	case MOVECODE_GIN:
		kind = def.GIN
	case MOVECODE_KEI:
		kind = def.KEI
	case MOVECODE_KYO:
		kind = def.KYO
	case MOVECODE_FU:
		kind = def.FU
	default:
		kind = -1
		err = errors.New("駒変換に失敗しました。")
	}

	return kind, err
}

func convertMoveCodeToIsPromoted(byList []byte) (bool, error) {

	if len(string(byList)) < 9 {
		return false, nil
	}

	byListKind := byList[8:9]
	strIsPromoted := string(byListKind)

	if strIsPromoted != MOVECODE_PROMOTED {
		return false, errors.New("不正な文字コードを検出しました。")
	}

	return true, nil
}

func (this *Move) ToMoveCode() string {
	bList := make([]byte, 0)

	appendStr := func(str string) {
		bList = append(bList, []byte(str)...)
	}

	if this.Player {
		appendStr(MOVECODE_BLACK)
	} else {
		appendStr(MOVECODE_WHITE)
	}

	if this.IsResigned {
		// 投了
		appendStr("_resigned")
		return string(bList)
	}

	nextX := strconv.Itoa(this.Next.X)
	nextY := strconv.Itoa(this.Next.Y)

	appendStr(nextX)
	appendStr(nextY)

	appendStr(getMoveCodeByKOP(this.KindOfPiece))

	appendStr("_")

	prevX := strconv.Itoa(this.Prev.X)
	prevY := strconv.Itoa(this.Prev.Y)

	appendStr(prevX)
	appendStr(prevY)

	if this.IsPromoted {
		appendStr(MOVECODE_PROMOTED)
	}

	return string(bList)
}
func (this *Move) ToJpnCode() string {
	bList := make([]byte, 0)

	appendStr := func(str string) {
		bList = append(bList, []byte(str)...)
	}

	if this.IsResigned {
		appendStr("投了")
		return string(bList)
	}
	appendStr(this.Player.OutputMark())
	appendStr(this.Next.OutputJpnCode())
	appendStr(this.GetPieceName())

	if this.Prev.IsCaptured() {
		appendStr("打")
	} else {
		appendStr("(")
		appendStr(strconv.Itoa(this.Prev.X))
		appendStr(strconv.Itoa(this.Prev.Y))
		appendStr(")")
	}
	return string(bList)
}

func (this *Move) ToJsCode() string {

	if this.IsResigned {
		return ""
	}

	bList := make([]byte, 0)

	appendStr := func(str string) {
		bList = append(bList, []byte(str)...)
	}

	appendStr(this.Prev.OutputJsCode())
	appendStr(JS_MOVECODE_SEPARATOR)
	appendStr(this.Next.OutputJsCode())
	appendStr(JS_MOVECODE_SEPARATOR)
	appendStr(this.KindOfPiece.ToString())
	appendStr(JS_MOVECODE_SEPARATOR)
	if this.IsPromoted {
		appendStr("true")
	} else {
		appendStr("false")
	}
	return string(bList)
}

func getMoveCodeByKOP(kop def.KindOfPiece) string {
	switch kop {
	case def.OH:
		return MOVECODE_OH
	case def.KIN:
		return MOVECODE_KIN
	case def.GIN:
		return MOVECODE_GIN
	case def.KEI:
		return MOVECODE_KEI
	case def.KYO:
		return MOVECODE_KYO
	case def.KAKU:
		return MOVECODE_KAKU
	case def.HISHA:
		return MOVECODE_HISHA
	case def.FU:
		return MOVECODE_FU
	default:
		return ""
	}
}
