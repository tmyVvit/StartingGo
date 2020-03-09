package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)
var startTime time.Time

func Init() {
	startTime = time.Now()
}
func ArraySource(a ...int)  <- chan int {
	out := make(chan int)
	go func() {
		for _, value := range a {
			out <- value
		}
		close(out)
	}()
	return out
}

func InMemSort(in <- chan int) <- chan int {
	out := make(chan int, 1024)
	go func() {
		var a []int
		for v := range in  {
			a = append(a, v)
		}
		sort.Ints(a)
		fmt.Println("sort done:", time.Now().Sub(startTime))

		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <- chan int) <- chan int {
	out := make(chan int, 1024)
	go func() {
		v1, ok1 := <- in1
		v2, ok2 := <- in2
		for ok1 || ok2 {
			if !ok1 || (ok2 && v2 <= v1) {
				out <- v2
				v2, ok2 = <- in2
			} else {
				out <- v1
				v1, ok1 = <- in1
			}
		}
		close(out)
		fmt.Println("merge done:", time.Now().Sub(startTime))
	}()
	return out
}

func MergeN(inputs ...<- chan int) <- chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}
func ReaderSource(reader io.Reader, chunkSize int) <- chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		i := 0
		for  {
			n, err := reader.Read(buffer)
			if n > 0 {
				out <- int(binary.BigEndian.Uint64(buffer))
			}
			i++
			if err != nil || (chunkSize > -1 && i*8 >= chunkSize) {
				break
			}
		}
		close(out)
		fmt.Println("read done", time.Now().Sub(startTime))
	}()
	return out
}

func WriterSink(writer io.Writer, in <- chan int) {
	buffer := make([]byte, 8, 1024)
	for v := range in {
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <- chan int {
	out := make(chan int, 1024)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

