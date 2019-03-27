package palmtree_test

import (
	"errors"
	"io"
	"testing"

	"github.com/kavehmz/palmtree"
)

type mock struct{}

func (s *mock) Close() error {
	return errors.New("test_closed")
}

func TestPalmTree_Get(t *testing.T) {
	count := 0
	p := palmtree.PalmTree{
		New: func() io.Closer {
			count++
			return &mock{}
		},
		Buffer: 1,
	}

	c1 := p.Get()
	if c1 == nil || count != 1 {
		t.Error("wrong result", c1, count)
	}
	err := p.Put(c1)
	if err != nil {
		t.Error("Expected no error")
	}

	c2 := p.Get()
	if c2 == nil || count != 1 || c1 != c2 {
		t.Error("wrong result. count must stay 1. c1 and c2 must be the same", c1, count)
	}
	err = p.Put(c2)
	if err != nil {
		t.Error("Expected no error")
	}

	c3 := p.Get()
	c4 := p.Get()
	if count != 2 {
		t.Error("we must have 2 connections now", count, c3, c4)
	}

	err = p.Put(c3)
	if err != nil {
		t.Error("Expected to put back the connection")
	}
	err = p.Put(c4)
	if err.Error() != "test_closed" {
		t.Error("Expected to close the connection because of buffer full")
	}

	err = p.Get().Close()
	if err.Error() != "test_closed" {
		t.Error("Expected to get the same mock error ")
	}
}
