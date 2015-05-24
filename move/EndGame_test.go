package move

import (
	def "KifuLibrary-Logic/define"
	"testing"
)

func Test_en_ToJpnCode(t *testing.T) {
	asrt := func(testName string, input *EndGame, expect string) {
		actual := input.ToJpnCode()
		if expect != actual {
			t.Errorf("[%s]テスト失敗 期待値 = %s, 結果値 = %s", testName, expect, actual)
		}
	}

	var input *EndGame

	input = new(EndGame)
	input.Winner = def.BLACK
	asrt("1:先手勝ち", input, "投了")

	input = new(EndGame)
	input.Winner = def.WHITE
	asrt("2:後手勝ち", input, "投了")

	input = new(EndGame)
	input.IsDrawGame = true
	asrt("3:引き分け", input, "引き分け")

	input = new(EndGame)
	input.IsSuspended = true
	asrt("4:中断", input, "中断")

	input = new(EndGame)
	input.Winner = def.BLACK
	input.ByOutOfTime = true
	asrt("5:時間切れ負け", input, "時間切れ負け")

	input = new(EndGame)
	input.Winner = def.WHITE
	input.ByFoul = true
	asrt("6:反則負け", input, "反則負け")

}
