	// creat a slice
	var slice1 []int = make([]int, 3, 3)
	slice2 := make([]int, 5)
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	// 多个切片共享底层数组时，任一切片改变都会影响底层数组
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice3 := slice1
	slice3[0] = 0
	fmt.Println("slice1:", slice1)
	// len(),cap()
	fmt.Println("length:", len(slice1), "capacity:", cap(slice1))
	// append(),copy()
	slice1 = append(slice1, 4, 5)
	fmt.Println("slice1:", slice1)
	copy(slice2, slice1)
	fmt.Println("slice2", slice2)
  
