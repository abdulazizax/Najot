package user

import (
    "strings"
)

type ValidationError struct {
    Field string
    Error string
}

func ValidateUser(user User) []ValidationError {
    var errors []ValidationError

    if user.Name == "" {
        errors = append(errors, ValidationError{"Ism", "Ismni kiritishingiz shart"})
    } else if len(user.Name) < 6 {
        errors = append(errors, ValidationError{"Ism", "Ism kamida 6 belgidan iborat bo'lishi kerak"})
    }

    if user.Age < 0 || user.Age > 120 {
        errors = append(errors, ValidationError{"Yosh", "Yosh 0 dan katta va 120 dan kichik bo'lishi kerak"})
    }

    if user.Email == "" {
        errors = append(errors, ValidationError{"Email", "Emailni kiritishingiz shart"})
    } else if !checkEmailFormat(user.Email) {
        errors = append(errors, ValidationError{"Email", "To'g'ri email formatini kiriting (masalan example@domain.com)"})
    }

    return errors
}

func checkEmailFormat(email string) bool {
    if strings.Count(email, "@") != 1 || strings.Count(email, ".") < 1 {
        return false
    }
    parts := strings.Split(email, "@")
    if len(parts) != 2 || strings.Count(parts[1], ".") < 1 {
        return false
    }
    return true
}
