package manager

import (
	"fmt"
	"log"
	"net"
	"time"

	//"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/lizhongz/dmoni/common"
	mPb "github.com/lizhongz/dmoni/proto/manager"
)

var ()

type masterServer struct {
	mng *Manager
}

func newMasterServer(mng *Manager) *masterServer {
	s := new(masterServer)
	s.mng = mng
	return s
}

func (s *masterServer) Run() {
	// Run clustering server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.mng.me.Port))
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	mPb.RegisterManagerServer(grpcServer, s)
	grpcServer.Serve(lis)
}

// SayHi registers an agent or updates its info, particularly heartbeat.
func (s *masterServer) SayHi(ctx context.Context, ni *mPb.NodeInfo) (*mPb.NodeInfo, error) {
	s.mng.agents.Lock()
	if _, present := s.mng.agents.m[ni.Id]; !present {
		// Create a new agent
		s.mng.agents.m[ni.Id] = &common.Node{
			Id:   ni.Id,
			Ip:   ni.Ip,
			Port: ni.Port,
		}
		grpclog.Printf("New agent %s %s:%d", ni.Id, ni.Ip, ni.Port)
	}

	grpclog.Printf("SayHi from %s", ni.Id)

	// Update agent's node info
	n := s.mng.agents.m[ni.Id]
	s.mng.agents.Unlock()

	//n.Ip = ni.Ip
	//n.Port = ni.Port
	n.Heartbeat = ni.Heartbeat
	n.Timestamp = time.Now()

	//grpclog.Printf("Hi from %s", n.Id)

	// Return my node info
	return &mPb.NodeInfo{
		Id:        s.mng.me.Id,
		Ip:        s.mng.me.Ip,
		Port:      s.mng.me.Port,
		Heartbeat: s.mng.me.Heartbeat,
	}, nil
}

// NotifyDone signifies the finish of an applicaiton and
// triggers application deregistration.
func (s *masterServer) NotifyDone(ctx context.Context, in *mPb.NDRequest) (*mPb.NDReply, error) {
	log.Printf("App %s exited", in.AppId)

	app := s.mng.getApp(in.AppId)
	if app.monitored {
		// Stop monitoring the app on agents
		newCtx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		err := s.mng.deregister(newCtx, app, true)
		if err != nil {
			log.Printf("Failed to deregister app %s: %v", in.AppId, err)
		}
	}

	// Store app info in db
	st := time.Unix(in.StartTime.Seconds, 0).Format(time.RFC3339)
	et := time.Unix(in.EndTime.Seconds, 0).Format(time.RFC3339)
	data := map[string]interface{}{
		"app_id":     app.Id,
		"entry_node": app.EntryNode,
		"exec":       app.ExecName,
		"args":       app.ExecArgs,
		"start_at":   st,
		"end_at":     et,
		"stdout":     in.Stdout,
		"stderr":     in.Stderr,
		"timestamp":  time.Now().Format(time.RFC3339),
	}
	go func() {
		if err := s.mng.logApp(data); err != nil {
			log.Printf("Failed to log application: %v", err)
		}
	}()

	s.mng.apps.Lock()
	delete(s.mng.apps.m, in.AppId)
	s.mng.apps.Unlock()

	return &mPb.NDReply{}, nil
}
