package main

import (
	"fmt"
	"os"
	features "student-db/features"
) 
func displayChoices(){
	fmt.Println("1. Add Students")
	fmt.Println("2. Delete Students ")
	fmt.Println("3. update Students ")
	fmt.Println("4. Show All Students ")
	fmt.Println("5. getstats  ")
	fmt.Println("6. Exit")
}
func main() {
	var myclass features.Class
	myclass.NewClass()
	fmt.Println(myclass)
	fmt.Println(myclass.Students)


	for {
		displayChoices()
		var input int
		fmt.Scanln(&input)

		switch input {
		case 1:
			myclass.AddStudent()
		case 2:
			myclass.DeleteStudent()
		case 3:
			myclass.UpdateStudent()
		case 4:
			myclass.ShowStudents()
		case 5:
			myclass.GetStats()
		case 6:
			os.Exit(0)
		}
	}
}