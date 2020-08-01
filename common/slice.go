package common

// StringRemove 删除指定的元素
func StringRemove(slice []string, delete string) {

	for index, s := range slice {

		if s == delete {
			slice = append(slice[:index], slice[index+1:]...)
		}
	}
}
