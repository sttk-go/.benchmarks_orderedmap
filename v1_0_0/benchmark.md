# Benchmark of orderedmap v1.0.0

This measurement compared [github.com/sttk/orderedmap](https://github.com/sttk/orderedmap) v1.0.0 with v0.6.0.

> BenchmarkNew_* ... v1.0.0, BenchmarkOld_* ... v0.6.0

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_orderedmap/v1_0_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkNew_OrderedMap_New-12                                   	233603660	         4.991 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_New-12                                   	240102171	         4.977 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Store_newOneEntry-12                     	 8433878	       140.4 ns/op	     320 B/op	       3 allocs/op
BenchmarkOld_OrderedMap_Store_newOneEntry-12                     	 8497327	       142.4 ns/op	     320 B/op	       3 allocs/op
BenchmarkNew_OrderedMap_Store_newFiveEntries-12                  	 2605658	       475.2 ns/op	     576 B/op	       7 allocs/op
BenchmarkOld_OrderedMap_Store_newFiveEntries-12                  	 2727255	       445.0 ns/op	     576 B/op	       7 allocs/op
BenchmarkNew_OrderedMap_Store_rewriteOneEntry-12                 	195176887	         6.124 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Store_rewriteOneEntry-12                 	197769052	         6.593 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Store_rewriteFiveEntries-12              	16030110	        73.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Store_rewriteFiveEntries-12              	16294638	        79.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Load_oneEntry-12                         	261393780	         4.261 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Load_oneEntry-12                         	307510924	         3.918 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Load_fiveEntries-12                      	19714333	        60.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Load_fiveEntries-12                      	19074674	        63.71 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Delete_oneEntry-12                       	13720958	        81.92 ns/op	      64 B/op	       1 allocs/op
BenchmarkOld_OrderedMap_Delete_oneEntry-12                       	14161364	        78.29 ns/op	      64 B/op	       1 allocs/op
BenchmarkNew_OrderedMap_Ldelete_oneEntry-12                      	56084493	        21.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Ldelete_oneEntry-12                      	55636460	        21.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_Delete_fiveEntries-12                    	 1947877	       542.3 ns/op	     320 B/op	       5 allocs/op
BenchmarkOld_OrderedMap_Delete_fiveEntries-12                    	 2365594	       502.5 ns/op	     320 B/op	       5 allocs/op
BenchmarkNew_OrderedMap_Ldelete_fiveEntries-12                   	 5914256	       201.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_Ldelete_fiveEntries-12                   	 5755848	       201.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithRange_oneEntry-12             	1000000000	         1.097 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithRange_oneEntry-12             	1000000000	         1.105 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithFront_oneEntry-12             	1000000000	         1.105 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithFront_oneEntry-12             	1000000000	         1.102 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithRange_fiveEntries-12          	491038879	         2.419 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithRange_fiveEntries-12          	484092074	         2.438 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_IterateWithFront_fiveEntries-12          	482437434	         2.440 ns/op	       0 B/op	       0 allocs/op
BenchmarkOld_OrderedMap_IterateWithFront_fiveEntries-12          	483978423	         2.452 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_empty-12                     	30511695	        38.86 ns/op	      64 B/op	       1 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_empty-12                     	30069876	        38.91 ns/op	      64 B/op	       1 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsString-12             	 1000000	      1123 ns/op	     264 B/op	      16 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsString-12             	 1000000	      1119 ns/op	     264 B/op	      16 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsStringPointer-12      	 1000000	      1096 ns/op	     312 B/op	      12 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsStringPointer-12      	 1000000	      1105 ns/op	     312 B/op	      12 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsInt-12                	 1334815	       881.0 ns/op	     184 B/op	      11 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsInt-12                	 1358151	       929.6 ns/op	     184 B/op	      11 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsIntPointer-12         	 1210636	       976.6 ns/op	     184 B/op	      11 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsIntPointer-12         	 1000000	      1027 ns/op	     184 B/op	      11 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsStruct-12             	  632274	      1947 ns/op	     928 B/op	      18 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsStruct-12             	  619155	      1956 ns/op	     928 B/op	      18 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsStructPointer-12      	  652317	      1822 ns/op	     768 B/op	      13 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsStructPointer-12      	  669128	      1735 ns/op	     768 B/op	      13 allocs/op
BenchmarkNew_OrderedMap_MarshalJSON_valueIsAny-12                	  696582	      1664 ns/op	     768 B/op	      13 allocs/op
BenchmarkOld_OrderedMap_MarshalJSON_valueIsAny-12                	  711710	      1701 ns/op	     768 B/op	      13 allocs/op
BenchmarkNew_OrderedMap_UnmarshalJSON_empty-12                   	 3837090	       307.9 ns/op	     888 B/op	       7 allocs/op
BenchmarkOld_OrderedMap_UnmarshalJSON_empty-12                   	 3928270	       307.7 ns/op	     888 B/op	       7 allocs/op
BenchmarkNew_OrderedMap_UnmarshalJSON_valueIsString-12           	  268245	      4416 ns/op	    2232 B/op	      86 allocs/op
BenchmarkOld_OrderedMap_UnmarshalJSON_valueIsString-12           	  264004	      4656 ns/op	    2232 B/op	      86 allocs/op
BenchmarkNew_OrderedMap_UnmarshalJSON_valueIsStringPointer-12    	  221761	      7059 ns/op	    2256 B/op	      90 allocs/op
BenchmarkOld_OrderedMap_UnmarshalJSON_valueIsStringPointer-12    	  239846	      5382 ns/op	    2256 B/op	      90 allocs/op
BenchmarkNew_OrderedMap_UnmarshalJSON_valueIsInt-12              	  215737	      5607 ns/op	    2128 B/op	      86 allocs/op
BenchmarkOld_OrderedMap_UnmarshalJSON_valueIsInt-12              	  165045	      6365 ns/op	    2128 B/op	      86 allocs/op
BenchmarkNew_OrderedMap_UnmarshalJSON_valueIsIntPointer-12       	  258334	      4641 ns/op	    2176 B/op	      91 allocs/op
BenchmarkOld_OrderedMap_UnmarshalJSON_valueIsIntPointer-12       	  259418	      4706 ns/op	    2176 B/op	      91 allocs/op
BenchmarkNew_OrderedMap_UnmarshalJSON_valueIsStruct-12           	  155064	      7652 ns/op	    2200 B/op	      74 allocs/op
BenchmarkOld_OrderedMap_UnmarshalJSON_valueIsStruct-12           	  153460	      7715 ns/op	    2200 B/op	      74 allocs/op
BenchmarkNew_OrderedMap_UnmarshalJSON_valueIsStructPointer-12    	  148270	      8149 ns/op	    2240 B/op	      79 allocs/op
BenchmarkOld_OrderedMap_UnmarshalJSON_valueIsStructPointer-12    	  148978	      8053 ns/op	    2240 B/op	      79 allocs/op
BenchmarkNew_OrderedMap_UnmarshalJSON_valueIsAny-12              	  147123	      8020 ns/op	    5576 B/op	     121 allocs/op
BenchmarkOld_OrderedMap_UnmarshalJSON_valueIsAny-12              	  145029	      8070 ns/op	    5576 B/op	     121 allocs/op
PASS
ok  	github.com/sttk/benchmarks_orderedmap/v1_0_0	90.479s
```
