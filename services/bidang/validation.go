package bidang

import (
	"agungdh.com/bpkad_disposisi_be/models"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate = validator.New()
var uni *ut.UniversalTranslator

func ValidateStoreUpdate(bidang models.Bidang) map[string]string {
	rules := map[string]string{
		"Bidang": "required",
	}

	validate.RegisterStructValidationMapRules(rules, models.Bidang{})

	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(bidang)

	var errors map[string]string = map[string]string{}

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errors[e.Field()] = e.Translate(trans)
		}
	}

	return errors
}
