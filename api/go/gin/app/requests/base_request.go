package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type BaseRequest interface {
	//
}

func Validate(s BaseRequest) error {
	validate := validator.New()
	return validate.Struct(s)
}

func NewUnprocessableEntityException(err error, s BaseRequest, ctx *gin.Context) {
	jsonFields := getStructJsonFields(s)
	fieldErrors := err.(validator.ValidationErrors)

	errors := make(map[string]string)
	for _, e := range fieldErrors {
		key := jsonFields[e.Field()]

		errors[key+"."+e.Tag()] = getTagText(key, e.Tag(), e.Param(), jsonFields)
	}

	ctx.JSON(422, errors)
}

func getStructJsonFields(s BaseRequest) map[string]string {
	jsonFields := make(map[string]string)

	reflection := reflect.ValueOf(s)
	for i := 0; i < reflection.Type().NumField(); i++ {
		field := reflection.Type().Field(i)
		fieldName := field.Name
		jsonTag := field.Tag.Get("json")

		switch jsonTag {
		case "-":
			continue

		case "":
			jsonFields[fieldName] = fieldName

		default:
			jsonFields[fieldName] = jsonTag
		}
	}

	return jsonFields
}

func getTagText(field string, tag string, params string, jsonFields map[string]string) string {
	variables := strings.Split(params, ",")
	switch tag {
	case "required":
		return fmt.Sprintf("O campo %s é obrigatório.", field)

	case "min":
		return fmt.Sprintf("O campo %s deve conter no mínimo %s caracteres", field, variables[0])

	case "max":
		return fmt.Sprintf("O campo %s deve conter no máximo %s caracteres", field, variables[0])

	case "eqfield":
		return fmt.Sprintf("Os campos %s e %s devem conter valores iguais", field, jsonFields[variables[0]])
	default:
		return tag
	}
}
