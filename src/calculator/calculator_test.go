package calculator

import "testing"

var numberset = [] struct {
	x int
	y int
	result int
}{
	{1,2,3},
	{2,2,4},
	{3,3,6},
}

func TestAdd(t *testing.T) {
	//result := Add(1, 3)
	//if (result !=4 ) {
		//t.Fail()
		//t.Errorf("Expected 1 + 3 == 4, got %d instead", result)
	//}
	// Index and set
	for _, set := range numberset {
		aresult := Add(set.x, set.y)
		if aresult != set.result {
			t.Errorf("Expected %d + %d == %d, got %d instead.", set.x, set.y, set.result, aresult)
		}
	}
}
