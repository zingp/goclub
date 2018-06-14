package logic

import "sync"

type BookMgr struct {
	BookList []*Book
	// 存储bookId 到借书学生列表的信息
	BookStudentMap map[string][]*Student
	//书籍名字到书籍列表的索引
	BookNameMap map[string][]*Book
	//书籍作者到书籍列表的索引
	BookAuthorMap map[string][]*Book

	lock sync.Mutex
}

func NewBookMgr() (bookMgr *BookMgr) {
	bookMgr = &BookMgr {
		BookStudentMap: make(map[string][]*Student, 16),
		BookNameMap: make(map[string][]*Book, 16),
		BookAuthorMap: make(map[string][]*Book, 16),
	}

	return
}

func (b *BookMgr) AddBook(book *Book)(err error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.BookList = append(b.BookList, book)
	bookList, ok := b.BookNameMap[book.Name]
	if !ok {
		var tmp []*Book
		b.BookNameMap[book.Name] = append(tmp, book)
	} else {
		bookList = append(bookList, book)
		b.BookNameMap[book.Name] = bookList
	}

	return
}

