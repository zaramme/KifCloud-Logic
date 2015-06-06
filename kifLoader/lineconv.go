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
	DROP      = "打"
	PROMOTION = "成"
	BRACKET   = "("
)

const RESIGNED = "投了"
const END_PREFIX = "まで"

func convertKifCodeToMoveCode(kifCode string, turn int) (move m.PlayingDescriptor, err error) {

	kifChars := []rune(kifCode)

	var localErr error
	var nextX, nextY int
	var prevX, prevY int

	isPromotion := def.IsPromoted(false)
	var kop def.KindOfPiece

	var prevStart int
	var kopStart int

	player := def.Player(turn%2 == 1)

	endGame := m.NewEndGameWithString(kifCode, player)

	if endGame != nil {
		return endGame, nil
	}

	mv := new(m.Move)

	// 先頭の全角文字をintに変換
	nextX, localErr = convertMBtoInt(kifChars[0])
	if localErr != nil {
		return nil, localErr
	}

	// 二番目の漢数字をintに変換
	nextY, localErr = convertCCtoInt(kifChars[1])
	if localErr != nil {
		return nil, localErr
	}

	if string(kifChars[2]) == PROMOTION {
		// ３文字目が「成」の場合
		kopStart = 3
		isPromotion = true
		prevStart = 5
	} else {
		// ３文字目が「成」以外の場合
		kopStart = 2
		switch string(kifChars[3]) {
		case BRACKET:
			prevStart = 4
		case PROMOTION:
			isPromotion = true
			prevStart = 5
		case DROP:
			prevStart = 0
		default:
			return nil, fmt.Errorf("指し手の変換に失敗しました")
		}
	}

	tmpPromotion := def.IsPromoted(false)
	kop, tmpPromotion, localErr = convertKOPtoKindOfPeace(kifChars[kopStart])
	if localErr != nil {
		return nil, localErr
	}

	if isPromotion == false && tmpPromotion == true {
		isPromotion = true
	}

	if prevStart == 0 {
		prevX, prevY = 0, 0
		if localErr != nil {
			return nil, localErr
		}
	} else {
		prevX, prevY, localErr = convertDuoToPrevPos(kifChars[prevStart : prevStart+2])
	}

	mv.Player = def.Player(player)
	mv.Prev = s.Position{prevX, prevY}
	mv.Next = s.Position{nextX, nextY}
	mv.IsPromoted = isPromotion
	mv.KindOfPiece = kop
	return mv, nil
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

func convertDuoToPrevPos(r []rune) (x, y int, err error) {

	x, err = strconv.Atoi(string(r[0:1]))
	if err != nil {
		return 0, 0, fmt.Errorf("不正な文字コードを検出しました...　%s", string(r[0:1]))
	}

	y, err = strconv.Atoi(string(r[1:2]))
	if err != nil {
		return 0, 0, fmt.Errorf("不正な文字コードを検出しました...　%s", string(r[1:2]))
	}

	return x, y, nil
}
