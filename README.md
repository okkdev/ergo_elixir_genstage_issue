# Ergo & Elixir GenStage Issue

Using Elixir 1.13.4 and Go 1.18.1

## How to reproduce

Start Elixir producer:
```sh
elixir --name producer@localhost --cookie cookiesecret producer.exs
```

Start Ergo consumer:
```sh
go run consumer.go
```

The Elixir producer should now throw this error:
```
[error] GenServer Producer terminating
** (FunctionClauseError) no function clause matching in GenStage.maybe_producer_cancel/2
```

In case of success expecting the Ergo consumer to output something like this:
```
[0, 1, 2, 3, 4]
[5, 6, 7, 8, 9]
[10, 11, 12, 13, 14]
[15, 16, 17, 18, 19]
[20, 21, 22, 23, 24]
...
```