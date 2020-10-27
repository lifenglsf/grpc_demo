package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/jinzhu/copier"
	"github.com/lifenglsf/grpc_demo/hw/pb"
)

//ErrAlreadyExists error info
var ErrAlreadyExists = errors.New("record already exists")

//LaptopStore defined
type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
}

//InMemoryLaptopStore defined
type InMemoryLaptopStore struct {
	mutext sync.RWMutex
	data   map[string]*pb.Laptop
}

//NewInMemoryLaptopStore create new memory laptop store
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

//Save save laptop
func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutext.Lock()
	defer store.mutext.Unlock()
	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data:%w", err)
	}
	store.data[other.Id] = other
	return nil
}

//Find find laptop
func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutext.Lock()
	defer store.mutext.Unlock()
	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("can not copy laptop data:%w", err)
	}
	return other, nil
}
