package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"time"
)

func walkDir(dir string, fileSizes chan<- int64)  {
	for _, entry := range dirents(dir){
		if entry.IsDir(){
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		}else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if nil != err{
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64)  {
	fmt.Fprintf(os.Stderr, "%d files  %.3f GB\n", nfiles, float64(nbytes)/1e9)
}

func timeUsed_1() func() {
	start := time.Now()
	f := func() {
		fmt.Println(time.Since(start))
	}
	return f
}

func Test_1()  {
	defer timeUsed_1()()
	pathSlice := []string{"/home/gaoyangang", "/usr/bin"}
	fileSize := make(chan int64)
	go func() {
		defer close(fileSize)
		for _, root := range pathSlice{
			walkDir(root, fileSize)
		}
	}()

	tick := time.Tick(50 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for  {
		select {
		case size, ok := <-fileSize:
			if !ok{
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func main() {
	Test_1()
}