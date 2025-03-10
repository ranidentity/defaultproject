2. Implement the API with a simple in-memory storage (map or slice). You will need to populate some books during the start-up of the program
3. The Library should have at least the following objects. You may add in any new objects or fields as you deemed necessary to complete the program
o BookDetail
    1. Title (string): Unique identifier for the book.
    2. AvailableCopies (int): No of available copies of the book that can be loaned.
o LoanDetail
    1. NameOfBorrower (string): Name of borrower.
    2. LoanDate (date): Date where the book was borrowed.
    3. ReturnDate (date): Date where the book should be returned.

4. Expose the following RESTful endpoints:
o GET /Book to retrieve the detail and available copies of a book title.
o POST /Borrow to borrow a book (loan period: 4 weeks) and display the detail
of the loan.
o POST /Extend to extend the loan of the book (extend 3 weeks from return
date).
o Post /Return to return the book.
5. Return appropriate HTTP status codes for success and error scenarios.
6. Write at least one unit test for each endpoint using Go’s net/http/httptest package.