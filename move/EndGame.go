package move

import (
	def "github.com/zaramme/KifCloud-Logic/define"
	s "strings"
)

type EndGame struct {
	Turn        def.Player
	Winner      def.Player
	IsDrawGame  bool
	IsSuspended bool
	ByOutOfTime bool
	ByFoul      bool
}

type PlayingDescriptor interface {
	ToMoveCode() string
	ToJsCode() string
	ToJpnCode() string
}

const ENDGAME_STRING_RESIGNED = "投了"
const ENDGAME_STRING_SUSPENDED = "中断"
const ENDGAME_STRING_DRAWGAME = "引き分け"
const ENDGAME_STRING_FOUL = "反則"
const ENDGAME_STRING_MADE = "まで"
const ENDGAME_STRING_BLACK = "先手"
const ENDGAME_STRING_WHITE = "後手"

func NewEndGameWithString(str string, turn def.Player) *EndGame {
	endGame := new(EndGame)
	endGame.Turn = turn

	if s.HasPrefix(str, ENDGAME_STRING_RESIGNED) {
		endGame.Winner = !turn
		return endGame
	}

	if s.HasPrefix(str, ENDGAME_STRING_MADE) {
		chars := []rune(str)

		str = string(chars[2:])

		if s.HasPrefix(str, ENDGAME_STRING_BLACK) {
			endGame.Winner = def.BLACK
			return endGame
		}

		if s.HasPrefix(str, ENDGAME_STRING_WHITE) {
			endGame.Winner = def.WHITE
			return endGame
		}

	}

	if s.HasPrefix(str, ENDGAME_STRING_SUSPENDED) {
		endGame.IsSuspended = true
		return endGame
	}

	if s.HasPrefix(str, ENDGAME_STRING_DRAWGAME) {
		endGame.IsDrawGame = true
		return endGame
	}

	if s.HasPrefix(str, ENDGAME_STRING_FOUL) {
		endGame.ByFoul = true
		return endGame
	}

	return nil
}

func (this *EndGame) ToMoveCode() string {
	if this.IsSuspended || this.IsDrawGame {
		return ""
	}
	if this.Winner == def.BLACK {
		return "w_resigned"
	} else {
		return "b_resigned"
	}
}

func (this *EndGame) ToJsCode() string {
	return "END_OF_GAME"
}

func (this *EndGame) ToJpnCode() string {
	if this.ByFoul {
		return "(反則負け)"
	}

	if this.ByOutOfTime {
		return "(時間切れ負け)"
	}

	if this.IsSuspended {
		return "(中断)"
	}

	if this.IsDrawGame {
		return "(引き分け)"
	}
	return "(投了)"
}
