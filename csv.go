package golangML

import (
	"bufio"
	"encoding/csv"
	"errors"
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
	if len(rows[0]) == 0 {
		panic(errors.New("No data"))
	}
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

// func WriteCsv(filename string, data ...[]float64) {
// 	filename = "./" + filename
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	wr := csv.NewWriter(bufio.NewWriter(file))
// 	for _, di := range data {
// 		writeData := make([]string, len(di))
// 		for i, d := range di {
// 			writeData[i] = strconv.FormatFloat(d, 'f', -1, 32)
// 		}
// 		wr.Write(writeData)
// 	}
// 	wr.Flush()
// }
