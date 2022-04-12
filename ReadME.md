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

- Write -  O(1)
- Read - O(K)

### Space Complexity
- O(N+K)

## Optimization
The time complexity of the reading is always O(K) since we are required to go through all the lists 
in order from the first priority to the last, however I found a simple solution to optimize it by 
adding the firstAvailable parameter which saves the first available list.

This solution reduces read time to O(1) in best case scenario, and statistically it should lower the 
overall read time on average.