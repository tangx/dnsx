package qcloud

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Test_regexp(t *testing.T) {

	regex("*", "*.demo")
	regex("*", "*.temo")
	regex("temo", "*.temo")
}

func regex(record string, str string) {
	_record := strings.ReplaceAll(record, "*", `\\*`)
	pattern := fmt.Sprintf(`.*%s.*`, _record)
	re := regexp.MustCompile(pattern)

	if re.MatchString(str) {
		fmt.Printf("%s -> %s \n", str, record)
	}
}
