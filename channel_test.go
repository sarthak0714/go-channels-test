package main

import(
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T){
	server:=NewServer()
	server.Start()
	
	for i:=0;i<10;i++{
		go func (i int)  {
			server.userchan<-fmt.Sprintf("user_%d",i)
		}(i)
	}
	println("end of for")
}