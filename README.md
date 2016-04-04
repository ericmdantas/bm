```boom -n 100000 -c 10000 http://localhost:PORT```

- Vanilla Node:

```
Summary:
  Total:        7.9484 secs
  Slowest:      2.9260 secs
  Fastest:      0.0000 secs
  Average:      0.2081 secs
  Requests/sec: 11311.4558
  Total data:   269724 bytes
  Size/request: 3 bytes

Status code distribution:
  [200] 89908 responses

Response time histogram:
  0.000 [5]     |
  0.293 [80314] |????????????????????????????????????????
  0.585 [2334]  |?
  0.878 [1822]  |
  1.170 [1145]  |
  1.463 [1760]  |
  1.756 [920]   |
  2.048 [328]   |
  2.341 [515]   |
  2.633 [644]   |
  2.926 [121]   |

Latency distribution:
  10% in 0.0155 secs
  25% in 0.0256 secs
  50% in 0.1271 secs
  75% in 0.1729 secs
  90% in 0.4510 secs
  95% in 1.1225 secs
  99% in 2.1780 secs
```


- Express Node:

```
Summary:
  Total:        12.4488 secs
  Slowest:      3.5801 secs
  Fastest:      0.0000 secs
  Average:      0.2787 secs
  Requests/sec: 6162.3572
  Total data:   230142 bytes
  Size/request: 3 bytes

Status code distribution:
  [200] 76714 responses

Response time histogram:
  0.000 [3]     |
  0.358 [65473] |????????????????????????????????????????
  0.716 [3690]  |??
  1.074 [989]   |
  1.432 [2000]  |?
  1.790 [2379]  |?
  2.148 [817]   |
  2.506 [919]   |
  2.864 [144]   |
  3.222 [142]   |
  3.580 [158]   |

Latency distribution:
  10% in 0.0230 secs
  25% in 0.0320 secs
  50% in 0.0870 secs
  75% in 0.2590 secs
  90% in 0.7052 secs
  95% in 1.5850 secs
  99% in 2.2660 secs
```


- Koa Node:

```
Summary:
  Total:        11.3190 secs
  Slowest:      3.6445 secs
  Fastest:      0.0000 secs
  Average:      0.2490 secs
  Requests/sec: 7093.7637
  Total data:   240882 bytes
  Size/request: 3 bytes

Status code distribution:
  [200] 80294 responses

Response time histogram:
  0.000 [59]    |
  0.364 [71260] |????????????????????????????????????????
  0.729 [857]   |
  1.093 [1354]  |
  1.458 [1592]  |
  1.822 [2276]  |?
  2.187 [970]   |
  2.551 [1302]  |
  2.916 [559]   |
  3.280 [53]    |
  3.645 [12]    |

Latency distribution:
  10% in 0.0245 secs
  25% in 0.0324 secs
  50% in 0.0861 secs
  75% in 0.1495 secs
  90% in 0.7480 secs
  95% in 1.6218 secs
  99% in 2.4988 secs
```


- Vanilla Go:

```
Summary:
  Total:        3.8404 secs
  Slowest:      3.1720 secs
  Fastest:      0.0000 secs
  Average:      0.1984 secs
  Requests/sec: 25675.9843
  Total data:   295815 bytes
  Size/request: 3 bytes

Status code distribution:
  [200] 98605 responses

Response time histogram:
  0.000 [662]   |
  0.317 [88030] |????????????????????????????????????????
  0.634 [1874]  |
  0.952 [1095]  |
  1.269 [2613]  |?
  1.586 [1098]  |
  1.903 [443]   |
  2.220 [213]   |
  2.538 [1968]  |
  2.855 [470]   |
  3.172 [139]   |

Latency distribution:
  10% in 0.0055 secs
  25% in 0.0160 secs
  50% in 0.0585 secs
  75% in 0.1181 secs
  90% in 0.3236 secs
  95% in 1.1637 secs
  99% in 2.4656 secs
```


- Httprouter Go:

```
Summary:
  Total:        3.6620 secs
  Slowest:      3.1656 secs
  Fastest:      0.0000 secs
  Average:      0.2157 secs
  Requests/sec: 27271.7348
  Total data:   299607 bytes
  Size/request: 3 bytes

Status code distribution:
  [200] 99869 responses

Response time histogram:
  0.000 [398]   |
  0.317 [89687] |????????????????????????????????????????
  0.633 [2057]  |
  0.950 [1422]  |
  1.266 [1358]  |
  1.583 [200]   |
  1.899 [436]   |
  2.216 [661]   |
  2.532 [2346]  |?
  2.849 [633]   |
  3.166 [671]   |

Latency distribution:
  10% in 0.0155 secs
  25% in 0.0490 secs
  50% in 0.0600 secs
  75% in 0.1121 secs
  90% in 0.2822 secs
  95% in 1.2513 secs
  99% in 2.5954 secs
```


- Negroni Go:

```
Summary:
  Total:        4.3422 secs
  Slowest:      3.2363 secs
  Fastest:      0.0000 secs
  Average:      0.1881 secs
  Requests/sec: 22420.2260
  Total data:   292059 bytes
  Size/request: 3 bytes

Status code distribution:
  [200] 97353 responses

Response time histogram:
  0.000 [1213]  |
  0.324 [86039] |????????????????????????????????????????
  0.647 [2608]  |?
  0.971 [1175]  |
  1.295 [1683]  |
  1.618 [982]   |
  1.942 [1680]  |
  2.265 [950]   |
  2.589 [615]   |
  2.913 [365]   |
  3.236 [43]    |

Latency distribution:
  10% in 0.0015 secs
  25% in 0.0105 secs
  50% in 0.0560 secs
  75% in 0.0951 secs
  90% in 0.3648 secs
  95% in 1.2271 secs
  99% in 2.2781 secs
```


- CSharp NancyFx.Self-Hosting:

```
Summary:
  Total:	13.7871 secs
  Slowest:	5.2326 secs
  Fastest:	0.0000 secs
  Average:	1.2363 secs
  Requests/sec:	5270.7107

Status code distribution:
  [200]	72668 responses

Response time histogram:
  0.000 [4]	|
  0.523 [7631]	|∎∎∎∎∎∎∎∎
  1.047 [37979]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1.570 [8677]	|∎∎∎∎∎∎∎∎∎
  2.093 [8075]	|∎∎∎∎∎∎∎∎
  2.616 [3974]	|∎∎∎∎
  3.140 [2834]	|∎∎
  3.663 [915]	|
  4.186 [1558]	|∎
  4.709 [568]	|
  5.233 [453]	|

Latency distribution:
  10% in 0.5092 secs
  25% in 0.7382 secs
  50% in 0.9143 secs
  75% in 1.5815 secs
  90% in 2.4777 secs
  95% in 3.0999 secs
  99% in 4.4563 secs
```


- CSharp WebApi.Owin.SelfHost

```
Summary:
  Total:	14.6994 secs
  Slowest:	4.9775 secs
  Fastest:	0.0000 secs
  Average:	1.2791 secs
  Requests/sec:	4842.0316
  Total data:	284700 bytes
  Size/request:	4 bytes

Status code distribution:
  [200]	71175 responses

Response time histogram:
  0.000 [2]	|
  0.498 [6080]	|∎∎∎∎∎∎∎
  0.995 [32556]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1.493 [12318]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1.991 [5799]	|∎∎∎∎∎∎∎
  2.489 [5930]	|∎∎∎∎∎∎∎
  2.986 [5716]	|∎∎∎∎∎∎∎
  3.484 [1662]	|∎∎
  3.982 [619]	|
  4.480 [484]	|
  4.977 [9]	|

Latency distribution:
  10% in 0.5282 secs
  25% in 0.7552 secs
  50% in 0.9593 secs
  75% in 1.7535 secs
  90% in 2.5968 secs
  95% in 2.9179 secs
  99% in 3.8352 secs
```  


- Pure python

```
Summary:
  Total:	25.5676 secs
  Slowest:	4.2101 secs
  Fastest:	0.9872 secs
  Average:	1.7560 secs
  Requests/sec:	380.0513

Status code distribution:
  [200]	9717 responses

Response time histogram:
  0.987 [1]	|
  1.309 [1792]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1.632 [2670]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1.954 [1681]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2.276 [2616]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2.599 [446]	|∎∎∎∎∎∎
  2.921 [283]	|∎∎∎∎
  3.243 [171]	|∎∎
  3.565 [31]	|
  3.888 [1]	|
  4.210 [25]	|

Latency distribution:
  10% in 1.0313 secs
  25% in 1.5239 secs
  50% in 1.7353 secs
  75% in 2.0630 secs
  90% in 2.2717 secs
  95% in 2.6155 secs
  99% in 3.1570 secs
```  