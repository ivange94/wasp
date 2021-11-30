package metrics

import (
	"fmt"
	"time"

	"github.com/iotaledger/wasp/packages/iscp"
	"github.com/prometheus/client_golang/prometheus"
)

type StateManagerMetrics interface {
	RecordBlockSize(blockIndex uint32, size float64)
}

type ChainMetrics interface {
	CountMessages()
	CountRequestAckMessages()
	MempoolMetrics
	ConsensusMetrics
	StateManagerMetrics
}

type MempoolMetrics interface {
	CountOffLedgerRequestIn()
	CountOnLedgerRequestIn()
	CountRequestOut()
	RecordRequestProcessingTime(iscp.RequestID, time.Duration)
	CountBlocksPerChain()
}

type ConsensusMetrics interface {
	RecordVMRunTime(time.Duration)
	CountVmRuns()
}

type chainMetricsObj struct {
	metrics *Metrics
	chainID *iscp.ChainID
}

var (
	_ ChainMetrics = &chainMetricsObj{}
	_ ChainMetrics = &defaultChainMetrics{}
)

func (c *chainMetricsObj) CountOffLedgerRequestIn() {
	c.metrics.offLedgerRequestCounter.With(prometheus.Labels{"chain": c.chainID.String()}).Inc()
}

func (c *chainMetricsObj) CountOnLedgerRequestIn() {
	c.metrics.onLedgerRequestCounter.With(prometheus.Labels{"chain": c.chainID.String()}).Inc()
}

func (c *chainMetricsObj) CountRequestOut() {
	c.metrics.processedRequestCounter.With(prometheus.Labels{"chain": c.chainID.String()}).Inc()
}

func (c *chainMetricsObj) CountMessages() {
	c.metrics.messagesReceived.With(prometheus.Labels{"chain": c.chainID.String()}).Inc()
}

func (c *chainMetricsObj) CountRequestAckMessages() {
	c.metrics.requestAckMessages.With(prometheus.Labels{"chain": c.chainID.String()}).Inc()
}

func (c *chainMetricsObj) RecordRequestProcessingTime(reqID iscp.RequestID, elapse time.Duration) {
	c.metrics.requestProcessingTime.With(prometheus.Labels{"chain": c.chainID.String(), "request": reqID.String()}).Set(elapse.Seconds())
}

func (c *chainMetricsObj) RecordVMRunTime(elapse time.Duration) {
	c.metrics.vmRunTime.With(prometheus.Labels{"chain": c.chainID.String()}).Set(elapse.Seconds())
}

func (c *chainMetricsObj) CountVmRuns() {
	c.metrics.vmRunCounter.With(prometheus.Labels{"chain": c.chainID.String()}).Inc()
}

func (c *chainMetricsObj) CountBlocksPerChain() {
	c.metrics.blocksPerChain.With(prometheus.Labels{"chain": c.chainID.String()}).Inc()
}

func (c *chainMetricsObj) RecordBlockSize(block_index uint32, block_size float64) {
	c.metrics.blockSizes.With(prometheus.Labels{"chain": c.chainID.String(), "block_index": fmt.Sprintf("%d", block_index)}).Set(block_size)
}

type defaultChainMetrics struct{}

func DefaultChainMetrics() ChainMetrics {
	return &defaultChainMetrics{}
}

func (m *defaultChainMetrics) CountOffLedgerRequestIn() {}

func (m *defaultChainMetrics) CountOnLedgerRequestIn() {}

func (m *defaultChainMetrics) CountRequestOut() {}

func (m *defaultChainMetrics) CountMessages() {}

func (m *defaultChainMetrics) CountRequestAckMessages() {}

func (m *defaultChainMetrics) RecordRequestProcessingTime(_ iscp.RequestID, _ time.Duration) {}

func (m *defaultChainMetrics) RecordVMRunTime(_ time.Duration) {}

func (m *defaultChainMetrics) CountVmRuns() {}

func (m *defaultChainMetrics) CountBlocksPerChain() {}

func (m *defaultChainMetrics) RecordBlockSize(_ uint32, _ float64) {}
