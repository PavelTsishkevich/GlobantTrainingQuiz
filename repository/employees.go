package repository

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Address   string
	Position  string
}

type EmployeeRepository struct {
	employees map[int]*Employee
	id        int
}

func (employee *Employee) Print() {
	fmt.Println("Employee:", employee.Id, employee.FirstName, employee.LastName, employee.Position, employee.Address)
}

func NewRepo() *EmployeeRepository {
	return &EmployeeRepository{make(map[int]*Employee), 0}
}

func (repository *EmployeeRepository) Create(employee *Employee) {
	repository.id++
	currentEmployee := employee
	currentEmployee.Id = repository.id

	repository.employees[repository.id] = employee
}

func (repository *EmployeeRepository) Update(employee *Employee) {
	repository.employees[employee.Id] = employee
}

func (repository *EmployeeRepository) Delete(id int) {
	delete(repository.employees, id)
}

func (repository *EmployeeRepository) FindById(id int) *Employee {
	return repository.employees[id]
}

func (repository *EmployeeRepository) FindAll() []*Employee {
	result := make([]*Employee, 0)
	for i := 0; i <= repository.id; i++ {
		element, ok := repository.employees[i]
		if ok {
			result = append(result, element)
		}
	}
	return result
}

func (repository *EmployeeRepository) PrintAll() {
	for _, value := range repository.FindAll() {
		value.Print()
	}
}
