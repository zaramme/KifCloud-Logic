package kifLoader

import (
	"bufio"
	"fmt"
	m "github.com/zaramme/KifCloud-Logic/move"
	"os"
	"strconv"
	"testing"
)

const debug_loadfile = false
const debug_readFile = false

const debug_sepalate = false
const debug_hasSeparator = false
const debug_mappingInfo = false
const debug_mappingMove = false

const filePath_normal = "testFiles/sampleKif.kif"
const filePath_hasRepeatMove = "testFiles/test_HasRepeatMove.kif"

const filePath_real_prefix = "testFiles/real/get_"
const filePath_real_extention = ".kif"
const filePath_shiftJisString = "testFiles/test_Shift−JIS.txt"

const filePath_error_void = "testFiles/test_void.kif"

const filepath_utf8 = "testFiles/test_utf8.kif"
const filePath_error_noSeparate1 = "testFiles/test_noSeparate_1.kif"
const filePath_error_noSeparate2 = "testFiles/test_noSeparate_2.kif"

const filePath_error_invalidMoves1 = "testFiles/test_invalidMoves_1.kif"
const filePath_error_invalidMoves2 = "testFiles/test_invalidMoves_2.kif"
const filePath_error_invalidMoves3 = "testFiles/test_invalidMoves_3.kif"
const filePath_error_invalidMoves4 = "testFiles/test_invalidMoves_4.kif"
const filePath_error_invalidMoves5 = "testFiles/test_invalidMoves_5.kif"

const filePath_error_invalidInfo1 = "testFiles/test_invalidInfo_1.kif"
const filePath_error_invalidInfo2 = "testFiles/test_invalidInfo_2.kif"

//////////////////////////////////////////////////////////////////////
// 正常系
func Test_LoadKifFile(t *testing.T) {

	file, err_of := testUtil_openFileStream(filePath_normal, t)
	if err_of != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err_of.Error())
	}
	KifFile, err_lk := LoadKifFile(file)
	if err_lk != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err_lk.Error())
		return
	}

	if debug_loadfile == false {
		return
	}

	fmt.Println(KifFile.Info)

	for n, move := range KifFile.Moves {
		fmt.Printf("[ %d ]", n)
		fmt.Printf(move.ToJpnCode())
		fmt.Printf("\n")
	}
}
func Test_LoadKifFile_realCases(t *testing.T) {
	testCase := func(filePath string) {
		file, err := testUtil_openFileStream(filePath, t)
		if err != nil {
			return
		}
		_, err = LoadKifFile(file)
		if err != nil {
			t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
			return
		}
		//fmt.Printf("%s is valid file \n", filePath)
	}
	for i := 1; i <= 310; i++ {
		id := []byte(strconv.Itoa(i))
		if i < 10 {
			id = append([]byte("0"), id...)
		}
		if i < 100 {
			id = append([]byte("0"), id...)
		}
		filename := append([]byte(filePath_real_prefix), id...)
		filename = append(filename, []byte(filePath_real_extention)...)
		testCase(string(filename))
	}
}
func Test_readFile(t *testing.T) {

	file, err_of := testUtil_openFileStream(filePath_normal, t)
	if err_of != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err_of.Error())
	}
	sList, err_rf := readFile(file)

	if err_rf != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err_rf.Error())
		return
	}

	if debug_readFile {
		for _, value := range sList {
			fmt.Println(value)
		}
	}
}

func Test_hasSeparator(t *testing.T) {
	file, err := testUtil_openFileStream(filePath_normal, t)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
	}
	sList, err := readFile(file)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
	}
	seq, b := hasSeparator(sList)

	if !b {
		t.Error("セパレータの検出に失敗しました")
	} else if debug_hasSeparator {
		fmt.Printf("セパレータ検出に成功しました seq...%d\n", seq)
	}
}

func Test_separate(t *testing.T) {
	sList, err := tetsUtil_readFile(filePath_normal, t)
	if err != nil {
		t.Error("ファイルの読み込みに失敗しました。(error=%s)", err.Error())
	}
	iList, mList, err := separate(sList)

	if err != nil {
		t.Error("棋譜情報と指し手情報の識別にしっぱしいました")
		return
	}

	if !debug_sepalate {
		return
	}
	fmt.Println("[[[情報]]]")
	for _, value := range iList {
		fmt.Println(value)
	}

	fmt.Println("[[[指し手]]]")

	for _, value := range mList {
		fmt.Println(value)
	}

}

func Test_mappingInfo(t *testing.T) {
	file, err := testUtil_openFileStream(filePath_normal, t)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
		return
	}
	sList, err := readFile(file)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
		return
	}
	iList, mList, err := separate(sList)
	if err != nil {
		t.Errorf("セパレータの検出に失敗しました。(error＝%s)", err.Error())
		return
	}

	iMap, err := mappingInfo(iList)

	if err != nil {
		t.Error("棋譜情報のマッピングに失敗しました error = %s", err.Error())
	}
	if len(iMap) < 1 {
		t.Error("棋譜情報のマッピングに失敗しました")
		return
	}

	if !debug_mappingInfo {
		return
	}
	for key, value := range iMap {
		fmt.Printf("[(key)%s] = %s\n", key, value)
	}
	moveList, err := mappingMoves(mList)
	if err != nil {
		t.Error("指し手情報のマッピングに失敗しました")
		return
	}
	if len(moveList) < 1 {
		t.Error("指し手情報のマッピングに失敗しました")
		return
	}

	if !debug_mappingMove {
		return
	}
	for key, value := range moveList {
		fmt.Printf("[%d] = %s\n", key, value)

	}
}

func Test_mappingMoves(t *testing.T) {
	file, err := testUtil_openFileStream(filePath_normal, t)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
		return
	}
	sList, err := readFile(file)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
		return
	}
	_, mList, err := separate(sList)
	if err != nil {
		t.Errorf("セパレータの検出に失敗しました。(error＝%s)", err.Error())
		return
	}

	//	fmt.Print("[test_mappingMoves]\n")
	//	fmt.Printf("%v\n", mList)

	moveList, err := mappingMoves(mList)
	if err != nil {
		t.Error("指し手情報のマッピングに失敗しました")
		return
	}
	if len(moveList) < 1 {
		t.Error("指し手情報のマッピングに失敗しました")
		return
	}
	if !debug_mappingMove {
		return
	}
	for key, value := range moveList {
		fmt.Printf("[%d] = %s\n", key, value)

	}
}

func Test_lineConvertFromShiftJisToUTF8(t *testing.T) {
	file, err := os.Open(filePath_shiftJisString)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました error = %s\n", err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	sJisText := scanner.Text()

	utf8Text, err2 := lineConvertFromShiftJisToUTF8(sJisText)

	if err2 != nil {
		t.Errorf("UTF−8変換に失敗しました。(システムエラー)")
		return
	}

	if utf8Text != "我輩は猫である。名前はまだ無い。" {
		t.Errorf("UTF−8変換に失敗しました。変換後 = %S", utf8Text)
	}

}

////////////////////////////////////////////////////////////////////
// エラーパターン
func Test_readFile_ファイルがUTF8(t *testing.T) {
	file, err := testUtil_openFileStream(filepath_utf8, t)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました error = %s\n", err.Error())
	}
	_, err = readFile(file)

	if err == nil {
		t.Errorf("エラー検出に失敗しました。filePath = %s", filepath_utf8)
	}
}

func Test_separate_エラーパターン(t *testing.T) {
	testCase := func(filePath string) {
		sList, err := tetsUtil_readFile(filePath, t)
		if err != nil {
			t.Errorf("ファイル読み込みに失敗しました error = %s\n", err.Error())
		}
		_, _, err = separate(sList)

		if err == nil {
			t.Errorf("エラー検出に失敗しました。filePath = %s", filePath)
		}
	}
	testCase(filePath_error_noSeparate1)
	testCase(filePath_error_noSeparate2)
	testCase(filePath_error_void)
}

func Test_mappingInfo_エラーパターン(t *testing.T) {
	testCase := func(filePath string) {
		sList, err := tetsUtil_readFile(filePath, t)
		if err != nil {
			t.Errorf("ファイル読み込みに失敗しました error = %s\n", err.Error())
			return
		}
		iList, _, err := separate(sList)
		if err != nil {
			t.Errorf("セパレータの検出に失敗しました。filePath = %s", filePath)
			return
		}
		_, err = mappingInfo(iList)

		if err == nil {
			t.Errorf("エラー検出に失敗しました。filePath = %s", filePath)
			return
		}
		//fmt.Printf("エラー検出：error = %s", err.Error())
	}
	testCase(filePath_error_invalidInfo1)
	testCase(filePath_error_invalidInfo2)
}

func Test_mappingMoves_エラーパターン(t *testing.T) {
	testCase := func(filePath string) {
		sList, err := tetsUtil_readFile(filePath, t)
		if err != nil {
			t.Errorf("ファイル読み込みに失敗しました error = %s\n", err.Error())
			return
		}
		_, mList, err := separate(sList)
		if err != nil {
			t.Errorf("セパレータの検出に失敗しました。filePath = %s", filePath)
			return
		}
		_, err = mappingMoves(mList)

		if err == nil {
			t.Errorf("エラー検出に失敗しました。filePath = %s", filePath)
			return
		}
		//fmt.Printf("エラー検出：error = %s", err.Error())
	}
	testCase(filePath_error_invalidMoves1)
	testCase(filePath_error_invalidMoves2)
	testCase(filePath_error_invalidMoves3)
	testCase(filePath_error_invalidMoves4)
	testCase(filePath_error_invalidMoves5)
}

func testUtil_openFileStream(filePath string, t *testing.T) (file *os.File, err error) {
	file, err = os.Open(filePath)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました error = %s\n", err.Error())
		return nil, err
	}
	return file, nil
}

func tetsUtil_readFile(filePath string, t *testing.T) (sList []string, err error) {
	file, err := testUtil_openFileStream(filePath, t)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error=%s)", err.Error())
		return nil, err
	}

	sList, err = readFile(file)
	if err != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err.Error())
		return nil, err
	}
	return sList, nil
}

func Test_hasRepeatMove(t *testing.T) {
	file, err_of := testUtil_openFileStream(filePath_hasRepeatMove, t)
	if err_of != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err_of.Error())
	}
	KifFile, err_lk := LoadKifFile(file)
	if err_lk != nil {
		t.Errorf("ファイル読み込みに失敗しました。(error＝%s)", err_lk.Error())
		return
	}

	fmt.Println(KifFile.Info)

	for n, move := range KifFile.Moves {
		fmt.Printf("[ %d ]\n", n)
		switch desc := move.(type) {
		case *m.Move:
			fmt.Println("type:move[%v]", desc)
		case *m.EndGame:
			fmt.Println("type:EndGame[%v]", desc)
		case *m.RepeatMove:
			println("type:RepeatMove[%v]", desc)
		}

		fmt.Printf("\n")
	}
}
