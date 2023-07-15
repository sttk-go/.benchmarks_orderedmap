# Benchmark of orderedmap JSON marshaling/unmarshaling v0.6.0

This measurement compared [github.com/sttk/orderedmap](https://github.com/sttk/orderedmap-go) with the following libraries:

- go standard map
- [github.com/wk8/go-ordered-map/v2](https://github.com/wk8/go-ordered-map)
- [github.com/iancoleman/orderedmap](https://github.com/iancoleman/orderedmap)

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_orderedmap/v0_6_0/json
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkOrderedMap_MarshalJSON_empty-12                     	30709935	        40.05 ns/op	      64 B/op	       1 allocs/op
BenchmarkMap_MarshalJSON_empty-12                            	 6902318	       172.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkOmW_MarshalJSON_empty-12                            	 5346298	       226.8 ns/op	     288 B/op	       5 allocs/op
BenchmarkOmI_MarshalJSON_empty-12                            	16630740	        69.51 ns/op	     112 B/op	       2 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsString-12             	 1000000	      1157 ns/op	     264 B/op	      16 allocs/op
BenchmarkMap_MarshalJSON_valueIsString-12                    	  923750	      1297 ns/op	     696 B/op	      15 allocs/op
BenchmarkOmW_MarshalJSON_valueIsString-12                    	  915439	      1309 ns/op	     488 B/op	      20 allocs/op
BenchmarkOmI_MarshalJSON_valueIsString-12                    	  879198	      1334 ns/op	     320 B/op	       8 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsStringPointer-12      	 1000000	      1121 ns/op	     312 B/op	      12 allocs/op
BenchmarkMap_MarshalJSON_valueIsStringPointer-12             	  986151	      1191 ns/op	     616 B/op	      10 allocs/op
BenchmarkOmW_MarshalJSON_valueIsStringPointer-12             	  937146	      1247 ns/op	     424 B/op	      15 allocs/op
BenchmarkOmI_MarshalJSON_valueIsStringPointer-12             	  879739	      1387 ns/op	     320 B/op	       8 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsInt-12                	 1298704	       927.2 ns/op	     184 B/op	      11 allocs/op
BenchmarkMap_MarshalJSON_valueIsInt-12                       	  967168	      1205 ns/op	     624 B/op	      15 allocs/op
BenchmarkOmW_MarshalJSON_valueIsInt-12                       	 1000000	      1081 ns/op	     408 B/op	      15 allocs/op
BenchmarkOmI_MarshalJSON_valueIsInt-12                       	  981204	      1252 ns/op	     192 B/op	       7 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsIntPointer-12         	 1000000	      1003 ns/op	     184 B/op	      11 allocs/op
BenchmarkMap_MarshalJSON_valueIsIntPointer-12                	 1000000	      1154 ns/op	     584 B/op	      10 allocs/op
BenchmarkOmW_MarshalJSON_valueIsIntPointer-12                	 1000000	      1149 ns/op	     408 B/op	      15 allocs/op
BenchmarkOmI_MarshalJSON_valueIsIntPointer-12                	  916609	      1268 ns/op	     192 B/op	       7 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsStruct-12             	  618319	      1974 ns/op	     928 B/op	      18 allocs/op
BenchmarkMap_MarshalJSON_valueIsStruct-12                    	  641107	      1950 ns/op	     952 B/op	      15 allocs/op
BenchmarkOmW_MarshalJSON_valueIsStruct-12                    	  574807	      2118 ns/op	    1384 B/op	      21 allocs/op
BenchmarkOmI_MarshalJSON_valueIsStruct-12                    	  609172	      1957 ns/op	     576 B/op	       9 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsStructPointer-12      	  684972	      1750 ns/op	     768 B/op	      13 allocs/op
BenchmarkMap_MarshalJSON_valueIsStructPointer-12             	  664712	      1791 ns/op	     792 B/op	      10 allocs/op
BenchmarkOmW_MarshalJSON_valueIsStructPointer-12             	  619893	      1913 ns/op	    1224 B/op	      16 allocs/op
BenchmarkOmI_MarshalJSON_valueIsStructPointer-12             	  593328	      2014 ns/op	     576 B/op	       9 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsAny-12                	  705843	      1683 ns/op	     768 B/op	      13 allocs/op
BenchmarkMap_MarshalJSON_valueIsAny-12                       	  562125	      2093 ns/op	     872 B/op	      15 allocs/op
BenchmarkOmW_MarshalJSON_valueIsAny-12                       	  642286	      1878 ns/op	    1224 B/op	      16 allocs/op
BenchmarkOmI_MarshalJSON_valueIsAny-12                       	  609439	      1951 ns/op	     576 B/op	       9 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_empty-12                   	 4163556	       279.6 ns/op	     880 B/op	       5 allocs/op
BenchmarkMap_UnmarshalJSON_empty-12                          	 7267635	       160.2 ns/op	     152 B/op	       2 allocs/op
BenchmarkOmW_UnmarshalJSON_empty-12                          	167325691	         7.147 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_UnmarshalJSON_empty-12                          	 2802804	       423.5 ns/op	    1040 B/op	       6 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsString-12           	  264420	      4388 ns/op	    2232 B/op	      85 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsString-12                  	 2260580	       525.2 ns/op	     264 B/op	       7 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsString-12                  	  763002	      1562 ns/op	     800 B/op	      24 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsString-12                  	 2224484	       530.2 ns/op	     264 B/op	       7 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsStringPointer-12    	  256683	      4705 ns/op	    2256 B/op	      89 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsStringPointer-12           	 2294766	       521.7 ns/op	     264 B/op	       7 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsStringPointer-12           	  580022	      1814 ns/op	     832 B/op	      28 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsStringPointer-12           	 2231823	       529.4 ns/op	     264 B/op	       7 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsInt-12              	  285204	      4225 ns/op	    2128 B/op	      84 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsInt-12                     	  564348	      2085 ns/op	     392 B/op	      22 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsInt-12                     	  698647	      1718 ns/op	     960 B/op	      29 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsInt-12                     	  174816	      6913 ns/op	    2632 B/op	     111 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsIntPointer-12       	  237504	      4536 ns/op	    2168 B/op	      89 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsIntPointer-12              	  541972	      2141 ns/op	     440 B/op	      27 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsIntPointer-12              	  574118	      2075 ns/op	    1008 B/op	      34 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsIntPointer-12              	  173380	      6728 ns/op	    2632 B/op	     111 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsStruct-12           	  150669	      7638 ns/op	    2192 B/op	      72 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsStruct-12                  	  192512	      6148 ns/op	     576 B/op	      33 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsStruct-12                  	  173208	      6827 ns/op	    1712 B/op	      60 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsStruct-12                  	   44370	     25747 ns/op	   11048 B/op	     427 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsStructPointer-12    	  150312	      7890 ns/op	    2232 B/op	      77 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsStructPointer-12           	  194050	      6077 ns/op	     712 B/op	      38 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsStructPointer-12           	  166336	      7173 ns/op	    1752 B/op	      65 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsStructPointer-12           	   45966	     25790 ns/op	   11048 B/op	     427 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsAny-12              	  142850	      8053 ns/op	    5576 B/op	     119 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsAny-12                     	  178378	      6296 ns/op	    4024 B/op	      80 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsAny-12                     	  177093	      6741 ns/op	    4720 B/op	      95 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsAny-12                     	   43582	     25755 ns/op	   11048 B/op	     427 allocs/op
PASS
ok  	github.com/sttk/benchmarks_orderedmap/v0_6_0/json	86.064s
```
