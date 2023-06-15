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
BenchmarkOrderedMap_New-12                             	233836898	         4.991 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_New-12                                    	162204830	         7.351 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_New-12                                    	239017436	         4.985 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_New-12                                    	 8625085	       132.7 ns/op	     144 B/op	       4 allocs/op
BenchmarkOmI_New-12                                    	193458746	         6.178 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_New-12                                    	32036374	        37.48 ns/op	      32 B/op	       1 allocs/op
BenchmarkOrderedMap_Store_newOneEntry-12               	 8192431	       141.3 ns/op	     320 B/op	       3 allocs/op
BenchmarkMap_Store_newOneEntry-12                      	72890980	        15.92 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_newOneEntry-12                      	 8175706	       144.2 ns/op	     320 B/op	       3 allocs/op
BenchmarkOmW_Store_newOneEntry-12                      	 4164040	       286.0 ns/op	     432 B/op	       7 allocs/op
BenchmarkOmI_Store_newOneEntry-12                      	 6095299	       196.4 ns/op	     376 B/op	       4 allocs/op
BenchmarkOmC_Store_newOneEntry-12                      	 3323998	       355.4 ns/op	     680 B/op	       7 allocs/op
BenchmarkOrderedMap_Store_newFiveEntries-12            	 2673966	       445.6 ns/op	     576 B/op	       7 allocs/op
BenchmarkMap_Store_newFiveEntries-12                   	16447040	        68.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_newFiveEntries-12                   	 2743116	       440.6 ns/op	     576 B/op	       7 allocs/op
BenchmarkOmW_Store_newFiveEntries-12                   	 1589464	       757.9 ns/op	     752 B/op	      15 allocs/op
BenchmarkOmI_Store_newFiveEntries-12                   	 1788235	       669.2 ns/op	     696 B/op	      11 allocs/op
BenchmarkOmC_Store_newFiveEntries-12                   	 1319750	       918.2 ns/op	     904 B/op	      15 allocs/op
BenchmarkOrderedMap_Store_rewriteOneEntry-12           	203976302	         5.846 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Store_rewriteOneEntry-12                  	100000000	        10.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_rewriteOneEntry-12                  	129922635	         9.174 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Store_rewriteOneEntry-12                  	190130857	         6.226 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Store_rewriteOneEntry-12                  	24440096	        49.40 ns/op	      24 B/op	       1 allocs/op
BenchmarkOmC_Store_rewriteOneEntry-12                  	14472936	        82.41 ns/op	      24 B/op	       1 allocs/op
BenchmarkOrderedMap_Store_rewriteFiveEntries-12        	17677873	        68.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Store_rewriteFiveEntries-12               	20418392	        57.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Store_rewriteFiveEntries-12               	 9598668	       122.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Store_rewriteFiveEntries-12               	17465239	        68.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Store_rewriteFiveEntries-12               	 3851229	       303.0 ns/op	     120 B/op	       5 allocs/op
BenchmarkOmC_Store_rewriteFiveEntries-12               	 2846256	       418.3 ns/op	     120 B/op	       5 allocs/op
BenchmarkOrderedMap_Load_oneEntry-12                   	330149112	         3.634 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Load_oneEntry-12                          	356583076	         3.372 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Load_oneEntry-12                          	322096626	         3.700 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Load_oneEntry-12                          	318117546	         3.726 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Load_oneEntry-12                          	339404163	         3.519 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Load_oneEntry-12                          	70996598	        16.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_Load_fiveEntries-12                	20171457	        58.36 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Load_fiveEntries-12                       	20829979	        58.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Load_fiveEntries-12                       	20327221	        58.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Load_fiveEntries-12                       	20387403	        58.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Load_fiveEntries-12                       	19466670	        58.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Load_fiveEntries-12                       	13398060	        87.34 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_Delete_oneEntry-12                 	15204871	        77.77 ns/op	      64 B/op	       1 allocs/op
BenchmarkOrderedMap_Ldelete_oneEntry-12                	57471456	        20.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_oneEntry-12                        	44500672	        27.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Delete_oneEntry-12                        	15385884	        77.62 ns/op	      64 B/op	       1 allocs/op
BenchmarkOmW_Delete_oneEntry-12                        	 9903642	       116.0 ns/op	      80 B/op	       2 allocs/op
BenchmarkOmI_Delete_oneEntry-12                        	14656464	        80.63 ns/op	      24 B/op	       1 allocs/op
BenchmarkOmC_Delete_oneEntry-12                        	 4911799	       235.8 ns/op	      56 B/op	       2 allocs/op
BenchmarkOrderedMap_Delete_fiveEntries-12              	 2425153	       502.9 ns/op	     320 B/op	       5 allocs/op
BenchmarkOrderedMap_Ldelete_fiveEntries-12             	 6080157	       196.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Delete_fiveEntries-12                     	 7854655	       149.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Delete_fiveEntries-12                     	 2423385	       494.8 ns/op	     320 B/op	       5 allocs/op
BenchmarkOmW_Delete_fiveEntries-12                     	 1692988	       706.3 ns/op	     400 B/op	      10 allocs/op
BenchmarkOmI_Delete_fiveEntries-12                     	 2199733	       527.4 ns/op	     120 B/op	       5 allocs/op
BenchmarkOmC_Delete_fiveEntries-12                     	  966180	      1239 ns/op	     280 B/op	      10 allocs/op
BenchmarkOrderedMap_IterateWithRange_oneEntry-12       	1000000000	         1.100 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_IterateWithFront_oneEntry-12       	1000000000	         1.085 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Iterate_oneEntry-12                       	40343478	        29.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Iterate_oneEntry-12                       	1000000000	         0.9323 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Iterate_oneEntry-12                       	651133122	         1.818 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Iterate_oneEntry-12                       	277199575	         4.282 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Iterate_oneEntry-12                       	19917408	        51.10 ns/op	      32 B/op	       1 allocs/op
BenchmarkOrderedMap_IterateWithRange_fiveEntries-12    	452908918	         2.621 ns/op	       0 B/op	       0 allocs/op
BenchmarkOrderedMap_IterateWithFront_fiveEntries-12    	450832488	         2.623 ns/op	       0 B/op	       0 allocs/op
BenchmarkMap_Iterate_fiveEntries-12                    	21047716	        57.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmE_Iterate_fiveEntries-12                    	421055880	         2.781 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmW_Iterate_fiveEntries-12                    	168688339	         7.151 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_Iterate_fiveEntries-12                    	20229650	        59.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmC_Iterate_fiveEntries-12                    	 4690234	       258.1 ns/op	     160 B/op	       5 allocs/op
PASS
ok  	github.com/sttk-go/.benchmarks_orderedmap/v0_1_0	102.449s
```

