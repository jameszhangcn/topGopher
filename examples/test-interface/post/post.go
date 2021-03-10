package post
import ("google.golang.org/grpc")

type Service interface {
	ListPosts() ([]string, error)
}

type service struct {
	conn *grpc.ClientConn
}

func NewService(conn *grpc.ClientConn) Service {
	return &service {
		conn: conn,
	}
}

func (s *service) ListPosts()([]string, error) {
	//posts, err := s.conn.conns
	//if err != nil {
		//return []*addrConn{}, err
	//}
	rt := make([]string, 0, 10)
	rt = append(rt,"aaa")
	rt = append(rt,"bbb")
	return rt, nil
}