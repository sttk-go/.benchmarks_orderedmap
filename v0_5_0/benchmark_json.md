# Benchmark of orderedmap JSON marshaling/unmarshaling v0.5.0

This measurement compared [github.com/sttk/orderedmap](https://github.com/sttk/orderedmap-go) with the following libraries:

- go standard map
- [github.com/wk8/go-ordered-map/v2](https://github.com/wk8/go-ordered-map)
- [github.com/iancoleman/orderedmap](https://github.com/iancoleman/orderedmap)

```
goos: darwin
goarch: amd64
pkg: github.com/sttk/benchmarks_orderedmap/v0_5_0/json
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkOrderedMap_MarshalJSON_empty-12                     	30356950	        38.30 ns/op	      64 B/op	       1 allocs/op
BenchmarkMap_MarshalJSON_empty-12                            	 6808545	       174.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkOmW_MarshalJSON_empty-12                            	 5465012	       223.5 ns/op	     288 B/op	       5 allocs/op
BenchmarkOmI_MarshalJSON_empty-12                            	17310634	        68.35 ns/op	     112 B/op	       2 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsString-12             	  717050	      1642 ns/op	     304 B/op	      21 allocs/op
BenchmarkMap_MarshalJSON_valueIsString-12                    	  955423	      1299 ns/op	     696 B/op	      15 allocs/op
BenchmarkOmW_MarshalJSON_valueIsString-12                    	  900835	      1336 ns/op	     488 B/op	      20 allocs/op
BenchmarkOmI_MarshalJSON_valueIsString-12                    	  902044	      1344 ns/op	     320 B/op	       8 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsStringPointer-12      	  749647	      1614 ns/op	     352 B/op	      17 allocs/op
BenchmarkMap_MarshalJSON_valueIsStringPointer-12             	  965278	      1213 ns/op	     616 B/op	      10 allocs/op
BenchmarkOmW_MarshalJSON_valueIsStringPointer-12             	  915316	      1281 ns/op	     424 B/op	      15 allocs/op
BenchmarkOmI_MarshalJSON_valueIsStringPointer-12             	  793724	      1427 ns/op	     320 B/op	       8 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsInt-12                	  810105	      1458 ns/op	     224 B/op	      16 allocs/op
BenchmarkMap_MarshalJSON_valueIsInt-12                       	  939774	      1218 ns/op	     624 B/op	      15 allocs/op
BenchmarkOmW_MarshalJSON_valueIsInt-12                       	  928370	      1119 ns/op	     408 B/op	      15 allocs/op
BenchmarkOmI_MarshalJSON_valueIsInt-12                       	  979503	      1237 ns/op	     192 B/op	       7 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsIntPointer-12         	  777591	      1488 ns/op	     224 B/op	      16 allocs/op
BenchmarkMap_MarshalJSON_valueIsIntPointer-12                	 1000000	      1158 ns/op	     584 B/op	      10 allocs/op
BenchmarkOmW_MarshalJSON_valueIsIntPointer-12                	  976424	      1170 ns/op	     408 B/op	      15 allocs/op
BenchmarkOmI_MarshalJSON_valueIsIntPointer-12                	  894009	      1341 ns/op	     192 B/op	       7 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsStruct-12             	  474433	      2511 ns/op	     968 B/op	      23 allocs/op
BenchmarkMap_MarshalJSON_valueIsStruct-12                    	  627913	      1921 ns/op	     952 B/op	      15 allocs/op
BenchmarkOmW_MarshalJSON_valueIsStruct-12                    	  523711	      2078 ns/op	    1384 B/op	      21 allocs/op
BenchmarkOmI_MarshalJSON_valueIsStruct-12                    	  570598	      1993 ns/op	     576 B/op	       9 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsStructPointer-12      	  513616	      2331 ns/op	     808 B/op	      18 allocs/op
BenchmarkMap_MarshalJSON_valueIsStructPointer-12             	  674902	      1779 ns/op	     792 B/op	      10 allocs/op
BenchmarkOmW_MarshalJSON_valueIsStructPointer-12             	  617948	      1963 ns/op	    1224 B/op	      16 allocs/op
BenchmarkOmI_MarshalJSON_valueIsStructPointer-12             	  563228	      1995 ns/op	     576 B/op	       9 allocs/op
BenchmarkOrderedMap_MarshalJSON_valueIsAny-12                	  541395	      2219 ns/op	     808 B/op	      18 allocs/op
BenchmarkMap_MarshalJSON_valueIsAny-12                       	  571113	      2051 ns/op	     872 B/op	      15 allocs/op
BenchmarkOmW_MarshalJSON_valueIsAny-12                       	  636100	      1870 ns/op	    1224 B/op	      16 allocs/op
BenchmarkOmI_MarshalJSON_valueIsAny-12                       	  588669	      2066 ns/op	     576 B/op	       9 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_empty-12                   	 4251066	       276.0 ns/op	     880 B/op	       5 allocs/op
BenchmarkMap_UnmarshalJSON_empty-12                          	 7109014	       161.8 ns/op	     152 B/op	       2 allocs/op
BenchmarkOmW_UnmarshalJSON_empty-12                          	167540503	         7.163 ns/op	       0 B/op	       0 allocs/op
BenchmarkOmI_UnmarshalJSON_empty-12                          	 2781500	       424.3 ns/op	    1040 B/op	       6 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsString-12           	  270926	      4466 ns/op	    2232 B/op	      85 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsString-12                  	 2233186	       520.0 ns/op	     264 B/op	       7 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsString-12                  	  736095	      1561 ns/op	     800 B/op	      24 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsString-12                  	 2281746	       524.7 ns/op	     264 B/op	       7 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsStringPointer-12    	  252920	      4703 ns/op	    2256 B/op	      89 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsStringPointer-12           	 2314224	       532.0 ns/op	     264 B/op	       7 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsStringPointer-12           	  651772	      1819 ns/op	     832 B/op	      28 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsStringPointer-12           	 2273912	       532.7 ns/op	     264 B/op	       7 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsInt-12              	  286004	      4186 ns/op	    2128 B/op	      84 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsInt-12                     	  564739	      2078 ns/op	     392 B/op	      22 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsInt-12                     	  689529	      1717 ns/op	     960 B/op	      29 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsInt-12                     	  163986	      6798 ns/op	    2632 B/op	     111 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsIntPointer-12       	  265920	      4574 ns/op	    2168 B/op	      89 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsIntPointer-12              	  542498	      2162 ns/op	     440 B/op	      27 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsIntPointer-12              	  578454	      2106 ns/op	    1008 B/op	      34 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsIntPointer-12              	  176238	      6803 ns/op	    2632 B/op	     111 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsStruct-12           	  146263	      7638 ns/op	    2192 B/op	      72 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsStruct-12                  	  192151	      6146 ns/op	     576 B/op	      33 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsStruct-12                  	  172959	      6836 ns/op	    1712 B/op	      60 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsStruct-12                  	   46285	     25882 ns/op	   11048 B/op	     427 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsStructPointer-12    	  139746	      7928 ns/op	    2232 B/op	      77 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsStructPointer-12           	  193263	      6130 ns/op	     712 B/op	      38 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsStructPointer-12           	  167046	      7285 ns/op	    1752 B/op	      65 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsStructPointer-12           	   47103	     25825 ns/op	   11048 B/op	     427 allocs/op
BenchmarkOrderedMap_UnmarshalJSON_valueIsAny-12              	  151137	      7892 ns/op	    5576 B/op	     119 allocs/op
BenchmarkMap_UnmarshalJSON_valueIsAny-12                     	  183066	      6531 ns/op	    4024 B/op	      80 allocs/op
BenchmarkOmW_UnmarshalJSON_valueIsAny-12                     	  163753	      6895 ns/op	    4720 B/op	      95 allocs/op
BenchmarkOmI_UnmarshalJSON_valueIsAny-12                     	   45026	     26518 ns/op	   11048 B/op	     427 allocs/op
PASS
ok  	github.com/sttk/benchmarks_orderedmap/v0_5_0/json	83.455s
```
