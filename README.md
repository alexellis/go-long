# go-long

Custom functions for testing timeouts in OpenFaaS

Example for node18:

```
time curl -s -i http://127.0.0.1:8080/function/node18-sleep \
-H "x-duration-ms: 120000"
```

Example for Go:

Edit `handler_wait_duration` in stack.yml to a Go duration i.e. `120s`, then:

```
time curl -s -i http://127.0.0.1:8080/function/go-long
```

Example for Python:

Edit `handler_wait_duration_secs` to a number of seconds in stack.yml, i.e. `120` then:

```
time curl -s -i http://127.0.0.1:8080/function/py-long
```
