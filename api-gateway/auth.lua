local jwt = require "resty.jwt"

local jwt_token = ngx.req.get_headers()["Authorization"]
local file = io.open("/run/secrets/pub_key", "r")
local jwt_result = jwt:verify(file:read("*a"), jwt_token)

if jwt_result.verified == false 
then
    ngx.exit(401)
end

