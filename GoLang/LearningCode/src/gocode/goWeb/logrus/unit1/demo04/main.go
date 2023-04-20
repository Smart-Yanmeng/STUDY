package main

import "fmt"

const (
	cBlack  = 0
	cRed    = 1
	cGreen  = 2
	cYellow = 3
	cBlue   = 4
	cPurple = 5
	cCyan   = 6
	cGray   = 7
)

// PrintColor - 自封装函数
func PrintColor(colorCode int, text string, isBackgroundColor bool) {
	if isBackgroundColor {
		fmt.Printf("\033[4%dm%s\033[0m", colorCode, text)

		return
	}
	fmt.Printf("\033[3%dm%s\033[0m", colorCode, text)
}

func main() {
	/* 控制台输出颜色 */
	// 前景色
	//fmt.Println("\033[30mBLACK\033[0m")
	//fmt.Println("\033[31mRED\033[0m")
	//fmt.Println("\033[32mGREEN\033[0m")
	//fmt.Println("\033[33mYELLOW\033[0m")
	//fmt.Println("\033[34mBLUE\033[0m")
	//fmt.Println("\033[35mPURPLE\033[0m")
	//fmt.Println("\033[36mCYAN\033[0m")
	//fmt.Println("\033[37mGRAY\033[0m")
	// 背景色
	//fmt.Println("\033[40mBLACK\033[0m")
	//fmt.Println("\033[41mRED\033[0m")
	//fmt.Println("\033[42mGREEN\033[0m")
	//fmt.Println("\033[43mYELLOW\033[0m")
	//fmt.Println("\033[44mBLUE\033[0m")
	//fmt.Println("\033[45mPURPLE\033[0m")
	//fmt.Println("\033[46mCYAN\033[0m")
	//fmt.Println("\033[47mGRAY\033[0m")

	PrintColor(cBlue, "York", true)
}
