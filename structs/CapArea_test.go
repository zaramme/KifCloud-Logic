package structs

import (
	def "github.com/zaramme/KifCloud-Logic/define"
	"testing"
)

func Test_caparea_construct(t *testing.T) {
	var ca CapArea = CapArea{def.BLACK, def.OH}

	if ca.KindOfPiece != def.OH {
		t.Errorf("CapAreaの初期化に失敗しました", ca.KindOfPiece)
	}

	if ca.Player != def.BLACK {
		t.Errorf("CapAreaの初期化に失敗しました", ca.Player)
	}

	ca = CapArea{def.WHITE, def.FU}
	if ca.KindOfPiece != def.FU {
		t.Errorf("CapAreaの初期化に失敗しました", ca.KindOfPiece)
	}

	if ca.Player != def.WHITE {
		t.Errorf("CapAreaの初期化に失敗しました", ca.Player)
	}
}

func Test_caparea_output(t *testing.T) {

	var expected string
	var actual string
	var cap CapArea

	cap = CapArea{def.BLACK, def.FU}
	expected = "{BLACK,FU}"
	actual = cap.Output()
	if expected != actual {
		t.Errorf("Caparea.Outputが期待値と異なります。 期待値＝%s, 実測値＝%s", expected, actual)
	}

	cap = CapArea{def.WHITE, def.KAKU}
	expected = "{WHITE,KAKU}"
	actual = cap.Output()
	if expected != actual {
		t.Errorf("Caparea.Outputが期待値と異なります。 期待値＝%s, 実測値＝%s", expected, actual)
	}

}
