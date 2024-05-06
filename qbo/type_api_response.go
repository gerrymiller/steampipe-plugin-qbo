package qbo

type ApiResponse[T interface{}] interface {
	GetResponse() T
}
