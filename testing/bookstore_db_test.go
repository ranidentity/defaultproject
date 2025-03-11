package testing

import (
	"bytes"
	"defaultproject/api"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDatabase(t *testing.T) *gorm.DB {
	// Use SQLite in-memory database for testing
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRESQL_DSN")), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func TestDbGetBookCase(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/book2", api.GetBook)

	// db := setupDatabase(t)
	// db.Create(&model.Book{Title: "valid_book", AvailableCopies: 3})

	tests := []struct {
		name           string
		requestBody    map[string]string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Book Title",
			requestBody:    map[string]string{"title": "valid_book"},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code": 0,"data": [{"id": 1,"created_at": "0001-01-01T00:00:00Z","updated_at": "0001-01-01T00:00:00Z","deleted_at": null,"title": "valid_book","available_copies": 3}],"msg": ""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal request body to JSON
			body, _ := json.Marshal(tt.requestBody)

			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodGet, "/book2", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			// Create a response recorder
			w := httptest.NewRecorder()
			// Create a new Gin context
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			// Call the handler
			api.GetBook(c) //inside this function will access DB

			// Convert to struct for easier comparison
			var actual MockedGetBookResponse
			err := json.Unmarshal(w.Body.Bytes(), &actual)
			if err != nil {
				t.Fatalf("Error unmarshaling response: %v", err)
			}
			var expected MockedGetBookResponse
			err = json.Unmarshal([]byte(tt.expectedBody), &expected)
			if err != nil {
				t.Fatalf("Error unmarshaling expected body: %v", err)
			}
			assert.Equal(t, expected, actual)
		})
	}
}
