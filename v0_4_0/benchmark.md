# Benchmark of orderedmap v0.4.0

This measurement compared [github.com/sttk/orderedmap](https://github.com/sttk/orderedmap-go) v0.4.0 with v0.1.0

> BenchmarkNew_* ... v0.4.0, BenchmarkOld_* ... v0.1.0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_orderedmap/v0_4_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkNew_OrderedMap_New-12                             	234283291	         4.972 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_New-12                             	237506594	         5.021 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Store_newOneEntry-12               	 8355214	       142.3 ns/op	     320 B/op	       3 allocs/op
BenchmarkOld_OrderedMap_Store_newOneEntry-12               	 8366300	       144.7 ns/op	     320 B/op	       3 allocs/op
BenchmarkNew_OrderedMap_Store_newFiveEntries-12            	 2740633	       443.4 ns/op	     576 B/op	       7 allocs/op
BenchmarkOld_OrderedMap_Store_newFiveEntries-12            	 2692292	       444.6 ns/op	     576 B/op	       7 allocs/op
BenchmarkNew_OrderedMap_Store_rewriteOneEntry-12           	205456267	         5.790 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Store_rewriteOneEntry-12           	197829985	         5.810 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Store_rewriteFiveEntries-12        	17945613	        65.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Store_rewriteFiveEntries-12        	18232474	        64.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Load_oneEntry-12                   	330242426	         3.596 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Load_oneEntry-12                   	330308085	         3.589 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Load_fiveEntries-12                	20910357	        57.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Load_fiveEntries-12                	20844068	        56.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Delete_oneEntry-12                 	15618668	        75.83 ns/op	      64 B/op	       1 allocs/op
BenchmarkOld_OrderedMap_Delete_oneEntry-12                 	15523722	        76.36 ns/op	      64 B/op	       1 allocs/op
BenchmarkNew_OrderedMap_Ldelete_oneEntry-12                	58400505	        20.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Ldelete_oneEntry-12                	58059978	        20.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Delete_fiveEntries-12              	 2374360	       504.5 ns/op	     320 B/op	       5 allocs/op
BenchmarkOld_OrderedMap_Delete_fiveEntries-12              	 2381367	       493.0 ns/op	     320 B/op	       5 allocs/op
BenchmarkNew_OrderedMap_Ldelete_fiveEntries-12             	 5948600	       197.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Ldelete_fiveEntries-12             	 5982142	       199.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithRange_oneEntry-12       	1000000000	         1.101 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithRange_oneEntry-12       	1000000000	         1.106 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithFront_oneEntry-12       	1000000000	         1.101 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithFront_oneEntry-12       	1000000000	         1.101 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithRange_fiveEntries-12    	478537676	         2.445 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithRange_fiveEntries-12    	480833822	         2.458 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithFront_fiveEntries-12    	482210365	         2.445 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithFront_fiveEntries-12    	465186624	         2.593 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/sttk/benchmarks_orderedmap/v0_4_0	43.100s
```
