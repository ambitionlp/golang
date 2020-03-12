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

	// crate a map
	//必须初始化
	var map1 = make(map[string]int) //
	// map2 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2

	for key := range map1 {
		fmt.Println("value:", map1[key])
	}
	//delete() function
	delete(map1, "a")
	//ok-idiom
	value, ok := map1["b"]
	fmt.Print("value:", value, "ok:", ok, "length:", len(map1))
	
//Swap change
func Swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

// Point 指针
func Point() {
	var a, b int
	var pa, pb *int
	var ppa **int
	pa = &a
	pb = &b
	ppa = &pa
	a = 10
	b = 20

	fmt.Println("a=", *pa, "a的地址为：", pa, "pa的地址为：", ppa)
	Swap(pa, pb)
	fmt.Println("a=", a, "b=", b)
}


