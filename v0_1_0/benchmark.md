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
BenchmarkOrderedMap_New-12                             	237535706	         4.993 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_New-12                                    	161318598	         7.359 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_New-12                                    	239920862	         4.975 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_New-12                                    	 9319274	       127.0 ns/op	     144 B/op	       4 allocs/op
BenchmarkOmI_New-12                                    	193338481	         6.190 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_New-12                                    	32856577	        36.00 ns/op	      32 B/op	       1 allocs/op
BenchmarkOrderedMap_Store_newOneEntry-12               	 8273097	       146.6 ns/op	     320 B/op	       3 allocs/op
BenchmarkMap_Store_newOneEntry-12                      	69303314	        16.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_newOneEntry-12                      	 8348668	       147.3 ns/op	     320 B/op	       3 allocs/op
BenchmarkOmW_Store_newOneEntry-12                      	 4332229	       278.3 ns/op	     432 B/op	       7 allocs/op
BenchmarkOmI_Store_newOneEntry-12                      	 6062400	       201.9 ns/op	     376 B/op	       4 allocs/op
BenchmarkOmC_Store_newOneEntry-12                      	 3434802	       346.9 ns/op	     680 B/op	       7 allocs/op
BenchmarkOrderedMap_Store_newFiveEntries-12            	 2751435	       434.6 ns/op	     576 B/op	       7 allocs/op
BenchmarkMap_Store_newFiveEntries-12                   	17141296	        67.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_newFiveEntries-12                   	 2744383	       438.8 ns/op	     576 B/op	       7 allocs/op
BenchmarkOmW_Store_newFiveEntries-12                   	 1617465	       745.3 ns/op	     752 B/op	      15 allocs/op
BenchmarkOmI_Store_newFiveEntries-12                   	 1776920	       669.1 ns/op	     696 B/op	      11 allocs/op
BenchmarkOmC_Store_newFiveEntries-12                   	 1323016	       917.7 ns/op	     904 B/op	      15 allocs/op
BenchmarkOrderedMap_Store_rewriteOneEntry-12           	204958378	         5.796 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Store_rewriteOneEntry-12                  	100000000	        10.15 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_rewriteOneEntry-12                  	130792699	         9.111 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Store_rewriteOneEntry-12                  	193313910	         6.206 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Store_rewriteOneEntry-12                  	24969692	        48.44 ns/op	      24 B/op	       1 allocs/op
BenchmarkOmC_Store_rewriteOneEntry-12                  	15243165	        79.32 ns/op	      24 B/op	       1 allocs/op
BenchmarkOrderedMap_Store_rewriteFiveEntries-12        	17918638	        65.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Store_rewriteFiveEntries-12               	19397265	        60.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_rewriteFiveEntries-12               	 9746870	       121.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Store_rewriteFiveEntries-12               	17989981	        66.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Store_rewriteFiveEntries-12               	 3787708	       305.7 ns/op	     120 B/op	       5 allocs/op
BenchmarkOmC_Store_rewriteFiveEntries-12               	 2856198	       419.2 ns/op	     120 B/op	       5 allocs/op
BenchmarkOrderedMap_Load_oneEntry-12                   	328271739	         3.586 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Load_oneEntry-12                          	355124528	         3.336 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Load_oneEntry-12                          	323068743	         3.679 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Load_oneEntry-12                          	304049613	         3.748 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Load_oneEntry-12                          	340718688	         3.514 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Load_oneEntry-12                          	70897368	        16.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_Load_fiveEntries-12                	21055326	        56.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Load_fiveEntries-12                       	20005021	        57.55 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Load_fiveEntries-12                       	20617323	        56.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Load_fiveEntries-12                       	21155775	        56.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Load_fiveEntries-12                       	20993260	        56.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Load_fiveEntries-12                       	13356096	        87.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_Delete_oneEntry-12                 	14702809	        78.41 ns/op	      64 B/op	       1 allocs/op
BenchmarkOrderedMap_Ldelete_oneEntry-12                	56408202	        20.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_oneEntry-12                        	43988404	        27.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Delete_oneEntry-12                        	15152523	        83.02 ns/op	      64 B/op	       1 allocs/op
BenchmarkOmW_Delete_oneEntry-12                        	 9730917	       116.1 ns/op	      80 B/op	       2 allocs/op
BenchmarkOmI_Delete_oneEntry-12                        	14041137	        80.77 ns/op	      24 B/op	       1 allocs/op
BenchmarkOmC_Delete_oneEntry-12                        	 4950622	       234.9 ns/op	      56 B/op	       2 allocs/op
BenchmarkOrderedMap_Delete_fiveEntries-12              	 2429157	       507.2 ns/op	     320 B/op	       5 allocs/op
BenchmarkOrderedMap_Ldelete_fiveEntries-12             	 5708865	       199.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_fiveEntries-12                     	 7415769	       155.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Delete_fiveEntries-12                     	 2475836	       479.4 ns/op	     320 B/op	       5 allocs/op
BenchmarkOmW_Delete_fiveEntries-12                     	 1713296	       691.0 ns/op	     400 B/op	      10 allocs/op
BenchmarkOmI_Delete_fiveEntries-12                     	 2225256	       529.4 ns/op	     120 B/op	       5 allocs/op
BenchmarkOmC_Delete_fiveEntries-12                     	  986916	      1208 ns/op	     280 B/op	      10 allocs/op
BenchmarkOrderedMap_IterateWithRange_oneEntry-12       	1000000000	         0.7207 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_IterateWithFront_oneEntry-12       	1000000000	         1.098 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Iterate_oneEntry-12                       	37918939	        31.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Iterate_oneEntry-12                       	1000000000	         0.7279 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Iterate_oneEntry-12                       	661500330	         1.822 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Iterate_oneEntry-12                       	276982584	         4.296 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Iterate_oneEntry-12                       	23685283	        50.67 ns/op	      32 B/op	       1 allocs/op
BenchmarkOrderedMap_IterateWithRange_fiveEntries-12    	482856134	         2.451 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_IterateWithFront_fiveEntries-12    	491418451	         2.416 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Iterate_fiveEntries-12                    	21247608	        55.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Iterate_fiveEntries-12                    	536706626	         2.177 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Iterate_fiveEntries-12                    	180753157	         6.560 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Iterate_fiveEntries-12                    	20082067	        58.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Iterate_fiveEntries-12                    	 4739180	       247.3 ns/op	     160 B/op	       5 allocs/op
PASS
ok  	github.com/sttk/benchmarks_orderedmap/v0_1_0	100.754s
```

