package main

/*
	Go언어의 기본 자료형
	bool : true, false를 저장합니다. 기본값은 false
	string : 문자 / 문자열을 저장합니다. 기본값은 ""

	int : 정수형. 기본값은 0
	int / int8 / int16 / int32 /
	int64 / uint / uint8 / uint16 /
	uint32 / uint64 / uintptr

	byte : uint8과 같습니다
	rune : int32와 같습니다. 유니코드 포인트를 나타냅니다.
	float32 / float 64 : 실수형을 저장합니다. 기본값은 0
	complex64 / complex128 : 복소수를 저장합니다. 기본값은 0+0i

	int, uint, uintptr 타입은 32-비트 시스템에서는 32비트 길이고,
	64-비트 시스템에서는 64비트 길이입니다. 특별히 정수의 크기나 부호
	(unsigned)를 지정할 이유가 없다면 int를 쓰면 됩니다.

	Go 강타입 언어이므로, 무시적 형 변환을 하지 않습니다.
	타입이 맞지 않으면 값을 할당할 수 없습니다.
	이럴 때는 변환할 타입(변환할 값)과 같이 형 변환을 꼭 해줘야 합니다.
	코드를 실행하면 19번째 줄에서 에러가 발생합니다.
	int타입 값을 float64 타입을 저장하는 변수에 할당하기 때문입니다.
	이 부분을 지우고 실행해보세요.
*/

import (
	"fmt"
	"math/cmplx"
)

var (
	i      int
	f      float64
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const format = "%T(%v)\n"
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, z, z)

	// int에서 float64로 묵시적 타입 변환을 할 수 없습니다
	// f = i

	// 다른 타입을 저장하려면 변환할타입(변수) 와 같이, 형 변환을 해줘야 합니다.
	f = float64(i)

	f = 0.0

	var z2 complex128
	fmt.Println("복소수의 기본값 : ", z2)
}
