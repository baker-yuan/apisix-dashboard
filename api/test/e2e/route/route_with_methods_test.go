package route_test

import (
	"net/http"

	"github.com/apisix/manager-api/test/e2e/base"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("route with methods", func() {
	DescribeTable("test route with methods",
		func(tc base.HttpTestCase) {
			base.RunTestCase(tc)
		},
		Entry("add route with invalid method", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
					 "name": "route1",
					 "uri": "/hello",
					 "methods": ["TEST"],
					 "upstream": {
						 "type": "roundrobin",
						 "nodes": {
							 "` + base.UpstreamIp + `:1980": 1
						 }
					 }
				 }`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusBadRequest,
		}),
		Entry("verify route", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusNotFound,
			Sleep:        base.SleepTime,
		}),
		Entry("add route with valid method", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
					 "name": "route1",
					 "uri": "/hello",
					 "methods": ["GET"],
					 "upstream": {
						 "type": "roundrobin",
						 "nodes": {
							 "` + base.UpstreamIp + `:1980": 1
						 }
					 }
				 }`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify route", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("update same route path", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
					 "name": "route1",
					 "uri": "/hello_",
					 "methods": ["GET"],
					 "upstream": {
						 "type": "roundrobin",
						 "nodes": {
							 "` + base.UpstreamIp + `:1980": 1
						 }
					 }
				 }`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify old route updated", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusNotFound,
			Sleep:        base.SleepTime,
		}),
		Entry("verify new update applied", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello_",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("delete route", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodDelete,
			Path:         "/apisix/admin/routes/r1",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("add route with valid methods", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
					 "name": "route1",
					 "uri": "/hello",
					 "methods": ["GET", "POST", "PUT", "DELETE", "PATCH"],
					 "upstream": {
						 "type": "roundrobin",
						 "nodes": {
							 "` + base.UpstreamIp + `:1980": 1
						 }
					 }
				 }`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify route by post", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodPost,
			Path:         "/hello",
			Body:         `test=test`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("verify route by put", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodPut,
			Path:         "/hello",
			Body:         `test=test`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("verify route by get", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("verify route by delete", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodDelete,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("verify route by patch", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodPatch,
			Path:         "/hello",
			Body:         `test=test`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("update route methods to GET only", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
					 "name": "route1",
					 "uri": "/hello",
					 "methods": ["GET"],
					 "upstream": {
						 "type": "roundrobin",
						 "nodes": {
							 "` + base.UpstreamIp + `:1980": 1
						 }
					 }
				 }`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify post method isn't working now", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodPost,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusNotFound,
			Sleep:        base.SleepTime,
		}),
		Entry("verify PUT method isn't working now", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodPut,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusNotFound,
			Sleep:        base.SleepTime,
		}),
		Entry("verify route by GET only", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("delete route", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodDelete,
			Path:         "/apisix/admin/routes/r1",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("add route with lower case methods", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
					 "name": "route1",
					 "uri": "/hello",
					 "methods": ["GET", "post"],
					 "upstream": {
						 "type": "roundrobin",
						 "nodes": {
							 "` + base.UpstreamIp + `:1980": 1
						 }
					 }
				}`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusBadRequest,
		}),
		Entry("verify route", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusNotFound,
			Sleep:        base.SleepTime,
		}),
		Entry("add route with methods GET", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
					 "name": "route1",
					 "uri": "/hello",
					 "methods": ["GET"],
					 "upstream": {
						 "type": "roundrobin",
						 "nodes": {
							 "` + base.UpstreamIp + `:1980": 1
						 }
					 }
				 }`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify route by get", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("verify route by post", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodPost,
			Path:         "/hello",
			Body:         `test=test`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusNotFound,
			Sleep:        base.SleepTime,
		}),
		Entry("update route methods to POST", base.HttpTestCase{
			Object: base.ManagerApiExpect(),
			Method: http.MethodPut,
			Path:   "/apisix/admin/routes/r1",
			Body: `{
					 "name": "route1",
					 "uri": "/hello",
					 "methods": ["POST"],
					 "upstream": {
						 "type": "roundrobin",
						 "nodes": {
							 "` + base.UpstreamIp + `:1980": 1
						 }
					 }
				 }`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
		}),
		Entry("verify route by get", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodGet,
			Path:         "/hello",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusNotFound,
			Sleep:        base.SleepTime,
		}),
		Entry("verify route by post", base.HttpTestCase{
			Object:       base.APISIXExpect(),
			Method:       http.MethodPost,
			Path:         "/hello",
			Body:         `test=test`,
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			ExpectBody:   "hello world",
			Sleep:        base.SleepTime,
		}),
		Entry("delete route", base.HttpTestCase{
			Object:       base.ManagerApiExpect(),
			Method:       http.MethodDelete,
			Path:         "/apisix/admin/routes/r1",
			Headers:      map[string]string{"Authorization": base.GetToken()},
			ExpectStatus: http.StatusOK,
			Sleep:        base.SleepTime,
		}),
	)
})
