module github.com/vcd-simple-blog/apps/backend/api-gateway

go 1.24.0

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/labstack/echo/v4 v4.11.3
)

require (
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
)

replace github.com/vcd-simple-blog/common => ../../../packages/go/common
