package string

const DefaultStringBase = 10

type Builder struct {
	Base int
}

func (b *Builder)SetBase(base int)  {
	b.Base = base
}
