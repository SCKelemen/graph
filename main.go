package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/sckelemen/c2/router"
	"github.com/sckelemen/graph/graph"

	msgs "github.com/SCKelemen/messages"
)

func main() {
	id := createIdentifier()
	fmt.Printf("I am identifying as %s \n", id)
	addr := GetOutboundIP()
	port, listener := GetPort()
	//api.Connect(listener, port)
	listener.Close()
	conn := fmt.Sprintf("%s:%v", addr.String(), port)
	self := graph.Vertex{ID: id, Address: conn}
	self.Print()

	lp := fmt.Sprintf(":%v", port)

	factory := msgs.New()
	birth := factory.CreateBirthMessage(self.ID, self.Address)

	jsond, err := json.Marshal(birth)
	if err == nil {
		http.Post("http://localhost:9999/graph", "application/json; charset=utf-8", bytes.NewBuffer(jsond))
	}
	go router.Run(lp)

	router := gin.Default()
	router.POST("/notify", func(c *gin.Context) {
		var msg msgs.Message
		c.BindJSON(&msg)

		switch msg.Type {
		case msgs.BirthMessageType:
			var data msgs.BirthData
			err = json.Unmarshal(msg.Data, &data)
			if err != nil {
				c.Status(400)
			}
			fmt.Printf("BIRTH: %s %s", data.Name, data.Address)
			//dispatcher.Dispatch(msg)
			//dispatcher.Subscribe(data.Name, data.Address)
			break
		case msgs.DeathMessageType:
			//fmt.Printf("DEATH: %s %s", data.Name, data.Address)
			//dispatcher.Dispatch(msg)
			break
		case msgs.SuicideMessageType:
			//fmt.Printf("SUICIDE: %s %s", data.Name, data.Address)
			//dispatcher.Dispatch(msg)
			break
		default:
			break
		}
		c.Status(200)
	})

	c := make(chan bool)
	<-c
}

func createIdentifier() string {
	user, err := user.Current()
	handleError(err)
	uuid, err := uuid.NewV4()
	handleError(err)

	id := fmt.Sprintf("%s@%s", user.Username, uuid)
	fmt.Printf("Username:\t%s \nName:\t\t%s\nGid:\t\t%s\nUid:\t\t%s\n", user.Username, user.Name, user.Gid, user.Uid)
	fmt.Println("Groups: ")
	groups, _ := user.GroupIds()
	for item := range groups {
		fmt.Printf("\t\t%v\n", item)
	}
	return id
}
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetPort() (int, net.Listener) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	port := listener.Addr().(*net.TCPAddr).Port
	fmt.Println("Using port:", port)
	return port, listener
}

func handleError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type vertex struct {
	*graph.Vertex
}

func (v vertex) Reload() bool {
	return true
}
