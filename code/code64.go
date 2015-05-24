package code

import (
	"fmt"
	"github.com/zaramme/KifCloud-Logic/math"
	"strings"
)

/// <summary>
/// URL表記可能な64進数文字コードを表現するクラス。
/// インスタンスを作成した場合は６４進数と、元の１０進数の両方の値を自動的に保持する。
/// また、getCode,getNumを静的メソッドとして直接使用することも可能。
/// </summary>

const p = 64
const CHARTABLE = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ_-"

type Code64 []*code64char

func NewCode64Nil() Code64 {
	return make(Code64, 0)
}

func NewCode64FromInt(num int) Code64 {
	ary := math.Convert64Ary(num) //64進数配列に変換
	c := make(Code64, 0)
	for i := 0; i < len(ary); i++ {
		c = append(c, newCode64CharfromInt(ary[i]))
	}
	return c
}

func NewCode64FromString(str string) (code64 Code64, err error) {
	code64 = make(Code64, len(str))
	err = nil

	byteList := []byte(str)
	for i := 0; i < len(str); i++ {
		code64char := newCode64CharfromString(string(byteList[i]))
		code64[i] = code64char
	}
	return code64, err
}

func (this Code64) ToInt() int {

	n64ary := make(math.N64ary, 0)

	for i := 0; i < len(this); i++ {
		n64ary = append(n64ary, this[i].Num)
	}

	return math.Reverse64Ary(n64ary)
}

func (this Code64) Length() int {

	return len([]*code64char(this))
}

// 特定の桁数になるように０でパディングする
func (this Code64) Padding(n int) Code64 {
	if len(this) > n {
		return this
	}

	zero := NewCode64FromInt(0)
	for i := len(this); i < n; i++ {
		this = append(zero, this...)
	}

	return this
}

func (this Code64) Unpadding() Code64 {
	for this[0].getNum() == 0 {
		this = this[1:]
	}

	return this
}

func (this Code64) ToString() string {
	b := make([]byte, 0)
	for i := 0; i <= len(this)-1; i++ {
		b = append(b, this[i].Code...)
	}
	return string(b)
}

type code64char struct {
	Num  int
	Code string
}

func newCode64CharfromInt(num int) *code64char {
	if num < 0 {
		fmt.Println("numが不正です")
	}
	instance := new(code64char)
	instance.Num = num
	instance.Code = instance.getCode()
	return instance
}

func newCode64CharfromString(code string) *code64char {
	instance := new(code64char)
	instance.Code = code
	instance.Num = instance.getNum()
	return instance
}

func (this code64char) getCode() string {
	n := this.Num
	str := string(CHARTABLE[n])
	return str
}

func (this code64char) getNum() int {
	num := int(strings.Index(CHARTABLE, this.Code))
	return num
}
