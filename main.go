package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"time"

	q3 "q3serverbrowser/q3query"
)

func main() {
	// make a connection
	udpConnection, err := net.DialTimeout("udp", q3.MASTER_SERVERS()[2], time.Duration(3)*time.Second)
	if err != nil {
		fmt.Println("Error connecting to master server:", err)
		return
	}
	defer udpConnection.Close()

	writer := bufio.NewWriter(udpConnection)
	_, err = writer.WriteString(q3.MSG_GETSERVERS)

	if err != nil {
		fmt.Println("Error sending request to master server:", err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing request to master server:", err)
		return
	}

	// Read response

	prefix := "\xFF\xFF\xFF\xFFgetserversResponse"

	// Must rely on this as only some master servers send an EOT msg
	udpConnection.SetReadDeadline(time.Now().Add(3 * time.Second))
	reader := bufio.NewReader(udpConnection)
	/*
		Reads in the maximum possible size of a dgram to process later
	*/
	datagram := make([]byte, 65507)
	/*
		Expanding byte array of processed bytes
	*/
	response := make([]byte, 0)
	for {
		n, err := reader.Read(datagram)
		if err != nil {
			if e, ok := err.(net.Error); !ok || !e.Timeout() {
				// handle error, it's not a timeout
				fmt.Println("Error receiving packets from  master server:", err)
			}
			break
		}

		// Remove getserversResponse prefix if present
		if bytes.Equal(datagram[:len(prefix)], []byte(prefix)) {
			datagram = datagram[len(prefix):n]
		} else {
			datagram = datagram[:n]
		}

		response = append(response, datagram...)

	}
	// Print Response
	var serverList []*q3.Server = parseServerList(response)
	fmt.Println("Found", len(serverList), "servers:")
	for _, server := range serverList {
		fmt.Println(server.IP)
	}

}

func parseServerList(data []byte) []*q3.Server {
	//check delimeter 0x5c or \\
	if data[0] != '\\' {
		return nil
	}

	// Parse server addresses
	var servers []*q3.Server
	for len(data) >= 6 {

		// Parsing for an EOT SHOULD be done with prefix removal
		// but only some master servers do this.
		if bytes.Equal(data[1:4], []byte("EOT")) {
			break
		}

		server, err := q3.NewServer(data[1:7])

		if err != nil {
			fmt.Println("Error parsing address:port byte data: ", err)
		}
		servers = append(servers, server)
		data = data[7:]
	}

	return servers
}
