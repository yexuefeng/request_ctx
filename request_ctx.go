package requestctx

// RequestCtx request context
type RequestCtx struct {
	index       int8
	middleWares []func(o *RequestCtx)
}

func NewRequestCtx() *RequestCtx {
	return &RequestCtx{
		index: -1,
	}
}

func (r *RequestCtx) Next() {
	r.index++
	for r.index < int8(len(r.middleWares)) {
		r.middleWares[r.index](r)
		r.index++
	}
}

func (r *RequestCtx) Use(args ...func(r *RequestCtx)) {
	r.middleWares = append(r.middleWares, args...)
}

func (r *RequestCtx) GetIndex() int8 {
	return r.index
}
