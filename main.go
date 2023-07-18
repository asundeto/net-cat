package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

var (
	clientsStr []string
	clientsNet []net.Conn
	leaving    = make(chan message)
	messages   = make(chan message)
	logo       = ReadWholeFile("./logo.txt")
	file       os.File
	boly       bool
	start      bool
)

type message struct {
	text    string
	address string
}

func main() {
	port := "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	// mux := http.NewServeMux()
	// err := http.ListenAndServe(":"+port, mux)
	// log.Fatal(err)
	listen, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Println("Incorrect port, try another one!")
		return
	}
	file1, err := os.Create("history.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	file = *file1
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on the port:", port)
	go broadcaster()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handle(conn)
	}
}

func handle(cl net.Conn) {
	var login bool
	if checkCnt(cl) {
		cl.Close()
		return
	}
	cl.Write([]byte(logo))
	cl.Write([]byte("Enter your name: "))
	input := bufio.NewScanner(cl)
	x := 0
	for input.Scan() {
		if x == 0 {
			name := string(input.Text())
			if checkUserName(name, cl) {
				continue
			}
			clientsStr = append(clientsStr, name)
			clientsNet = append(clientsNet, cl)
			history, err := ioutil.ReadFile(file.Name())
			if err != nil {
				fmt.Println("Failed to open history:", err)
				return
			}
			cl.Write(history)
			messages <- newMessage(" has joined.", cl, false)
			x++
			login = true
			continue
		}
		inpTxt := input.Text()
		if checkMsg(inpTxt, cl) {
			continue
		}
		messages <- newMessage(inpTxt, cl, true)
		x++
	}
	if login {
		leaving <- newMessage(" has left.", cl, false)

		deleteClient(cl)

		cl.Close()
	} else {
		cl.Close()
	}
}

func deleteClient(cl net.Conn) {
	var newClientsStr []string
	var newClientsNet []net.Conn
	for i := 0; i < len(clientsNet); i++ {
		if clientsNet[i] != cl {
			newClientsNet = append(newClientsNet, clientsNet[i])
			newClientsStr = append(newClientsStr, clientsStr[i])
		}
	}
	clientsStr, clientsNet = newClientsStr, newClientsNet
}

func saveHistory(s string, h *os.File) {
	if _, err := h.WriteString(s); err != nil {
		log.Fatal()
	}
}

func checkMsg(s string, cl net.Conn) bool {
	if s == "" {
		name := "[" + clientsStr[getId(cl)] + "]"
		cl.Write([]byte(getTime(name)))
		return true
	}
	return false
}

func checkCnt(cl net.Conn) bool {
	if len(clientsStr) >= 2 {
		cl.Write([]byte("Server is full! Try again latter."))
		return true
	}
	return false
}

func getId(cl net.Conn) int {
	for i := 0; i < len(clientsNet); i++ {
		if cl == clientsNet[i] {
			return i
		}
	}
	return 0
}

func checkUserName(s string, cl net.Conn) bool {
	if s == "" {
		cl.Write([]byte("Please enter correct name!\nEnter your name: "))
		return true
	}
	for i := 0; i < len(clientsStr); i++ {
		if s == clientsStr[i] {
			cl.Write([]byte("Username exists! Try another username.\nEnter your name: "))
			return true
		}
	}
	return false
}

func newMessage(msg string, conn net.Conn, b bool) message {
	addr := conn.RemoteAddr().String()
	name := clientsStr[getId(conn)]
	if b {
		name = "[" + name + "]"
		name = getTime(name)
		boly = false
	} else {
		boly = true
	}
	return message{
		text:    name + msg,
		address: addr,
	}
}

func getTime(name string) string {
	return fmt.Sprintf("\r[%s]%s: ", time.Now().Format("01-02-2006 15:04:05"), name)
}

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			for _, conn := range clientsNet {
				name := "[" + clientsStr[getId(conn)] + "]"
				name = getTime(name)
				if msg.address == conn.RemoteAddr().String() {
					conn.Write([]byte(name))
					continue
				}
				conn.Write([]byte("\n" + msg.text + "\n" + name))
			}
			if boly {
				if !start {
					saveHistory(msg.text+"\n", &file)
					start = true
					continue
				}
				saveHistory(msg.text+"\n", &file)
			} else {
				saveHistory(msg.text+"\n", &file)
			}
		case msg := <-leaving:
			for _, conn := range clientsNet {
				conn.Write([]byte("\n" + msg.text + "\n"))
				evrName := "[" + clientsStr[getId(conn)] + "]"
				evrName = getTime(evrName)
				conn.Write([]byte(evrName))
			}
			saveHistory(msg.text+"\n", &file)
		}
	}
}

func ReadWholeFile(logo string) string {
	contents, err := ioutil.ReadFile(logo)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(string(contents)) == 0 {
		fmt.Println("Error! Empty parametre!")
		return "ERROR"
	}
	return string(contents)
}
