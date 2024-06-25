package design

import (
	. "goa.design/goa/v3/dsl"
)

//var JWTAuth = JWTSecurity("jwt", func() {
//	Scope("api:read", "Read access")
//	Scope("api:write", "Write access")
//})

var APIKeyAuth = APIKeySecurity("api_key")
