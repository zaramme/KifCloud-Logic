package kifLoader

import (
	"bufio"
	"errors"
	m "github.com/zaramme/KifCloud-Logic/move"
	"os"
	"strconv"
	s "strings"
	//	t "time"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
)

const SEPALATOR = "手数----指手---------消費時間--"
const INFO_SEPALATOR = "："
const COMMENT_PREFIX = "*"

type KifFile struct {
	Info  map[string]string
	Moves moveList
}
type moveList []m.PlayingDescriptor

////////////////////////////////////////////////////////////
// Public Method
func LoadKifFile(file *os.File) (kif *KifFile, err error) {
	///////////////////////////////////////////////////////
	// ファイルを行ごとに格納
	sList, err := readFile(file)
	if err != nil {
		return nil, err
	}

	kif, err = LoadStringList(sList)

	if err != nil {
		return nil, err
	}

	return kif, nil
}

func LoadStringList(sList []string) (kif *KifFile, err error) {
	///////////////////////////////////////////////////////
	// セパレータを検出して対局情報と指し手情報に分割
	infoList, moveList, err := separate(sList)
	if err != nil {
		return nil, err
	}

	///////////////////////////////////////////////////////
	// 対局情報の検出
	info, err := mappingInfo(infoList)
	if err != nil {
		return nil, err
	}
	///////////////////////////////////////////////////////
	// 指し手情報を検出
	moves, err := mappingMoves(moveList)
	if err != nil {
		return nil, err
	}
	///////////////////////////////////////////////////////
	// オブジェクトを生成
	kif = newKifFile(info, moves)

	return kif, nil
}

////////////////////////////////////////////////////////////
// Private Method
func newKifFile(info map[string]string, moves moveList) *KifFile {
	obj := new(KifFile)
	obj.Info = info
	obj.Moves = moves

	return obj
}

func readFile(file *os.File) (sList []string, err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sOriginal := scanner.Text()
		sEncoded, err := lineConvertFromShiftJisToUTF8(sOriginal)
		if err != nil {
			return nil, err
		}
		sList = append(sList, sEncoded)
	}
	if err := scanner.Err(); err != nil {
		return sList, err
	}

	return sList, nil
}

// 棋譜情報を対局情報と指し手情報に分割
func separate(input []string) (infoList, moveList []string, err error) {

	infoList = make([]string, 0)
	moveList = make([]string, 0)
	seq, b := hasSeparator(input)

	if !b {
		return infoList, moveList, errors.New("ファイル形式が不正です：指し手が見つかりません")
	}

	infoList = input[:seq]
	moveList = input[seq+1:]

	return infoList, moveList, nil
}

// 対局情報のマッピング
func mappingInfo(iList []string) (iMap map[string]string, err error) {
	iMap = make(map[string]string)

	for _, line := range iList {
		separated := s.Split(line, INFO_SEPALATOR)
		if len(separated) != 2 {
			return nil, fmt.Errorf("対局情報のマッピングに失敗しました。s = %s", line)
		}
		iMap[separated[0]] = separated[1]
	}
	return iMap, nil
}

// 指し手情報のマッピング
func mappingMoves(mList []string) (rList moveList, err error) {

	rList = make(moveList, 0)
	for n := 0; n < len(mList); n++ {
		line := mList[n]
		// コメント行は無視
		if s.HasPrefix(line, COMMENT_PREFIX) {
			fmt.Printf("[n = %d]コメント行を無視します", n)
			continue
		}
		// 空行で分割
		fields := s.Fields(line)
		if len(fields) < 2 {
			err = fmt.Errorf("[n=%d]スプリットに失敗しました。= (%s)\n", n, line)
			return nil, err
		}
		// 手数を検出
		moveNum, err := strconv.Atoi(fields[0])
		if err != nil {
			err = fmt.Errorf("[n=%d]手数の検出に失敗しました。= (%s)\n", n, line)
			return nil, err
		}
		if len(rList) < moveNum {
			rList = extendSlice(rList, moveNum)
		}
		rList[moveNum-1], err = convertKifCodeToMoveCode(fields[1], moveNum)
		if err != nil {
			err = fmt.Errorf("[n=%d]記号の変換に失敗しました。。= (%s)\n", n, line)
			return nil, err
		}
	}

	return rList, err
}

func extendSlice(slice moveList, length int) moveList {
	res := make(moveList, length)

	for i, v := range slice {
		res[i] = v
	}
	return res
}

func hasSeparator(input []string) (seq int, hasSepalator bool) {
	// セパレーターを検出する
	hasSepalator = false

	for n, line := range input {
		if line == SEPALATOR {
			seq = n
			hasSepalator = true
			break
		}
	}
	return seq, hasSepalator
}

func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}

// Convert a string encoding from ShiftJIS to UTF-8
func lineConvertFromShiftJisToUTF8(str string) (string, error) {
	return transformEncoding(s.NewReader(str), japanese.ShiftJIS.NewDecoder())
}
