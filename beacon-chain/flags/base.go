// Package flags defines beacon-node specific runtime flags for
// setting important values such as ports, eth1 endpoints, and more.
package flags

import (
	"github.com/urfave/cli/v2"
)

var (
	// HTTPWeb3ProviderFlag provides an HTTP access endpoint to an ETH 1.0 RPC.
	HTTPWeb3ProviderFlag = &cli.StringFlag{
		Name:  "http-web3provider",
		Usage: "A mainchain web3 provider string http endpoint",
		Value: "https://goerli.prylabs.net",
	}
	// DepositContractFlag defines a flag for the deposit contract address.
	DepositContractFlag = &cli.StringFlag{
		Name:  "deposit-contract",
		Usage: "Deposit contract address. Beacon chain node will listen logs coming from the deposit contract to determine when validator is eligible to participate.",
		Value: "0x5cA1e00004366Ac85f492887AAab12d0e6418876",
	}
	// RPCHost defines the host on which the RPC server should listen.
	RPCHost = &cli.StringFlag{
		Name:  "rpc-host",
		Usage: "Host on which the RPC server should listen",
		Value: "0.0.0.0",
	}
	// RPCPort defines a beacon node RPC port to open.
	RPCPort = &cli.IntFlag{
		Name:  "rpc-port",
		Usage: "RPC port exposed by a beacon node",
		Value: 4000,
	}
	// RPCMaxPageSize defines the maximum numbers per page returned in RPC responses from this
	// beacon node (default: 500).
	RPCMaxPageSize = &cli.IntFlag{
		Name:  "rpc-max-page-size",
		Usage: "Max number of items returned per page in RPC responses for paginated endpoints.",
		Value: 500,
	}
	// MonitoringPortFlag defines the http port used to serve prometheus metrics.
	MonitoringPortFlag = &cli.Int64Flag{
		Name:  "monitoring-port",
		Usage: "Port used to listening and respond metrics for prometheus.",
		Value: 8080,
	}
	// CertFlag defines a flag for the node's TLS certificate.
	CertFlag = &cli.StringFlag{
		Name:  "tls-cert",
		Usage: "Certificate for secure gRPC. Pass this and the tls-key flag in order to use gRPC securely.",
	}
	// KeyFlag defines a flag for the node's TLS key.
	KeyFlag = &cli.StringFlag{
		Name:  "tls-key",
		Usage: "Key for secure gRPC. Pass this and the tls-cert flag in order to use gRPC securely.",
	}
	// DisableGRPCGateway for JSON-HTTP requests to the beacon node.
	DisableGRPCGateway = &cli.BoolFlag{
		Name:  "disable-grpc-gateway",
		Usage: "Disable the gRPC gateway for JSON-HTTP requests",
	}
	// GRPCGatewayPort enables a gRPC gateway to be exposed for Prysm.
	GRPCGatewayPort = &cli.IntFlag{
		Name:  "grpc-gateway-port",
		Usage: "Enable gRPC gateway for JSON requests",
		Value: 3000,
	}
	// GPRCGatewayCorsDomain serves preflight requests when serving gRPC JSON gateway.
	GPRCGatewayCorsDomain = &cli.StringFlag{
		Name: "grpc-gateway-corsdomain",
		Usage: "Comma separated list of domains from which to accept cross origin requests " +
			"(browser enforced). This flag has no effect if not used with --grpc-gateway-port.",
	}
	// MinSyncPeers specifies the required number of successful peer handshakes in order
	// to start syncing with external peers.
	MinSyncPeers = &cli.IntFlag{
		Name:  "min-sync-peers",
		Usage: "The required number of valid peers to connect with before syncing.",
		Value: 3,
	}
	// ContractDeploymentBlock is the block in which the eth1 deposit contract was deployed.
	ContractDeploymentBlock = &cli.IntFlag{
		Name:  "contract-deployment-block",
		Usage: "The eth1 block in which the deposit contract was deployed.",
		Value: 2523557,
	}
	// SetGCPercent is the percentage of current live allocations at which the garbage collector is to run.
	SetGCPercent = &cli.IntFlag{
		Name:  "gc-percent",
		Usage: "The percentage of freshly allocated data to live data on which the gc will be run again.",
		Value: 100,
	}
	// UnsafeSync starts the beacon node from the previously saved head state and syncs from there.
	UnsafeSync = &cli.BoolFlag{
		Name:  "unsafe-sync",
		Usage: "Starts the beacon node with the previously saved head state instead of finalized state.",
	}
	// SlasherCertFlag defines a flag for the slasher TLS certificate.
	SlasherCertFlag = &cli.StringFlag{
		Name:  "slasher-tls-cert",
		Usage: "Certificate for secure slasher gRPC connection. Pass this in order to use slasher gRPC securely.",
	}
	// SlasherProviderFlag defines a flag for a slasher RPC provider.
	SlasherProviderFlag = &cli.StringFlag{
		Name:  "slasher-provider",
		Usage: "A slasher provider string endpoint. Can either be an grpc server endpoint.",
		Value: "127.0.0.1:4002",
	}
	// SlotsPerArchivedPoint specifies the number of slots between the archived points, to save beacon state in the cold
	// section of DB.
	SlotsPerArchivedPoint = &cli.IntFlag{
		Name:  "slots-per-archive-point",
		Usage: "The slot durations of when an archived state gets saved in the DB.",
		Value: 2048,
	}
	// DisableDiscv5 disables running discv5.
	DisableDiscv5 = &cli.BoolFlag{
		Name:  "disable-discv5",
		Usage: "Does not run the discoveryV5 dht.",
	}
	// BlockBatchLimit specifies the requested block batch size.
	BlockBatchLimit = &cli.IntFlag{
		Name:  "block-batch-limit",
		Usage: "The amount of blocks the local peer is bounded to request and respond to in a batch.",
		Value: 64,
	}
	// BlockBatchLimitBurstFactor specifies the factor by which block batch size may increase.
	BlockBatchLimitBurstFactor = &cli.IntFlag{
		Name:  "block-batch-limit-burst-factor",
		Usage: "The factor by which block batch limit may increase on burst.",
		Value: 10,
	}
	// EnableDebugRPCEndpoints as /v1/beacon/state.
	EnableDebugRPCEndpoints = &cli.BoolFlag{
		Name:  "enable-debug-rpc-endpoints",
		Usage: "Enables the debug rpc service, containing utility endpoints such as /eth/v1alpha1/beacon/state. Requires --new-state-mgmt",
	}
)
