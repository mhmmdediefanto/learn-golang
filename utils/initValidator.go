package utils

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding" // Import ini penting
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
)

var (
	Trans ut.Translator
)

func InitValidator() {
	// 1. AMBIL instance validator bawaan Gin
	// Kita melakukan type assertion ke *validator.Validate
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// 2. Konfigurasi agar nama field di error message sesuai tag JSON
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 3. Setup Translator Bahasa Indonesia
		english := en.New()
		uni := ut.New(english, english)

		trans, _ := uni.GetTranslator("id")
		Trans = trans

		// 4. Register translation ke validator bawaan Gin
		idTranslations.RegisterDefaultTranslations(v, Trans)
	}
}
