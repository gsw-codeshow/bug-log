/*
本代码由模版自动生成请勿擅自修改！！！！！！！！！！！！！

模版创建时间： 20201215
*/

package document

import (
	"encoding/json"
	"io/ioutil"

	"github.com/GoAdminGroup/go-admin/modules/logger"
)

var documentTable *DocumentTable

type DocumentTable struct {
	Suger map[int32][]*Suger
}

type JsonTable struct {
	Suger []*Suger `json:"suger"`
}

type Suger struct {
	SugerId int32 ` json:"SugerId" `

	SugerName string ` json:"SugerName" `

	SugerName_translatecn string ` json:"SugerName_translatecn" `

	UserType int32 ` json:"UserType" `

	Factory string ` json:"Factory" `

	ValueType int32 ` json:"ValueType" `

	Value float32 ` json:"Value" `

	Attr []float32 ` json:"attr" `

	Pivot int32 ` json:"pivot" `

	Picture string ` json:"Picture" `

	_Layer1Color string ` json:"_Layer1Color" `

	_Layer2Color string ` json:"_Layer2Color" `

	_Color string ` json:"_Color" `
}

func GetSugerStruct() map[int32][]*Suger {
	return documentTable.Suger
}

func GetSugerByID(id int32) []*Suger {
	return documentTable.Suger[id]
}

func LoadFileFromPath(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func DocumentTableInit(path string) (err error) {
	data := &JsonTable{}
	body, err := LoadFileFromPath(path)
	err = json.Unmarshal(body, data)
	if nil != err {
		logger.Errorf("DocumentTableInit Error:%s", err)
		return
	}
	documentTable = &DocumentTable{

		Suger: make(map[int32][]*Suger),
	}

	for _, value := range data.Suger {

		m := documentTable.Suger[value.SugerId]
		if m == nil {
			m = make([]*Suger, 0)
		}
		m = append(m, value)
		documentTable.Suger[value.SugerId] = m
	}

	return
}

func main() {
	return
}
