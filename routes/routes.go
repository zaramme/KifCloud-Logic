package routes

import (
	"fmt"
	b "github.com/zaramme/KifCloud-Logic/board"
	k "github.com/zaramme/KifCloud-Logic/kifLoader"
	m "github.com/zaramme/KifCloud-Logic/move"
	r "github.com/zaramme/KifCloud-Logic/rsh"
)

type Routes []Set

type Set struct {
	// RSH code.rsh
	Prev    string
	Current string
	Move    m.PlayingDescriptor
}

type ErrorForUser struct {
	errString string
}

func newErrorForUser(str string) ErrorForUser {
	err := ErrorForUser{errString: str}
	return err
}

func (this ErrorForUser) Error() string {
	return this.errString
}

func (this Set) OutPut() string {
	bList := make([]byte, 0)
	appendStr := func(str string) {
		bList = append(bList, []byte(str)...)
	}
	appendStr(this.Prev)
	appendStr("/")
	appendStr(this.Move.ToMoveCode())
	appendStr(" ... ")
	appendStr(this.Current)

	return string(bList)
}

func NewRoutes() Routes {
	return make(Routes, 0)
}
func NewRoutesFromKifuFile(kif *k.KifFile) (rou Routes, err error) {
	currentboard := b.NewBoardInit()

	Routes := make(Routes, len(kif.Moves))

	for i, desk := range kif.Moves {
		rsh, err := r.ConvertRshFromBoard(currentboard)
		if err != nil {
			return nil, err
		}

		var move *m.Move
		returnFlg := false
		switch t := desk.(type) {
		case *m.EndGame:
			Routes[i] = Set{"", "", desk}
			returnFlg = true
		case *m.RepeatMove:
			move = t.GetMove()
		case *m.Move:
			move = t
		}

		if returnFlg {
			break
		}

		str, err := rsh.ToString()
		if err != nil {
			return nil, err
		}
		Routes[i] = Set{
			str, "", move}

		if i != 0 {
			str2, err := rsh.ToString()
			if err != nil {
				return nil, err
			}
			Routes[i-1].Current = str2
		}

		//fmt.Printf("[NewRouteFromKifuFile]rsh変換... %s => %s \n", rsh.ToString(), move.ToJpnCode())
		err = currentboard.AddMove(move)
		if err != nil {
			return Routes, err
		}

		isValid, errors := currentboard.IsValid()
		if !isValid {
			fmt.Printf("不正な盤面を検出しました。[ %s ]", errors)
			return nil, newErrorForUser(errors)
		}
	}

	return Routes, nil

}
