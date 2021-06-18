package main

import (
	"bytes"
	"fmt"
	"main/types"
	"net"
	"os"
	"strconv"
	"sync"
)


func scanPortTcp(address string, port int, semaphore *sync.WaitGroup, result *types.ListPorts) {

	var buffer = bytes.Buffer{}

	buffer.WriteString(address)
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(port))

	_, err := net.Dial("tcp", buffer.String())
	if err == nil {
		result.AddElement(port, true)
	} else {
		result.AddElement(port, false)
	}

	semaphore.Done()
}


func scanRangePortsTcp(address string, first int, second int) {

	diff := second - first
	i := 0

	if diff < 0 {
		fmt.Fprintln(os.Stderr, "La prima porta deve essere minore della seconda")
		return
	}

	semaphore := &sync.WaitGroup{}
	semaphore.Add(diff + 1)

	result := types.NewPortList()

	for first + i <= second {
		go scanPortTcp(address, first + i, semaphore, result)
		i++
	}

	semaphore.Wait()
	result.PrintResult()
}


func main() {
	scanRangePortsTcp("127.0.0.1", 7950, 8000)
}
