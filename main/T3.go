package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
	"bytes"
	"io"
)

func WriteTo()  {
	reader := bytes.NewReader([]byte("Go语言中文网"))
	reader.WriteTo(os.Stdout)
}

func ReadFrom()  {
	f, err := os.Open("t")
	if nil != err {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(os.Stdout)
	w.ReadFrom(f)
	w.Flush()
}

func WriterT()  {
	f, err := os.Create("t")
	if nil != err {
		panic(err)
	}
	defer f.Close()
	f.WriteString("Golang中文社区——这里是多余的")
	n, err := f.WriteAt([]byte("Go语言中文网"), 24)
	if nil != err {
		panic(err)
	}
	fmt.Println(n)
}


func ReaderT()  {
	reader := strings.NewReader("Go语言中文网")
	b := make([]byte, 100)
	n, err := reader.ReadAt(b, 2)
	if n > 0 {
		fmt.Printf("%s, %d\n", b, n)
		fmt.Println(err)
	}
}

func Seek()  {
	r := strings.NewReader("Go语言中文网")
	r.Seek(-6, os.SEEK_END)
	rs, _, _ := r.ReadRune()
	fmt.Printf("%c\n", rs)
}

func ByteRWT()  {
	var ch byte
	fmt.Scanf("%c\n", &ch)

	buffer := new(bytes.Buffer)
	err := buffer.WriteByte(ch)
	if err == nil {
		fmt.Println("写入一个字节成功！准备读取该字节……")
		newCh, _ := buffer.ReadByte()
		fmt.Printf("读取的字节：%c\n", newCh)
	} else {
		fmt.Println("写入错误")
	}
}

func main()  {

	//WriterT()
	//ReadFrom()
	//WriteTo()
	//Seek()
	//ByteRWT()
	io.Copy(os.Stdout, strings.NewReader("HHKSDF"))
}

