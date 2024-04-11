package validation

import (
	"log"
	"regexp"

	"github.com/BoruTamena/UserManagement/models"
)

func NameValidation(username string) bool {
	regex, _ := regexp.Compile(`^[a-zA-Z0-9_]{3,20}$`)
	return regex.MatchString(username)

}

func PasswordValidation(password string) bool {
	regex, err := regexp.Compile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`)
	if err != nil {
		log.Println("Error compiling regex:", err)
		return false
	}
	return regex.MatchString(password)
}

func PhoneNumberValidation(phoneNumber string) bool {

	log.Println(phoneNumber)
	regex := regexp.MustCompile(`^\+2519\d{8}$`)
	return regex.MatchString(phoneNumber)

}

func EmailValidation(email string) bool {

	regex, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)

}

func AddressValidation(address string) bool {
	regex, _ := regexp.Compile(`[^0-9]`)
	return regex.MatchString(address)
}

func ValidateUser(user models.User) bool {

	if !NameValidation(user.UserName) {

		log.Println("Invalid username ")
		return false
	}
	if !EmailValidation(user.Email) {
		log.Println("Invalid Email")
		return false
	}

	if !PhoneNumberValidation(user.PhoneNumber) {
		log.Println("Invalid Phonenumber", user.PhoneNumber)
		return false
	}

	// if !PasswordValidation(user.Password) {
	// 	log.Println("Invalid Password")
	// 	return false
	// }

	if !AddressValidation(user.Address) {
		log.Println("Invalid Address")
		return false
	}

	return true
}
