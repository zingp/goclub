package logic

import (
	"sync" 
	"errors"
)

type Book struct {
	BookId string
	Name string
	Num int
	Author string
	PublishDate int64
	lock sync.Mutex   // 互斥锁
}

// 构造函数
func NewBook(bookId string, name string, num int, author string, publishDate int64 ) (book *Book) {
	book = &Book {
		BookId: bookId,
		Name: name,
		Num: num,
		Author: author,
		PublishDate: publishDate,
	}
	return
}

func (b *Book) Borrow() (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.Num <= 0 {
		err = errors.New("books not enough ")
		return 
	}

	b.Num -= 1
	return
}

func (b *Book) Back() (err error){
	b.lock.Lock()
	defer b.lock.Unlock()

	b.Num += 1
	return
}