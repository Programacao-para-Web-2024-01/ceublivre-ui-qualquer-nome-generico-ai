package utils

import "fmt"

func Notify(email string, message string) {
    fmt.Printf("Enviando email para %s: %s\n", email, message)
}
