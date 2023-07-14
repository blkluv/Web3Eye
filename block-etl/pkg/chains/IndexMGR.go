package chains

import (
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/web3eye-io/Web3Eye/block-etl/pkg/chains/eth"
	endpointNMCli "github.com/web3eye-io/Web3Eye/nft-meta/pkg/client/v1/endpoint"

	common_eth "github.com/web3eye-io/Web3Eye/common/chains/eth"
	"github.com/web3eye-io/Web3Eye/proto/web3eye"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/cttype"
	"github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/endpoint"
	"golang.org/x/net/context"
)

type GetEndpointChainID func(ctx context.Context, endpoint string) (ChainID string, err error)

type Index interface {
	StartIndex(ctx context.Context)
	StopIndex()
}

type indexMGR struct {
	UpdateInterval          time.Duration
	EndpointChainIDHandlers map[basetype.ChainType]GetEndpointChainID
	Indexs                  map[basetype.ChainType]map[string]Index
}

type EndpointInfo struct {
	ChainType basetype.ChainType
	Endpoint  string
	chainID   string
}

var (
	pMGR           *indexMGR
	UpdateInterval = time.Second * 10
)

func init() {
	pMGR = &indexMGR{
		UpdateInterval:          UpdateInterval,
		EndpointChainIDHandlers: make(map[basetype.ChainType]GetEndpointChainID),
	}

	pMGR.EndpointChainIDHandlers[basetype.ChainType_Ethereum] = common_eth.GetEndpointChainID
}

func GetIndexMGR() *indexMGR {
	return pMGR
}

func (pmgr *indexMGR) StartRunning(ctx context.Context) {
	// resp, err := endpointNMCli.GetEndpoints(ctx, &endpoint.GetEndpointsRequest{})
	// if err != nil {
	// 	logger.Sugar().Errorf("get endpoints from nft-meta failed, err: %v", err)
	// }
	// for _, info := range eInfos {
	// 	if _, ok := pmgr.EndpointChainIDHandlers[info.ChainType]; !ok {
	// 		logger.Sugar().Errorf("have not support chain type: %v", info.ChainType.String())
	// 	}
	// 	chainID, err := pmgr.EndpointChainIDHandlers[info.ChainType](ctx, info.Endpoint)
	// 	if err != nil {
	// 		logger.Sugar().Errorf("cannot get chainID for chainType: %v,endpoint: %v,err: %v", info.ChainType.String(), info.Endpoint, err)
	// 	}
	// 	info.chainID = chainID
	// }
}

func (pmgr *indexMGR) checkNewEndpoints(ctx context.Context) {
	conds := &endpoint.Conds{
		State: &web3eye.StringVal{
			Op:    "eq",
			Value: cttype.EndpointState_EndpointDefault.String(),
		},
	}

	getEResp, err := endpointNMCli.GetEndpoints(ctx, &endpoint.GetEndpointsRequest{Conds: conds})
	if err != nil {
		logger.Sugar().Errorf("get endpoints from nft-meta failed, err: %v", err)
		return
	}

	infos := getEResp.GetInfos()
	updateInfos := []*endpoint.EndpointReq{}
	for _, info := range infos {
		handler, ok := pmgr.EndpointChainIDHandlers[info.ChainType]
		if !ok {
			logger.Sugar().Warnf("have not handler for chain type: %v", info.ChainType)
		}

		chainID, err := handler(ctx, info.Address)
		if err != nil {
			info.State = cttype.EndpointState_EndpointError
			continue
		}
		info.ChainID = chainID
		info.State = cttype.EndpointState_EndpointAvaliable

		updateInfos = append(updateInfos, &endpoint.EndpointReq{
			ID:        &info.ID,
			ChainType: &info.ChainType,
			ChainID:   &info.ChainID,
			Address:   &info.Address,
			State:     &info.State,
		})
	}

	updateEResp, err := endpointNMCli.UpdateEndpoints(ctx, &endpoint.UpdateEndpointsRequest{
		Infos: updateInfos,
	})
	if err != nil {
		logger.Sugar().Errorf("update endpoints to nft-meta failed, err: %v", err)
		return
	}
	if len(updateInfos) != 0 {
		for _, v := range updateEResp.Infos {
			logger.Sugar().Errorf("update endpoint %v failed, err: %v", v.ID, v.MSG)
		}
	}
}

func (pmgr *indexMGR) AddEndpoint(info EndpointInfo) {
	if _, ok := pmgr.Indexs[info.ChainType]; !ok {
		pmgr.Indexs[info.ChainType] = make(map[string]Index)
	}

	if _, ok := pmgr.Indexs[info.ChainType][info.chainID]; !ok {
		pmgr.Indexs[info.ChainType][info.chainID] = &eth.EthIndexer{}
	}
}
