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

// 增加书
func (b *BookMgr) AddBook(book *Book)(err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	//1 添加到book列表中
	b.BookList = append(b.BookList, book)

	//2 更新书籍名字到同一个书籍名字对应的book列表
	bookList, ok := b.BookNameMap[book.Name]
	if !ok {
		var tmp []*Book
		b.BookNameMap[book.Name] = append(tmp, book)
	} else {
		bookList = append(bookList, book)
		b.BookNameMap[book.Name] = bookList
	}

	bookList, ok = b.BookNameMap[book.Author]
	if !ok {
		var tmp []*Book
		b.BookNameMap[book.Author] = append(tmp, book)
	} else {
		bookList = append(bookList, book)
		b.BookNameMap[book.Author] = bookList
	}

	return
} 

//检索
func (b *BookMgr) SearchByBookName(bookName string)(bookList []*Book) {
	b.lock.Lock()
	defer b.lock.Unlock()
	bookList = b.BookNameMap[bookName]
	return
}