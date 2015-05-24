package structs

import (
	def "github.com/zaramme/KifCloud-Logic/define"
	"strconv"
)

type Position struct {
	X int
	Y int
}

func (p Position) IsCaptured() bool {
	if p.X == 0 && p.Y == 0 {
		return true
	}
	return false
}

func (p Position) Output() string {
	var strList []byte

	strList = append(strList, []byte("{")...)

	if p.IsCaptured() {
		strList = append(strList, []byte("Captured")...)
	} else {
		strList = append(strList, []byte(strconv.Itoa(p.X))...)
		strList = append(strList, []byte(",")...)
		strList = append(strList, []byte(strconv.Itoa(p.Y))...)
	}

	strList = append(strList, []byte("}")...)

	return string(strList)
}

func (p Position) OutputJpnCode() string {
	bList := make([]byte, 0)

	appendStr := func(str string) {
		bList = append(bList, []byte(str)...)
	}

	mb := def.MB(p.X)
	appendStr(mb.ToString())
	cc := def.CC(p.Y)
	appendStr(cc.ToString())

	return string(bList)
}

func (p Position) OutputJsCode() string {
	var strList []byte

	strList = append(strList, []byte(strconv.Itoa(p.X))...)
	strList = append(strList, []byte(strconv.Itoa(p.Y))...)

	return string(strList)
}

type PositionMap map[Position]PieceStates
