package encry

import (
	"regexp"
)

var emailRegexp = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?")
var phoneRegexp = regexp.MustCompile("\\d{3}\\d{8}|\\d{4}\\{7,8}")

func CheckEmail(email string) bool {
	return emailRegexp.MatchString(email)
}

func CheckPhone(phone string) bool {
	return phoneRegexp.MatchString(phone)
}
