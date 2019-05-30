package engine_monitor
import (
	"encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"matching-engine/engine"
	pb "matching-engine/engine_monitor/engine_monitor"
	"net"
)

const (
	port = ":50051"
)

type Monitor_Server struct {}
// server is used to implement helloworld.GreeterServer.
type server struct{
	OrderBook *engine.OrderBook
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	out, err := json.Marshal(s.OrderBook)
	if err != nil {
		panic(err)
	} else {
		return &pb.HelloReply{Message: "Hello " + in.Name + string(out)}, nil
	}
}

//func main() {
//	start_monitor_server()
//}

func (ms *Monitor_Server) StartMointor(book *engine.OrderBook) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{book})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return err
}
