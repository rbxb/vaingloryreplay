package server

import (
	"net/http"
	"strconv"
	"math/rand"
	"time"
	"github.com/rbxb/cmux"
	"github.com/rbxb/customprot"
	"github.com/rbxb/vaingloryreplay/vainglorystream"
	"github.com/rbxb/vaingloryreplay/vainglorystream/streamcore"
)

type Server struct {
	streams [] * streamcore.Stream
	ids []int
	lock chan byte
	mux cmux.Handler
}

func NewServer() * Server {
	srvr := &Server{
		streams: make([] * streamcore.Stream, 0),
		ids: make([]int, 0),
		lock: make(chan byte, 1),
	}
	srvr.mux = cmux.NewBranch("", []cmux.Handler{
		cmux.NewLeaf("stream", srvr.handleStream),
		cmux.NewLeaf("view", srvr.handleView),
	})
	rand.Seed(time.Now().UnixNano())
	return srvr
}

func(srvr * Server) ServeHTTP(w http.ResponseWriter, req * http.Request) {
	srvr.mux.ServeHTTP(w,req)
}

func(srvr * Server) generateId() int {
	var id int
	same := true
	for same {
		same = false
		id = rand.Intn(900000) + 100000
		for _, i := range srvr.ids {
			if i == id {
				same = true
				break
			}
		}
	}
	return id
}

func(srvr * Server) handleStream(w http.ResponseWriter, req * http.Request) {
	conn, err := customprot.Upgrade(w, req, map[string]string {
		"vainglory-replay-stream": "stream",
	}, map[string]string {
		"vainglory-replay-stream": "stream",
	})
	if err != nil {
		return
	}
	srvr.lock <- 0
	id := srvr.generateId()
	stream := streamcore.NewStream(conn, srvr.closeStream)
	srvr.streams = append(srvr.streams, stream)
	srvr.ids = append(srvr.ids, id)
	<- srvr.lock
	if err := vainglorystream.WriteFrame(conn, []byte(strconv.Itoa(id))); err != nil {
		stream.Close()
	}
}

func(srvr * Server) handleView(w http.ResponseWriter, req * http.Request) {
	str := req.URL.Query().Get("id")
	if id, err := strconv.Atoi(str); err == nil {
		var stream * streamcore.Stream = nil
		srvr.lock <- 0
		for i, x := range srvr.ids {
			if x == id {
				stream = srvr.streams[i]
				break
			}
		}
		<- srvr.lock
		if stream != nil {
			conn, err := customprot.Upgrade(w, req, map[string]string {
				"vainglory-replay-stream": "view",
			}, map[string]string {
				"vainglory-replay-stream": "view",
			})
			if err == nil {
				stream.Join(streamcore.NewChild(conn))
			}
			return
		}
	}
	http.Error(w, "Invalid stream id.", 400)
}

func(srvr * Server) closeStream(stream * streamcore.Stream) {
	srvr.lock <- 0
	for i, s := range srvr.streams {
		if s == stream {
			srvr.streams = append(srvr.streams[:i], srvr.streams[i+1:]...)
			srvr.ids = append(srvr.ids[:i], srvr.ids[i+1:]...)
			break
		}
	}
	<- srvr.lock
}