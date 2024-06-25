package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = Service("common", func() {
	Description("service exposes public endpoints that defines CORS policy.")
	cors.Origin("*", func() {
		cors.Headers("*")
		cors.Methods("GET", "POST")
		cors.Expose("X-Time", "X-Api-Version")
		cors.MaxAge(100)
		cors.Credentials()
	})

	Method("health", func() {
		HTTP(func() {
			GET("/health")
			Response(StatusOK)
		})
	})
})
