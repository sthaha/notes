# Golang for Microservices


## Problems it solves

```
  +----------+    +----------+
  | Service  |    | Service  |
  +----------+    +----------+
       |
       |
     .-V-.
    (     )
    |`---'|
    |     |
     `---'
```

Services are small and isolated so it helps solves some **organisational** problems


- Team is too large to work effectively on a shared codebase
- Team `A` blocked on `B` to get work done
- Product velocity; again due to isolation

## Technical problems it creates


Need stable api ...
- No shared Database; so distributed transaction is **HARD**
   - should go for eventual consistency instead ...

- Testing the Mircroservice fleet becomes way too **Hard**
  - solution: - MTTR: optimise mean time to recover
  - canary deployments
  - blue/green deployments

- Service discovery: How does service `A` discover service `B` ?
- Monitoring / instrumentation
- Distributed tracing
- CI/CD

## Architecture

### core service

```go
type Adder interface {
  Sum(a, b int) (int, error)
  Concat(a, b string) (string, error)
}


type BasicService struct{}


func (s BasicService) Sum(a, b int) (int, error) {
  return a + b, nil
}

func (s BasicService) Concat(a, b string) (string, error) {
  return  a + b, nil
}

```

###  Transport: how do we talk to it

- json/http
