package test

func Equal(a, b []byte) bool {

	for i := range a {

		if a[i] != b[i] {
			return false
		}
	}

	return true

}
