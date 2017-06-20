package main

import (
	"path/filepath"
	"time"
	"sync"
	"os"
	"io/ioutil"
	"fmt"
)

var sema = make(chan struct{}, 20)

func dirents_1(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {<-sema}()
	entries, err := ioutil.ReadDir(dir)
	if nil != err{
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func walkDir_1(dir string, n *sync.WaitGroup, fileSizes chan<- int64)  {
	defer n.Done()
	for _, entry := range dirents_1(dir){
		if entry.IsDir(){
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir_1(subdir, n, fileSizes)
		}else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage_1(nfiles, nbytes int64)  {
	fmt.Fprintf(os.Stderr, "%d files  %.3f GB\n", nfiles, float64(nbytes)/1e9)
}

func timeUsed() func() {
	start := time.Now()
	f := func() {
		fmt.Println(time.Since(start))
	}
	return f
}

func Test()  {
	defer timeUsed()()
	pathSlice := []string{"/home/gaoyangang", "/usr/bin"}
	fileSize := make(chan int64)
	var n sync.WaitGroup
	for _, root := range pathSlice{
		n.Add(1)
		go walkDir_1(root, &n, fileSize)
	}

	go func() {
		n.Wait()
		close(fileSize)
	}()

	tick := time.Tick(10 * time.Millisecond)
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
			printDiskUsage_1(nfiles, nbytes)
		}
	}
	printDiskUsage_1(nfiles, nbytes)
}


func main() {
	Test()
}