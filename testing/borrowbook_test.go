package testing

import (
	"bytes"
	"defaultproject/model"
	"defaultproject/response"
	"defaultproject/serializer"
	"defaultproject/status"
	"defaultproject/util"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TESTING DATE ISSUE
var NowFunc = time.Now // Default to `time.Now`

func TestBorrowBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/borrow", GetBook)

	fixedTime := time.Date(2025, 3, 10, 12, 0, 0, 0, time.UTC)
	NowFunc = func() time.Time { return fixedTime } // Override NowFunc
	defer func() { NowFunc = time.Now }()

	tests := []struct {
		name           string
		requestBody    map[string]string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Input",
			requestBody:    map[string]string{"title": "valid_book", "name_of_borrower": "user1"},
			expectedStatus: http.StatusOK,
			expectedBody: fmt.Sprintf(`{
				"code":0,
				"data":{
						"id": 1,
						"created_at": "0001-01-01T00:00:00Z",
						"updated_at": "0001-01-01T00:00:00Z",
						"deleted_at": null,
						"name_of_borrower": "user1",
						"loan_date": "%s",
						"return_date": "%s",
						"book_return_on": null,
						"BookId": 1,
						"book": null
					},
				"msg":"Please return before %s"
			}`, fixedTime.Truncate(time.Minute).UTC().Format(time.RFC3339), fixedTime.Truncate(time.Minute).AddDate(0, 0, 10).UTC().Format(time.RFC3339), fixedTime.Truncate(time.Minute).AddDate(0, 0, 10).UTC()),
		},
		{
			name:           "Invalid Book Title",
			requestBody:    map[string]string{},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   DataSet["invalid_borrow_result"],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal request body to JSON
			body, _ := json.Marshal(tt.requestBody)
			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodPost, "/borrow", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			// Create a response recorder
			w := httptest.NewRecorder()
			// Create a new Gin context
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			// Call the handler
			BorrowBook(c)
			// Validate response
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func MockBookStoreBorrow(borrower string, title string) (MockedBookStoreResponse, error) {
	from := NowFunc().Truncate(time.Minute).UTC()
	to := from.AddDate(0, 0, 10) // Add 10 days
	if title == "valid_book" && borrower != "" {
		return MockedBookStoreResponse{
			Data: model.LoanDetail{BaseModel: model.BaseModel{ID: 1}, NameOfBorrower: borrower, BookId: 1, ReturnDate: to, LoanDate: from, BookReturnedOn: nil},
			Msg:  fmt.Sprintf("Please return before %s", to),
		}, nil
	} else if title == "" && borrower == "" {
		return MockedBookStoreResponse{}, nil
	}
	return MockedBookStoreResponse{}, nil
}

func BorrowBook(c *gin.Context) {
	var reqbody response.BookStoreBorrowBookRequest
	if err := c.ShouldBind(&reqbody); err != nil {
		translations := map[string]string{
			"Title.required":          "Title is required",
			"NameOfBorrower.required": "Name of borrower is required",
		}
		customErrors := util.HandleValidationErrors(err, translations)
		c.JSON(status.CodeGeneralError, serializer.ErrRequestFormat(status.CodeGeneralError, customErrors))
		c.Abort()
		return
	}
	if res, err := MockBookStoreBorrow(reqbody.NameOfBorrower, reqbody.Title); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}
