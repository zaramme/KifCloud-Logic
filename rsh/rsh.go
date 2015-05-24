package rsh

import (
	b "KifuLibrary-Logic/board"
	"KifuLibrary-Logic/code"
	def "KifuLibrary-Logic/define"
	"KifuLibrary-Logic/math"
	s "KifuLibrary-Logic/structs"
	//"fmt"
)

type Position s.Position
type CapArea s.CapArea
type PieceStates s.PieceStates

//     #region 説明
//     /*
//【SRHコード概要】 RSH
//　将棋の局面情報を、可逆変換可能な文字列で表現するコード体系です。
//
//　（１）仕様
//　・英数、記号を組み合わせた64文字コードを使用(0-1、a-z、A-Z、-(ハイフン)、_(アンダーバー)　)
//　・固定長40文字＋可変長(最大40文字)で構成。
//　・URL対応文字のみ使用なので、Webサービスのアドレスとしてそのまま使用可能
//　・また、駒の種類ごとにブロック化して文字列化されているので
//　　コード同士の数値比較で、盤面の差分を算出可能。
//
// （２）特徴
// ・出現可能性の低い局面パターンの表現を可変長部分に割り当てることで、
// 　平均40～45文字程度という長さで可逆コードを実現しました。
//　　可変長部分の出現条件は以下のとおりです。
//　　　１．成りゴマ一枚につき、１文字追加
//　　　２．「と金の二歩」一枚につき、２文字追加
//　　　３．歩の持ち駒が８枚を超えた場合に、１文字追加
//
//　（３）表現できる局面の範囲
//　以下の局面は表現できません。
//　・通常の本将棋よりも、駒数が多い局面（玉が３枚ある等）
//　・二歩のある局面（と金が含まれる場合を除く）
//　・先手後手のどちらかが、歩を１８枚すべて持っている局面
//  *
//　（その他）桁数のさらなる削減可能性について
//　・現状では、４０種類の駒を８１マス＋先手後手＋持ち駒に対応させている方式ので、
//　盤上の１マスに複数の駒が存在する状態が表現可能となっており、これを削減することで、
//　局面数オーダーを163^40 → 163C40 まで減少させることが可能。
//　ただ、持ち駒は重複可能であったりなどの例外パターンも大きく、
//　変換アルゴリズムが難解となり、さらにコードの可読性、比較性が損なわれるため、採用を見送った。

//　・桂馬・香車に対して、「行き場のない駒」の場合を削除して、
//　　成りゴマの場合は、PropSの空き情報を利用して表現することで、組み合わせ数の削減が可能。
//　　ただし、現行の桂馬・香車の表現文字数(それぞれ5桁)を4桁に削減するためには、
//　　組み合わせ数を1/61まで減らす必要があり、上記の削減では足りない。
//  */
//     #endregion

/// 定数値クラス
const BLACK_CARRY = 81 // 後手番の桁あげ数値
const SEPARETOR = '|'  // 区切り文字

type RshCode struct {
	PropSCodeByKoP map[def.KindOfPiece]int
	BaseIndex      map[def.KindOfPiece]int
	Board          *b.Board
	Code           string
	Int            int
	Base_TK        code.Code64
	Base_M2        code.Code64
	Base_KIN       code.Code64
	Base_GIN       code.Code64
	Base_KEI       code.Code64
	Base_KYO       code.Code64
	Base_P18Black  code.Code64
	Base_P18White  code.Code64
	Base_P18Cap    code.Code64
	Add_TK         code.Code64
	Add_P18ExCap   code.Code64
	Add_P18Prom    code.Code64
	Add_Prom       code.Code64
}

func NewRshCodeInit() *RshCode {
	rsh := new(RshCode)
	rsh.Board = b.NewBoard()
	rsh.Base_TK = make(code.Code64, 3)
	rsh.Base_M2 = make(code.Code64, 5)
	rsh.Base_KIN = make(code.Code64, 5)
	rsh.Base_GIN = make(code.Code64, 5)
	rsh.Base_KEI = make(code.Code64, 5)
	rsh.Base_KYO = make(code.Code64, 5)
	rsh.Base_P18Black = make(code.Code64, 5)
	rsh.Base_P18White = make(code.Code64, 5)
	rsh.Base_P18Cap = make(code.Code64, 2)
	rsh.Add_TK = make(code.Code64, 0)
	rsh.Add_Prom = make(code.Code64, 0)
	rsh.Add_P18ExCap = make(code.Code64, 0)
	rsh.Add_P18Prom = make(code.Code64, 0)

	return rsh
}

// CAUTION : add未対応
func NewRshCodeFromString(s string) (rsh *RshCode, err error) {
	err = nil
	rsh = NewRshCodeInit()

	extract := func(start, end int) string {
		r := []byte(s)
		return string(r[start:end])
	}

	rsh.Base_TK, _ = code.NewCode64FromString(extract(0, 2))
	rsh.Base_M2, _ = code.NewCode64FromString(extract(2, 7))
	rsh.Base_KIN, _ = code.NewCode64FromString(extract(7, 12))
	rsh.Base_GIN, _ = code.NewCode64FromString(extract(12, 17))
	rsh.Base_KEI, _ = code.NewCode64FromString(extract(17, 22))
	rsh.Base_KYO, _ = code.NewCode64FromString(extract(22, 27))
	rsh.Base_P18Black, _ = code.NewCode64FromString(extract(27, 32))
	rsh.Base_P18White, _ = code.NewCode64FromString(extract(32, 37))
	rsh.Base_P18Cap, _ = code.NewCode64FromString(extract(37, 38))

	if len(s) <= 39 {
		return rsh, nil
	}

	composite, _ := code.NewCode64FromString(extract(39, 40))
	existAddProm := false
	//fmt.Printf("composite=%d \n", composite.ToInt())
	rsh.Add_TK, rsh.Add_P18ExCap, existAddProm = divideADDTkAndAddP16ExCap(composite)
	rshLength := len(s)

	if existAddProm {
		//fmt.Printf("Add_Promを読み込んでいます・・・。\n")
		add_PromPadding, _ := code.NewCode64FromString(extract(40, 43))
		rsh.Add_Prom = add_PromPadding.Unpadding()
		if rshLength > 42 {
			//fmt.Printf("(#1)Add_P18Promを読み込んでいます・・・。\n")
			rsh.Add_P18Prom, _ = code.NewCode64FromString(extract(43, rshLength))
		}
	} else {
		if rshLength > 39 {
			//fmt.Printf("(#2)Add_P18Promを読み込んでいます・・・。\n")
			rsh.Add_P18Prom, _ = code.NewCode64FromString(extract(40, rshLength))
		}
	}

	return rsh, nil
}

func (this *RshCode) HasAdditionalCodes() bool {
	result := false

	assertExist := func(code code.Code64) {
		if len(code) > 0 && code.ToInt() != 0 {
			result = true
		}
	}

	assertExist(this.Add_Prom)
	assertExist(this.Add_TK)
	assertExist(this.Add_P18Prom)
	assertExist(this.Add_P18ExCap)

	return result
}

func (this *RshCode) ToString() string {
	strList := make([]byte, 0)

	// base
	strList = append(strList, this.Base_TK.ToString()...)
	strList = append(strList, this.Base_M2.ToString()...)
	strList = append(strList, this.Base_KIN.ToString()...)
	strList = append(strList, this.Base_GIN.ToString()...)
	strList = append(strList, this.Base_KEI.ToString()...)
	strList = append(strList, this.Base_KYO.ToString()...)
	strList = append(strList, this.Base_P18Black.ToString()...)
	strList = append(strList, this.Base_P18White.ToString()...)
	strList = append(strList, this.Base_P18Cap.ToString()...)

	if !this.HasAdditionalCodes() {
		return string(strList) // addが無い場合はスキップ
	}

	// 付加コードの追加
	strList = append(strList, SEPARETOR)

	isExistAddProm := len(this.Add_Prom) > 0
	composite := composeAddTkAndAddP16ExCap(this.Add_TK, this.Add_P18ExCap, isExistAddProm)
	strList = append(strList, composite.ToString()...)

	if len(this.Add_Prom) > 0 {
		addPromPadding := this.Add_Prom.Padding(3)
		strList = append(strList, addPromPadding.ToString()...)
	}

	if len(this.Add_P18Prom) > 0 {
		strList = append(strList, this.Add_P18Prom.ToString()...)
	}

	return string(strList)

}

func composeAddTkAndAddP16ExCap(add_TK, add_P16ExCap code.Code64, isExistAddProm bool) code.Code64 {

	digitTK := 0
	if len(add_TK) > 0 {
		digitTK = add_TK.ToInt()
	}
	digitCap := 0
	if len(add_P16ExCap) > 0 {
		digitCap = add_P16ExCap.ToInt()
	}

	digit := digitTK*6 + digitCap

	if isExistAddProm {
		digit = digit + 32
	}

	composite := code.NewCode64FromInt(digit)

	return composite
}

func divideADDTkAndAddP16ExCap(composite code.Code64) (add_tk, add_p16exCap code.Code64, ExistAddProm bool) {
	compositeInt := composite.ToInt()

	ExistAddProm = false
	if compositeInt >= 32 {
		compositeInt = compositeInt - 32
		ExistAddProm = true
	}

	add_tk = code.NewCode64FromInt((compositeInt / 6))
	add_p16exCap = code.NewCode64FromInt((compositeInt % 6))

	return add_tk, add_p16exCap, ExistAddProm
}

// func (this *RshCode) getTKfromBoard() {
// 	brd := this.board

// 	tkn164Array := this.getPieceCodesByKindOfPiece(def.OH, brd)
// 	tkdigit := math.Reverse164Ary(tkn164Array)
// 	if brd.Turn == def.Player(def.WHITE) {
// 		tkdigit += TK_CARRY
// 	}

// 	this.Base_TK = code.NewCode64FromInt(tkdigit)
// }

// func (this *RshCode) getM2fromBoard() {
// 	brd := this.board

// 	n164Array := make(math.N164ary, 0)
// 	n164Array = append(n164Array, this.getPieceCodesByKindOfPiece(def.HISHA, brd)...)
// 	n164Array = append(n164Array, this.getPieceCodesByKindOfPiece(def.KAKU, brd)...)

// 	this.Base_M2 = this.convert164AryToBase64(n164Array)
// }

/// <summary>
/// 持ち駒の状態を１６４進数コードに変換
/// </summary>
/// <param name="keyp"></param>
/// <param name="lst"></param>
func ConverCapturedUnitto164(cap CapArea, num int) (res math.N164ary) {
	for i := 0; i < num; i++ {
		if cap.Player == def.BLACK {
			res = append(res, 162)
		} else {
			res = append(res, 163)
		}
	}
	return res
}
