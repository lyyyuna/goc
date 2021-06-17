package server

import (
	"net/rpc"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// gocServer represents a goc server
type gocServer struct {
	port      int
	storePath string
	upgrader  websocket.Upgrader

	rpcClients sync.Map
	// mu         sync.Mutex // used to protect rpcClients access
}

type gocCliendId string

// gocCoveredClient represents a covered client
type gocCoveredClient struct {
	Id       string      `json:"id"`
	RemoteIP string      `json:"remoteip"`
	Hostname string      `json:"hostname"`
	CmdLine  string      `json:"cmdline"`
	Pid      string      `json:"pid"`
	rpc      *rpc.Client `json:"-"`
	exitCh   chan int    `json:"-"`
}

func RunGocServerUntilExit(port int) {
	gs := gocServer{
		port:      port,
		storePath: "",
		upgrader: websocket.Upgrader{
			ReadBufferSize:   4096,
			WriteBufferSize:  4096,
			HandshakeTimeout: 45 * time.Second,
		},
	}

	r := gin.Default()
	v2 := r.Group("/v2")
	{
		v2.GET("/cover/profile", gs.getProfiles)
		v2.DELETE("/cover/profile", nil)
		v2.GET("/services", gs.listServices)

		v2.GET("/cover/ws/watch", nil)

		// internal use only
		v2.GET("/internal/ws/rpcstream", gs.serveRpcStream)
		v2.GET("/internal/ws/watchstream", nil)
	}

	r.Run(":" + strconv.Itoa(port))
}