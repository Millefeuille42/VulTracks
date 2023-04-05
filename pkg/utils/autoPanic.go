package utils

func AutoPanic(err error) {
	if err != nil {
		panic(err)
	}
}
