package logic

import (
	"sync"
	"fmt"
)

type Student struct {
	Id int
	Name string
	Grade int
	Identify string
	Sex int
	BookMap map[string]*Book
	lock sync.Mutex
}

func NewStudent(id int, name string, grade int, identify string, sex int) (stu *Student) {
	stu = &Student {
		Id: id,
		Name: name,
		Grade: grade,
		Identify: identify,
		Sex:sex,
		BookMap:make(map[string]*Book, 32),
	}
	return
}


func (s *Student) AddBook (b *Book) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.BookMap[b.BookId] = b
	return
}

func (s *Student) BackBook (bookId string)(err error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	_, ok := s.BookMap[bookId]
	if !ok {
		err = fmt.Errorf("stu id:%d not exist book, book id:%s", s.Id, bookId)
		return
	}
	
	delete(s.BookMap, bookId)
	return
}