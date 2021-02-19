package stringmatch

import (
    "fmt"
    "testing"
)

func TestViolenceSearch(t *testing.T) {
    str := "CBC DCABCABABCABD BBCCA"
    match := "ABCABD"
    index := violenceSearch(str, match)
    fmt.Printf("%s 在 %s 中的位置为 %d\n", match, str, index)
}
