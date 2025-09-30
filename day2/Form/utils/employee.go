package utils

import (
	"fmt"
	"strings"
)

type Employee struct {
	name            string // if instance variable start with small case they are private
	email           string
	department      string
	dob             string
	password        string
	confirmPassword string
	empNo           int
}

func NewEmployee(empNo int, name, email, password, confirmPassword, department, dob string) (*Employee, error) {

	return &Employee{
		empNo:           empNo,
		name:            strings.Trim(name, " "),
		email:           email,
		department:      strings.ToUpper(department),
		dob:             dob,
		password:        password,
		confirmPassword: confirmPassword,
	}, nil
}

//func (e *Employee) PopulateEmployee() []*Employee {
//	var employees []*Employee = make([]*Employee, 0)
//	employees = append(employees, NewEmployee(101, "Prady", "", "123456", "123456", "IT", "1999-01-01"))
//
//	return employees
//}

func (e *Employee) SetDepartment(department string) {
	e.department = strings.ToUpper(department)
}

// for handling naming conflict
// On the left → name: means the struct field.
// On the right → name means the constructor parameter.

func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e Employee) Display() {
	fmt.Printf("employee details are EmpNo - %d, Name - %s, Email - %s, Password - %s", e.empNo, e.name, e.email, e.password)
}

func (e Employee) GetEmpNo() int {
	return e.empNo
}

func (e Employee) GetEmail() string {
	return e.email
}
func (e Employee) GetPassword() string {
	return e.password
}
func (e Employee) GetConfirmPassword() string {
	return e.confirmPassword
}

func (e Employee) GetDepartment() string {
	return e.department
}
func (e Employee) GetDob() string {
	return e.dob
}

func (e *Employee) SetEmail(email string) {
	e.email = email
}

func (e *Employee) SetPassword(password string) {
	e.password = password
}
func (e *Employee) SetDob(dob string) {
	e.dob = dob
}
