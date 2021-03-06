package promgrpc

import (
	"context"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
)

type StatsHandlerCollector interface {
	// Init reallocates possible dimensions for given metric.
	Init(map[string]grpc.ServiceInfo) error

	stats.Handler
	prometheus.Collector
}

type StatsHandler struct {
	handlers []StatsHandlerCollector
}

var _ StatsHandlerCollector = &StatsHandler{}

// NewStatsHandler allows to pass various number of handlers.
func NewStatsHandler(handlers ...StatsHandlerCollector) *StatsHandler {
	return &StatsHandler{
		handlers: handlers,
	}
}

func ClientStatsHandler() *StatsHandler {
	return defaultStatsHandler(Client)
}

func ServerStatsHandler() *StatsHandler {
	return defaultStatsHandler(Server)
}

func defaultStatsHandler(sub Subsystem) *StatsHandler {
	return NewStatsHandler(
		NewConnectionsStatsHandler(sub, NewConnectionsGaugeVec(sub)),
		NewRequestsTotalStatsHandler(sub, NewRequestsTotalCounterVec(sub)),
		NewRequestsStatsHandler(sub, NewRequestsGaugeVec(sub)),
		NewRequestDurationStatsHandler(sub, NewRequestDurationHistogramVec(sub)),
		NewResponsesTotalStatsHandler(sub, NewResponsesTotalCounterVec(sub)),
		NewMessagesReceivedTotalStatsHandler(sub, NewMessagesReceivedTotalCounterVec(sub)),
		NewMessagesSentTotalStatsHandler(sub, NewMessagesSentTotalCounterVec(sub)),
		NewMessageSentSizeStatsHandler(sub, NewMessageSentSizeHistogramVec(sub)),
		NewMessageReceivedSizeStatsHandler(sub, NewMessageReceivedSizeHistogramVec(sub)),
	)
}

// Init implements StatsHandlerCollector interface.
// TODO: implement
func (h *StatsHandler) Init(info map[string]grpc.ServiceInfo) error {
	return nil
}

func (h *StatsHandler) TagRPC(ctx context.Context, inf *stats.RPCTagInfo) context.Context {
	service, method := split(inf.FullMethodName)

	ctx = context.WithValue(ctx, tagRPCKey, prometheus.Labels{
		labelFailFast: strconv.FormatBool(inf.FailFast),
		labelService:  service,
		labelMethod:   method,
	})

	for _, c := range h.handlers {
		ctx = c.TagRPC(ctx, inf)
	}
	return ctx
}

// HandleRPC processes the RPC stats.
func (h *StatsHandler) HandleRPC(ctx context.Context, sts stats.RPCStats) {
	for _, c := range h.handlers {
		c.HandleRPC(ctx, sts)
	}
}

func (h *StatsHandler) TagConn(ctx context.Context, inf *stats.ConnTagInfo) context.Context {
	for _, c := range h.handlers {
		ctx = c.TagConn(ctx, inf)
	}
	return ctx
}

// HandleConn processes the Conn stats.
func (h *StatsHandler) HandleConn(ctx context.Context, sts stats.ConnStats) {
	for _, c := range h.handlers {
		c.HandleConn(ctx, sts)
	}
}

// Describe implements prometheus Collector interface.
func (h *StatsHandler) Describe(in chan<- *prometheus.Desc) {
	for _, c := range h.handlers {
		c.Describe(in)
	}
}

// Collect implements prometheus Collector interface.
func (h *StatsHandler) Collect(in chan<- prometheus.Metric) {
	for _, c := range h.handlers {
		c.Collect(in)
	}
}

type baseStatsHandler struct {
	subsystem Subsystem
	collector prometheus.Collector
}

// HandleRPC implements stats Handler interface.
func (h *baseStatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	service, method := split(info.FullMethodName)

	return context.WithValue(ctx, tagRPCKey, prometheus.Labels{
		labelFailFast:        strconv.FormatBool(info.FailFast),
		labelService:         service,
		labelMethod:          method,
		labelClientUserAgent: userAgent(ctx),
	})
}

// TagRPC implements stats Handler interface.
func (h *baseStatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return context.WithValue(ctx, tagConnKey, prometheus.Labels{
		labelRemoteAddr:      info.RemoteAddr.String(),
		labelLocalAddr:       info.LocalAddr.String(),
		labelClientUserAgent: userAgent(ctx),
	})
}

// HandleRPC implements stats Handler interface.
func (h *baseStatsHandler) HandleConn(ctx context.Context, stat stats.ConnStats) {
}

// HandleRPC implements stats Handler interface.
func (h *baseStatsHandler) HandleRPC(ctx context.Context, stat stats.RPCStats) {
}

// Describe implements prometheus Collector interface.
func (h *baseStatsHandler) Describe(in chan<- *prometheus.Desc) {
	h.collector.Describe(in)
}

// Collect implements prometheus Collector interface.
func (h *baseStatsHandler) Collect(in chan<- prometheus.Metric) {
	h.collector.Collect(in)
}
