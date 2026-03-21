package util

type Comperable[T any] interface {
	Equals(other T) bool
}