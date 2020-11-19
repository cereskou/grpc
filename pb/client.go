package pb

import (
	"fmt"

	grpc "google.golang.org/grpc"
)

//LocalService -
type LocalService struct {
	client LocalServiceClient
	conn   *grpc.ClientConn
}

//GetLocalService -
func GetLocalService(port int) (*LocalService, error) {
	url := fmt.Sprintf("localhost:%v", port+1)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := NewLocalServiceClient(conn)

	return &LocalService{
		client: client,
		conn:   conn,
	}, nil
}

//RPC -
func (s *LocalService) RPC() LocalServiceClient {
	return s.client
}

//Close -
func (s *LocalService) Close() {
	s.conn.Close()
}
