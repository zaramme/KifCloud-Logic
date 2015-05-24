package math

import (
	"sort"
)

type N64ary []int

func Convert64Ary(digit int) (result N64ary) {
	return GetNary(digit, 64)
}

func Reverse64Ary(lst N64ary) int {
	return ReverseNary(lst, 64)
}

type N81ary []int

func Convert81Ary(digit int) (result N81ary) {
	return GetNary(digit, 81)
}

type N164ary []int

func (this N164ary) Sort() N164ary {
	p := []int(this)
	sort.Ints(p)
	return N164ary(p)
}

func (this N164ary) Output() string {
	// var byList []byte

	// fmt.Print(this)
	// for i := 0; i < len(byList); i++ {
	// 	byList = append(byList, []byte(strconv.Itoa(this[i]))...)
	// 	if i != len(byList)-1 {
	// 		byList = append(byList, []byte(",")...)
	// 	}
	// }

	return ""
}

func Convert164Ary(digit int) (result N164ary) {
	return GetNary(digit, 164)
}

func Reverse164Ary(lst N164ary) (result int) {
	return ReverseNary(lst, 164)
}

func GetNary(digit int, ary int) (value []int) {

	value = make([]int, 0)

	for {
		value = append(value, digit%ary) //nで割ったあまりを配列に追加して行く
		digit = digit / ary
		if digit == 0 {
			break
		}
	}
	return value
}

func ReverseNary(lst []int, ary int) (result int) {
	t := 1
	for i := 0; i < len(lst); i++ {
		result += lst[i] * t
		t = t * ary
	}
	return result
}

func SliceInsert(s []int, i int, v int) (r []int) {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}
