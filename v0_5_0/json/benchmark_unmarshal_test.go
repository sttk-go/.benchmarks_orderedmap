package json_test

import (
  "encoding/json"
  "testing"

  orderedmap "github.com/sttk/benchmarks_orderedmap/v0_5_0"

  om_i "github.com/iancoleman/orderedmap"
  om_w "github.com/wk8/go-ordered-map/v2"
)

func BenchmarkOrderedMap_UnmarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{}`)

  om := orderedmap.New[string, string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkMap_UnmarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{}`)

  m := make(map[string]string)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := json.Unmarshal(bs, &m)
    _ = err
  }
}

func BenchmarkOmW_UnmarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{}`)

  om := om_w.New[string, string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOmI_UnmarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{}`)

  om := om_i.New()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOrderedMap_UnmarshalJSON_valueIsString(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":"ABCD",
    "bar":"EFG",
    "Baz":"HIJK",
    "qux":"LMN",
    "quux":OPQ"}`)

  om := orderedmap.New[string, string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkMap_UnmarshalJSON_valueIsString(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":"ABCD",
    "bar":"EFG",
    "Baz":"HIJK",
    "qux":"LMN",
    "quux":OPQ"}`)

  m := make(map[string]string)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := json.Unmarshal(bs, &m)
    _ = err
  }
}

func BenchmarkOmW_UnmarshalJSON_valueIsString(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":"ABCD",
    "bar":"EFG",
    "Baz":"HIJK",
    "qux":"LMN",
    "quux":OPQ"}`)

  om := om_w.New[string, string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOmI_UnmarshalJSON_valueIsString(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":"ABCD",
    "bar":"EFG",
    "Baz":"HIJK",
    "qux":"LMN",
    "quux":OPQ"}`)

  om := om_i.New()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOrderedMap_UnmarshalJSON_valueIsStringPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":"ABCD",
    "bar":"EFG",
    "Baz":"HIJK",
    "qux":"LMN",
    "quux":OPQ"}`)

  om := orderedmap.New[string, *string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkMap_UnmarshalJSON_valueIsStringPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":"ABCD",
    "bar":"EFG",
    "Baz":"HIJK",
    "qux":"LMN",
    "quux":OPQ"}`)

  m := make(map[string](*string))

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := json.Unmarshal(bs, &m)
    _ = err
  }
}

func BenchmarkOmW_UnmarshalJSON_valueIsStringPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":"ABCD",
    "bar":"EFG",
    "Baz":"HIJK",
    "qux":"LMN",
    "quux":OPQ"}`)

  om := om_w.New[string, *string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOmI_UnmarshalJSON_valueIsStringPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":"ABCD",
    "bar":"EFG",
    "Baz":"HIJK",
    "qux":"LMN",
    "quux":OPQ"}`)

  om := om_i.New()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOrderedMap_UnmarshalJSON_valueIsInt(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{"foo":12,"bar":34,"Baz":56,"qux":78,"quux":9}`)

  om := orderedmap.New[string, int]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkMap_UnmarshalJSON_valueIsInt(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{"foo":12,"bar":34,"Baz":56,"qux":78,"quux":9}`)

  m := make(map[string]int)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := json.Unmarshal(bs, &m)
    _ = err
  }
}

func BenchmarkOmW_UnmarshalJSON_valueIsInt(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{"foo":12,"bar":34,"Baz":56,"qux":78,"quux":9}`)

  om := om_w.New[string, int]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOmI_UnmarshalJSON_valueIsInt(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{"foo":12,"bar":34,"Baz":56,"qux":78,"quux":9}`)

  om := om_i.New()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOrderedMap_UnmarshalJSON_valueIsIntPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{"foo":12,"bar":34,"Baz":56,"qux":78,"quux":9}`)

  om := orderedmap.New[string, *int]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkMap_UnmarshalJSON_valueIsIntPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{"foo":12,"bar":34,"Baz":56,"qux":78,"quux":9}`)

  m := make(map[string](*int))

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := json.Unmarshal(bs, &m)
    _ = err
  }
}

func BenchmarkOmW_UnmarshalJSON_valueIsIntPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{"foo":12,"bar":34,"Baz":56,"qux":78,"quux":9}`)

  om := om_w.New[string, *int]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOmI_UnmarshalJSON_valueIsIntPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{"foo":12,"bar":34,"Baz":56,"qux":78,"quux":9}`)

  om := om_i.New()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

// type A1 struct {
//   Flg bool
//   Str string
// }
// 
// type A2 struct {
//   Num int
//   Obj A1
// }

func BenchmarkOrderedMap_UnmarshalJSON_valueIsStruct(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
    "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
    "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
    "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
    "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  om := orderedmap.New[string, A2]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkMap_UnmarshalJSON_valueIsStruct(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
    "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
    "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
    "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
    "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  m := make(map[string]A2)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := json.Unmarshal(bs, &m)
    _ = err
  }
}

func BenchmarkOmW_UnmarshalJSON_valueIsStruct(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
    "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
    "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
    "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
    "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  om := om_w.New[string, A2]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOmI_UnmarshalJSON_valueIsStruct(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
    "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
    "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
    "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
    "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
    "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  om := om_i.New()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOrderedMap_UnmarshalJSON_valueIsStructPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
  "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
  "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
  "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
  "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
  "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  om := orderedmap.New[string, *A2]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkMap_UnmarshalJSON_valueIsStructPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
  "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
  "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
  "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
  "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
  "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  m := make(map[string](*A2))

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := json.Unmarshal(bs, &m)
    _ = err
  }
}

func BenchmarkOmW_UnmarshalJSON_valueIsStructPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
  "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
  "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
  "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
  "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
  "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  om := om_w.New[string, *A2]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOmI_UnmarshalJSON_valueIsStructPointer(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
  "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
  "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
  "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
  "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
  "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  om := om_i.New()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOrderedMap_UnmarshalJSON_valueIsAny(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
  "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
  "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
  "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
  "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
  "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  om := orderedmap.New[string, any]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkMap_UnmarshalJSON_valueIsAny(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
  "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
  "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
  "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
  "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
  "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  m := make(map[string]any)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := json.Unmarshal(bs, &m)
    _ = err
  }
}

func BenchmarkOmW_UnmarshalJSON_valueIsAny(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
  "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
  "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
  "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
  "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
  "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  om := om_w.New[string, any]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}

func BenchmarkOmI_UnmarshalJSON_valueIsAny(b *testing.B) {
  b.StopTimer()

  bs := []byte(`{
  "foo":{"Num":12,"Obj":{"Flg":true,"Str":"ABC"}},
  "bar":{"Num":34,"Obj":{"Flg":true,"Str":"DEF"}},
  "Baz":{"Num":56,"Obj":{"Flg":false,"Str":"GH"}},
  "qux":{"Num":78,"Obj":{"Flg":true,"Str":"IJK"}},
  "quux":{"Num":99,"Obj":{"Flg":false,"Str":"LMN"}}}`)

  om := om_i.New()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    err := om.UnmarshalJSON(bs)
    _ = err
  }
}
