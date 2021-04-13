package main

func if_test() int {
	if 996 > 404 {
		return 996
	} else {
		return 404
	}
}

func for_test() int {
	count := 42
	for i := 0; i < count; i++ {
		i += 1
	}

	return count
}

func for_range_test() int {
	array := make([]int, 42)
	for _, elem := range array {
		elem += 1
	}
	return array[0]
}

func switch_test(a int) int {
	switch a {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return 4
	default:
		/* do nothing */
	}
	return 996
}
