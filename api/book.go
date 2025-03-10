package api

import (
	"defaultproject/response"
	"defaultproject/serializer"
	"defaultproject/service"
	"defaultproject/status"
	"defaultproject/util"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	var reqbody response.BookStoreRequest
	if err := c.ShouldBind(&reqbody); err != nil {
		// shouldn't be any problem since it will either take Title or just none
		c.Abort()
		return
	}
	if res, err := service.BookStoreGetBook(reqbody.Title); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
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
	if res, err := service.BookStoreBorrowBook(reqbody.Title, reqbody.NameOfBorrower); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}

// TODO add proper response
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
	if res, err := service.BookStoreExtendLoan(reqbody.LoanId); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}

func ReturnBook(c *gin.Context) {
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
	if res, err := service.BookStoreReturn(reqbody.LoanId); err == nil {
		c.JSON(status.CodeOk, res)
	} else {
		c.JSON(status.CodeGeneralError, res)
	}
}
