package main

import (
	"testing"
)

func TestGetFirstNumber(t *testing.T) {
	cases := []struct {
		testString    string
		expectedInt   string
		expectedIndex int
	}{
		{testString: "19321907321",
			expectedInt:   "1",
			expectedIndex: 0},
		{testString: "fourp783fiveseventhree",
			expectedInt:   "7",
			expectedIndex: 5},
		{testString: "foursix5eightfivezvnbsevenjcrzhxdzfb2",
			expectedInt:   "5",
			expectedIndex: 7},
		{testString: "sixfconesix6three1sixsix",
			expectedInt:   "6",
			expectedIndex: 11},
		{testString: "fconesixthreesixsix",
			expectedInt:   "",
			expectedIndex: -1},
	}
	for _, c := range cases {
		actualInt, actualIndex := GetFirstNumber(c.testString)
		if actualInt != c.expectedInt || actualIndex != c.expectedIndex {
			t.Errorf("getFirstNumber(%s) = %s, %d; expected %s, %d", c.testString, actualInt, actualIndex, c.expectedInt, c.expectedIndex)
		}
	}
}

func TestGetLastNumber(t *testing.T) {
	cases := []struct {
		testString    string
		expectedInt   string
		expectedIndex int
	}{
		{testString: "19321907321",
			expectedInt:   "1",
			expectedIndex: 10},
		{testString: "fourp783fiveseventhree",
			expectedInt:   "3",
			expectedIndex: 7},
		{testString: "foursix5eightfivezvnbsevenjcrzhxdzfb2",
			expectedInt:   "2",
			expectedIndex: 36},
		{testString: "sixfconesix6three1sixsix",
			expectedInt:   "1",
			expectedIndex: 17},
		{testString: "fconesixthreesixsix",
			expectedInt:   "",
			expectedIndex: -1},
	}
	for _, c := range cases {
		actualInt, actualIndex := GetLastNumber(c.testString)
		if actualInt != c.expectedInt || actualIndex != c.expectedIndex {
			t.Errorf("getFirstNumber(%s) = %s, %d; expected %s, %d", c.testString, actualInt, actualIndex, c.expectedInt, c.expectedIndex)
		}
	}
}

func TestGetFirstCharInteger(t *testing.T) {

	cases := []struct {
		testString    string
		expectedInt   string
		expectedIndex int
	}{
		{testString: "19321907321",
			expectedInt:   "",
			expectedIndex: -1},
		{testString: "fourp783fiveseventhree",
			expectedInt:   "4",
			expectedIndex: 0},
		{testString: "foursix5eightfivezvnbsevenjcrzhxdzfb2",
			expectedInt:   "4",
			expectedIndex: 0},
		{testString: "sixfconesix6three1sixsix",
			expectedInt:   "6",
			expectedIndex: 0},
		{testString: "2fconesix6three1sixsix",
			expectedInt:   "1",
			expectedIndex: 3},
	}

	for _, c := range cases {
		actualInt, actualIndex := GetFirstCharInteger(c.testString)
		if actualInt != c.expectedInt || actualIndex != c.expectedIndex {
			t.Errorf("getLastCharInteger(%s) = %s, %d; expected %s, %d", c.testString, actualInt, actualIndex, c.expectedInt, c.expectedIndex)
		}
	}
}

func TestGetLastCharInteger(t *testing.T) {

	cases := []struct {
		testString    string
		expectedInt   string
		expectedIndex int
	}{
		{testString: "19321907321",
			expectedInt:   "",
			expectedIndex: -1},
		{testString: "fourp783fiveseventhree",
			expectedInt:   "3",
			expectedIndex: 17},
		{testString: "foursix5eightfivezvnbsevenjcrzhxdzfb2",
			expectedInt:   "7",
			expectedIndex: 21},
		{testString: "sixfconesix6three1sixsix",
			expectedInt:   "6",
			expectedIndex: 21},
	}

	for _, c := range cases {
		actualInt, actualIndex := GetLastCharInteger(c.testString)
		if actualInt != c.expectedInt || actualIndex != c.expectedIndex {
			t.Errorf("getLastCharInteger(%s) = %s, %d; expected %s, %d", c.testString, actualInt, actualIndex, c.expectedInt, c.expectedIndex)
		}
	}
}
