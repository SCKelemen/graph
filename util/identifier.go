package util

import (
	"fmt"
	"log"
	"net"
	"os/user"
	"runtime"

	uuid "github.com/satori/go.uuid"
)

func CreateID() string {
	usr, _ := user.Current()
	uid, _ := uuid.NewV4()
	return fmt.Sprintf("%s@%s", usr.Uid, uid)
}

func GatherInfo() (UserInfo, RunTimeInfo) {
	usr := getUserInfo()
	gop := getRuntimeInfo()
	return usr, gop
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

type UserInfo struct {
	Username      string
	UserID        string
	GroupID       string
	GroupIDs      []string
	HomeDirectory string
	Name          string
}

type RunTimeInfo struct {
	GOOS   string
	GOARCH string
	CPUs   int
}

func getRuntimeInfo() RunTimeInfo {
	numCPUs := runtime.NumCPU()
	return RunTimeInfo{GOOS: runtime.GOOS, GOARCH: runtime.GOARCH, CPUs: numCPUs}
}

func getUserInfo() UserInfo {
	info := UserInfo{}
	usr, err := user.Current()
	if err == nil {
		groups, _ := usr.GroupIds()
		info = UserInfo{Username: usr.Username, UserID: usr.Uid, Name: usr.Name, HomeDirectory: usr.HomeDir, GroupID: usr.Gid, GroupIDs: groups}
	}
	return info
}
