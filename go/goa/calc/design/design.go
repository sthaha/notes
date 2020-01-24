package design

import . "goa.design/goa/v3/dsl"

var _ = API("calc", func() {
	Title("Calculator Service")
	// Must define Server
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var calc = Service("calc", func() {
	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "left")
			Field(2, "b", Int, "right")

			Required("a", "b")
		})

		Result(Int)

		HTTP(func() {
			GET("/add/{a}/{b}")
		})

	})
})
