package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"../bubblesort"
	"../qsort"
)

var infile *string = flag.String("i", "infile", "File contains value for sorting")
var outfile *string = flag.String("o", "outfile", "File to recieve sort value")
var algorithm *string = flag.String("a", "qsort", "Sort Algorithm")

func main()  {
	flag.Parse()

	if nil != infile{
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile)
	if nil != err{
		fmt.Println(err)
	}else {
		fmt.Println(values)
		writeValues(values, "data_result")
	}

	fmt.Println(bubblesort.BubbleSort([]int{1,5,2,4,3}))
	data := []int{1, 4, 2, 60, 12}
	qsort.Qsort(data, 0, len(data)-1)
	fmt.Println(data)
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if nil != err{
		fmt.Println("fail to open faile ", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for  {
		line, isPrefix, err1 := br.ReadLine()
		if nil != err1{
			if err1 != io.EOF{
				err = err1
			}
			break
		}
		if isPrefix{
			fmt.Println("A too long line, seems unexpected")
			break
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if nil != err1{
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) (err error) {
	file, err := os.Create(outfile)
	if nil != err{
		fmt.Println("can not create file", outfile)
		return
	}
	defer file.Close()

	for _, value := range values{
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}



