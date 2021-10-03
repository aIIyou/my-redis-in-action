package chapter2

const (
	LOG_KEY = "login:"
)

type User struct {
	Name          string
	UID           int64
	LastLoginTime int64
}
