package single

import (
	"ditto.co.jp/grpc/w32"
)

//LimitSingleInstance -
type LimitSingleInstance struct {
	hMutex uintptr
	name   string
}

//New -
func New(uid string) *LimitSingleInstance {
	return &LimitSingleInstance{
		name: uid,
	}
}

//Open -
func (s *LimitSingleInstance) Open() (err error) {
	s.hMutex, err = w32.CreateMutex(s.name)
	if err != nil {
		return
	}

	return nil
}
