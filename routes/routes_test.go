package routes

import (
	kl "KifuLibrary-Logic/kifLoader"
	"fmt"
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
	for _, v := range routes {
		//		fmt.Printf("  %s  |  %s \n", v.Move.ToMoveCode(), v.Rsh_Current)
		fmt.Printf("%s \n", v.Current)
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
