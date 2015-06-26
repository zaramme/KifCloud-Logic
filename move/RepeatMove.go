package move

import (
	s "github.com/zaramme/KifCloud-Logic/structs"
)

type RepeatMove struct {
	previousMove *Move
	Prev         s.Position
	s.PieceStates
}

func (this *RepeatMove) SetMove(m *Move) {
	this.previousMove = m
	return
}

func (this RepeatMove) GetMove() *Move {

	if this.previousMove == nil {
		return nil
	}
	move := new(Move)

	move.Prev = this.Prev
	move.Next = this.previousMove.Next

	move.PieceStates = this.PieceStates

	return move

}

func (this RepeatMove) ToMoveCode() string {
	mv := this.GetMove()

	if mv == nil {
		return ""
	}

	return mv.ToMoveCode()
}

func (this RepeatMove) ToJsCode() string {
	mv := this.GetMove()

	if mv == nil {
		return ""
	}

	return mv.ToJsCode()
}

func (this RepeatMove) ToJpnCode() string {
	mv := this.GetMove()

	if mv == nil {
		return ""
	}
	return mv.ToJpnCode()
}
