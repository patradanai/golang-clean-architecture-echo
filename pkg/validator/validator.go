package validator

import (
	errs "movie-service/pkg/errors"
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
	trans    ut.Translator
)

func init() {
	Validate = validator.New()
}

func ValidateStruct[T any](structRule T) errs.Errors {
	if err := Validate.Struct(structRule); err != nil {
		var fieldErrors []errs.FieldError

		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				jsonFieldName := e.Field()
				if field, ok := reflect.TypeOf(&structRule).Elem().FieldByName(e.Field()); ok {
					if jsonTag, ok := field.Tag.Lookup("json"); ok {
						jsonFieldName = strings.Split(jsonTag, ",")[0]
					}
				}

				fieldErrors = append(fieldErrors, errs.FieldError{
					FieldName:   jsonFieldName,
					Description: e.Error(),
				})
			}
		}

		return errs.WrapFieldError(errs.ValidateFieldError, fieldErrors)
	}

	return nil

}
