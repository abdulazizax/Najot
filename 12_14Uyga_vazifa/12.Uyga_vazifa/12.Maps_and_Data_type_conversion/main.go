//  1. Oldindan belgilangan qoidalarga muvofiq foydalanuvchi
//     kiritishini ValidateUser digan method yarating.
//
// Qoidalar:
// Name empty bo'lishi kerak emas
// Name uzunligi kamida 6 belgigi bo'lishi kerak
// Age 0 dan kotta va 120 dan kichik bo'lishi
// Email empty bo'lishi kerak emas
// Email formatiga mos bolishi kerak (masalan example@domain.com)
//
// 2. Error slice yaratilgan holda barcha paydo bo'lgan errorlarni qoshib yuvorin
// 3. Foydalanuvchi ma'lumotlarni terminaldan oqib oling
// 4. Oqib oliniyatgan jarayonda errorlarni ohirida chiqarib berin
//
// Masalan:
// Name: asd
// Age: 123
// Email: ""

// Errors:
// Name: length cannot be less than a 6 characters
// Age: couldn't be more than 120
// Email: couldn't be empty

package main

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field string
	Error string
}

type User struct {
	Name     string
	Age      int
	Email    string
	Password string
}

func check(str string) bool {
	if strings.Count(str, "@") != 1 || strings.Count(str, ".") != 1 {
		return false
	}
	parts := strings.Split(str, "@")
	if len(parts) != 2 || strings.Count(parts[1], ".") != 1 {
		return false
	}
	return true
}

func ValidateUser(user User) []ValidationError {
	var errors []ValidationError

	// Validate Name
	if user.Name == "" {
		errors = append(errors, ValidationError{"Ism", "Ismni kiritishingiz shart"})
	} else if len(user.Name) < 6 {
		errors = append(errors, ValidationError{"Ism", "Ism kamida 6 belgidan iborat bo'lishi kerak"})
	}

	// Validate Age
	if user.Age < 0 || user.Age > 120 {
		errors = append(errors, ValidationError{"Yosh", "Yosh 0 dan katta va 120 dan kichik bo'lishi kerak"})
	}

	// Validate Email
	if user.Email == "" {
		errors = append(errors, ValidationError{"Email", "Emailni kiritishingiz shart"})
	} else {
		if !check(user.Email) {
			errors = append(errors, ValidationError{"Email", "To'g'ri email formatini kiriting (masalan example@domain.com)"})
		}
	}

	return errors
}

func main() {
	var user User

	fmt.Print("Ism: ")
	fmt.Scanln(&user.Name)

	fmt.Print("Yosh: ")
	fmt.Scanln(&user.Age)

	fmt.Print("Email: ")
	fmt.Scanln(&user.Email)

	errors := ValidateUser(user)

	if len(errors) > 0 {
		fmt.Println("\nXatolik(lar):")
		for _, err := range errors {
			fmt.Printf("%s: %s\n", err.Field, err.Error)
		}
	} else {
		fmt.Println("\nSiz muvaffaqiyatli ro'yxatdan o'tdingiz!")
	}
}
