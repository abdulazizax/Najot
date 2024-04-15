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
    "homework/user"
)

func main() {
    var newUser user.User

    fmt.Print("Ism: ")
    fmt.Scanln(&newUser.Name)

    fmt.Print("Yosh: ")
    fmt.Scanln(&newUser.Age)

    fmt.Print("Email: ")
    fmt.Scanln(&newUser.Email)

    validationErrors := user.ValidateUser(newUser)

    if len(validationErrors) > 0 {
        fmt.Println("\nXatolik(lar):")
        for _, err := range validationErrors {
            fmt.Printf("%s: %s\n", err.Field, err.Error)
        }
    } else {
        fmt.Println("\nSiz muvaffaqiyatli ro'yxatdan o'tdingiz!")
    }
}
