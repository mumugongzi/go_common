package tool

func StrInList(str string, list []string) bool {
	for _, candidate := range list {
		if str == candidate {
			return true
		}
	}
	return false
}

func IntInList(i int, list []int) bool {
	for _, candidate := range list {
		if i == candidate {
			return true
		}
	}
	return false
}

func Int64InList(i int64, list []int64) bool {
	for _, candidate := range list {
		if i == candidate {
			return true
		}
	}
	return false
}
