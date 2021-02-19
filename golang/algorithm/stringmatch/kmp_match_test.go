package stringmatch

import (
    "fmt"
    "testing"
)

func TestKmpSearch(t *testing.T) {
    str := "CBC DCABCABABCABD BBCCA"
    match := "ABCABD"
    index := kmpSearch(str, match)
    fmt.Printf("%s 在 %s 中的位置为 %d\n", match, str, index)
}
