package routes

import (
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
		rsh := r.ConvertRshFromBoard(currentboard)

		move, ok := desk.(*m.Move)

		if !ok {
			Routes[i] = Set{
				"", "", desk}
			break
		}

		Routes[i] = Set{
			rsh.ToString(), "", move}
		if i != 0 {
			Routes[i-1].Current = rsh.ToString()
		}
		err := currentboard.AddMove(move)

		if err != nil {
			return Routes, err
		}
	}

	return Routes, nil

}
