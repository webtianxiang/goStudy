package main

import "fmt"

// 赋值运算符（=、+=、-=、*=、/=、%=、<<=、>>=、&=、|=、^=）

func main() {
	var (
		x = 2
	)
	x += 2  // x = x + 2
	x -= 2  // x = x - 2
	x *= 2  // x = x * 2
	x /= 2  // x = x / 2
	x <<= 2 // x = x << 2
	x >>= 2 // x = x >> 2
	x &= 2  // x = x & 2
	x |= 2  // x = x | 2
	x ^= 2  // x = x ^ 2
	fmt.Println(x)
}
