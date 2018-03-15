# fproto-wrap-headers

[fproto-wrap](https://github.com/RangelReale/fproto-wrap) converter for HTTP-like header values. (map[string][]string).

Currently there is only a Golang wrapper generator.

### proto

```proto
message Headers {
    message Values {
        repeated string values = 1;
    }
    map<string, Values> headers = 1;
}
```

In Golang, this would convert to and from "map[string][]string".

### author

Rangel Reale (rangelspam@gmail.com)

