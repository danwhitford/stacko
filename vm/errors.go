package vm

type DoNotPop struct{}

func (m DoNotPop) Error() string {
	return "please vm do not pop me"
}
