package mysql

type mockClient struct {
	realClient
	//mock go_sqlmock.
}
