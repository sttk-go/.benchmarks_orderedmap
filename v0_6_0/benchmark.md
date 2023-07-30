# Benchmark of orderedmap v0.6.0

This measurement compared [github.com/sttk/orderedmap](https://github.com/sttk/orderedmap) v0.6.0 with v0.5.0

> BenchmarkNew_* ... v0.6.0, BenchmarkOld_* ... v0.5.0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_orderedmap/v0_6_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkNew_OrderedMap_New-12                                 	237650155	         5.163 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_New-12                                 	230594876	         5.259 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Store_newOneEntry-12                   	 8389599	       141.6 ns/op	     320 B/op	       3 allocs/op
BenchmarkOld_OrderedMap_Store_newOneEntry-12                   	 8477829	       139.3 ns/op	     320 B/op	       3 allocs/op
BenchmarkNew_OrderedMap_Store_newFiveEntries-12                	 2697637	       445.6 ns/op	     576 B/op	       7 allocs/op
BenchmarkOld_OrderedMap_Store_newFiveEntries-12                	 2625998	       449.4 ns/op	     576 B/op	       7 allocs/op
BenchmarkNew_OrderedMap_Store_rewriteOneEntry-12               	196288118	         6.076 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Store_rewriteOneEntry-12               	195903399	         6.036 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Store_rewriteFiveEntries-12            	15446953	        70.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Store_rewriteFiveEntries-12            	17302159	        69.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Load_oneEntry-12                       	302443078	         3.967 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Load_oneEntry-12                       	307310734	         3.886 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Load_fiveEntries-12                    	20197333	        60.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Load_fiveEntries-12                    	19242021	        60.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Delete_oneEntry-12                     	15535344	        75.77 ns/op	      64 B/op	       1 allocs/op
BenchmarkOld_OrderedMap_Delete_oneEntry-12                     	15177268	        76.21 ns/op	      64 B/op	       1 allocs/op
BenchmarkNew_OrderedMap_Ldelete_oneEntry-12                    	57810558	        20.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Ldelete_oneEntry-12                    	58294222	        20.79 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Delete_fiveEntries-12                  	 2413815	       505.9 ns/op	     320 B/op	       5 allocs/op
BenchmarkOld_OrderedMap_Delete_fiveEntries-12                  	 2406757	       520.9 ns/op	     320 B/op	       5 allocs/op
BenchmarkNew_OrderedMap_Ldelete_fiveEntries-12                 	 6165386	       196.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Ldelete_fiveEntries-12                 	 5937308	       198.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithRange_oneEntry-12           	1000000000	         1.120 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithRange_oneEntry-12           	1000000000	         1.084 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithFront_oneEntry-12           	1000000000	         1.083 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithFront_oneEntry-12           	1000000000	         1.085 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithRange_fiveEntries-12        	495839900	         2.425 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithRange_fiveEntries-12        	486771297	         2.423 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithFront_fiveEntries-12        	458716870	         2.486 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithFront_fiveEntries-12        	463191949	         2.597 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_empty-12                   	30772724	        38.58 ns/op	      64 B/op	       1 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_empty-12                   	31277992	        38.03 ns/op	      64 B/op	       1 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsString-12           	 1000000	      1087 ns/op	     264 B/op	      16 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsString-12           	  767989	      1553 ns/op	     304 B/op	      21 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsStringPointer-12    	 1000000	      1053 ns/op	     312 B/op	      12 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsStringPointer-12    	  793401	      1483 ns/op	     352 B/op	      17 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsInt-12              	 1387006	       865.8 ns/op	     184 B/op	      11 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsInt-12              	  916072	      1321 ns/op	     224 B/op	      16 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsIntPointer-12       	 1300142	       912.9 ns/op	     184 B/op	      11 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsIntPointer-12       	  856185	      1353 ns/op	     224 B/op	      16 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsStruct-12           	  650642	      1825 ns/op	     928 B/op	      18 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsStruct-12           	  505735	      2365 ns/op	     968 B/op	      23 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsStructPointer-12    	  692540	      1725 ns/op	     768 B/op	      13 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsStructPointer-12    	  531334	      2254 ns/op	     808 B/op	      18 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsAny-12              	  720528	      1633 ns/op	     768 B/op	      13 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsAny-12              	  538660	      2204 ns/op	     808 B/op	      18 allocs/op
PASS
ok  	github.com/sttk/benchmarks_orderedmap/v0_6_0	64.266s
```
