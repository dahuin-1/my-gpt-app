package design

import (
	. "goa.design/goa/v3/dsl"
)

var ErrorType = Type("ErrorType", func() {
	Attribute("error_code", String, func() {
		Example("3000")
	})
	Attribute("message", String, func() {
		Example("something error")
	})
	Attribute("status", String, func() {
		Meta("struct:error:name")
		Meta("struct:tag:json", "-")
		Example("NotAcceptable")
	})
	Required("error_code", "message", "status")
})
