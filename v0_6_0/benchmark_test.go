package v0_6_0_test

import (
	"testing"

	om_old "github.com/sttk/benchmarks_orderedmap/v0_5_0"
	orderedmap "github.com/sttk/benchmarks_orderedmap/v0_6_0"
)

type Foo struct {
	Bar string
	Baz int
}

func BenchmarkNew_OrderedMap_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		om := orderedmap.New[string, Foo]()
		_ = om
	}
}

func BenchmarkOld_OrderedMap_New(b *testing.B) {
	for i := 0; i < b.N; i++ {
		om := om_old.New[string, Foo]()
		_ = om
	}
}

func BenchmarkNew_OrderedMap_Store_newOneEntry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		om := orderedmap.New[string, Foo]()
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
}

func BenchmarkOld_OrderedMap_Store_newOneEntry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		om := om_old.New[string, Foo]()
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
}

func BenchmarkNew_OrderedMap_Store_newFiveEntries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		om := orderedmap.New[string, Foo]()
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
}

func BenchmarkOld_OrderedMap_Store_newFiveEntries(b *testing.B) {
	for i := 0; i < b.N; i++ {
		om := om_old.New[string, Foo]()
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
}

func BenchmarkNew_OrderedMap_Store_rewriteOneEntry(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
}

func BenchmarkOld_OrderedMap_Store_rewriteOneEntry(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
}

func BenchmarkNew_OrderedMap_Store_rewriteFiveEntries(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
}

func BenchmarkOld_OrderedMap_Store_rewriteFiveEntries(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
}

func BenchmarkNew_OrderedMap_Load_oneEntry(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		value, exists := om.Load("foo-0")
		_ = value
		_ = exists
	}
}

func BenchmarkOld_OrderedMap_Load_oneEntry(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		value, exists := om.Load("foo-0")
		_ = value
		_ = exists
	}
}

func BenchmarkNew_OrderedMap_Load_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v0, exists := om.Load("foo-0")
		v1, exists := om.Load("foo-1")
		v2, exists := om.Load("foo-2")
		v3, exists := om.Load("foo-3")
		v4, exists := om.Load("foo-4")
		_ = v0
		_ = v1
		_ = v2
		_ = v3
		_ = v4
		_ = exists
	}
}

func BenchmarkOld_OrderedMap_Load_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v0, exists := om.Load("foo-0")
		v1, exists := om.Load("foo-1")
		v2, exists := om.Load("foo-2")
		v3, exists := om.Load("foo-3")
		v4, exists := om.Load("foo-4")
		_ = v0
		_ = v1
		_ = v2
		_ = v3
		_ = v4
		_ = exists
	}
}

func BenchmarkNew_OrderedMap_Delete_oneEntry(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Delete("foo-0")
	}
}

func BenchmarkOld_OrderedMap_Delete_oneEntry(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Delete("foo-0")
	}
}

func BenchmarkNew_OrderedMap_Ldelete_oneEntry(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Ldelete("foo-0")
	}
}

func BenchmarkOld_OrderedMap_Ldelete_oneEntry(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Ldelete("foo-0")
	}
}

func BenchmarkNew_OrderedMap_Delete_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})
		om.Delete("foo-0")
		om.Delete("foo-1")
		om.Delete("foo-2")
		om.Delete("foo-3")
		om.Delete("foo-4")
	}
}

func BenchmarkOld_OrderedMap_Delete_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})
		om.Delete("foo-0")
		om.Delete("foo-1")
		om.Delete("foo-2")
		om.Delete("foo-3")
		om.Delete("foo-4")
	}
}

func BenchmarkNew_OrderedMap_Ldelete_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})
		om.Ldelete("foo-0")
		om.Ldelete("foo-1")
		om.Ldelete("foo-2")
		om.Ldelete("foo-3")
		om.Ldelete("foo-4")
	}
}

func BenchmarkOld_OrderedMap_Ldelete_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})
		om.Ldelete("foo-0")
		om.Ldelete("foo-1")
		om.Ldelete("foo-2")
		om.Ldelete("foo-3")
		om.Ldelete("foo-4")
	}
}

func BenchmarkNew_OrderedMap_IterateWithRange_oneEntry(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Range(func(k string, v Foo) bool {
			return true
		})
	}
}

func BenchmarkOld_OrderedMap_IterateWithRange_oneEntry(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Range(func(k string, v Foo) bool {
			return true
		})
	}
}

func BenchmarkNew_OrderedMap_IterateWithFront_oneEntry(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for ent := om.Front(); ent != nil; ent = ent.Next() {
			_ = ent.Key()
			_ = ent.Value()
		}
	}
}

func BenchmarkOld_OrderedMap_IterateWithFront_oneEntry(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for ent := om.Front(); ent != nil; ent = ent.Next() {
			_ = ent.Key()
			_ = ent.Value()
		}
	}
}

func BenchmarkNew_OrderedMap_IterateWithRange_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Range(func(k string, v Foo) bool {
			return true
		})
	}
}

func BenchmarkOld_OrderedMap_IterateWithRange_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Range(func(k string, v Foo) bool {
			return true
		})
	}
}

func BenchmarkNew_OrderedMap_IterateWithFront_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for ent := om.Front(); ent != nil; ent = ent.Next() {
			_ = ent.Key()
			_ = ent.Value()
		}
	}
}

func BenchmarkOld_OrderedMap_IterateWithFront_fiveEntries(b *testing.B) {
	b.StopTimer()
	om := om_old.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for ent := om.Front(); ent != nil; ent = ent.Next() {
			_ = ent.Key()
			_ = ent.Value()
		}
	}
}

func BenchmarkNew_OrderedMap_MarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  om := orderedmap.New[string, string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkOld_OrderedMap_MarshalJSON_empty(b *testing.B) {
  b.StopTimer()

  om := om_old.New[string, string]()

  b.StartTimer()
  for i := 0; i < b.N; i++ {
    bs, err := om.MarshalJSON()
    _ = bs
    _ = err
  }
}

func BenchmarkNew_OrderedMap_MarshalJSON_valueIsString(b *testing.B) {
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

func BenchmarkOld_OrderedMap_MarshalJSON_valueIsString(b *testing.B) {
  b.StopTimer()

  om := om_old.New[string, string]()
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

func BenchmarkNew_OrderedMap_MarshalJSON_valueIsStringPointer(b *testing.B) {
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

func BenchmarkOld_OrderedMap_MarshalJSON_valueIsStringPointer(b *testing.B) {
  b.StopTimer()

  v0 := "ABCD"
  v1 := "EFG"
  v2 := "HIJK"
  v3 := "LMN"
  v4 := "OPQR"

  om := om_old.New[string, *string]()
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

func BenchmarkNew_OrderedMap_MarshalJSON_valueIsInt(b *testing.B) {
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

func BenchmarkOld_OrderedMap_MarshalJSON_valueIsInt(b *testing.B) {
  b.StopTimer()

  om := om_old.New[string, int]()
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

func BenchmarkNew_OrderedMap_MarshalJSON_valueIsIntPointer(b *testing.B) {
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

func BenchmarkOld_OrderedMap_MarshalJSON_valueIsIntPointer(b *testing.B) {
  b.StopTimer()

  v0 := 12
  v1 := 34
  v2 := 56
  v3 := 78
  v4 := 9

  om := om_old.New[string, *int]()
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
type A1 struct {
  Flg bool
  Str string
}

type A2 struct {
  Num int
  Obj A1
}

func BenchmarkNew_OrderedMap_MarshalJSON_valueIsStruct(b *testing.B) {
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

func BenchmarkOld_OrderedMap_MarshalJSON_valueIsStruct(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := om_old.New[string, A2]()
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

func BenchmarkNew_OrderedMap_MarshalJSON_valueIsStructPointer(b *testing.B) {
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

func BenchmarkOld_OrderedMap_MarshalJSON_valueIsStructPointer(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := om_old.New[string, *A2]()
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

func BenchmarkNew_OrderedMap_MarshalJSON_valueIsAny(b *testing.B) {
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

func BenchmarkOld_OrderedMap_MarshalJSON_valueIsAny(b *testing.B) {
  b.StopTimer()

  v0 := A2{Num:12, Obj:A1{Flg:true, Str:"ABC"}}
  v1 := A2{Num:34, Obj:A1{Flg:true, Str:"DEF"}}
  v2 := A2{Num:56, Obj:A1{Flg:false, Str:"GH"}}
  v3 := A2{Num:78, Obj:A1{Flg:true, Str:"IJK"}}
  v4 := A2{Num:99, Obj:A1{Flg:false, Str:"LMN"}}

  om := om_old.New[string, any]()
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
