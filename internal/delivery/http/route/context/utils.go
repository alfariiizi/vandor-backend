package httpctx

func GetBody[T any](ctx HttpContext) (T, error) {
	var t T
	if err := ctx.BindBody(&t); err != nil {
		var zero T
		return zero, err
	}
	return t, nil
}
