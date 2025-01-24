package main

import "studentManagementSystem/api"

func main() {
	// @title 学生管理系统 API
	// @version 1.0
	// @description
	// @host localhost:8080
	// @BasePath
	api.InitRouter()
	/*err := utils.DealWithCSV("./student.csv", false)
	if err != nil {
		panic(err)
	}
	db := mapper.GetDatabase()
	fmt.Printf("%+v\n", db)*/
}
