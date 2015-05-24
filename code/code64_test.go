package code

import (
	"testing"
)

func Test_NewCode64FromInt_複数桁をコンストラクタ(t *testing.T) {
	var target = NewCode64FromInt(64*64*58 + 64*20 + 33)

	if target[0].Num != 33 {
		t.Errorf("期待値と数値が異なります i=0", target[0].Num)
	}
	if target[1].Num != 20 {
		t.Errorf("期待値と数値が異なります i=1", target[1].Num)
	}
	if target[2].Num != 58 {
		t.Errorf("期待値と数値が異なります i=2", target[2].Num)
	}
}

func Test_NewCode64FromInt_一桁をコンストラクタ(t *testing.T) {
	var target = NewCode64FromInt(30)

	if target[0].Num != 30 {
		t.Errorf("期待値と数値が異なります", target[0].Num)
	}
	if target[0].Code != "k" {
		t.Errorf("期待値と数値が異なります", target[0].Code)
	}
}

// func Test_NewCode64FromString_ゼロプレフィックス(t *testing.T) {
// 	var target, _ = NewCode64FromString("004")

// 	if target.ToInt() != 4 {
// 		t.Errorf("期待値と数値が異なります。%s = %s", target.ToInt(), 4)
// 	}

// }

func Test_Unpadding(t *testing.T) {
	var target, _ = NewCode64FromString("0004")

	actual := target.Unpadding()

	if actual.ToInt() != 4 {
		t.Errorf("期待値と数値が異なります。%s = %s", target.ToInt(), 4)
	}
}

func Test_newCode64charfromInt(t *testing.T) {
	var actual *code64char
	actual = newCode64CharfromInt(0)
	assert_newCode64CharfromInt(actual, "0", t)

	actual = newCode64CharfromInt(10)
	assert_newCode64CharfromInt(actual, "a", t)

	actual = newCode64CharfromInt(20)
	assert_newCode64CharfromInt(actual, "f", t)

	actual = newCode64CharfromInt(30)
	assert_newCode64CharfromInt(actual, "k", t)

	actual = newCode64CharfromInt(40)
	assert_newCode64CharfromInt(actual, "p", t)

	actual = newCode64CharfromInt(50)
	assert_newCode64CharfromInt(actual, "u", t)

	actual = newCode64CharfromInt(60)
	assert_newCode64CharfromInt(actual, "z", t)

	actual = newCode64CharfromInt(63)
	assert_newCode64CharfromInt(actual, "-", t)

}

func Test_newCode64charfromString(t *testing.T) {
	var actual *code64char

	actual = newCode64CharfromString("0")
	assert_newCode64CharfromString(actual, 0, t)

	actual = newCode64CharfromString("A")
	assert_newCode64CharfromString(actual, 11, t)

	actual = newCode64CharfromString("F")
	assert_newCode64CharfromString(actual, 21, t)

	actual = newCode64CharfromString("K")
	assert_newCode64CharfromString(actual, 31, t)

	actual = newCode64CharfromString("P")
	assert_newCode64CharfromString(actual, 41, t)

	actual = newCode64CharfromString("U")
	assert_newCode64CharfromString(actual, 51, t)

	actual = newCode64CharfromString("Z")
	assert_newCode64CharfromString(actual, 61, t)

	actual = newCode64CharfromString("-")
	assert_newCode64CharfromString(actual, 63, t)

}

func Test_ToString(t *testing.T) {
	c := NewCode64FromInt(64*64*11 + 64*21 + 31) //KFA

	actual := c.ToString()
	expected := "KFA"
	if actual != expected {
		t.Errorf("期待値と異なる値に変換されました。expected=%s, actual=%s", expected, actual)
	}

}

func Test_ToInt(t *testing.T) {
	c := NewCode64FromInt(64*64*64*64*64*5 + 64*64*3 + 64*2 + 9)
	actual := c.ToInt()
	expected := 64*64*64*64*64*5 + 64*64*3 + 64*2 + 9

	if actual != expected {
		t.Errorf("期待値と異なる値に変換されました。expected=%d, actual=%d", expected, actual)
	}

}

func Test_Length(t *testing.T) {
	// テストケース
	testCase := func(str string, expected int) {
		c, _ := NewCode64FromString(str)
		if c.Length() != expected {
			t.Errorf("期待値と異なるlengthが検出されました。testString=%s, length=%d", str, c.Length())
		}
	}

	// １〜１０００桁までの文字列をテスト
	strList := make([]byte, 0)
	for i := 0; i < 1000; i++ {
		strList = append(strList, []byte(string(CHARTABLE[i%64]))...)
		str := string(strList)
		testCase(str, i+1)
	}
}

func assert_newCode64CharfromInt(actual *code64char, expected string, t *testing.T) {
	if actual.Code != expected {
		t.Errorf("期待値と異なる値に変換されました。expected=%d, actual=%d", expected, actual.Code)
	}
}

func assert_newCode64CharfromString(actual *code64char, expected int, t *testing.T) {
	if actual.Num != expected {
		t.Errorf("期待値と異なる値に変換されました。", actual.Num, expected)
	}
}
