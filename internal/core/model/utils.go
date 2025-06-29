package model

func ToList[T any, U ~*T](items []*T) []U {
	res := make([]U, len(items))
	for i, item := range items {
		res[i] = U(item)
	}
	return res
}
