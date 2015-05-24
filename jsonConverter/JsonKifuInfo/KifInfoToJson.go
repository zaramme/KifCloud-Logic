package JsonKifuInfo

import ()

type Shell struct {
	KifuInfos []*KifuInfo
}

type KifuInfo struct {
	UserID      string
	KifuID      int
	BlackPlayer string
	WhitePlayer string
	Winner      bool
	PlayingDate string
}

func NewKifuInfo(n int) *Shell {
	i := make([]*KifuInfo, n)
	r := new(Shell)
	r.KifuInfos = i
	return r
}
