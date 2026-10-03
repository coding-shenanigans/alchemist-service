package domain

type UserPageCursor struct {
	Q        string
	PageSize int
	LastId   int
}
