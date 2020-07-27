package future

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type IFuture interface {
	Success(SuccessFunc) IFuture
	Fail(FailFunc) IFuture
	Execute(ExecuteStringFunc)
}

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (s *MaybeString) Success(f SuccessFunc) IFuture {
	s.successFunc = f
	return s
}
func (s *MaybeString) Fail(f FailFunc) IFuture {
	s.failFunc = f
	return s
}
func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func() {
		str, err := f()
		if err != nil {
			s.failFunc(err)
			return
		}
		s.successFunc(str)
	}()
}
