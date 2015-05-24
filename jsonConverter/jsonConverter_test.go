package jsonConverter

import (
	//	b "KifuLibrary-Logic/board"
	r "KifuLibrary-Logic/rsh"
	"fmt"
	"testing"
)

//func Test_初期盤面変換(t *testing.T) {
// 	brd := b.NewBoardInit()

// 	json, _ := BoardToJson(brd)
// }

func Test_特定盤面変換(t *testing.T) {

	testCase := func(str string) {
		fmt.Printf("rsh = %s\n", str)
		actualRsh, rshErr := r.NewRshCodeFromString(str)
		if rshErr != nil {
			t.Errorf("rsh変換に失敗しました。= %s\n", str)
			return
		}

		board := r.BuildBoardFromRshCode(actualRsh)
		json, jsonErr := BoardToJson(board)

		if jsonErr != nil {
			t.Errorf("json変換に失敗しました。= %s\n", str)
			return
		}
		//		fmt.Printf("rsh = %s turn = %b\n", str, board.Turn)
		if json.Turn != bool(board.Turn) {
			t.Errorf("Turnの値が不正です。 expect = %s, actual = %s \n", board.Turn, json.Turn)
		}

		if json.Rsh != str {
			t.Errorf("Rshの値が不正です。 expect = %s, actual = %s \n", str, json.Rsh)
		}

		for _, piece := range json.Pieces {
			if piece == nil {
				t.Errorf("pieceの値が不正です。 = %s\n", str)
				t.Errorf("%s", board.CapturedMap)
				t.Errorf("----positionMap = %d\n", len(board.PositionMap))
				var sum int
				for _, i := range board.CapturedMap {
					sum += i
				}
				t.Errorf("----capturedMap = %d\n", sum)
			}
		}
	}
	testCase("4sflk2hGLCEL2UtPMR4e0ohgUgpT0AEsFBnWE0")
	testCase("Icflk2hGLCEL2UtPMR4e0ohgUgpT0AEsFORJf0")
	testCase("4sflk2hWVCEL2UtPMR4e0ohgUgpT0AEsFORJf0")
	testCase("Icflk2hWVCEL2UtPMR4e0ohgUgpT0AEsXpRJf0")
	testCase("4sft-oqWVCEL2UtPMR4e0ohgUgpT0AEsXpRJf0|l008")
	testCase("IcfSi7QWVCEL2DzpMR4e0ohgUgpT0AEsXpRJf0")
	testCase("4sfSi7QWVCELmNzpMR4e0ohgUgpT0AEsXpRJf0")
	testCase("IcfSi7QWVCEL2d_PMR4e0ohgUgpT0AEsXpRJf0")
	testCase("4sfSi7QWVCELiwZPMR4e0ohgUgpT0AEsXpRJf0")
	testCase("OAfSi7QWVCELiwZPMR4e0ohgUgpT0AEsXpRJf0")
	testCase("fQfSi7QWVCELiwZPMR4e0ohgUgpT0AEsXpRJf0")
	testCase("QAfSi7QWVCELiSuyMR4e0ohgUgpT0AEsXpRJf0")
	testCase("fQfSi7Q-VCELiSuyMR4e0ohgUgpT0AEsXpRJf0")
	testCase("QAfSi7Q-VCELiSuyMR4e0ohgUgpT0AEsXrAIi0")
	testCase("fQfSi7Q-VCELiSuyMR4e0ohgUgpTzrFpXrAIi0")
	testCase("QAfSi7Q-VCELiSuyMR4e0ohgUgpTzrFpH7nIi0")
	testCase("iQfSi7Q-VCELiSuyMR4e0ohgUgpTzrFpH7nIi0")
	testCase("TAfSi7Q-VCELiOAvLR4e0ohgUgpTzrFpH7nIi0")
	testCase("kQfSi7Q-VCELiOAvLR4e0ohgUgpTzrFpH7nIi0")
	testCase("VAfSi7Q-RpbKiOAvLR4e0ohgUgpTzrFpH7nIi0")
	testCase("kQfSi7Q-RpbKmOAvLR4e0ohgUgpTzrFpH7nIi0")
	testCase("f8fSi7Q-RpbKmOAvLR4e0ohgUgpTzrFpH7nIi0")
	testCase("ZOfSi7Q-RpbKmOAvLR4e0ohgUgp9RrFpH7nIi0")
	testCase("f8fSi7Q-RpbKmOAvLR4e0ohgUgp9RrFp1GnIi0")
	testCase("ZOfSi7Q-RpbKLOAvLR4e0ohgUgp9RrFp1GnIi0")
	testCase("U4fSi7Q-RpbKLOAvLR4e0ohgUgp9RrFp1GnIi0")
	testCase("jnfSi7Q-RpbKLOAvLR4e0ohgUgpNQrFp1GnIi0")
	testCase("U4fSi7Q-AvAKLOAvLR4e0ohgUgpNQrFp1GnIi0")
	testCase("jnfSi7Q-AvAKLOAvLv4e0ohgUgpNQrFp1GnIi0")
	testCase("U4fSi7Q-AvAKLKnRKv4e0ohgUgpNQrFp1GnIi0")
	testCase("jnfSi7Q-AvAKPKnRKv4e0ohgUgpNQrFp1GnIi0")
	testCase("U4fSi7Q-ahvjPKnRKv4e0ohgUgpNQrFp1GnIi0")
	testCase("jnfSi7Q-ahvjPKnRKv4e0ohgUgp5DfFp1GnIi0")
	testCase("U4fSi7Q-ahvjPKnRKv4e0ohgUgp5DfFp2GnIi0")
	testCase("jnoSi7Q-ahvjPKnRKv4e0ohgUgp5DfFp2GnIi0")
	testCase("U4ofTVq-ahvjPKnRKv4e0ohgUgp5DfFp2GnIi0")
	testCase("jnofTVq-ahvjPKnRKv4e0ohgUgpNwYfp2GnIi0")
	testCase("U4ofTVq-ahvjPHVFmv4e0ohgUgp5u1Ep2GnIi8")
	testCase("jnofTVq-ahvjTXDtqv4e0ohgUgp5u1Ep2GnIi8")
	testCase("U4osEUq-ahvjd-h7Qv4e0ohgUgp5u1Ep2GnIi8")
	testCase("jngvsCm-ahvjd-h7Qv4e0ohgUgp5u1Ep2GnIi8")
	testCase("U4o5f3Q-ahvjd-h7Qv4e0ohgUgp5u1Ep2GnIi8")
	testCase("jnoSi7Q-ahvjd-h7Qv4e0ohgUgpVF-ep2GnIi8")
	testCase("U4oSi7Q-ahvjd-h7Qv4e0ohgUgpVF-epohnIi8")
	testCase("jnoSi7Q-ahvjd-h7QP4e0ohgUgpVF-epohnIi8")
	testCase("U4oSi7Q-ahvjtqh7QP4e0ohgUgpVF-epohnIi8")
	testCase("jnoSi7Q-ahvjtqh7QP4e0ohgUgpQF-epohnIi8")
	testCase("U4gnStq-ahvjtqh7QP4e0ohgUgpQF-epohnIi8")
	testCase("jngnStq-ahvjtqh7QP4e0ohgUgp7f-epudnIi9")
	testCase("U4wXRSq-ahvjtqh7QP4e0ohgUgpAE-epudnIiD|l00l")
	testCase("jnwXRSq-ahvjtqh7Qt4e0ohgUgpAE-epudnIiD|l00t")
	testCase("U4wXRSq-ahvjd-h7QSgm5QhgUgpAE-epudnIiD|l008")
	testCase("jnwXRSq-ahvjd-h7QSgm5QhgUgpM3-epudnIiD|l008")
	testCase("U4wXRSq-ahvjdCrUqSgm5QhgUgpM3-epudnIiD|l008")
	testCase("jnwXRSq-ahvjdCrUqSgm5QhgUgpYU_epeeMIie|l008")
	testCase("U4wXRSqAYz1QdTlvqSgm5QhgUgpYU_epeeMIie|l082")
	testCase("GnwXRSqAYz1Qd6nwqSgm5QhgUgpYU_epeeMIie|l008")
	testCase("s4wXRSqAEXlld6nwqSgm5QhgUgpYU_epeeMIie|l008")
	testCase("GnnXRSqAEXlld6nwqSgm5QhgUgpYU_epeeMIie|l008")
	testCase("s44nStqAEXlld6nwqSgm5QhgUgpYU_epeeMIie|l00l")
	testCase("GnYMStqAEXlld6nwqSgm5QhgUgpYU_epeeMIie|l00l")
	testCase("s4YMStqxuK3Qd6nwqSgm5QhgUgpYU_epeeMIie|l00l")
	testCase("GnKnStqhfH7Qd6nwqSgm5QhgUgpYU_epeeMIie|l00d")
	testCase("s4SqqTqhfH7Qd6nwqSgm5QhgUgpYU_epeeMIie|l00d")
	testCase("GnxqqTqhfH7Qd6nwqSgm5QhgUgpYU_epeeMIie|l00d")
	testCase("s4xqqTqhfH7Qd6nwqSgm5QhgUgpYU_epurYIie|l00d")
	testCase("GnxqqTqhfH7QFXDtqSgm5QhgUgpYU_epurYIie|l00l")
	testCase("s4xqqTqpFnuqFXDtqSgm5QhgUgpYU_epurYIie|l00l")
	testCase("GnAqqTqpFnuqFXDtqSgm5QhgUgpYU_epurYIie|l001")
	testCase("s4AbKuqpFnuqFXDtqSgm5QhgUgpIatepurYIii|l001")
	testCase("jnAbKuqpFnuqFXDtqSgm5QhgUgpIatepurYIii|l001")
	testCase("U4AbKuqpFnuqv1H7QSgm5QhgUgpIatepurYIii|l00l")
	testCase("jnAbKuqpFnuqnII7QSgm5QhgUgpW8tepurYIii|l00d")
	testCase("U4AWItqpFnuqnII7QSgm5QhgUgpO2tepurYIim|l00d")
	testCase("jnAWItqpFnuqRFG7QSgm5QhgUgpO2tepurYIim|l00d")
	testCase("U4IZgTqpFnuqRFG7QSgm5QhgUgp74SepurYIiq|l00l")
	testCase("jnIZgTqd-nwqrFG7QSgm5QhgUgp74SepurYIiq|l00t")
	testCase("U4IZgTqH-nwqv-L7QSgm5QhgUgp74SepurYIiq|l00d")
	testCase("jnIZgTqH-nwqv-L7QSgm5QhgUgpCetepurYIiP|l00d")
	testCase("U4QLJTqH-nwqv-L7QSgm5QhgUgp74SepurYIiT|l00d")
	testCase("jnQLJTqH-nwqv-L7QSgm5QhgUgpWLtepurYIit|l00d")
	testCase("U4ZQ0SqH-nwqv-L7QSgm5QhgUgpWLtepurYIit|l002")
	testCase("jnJoowqH-nwqv-L7QSgm5QGgUgpWLtepurYIit")
	testCase("U4sWNwqH-nwqv-L7QSgm5QGgUgpWLtepurYIit")
	testCase("jnsWNwq1Fnuqv-L7QSgm5QGgUgpWLtepurYIit")
	testCase("U4sWNwq1FnuqvEi7QSgm5QGgUgpWLtepurYIit")
	testCase("jnsWNwq1FnuqfnK4QSgm5QGgUgpWLtepurYIit")
	testCase("U4sWNwqHxH7Qf6f5QSgm5QGgUgpWLtepurYIit|l00l")
	testCase("jnsWNwqHxH7QthG7QSgm5QGgUgpWLtepurYIit")
	testCase("U4sWNwq9noUqthG7QSgm5QGgUgpWLtepurYIit")
	testCase("jnsWNwqH-nwqphG7QSgm5QGgUgpWLtepurYIit")
	testCase("U4sWNwq1HNwqV-L7QSgm5QGgUgpWLtepurYIit")
}
