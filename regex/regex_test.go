package regex

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var validate *validator.Validate

func TestRegexLib(t *testing.T) {
	validate = validator.New()
	registerRegex(validate)
	registerSp(validate)
	//myEmail := "joeybloggs.gmail.com"

	proxyList := "http://http://asdfasdf:asdadfa@1.1.1.1:90,asdfasdf"

	//email := "asdfadsf@asdf.com"
	//errs := validate.Var(email, "contains=tomsdf")

	errs := validate.Var(proxyList, "omitempty,sp=0x2C regex=https?://([\\w-]+:\\w+@)?(\\d+\\.?){4}:\\d+")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}
}

// 正则支持
func registerRegex(validate *validator.Validate) {
	err := validate.RegisterValidation("regex", func(fl validator.FieldLevel) bool {
		match, err := regexp.MatchString(fl.Param(), fl.Field().String())
		if err != nil {
			logrus.Errorf("validate: regex failed, pattern: %s, value: %s, err: %s", fl.Param(), fl.Field().String(), err)
		}
		return match
	})
	if err != nil {
		panic(err)
	}
}

// 切分字符串并校验每个元素
func registerSp(validate *validator.Validate) {
	// eg: sp=0x2C ip   0x2c 为 , 转义，表示 , 切分并对每个元素 ip 进行校验
	err := validate.RegisterValidation("sp", func(fl validator.FieldLevel) bool {
		groups := strings.SplitN(fl.Param(), " ", 2)
		if len(groups) < 2 {
			return false
		}
		// 第一段为分隔符，第二段为规则
		sep := groups[0]
		tag := groups[1]
		sp := strings.Split(fl.Field().String(), sep)
		if len(sp) == 0 {
			logrus.Errorf("validate: sp failed, no items found, param: %s, value: %s", fl.Param(), fl.Field().String())
			return false
		}
		for _, e := range sp {
			errs := validate.Var(e, tag)
			if errs != nil {
				logrus.Errorf("validate: sp failed, param: %s, value: %s, err: %s", fl.Param(), fl.Field().String(), errs)
				return false
			}
		}
		return true
	})
	if err != nil {
		panic(err)
	}
}
