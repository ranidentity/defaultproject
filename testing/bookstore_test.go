package testing

import (
	"defaultproject/model"
	"defaultproject/response"
	"defaultproject/status"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockedBookStoreResponse struct {
	Code  int               `json:"code"`
	Data  interface{}       `json:"data"`
	Msg   string            `json:"msg"`
	Error map[string]string `json:"error,omitempty"`
}

func TestGetBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/book", GetBook)
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
			expectedBody:   DataSet["valid_data_expected_result"],
		},
		{
			name:           "Invalid Book Title",
			requestBody:    map[string]string{"title": "unknown_book"},
			expectedStatus: http.StatusOK,
			expectedBody:   DataSet["empty_result"],
		},
		{
			name:           "Empty Request Body",
			requestBody:    map[string]string{},
			expectedStatus: http.StatusOK, // If no title, it shouldn't be an error per your logic
			expectedBody:   DataSet["all_books"],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal request body to JSON
			// body, _ := json.Marshal(tt.requestBody)// Using GET
			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodGet, "/book?title="+tt.requestBody["title"], nil)
			req.Header.Set("Content-Type", "application/json")
			// Create a response recorder
			w := httptest.NewRecorder()
			// Create a new Gin context
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			// Call the handler
			GetBook(c)
			// Validate response
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

// Controlled environment
var mockBookStoreGetBook = func(title string) (MockedBookStoreResponse, error) {
	var data []model.Book
	data = append(data, model.Book{BaseModel: model.BaseModel{ID: 1}, Title: "valid_book", AvailableCopies: 3})
	data = append(data, model.Book{BaseModel: model.BaseModel{ID: 2}, Title: "valid_book_2", AvailableCopies: 3})
	data = append(data, model.Book{BaseModel: model.BaseModel{ID: 3}, Title: "valid_book_3", AvailableCopies: 3})
	if title == "valid_book" {
		return MockedBookStoreResponse{
			Code: 0,
			Data: []model.Book{{BaseModel: model.BaseModel{ID: 1}, Title: "valid_book", AvailableCopies: 3}},
			Msg:  "",
		}, nil
	} else if title == "" {
		return MockedBookStoreResponse{
			Code: 0,
			Data: data,
			Msg:  "",
		}, nil
	}
	return MockedBookStoreResponse{}, nil
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
