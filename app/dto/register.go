package dto

import (
	"fmt"
	"github.com/faneaatiku/auth_api/app/utils"
	"unicode"
)

type EmailAddress string

func (ea EmailAddress) Validate() error {
	return utils.ValidateEmail(ea.String())
}

func (ea EmailAddress) String() string {
	return string(ea)
}

type RegisterRequest struct {
	Email          EmailAddress `json:"email"`
	Password       string       `json:"password"`
	RepeatPassword string       `json:"repeat_password"`
}

func (r RegisterRequest) Validate() error {
	err := r.Email.Validate()
	if err != nil {
		return err
	}

	if r.Password != r.RepeatPassword {
		return fmt.Errorf("passwords do not match")
	}

	valid := r.isValidPassword(r.Password)
	if !valid {
		return fmt.Errorf("password must include at least 1 uppercase letter, 1 lowercase letter, one number and one cpecial character (?.,:;!@#$%%^&*)")
	}

	return nil
}

// validatePassword checks if the password meets the specified criteria:
// at least one uppercase letter, one lowercase letter, one digit,
// one special character, and a minimum length of 8 characters.
func (r RegisterRequest) isValidPassword(password string) bool {
	// Check length
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}
