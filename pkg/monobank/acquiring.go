package monobank

type Acquiring struct {
	xToken string
}

func NewAcquiring(xToken string) *Acquiring {
	return &Acquiring{
		xToken: xToken,
	}
}
