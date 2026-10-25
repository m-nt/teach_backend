# Booking rule CLI

```sh
go run .
go test ./...
```

Expected output:

```text
attempt 1: accepted; remaining=1
attempt 2: accepted; remaining=0
attempt 3: rejected (class is full)
```

Follow the rule before studying the tests. A failed reservation returns the unchanged booked count. `findSession` returns a value copy; changing it does not update the slice. Intermediate lesson 01 revisits this deliberately visible ownership detail.

Exercise: add a cancellation rule that rejects an empty class and test the zero/one-booking boundary. No input persistence or interactive prompt is implemented here.
