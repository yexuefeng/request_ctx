package requestctx

import (
	"fmt"
	"testing"
)

func TestNext(t *testing.T) {
	r := NewRequestCtx()
	f1 := func(r *RequestCtx) {
		fmt.Println("start f1")
		r.Next()
		fmt.Println("end f1")
	}
	f2 := func(r *RequestCtx) {
		fmt.Println("start f2")
		r.Next()
		fmt.Println("end f2")
	}
	r.Use(f1, f2)
	r.Next()
	if r.GetIndex() != 4 {
		t.Errorf("index value is wrong, index val: %d", r.GetIndex())
	}
}
