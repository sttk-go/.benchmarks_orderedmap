# Benchmark of orderedmap v0.1.0

This measurement compared [github.com/sttk/orderedmap](https://github.com/sttk/orderedmap-go) with the following libraries:

- go standard map
- [github.com/elliotchance/orderedmap/v2](https://github.com/elliotchance/orderedmap)
- [github.com/wk8/go-ordered-map/v2](https://github.com/wk8/go-ordered-map)
- [github.com/iancoleman/orderedmap](https://github.com/iancoleman/orderedmap)
- [github.com/cevaris/ordered_map](https://github.com/cevaris/ordered_map)


```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_orderedmap/v0_1_0
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkOrderedMap_New-12                             	225976876	         5.255 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_New-12                                    	159506923	         7.449 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_New-12                                    	237500554	         5.055 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_New-12                                    	 8904388	       132.3 ns/op	     144 B/op	       4 allocs/op
BenchmarkOmI_New-12                                    	184292869	         6.632 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_New-12                                    	26627523	        38.50 ns/op	      32 B/op	       1 allocs/op
BenchmarkOrderedMap_Store_newOneEntry-12               	 8231464	       143.8 ns/op	     320 B/op	       3 allocs/op
BenchmarkMap_Store_newOneEntry-12                      	71484741	        18.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_newOneEntry-12                      	 8107554	       144.9 ns/op	     320 B/op	       3 allocs/op
BenchmarkOmW_Store_newOneEntry-12                      	 4289887	       280.7 ns/op	     432 B/op	       7 allocs/op
BenchmarkOmI_Store_newOneEntry-12                      	 6065775	       199.0 ns/op	     376 B/op	       4 allocs/op
BenchmarkOmC_Store_newOneEntry-12                      	 3377782	       355.4 ns/op	     680 B/op	       7 allocs/op
BenchmarkOrderedMap_Store_newFiveEntries-12            	 2687220	       439.4 ns/op	     576 B/op	       7 allocs/op
BenchmarkMap_Store_newFiveEntries-12                   	17155299	        67.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_newFiveEntries-12                   	 2705053	       445.6 ns/op	     576 B/op	       7 allocs/op
BenchmarkOmW_Store_newFiveEntries-12                   	 1536627	       750.8 ns/op	     752 B/op	      15 allocs/op
BenchmarkOmI_Store_newFiveEntries-12                   	 1774276	       683.9 ns/op	     696 B/op	      11 allocs/op
BenchmarkOmC_Store_newFiveEntries-12                   	 1257104	       935.4 ns/op	     904 B/op	      15 allocs/op
BenchmarkOrderedMap_Store_rewriteOneEntry-12           	183234828	         6.648 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Store_rewriteOneEntry-12                  	100000000	        10.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_rewriteOneEntry-12                  	127495576	         9.470 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Store_rewriteOneEntry-12                  	184913955	         6.470 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Store_rewriteOneEntry-12                  	23207211	        50.04 ns/op	      24 B/op	       1 allocs/op
BenchmarkOmC_Store_rewriteOneEntry-12                  	15092332	        81.72 ns/op	      24 B/op	       1 allocs/op
BenchmarkOrderedMap_Store_rewriteFiveEntries-12        	17165271	        67.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Store_rewriteFiveEntries-12               	18574891	        62.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_rewriteFiveEntries-12               	 9544726	       123.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Store_rewriteFiveEntries-12               	17717774	        67.81 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Store_rewriteFiveEntries-12               	 3901704	       301.1 ns/op	     120 B/op	       5 allocs/op
BenchmarkOmC_Store_rewriteFiveEntries-12               	 2805021	       422.5 ns/op	     120 B/op	       5 allocs/op
BenchmarkOrderedMap_Load_oneEntry-12                   	328080152	         3.612 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Load_oneEntry-12                          	351936919	         3.373 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Load_oneEntry-12                          	332938328	         3.603 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Load_oneEntry-12                          	331687356	         3.588 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Load_oneEntry-12                          	337818858	         3.548 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Load_oneEntry-12                          	71480734	        16.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_Load_fiveEntries-12                	20272993	        57.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Load_fiveEntries-12                       	20325537	        57.70 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Load_fiveEntries-12                       	20562991	        57.17 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Load_fiveEntries-12                       	20942077	        56.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Load_fiveEntries-12                       	21052701	        57.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Load_fiveEntries-12                       	13191853	        89.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_Delete_oneEntry-12                 	15337981	        77.81 ns/op	      64 B/op	       1 allocs/op
BenchmarkOrderedMap_Ldelete_oneEntry-12                	55733973	        21.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_oneEntry-12                        	43589364	        27.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Delete_oneEntry-12                        	14950190	        79.14 ns/op	      64 B/op	       1 allocs/op
BenchmarkOmW_Delete_oneEntry-12                        	10185577	       116.9 ns/op	      80 B/op	       2 allocs/op
BenchmarkOmI_Delete_oneEntry-12                        	14293220	        81.99 ns/op	      24 B/op	       1 allocs/op
BenchmarkOmC_Delete_oneEntry-12                        	 5019954	       235.1 ns/op	      56 B/op	       2 allocs/op
BenchmarkOrderedMap_Delete_fiveEntries-12              	 2381502	       495.4 ns/op	     320 B/op	       5 allocs/op
BenchmarkOrderedMap_Ldelete_fiveEntries-12             	 5832417	       199.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_fiveEntries-12                     	 7799774	       151.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Delete_fiveEntries-12                     	 2475244	       482.6 ns/op	     320 B/op	       5 allocs/op
BenchmarkOmW_Delete_fiveEntries-12                     	 1667479	       715.6 ns/op	     400 B/op	      10 allocs/op
BenchmarkOmI_Delete_fiveEntries-12                     	 2230762	       533.6 ns/op	     120 B/op	       5 allocs/op
BenchmarkOmC_Delete_fiveEntries-12                     	  966979	      1249 ns/op	     280 B/op	      10 allocs/op
BenchmarkOrderedMap_IterateWithRange_oneEntry-12       	1000000000	         1.130 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_IterateWithFront_oneEntry-12       	1000000000	         1.093 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Iterate_oneEntry-12                       	39162606	        30.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Iterate_oneEntry-12                       	1000000000	         0.9457 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Iterate_oneEntry-12                       	662453289	         1.746 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Iterate_oneEntry-12                       	258564163	         4.564 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Iterate_oneEntry-12                       	22714168	        51.93 ns/op	      32 B/op	       1 allocs/op
BenchmarkOrderedMap_IterateWithRange_fiveEntries-12    	445619594	         2.655 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_IterateWithFront_fiveEntries-12    	435975446	         2.645 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Iterate_fiveEntries-12                    	20967750	        55.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Iterate_fiveEntries-12                    	400428043	         2.953 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Iterate_fiveEntries-12                    	179061030	         6.625 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Iterate_fiveEntries-12                    	18890778	        59.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Iterate_fiveEntries-12                    	 4745277	       251.8 ns/op	     160 B/op	       5 allocs/op
PASS
ok  	github.com/sttk/benchmarks_orderedmap/v0_1_0	102.293s
```

