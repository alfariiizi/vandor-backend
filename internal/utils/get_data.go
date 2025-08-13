package utils

func GetData[T any](data *T, err error) (T, error) {
	var zero T
	if err != nil || data == nil {
		return zero, err
	}
	return *data, nil
}
