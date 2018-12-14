# strset

Go library providing methods for string sets

[![GoDoc](https://godoc.org/gitlab.com/shaydo/strset?status.svg)](http://godoc.org/gitlab.com/shaydo/strset)

```go
	set := New()
        // add elements to the set
	set.Add("foo", "bar", "baz")
        // check if the set contains the string
	if set.Contains("bar") {
		fmt.Println("bar is here")
	}
        // remove string from the set
	set.Del("baz")
        // get all elements of the set as string slice
	elements := set.AsSlice()
	sort.Strings(elements)
	fmt.Println(elements)
```
