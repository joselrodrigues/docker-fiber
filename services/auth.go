//go:generate mockgen -source auth.go -destination mock/auth_mock.go -package mock;
package services

type sendString interface {
	SendString(v string) error
}

func SingIn(c sendString) error {
	return c.SendString("Sing")
}

func SingUp(c sendString) error {
	return c.SendString("SingUp")
}
