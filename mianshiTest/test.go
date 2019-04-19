package main

import (
	"fmt"
	"time"
)

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

//type student struct {
//	Name string
//}
//
//func zhoujielun(v interface{}) {
//	switch msg := v.(type) {
//	case *student, student:
//		msg.Name
//	}
//}

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func ArrayAppend() []int {
	arr := make([]int, 5)
	fmt.Printf("arr.len: %d; arr.cap: %d \n", len(arr), cap(arr))
	arr = append(arr, 10)
	//问现在 arr 结果是什么
	fmt.Printf("arr.len: %d; arr.cap: %d \n", len(arr), cap(arr))
	return arr
}

/**
 * 考察点Slices的变量储存方式 (切片是引用类型)
 * 所以每次对 array.append 做的修改,本身会对 array 指针指向的变量地址的值做修改.
 */
func ArrayEvenNumber(array []int) []int {
	for index, arr := range array {
		if arr%2 == 0 {
			array = append(array[:index], array[index+1:]...)
		}
	}
	return array
}

// panic先于defer执行，但是panic的向上传递是需要在defer执行完之后。
// 一个函数最终只会向上传递一个panic，后发生的panic会覆盖之前的panic

func multi_panic_defer() {
	// panic one执行之后panic two也就不会被执行了
	// 由于panic的执行咸鱼defer，但是defer发生于panic向上传递之前，
	// 因此panic one在被向上传递之前defer中的panic defer发生了，
	// 并且覆盖了panic one，因此最终向上传递被捕获的只有panic defer

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	defer func() {
		panic("panic defer")
	}()

	panic("panic one")
	panic("panic two") // 第一个panic后的代码就不会被执行
}

func defer_panic() {
	defer func() {
		fmt.Println("A")
	}()
	defer func() {
		fmt.Println("B")
	}()
	defer func() {
		fmt.Println("C")
	}()

	panic("异常")
}

type defintion_int int // 新定义一个类型，赋值需要强转
type alias_int = int   // 定义类型别名，是同一个类型

func (self *defintion_int) add(dela defintion_int) defintion_int {
	*self += dela
	return *self
}

// 直接对不同的位数进行判断
func P2(num int) (newNum int) {
	for num != 0 {
		temp := newNum*10 + num%10
		if temp/10 != newNum {
			return 0 // 溢出
		}
		newNum, num = temp, num/10
	}
	return
}

func main() {

	newNum := P2(-123214)
	fmt.Println(newNum)

	ArrayAppend()

	//p := &People{}
	//p.String()

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))

	// 0 1 zz zz 4
	fmt.Println(x, y, z, k, p)

	time.Now()

	// list *[]int
	//list := new([]int)
	//list = append(list, 1)
	//fmt.Println(list)
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
