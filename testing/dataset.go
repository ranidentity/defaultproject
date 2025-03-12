package testing

var (
	DataSet = map[string]string{
		"valid_data_expected_result": `{
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
		"all_books": `{"code":0,"data":[{"id":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null,"title":"valid_book","available_copies":3},{"id":2,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null,"title":"valid_book_2","available_copies":3},{"id":3,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null,"title":"valid_book_3","available_copies":3}],"msg":""}`,
		"valid_borrow_book_result": `{
			"code":0,
			"data":{
					"id": 1,
					"created_at": "2025-03-12T10:10:21.897682+08:00",
					"updated_at": "2025-03-12T10:10:21.897682+08:00",
					"deleted_at": null,
					"name_of_borrower": "user1",
					"loan_date": "2025-03-12T02:10:00Z",
					"return_date": "2025-03-22T02:10:00Z",
					"book_return_on": null,
					"BookId": 1,
					"book": null
				}
		}`,
		"invalid_borrow_result": `{
			"code": 500,
			"data": null,
			"msg": "",
			"error": {
				"NameOfBorrower": "Name of borrower is required",
				"Title": "Title is required"
			}
		}`,
		"valid_loan_book": `{
			"code": 0,
			"data": null,
			"msg": "Successfully extended loan for another 3 weeks"
		}`,
		"invalid_loan_id": `{
			"code": 500,
			"data": null,
			"msg": "",
			"error": {
				"LoanId": "Loan id is required"
			}
		}`,
		"valid_return_book": `{
			"code": 0,
			"data": null,
			"msg": "Book returned"
		}`,
		"invalid_return_book": `{
			"code": 0,
			"data": null,
			"msg": "failed to lock and fetch loan: record not found"
		}`,
		"empty_result": `{"code":0,"data":null,"msg":""}`,
	}
)
