package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	Validate   *validator.Validate
	Translator ut.Translator
}

func InitValidator() *Validator {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	return &Validator{
		Validate:   validate,
		Translator: trans,
	}
}

func TranslateError(err error, trans ut.Translator) (errs string) {
	if err == nil {
		return ""
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs += translatedErr.Error() + ","
	}
	return errs
}
