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

func main()  {
	flag.Usage = func() {
		fmt.Println(`yp converts ipv4 from bin->dec and vice versa. You can provide as many ips as you want, as args
ex:	
	yp -b 11111111.11111111.11111111.11111111 11111111.00000000.11111111.00000000 ... [bin->dec]
	yp -d 255.255.255.255 127.0.0.1 192.168.192.168 ... [dec->bin]
	yp -v [for verbose output]`)
	}

	flag.BoolVar(&verbose, "v", false, "verbose mode")
	var bin bool
	flag.BoolVar(&bin,"b", false, "ip bin to dec")
	var dec bool
	flag.BoolVar(&dec, "d", false, "ip dec to bin")
	flag.Parse()

	x := flag.Args()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}

	if bin == true && dec == true {
		log.Fatalln("You can only do one operation at a time")

	} else if bin == true {
		for _, i := range x {
			binToIp(i)
		}

	} else if dec == true {
		for _, i := range x {
			ipToBin(i)
		}
	}


	//ipToBin("255.255.255.0")
	//binToIp("11111111.11111111.11111111.11111111")
}

func binToIp(ipBin string)  {
	octalSplit := strings.Split(ipBin, ".")

	// Checks whether the octets are 8 bits
	for _, i := range octalSplit {
		if len(i) != 8 {
			log.Fatalln("Make sure to provide valid ipv4 addresses in binary")
		}
	}

	dec0 := coreBiP(octalSplit[0])
	dec1 := coreBiP(octalSplit[1])
	dec2 := coreBiP(octalSplit[2])
	dec3 := coreBiP(octalSplit[3])

	if verbose {
		fmt.Printf("%v.%v.%v.%v			->	%v . %v . %v . %v\n", dec0, dec1, dec2, dec3, octalSplit[0], octalSplit[1], octalSplit[2], octalSplit[3])
	} else {
		fmt.Printf("%v.%v.%v.%v\n", dec0, dec1, dec2, dec3 )
	}
}

func coreBiP(binaryString string) (decimal int64) {
	i, err := strconv.ParseInt(binaryString, 2, 64)

	if err != nil {
		log.Fatal(err)
	} else {
		decimal = i
	}
	return
}

func ipToBin(ipDec string) {
	realI := net.ParseIP(ipDec)
	if realI == nil {
		log.Fatalln("Make sure to provide valid ipv4 addresses in decimal")
	}

	gNet := strings.Split(ipDec, ".")
	bin0, _ := strconv.Atoi(gNet[0])
	bin1, _ := strconv.Atoi(gNet[1])
	bin2, _ := strconv.Atoi(gNet[2])
	bin3, _ := strconv.Atoi(gNet[3])

	if verbose {
		fmt.Printf("%08b . %08b . %08b . %08b	->	%v.%v.%v.%v\n", byte(bin0), byte(bin1), byte(bin2), byte(bin3), byte(bin0), byte(bin1), byte(bin2), byte(bin3))
	} else {
		fmt.Printf("%08b . %08b . %08b . %08b\n", byte(bin0), byte(bin1), byte(bin2), byte(bin3))
	}
}
