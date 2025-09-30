package utils

import (
	"fmt"
	"regexp"
)

// function first letter must be capital

func ValidateEmployee(emp Employee) error {
	if emp.name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	// Simple email validation regex
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, emp.email)
	if !match {
		return fmt.Errorf("invalid email format")
	}

	if emp.password == "" {
		return fmt.Errorf("password cannot be empty")
	}

	if emp.password != emp.confirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	if emp.empNo <= 0 {
		return fmt.Errorf("Employee number must be a positive integer")
	}

	return nil
}

