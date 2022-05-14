package golib

import (
	"bufio"
	"encoding/csv"
	"os"

	iconv "github.com/djimenez/iconv-go" //콘솔에서 한글사용
)

/* --------------------------------------------------
    CVS Functions
---------------------------------------------------- */

func Read_CSV_File(filename string) ([][]string, error) { // 라이스로 반환
	file, err := os.Open(filename)              // 파일 오픈
	rdr := csv.NewReader(bufio.NewReader(file)) // csv reader 생성
	rows, _ := rdr.ReadAll()                    // csv 내용 모두 읽기
	return rows, err
}

/* ----------------------------------------------------
    System Functions
---------------------------------------------------- */

func Hangul(str string) string { //콘솔에 출력시 한글깨짐 방지를 위해 utf-8로 변환하여 반환함
	out, _ := iconv.ConvertString(string(str), "euc-kr", "utf-8")
	return out
}

func setFyneHangul() {
	os.Setenv("FYNE_FONT", "NanumGothic.ttf") // Fyne에서 한글 깨짐을 방지 한다.
}
