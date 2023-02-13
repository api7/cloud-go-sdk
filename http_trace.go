// Copyright 2022 API7.ai, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cloud

import (
	"fmt"
	"net/http"
	"time"
)

type TraceSeriesKey struct{}

// TraceSeries contains a series of events (ordered by their happening time).
type TraceSeries struct {
	// ID indicates this series.
	ID ID
	// Request is the outgoing request that will send to API7 Cloud.
	// It's the context of the trace series.
	// NOTE: This request object is cloned from the original one, so
	// please avoid reading the Request.Body reader. Instead, use the
	// RequestBody field.
	Request *http.Request
	// RequestBody contains a copy of the outgoing HTTP request body.
	RequestBody []byte
	// Response indicates the response that will receive from API7 Cloud.
	// It's the context of the trace series.
	Response *http.Response
	// ResponseBody contains a copy of the incoming HTTP response body.
	ResponseBody []byte
	// Events contains a series of trace events.
	Events []*TraceEvent
}

func (series *TraceSeries) appendEvent(ev *TraceEvent) {
	series.Events = append(series.Events, ev)
}

// TraceEvent indicates an event occurred during the communication with API7 Cloud.
type TraceEvent struct {
	// Message indicates the event message level.
	Message string
	// HappenedAt indicates the time that this event occurred.
	HappenedAt time.Time
}

func generateEvent(tpl string, a ...any) *TraceEvent {
	// TODO may use another way to generate the message for avoiding the
	// reflection.
	message := fmt.Sprintf(tpl, a...)
	return &TraceEvent{
		Message:    message,
		HappenedAt: time.Now(),
	}
}

// TraceInterface is the interface for http trace.
type TraceInterface interface {
	sendSeries(series *TraceSeries)

	// TraceChan returns a readonly channel which returns *TraceSeries object.
	TraceChan() <-chan *TraceSeries
}

type tracer struct {
	c chan *TraceSeries
}

func newTracer() *tracer {
	return &tracer{
		c: make(chan *TraceSeries),
	}
}

func (t *tracer) sendSeries(series *TraceSeries) {
	t.c <- series
}

func (t *tracer) TraceChan() <-chan *TraceSeries {
	return t.c
}
