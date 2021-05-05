package main

import "sync"

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{map[int]string{}, sync.RWMutex{}}
}

type InMemoryUserStore struct {
	user map[int]string
	lock sync.RWMutex
}

func (i *InMemoryUserStore) GetUserName(id int) string {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.user[id]
}

func (i *InMemoryUserStore) CreateUserByName(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	id := len(i.user) + 1
	i.user[id] = name
}
