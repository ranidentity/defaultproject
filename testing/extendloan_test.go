package testing

import (
	"bytes"
	"defaultproject/response"
	"defaultproject/serializer"
	"defaultproject/status"
	"defaultproject/util"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExtendLoan(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/Extend", GetBook)
	tests := []struct {
		name           string
		requestBody    map[string]uint
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Input",
			requestBody:    map[string]uint{"loan_id": 1},
			expectedStatus: http.StatusOK,
			expectedBody:   DataSet["valid_loan_book"],
		},
		{
			name:           "Invalid Loan Id",
			requestBody:    map[string]uint{},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   DataSet["invalid_loan_id"],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal request body to JSON
			body, _ := json.Marshal(tt.requestBody)
			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodPost, "/Extend", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			// Create a response recorder
			w := httptest.NewRecorder()
			// Create a new Gin context
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			// Call the handler
			ExtendLoan(c)
			// Validate response
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}

func MockExtendLoan(loan_id int) (MockedBookStoreResponse, error) {
	if loan_id == 1 {
		return MockedBookStoreResponse{
			Msg: "Successfully extended loan for another 3 weeks",
		}, nil
	}
	return MockedBookStoreResponse{}, nil
}

func ExtendLoan(c *gin.Context) {
	var reqbody response.BookStoreExtendLoanRequest
	if err := c.ShouldBind(&reqbody); err != nil {
		translations := map[string]string{
			"LoanId.required": "Loan id is required",
		}
		customErrors := util.HandleValidationErrors(err, translations)
		c.JSON(status.CodeGeneralError, serializer.ErrRequestFormat(status.CodeGeneralError, customErrors))
		c.Abort()
		return
	}
	if res, err := MockExtendLoan(int(reqbody.LoanId)); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}
