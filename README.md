# Logging Test for AWS Lambda provided.al2023 Runtime

This repository contains test code to examine how the `provided.al2023` runtime handles logs.

The test results revealed that the value of the `level` key in JSON-formatted logs is treated as the log level. 
`LEVEL` and `Level` are invalid. Valid values are `error`, `Error`, or `ERROR`, `erroR` is invalid.

See also: https://docs.aws.amazon.com/lambda/latest/dg/monitoring-cloudwatchlogs-advanced.html#monitoring-cloudwatchlogs-log-level

## Result
```
{
    "input": "{\"Context\":{\"Context\":{\"Context\":{}}}}",
    "level": "error",
    "msg": "error",
    "time": "2024-09-11T12:06:45Z"
}
```

```
{"error":"invalid character '}' after top-level value","input":"{\"Context\":{\"Context\":{\"Context\":{}}}}}","level":"error","msg":"input json unmarshal error","time":"2024-09-11T12:06:45Z"}
```

```
{
    "msg": "mylog",
    "out": "stderr",
    "level": "error"
}
```

```
{
    "msg": "mylog",
    "out": "stdout",
    "level": "error"
}
```

```
{
    "msg": "mylog",
    "out": "stdout",
    "level": "ERROR"
}
```

```
{
    "msg": "mylog",
    "out": "stdout",
    "level": "Error"
}
```