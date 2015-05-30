package code

import (
	"testing"
)

func Test_NewCode64FromInt_複数桁をコンストラクタ(t *testing.T) {
	var target, err = NewCode64FromInt(64*64*58 + 64*20 + 33)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}
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

func Test_NewCode64FromInt_桁上がりチェック(t *testing.T) {

	asrt := func(n int, exp string) {
		var target, err = NewCode64FromInt(n)
		if err != nil {
			t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
		}

		if target.ToString() != exp {
			t.Errorf("期待値と出力文字列が異なります。期待値 = %s 実際値 = %s", exp, target.ToString())

		}
	}

	asrt(1, "1")
	asrt(63, "-")
	asrt(64, "10")
	asrt(65, "11")
	asrt(64+63, "1-")
}

func Test_NewCode64FromInt_一桁をコンストラクタ(t *testing.T) {
	var target, err = NewCode64FromInt(30)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}

	if target[0].Num != 30 {
		t.Errorf("期待値と数値が異なります", target[0].Num)
	}
	if target[0].Code != "k" {
		t.Errorf("期待値と数値が異なります", target[0].Code)
	}
}

func Test_NewCode64FromString(t *testing.T) {

	asrt := func(input string, exp string) {
		target, err := NewCode64FromString(input)
		if err != nil {
			t.Errorf("変換エラーを検出しました。 input = %s [Error]%s", input, err.Error())
		}
		if target.ToString() != exp {
			t.Errorf("期待値と出力文字列が異なります。期待値 = %s 実際値 = %s", exp, target.ToString())
		}
	}

	asrt("1", "1")
	asrt("123456", "123456")
	asrt("7890a", "7890a")
	asrt("aAbBcCdD", "aAbBcCdD")
	asrt("EeFfGgHhIiJjKk", "EeFfGgHhIiJjKk")
	asrt("opqrstu", "opqrstu")
	asrt("vwxyz_-", "vwxyz_-")

}

func Test_NewCode64FromString_ゼロプレフィックス(t *testing.T) {
	var target, _ = NewCode64FromString("004")

	if target.ToInt() != 4 {
		t.Errorf("期待値と数値が異なります。%s = %s", target.ToInt(), 4)
	}

}

func Test_Unpadding(t *testing.T) {
	var target, _ = NewCode64FromString("0004")

	actual := target.Unpadding()

	if actual.ToInt() != 4 {
		t.Errorf("期待値と数値が異なります。%d = %d", target.ToInt(), 4)
	}
}

func Test_Unpadding_全て０(t *testing.T) {
	var target, _ = NewCode64FromString("000")

	a := target.Unpadding()

	if a.ToInt() != 0 {
		t.Errorf("期待値と数値が異なります。%s = %s", a.ToInt(), 0)
	}
}

func Test_Padding_Unpadding_逆関数チェック(t *testing.T) {
	// 64 * 64 * 64 - 1 までの数値をパディング〜パディング解除して数値変化がないことを確認

	errCount := 0

	for i := 0; i < 64*64*64; i++ {
		before, err := NewCode64FromInt(i)
		if err != nil {
			t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
		}
		padding := before.Padding(3)
		//		after := padding.Unpadding()

		after := padding.Unpadding()

		if after.ToInt() != i {
			t.Errorf("パディングによる破壊を検出しました。 i = %s(%d), padding = %s, result = %s(%d)", before.ToString(), i, padding.ToString(), after.ToString(), after.ToInt())

			errCount++

			if errCount > 100 {
				t.Errorf("too many errors")
				return
			}
		}

	}
}

func Test_newCode64charfromInt(t *testing.T) {
	var actual *code64char
	var err error
	actual, err = newCode64CharfromInt(0)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}
	assert_newCode64CharfromInt(actual, "0", t)

	actual, err = newCode64CharfromInt(10)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}
	assert_newCode64CharfromInt(actual, "a", t)

	actual, err = newCode64CharfromInt(20)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}
	assert_newCode64CharfromInt(actual, "f", t)

	actual, err = newCode64CharfromInt(30)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}
	assert_newCode64CharfromInt(actual, "k", t)

	actual, err = newCode64CharfromInt(40)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}
	assert_newCode64CharfromInt(actual, "p", t)

	actual, err = newCode64CharfromInt(50)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}
	assert_newCode64CharfromInt(actual, "u", t)

	actual, err = newCode64CharfromInt(60)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}
	assert_newCode64CharfromInt(actual, "z", t)

	actual, err = newCode64CharfromInt(63)
	if err != nil {
		t.Errorf("エラー値を返しました。 = [ %s ]", err.Error())
	}
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
	c, _ := NewCode64FromInt(64*64*11 + 64*21 + 31) //KFA

	actual := c.ToString()
	expected := "AFK"
	if actual != expected {
		t.Errorf("期待値と異なる値に変換されました。expected=%s, actual=%s", expected, actual)
	}

}

func Test_ToInt(t *testing.T) {
	c, _ := NewCode64FromInt(64*64*64*64*64*5 + 64*64*3 + 64*2 + 9)
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
