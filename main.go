package main

import (
	"fmt"
	"os"
	"time"

	"github.com/chrisbcaldwell/trim"

	"github.com/go-gota/gota/dataframe"
)

const filepath = "random.csv"
const quantile = 0.05

func main() {
	start := time.Now()
	df, err := readToDataFrame(filepath)
	if err != nil {
		fmt.Println("could not read", filepath)
	}
	df.SetNames("integers", "floats")

	ints, err := df.Col("integers").Int()
	if err != nil {
		fmt.Println("Column 'integers' not of integer type")
	}
	floats := df.Col("floats").Float()

	intTM := trim.TrimmedMean(ints, quantile)
	floatTM := trim.TrimmedMean(floats, quantile)
	elapsed := time.Since(start)
	fmt.Printf("Integer data trimmed mean: %.2f\n", intTM)
	fmt.Printf("Float data trimmed mean: %.2f\n", floatTM)
	fmt.Println("Total run time:", elapsed)
}

func readToDataFrame(p string) (dataframe.DataFrame, error) {
	f, err := os.Open(p)
	if err != nil {
		fmt.Println("Unable to read input file " + filepath)
		return dataframe.New(), err
	}
	defer f.Close()

	df := dataframe.ReadCSV(f)
	return df, nil
}
