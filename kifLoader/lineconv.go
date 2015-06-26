package kifLoader

import (
	"fmt"
	def "github.com/zaramme/KifCloud-Logic/define"
	m "github.com/zaramme/KifCloud-Logic/move"
	s "github.com/zaramme/KifCloud-Logic/structs"
	"strconv"
)

const (
	KOP_OH      = "玉"
	KOP_KIN     = "金"
	KOP_GIN     = "銀"
	KOP_KEI     = "桂"
	KOP_KYO     = "香"
	KOP_KAKU    = "角"
	KOP_HISHA   = "飛"
	KOP_FU      = "歩"
	KOP_GIN_P   = "成銀"
	KOP_KEI_P   = "成桂"
	KOP_KYO_P   = "成香"
	KOP_KAKU_P  = "馬"
	KOP_HISHA_P = "龍"
	KOP_FU_P    = "と"
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

const (
	REPEAT        = "同"
	DROP          = "打"
	PROMOTION     = "成"
	BRACKET_START = "("
	BRACKET_END   = ")"
	BLANK_SB      = " "
	BLANK_MB      = "　"
)

const RESIGNED = "投了"
const END_PREFIX = "まで"

func convertKifCodeToMoveCode(kifCode string, turn int) (move m.PlayingDescriptor, err error) {

	kifChars := []rune(kifCode)
	seq := 0 //現在の文字認識位置

	var next s.Position
	var prev s.Position
	var kop def.KindOfPiece
	var currentP def.IsPromoted
	player := def.Player(turn%2 == 1)

	// 行頭の半角数字をスキップ
	skipSbNums(kifChars, &seq)

	///////////////////////////////////////////////////
	// 対局終了の検出

	skipBlank(kifChars, &seq)
	endGame := m.NewEndGameWithString(string(kifChars[seq:]), player)

	if endGame != nil {
		return endGame, nil
	}

	///////////////////////////////////////////////////
	// 「同」記号の検出
	isRepeat := checkRepeat(kifChars, &seq)
	if isRepeat {
		currentP = false
		currentP = recognizePromotion(kifChars, &seq, currentP)
		kop, currentP = recognizePiece(kifChars, &seq, currentP)
		currentP = recognizePromotion(kifChars, &seq, currentP)
		prev, ok := recognizePrev(kifChars, &seq)
		if !ok {
			return nil, fmt.Errorf("符号の認識に失敗しました。input = %s", kifCode)
		}

		rMove := new(m.RepeatMove)

		rMove.PieceStates = s.PieceStates{player, currentP, kop}
		rMove.Prev = prev

		return rMove, nil
	}

	///////////////////////////////////////////////////
	// 通常の指し手の検出

	mv := new(m.Move)
	currentIsP := def.IsPromoted(false)
	var ok bool

	next, ok = recognizeNext(kifChars, &seq)
	if !ok {
		return nil, fmt.Errorf("符号の認識に失敗しました。input = %s", kifCode)
	}

	currentIsP = recognizePromotion(kifChars, &seq, currentIsP)

	kop, currentIsP = recognizePiece(kifChars, &seq, currentIsP)

	currentIsP = recognizePromotion(kifChars, &seq, currentIsP)

	prev, ok = recognizePrev(kifChars, &seq)
	if !ok {
		return nil, fmt.Errorf("符号の認識に失敗しました。input = %s", kifCode)
	}

	mv.Player = player
	mv.Prev = prev
	mv.Next = next
	mv.IsPromoted = currentIsP
	mv.KindOfPiece = kop

	return mv, nil
}

func skipSbNums(rList []rune, seq *int) {
	skipBlank(rList, seq)
	for {
		_, err := strconv.Atoi(string(rList[*seq]))
		if err != nil {
			break
		}
		*seq++
	}
}

func skipBlank(rList []rune, seq *int) {
	for {
		if len(rList) == *seq {
			break
		}

		str := string(rList[*seq])
		if str == BLANK_MB || str == BLANK_SB {
			*seq++
			continue
		}
		break
	}
}

func recognizePiece(rList []rune, seq *int, currentP def.IsPromoted) (kop def.KindOfPiece, isP def.IsPromoted) {
	skipBlank(rList, seq)

	isP = currentP

	hit := false

	switch string(rList[*seq]) {
	case KOP_OH:
		kop = def.OH
		hit = true
	case KOP_KIN:
		kop = def.KIN
		hit = true
	case KOP_GIN:
		kop = def.GIN
		hit = true
	case KOP_KEI:
		kop = def.KEI
		hit = true
	case KOP_KYO:
		kop = def.KYO
		hit = true
	case KOP_KAKU:
		kop = def.KAKU
		hit = true
	case KOP_HISHA:
		kop = def.HISHA
		hit = true
	case KOP_FU:
		kop = def.FU
		hit = true
	case KOP_GIN_P:
		kop = def.GIN
		isP = true
		hit = true
	case KOP_KEI_P:
		kop = def.KEI
		isP = true
		hit = true
	case KOP_KYO_P:
		kop = def.KYO
		isP = true
		hit = true
	case KOP_KAKU_P:
		kop = def.KAKU
		isP = true
		hit = true
	case KOP_HISHA_P:
		kop = def.HISHA
		isP = true
		hit = true
	case KOP_FU_P:
		kop = def.FU
		isP = true
		hit = true
	}

	if hit {
		*seq++
	}
	return

}

func recognizePromotion(rList []rune, seq *int, currentP def.IsPromoted) (isP def.IsPromoted) {
	skipBlank(rList, seq)
	isP = currentP

	if string(rList[*seq]) == PROMOTION {
		*seq++
		isP = true
	}

	return isP
}

func recognizeNext(rList []rune, seq *int) (next s.Position, ok bool) {
	skipBlank(rList, seq)

	X, err := convertMBtoInt(rList[*seq])
	if err != nil {
		return next, false
	}
	*seq++

	Y, err := convertCCtoInt(rList[*seq])
	if err != nil {
		return next, false
	}

	*seq++
	next = s.Position{X, Y}

	return next, true

}

func recognizePrev(rList []rune, seq *int) (prev s.Position, ok bool) {
	skipBlank(rList, seq)

	if string(rList[*seq]) == DROP {
		*seq++
		return s.Position{0, 0}, true
	}

	if string(rList[*seq]) != BRACKET_START {
		return prev, false
	}
	*seq++

	X, err := strconv.Atoi(string(rList[*seq]))
	if err != nil {
		return prev, false
	}
	*seq++

	Y, err := strconv.Atoi(string(rList[*seq]))
	if err != nil {
		return prev, false
	}
	*seq++

	if string(rList[*seq]) != BRACKET_END {
		return prev, false
	}
	*seq++

	prev = s.Position{X, Y}

	return prev, true

}

func convertMBtoInt(r rune) (i int, err error) {
	switch string(r) {
	case MB_1:
		return 1, nil
	case MB_2:
		return 2, nil
	case MB_3:
		return 3, nil
	case MB_4:
		return 4, nil
	case MB_5:
		return 5, nil
	case MB_6:
		return 6, nil
	case MB_7:
		return 7, nil
	case MB_8:
		return 8, nil
	case MB_9:
		return 9, nil
	}
	return 0, fmt.Errorf("指し手の行頭は数字である必要があります", string(r))

}

func convertCCtoInt(r rune) (i int, err error) {
	switch string(r) {
	case CC_1:
		return 1, nil
	case CC_2:
		return 2, nil
	case CC_3:
		return 3, nil
	case CC_4:
		return 4, nil
	case CC_5:
		return 5, nil
	case CC_6:
		return 6, nil
	case CC_7:
		return 7, nil
	case CC_8:
		return 8, nil
	case CC_9:
		return 9, nil
	}
	return 0, fmt.Errorf("指し手の変換に失敗しました", string(r))

}

func convertKOPtoKindOfPeace(r rune) (kop def.KindOfPiece, isP def.IsPromoted, err error) {
	switch string(r) {
	case KOP_OH:
		return def.OH, false, nil
	case KOP_KIN:
		return def.KIN, false, nil
	case KOP_GIN:
		return def.GIN, false, nil
	case KOP_KEI:
		return def.KEI, false, nil
	case KOP_KYO:
		return def.KYO, false, nil
	case KOP_KAKU:
		return def.KAKU, false, nil
	case KOP_HISHA:
		return def.HISHA, false, nil
	case KOP_FU:
		return def.FU, false, nil
	case KOP_GIN_P:
		return def.GIN, true, nil
	case KOP_KEI_P:
		return def.KEI, true, nil
	case KOP_KYO_P:
		return def.KYO, true, nil
	case KOP_KAKU_P:
		return def.KAKU, true, nil
	case KOP_HISHA_P:
		return def.HISHA, true, nil
	case KOP_FU_P:
		return def.FU, true, nil
	}
	return 0, false, fmt.Errorf("不正な文字コードを検出しました・・・%s", string(r))

}

func checkRepeat(rList []rune, seq *int) bool {
	switch string(rList[*seq]) {
	case REPEAT:
		*seq++
		return true
	}
	return false
}

func checkBlank(r rune) bool {
	switch string(r) {
	case BLANK_SB:
		return true
	case BLANK_MB:
		return true
	}
	return false
}
