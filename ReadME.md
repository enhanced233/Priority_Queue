# Priority Queue 
A simple to use priority queue with K priorities (K determined by user) implemented by GO.

	- Data type is an empty interface to accommodate all data types, therefore its up to the user to convert them back and check for errors (as shown in usage).
	- Safe for concurrency.
	- Inserting multiple items with the same priority number will fetch them in order of FIFO. 
## Example

### Import

```go
import "github.com/enhanced233/Priority_Queue/pq"
```

### Usage

```go
q := pq.NewPq(4)	
q.Insert(byte('e'), 1)
q.Insert(byte('l'), 1)
q.Insert(byte('o'), 3)
q.Insert(byte('H'), 0)
q.Insert(byte('l'), 2)
var s string
for !q.IsEmpty() {
	data, ok := q.Fetch().(byte)
	if ok {
		s = s + string(data)
	}
}
fmt.Println(s)
```

## Complexity
For N data points and a queue with K priorities:
### Time Complexity
	Insert -  O(1) 

	Fetch - O(K) with an asymptotic complexity of O(1)

	IsEmpty - O(K) with an asymptotic complexity of O(1)

### Space Complexity
	O(N+K)
