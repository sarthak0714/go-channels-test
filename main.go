package main

import "fmt"

type Server struct{
	users map[string]string
	userchan chan string
}

func NewServer() *Server{
	return &Server{
	users: make(map[string]string),
	userchan: make(chan string),
	}
}

func (s *Server) Start(){
	go s.loop()
}

func (s *Server) loop(){
	for{
		user:=<-s.userchan
		s.users[user]=user
		fmt.Printf("adding new user %s\n",user)
	}
}

func (s *Server) AddUser(user string){
	s.users[user]=user
}

func main(){
	msgch:=make(chan string)
	go sendMsg(msgch)
	go readMsg(msgch)
	rvcdmsg:=<-msgch
	fmt.Printf("the messgae is: %s",rvcdmsg)
}

// only read the send the value to the channel unable to modify or read just add

func sendMsg(msgc chan <- string){
	msgc<-"Nyaaa"
}

// acess the comming message and interact with the message

func readMsg(msgc <- chan string){
	message:=<-msgc
	fmt.Println(message)
}