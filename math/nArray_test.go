package math

import (
	"fmt"
	"testing"
)

func Test_SliceInsert(t *testing.T) {
	act1 := []int{1, 2, 3, 4, 5}

	act1 = SliceInsert(act1, 0, 0)

	exp1 := []int{0, 1, 2, 3, 4, 5}

	if !assertArrayEquarl(exp1, act1) {
		t.Errorf("値が不正です(e1)")
	}

	act2 := []int{1, 2, 3, 4, 5}
	act2 = SliceInsert(act2, 2, 0)
	exp2 := []int{1, 2, 0, 3, 4, 5}
	if !assertArrayEquarl(exp2, act2) {
		t.Errorf("値が不正です(e1)")
	}

	act3 := []int{1, 2, 3, 4, 5}
	act3 = SliceInsert(act3, 5, 0)
	exp3 := []int{1, 2, 3, 4, 5, 0}
	if !assertArrayEquarl(exp3, act3) {
		t.Errorf("値が不正です(e1)")
	}

}

func Test_GetNarray(t *testing.T) {

	act1 := GetNary(4, 2)
	exp1 := []int{0, 0, 1}
	if !assertArrayEquarl(exp1, act1) {
		t.Errorf("値が不正です(e1)")
		for _, v := range act1 {
			fmt.Println(v)
		}
	}

	act2 := GetNary(100, 4)
	exp2 := []int{0, 1, 2, 1}
	if !assertArrayEquarl(exp2, act2) {
		t.Errorf("値が不正です(e2)")
		for _, v := range act2 {
			fmt.Println(v)
		}
	}
}

func Test_ReverseNary(t *testing.T) {
	assert_ReverseNary(t, []int{0, 0, 1}, 100, 10)
	assert_ReverseNary(t, []int{1, 2, 3}, 1+12*2+12*12*3, 12)
	assert_ReverseNary(t, []int{1, 2, 4, 4}, 1+2*163+4*163*163+4*163*163*163, 163)
}

func Test_Reverse164Ary(t *testing.T) {
	assert_Reverse164ary(t, N164ary{0, 0, 1}, 164*164)
	assert_Reverse164ary(t, N164ary{1, 2, 3}, 1+164*2+164*164*3)
}

func Test_Reverse64Ary(t *testing.T) {
	assert_Reverse64ary(t, N64ary{0, 0, 1}, 64*64)
	assert_Reverse64ary(t, N64ary{1, 2, 3}, 1+64*2+64*64*3)
}

func assert_ReverseNary(t *testing.T, target []int, exp int, array int) {
	act := ReverseNary(target, array)

	if exp != act {
		t.Errorf("値が不正です", act, exp)
	}
}

func assert_Reverse164ary(t *testing.T, target N164ary, exp int) {
	act := Reverse164Ary(target)

	if exp != act {
		t.Errorf("値が不正です", act, exp)
	}
}

func assert_Reverse64ary(t *testing.T, target N64ary, exp int) {
	act := Reverse64Ary(target)

	if exp != act {
		t.Errorf("値が不正です", act, exp)
	}
}

func Test_Convert64ary(t *testing.T) {
	act1 := Convert64Ary(64)
	exp1 := []int{0, 1}
	if !assertArrayEquarl(exp1, act1) {
		t.Errorf("値が不正です(e1")
		for _, v := range act1 {
			fmt.Println(v)
		}
	}
}

func Test_Convert164ary(t *testing.T) {
	act1 := Convert164Ary(164 * 164 * 2)
	exp1 := []int{0, 0, 2}
	if !assertArrayEquarl(exp1, act1) {
		t.Errorf("値が不正です(e1")
		for _, v := range act1 {
			fmt.Println(v)
		}
	}
}

func Test_N164ary_Sort(t *testing.T) {
	test := N164ary{4, 108, 1, 0, 25}
	actual := test.Sort()
	expected := N164ary{0, 1, 4, 25, 108}
	if !assertArrayEquarl(expected, actual) {
		t.Errorf("値が不正です(実測値…)", actual)
	}
}

// func Test_sortEachNary(t *testing.T) {
// 	actual64 := n64ary([]int{18, 32, 63, 21, 44})
// 	expected64 := []int{63, 44, 32, 21, 18}
// 	actual64.Sort()

// 	if !assertArrayEquarl(expected64, actual64) {
// 		t.Errorf("値が不正です(e1")
// 		for _, v := range actual64 {
// 			fmt.Println(v)
// 		}
// 	}
// }

// 配列の値を比較するアサーションメソッド
func assertArrayEquarl(expected []int, actual []int) (result bool) {
	if len(expected) != len(actual) {
		return false
	}
	for i, v := range actual {
		if v != expected[i] {
			return false
		}
	}
	return true
}
