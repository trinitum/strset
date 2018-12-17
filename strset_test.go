package strset

import (
	"fmt"
	"sort"
	"testing"
)

func Example() {
	set := New()
	set.Add("foo", "bar", "baz")
	if set.Contains("bar") {
		fmt.Println("bar is here")
	}
	set.Del("baz")
	elements := set.AsSlice()
	sort.Strings(elements)
	fmt.Println(elements)
	// Output:
	// bar is here
	// [bar foo]
}

func TestSet(t *testing.T) {
	s1 := New()
	if s1.Contains("foo") {
		t.Fatal("s1 contains foo")
	}

	s1.Add("foo")
	if !s1.Contains("foo") {
		t.Fatal("s1 doesn't contain foo")
	}

	s1.Add("bar", "baz")
	if !s1.Contains("bar") {
		t.Fatal("s1 doesn't contain bar")
	}

	s1.Del("baz")
	if s1.Contains("baz") {
		t.Fatal("s1 contains baz")
	}

	s2 := New("bar", "boo")
	if len(s2) != 2 {
		t.Fatalf("expected len of s2 to be 2")
	}
	s2s := s2.AsSlice()
	sort.Strings(s2s)
	if len(s2s) != 2 || s2s[0] != "bar" || s2s[1] != "boo" {
		t.Fatalf("expected to get bar and boo, but got %v", s2s)
	}

	s3 := s1.Intersect(s2)
	s3s := s3.AsSlice()
	if len(s3s) != 1 || s3s[0] != "bar" {
		t.Fatalf("expected intersection to contain bar, but got %v", s3s)
	}
}
