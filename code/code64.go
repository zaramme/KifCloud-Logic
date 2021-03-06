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

func NewCode64FromInt(num int) (c Code64, err error) {
	ary := math.Convert64Ary(num) //64進数配列に変換
	c = make(Code64, 0)
	for i := 0; i < len(ary); i++ {
		c64char, err := newCode64CharfromInt(ary[i])
		if err != nil {
			return nil, err
		}
		c = append(c, c64char)
	}
	return c, nil
}

func NewCode64FromString(str string) (code64 Code64, err error) {
	code64 = make(Code64, len(str))
	err = nil

	byteList := []byte(str)
	for i := 0; i < len(str); i++ {
		code64char := newCode64CharfromString(string(byteList[len(str)-1-i]))
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

	zero, _ := NewCode64FromInt(0)
	for i := len(this); i < n; i++ {

		this = append(this, zero...)
	}

	return this
}

func (this Code64) Unpadding() Code64 {
	for this[len(this)-1].getNum() == 0 {
		this = this[:(len(this) - 1)]

		if len(this) == 1 {
			return this
		}
	}

	return this
}

func (this Code64) ToString() string {
	b := make([]byte, 0)

	pushFirst := func(s string) []byte {
		first := []byte(s)
		_b := append(first, b...)
		return _b
	}

	for i := 0; i <= len(this)-1; i++ {
		b = pushFirst(this[i].Code)
	}
	return string(b)
}

type code64char struct {
	Num  int
	Code string
}

func newCode64CharfromInt(num int) (c64Char *code64char, err error) {
	if num < 0 {
		err = fmt.Errorf("newCode64fromIntに負の値が入力されました num = %d", num)
		return nil, err
	}
	instance := new(code64char)
	instance.Num = num
	instance.Code = instance.getCode()
	return instance, nil
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
