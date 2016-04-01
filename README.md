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


- CSharp - NancyFx.SelfHosting

```
Summary:
  Total:	15.9582 secs
  Slowest:	7.1564 secs
  Fastest:	0.0000 secs
  Average:	1.4180 secs
  Requests/sec:	3890.6039

Status code distribution:
  [200]	62087 responses

Response time histogram:
  0.000 [1]	|
  0.716 [11187]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  1.431 [28361]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2.147 [8188]	|∎∎∎∎∎∎∎∎∎∎∎
  2.863 [10084]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  3.578 [2265]	|∎∎∎
  4.294 [952]	|∎
  5.010 [543]	|
  5.725 [291]	|
  6.441 [206]	|
  7.156 [9]	|

Latency distribution:
  10% in 0.3861 secs
  25% in 0.8112 secs
  50% in 1.1062 secs
  75% in 2.0604 secs
  90% in 2.6815 secs
  95% in 3.1766 secs
  99% in 4.7580 secs
```
