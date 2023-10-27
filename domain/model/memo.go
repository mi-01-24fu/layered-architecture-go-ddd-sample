package model

type Memo struct {
	Id      int
	Title   Title
	Content Content
	Date    Date
}

func NewMemo(title Title, content Content, date Date) *Memo {
	return &Memo{
		Title:   title,
		Content: content,
		Date:    date,
	}
}
