package service

import (
	store "goServer/entity"
	"sync"
)


type DataStore struct {
	sync.Mutex
	Books map[string]*store.Book
}

func (ms *DataStore) Create(book *store.Book) error{
	ms.Lock()
	defer ms.Unlock()
	if _, ok := ms.Books[book.Id]; ok{
		return store.ErrExist
	}
	nBook := *book
	ms.Books[book.Id] = &nBook
	return nil
}

func (ms *DataStore) Get(id string) (store.Book, error){
	ms.Lock()
	defer ms.Unlock()
	t, ok := ms.Books[id]
	if ok {
		return *t, nil
	}
	return store.Book{}, store.ErrNotFound
}


func (ms *DataStore) Delete(id string) error{
	ms.Lock()
	defer ms.Unlock()
	if _ ,ok := ms.Books[id]; ok{
		return store.ErrNotFound
	}
	
	delete(ms.Books, id)
	return nil
}