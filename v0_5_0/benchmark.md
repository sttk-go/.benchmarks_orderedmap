# Benchmark of orderedmap v0.5.0

This measurement compared [github.com/sttk/orderedmap](https://github.com/sttk/orderedmap-go) v0.5.0 with v0.4.0

> BenchmarkNew_* ... v0.5.0, BenchmarkOld_* ... v0.4.0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_orderedmap/v0_5_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkNew_OrderedMap_New-12                             	233890573	         4.988 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_New-12                             	238625810	         4.984 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Store_newOneEntry-12               	 8278531	       142.9 ns/op	     320 B/op	       3 allocs/op
BenchmarkOld_OrderedMap_Store_newOneEntry-12               	 8313032	       144.0 ns/op	     320 B/op	       3 allocs/op
BenchmarkNew_OrderedMap_Store_newFiveEntries-12            	 2725975	       439.9 ns/op	     576 B/op	       7 allocs/op
BenchmarkOld_OrderedMap_Store_newFiveEntries-12            	 2762152	       439.8 ns/op	     576 B/op	       7 allocs/op
BenchmarkNew_OrderedMap_Store_rewriteOneEntry-12           	191053759	         6.260 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Store_rewriteOneEntry-12           	186208292	         6.402 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Store_rewriteFiveEntries-12        	17172411	        69.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Store_rewriteFiveEntries-12        	17335756	        66.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Load_oneEntry-12                   	312376158	         3.828 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Load_oneEntry-12                   	331935686	         3.592 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Load_fiveEntries-12                	20690523	        57.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Load_fiveEntries-12                	20949424	        56.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Delete_oneEntry-12                 	15266077	        77.29 ns/op	      64 B/op	       1 allocs/op
BenchmarkOld_OrderedMap_Delete_oneEntry-12                 	15174847	        76.79 ns/op	      64 B/op	       1 allocs/op
BenchmarkNew_OrderedMap_Ldelete_oneEntry-12                	57612496	        20.83 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Ldelete_oneEntry-12                	59197850	        20.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Delete_fiveEntries-12              	 2353731	       500.5 ns/op	     320 B/op	       5 allocs/op
BenchmarkOld_OrderedMap_Delete_fiveEntries-12              	 2416612	       492.7 ns/op	     320 B/op	       5 allocs/op
BenchmarkNew_OrderedMap_Ldelete_fiveEntries-12             	 5922348	       197.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Ldelete_fiveEntries-12             	 6016850	       196.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithRange_oneEntry-12       	1000000000	         1.099 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithRange_oneEntry-12       	1000000000	         1.103 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithFront_oneEntry-12       	1000000000	         1.100 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithFront_oneEntry-12       	1000000000	         1.102 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithRange_fiveEntries-12    	483804207	         2.453 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithRange_fiveEntries-12    	489303610	         2.422 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithFront_fiveEntries-12    	474407106	         2.575 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithFront_fiveEntries-12    	487704862	         2.432 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/sttk/benchmarks_orderedmap/v0_5_0	43.495s
```
