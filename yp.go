package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var verbose bool

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stdout, "It'll take IP addresses and show the binary format. Set -v before for verbose mode\nexample: \n\typ 10.10.10.1 14.15.15.3 127.0.0.1\n\typ -v 10.10.10.1 14.15.15.3 127.0.0.1\n")

	}
	flag.BoolVar(&verbose, "v", false, "Verbose output will put show the corresponding ip along")
	flag.Parse()

	x := flag.Args()
	if flag.NArg() == 0 {
		usage()
	}

	for _, i := range x {
		ipTobin(i)
	}
	fmt.Println()

}

func ipTobin(ip string) {
	realI := net.ParseIP(ip)
	if realI == nil {
		log.Fatalln("Make sure to provide valid ip addresses")
	}

	gNet := strings.Split(ip, ".")
	bin0, _ := strconv.Atoi(gNet[0])
	bin1, _ := strconv.Atoi(gNet[1])
	bin2, _ := strconv.Atoi(gNet[2])
	bin3, _ := strconv.Atoi(gNet[3])

	if !verbose {
		fmt.Printf("%08b . %08b . %08b . %08b\n", byte(bin0), byte(bin1), byte(bin2), byte(bin3))
	} else if verbose {
		fmt.Printf("%08b . %08b . %08b . %08b	->	%v.%v.%v.%v\n", byte(bin0), byte(bin1), byte(bin2), byte(bin3), byte(bin0), byte(bin1), byte(bin2), byte(bin3))
	}
}

func usage()  {
	_, _ = fmt.Fprintf(os.Stderr, "It'll take IP addresses and show the binary format. Set -v before for verbose mode\nexample: \n\typ 10.10.10.1 14.15.15.3 127.0.0.1\n\typ -v 10.10.10.1 14.15.15.3 127.0.0.1")

}