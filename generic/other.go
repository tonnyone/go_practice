package generic

// type constraint
type Number interface {
	int64 | float64
}

type A interface {
}
type B interface {
}
type AB interface {
	A
}

func add(i, j AB) AB {
	return nil
}
func add2[T AB](i, j A) T {
	return nil
}

func test() any {
	return "asdfa"
}
