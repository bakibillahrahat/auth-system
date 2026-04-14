package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bakibillahrahat/auth-system/internal/database"
	"github.com/bakibillahrahat/auth-system/internal/models"
	"github.com/bakibillahrahat/auth-system/pkg/utils"
)

// type ResisterInput: Frontend will send a JSON payload with the following fields when a user tries to register. This struct is used to bind the incoming JSON data to a Go struct for easier processing.
type RegisterInput struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"  validate:"required"`
	Email     string `json:"email"      validate:"required,email" gorm:"unique,not null"` 
	Password  string `json:"password"   validate:"required,min=6" gorm:"not null"`
	Address   models.Address `json:"address" gorm:"embedded"`
	AvatarURL string `json:"avatar_url" validate:"omitempty,url"`
}

// Register API: This struct defines the expected input for the user registration endpoint. It includes fields for the user's first name, last name, email, password, address, and an optional avatar URL. The struct tags specify how the JSON payload should be mapped to the struct fields and include validation rules to ensure that the input data is in the correct format and meets certain criteria (e.g., email must be valid, password must be at least 6 characters long).

func Register(w http.ResponseWriter, r *http.Request) {
	// API response will be sent back to the client in JSON format, so we set the Content-Type header to application/json.
	w.Header().Set("Content-Type", "application/json")

	// 1. Parse the incoming JSON request body and bind it to the RegisterInput struct. If there is an error during parsing (e.g., invalid JSON format), we return a 400 Bad Request response with an error message.
	var input RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input formatted JSON"})
		return
	}

	// 2. Validate the input data using the validator package. If the validation fails (e.g., missing required fields, invalid email format), we return a 400 Bad Request response with an error message.
	if input.Email == "" || input.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Email and password are required"})
		return
	}

	// 3. Convert The password in hashed
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Could not secure password"})
		return
	}

	// 4. User Model for database
	user := models.User {
		FirstName: input.FirstName,
		LastName: input.LastName,
		Email: input.Email,
		Password: hashedPassword,
		Address: input.Address,
		AvatarURL: input.AvatarURL,
	}

	// 5. Data save on Database Using GORM
	if err := database.DB.Create(&user).Error; err != nil {
		// As Email field using gorm:'unique", So it get any duplicate email it will give error
		w.WriteHeader(http.StatusConflict) // 409 Confilict
		json.NewEncoder(w).Encode(map[string]string{"error": "Email already exists"})
	}

	// 6. If data save successfully. User will get Success Message
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registerd successfully!",
	})
}