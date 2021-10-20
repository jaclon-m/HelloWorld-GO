package headerz

import "httpserver-demo/handler"

func init() {
	headerz := HeaderHandler{}
	handler.DefaultDispatchHandler.AddHandler(&headerz)
}
