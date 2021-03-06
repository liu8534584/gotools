package ArrayUtils

func Get(key int, arr []interface{}) (interface{}, bool) {
	if len(arr) == 0 {
		return nil, false
	}
	if len(arr) < key-1 {
		return nil, false
	}
	return arr[key], true
}

func StringArrayToInterfaceArray(names []string) []interface{} {
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	return vals
}

//是否在数组中
func InArray(value interface{}, arr []interface{}) bool {
	for _, v := range arr {
		if value == v {
			return true
		}
	}

	return false
}

// a数组中是否包含b数组的全部元素
func ArrayInArray(aArray, bArray []interface{}) bool {
	if len(aArray) == 0 {
		return false
	}

	isAllInArray := true
	for _, b := range aArray {
		if !InArray(b, aArray) {
			isAllInArray = false
		}
	}
	return isAllInArray
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
