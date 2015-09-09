/** **************************************************************************
 *  Copyright (c) 2015 Lobaro UG (haftungsbeschränkt)
 *  T.Rohde
 *
 *  ### Lobaro.com Simple UDP Test ###
 *
 *  Description:
 *  - This app sends out the current time as string to an given IPv6 address via UDP.
 *	- The interval for transmison of these udp packets is one second.
 * 	- The remote address and port must be set as commandline parameter.
 *	- Local port will be choosen randomly by the operating system.
 *	- Any received data will be simly shown on the command line as string.
 *
 *****************************************************************************/

package main

import (
	"net"
	"log"
	"fmt"
	"time"
	"os"
)

func main() {
	
	
	argcnt := len(os.Args)
	
	if argcnt != 2 {
		fmt.Println("usage: udptest.exe [IPv6_of_LobaroBox:Port]\r\nExample: udptest.exe [fe80:0000:0000:0000:0211:7d00:0030:8e3f]:5684")
		return
	}

	//Remote Addr
	BoxAddr, err := net.ResolveUDPAddr("udp6", os.Args[1])
	if(err != nil) {
		log.Fatal(err)
		return
	}
	
	LocalAddr, err := net.ResolveUDPAddr("udp6", ":0") //:0 => OS sets local port
	if(err != nil) {
		log.Fatal(err)
		return
	}
	
	c, err := net.ListenUDP("udp6", LocalAddr)
	if(err != nil) {
		log.Fatal(err)
		return
	}
	defer c.Close()
	
	fmt.Println("Start Listening to:", c.LocalAddr())

	//readsocket loop
	go func(c *net.UDPConn) {
		for {
			rxMsg := make([]byte, 512);
			n,remote, err := c.ReadFromUDP(rxMsg)

			if(err != nil) {
				log.Fatal(err)
				return
			} else {
				fmt.Println("Got from ",remote," ", n, "Bytes: ", string(rxMsg[:n]))
			}
		}
	}(c)
	
	//write every second current time string to remote "LOBARO IPV6 UNIVERSAL BOX"
	for {
		b := []byte(time.Now().String())
		n, err :=c.WriteTo(b,BoxAddr)
		if(err != nil) {
			log.Fatal(err)
			return
		} else {
			fmt.Println("Wrote to ", BoxAddr," ",n, " Bytes: ", string(b[:n]))
		}
		
		time.Sleep(1000 * time.Millisecond)
	}
}