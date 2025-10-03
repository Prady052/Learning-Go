package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"example.com/Form/utils"
)

var departments = map[string]int{
	"IT":        101,
	"HR":        102,
	"Finance":   103,
	"Marketing": 105,
	"Support":   106,
	"R&D":       107,
}

func main() {
	var employees []utils.Employee
	flag := true
	reader := bufio.NewReader(os.Stdin)

	utils.Greet()
	for flag {
		fmt.Println("\n1. Register Employee\n2. Print Employees\n3. Store to File\n4. update \n5. Delete\n 0. Exit\nEnter option: ")

		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			fmt.Println("Invalid option, try again")
			continue
		}

		switch option {
		case 1:
			var empNo int
			fmt.Print("Enter Employee No: ")
			_, err = fmt.Scan(&empNo)
			fmt.Scanln()
			// empNo validation
			if empNo < 0 || err != nil {
				fmt.Println("Invalid input, try again")
				continue
			}

			fmt.Print("Enter Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			//name validation
			if name == "" {
				fmt.Println("name cannot be empty")
				continue
			}

			fmt.Print("Enter Email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			// email validation
			//if email == "" {
			//	return nil, errors.New("email cannot be empty")
			//}

			emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
			match, _ := regexp.MatchString(emailRegex, email)
			if !match {
				fmt.Println("invalid email format")
				continue
			}

			fmt.Print("Enter Department: ")
			department, _ := reader.ReadString('\n')
			department = strings.TrimSpace(department)

			// department validation
			if departments[strings.ToUpper(department)] == 0 {
				fmt.Println("invalid department")
				continue
			}

			fmt.Print("Enter DOB (YYYY-MM-DD): ")
			dob, _ := reader.ReadString('\n')
			dob = strings.TrimSpace(dob)

			//Dob validation
			_, err := time.Parse("2006-01-02", dob)

			if err != nil {
				fmt.Println("invalid date format")
				continue
			}

			fmt.Print("Enter Password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			//password validation
			if password == "" {
				fmt.Println("password cannot be empty")
				continue
			}

			fmt.Print("Confirm Password: ")
			confirmPassword, _ := reader.ReadString('\n')
			confirmPassword = strings.TrimSpace(confirmPassword)

			if password != confirmPassword {
				fmt.Println("passwords do not match")
				continue
			}

			emp, err := utils.NewEmployee(empNo, name, email, password, confirmPassword, department, dob)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			employees = append(employees, *emp)
			fmt.Println("Employee registered successfully!")

		case 2:
			fmt.Println("\nAll Employees:")
			for _, e := range employees {
				fmt.Printf("EmpNo: %d, Name: %s, Email: %s, Dept: %s, DOB: %s\n",
					e.GetEmpNo(), e.GetName(), e.GetEmail(), e.GetDepartment(), e.GetDob())
			}
		case 3:
			os.WriteFile("employeeDetails.txt", []byte(fmt.Sprintf("%+v", employees)), 0644)

		case 4:
			fmt.Println("Enter the employee number to update: ")
			var empNo int
			_, err := fmt.Scan(&empNo)
			fmt.Scanln()
			if err != nil {
				fmt.Println("Invalid option, try again")
				continue
			}
			var data *utils.Employee
			for index, e := range employees {
				if e.GetEmpNo() == empNo {
					data = &employees[index]
				}
			}
			fmt.Println("Enter the field to update else keep blank to skip: ")
			var cn string
			fmt.Println("Enter name: ")
			cn, _ = reader.ReadString('\n')
			cn = strings.TrimSpace(cn)
			if cn != "" {
				data.SetName(cn)
			}
			var email string
			fmt.Println("Enter email: ")
			email, _ = reader.ReadString('\n')
			email = strings.TrimSpace(email)

			// email validation
			//if email == "" {
			//	return nil, errors.New("email cannot be empty")
			//}

			emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
			if email != "" {
				match, _ := regexp.MatchString(emailRegex, email)
				if !match {
					fmt.Println("invalid email format")
					continue
				}
				data.SetEmail(email)
			}
			var dept string
			fmt.Println("Enter department: ")
			dept, _ = reader.ReadString('\n')
			dept = strings.TrimSpace(dept)

			// department validation

			if dept != "" {
				if departments[strings.ToUpper(dept)] == 0 {
					fmt.Println("invalid department")
					continue
				}
				data.SetDepartment(dept)
			}
			var dob string
			fmt.Println("Enter dob: ")
			dob, _ = reader.ReadString('\n')
			dob = strings.TrimSpace(dob)

			//Dob validation

			if dob != "" {
				_, err = time.Parse("2006-01-02", dob)

				if err != nil {
					fmt.Println("invalid date format")
					continue
				}
				data.SetDob(dob)
			}
			var password string
			fmt.Println("Enter password: ")
			fmt.Scanln(&password)
			if password != "" {
				data.SetPassword(password)
			}

			fmt.Println("Employee updated successfully!")

		case 5:
			fmt.Println("Enter the employee number to delete: ")
			var empNo int
			_, err := fmt.Scan(&empNo)
			if err != nil {
				fmt.Println("Invalid option, try again")
				continue
			}

			var pos int
			for index, e := range employees {
				if e.GetEmpNo() == empNo {
					pos = index
					break
				}
			}
			copy(employees[pos:], employees[pos+1:]) // copy all elements from pos+1 to the end of the slice
			// to pos to len(employees)-1
			// and then delete the last element
			employees = employees[:len(employees)-1]
			fmt.Println("Employee deleted successfully!")
		case 0:
			flag = false
			fmt.Println("Exiting program...")

		default:
			fmt.Println("Invalid option, try again")
		}
	}
}
