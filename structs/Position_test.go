package structs

import (
	"testing"
)

func Test_position_construct(t *testing.T) {
	var actual Position = Position{1, 2}
	expectedX := 1
	expectedY := 2
	if actual.X != expectedX {
		t.Errorf("x座標の代入に失敗しました")
	}
	if actual.Y != expectedY {
		t.Errorf("x座標の代入に失敗しました")
	}
}

func Test_position_isCaptured(t *testing.T) {
	var actual Position = Position{0, 0}
	if !actual.IsCaptured() {
		t.Errorf("駒台判定に失敗しました(x,y)=(0,0)")
	}

	actual = Position{0, 1}
	if actual.IsCaptured() {
		t.Errorf("駒台判定に失敗しました(x,y)=(0,1)")
	}

	actual = Position{1, 0}
	if actual.IsCaptured() {
		t.Errorf("駒台判定に失敗しました(x,y)=(1,0)")
	}

	actual = Position{1, 1}
	if actual.IsCaptured() {
		t.Errorf("駒台判定に失敗しました(x,y)=(1,1)")
	}
}

func Test_position_compare(t *testing.T) {
	var p1 Position = Position{1, 2}
	var p2 Position = Position{1, 2}

	if p1 != p2 {
		t.Errorf("point構造体の同値判定に失敗しました")
	}
}

func Test_position_output(t *testing.T) {

	var expected, actual string
	var pos Position
	expected = "{3,7}"
	pos = Position{3, 7}
	actual = pos.Output()

	if expected != actual {
		t.Errorf("Position.Outputに失敗しました。期待値＝%s、実測値＝%s", expected, actual)
	}

	expected = "{Captured}"
	pos = Position{0, 0}
	actual = pos.Output()

	if expected != actual {
		t.Errorf("Position.Outputに失敗しました。期待値＝%s、実測値＝%s", expected, actual)
	}

}

func Test_position_outputJpnCode(t *testing.T) {

	var pos Position
	pos = Position{1, 1}
	if pos.OutputJpnCode() != "１一" {
		t.Errorf("座標の変換に失敗しました。%s", pos.Output())
	}

	pos = Position{5, 3}
	if pos.OutputJpnCode() != "５三" {
		t.Errorf("座標の変換に失敗しました。%s", pos.Output())
	}

	pos = Position{9, 9}
	if pos.OutputJpnCode() != "９九" {
		t.Errorf("座標の変換に失敗しました。%s", pos.Output())
	}
}
func Test_position_outputJSCode(t *testing.T) {
	asrt := func(pos Position, exp string) {
		act := pos.OutputJsCode()
		if act != exp {
			t.Errorf("座標の変換に失敗しました。%s", pos.Output())
		}
	}

	asrt(Position{1, 1}, "11")
	asrt(Position{2, 2}, "22")
	asrt(Position{3, 4}, "34")
	asrt(Position{4, 6}, "46")
	asrt(Position{5, 8}, "58")
	asrt(Position{6, 1}, "61")
	asrt(Position{7, 3}, "73")
	asrt(Position{8, 5}, "85")
	asrt(Position{9, 9}, "99")
}
