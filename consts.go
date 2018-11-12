package main

import "fmt"

const SingleString = "SingleString Constant"

const (
	MultilineFirst = 1
	MultilineSecond = "second"
	MultilineEmpty
)

const (
	IotaFirst = iota
	IotaSecond
	IotaThird
	IotaFourth
)

const (
	IotaUseCaseFirst = 1 << iota
	IotaUseCaseSecond
	IotaUseCaseThird
	IotaUseCaseFourth
)

const (
	IotaUseCase2First = iota * 3 - 2
	IotaUseCase2Second
	IotaUseCase2Third
	IotaUseCase2Fourth
)

func main () {
	fmt.Println("SingleString: ", SingleString, "\n");

	fmt.Println("MultilineFirst: ", MultilineFirst, "\n");
	fmt.Println("MultilineSecond: ", MultilineSecond, "\n");
	fmt.Println("MultilineEmpty: ", MultilineEmpty, "\n");

	fmt.Println("IotaFirst: ", IotaFirst, "\n");
	fmt.Println("IotaSecond: ", IotaSecond, "\n");
	fmt.Println("IotaThird: ", IotaThird, "\n");
	fmt.Println("IotaFourth: ", IotaFourth, "\n");

	fmt.Println("IotaUseCaseFirst: ", IotaUseCaseFirst, "\n");
	fmt.Println("IotaUseCaseSecond: ", IotaUseCaseSecond, "\n");
	fmt.Println("IotaUseCaseThird: ", IotaUseCaseThird, "\n");
	fmt.Println("IotaUseCaseFourth: ", IotaUseCaseFourth, "\n");

	fmt.Println("IotaUseCase2First: ", IotaUseCase2First, "\n");
	fmt.Println("IotaUseCase2Second: ", IotaUseCase2Second, "\n");
	fmt.Println("IotaUseCase2Third: ", IotaUseCase2Third, "\n");
	fmt.Println("IotaUseCase2Fourth: ", IotaUseCase2Fourth, "\n");
}
