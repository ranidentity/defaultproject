package testing

import (
	"bytes"
	"defaultproject/model"
	"defaultproject/response"
	"defaultproject/status"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/book", GetBook)
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
			expectedBody: `{
    "code": 0,
    "data": [
        {
            "id": 1,
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z",
            "deleted_at": null,
            "title": "valid_book",
            "available_copies": 3
        }
    ],
    "msg": ""
}`,
		},
		{
			name:           "Invalid Book Title",
			requestBody:    map[string]string{"title": "unknown_book"},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":0,"data":null,"msg":""}`,
		},
		{
			name:           "Empty Request Body",
			requestBody:    map[string]string{},
			expectedStatus: http.StatusOK, // If no title, it shouldn't be an error per your logic
			expectedBody:   `{"code":0,"data":[{"id":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null,"title":"valid_book","available_copies":3},{"id":2,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null,"title":"valid_book_2","available_copies":3},{"id":3,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null,"title":"valid_book_3","available_copies":3}],"msg":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal request body to JSON
			body, _ := json.Marshal(tt.requestBody)

			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodPost, "/book", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			// Create a response recorder
			w := httptest.NewRecorder()
			// Create a new Gin context
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			// Call the handler
			GetBook(c)

			// Convert to struct for easier comparison
			// var expected, actual MockedGetBookResponse
			// json.Unmarshal([]byte(w.Body.String()), &expected)
			// json.Unmarshal(w.Body.Bytes(), &actual)

			// Validate response
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

type MockedBook struct {
	ID              int    `json:"ID"`
	Title           string `json:"title"`
	AvailableCopies int    `json:"availableCopies"`
}
type MockedGetBookResponse struct {
	Code int          `json:"code"`
	Data []model.Book `json:"data"`
	Msg  string       `json:"msg"`
}

// Controlled environment
var mockBookStoreGetBook = func(title string) (MockedGetBookResponse, error) {
	var data []model.Book
	data = append(data, model.Book{BaseModel: model.BaseModel{ID: 1}, Title: "valid_book", AvailableCopies: 3})
	data = append(data, model.Book{BaseModel: model.BaseModel{ID: 2}, Title: "valid_book_2", AvailableCopies: 3})
	data = append(data, model.Book{BaseModel: model.BaseModel{ID: 3}, Title: "valid_book_3", AvailableCopies: 3})
	if title == "valid_book" {
		return MockedGetBookResponse{
			Code: 0,
			Data: []model.Book{{BaseModel: model.BaseModel{ID: 1}, Title: "valid_book", AvailableCopies: 3}},
			Msg:  "",
		}, nil
	} else if title == "" {
		return MockedGetBookResponse{
			Code: 0,
			Data: data,
			Msg:  "",
		}, nil
	}
	return MockedGetBookResponse{}, nil
}

func GetBook(c *gin.Context) {
	var reqbody response.BookStoreRequest
	if err := c.ShouldBind(&reqbody); err != nil {
		c.Abort()
		return
	}
	if res, err := mockBookStoreGetBook(reqbody.Title); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}

func TestBorrowBook(t *testing.T) {

}
