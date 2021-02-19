package recursion

import  "testing"

func TestFactorial(t *testing.T) {
    res := factorial(5)
    t.Logf("%d\n", res) // 120
}
