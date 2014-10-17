package utility

func In_Slice(needle interface{}, haystack []interface{}) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}
