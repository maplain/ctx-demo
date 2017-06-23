This repo shows the common usage of golang/context and some concurrent patterns.

In includes a server which will start a user specified number of workers to do
its job. The whole thing is cancellable:

1. The initiazliation of workers can be cancelled, if cancelled in the middle of
initialization of all workers, the rest of them will not be initialized;
2. The worker can be cancelled, if cancelled while processing, it'll drop its
job immediately;

What's more:
1. any failure of any worker will cancel the whole job;
2. our server will be able to return any partial finished results if failure
appeared and some workers already successfully returned their part of results

Concurrent pattern:
Based on example here: https://blog.golang.org/pipelines with improvement:
1. downstream worker is able to notify upstream node to cancel the whole job which is important in reality and tricky to implement correctly because:
   1. generally only sender close the sending channel, but in this case, sender of errChan is downstream worker not upstream routine;
     think about the quote from above blog:
     ```
     There is a pattern to our pipeline functions:
          1. stages close their outbound channels when all the send operations are done.
          2. stages keep receiving values from inbound channels until those channels are closed.
     ```
   2. we also want to collect and return all returned results when whole job is cancelled in the middle


try:
```
cd pkg/example/
go build main.go
./main --help
./main -wn 20 -wd 20 -lo 50 -hi 100
```

start another windown is terminal
```
curl -X POST http://127.0.0.1:8080/number\?q\=1\&timeout\=100ms
```

change parameters and see what you got

Also:
1. pkg/example, every file except main.go can be directly 'go run *', try it yourself;
2. pkg/printer includes a tiny client-server style printer which displays your output to terminal. itself is also an
example.
