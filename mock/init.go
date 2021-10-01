package mock

var Mock *MockCenter

func init() {
	Mock = &MockCenter{}
	Mock.Init()
}
