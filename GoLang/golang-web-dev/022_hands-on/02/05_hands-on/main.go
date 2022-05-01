package main

import (
	"bufio"
	"fmt"
	"log"
	"io"
	"net"	
)


func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				fmt.Println("ENDOF")
				break
			}
		}
		defer conn.Close()

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")
	}
}
