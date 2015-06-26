package routes

import (
	"fmt"
	kl "github.com/zaramme/KifCloud-Logic/kifLoader"
	rsh "github.com/zaramme/KifCloud-Logic/rsh"
	"os"
	"testing"
)

func Test_newRoutesFromKifFile(t *testing.T) {
	file, err_of := testUtil_openFileStream("../kifLoader/testFiles/sampleKif.kif", t)
	if err_of != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err_of.Error())
	}
	kif, err := kl.LoadKifFile(file)

	if err != nil {
		t.Errorf("棋譜ファイルの読み込みに失敗しました。 error = %s", err.Error())
		return
	}
	routes, rErr := NewRoutesFromKifuFile(kif)
	if rErr != nil {
		t.Errorf("棋譜の変換に失敗しました。 error = %s", rErr.Error())
		return
	}
	for _, set := range routes {
		if set.Current == "" {
			continue
		}
		rshCurrent, err := rsh.NewRshCodeFromString(set.Current)
		if err != nil {
			t.Errorf("RSH変換に失敗しました(生成されたRshCurrentが不正です)。 error = %s", err.Error())
			return
		}
		boardCurrent := rsh.BuildBoardFromRshCode(rshCurrent)

		isValid, errString := boardCurrent.IsValid()
		if !isValid {
			t.Errorf("BOARD変換に失敗しました(生成されたBOARDの整合性がおかしいです)。 error = %s", errString)
		}
	}
}

func Test_newRoutesFromKifFile_特定ケース１(t *testing.T) {
	file, err_of := testUtil_openFileStream("get_003.kif", t)
	if err_of != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err_of.Error())
	}
	kif, err := kl.LoadKifFile(file)

	if err != nil {
		t.Errorf("棋譜ファイルの読み込みに失敗しました。 error = %s", err.Error())
		return
	}
	routes, rErr := NewRoutesFromKifuFile(kif)
	if rErr != nil {
		t.Errorf("棋譜の変換に失敗しました。 error = %s", rErr.Error())
		return
	}
	for _, set := range routes {
		if set.Current == "" {
			continue
		}
		rshCurrent, err := rsh.NewRshCodeFromString(set.Current)
		if err != nil {
			t.Errorf("RSH変換に失敗しました(生成されたRshCurrentが不正です)。 error = %s", err.Error())
			return
		}
		boardCurrent := rsh.BuildBoardFromRshCode(rshCurrent)

		isValid, errString := boardCurrent.IsValid()
		if !isValid {
			t.Errorf("BOARD変換に失敗しました(生成されたBOARDの整合性がおかしいです)。 error = %s", errString)
		}
	}
}

func Test_newRoutesFromKifFile_HotFix2(t *testing.T) {
	file, err := testUtil_openFileStream("get_003.kif", t)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
	}
	kif, err := kl.LoadKifFile(file)

	if err != nil {
		t.Errorf("棋譜ファイルの読み込みに失敗しました。 error = %s", err.Error())
		return
	}
	routes, rErr := NewRoutesFromKifuFile(kif)
	if rErr != nil {
		t.Errorf("棋譜の変換に失敗しました。 error = %s", rErr.Error())
		return
	}

	id := 0

	asrtRsh := func(rshString string) error {
		id++
		r, err := rsh.NewRshCodeFromString(rshString)
		if err != nil {
			t.Errorf("rshの変換に失敗しました。rsh = %s error = %s \n", rshString, err.Error())
			return fmt.Errorf("%s", rshString)
		}
		b := rsh.BuildBoardFromRshCode(r)
		for key, value := range b.CapturedMap {
			if value < 0 {
				t.Errorf("[%s : %d]持ち駒の値が不正です。[%s] => %d \n", rshString, id, key.Output(), value)
				return fmt.Errorf("%s", rshString)
			}
		}
		return nil
	}

	for i := 0; i < len(routes)-1; i++ {
		value := routes[i]
		err = asrtRsh(value.Prev)
		if err != nil {
			t.Errorf("rshの変換に失敗しました。seq = %d rsh= %s error = %s", i, value.Prev, err.Error())
			break
		}
	}
}

func Test_newRoutesFromKifFile_HotFix3(t *testing.T) {
	file, err := testUtil_openFileStream("test_HasRepeatMove.kif", t)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
	}
	kif, err := kl.LoadKifFile(file)

	if err != nil {
		t.Errorf("棋譜ファイルの読み込みに失敗しました。 error = %s", err.Error())
		return
	}
	routes, rErr := NewRoutesFromKifuFile(kif)
	if rErr != nil {
		t.Errorf("棋譜の変換に失敗しました。 error = %s", rErr.Error())
		return
	}

	id := 0

	asrtRsh := func(rshString string) error {
		id++
		r, err := rsh.NewRshCodeFromString(rshString)
		if err != nil {
			t.Errorf("rshの変換に失敗しました。rsh = %s error = %s \n", rshString, err.Error())
			return fmt.Errorf("%s", rshString)
		}
		b := rsh.BuildBoardFromRshCode(r)
		for key, value := range b.CapturedMap {
			if value < 0 {
				t.Errorf("[%s : %d]持ち駒の値が不正です。[%s] => %d \n", rshString, id, key.Output(), value)
				return fmt.Errorf("%s", rshString)
			}
		}
		return nil
	}

	for i := 0; i < len(routes); i++ {
		value := routes[i]
		err = asrtRsh(value.Prev)
		if err != nil {
			t.Errorf("rshの変換に失敗しました。seq = %d rsh= %s error = %s", i, value.Prev, err.Error())
			break
		}
	}
}

func testUtil_openFileStream(filePath string, t *testing.T) (file *os.File, err error) {
	file, err = os.Open(filePath)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました error = %s\n", err.Error())
		return nil, err
	}
	return file, nil
}
