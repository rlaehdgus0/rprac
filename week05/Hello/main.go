package main // 패키지를 main에 속하게 함

import "fmt" // fmt라는 모듈을 임포트

func main() { // 함수 생성
	// name := "Go"               // name 이라는 변수에 "Go"라는 문자 저장
	// fmt.Println("Hello", name) // fmt를 이용하여 출력, ,을 통해서 변수와 문자열 구분
	number := 7
	fmt.Printf("Number is %d\n", number)
}
