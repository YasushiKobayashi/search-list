package utils

import "net/url"

func UniqStringArray(list []string) []string {
	m := make(map[string]struct{})

	newList := make([]string, 0)

	for _, element := range list {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; ok == false {
			m[element] = struct{}{}
			newList = append(newList, element)
		}
	}

	return newList
}

func IsValidUrl(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	} else {
		return true
	}
}
