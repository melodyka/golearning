package main

//"errors"

/*
type error interface {
	Error() string
}

func Sqrt() error {
	return errors.New("math: square root of negative number")
}
*/
/*
func recur(n uint64) (result uint64) {
	if n > 0 {
		result = n * recur(n-1)
		return result
	}
	return 1
}
*/
/*
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
*/

func main() {

	//var myArray1 [10]int
	/*
		myArray1 := [10]int{1, 2, 3}
		fmt.Println(myArray1)

		var slice1 []int = make([]int, 3)
		fmt.Println(slice1, cap(slice1), len(slice1))
	*/
	/*
	   for i:=0;i<len(myArray1);i++{
	   	fmt.Println(myArray1[i])
	   }
	*/
	//尝试将数组赋值给数组，长度一致可以
	//oldMap := [10]int{1, 2, 3, 4}
	//var newMap [10]int

	//普通遍历赋值
	//for key, value := range oldMap {
	//	newMap[key] = value
	//}

	//尝试用slice赋值,两个切片即可
	//oldMap := [10]int{1, 2, 3, 4}
	//oldMap := []int{1, 2, 3, 4}
	//newMap := make([]int, 10)
	//var newMap []int
	//newMap = oldMap[0:10]
	/*
		for key, value := range oldMap {
			//fmt.Println(key, value)
			newMap[key] = value
		}
	*/
	//fmt.Println(newMap)

	//递归
	//var i uint64 = 3

	//fmt.Printf("test for recur %d\n", recur(i))

	//数据类型转换,ATOI有两个返回值，ITOA,ParseFloat，formatfloat
	/*
		var str string = "10"
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("转换错误:", err)
		} else {
			fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
		}
	*/
	/*
		var i interface{} = "Hello, World"
		str, ok := i.(string)
		if ok {
			fmt.Printf("'%s' is a string\n", str)
		} else {
			fmt.Println("conversion failed")
		}
	*/
	/*
		err := Sqrt()
		fmt.Println(err)
	*/
	/*
		s := []int{7, 2, 8, -9, 4, 0}
		c := make(chan int)
		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:], c)
		x, y := <-c, <-c
		fmt.Println(x, y, x+y)
	*/
}
