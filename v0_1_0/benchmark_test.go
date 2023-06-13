package v0_1_0_test

import (
	"testing"

	orderedmap "github.com/sttk-go/.benchmarks_orderedmap/v0_1_0"

	om_c "github.com/cevaris/ordered_map"
	om_e "github.com/elliotchance/orderedmap/v2"
	om_i "github.com/iancoleman/orderedmap"
	om_w "github.com/wk8/go-ordered-map/v2"
)

type Foo struct {
	Bar string
	Baz int
}

func BenchmarkOrderedMap_New(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := orderedmap.New[string, Foo]()
		_ = om
	}
	b.StopTimer()
}

func BenchmarkMap_New(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]Foo)
		_ = m
	}
	b.StopTimer()
}

func BenchmarkOmE_New(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_e.NewOrderedMap[string, Foo]()
		_ = om
	}
	b.StopTimer()
}

func BenchmarkOmW_New(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_w.New[string, Foo]()
		_ = om
	}
	b.StopTimer()
}

func BenchmarkOmI_New(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_i.New()
		_ = om
	}
	b.StopTimer()
}

func BenchmarkOmC_New(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_c.NewOrderedMap()
		_ = om
	}
	b.StopTimer()
}

func BenchmarkOrderedMap_Store_newOneEntry(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := orderedmap.New[string, Foo]()
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkMap_Store_newOneEntry(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]Foo)
		m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}
	}
	b.StopTimer()
}

func BenchmarkOmE_Store_newOneEntry(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_e.NewOrderedMap[string, Foo]()
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkOmW_Store_newOneEntry(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_w.New[string, Foo]()
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkOmI_Store_newOneEntry(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_i.New()
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkOmC_Store_newOneEntry(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_c.NewOrderedMap()
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkOrderedMap_Store_newFiveEntries(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := orderedmap.New[string, Foo]()
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Store("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Store("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Store("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Store("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
	b.StopTimer()
}

func BenchmarkMap_Store_newFiveEntries(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]Foo)
		m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}
		m["foo-1"] = Foo{Bar: "bar-1", Baz: 1}
		m["foo-2"] = Foo{Bar: "bar-2", Baz: 2}
		m["foo-3"] = Foo{Bar: "bar-3", Baz: 3}
		m["foo-4"] = Foo{Bar: "bar-4", Baz: 4}
	}
	b.StopTimer()
}

func BenchmarkOmE_Store_newFiveEntries(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_e.NewOrderedMap[string, Foo]()
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
	b.StopTimer()
}

func BenchmarkOmW_Store_newFiveEntries(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_w.New[string, Foo]()
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
	b.StopTimer()
}

func BenchmarkOmI_Store_newFiveEntries(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_i.New()
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
	b.StopTimer()
}

func BenchmarkOmC_Store_newFiveEntries(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om := om_c.NewOrderedMap()
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
	b.StopTimer()
}

func BenchmarkOrderedMap_Store_rewriteOneEntry(b *testing.B) {
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkMap_Store_rewriteOneEntry(b *testing.B) {
	m := make(map[string]Foo)
	m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}
	}
	b.StopTimer()
}

func BenchmarkOmE_Store_rewriteOneEntry(b *testing.B) {
	om := om_e.NewOrderedMap[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkOmW_Store_rewriteOneEntry(b *testing.B) {
	om := om_w.New[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkOmI_Store_rewriteOneEntry(b *testing.B) {
	om := om_i.New()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkOmC_Store_rewriteOneEntry(b *testing.B) {
	om := om_c.NewOrderedMap()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	}
	b.StopTimer()
}

func BenchmarkOrderedMap_Store_rewriteFiveEntries(b *testing.B) {
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
	b.StopTimer()
}

func BenchmarkMap_Store_rewriteFiveEntries(b *testing.B) {
	m := make(map[string]Foo)
	m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}
	m["foo-1"] = Foo{Bar: "bar-1", Baz: 1}
	m["foo-2"] = Foo{Bar: "bar-2", Baz: 2}
	m["foo-3"] = Foo{Bar: "bar-3", Baz: 3}
	m["foo-4"] = Foo{Bar: "bar-4", Baz: 4}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}
		m["foo-1"] = Foo{Bar: "bar-1", Baz: 1}
		m["foo-2"] = Foo{Bar: "bar-2", Baz: 2}
		m["foo-3"] = Foo{Bar: "bar-3", Baz: 3}
		m["foo-4"] = Foo{Bar: "bar-4", Baz: 4}
	}
	b.StopTimer()
}

func BenchmarkOmE_Store_rewriteFiveEntries(b *testing.B) {
	om := om_e.NewOrderedMap[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
	b.StopTimer()
}

func BenchmarkOmW_Store_rewriteFiveEntries(b *testing.B) {
	om := om_w.New[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
	b.StopTimer()
}

func BenchmarkOmI_Store_rewriteFiveEntries(b *testing.B) {
	om := om_i.New()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
	b.StopTimer()
}

func BenchmarkOmC_Store_rewriteFiveEntries(b *testing.B) {
	om := om_c.NewOrderedMap()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
		om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
		om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
		om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
		om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})
	}
	b.StopTimer()
}

func BenchmarkOrderedMap_Load_oneEntry(b *testing.B) {
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		value, exists := om.Load("foo-0")
		_ = value
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkMap_Load_oneEntry(b *testing.B) {
	m := make(map[string]Foo)
	m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		value, exists := m["foo-0"]
		_ = value
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOmE_Load_oneEntry(b *testing.B) {
	om := om_e.NewOrderedMap[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		value, exists := om.Get("foo-0")
		_ = value
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOmW_Load_oneEntry(b *testing.B) {
	om := om_w.New[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		value, exists := om.Get("foo-0")
		_ = value
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOmI_Load_oneEntry(b *testing.B) {
	om := om_i.New()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		value, exists := om.Get("foo-0")
		_ = value
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOmC_Load_oneEntry(b *testing.B) {
	om := om_c.NewOrderedMap()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		value, exists := om.Get("foo-0")
		_ = value
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOrderedMap_Load_fiveEntries(b *testing.B) {
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
	b.StopTimer()
}

func BenchmarkMap_Load_fiveEntries(b *testing.B) {
	m := make(map[string]Foo)
	m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}
	m["foo-1"] = Foo{Bar: "bar-1", Baz: 1}
	m["foo-2"] = Foo{Bar: "bar-2", Baz: 2}
	m["foo-3"] = Foo{Bar: "bar-3", Baz: 3}
	m["foo-4"] = Foo{Bar: "bar-4", Baz: 4}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v0, exists := m["foo-0"]
		v1, exists := m["foo-1"]
		v2, exists := m["foo-2"]
		v3, exists := m["foo-3"]
		v4, exists := m["foo-4"]
		_ = v0
		_ = v1
		_ = v2
		_ = v3
		_ = v4
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOmE_Load_fiveEntries(b *testing.B) {
	om := om_e.NewOrderedMap[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v0, exists := om.Get("foo-0")
		v1, exists := om.Get("foo-1")
		v2, exists := om.Get("foo-2")
		v3, exists := om.Get("foo-3")
		v4, exists := om.Get("foo-4")
		_ = v0
		_ = v1
		_ = v2
		_ = v3
		_ = v4
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOmW_Load_fiveEntries(b *testing.B) {
	om := om_w.New[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v0, exists := om.Get("foo-0")
		v1, exists := om.Get("foo-1")
		v2, exists := om.Get("foo-2")
		v3, exists := om.Get("foo-3")
		v4, exists := om.Get("foo-4")
		_ = v0
		_ = v1
		_ = v2
		_ = v3
		_ = v4
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOmI_Load_fiveEntries(b *testing.B) {
	om := om_i.New()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v0, exists := om.Get("foo-0")
		v1, exists := om.Get("foo-1")
		v2, exists := om.Get("foo-2")
		v3, exists := om.Get("foo-3")
		v4, exists := om.Get("foo-4")
		_ = v0
		_ = v1
		_ = v2
		_ = v3
		_ = v4
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOmC_Load_fiveEntries(b *testing.B) {
	om := om_c.NewOrderedMap()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		v0, exists := om.Get("foo-0")
		v1, exists := om.Get("foo-1")
		v2, exists := om.Get("foo-2")
		v3, exists := om.Get("foo-3")
		v4, exists := om.Get("foo-4")
		_ = v0
		_ = v1
		_ = v2
		_ = v3
		_ = v4
		_ = exists
	}
	b.StopTimer()
}

func BenchmarkOrderedMap_IterateWithRange_oneEntry(b *testing.B) {
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		om.Range(func(k string, v Foo) bool {
			return true
		})
	}
	b.StopTimer()
}

func BenchmarkOrderedMap_IterateWithFront_oneEntry(b *testing.B) {
	om := orderedmap.New[string, Foo]()
	om.Store("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for ent := om.Front(); ent != nil; ent = ent.Next() {
			_ = ent.Key()
			_ = ent.Value()
		}
	}
	b.StopTimer()
}

func BenchmarkMap_Iterate_oneEntry(b *testing.B) {
	m := make(map[string]Foo)
	m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for k, v := range m {
			_ = k
			_ = v
		}
	}
	b.StopTimer()
}

func BenchmarkOmE_Iterate_oneEntry(b *testing.B) {
	om := om_e.NewOrderedMap[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for el := om.Front(); el != nil; el = el.Next() {
			_ = el.Key
			_ = el.Value
		}
	}
	b.StopTimer()
}

func BenchmarkOmW_Iterate_oneEntry(b *testing.B) {
	om := om_w.New[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for pair := om.Oldest(); pair != nil; pair = pair.Next() {
			_ = pair.Key
			_ = pair.Value
		}
	}
	b.StopTimer()
}

func BenchmarkOmI_Iterate_oneEntry(b *testing.B) {
	om := om_i.New()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range om.Keys() {
			v, _ := om.Get(k)
			_ = v
		}
	}
	b.StopTimer()
}

func BenchmarkOmC_Iterate_oneEntry(b *testing.B) {
	om := om_c.NewOrderedMap()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		iter := om.IterFunc()
		for kv, ok := iter(); ok; kv, ok = iter() {
			_ = kv.Key
			_ = kv.Value
		}
	}
	b.StopTimer()
}

func BenchmarkOrderedMap_IterateWithRange_fiveEntries(b *testing.B) {
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
	b.StopTimer()
}

func BenchmarkOrderedMap_IterateWithFront_fiveEntries(b *testing.B) {
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
	b.StopTimer()
}

func BenchmarkMap_Iterate_fiveEntries(b *testing.B) {
	m := make(map[string]Foo)
	m["foo-0"] = Foo{Bar: "bar-0", Baz: 0}
	m["foo-1"] = Foo{Bar: "bar-1", Baz: 1}
	m["foo-2"] = Foo{Bar: "bar-2", Baz: 2}
	m["foo-3"] = Foo{Bar: "bar-3", Baz: 3}
	m["foo-4"] = Foo{Bar: "bar-4", Baz: 4}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for k, v := range m {
			_ = k
			_ = v
		}
	}
	b.StopTimer()
}

func BenchmarkOmE_Iterate_fiveEntries(b *testing.B) {
	om := om_e.NewOrderedMap[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for el := om.Front(); el != nil; el = el.Next() {
			_ = el.Key
			_ = el.Value
		}
	}
	b.StopTimer()
}

func BenchmarkOmW_Iterate_fiveEntries(b *testing.B) {
	om := om_w.New[string, Foo]()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for pair := om.Oldest(); pair != nil; pair = pair.Next() {
			_ = pair.Key
			_ = pair.Value
		}
	}
	b.StopTimer()
}

func BenchmarkOmI_Iterate_fiveEntries(b *testing.B) {
	om := om_i.New()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range om.Keys() {
			v, _ := om.Get(k)
			_ = v
		}
	}
	b.StopTimer()
}

func BenchmarkOmC_Iterate_fiveEntries(b *testing.B) {
	om := om_c.NewOrderedMap()
	om.Set("foo-0", Foo{Bar: "bar-0", Baz: 0})
	om.Set("foo-1", Foo{Bar: "bar-1", Baz: 1})
	om.Set("foo-2", Foo{Bar: "bar-2", Baz: 2})
	om.Set("foo-3", Foo{Bar: "bar-3", Baz: 3})
	om.Set("foo-4", Foo{Bar: "bar-4", Baz: 4})

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		iter := om.IterFunc()
		for kv, ok := iter(); ok; kv, ok = iter() {
			_ = kv.Key
			_ = kv.Value
		}
	}
	b.StopTimer()
}
