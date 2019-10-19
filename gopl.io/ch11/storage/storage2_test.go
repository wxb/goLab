package storage

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		fmt.Println("++++++++")
		notifiedUser, notifiedMsg = user, msg
	}

	// ...simulate a 980MB-used condition...

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s",
			notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}

var echo01 = func() string {
	fmt.Println("01")
	return "01"
}

func TestEcho01(t *testing.T) {
	// saved := echo01
	// defer func() { echo01 = saved }()
	echo01 := func() string {
		fmt.Println("0101")
		return "0101"
	}
	echo01()
	fmt.Println("TestEcho01")
}

func TestEcho0101(t *testing.T) {

	echo01()
	fmt.Println("TestEcho0101")
}
