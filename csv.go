package golangML

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
)

func ReadCsv(filename string) [][]float64 {
	file, err1 := os.Open(filename)
	if err1 != nil {
		panic(err1)
	}
	defer file.Close()
	rdr := csv.NewReader(bufio.NewReader(file))
	rows, _ := rdr.ReadAll()
	x := make([][]float64, len(rows[0]))
	for i, row := range rows {
		for j := range row {
			num, err2 := strconv.ParseFloat(rows[i][j], 64)
			if err2 != nil {
				panic(err2)
			}
			x[j] = append(x[j], num)
		}
	}
	return x
}

func WriteCsv() {

}
