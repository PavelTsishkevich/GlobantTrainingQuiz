package main

import (
	"GlobantTrainingQuiz/models"
	"GlobantTrainingQuiz/repository"
	"GlobantTrainingQuiz/util"
	"fmt"
)

func main() {
	stack := models.NewStack()
	stack.Push(1)
	fmt.Println(stack.Size())
	stack.Push(2)
	fmt.Println(stack.Size())
	stack.Push(3)
	fmt.Println(stack.Size())
	result, _ := stack.Pull()
	fmt.Println("First out", result)
	result, _ = stack.Pull()
	fmt.Println("Second out", result)
	result, _ = stack.Pull()
	fmt.Println("Third out", result)
	result, err := stack.Pull()
	fmt.Println(err)

	queue := models.NewQueue()
	queue.Push(1)
	fmt.Println(queue.Size())
	queue.Push(2)
	fmt.Println(queue.Size())
	queue.Push(3)
	fmt.Println(queue.Size())
	result, _ = queue.Pull()
	fmt.Println("First out", result)
	result, _ = queue.Pull()
	fmt.Println("Second out", result)
	result, _ = queue.Pull()
	fmt.Println("Third out", result)
	result, err = queue.Pull()
	fmt.Println(err)

	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 11

	result, err = util.Search(x, a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Index of element is", result)
	}

	firstEmployee := repository.Employee{FirstName: "Pavel", LastName: "Tsishkevich", Position: "Developer", Address: "Minsk"}
	secondEmployee := repository.Employee{FirstName: "Denis", LastName: "Sedun", Position: "QA", Address: "Astana"}
	thirdEmployee := repository.Employee{FirstName: "Dzmitry", LastName: "Prokurat", Position: "Leader", Address: "Minsk"}

	repo := repository.NewRepo()
	repo.Create(&firstEmployee)
	repo.Create(&secondEmployee)
	repo.Create(&thirdEmployee)
	repo.PrintAll()

	secondEmployee.Address = "Buenos Aires"
	repo.Update(&secondEmployee)
	repo.PrintAll()

	fmt.Println(repo.FindById(1))

	repo.Delete(2)
	repo.PrintAll()

}
