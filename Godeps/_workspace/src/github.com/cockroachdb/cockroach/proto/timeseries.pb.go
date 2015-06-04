// Code generated by protoc-gen-gogo.
// source: cockroach/proto/timeseries.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

// TimeSeriesQueryAggregator describes a set of aggregation functions which are
// applied to data points before returning them as part of a query.
//
// Cockroach does not store data points at full fidelity, instead "downsampling"
// data points into fixed-length sample periods. The value returned for each
// sample period is equivalent to applying the supplied aggregator function to
// the original data points that fell within the sample period.
type TimeSeriesQueryAggregator int32

const (
	// AVG returns the average value of points within the sample period.
	TimeSeriesQueryAggregator_AVG TimeSeriesQueryAggregator = 1
	// AVG_RATE returns the rate of change of the average over the sample period's
	// duration.  This is computed via linear regression with the previous sample
	// period's average value.
	TimeSeriesQueryAggregator_AVG_RATE TimeSeriesQueryAggregator = 2
)

var TimeSeriesQueryAggregator_name = map[int32]string{
	1: "AVG",
	2: "AVG_RATE",
}
var TimeSeriesQueryAggregator_value = map[string]int32{
	"AVG":      1,
	"AVG_RATE": 2,
}

func (x TimeSeriesQueryAggregator) Enum() *TimeSeriesQueryAggregator {
	p := new(TimeSeriesQueryAggregator)
	*p = x
	return p
}
func (x TimeSeriesQueryAggregator) String() string {
	return proto1.EnumName(TimeSeriesQueryAggregator_name, int32(x))
}
func (x *TimeSeriesQueryAggregator) UnmarshalJSON(data []byte) error {
	value, err := proto1.UnmarshalJSONEnum(TimeSeriesQueryAggregator_value, data, "TimeSeriesQueryAggregator")
	if err != nil {
		return err
	}
	*x = TimeSeriesQueryAggregator(value)
	return nil
}

// TimeSeriesDatapoint is a single point of time series data; a value associated
// with a timestamp.
type TimeSeriesDatapoint struct {
	// The timestamp when this datapoint is located, expressed in nanoseconds
	// since the unix epoch.
	TimestampNanos int64 `protobuf:"varint,1,opt,name=timestamp_nanos" json:"timestamp_nanos"`
	// A floating point representation of the value of this datapoint.
	Value            float64 `protobuf:"fixed64,2,opt,name=value" json:"value"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TimeSeriesDatapoint) Reset()         { *m = TimeSeriesDatapoint{} }
func (m *TimeSeriesDatapoint) String() string { return proto1.CompactTextString(m) }
func (*TimeSeriesDatapoint) ProtoMessage()    {}

func (m *TimeSeriesDatapoint) GetTimestampNanos() int64 {
	if m != nil {
		return m.TimestampNanos
	}
	return 0
}

func (m *TimeSeriesDatapoint) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// TimeSeriesData is a set of measurements of a single named variable at
// multiple points in time. This message contains a name and a source which, in
// combination, uniquely identify the time series being measured. Measurement
// data is represented as a repeated set of TimeSeriesDatapoint messages.
type TimeSeriesData struct {
	// A string which uniquely identifies the variable from which this data was
	// measured.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name"`
	// A string which identifies the unique source from which the variable was measured.
	Source string `protobuf:"bytes,2,opt,name=source" json:"source"`
	// Datapoints representing one or more measurements taken from the variable.
	Datapoints       []*TimeSeriesDatapoint `protobuf:"bytes,3,rep,name=datapoints" json:"datapoints,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TimeSeriesData) Reset()         { *m = TimeSeriesData{} }
func (m *TimeSeriesData) String() string { return proto1.CompactTextString(m) }
func (*TimeSeriesData) ProtoMessage()    {}

func (m *TimeSeriesData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimeSeriesData) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *TimeSeriesData) GetDatapoints() []*TimeSeriesDatapoint {
	if m != nil {
		return m.Datapoints
	}
	return nil
}

// TimeSeriesQueryRequest is the standard incoming time series query request
// accepted from cockroach clients.
type TimeSeriesQueryRequest struct {
	// A timestamp in nanoseconds which defines the early bound of the time span
	// for this query.
	StartNanos int64 `protobuf:"varint,1,opt,name=start_nanos" json:"start_nanos"`
	// A timestamp in nanoseconds which defines the late bound of the time span
	// for this query. Must be greater than start_nanos.
	EndNanos int64 `protobuf:"varint,2,opt,name=end_nanos" json:"end_nanos"`
	// A set of Queries for this request. A request must have at least one
	// Query.
	Queries          []TimeSeriesQueryRequest_Query `protobuf:"bytes,3,rep,name=queries" json:"queries"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *TimeSeriesQueryRequest) Reset()         { *m = TimeSeriesQueryRequest{} }
func (m *TimeSeriesQueryRequest) String() string { return proto1.CompactTextString(m) }
func (*TimeSeriesQueryRequest) ProtoMessage()    {}

func (m *TimeSeriesQueryRequest) GetStartNanos() int64 {
	if m != nil {
		return m.StartNanos
	}
	return 0
}

func (m *TimeSeriesQueryRequest) GetEndNanos() int64 {
	if m != nil {
		return m.EndNanos
	}
	return 0
}

func (m *TimeSeriesQueryRequest) GetQueries() []TimeSeriesQueryRequest_Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

// Each Query defines a specific metric to query over the time span of
// this request.
type TimeSeriesQueryRequest_Query struct {
	// The name of the time series to query.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name"`
	// The aggregation function to apply to points in the result.
	Aggregator       *TimeSeriesQueryAggregator `protobuf:"varint,2,opt,name=aggregator,enum=cockroach.proto.TimeSeriesQueryAggregator,def=1" json:"aggregator,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *TimeSeriesQueryRequest_Query) Reset()         { *m = TimeSeriesQueryRequest_Query{} }
func (m *TimeSeriesQueryRequest_Query) String() string { return proto1.CompactTextString(m) }
func (*TimeSeriesQueryRequest_Query) ProtoMessage()    {}

const Default_TimeSeriesQueryRequest_Query_Aggregator TimeSeriesQueryAggregator = TimeSeriesQueryAggregator_AVG

func (m *TimeSeriesQueryRequest_Query) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimeSeriesQueryRequest_Query) GetAggregator() TimeSeriesQueryAggregator {
	if m != nil && m.Aggregator != nil {
		return *m.Aggregator
	}
	return Default_TimeSeriesQueryRequest_Query_Aggregator
}

// TimeSeriesQueryResponse is the standard response for time series queries
// returned to cockroach clients.
type TimeSeriesQueryResponse struct {
	// A set of Results; there will be one result for each Query in the matching
	// TimeSeriesQueryRequest, in the same order. A Result will be present for
	// each Query even if there are zero datapoints to return.
	Results          []*TimeSeriesQueryResponse_Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *TimeSeriesQueryResponse) Reset()         { *m = TimeSeriesQueryResponse{} }
func (m *TimeSeriesQueryResponse) String() string { return proto1.CompactTextString(m) }
func (*TimeSeriesQueryResponse) ProtoMessage()    {}

func (m *TimeSeriesQueryResponse) GetResults() []*TimeSeriesQueryResponse_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

// Result is the data returned from a single metric query over a time span.
type TimeSeriesQueryResponse_Result struct {
	// A string which uniquely identifies the variable from which this data was
	// measured.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name"`
	// A list of sources from which the data was aggregated.
	Sources []string `protobuf:"bytes,2,rep,name=sources" json:"sources"`
	// The aggregation function applied to points in the result.
	Aggregator *TimeSeriesQueryAggregator `protobuf:"varint,3,opt,name=aggregator,enum=cockroach.proto.TimeSeriesQueryAggregator,def=1" json:"aggregator,omitempty"`
	// Datapoints describing the queried data.
	Datapoints       []*TimeSeriesDatapoint `protobuf:"bytes,4,rep,name=datapoints" json:"datapoints,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *TimeSeriesQueryResponse_Result) Reset()         { *m = TimeSeriesQueryResponse_Result{} }
func (m *TimeSeriesQueryResponse_Result) String() string { return proto1.CompactTextString(m) }
func (*TimeSeriesQueryResponse_Result) ProtoMessage()    {}

const Default_TimeSeriesQueryResponse_Result_Aggregator TimeSeriesQueryAggregator = TimeSeriesQueryAggregator_AVG

func (m *TimeSeriesQueryResponse_Result) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimeSeriesQueryResponse_Result) GetSources() []string {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *TimeSeriesQueryResponse_Result) GetAggregator() TimeSeriesQueryAggregator {
	if m != nil && m.Aggregator != nil {
		return *m.Aggregator
	}
	return Default_TimeSeriesQueryResponse_Result_Aggregator
}

func (m *TimeSeriesQueryResponse_Result) GetDatapoints() []*TimeSeriesDatapoint {
	if m != nil {
		return m.Datapoints
	}
	return nil
}

func init() {
	proto1.RegisterEnum("cockroach.proto.TimeSeriesQueryAggregator", TimeSeriesQueryAggregator_name, TimeSeriesQueryAggregator_value)
}
