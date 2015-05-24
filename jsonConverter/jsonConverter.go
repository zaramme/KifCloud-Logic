package jsonConverter

import (
	b "github.com/zaramme/KifCloud-Logic/board"
	//	def "../define"
	m "github.com/zaramme/KifCloud-Logic/move"
	r "github.com/zaramme/KifCloud-Logic/rsh"
	s "github.com/zaramme/KifCloud-Logic/structs"
	//	"encoding/json"
	//	"fmt"
)

type shell struct {
	Rsh    string
	Turn   bool
	Pieces [40]*piece
	Info   info
}

type info struct {
	RshCurrent   string
	RshPrev      string
	LastMoveCode string
	MoveText     string
	LastJsCode   string
}

type piece struct {
	Pos   int
	Ply   bool
	IsPrm bool
	Kop   string
	IsCpt bool
}

func BoardToJson(brd *b.Board) (output shell, err error) {

	var shl shell

	shl.Turn = bool(brd.Turn)
	rsh := r.ConvertRshFromBoard(brd)
	shl.Rsh = rsh.ToString()
	shl.Info.RshCurrent = rsh.ToString()
	index := 0
	addPiece := func(p *piece) {
		shl.Pieces[index] = p
		index++
	}

	getPos := func(pos s.Position) int {
		return pos.X*10 + pos.Y
	}

	// 盤上の情報を出力
	for pos, ps := range brd.PositionMap {
		var p piece
		p.Pos = getPos(pos)
		p.IsPrm = bool(ps.IsPromoted)
		p.IsCpt = false
		p.Ply = bool(ps.Player)
		p.Kop = ps.KindOfPiece.ToString()
		addPiece(&p)
	}

	// 駒台の情報を出力
	for cap, value := range brd.CapturedMap {
		var p piece
		p.Pos = 0
		p.IsPrm = false
		p.IsCpt = true
		p.Kop = cap.KindOfPiece.ToString()
		p.Ply = bool(cap.Player)
		for i := 0; i < value; i++ {
			//fmt.Printf("(boardtoJson)持ち駒を配置しています・・・%s\n", p)
			addPiece(&p) // 持ち駒の数だけ繰り返す
		}
	}

	return shl, nil
}

func (this *shell) AppendLastMove(m *m.Move) {
	this.Info.RshPrev = this.Info.RshCurrent
	this.Info.LastMoveCode = m.ToMoveCode()
	this.Info.LastJsCode = m.ToJsCode()
	this.Info.MoveText = m.ToJpnCode()
	this.Info.RshCurrent = ""

	rsh, err := r.NewRshCodeFromString(this.Info.RshPrev)
	if err != nil {
		return
	}
	brd := r.BuildBoardFromRshCode(rsh)

	brd.AddMove(m)

	this.Info.RshCurrent = r.ConvertRshFromBoard(brd).ToString()
}
