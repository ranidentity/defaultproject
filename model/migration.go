package model

func migration() {
	_ = DB.AutoMigrate(&Book{}, &LoanDetail{})
}
