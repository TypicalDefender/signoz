package app

import (
	"context"

	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/util/stats"
	"go.signoz.io/query-service/model"
)

type Reader interface {
	GetChannel(id string) (*model.ChannelItem, *model.ApiError)
	GetChannels() (*[]model.ChannelItem, *model.ApiError)
	DeleteChannel(id string) *model.ApiError
	CreateChannel(receiver *model.Receiver) (*model.Receiver, *model.ApiError)
	EditChannel(receiver *model.Receiver, id string) (*model.Receiver, *model.ApiError)

	GetRule(id string) (*model.RuleResponseItem, *model.ApiError)
	ListRulesFromProm() (*model.AlertDiscovery, *model.ApiError)
	CreateRule(alert string) *model.ApiError
	EditRule(alert string, id string) *model.ApiError
	DeleteRule(id string) *model.ApiError

	GetInstantQueryMetricsResult(ctx context.Context, query *model.InstantQueryMetricsParams) (*promql.Result, *stats.QueryStats, *model.ApiError)
	GetQueryRangeResult(ctx context.Context, query *model.QueryRangeParams) (*promql.Result, *stats.QueryStats, *model.ApiError)
	GetServiceOverview(ctx context.Context, query *model.GetServiceOverviewParams) (*[]model.ServiceOverviewItem, error)
	GetServices(ctx context.Context, query *model.GetServicesParams) (*[]model.ServiceItem, error)
	// GetApplicationPercentiles(ctx context.Context, query *model.ApplicationPercentileParams) ([]godruid.Timeseries, error)
	GetTopEndpoints(ctx context.Context, query *model.GetTopEndpointsParams) (*[]model.TopEndpointsItem, error)
	GetUsage(ctx context.Context, query *model.GetUsageParams) (*[]model.UsageItem, error)
	GetServicesList(ctx context.Context) (*[]string, error)
	GetServiceMapDependencies(ctx context.Context, query *model.GetServicesParams) (*[]model.ServiceMapDependencyResponseItem, error)
	GetTTL(ctx context.Context, ttlParams *model.GetTTLParams) (*model.GetTTLResponseItem, *model.ApiError)
	GetSpanFilters(ctx context.Context, query *model.SpanFilterParams) (*model.SpanFiltersResponse, *model.ApiError)
	GetTagFilters(ctx context.Context, query *model.TagFilterParams) (*[]model.TagFilters, *model.ApiError)
	GetTagValues(ctx context.Context, query *model.TagFilterParams) (*[]model.TagValues, *model.ApiError)
	GetFilteredSpans(ctx context.Context, query *model.GetFilteredSpansParams) (*model.GetFilterSpansResponse, *model.ApiError)
	GetFilteredSpansAggregates(ctx context.Context, query *model.GetFilteredSpanAggregatesParams) (*model.GetFilteredSpansAggregatesResponse, *model.ApiError)

	GetErrors(ctx context.Context, params *model.GetErrorsParams) (*[]model.Error, *model.ApiError)
	GetErrorForId(ctx context.Context, params *model.GetErrorParams) (*model.ErrorWithSpan, *model.ApiError)
	GetErrorForType(ctx context.Context, params *model.GetErrorParams) (*model.ErrorWithSpan, *model.ApiError)
	// Search Interfaces
	SearchTraces(ctx context.Context, traceID string) (*[]model.SearchSpansResult, error)

	// Setter Interfaces
	SetTTL(ctx context.Context, ttlParams *model.TTLParams) (*model.SetTTLResponseItem, *model.ApiError)

	GetMetricAutocompleteMetricNames(ctx context.Context, matchText string) (*[]string, *model.ApiError)
	GetMetricAutocompleteTagKey(ctx context.Context, params *model.MetricAutocompleteTagParams) (*[]string, *model.ApiError)
	GetMetricAutocompleteTagValue(ctx context.Context, params *model.MetricAutocompleteTagParams) (*[]string, *model.ApiError)
}
