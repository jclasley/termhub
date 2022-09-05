package internal

func Every[T any](m []T, f func(int, T)) {
	for k, v := range m {
		f(k, v)
	}
}
