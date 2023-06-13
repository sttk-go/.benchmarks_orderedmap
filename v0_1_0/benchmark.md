# Benchmark of orderedmap v0.1.0

This measurement compared [github.com/sttk-go/orderedmap](https://github.com/sttk-go/orderedmap) with the following libraries:

- go standard map
- [github.com/elliotchance/orderedmap/v2](https://github.com/elliotchance/orderedmap)
- [github.com/wk8/go-ordered-map/v2](https://github.com/wk8/go-ordered-map)
- [github.com/iancoleman/orderedmap](https://github.com/iancoleman/orderedmap)
- [github.com/cevaris/ordered_map](https://github.com/cevaris/ordered_map)


```
goos: darwin
goarch: amd64
pkg: github.com/sttk-go/.benchmarks_orderedmap/v0_1_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkOrderedMap_New-12                             	237307930	         4.980 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_New-12                                    	160368148	         7.352 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_New-12                                    	235260141	         4.995 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_New-12                                    	 9298748	       127.3 ns/op	     144 B/op	       4 allocs/op
BenchmarkOmI_New-12                                    	191895081	         6.279 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_New-12                                    	32636377	        36.81 ns/op	      32 B/op	       1 allocs/op
BenchmarkOrderedMap_Store_newOneEntry-12               	 8568556	       137.1 ns/op	     320 B/op	       3 allocs/op
BenchmarkMap_Store_newOneEntry-12                      	72915778	        19.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_newOneEntry-12                      	 7895223	       142.8 ns/op	     320 B/op	       3 allocs/op
BenchmarkOmW_Store_newOneEntry-12                      	 4241401	       276.6 ns/op	     432 B/op	       7 allocs/op
BenchmarkOmI_Store_newOneEntry-12                      	 6017940	       196.8 ns/op	     376 B/op	       4 allocs/op
BenchmarkOmC_Store_newOneEntry-12                      	 3355044	       357.7 ns/op	     680 B/op	       7 allocs/op
BenchmarkOrderedMap_Store_newFiveEntries-12            	 2930772	       408.4 ns/op	     576 B/op	       7 allocs/op
BenchmarkMap_Store_newFiveEntries-12                   	17605202	        66.59 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_newFiveEntries-12                   	 2768826	       421.1 ns/op	     576 B/op	       7 allocs/op
BenchmarkOmW_Store_newFiveEntries-12                   	 1618383	       737.4 ns/op	     752 B/op	      15 allocs/op
BenchmarkOmI_Store_newFiveEntries-12                   	 1773447	       683.9 ns/op	     696 B/op	      11 allocs/op
BenchmarkOmC_Store_newFiveEntries-12                   	 1319985	       918.6 ns/op	     904 B/op	      15 allocs/op
BenchmarkOrderedMap_Store_rewriteOneEntry-12           	280519699	         4.162 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Store_rewriteOneEntry-12                  	100000000	        10.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_rewriteOneEntry-12                  	129373201	         9.201 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Store_rewriteOneEntry-12                  	183779058	         6.369 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Store_rewriteOneEntry-12                  	24460201	        49.36 ns/op	      24 B/op	       1 allocs/op
BenchmarkOmC_Store_rewriteOneEntry-12                  	14829750	        79.39 ns/op	      24 B/op	       1 allocs/op
BenchmarkOrderedMap_Store_rewriteFiveEntries-12        	19751548	        59.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Store_rewriteFiveEntries-12               	20631087	        56.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_rewriteFiveEntries-12               	 9535741	       124.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Store_rewriteFiveEntries-12               	17370747	        66.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Store_rewriteFiveEntries-12               	 3935174	       307.4 ns/op	     120 B/op	       5 allocs/op
BenchmarkOmC_Store_rewriteFiveEntries-12               	 2815861	       424.3 ns/op	     120 B/op	       5 allocs/op
BenchmarkOrderedMap_Load_oneEntry-12                   	319221229	         3.691 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Load_oneEntry-12                          	346303748	         3.401 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Load_oneEntry-12                          	319800240	         3.684 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Load_oneEntry-12                          	316819644	         3.731 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Load_oneEntry-12                          	339030615	         3.545 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Load_oneEntry-12                          	72521364	        16.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_Load_fiveEntries-12                	19357840	        60.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Load_fiveEntries-12                       	20183835	        57.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Load_fiveEntries-12                       	20135984	        58.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Load_fiveEntries-12                       	20108024	        58.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Load_fiveEntries-12                       	20246328	        57.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Load_fiveEntries-12                       	13440994	        88.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_IterateWithRange_oneEntry-12       	1000000000	         0.9395 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_IterateWithFront_oneEntry-12       	1000000000	         0.7652 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Iterate_oneEntry-12                       	38783856	        30.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Iterate_oneEntry-12                       	1000000000	         0.7211 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Iterate_oneEntry-12                       	645186036	         1.836 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Iterate_oneEntry-12                       	259601649	         4.375 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Iterate_oneEntry-12                       	24034762	        49.22 ns/op	      32 B/op	       1 allocs/op
BenchmarkOrderedMap_IterateWithRange_fiveEntries-12    	357135253	         3.316 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_IterateWithFront_fiveEntries-12    	535868139	         2.147 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Iterate_fiveEntries-12                    	21201147	        55.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Iterate_fiveEntries-12                    	527567136	         2.163 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Iterate_fiveEntries-12                    	181595574	         6.579 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Iterate_fiveEntries-12                    	19704794	        59.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Iterate_fiveEntries-12                    	 4578742	       253.2 ns/op	     160 B/op	       5 allocs/op
PASS
ok  	github.com/sttk-go/.benchmarks_orderedmap/v0_1_0	80.585s
```
