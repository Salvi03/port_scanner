package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	_ "github.com/akamensky/argparse"
	"main/types"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)


func scanPortTcp(address string, port int, semaphore *sync.WaitGroup, result *types.ListPorts) {

	_, err := net.DialTimeout("tcp", address + ":" + strconv.Itoa(port), 1 * time.Second)
	if err == nil {
		result.AddElement(port, true)
	} else {
		result.AddElement(port, false)
	}

	if semaphore != nil {
		semaphore.Done()
	} else {
		result.PrintResult()
	}
}


func scanRangePortsTcp(address string, first int, second int) {

	diff := second - first
	i := 0

	if second < first {
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
	parser := argparse.NewParser("port_scanner",
						"questo programma serve a controllare le porte aperte di un host")

	host := parser.String("H", "host", &argparse.Options{
		Required: false,
		Help: "Determina l'host su cui si esegue lo scan",
	})

	rangePort := parser.IntList("r", "range", &argparse.Options{
		Required: false,
		Help: "Inserisci due numeri separati da uno spazio per scannerizzare le porte comprese tra queste",
	})

	port := parser.Int("p", "port", &argparse.Options{
		Required: false,
		Help: "Scannerizza solo una porta",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}


	if *port != 0 {
		scanRangePortsTcp(*host, *port, *port)
		return
	}

	scanRangePortsTcp(*host, (*rangePort)[0], (*rangePort)[1])
}
