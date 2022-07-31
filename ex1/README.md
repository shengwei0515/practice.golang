# Request

1. 
command example
```
./ex1 -h
```

result
```
Usage of ./ex1:
    -u string
        url to parse
```

2. 
command example
```
./ex1 -u https://example.com:443/path?arg=value
```

result
```
Schema: https
Hostname: example.com
Port: 443
Path: /path
Query String: arg=value
```