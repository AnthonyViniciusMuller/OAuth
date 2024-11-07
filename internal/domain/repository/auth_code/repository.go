package authcode

type Repository interface {
	Insert(code string, userID int64) error
	GetUserID(code string) (int64, error)
}
