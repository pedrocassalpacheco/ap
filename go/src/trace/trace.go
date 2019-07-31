package trace

import (
	instana "../../../go-sensor"
	"context"
	"fmt"
	opentrace "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
	"net/http"
	//"os"
)

const (
	Service = "serviceB"
)

var GoSensor *instana.Sensor
var GoSensorOpts = instana.Options {
	Service:	Service,
	LogLevel:	instana.Error,
}
var OTTracer opentrace.Tracer

func init() {

	// Instana
	GoSensor = instana.NewSensorWithOption(&GoSensorOpts)

	// OT tracing
	opentrace.InitGlobalTracer(instana.NewTracerWithOptions(&GoSensorOpts))

}

func GetTraceFromContext(ctx context.Context) {
	parentSpan,ok := ctx.Value("parentSpan").(opentrace.Span)

	if ok == false {
		OTTracer = nil
		log.Panic("Unable to obtain parent span")
	} else {
		OTTracer = parentSpan.Tracer()
		fmt.Printf("Retrieved tracer from context %T", parentSpan.Tracer())
	}
	return
}

func TraceSQLExecution(r *http.Request, sql string, comments string) opentrace.Span {

	// Parent span
	parentSpan, ok := r.Context().Value("parentSpan").(opentrace.Span)

	if ok == false {
		OTTracer = nil
		log.Panic("Unable to obtain parent span")
	}

	// Does this work?
	tracer := parentSpan.Tracer()

	childSpan := tracer.StartSpan("SQL", opentrace.ChildOf(parentSpan.Context()))
	childSpan.SetTag(string(ext.SpanKind), 	"client")
	childSpan.SetTag(string(ext.DBType), 	"MySQL")
	childSpan.SetTag(string(ext.DBInstance),"people")
	childSpan.SetTag(string(ext.DBUser), 	"pedro")
	childSpan.SetTag(string(ext.DBStatement), sql)
	childSpan.SetBaggageItem("Comments", comments)

	return childSpan

}

func TraceDBConnection(r *http.Request, comments string) opentrace.Span {

	// Parent span
	parentSpan, ok := r.Context().Value("parentSpan").(opentrace.Span)

	if ok == false {
		OTTracer = nil
		log.Panic("Unable to obtain parent span")
	}

	// Does this work?
	tracer := parentSpan.Tracer()

	childSpan := tracer.StartSpan("Connection", opentrace.ChildOf(parentSpan.Context()))
	childSpan.SetTag(string(ext.SpanKind), 		"client")
	childSpan.SetTag(string(ext.DBType),		"MySQL" )
	childSpan.SetTag(string(ext.DBInstance), 	"people")
	childSpan.SetTag(string(ext.DBUser), 		"pedro")
	childSpan.SetTag(string(ext.DBStatement), 	"connect")
	childSpan.SetBaggageItem("Comments", comments)

	return childSpan
}

func TraceFunctionExecution(r *http.Request, f func(), comments string) {

	// Parent span
	parentSpan, ok := r.Context().Value("parentSpan").(opentrace.Span)

	if ok == false {
		OTTracer = nil
		log.Panic("Unable to obtain parent span")
	}

	tracer := parentSpan.Tracer()

	childSpan := tracer.StartSpan("method", opentrace.ChildOf(parentSpan.Context()))
	childSpan.SetTag(string(ext.SpanKind), 	"intermediate")
	childSpan.SetTag(string(ext.Component), "method")
	childSpan.SetBaggageItem("Comments", comments)

	// Execute function
	f()

	childSpan.Finish()

}


