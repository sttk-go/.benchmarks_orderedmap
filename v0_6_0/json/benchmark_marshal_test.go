package json_test

import (
  "encoding/json"
  "testing"

  orderedmap "github.com/sttk/benchmarks_orderedmap/v0_6_0"

  om_i "github.com/iancoleman/orderedmap"
  om_w "github.com/wk8/go-ordered-map/v2"
)

func BenchmarkOrderedMap_MarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  om := orderedmap.New[string, string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkMap_MarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  m := make(map[string]string)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := json.Marshal(m)
    _ = bs
    _ = err
  }
}

func BenchmarkOmW_MarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  om := om_w.New[string, string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOmI_MarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  om := om_i.New()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOrderedMap_MarshalJSON_valueIsString(b *testing.B) {
  b.StopTimer()

  om := orderedmap.New[string, string]()
  om.Store("foo", "ABCD")
  om.Store("bar", "EFG")
  om.Store("baz", "HIJK")
  om.Store("qux", "LMN")
  om.Store("quux", "OPQ")

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkMap_MarshalJSON_valueIsString(b *testing.B) {
  b.StopTimer()

  m := make(map[string]string)
  m["foo"] = "ABCD"
  m["bar"] = "EFG"
  m["baz"] = "HIJK"
  m["qux"] = "LMN"
  m["quux"] = "OPQR"

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := json.Marshal(m)
    _ = bs
    _ = err
  }
}

func BenchmarkOmW_MarshalJSON_valueIsString(b *testing.B) {
  b.StopTimer()

  om := om_w.New[string, string]()
  om.Set("foo", "ABCD")
  om.Set("bar", "EFG")
  om.Set("baz", "HIJK")
  om.Set("qux", "LMN")
  om.Set("quux", "OPQ")

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOmI_MarshalJSON_valueIsString(b *testing.B) {
  b.StopTimer()

  om := om_i.New()
  om.Set("foo", "ABCD")
  om.Set("bar", "EFG")
  om.Set("baz", "HIJK")
  om.Set("qux", "LMN")
  om.Set("quux", "OPQ")

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOrderedMap_MarshalJSON_valueIsStringPointer(b *testing.B) {
  b.StopTimer()

  v0 := "ABCD"
  v1 := "EFG"
  v2 := "HIJK"
  v3 := "LMN"
  v4 := "OPQR"

  om := orderedmap.New[string, *string]()
  om.Store("foo", &v0)
  om.Store("bar", &v1)
  om.Store("Baz", &v2)
  om.Store("qux", &v3)
  om.Store("quux", &v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkMap_MarshalJSON_valueIsStringPointer(b *testing.B) {
  b.StopTimer()

  v0 := "ABCD"
  v1 := "EFG"
  v2 := "HIJK"
  v3 := "LMN"
  v4 := "OPQR"

  m := make(map[string](*string))
  m["foo"] = &v0
  m["bar"] = &v1
  m["baz"] = &v2
  m["qux"] = &v3
  m["quux"] = &v4

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := json.Marshal(m)
    _ = bs
    _ = err
  }
}

func BenchmarkOmW_MarshalJSON_valueIsStringPointer(b *testing.B) {
  b.StopTimer()

  v0 := "ABCD"
  v1 := "EFG"
  v2 := "HIJK"
  v3 := "LMN"
  v4 := "OPQR"

  om := om_w.New[string, *string]()
  om.Set("foo", &v0)
  om.Set("bar", &v1)
  om.Set("Baz", &v2)
  om.Set("qux", &v3)
  om.Set("quux", &v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOmI_MarshalJSON_valueIsStringPointer(b *testing.B) {
  b.StopTimer()

  v0 := "ABCD"
  v1 := "EFG"
  v2 := "HIJK"
  v3 := "LMN"
  v4 := "OPQR"

  om := om_i.New()
  om.Set("foo", &v0)
  om.Set("bar", &v1)
  om.Set("Baz", &v2)
  om.Set("qux", &v3)
  om.Set("quux", &v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOrderedMap_MarshalJSON_valueIsInt(b *testing.B) {
  b.StopTimer()

  om := orderedmap.New[string, int]()
  om.Store("foo", 12)
  om.Store("bar", 34)
  om.Store("baz", 56)
  om.Store("qux", 78)
  om.Store("quux", 9)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkMap_MarshalJSON_valueIsInt(b *testing.B) {
  b.StopTimer()

  m := make(map[string]int)
  m["foo"] = 12
  m["bar"] = 34
  m["baz"] = 56
  m["qux"] = 78
  m["quux"] = 9

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := json.Marshal(m)
    _ = bs
    _ = err
  }
}

func BenchmarkOmW_MarshalJSON_valueIsInt(b *testing.B) {
  b.StopTimer()

  om := om_w.New[string, int]()
  om.Set("foo", 12)
  om.Set("bar", 34)
  om.Set("baz", 56)
  om.Set("qux", 78)
  om.Set("quux", 9)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOmI_MarshalJSON_valueIsInt(b *testing.B) {
  b.StopTimer()

  om := om_i.New()
  om.Set("foo", 12)
  om.Set("bar", 34)
  om.Set("baz", 56)
  om.Set("qux", 78)
  om.Set("quux", 9)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOrderedMap_MarshalJSON_valueIsIntPointer(b *testing.B) {
  b.StopTimer()

  v0 := 12
  v1 := 34
  v2 := 56
  v3 := 78
  v4 := 9

  om := orderedmap.New[string, *int]()
  om.Store("foo", &v0)
  om.Store("bar", &v1)
  om.Store("Baz", &v2)
  om.Store("qux", &v3)
  om.Store("quux", &v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkMap_MarshalJSON_valueIsIntPointer(b *testing.B) {
  b.StopTimer()

  v0 := 12
  v1 := 34
  v2 := 56
  v3 := 78
  v4 := 9

  m := make(map[string](*int))
  m["foo"] = &v0
  m["bar"] = &v1
  m["baz"] = &v2
  m["qux"] = &v3
  m["quux"] = &v4

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := json.Marshal(m)
    _ = bs
    _ = err
  }
}

func BenchmarkOmW_MarshalJSON_valueIsIntPointer(b *testing.B) {
  b.StopTimer()

  v0 := 12
  v1 := 34
  v2 := 56
  v3 := 78
  v4 := 9

  om := om_w.New[string, *int]()
  om.Set("foo", &v0)
  om.Set("bar", &v1)
  om.Set("Baz", &v2)
  om.Set("qux", &v3)
  om.Set("quux", &v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOmI_MarshalJSON_valueIsIntPointer(b *testing.B) {
  b.StopTimer()

  v0 := 12
  v1 := 34
  v2 := 56
  v3 := 78
  v4 := 9

  om := om_i.New()
  om.Set("foo", &v0)
  om.Set("bar", &v1)
  om.Set("Baz", &v2)
  om.Set("qux", &v3)
  om.Set("quux", &v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

type A1 struct {
  Flg bool
  Str string
}

type A2 struct {
  Num int
  Obj A1
}

func BenchmarkOrderedMap_MarshalJSON_valueIsStruct(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := orderedmap.New[string, A2]()
  om.Store("foo", v0)
  om.Store("bar", v1)
  om.Store("Baz", v2)
  om.Store("qux", v3)
  om.Store("quux", v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkMap_MarshalJSON_valueIsStruct(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  m := make(map[string]A2)
  m["foo"] = v0
  m["bar"] = v1
  m["Baz"] = v2
  m["qux"] = v3
  m["quux"] = v4

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := json.Marshal(m)
    _ = bs
    _ = err
  }
}

func BenchmarkOmW_MarshalJSON_valueIsStruct(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := om_w.New[string, A2]()
  om.Set("foo", v0)
  om.Set("bar", v1)
  om.Set("Baz", v2)
  om.Set("qux", v3)
  om.Set("quux", v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOmI_MarshalJSON_valueIsStruct(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := om_i.New()
  om.Set("foo", v0)
  om.Set("bar", v1)
  om.Set("Baz", v2)
  om.Set("qux", v3)
  om.Set("quux", v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOrderedMap_MarshalJSON_valueIsStructPointer(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := orderedmap.New[string, *A2]()
  om.Store("foo", &v0)
  om.Store("bar", &v1)
  om.Store("Baz", &v2)
  om.Store("qux", &v3)
  om.Store("quux", &v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkMap_MarshalJSON_valueIsStructPointer(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  m := make(map[string](*A2))
  m["foo"] = &v0
  m["bar"] = &v1
  m["Baz"] = &v2
  m["qux"] = &v3
  m["quux"] = &v4

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := json.Marshal(m)
    _ = bs
    _ = err
  }
}

func BenchmarkOmW_MarshalJSON_valueIsStructPointer(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := om_w.New[string, *A2]()
  om.Set("foo", &v0)
  om.Set("bar", &v1)
  om.Set("Baz", &v2)
  om.Set("qux", &v3)
  om.Set("quux", &v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOmI_MarshalJSON_valueIsStructPointer(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := om_i.New()
  om.Set("foo", &v0)
  om.Set("bar", &v1)
  om.Set("Baz", &v2)
  om.Set("qux", &v3)
  om.Set("quux", &v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOrderedMap_MarshalJSON_valueIsAny(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := orderedmap.New[string, any]()
  om.Store("foo", v0)
  om.Store("bar", v1)
  om.Store("Baz", v2)
  om.Store("qux", v3)
  om.Store("quux", v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkMap_MarshalJSON_valueIsAny(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  m := make(map[string]any)
  m["foo"] = v0
  m["bar"] = v1
  m["Baz"] = v2
  m["qux"] = v3
  m["quux"] = v4

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := json.Marshal(m)
    _ = bs
    _ = err
  }
}

func BenchmarkOmW_MarshalJSON_valueIsAny(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := om_w.New[string, any]()
  om.Set("foo", v0)
  om.Set("bar", v1)
  om.Set("Baz", v2)
  om.Set("qux", v3)
  om.Set("quux", v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOmI_MarshalJSON_valueIsAny(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := om_i.New()
  om.Set("foo", v0)
  om.Set("bar", v1)
  om.Set("Baz", v2)
  om.Set("qux", v3)
  om.Set("quux", v4)

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}
