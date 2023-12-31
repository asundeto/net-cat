# NET-CAT

## Authors:

atemerzh and asundeto

## Description:

This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

* NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports and many more.

### The project must have the following features :

* TCP connection between server and multiple clients (relation of 1 to many).
* A name requirement to the client.
* Control connections quantity.
* Clients must be able to send messages to the chat.
* Do not broadcast EMPTY messages from a client.
* Messages sent, must be identified by the time that was sent and the user name of who sent the message, example : [2020-01-20 15:48:41][client.name]:[client.message]
* If a Client joins the chat, all the previous messages sent to the chat must be uploaded to the new Client.
* If a Client connects to the server, the rest of the Clients must be informed by the server that the Client joined the group.
* If a Client exits the chat, the rest of the Clients must be informed by the server that the Client left.
* All Clients must receive the messages sent by other Clients.
* If a Client leaves the chat, the rest of the Clients must not disconnect.
* If there is no port specified, then set as default the port 8989. Otherwise, program must respond with usage message: [USAGE]: ./TCPChat $port

## Usage: how to run

student@ALEM~/net-cat$ **go run .**

Listening on the port :8989

**or** 

student@ALEM~/net-cat$ **go run . $port**

Listening on the port :$port

* Open the terminal

* Type **nc localhost $port**

* If you have done everything correctly, the TCP-Chat logo will appear!

* Example:
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:

* Repeat with another terminals