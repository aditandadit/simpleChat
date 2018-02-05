package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

//RunHost Takes ip as Arg, Listens for connection on ip:8080
func RunHost(ip string) {
	ipPort := ip + ":" + port

	listener, err := net.Listen("tcp", ipPort)
	// usually functions Return returnVal, Err

	if err != nil {
		//fmt.Println("Error", err);
		//os.Exit(1);
		log.Fatal("Error : ", err)
		// This prints and exists with status 1 same as above
	}

	conn, err := listener.Accept()

	if err != nil {
		log.Fatal("Error :", err)
	}
	for {
		handleHost(conn)
	}
}

//RunGuest takes destination ip as argument and connects to it
func RunGuest(ip string) {
	ipPort := ip + ":" + port
	conn, err := net.Dial("tcp", ipPort)

	if err != nil {
		log.Fatal("Error Dialing : ", err)
		// Print error and Exit with os.Exit(1)
	}

	for {
		handleGuest(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readError := reader.ReadString('\n')
	// Delimited is Return ie when Host String has \n at end
	if readError != nil {
		log.Fatal("Error : ", readError)
	}

	fmt.Print("Message Received : ", message)

	// Now Host turn to send msg
	fmt.Print("Send Message : ")
	replyReader := bufio.NewReader(os.Stdin)
	msg, err := replyReader.ReadString('\n')
	if err != nil {
		log.Fatal("Error : ", err)
	}
	fmt.Fprint(conn, msg)
}

func handleGuest(conn net.Conn) {
	fmt.Print("Send Message : ")

	// Create Reader on StdIn of the process
	reader := bufio.NewReader(os.Stdin)

	// \n is the delimiter
	message, readErr := reader.ReadString('\n')

	if readErr != nil {
		log.Fatal("Error reading stdin : ", readErr)
	}

	// Use connection to Send Message;
	// takes in io.Writer -> Interface that implements Write
	// Connection implements the Writer interface
	fmt.Fprint(conn, message)


	//Now guest will read msg
	bufio.NewReader(conn);
	replyReader := bufio.NewReader(conn);
	msg, err := replyReader.ReadString('\n');

	if err != nil {
		log.Fatal("Error", nil);
	}

	fmt.Print("Message Received : ");
	fmt.Fprint(os.Stdin, msg);
}
