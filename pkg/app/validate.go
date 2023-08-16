package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	transfer()
}

func transfer() {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	//注册翻译器
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
}

func BindJSONAndValidate(c *gin.Context, v interface{}) error {
	b := binding.Default(c.Request.Method, "application/json")
	if err := c.ShouldBindWith(v, b); err != nil {
		if err.Error() == "EOF" {
			return errors.New("请传入参数")
		}
		if ev, ok := err.(validator.ValidationErrors); ok {
			for _, errc := range ev {
				msg := errc.Translate(trans)
				return errors.New(msg)
			}
		} else {
			return errors.New(err.Error())
		}
	}
	return nil
}

func BindHeaderAndValidate(c *gin.Context, v interface{}) error {
	if err := c.ShouldBindHeader(v); err != nil {
		for _, errc := range err.(validator.ValidationErrors) {
			msg := errc.Translate(trans)
			return errors.New(msg)
		}
	}
	return nil
}
