package errorx

type Errorx interface {
	Error() string
	ErrTag() string
}
