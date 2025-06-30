package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	isTrue := flag.Bool("isTrue", false, "是否正确")
	name := flag.String("name", "xx", "名字")
	dur := flag.Duration("time", time.Second, "时间差")
	// 必须在所有测试运行前解析flag
	flag.Parse()
	fmt.Println("=== TestFlag ===")

	// 打印解析后的flag值
	fmt.Printf("isTrue: %v\n", *isTrue)
	fmt.Printf("name: %s\n", *name)
	fmt.Printf("time: %v\n", *dur)

	// 打印其他参数
	fmt.Println("flag.Args():", flag.Args())
	fmt.Println("flag.NArg():", flag.NArg())
	fmt.Println("flag.NFlag():", flag.NFlag())
	fmt.Println("=== TestName ===")
	for i, arg := range os.Args {
		fmt.Printf("os.Args[%d]: %v\n", i, arg)
	}
}
