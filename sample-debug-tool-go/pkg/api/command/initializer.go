package command

import (
	"reflect"
	"sort"
	"strings"
)

type Handler interface {
	GetBaseCommandGroup() *Group
}

func InitCommand(group *Group) {
	for _, c := range group.Commands {
		c.Schema = createSchema(c)
	}
}

func createSchema(c *Command) *Schema {
	rt := reflect.TypeOf(c.Model)
	schema := &Schema{
		Fields: make([]*Field, 0, rt.NumField()),
	}
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		schema.Fields = append(schema.Fields, createField(&f))
	}
	return schema
}

func createField(f *reflect.StructField) *Field {
	field := prepareBaseField(f)
	field.Label = createLabel(f)
	field.Model = createModel(f)
	field.Required = required(f)
	field.Values = createValues(f)
	return &field
}

func prepareBaseField(f *reflect.StructField) Field {
	if typeTag := f.Tag.Get("type"); typeTag != "" {
		return typeFieldMap[FormFieldType(typeTag)]
	}
	// typeの指定がなければデフォルトの設定を使用
	return typeFieldMap[defaultTypeMap[f.Type.Name()]]
}

func createLabel(f *reflect.StructField) string {
	if label := f.Tag.Get("label"); label != "" {
		return label
	}
	// labelの指定がなければ変数名をlabelとして設定
	return f.Name
}

func createModel(f *reflect.StructField) string {
	// json名をmodel名として設定
	return f.Tag.Get("json")
}

func required(f *reflect.StructField) bool {
	// validate情報にrequiredが含まれるかで判定
	return strings.Contains(f.Tag.Get("validate"), "required")
}

func createValues(f *reflect.StructField) Values {
	// enum指定
	if enum := f.Tag.Get("enum"); enum != "" {
		return createEnumValues(enum)
	}
	// マスター指定
	if master := f.Tag.Get("master"); master != "" {
		return createMasterValues(master)
	}
	return Values{}
}

func createEnumValues(enumName string) Values {
	enumMap := EnumTypeMap[enumName]
	values := make(Values, 0, len(enumMap))
	for k, v := range enumMap {
		values = append(values, &Value{
			ID:    v,
			Name:  k,
			Value: v,
		})
	}
	// ID昇順にソート
	sort.Slice(values, func(i, j int) bool {
		return values[i].ID.(int32) < values[j].ID.(int32)
	})
	return values
}

var EnumTypeMap = map[string]map[string]int32{
	"RarityType": {
		"ノーマル": 1,
		"レア":   2,
		"Sレア":  3,
	},
}

func createMasterValues(tableName string) Values {
	var values Values
	switch tableName {
	case "Card":
		cards := GetCards()
		values = make(Values, 0, len(cards))
		for _, card := range cards {
			values = append(values, &Value{
				ID:    card.ID,
				Name:  card.Name,
				Value: card.ID,
			})
		}
	}
	return values
}

type Card struct {
	ID      string
	Name    string
	SkillID string
}

type Cards []*Card

func GetCards() Cards {
	// 実際にはDBやキャッシュ等から取得する
	return Cards{
		{
			ID:      "cardID1",
			Name:    "カード1",
			SkillID: "skillID1",
		},
		{
			ID:      "cardID2",
			Name:    "カード2",
			SkillID: "skillID2",
		},
		{
			ID:      "cardID3",
			Name:    "カード3",
			SkillID: "skillID3",
		},
	}
}
