package mock_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

func TestSomething(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedObject)

	testObj.On("DoSomething", 123).Return(true, nil)

	// targetFuncThatDoesSomethingWithObj(testObj)

	testObj.AssertExpectations(t)
}
