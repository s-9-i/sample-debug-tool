package command

import "github.com/labstack/echo/v4"

type FieldType string

const (
	FieldTypeInput     FieldType = "input"
	FieldTypeCheckbox  FieldType = "checkbox"
	FieldTypeSelect    FieldType = "select"
	FieldTypeRadios    FieldType = "radios"
	FieldTypeChecklist FieldType = "checklist"
)

type InputType string

const (
	InputTypeNumber   InputType = "number"
	InputTypeText     InputType = "text"
	InputTypeDatetime InputType = "datetime-local"
)

type ValidatorType string

const (
	ValidatorTypeInteger ValidatorType = "integer"
	ValidatorTypeString  ValidatorType = "string"
)

type Schema struct {
	Fields []*Field `json:"fields"`
}

type Field struct {
	Type          FieldType     `json:"type"`
	InputType     InputType     `json:"inputType"`
	Label         string        `json:"label"`
	Model         string        `json:"model"`
	Required      bool          `json:"required"`
	ValidatorType ValidatorType `json:"validator"`
	Step          int32         `json:"step"`
	Values        Values        `json:"values"`
}

type Value struct {
	// 文字列, 数値どちらでも設定できるようにinterface{}型とする
	ID    interface{} `json:"id"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type Values []*Value

type Command struct {
	HandlerFunc echo.HandlerFunc `json:"-"`
	URL         string           `json:"url"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Model       interface{}      `json:"model"`
	Schema      *Schema          `json:"schema"`
}

type Commands []*Command

type Group struct {
	Name     string   `json:"name"`
	URL      string   `json:"url"`
	Commands Commands `json:"commands"`
}

type Groups []*Group

type FormFieldType string

const (
	FormFieldTypeNumber    FormFieldType = "number"
	FormFieldTypeText      FormFieldType = "text"
	FormFieldTypeCheckbox  FormFieldType = "checkbox"
	FormFieldTypeDatetime  FormFieldType = "datetime"
	FormFieldTypeSelect    FormFieldType = "select"
	FormFieldTypeRadios    FormFieldType = "radios"
	FormFieldTypeChecklist FormFieldType = "checklist"
)

var (
	typeFieldMap = map[FormFieldType]Field{
		FormFieldTypeNumber: {
			Type:          FieldTypeInput,
			InputType:     InputTypeNumber,
			ValidatorType: ValidatorTypeInteger,
		},
		FormFieldTypeText: {
			Type:          FieldTypeInput,
			InputType:     InputTypeText,
			ValidatorType: ValidatorTypeString,
		},
		FormFieldTypeCheckbox: {
			Type: FieldTypeCheckbox,
		},
		FormFieldTypeDatetime: {
			Type:      FieldTypeInput,
			InputType: InputTypeDatetime,
			Step:      1,
		},
		FormFieldTypeSelect: {
			Type:   FieldTypeSelect,
			Values: Values{},
		},
		FormFieldTypeRadios: {
			Type:   FieldTypeRadios,
			Values: Values{},
		},
		FormFieldTypeChecklist: {
			Type:   FieldTypeChecklist,
			Values: Values{},
		},
	}

	defaultTypeMap = map[string]FormFieldType{
		"int32":  FormFieldTypeNumber,
		"int64":  FormFieldTypeNumber,
		"string": FormFieldTypeText,
		"bool":   FormFieldTypeCheckbox,
	}
)
