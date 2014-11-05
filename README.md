fenwick [![GoDoc](https://godoc.org/github.com/abursavich/fenwick?status.svg)](https://godoc.org/github.com/abursavich/fenwick)
====

Package fenwick provides an implementation of a Fenwick Tree or Binary Indexed Tree which provides efficient manipulation and calculation of prefix sums of a table of values.

Where n is the size of the table, the tree requires O(n) space and operations take O(log n) time.

Example:

```Go
tree := fenwick.NewTree(0, 10)
tree.Add(0, 3)
tree.Add(1, -2)
tree.Add(5, 10)
tree.Add(7, 3)
tree.Add(7, 4)
tree.Add(10, 100)
fmt.Println(tree.Sum(0))       // 3
fmt.Println(tree.Sum(1))       // 1
fmt.Println(tree.Sum(6))       // 11
fmt.Println(tree.Value(5))     // 10
fmt.Println(tree.Value(7))     // 7
fmt.Println(tree.Range(1, 5))  // 8
fmt.Println(tree.Range(7, 10)) // 107
```
