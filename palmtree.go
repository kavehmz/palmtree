package palmtree

import (
	"sync"
)

type closer = interface {
	Close() error
}

// PalmTree is a special pool for keeping closable connections
type PalmTree struct {
	// Buffer sets the maximum number of connection which the pool will keep alive
	// If buffer gets full new connections get closed and discarded
	Buffer uint64
	// A function which will generate new connections if needed
	New func() closer

	conns chan *closer
	sync.Mutex
}

// Get will return a new connection
func (s *PalmTree) Get() closer {
	s.Lock()
	if s.conns == nil {
		s.conns = make(chan *closer, s.Buffer)
	}
	s.Unlock()

	var con *closer
	select {
	case con = <-s.conns:
	default:
	}

	if con == nil {
		return s.New()
	}

	return *con
}

// Put returns the connection to the pool. If pool is full it will close the connection and discard it.
// Return error is set only if Put tries to close a connection and faces any error.
func (s *PalmTree) Put(con closer) error {
	select {
	case s.conns <- &con:
	default:
		return con.Close()
	}
	return nil
}
