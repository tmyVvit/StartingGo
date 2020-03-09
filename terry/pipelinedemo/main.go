package main

import (
	"bufio"
	"fmt"
	"os"
	"pipeline"
	"strconv"
)

func main() {
	p, files := CreateNetworkPipeline("small.in", 800, 4)
	for i := 0; i < len(files); i++ {
		defer files[i].Close()
	}
	WriteToFile("large.out", p)
	//printFile("large.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	p := pipeline.ReaderSource(file, -1)
	i := 0
	for v := range p {
		fmt.Printf("%d: %d\n", i, v)
		i++
		if i >= 100 {
			break
		}
	}
}

func WriteToFile(filename string, p <- chan int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	pipeline.WriterSink(writer, p)
}

func CreatePipeline(filename string, fileSize int, chunkCount int) (<- chan int, []*os.File){

	var out []<-chan int
	var files []*os.File
	chunkSize := fileSize / chunkCount
	pipeline.Init()
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sorted := pipeline.InMemSort(source)
		out = append(out, sorted)
		files = append(files, file)
	}
	return pipeline.MergeN(out...), files
}

func CreateNetworkPipeline(filename string, fileSize int, chunkCount int) (<- chan int, []*os.File){

	var out []<-chan int
	var files []*os.File
	var addrs []string
	chunkSize := fileSize / chunkCount
	pipeline.Init()
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		addr := ":" + strconv.Itoa(7000 + i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		addrs = append(addrs, addr)
		files = append(files, file)
	}
	for _, addr := range addrs {
		out = append(out, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(out...), files
}

func CreateRandomFile(name string, count int) error {
	file, err := os.Create(name)

	if err != nil {
		return err
	}
	defer file.Close()
	p := pipeline.RandomSource(count)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	_ = writer.Flush()
	return nil
}

func demo() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3, 4, 6, 1, 9, 5)),
		pipeline.InMemSort(pipeline.ArraySource(2, 10, 3, 5, 8)) )
	for v := range p  {
		fmt.Println(v)
	}
}
