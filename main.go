package main

import (
	"fmt"
	simpleconnection "study/feater_postgres/simple_connection"
	"study/feature1"
	"study/feature2"
)

func main() {
	fmt.Println("Hello GIT")
	feature1.Feature1()
	feature2.Feature2()

	simpleconnection.CheckConnection()
}
