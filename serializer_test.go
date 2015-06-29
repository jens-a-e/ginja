package ginja

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func NewTestApi() *Api {
	api := &Api{}
	return api
}

type TestItem struct {
	Name string `json:"name"`
}

var testItem = TestItem{
	Name: "A Name",
}

var testItemPayload = map[string]interface{}{
	"data": map[string]interface{}{
		"type": "testitem",
		"id":   "0",
		"attributes": map[string]interface{}{
			"name": "A Name",
		},
	},
}

func TestStoreRegister(t *testing.T) {
	Convey("Api can register arbitrary types", t, func() {
		api := NewTestApi()
		api.Register(TestItem{})

		So(api.types[reflect.TypeOf(TestItem{})], ShouldResemble, [2]string{"testitem", "testitems"})
		So(api.NameFor(TestItem{}), ShouldEqual, "testitem")
	})
}

func TestDocumentMarshalJSON(t *testing.T) {
	Convey("Empty document has data:null", t, func() {
		d := NewDocument()

		// So(d, ShouldImplement, (*json.Marshaler)(nil))

		payload, err := json.Marshal(&d)

		So(string(payload), ShouldEqual, `{"data":null}`)
		So(err, ShouldBeNil)

	})

}
