package env

type Env string

const (
	Dev  = Env("dev")
	Prod = Env("prod")
)

func (e Env) isValid() bool {
	switch e {
	case Dev, Prod:
		return true
	default:
		return false
	}
}

func (e Env) String() string {
	return string(e)
}
