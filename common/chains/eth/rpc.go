package eth

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/web3eye-io/Web3Eye/common/chains/eth/contracts"
	basetype "github.com/web3eye-io/Web3Eye/proto/web3eye/basetype/v1"
)

func (ethCli ethClients) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	logs := []types.Log{}

	var err error
	err = ethCli.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		logs, err = c.FilterLogs(ctx, query)
		if err != nil {
			return false, err
		}
		return false, nil
	})

	return logs, err
}

func (ethCli ethClients) CurrentBlockNum(ctx context.Context) (uint64, error) {
	var num uint64

	var err error
	err = ethCli.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		num, err = c.BlockNumber(ctx)
		if err != nil {
			return false, err
		}
		return false, nil
	})

	return num, err
}

func (e ethClients) TokenURI(ctx context.Context, tokenType basetype.TokenType, contractAddr, tokenID string, blockNum uint64) (string, error) {

	var uri string
	var err error
	err = e.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		uri, err = e.tokenURI(c, tokenType, contractAddr, tokenID, blockNum)
		return false, err
	})
	uri = e.ReplaceID(uri, tokenID)
	return uri, err
}

func (ethCli ethClients) tokenURI(
	ethClient *ethclient.Client,
	tokenType basetype.TokenType, contractAddr,
	tokenID string,
	blockNum uint64) (string, error) {
	if !common.IsHexAddress(contractAddr) {
		return "", errors.New("contract address is invalid")
	}

	baseNum := 10
	contract := common.HexToAddress(contractAddr)
	token, ok := big.NewInt(0).SetString(tokenID, baseNum)
	if !ok {
		return "", errors.New("parse tokenID failed")
	}

	switch tokenType {
	case basetype.TokenType_ERC721:
		erc721, err := contracts.NewIERC721MetadataCaller(contract, ethClient)
		if err != nil {
			return "", err
		}
		return erc721.TokenURI(&bind.CallOpts{BlockNumber: big.NewInt(int64(blockNum))}, token)
	case basetype.TokenType_ERC1155:
		erc1155, err := contracts.NewIERC1155MetadataURICaller(contract, ethClient)
		if err != nil {
			return "", err
		}
		return erc1155.Uri(&bind.CallOpts{BlockNumber: big.NewInt(int64(blockNum))}, token)
	}

	return "", nil
}

func (ethCli ethClients) BlockByNumber(ctx context.Context, blockNum uint64) (*types.Block, error) {
	var block *types.Block
	var err error
	err = ethCli.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		block, err = c.BlockByNumber(ctx, big.NewInt(int64(blockNum)))
		return false, err
	})
	return block, err
}

type ContractCreator struct {
	From        common.Address
	BlockNumber uint64
	TxHash      common.Hash
	TxTime      uint64
}

func (ethCli ethClients) GetContractCreator(ctx context.Context, contractAddr string) (*ContractCreator, error) {

	var creator *ContractCreator
	var err error
	err = ethCli.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		creator, err = ethCli.getContractCreator(ctx, c, contractAddr)
		return false, err
	})
	return creator, err
}

func (ethCli ethClients) getContractCreator(ctx context.Context, ethClient *ethclient.Client, contractAddr string) (*ContractCreator, error) {
	rHeight, err := ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	contract := common.HexToAddress(contractAddr)
	var lHeight uint64 = 0
	for lHeight <= rHeight {
		mHeight := (rHeight + lHeight) / 2
		codeAt, err := ethClient.CodeAt(ctx, contract, big.NewInt(0).SetUint64(mHeight))
		if err != nil {
			return nil, err
		}
		if len(codeAt) > 0 {
			rHeight = mHeight
		} else {
			lHeight = mHeight
		}
		if lHeight+1 == rHeight {
			break
		}
	}
	block, err := ethClient.BlockByNumber(ctx, big.NewInt(0).SetUint64(rHeight))
	if err != nil {
		return nil, err
	}

	txs := block.Transactions()
	for _, tx := range txs {
		if tx.To() != nil || len(tx.Data()) == 0 {
			continue
		}
		receipt, err := ethClient.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, err
		}

		if receipt.ContractAddress == contract {
			from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
			if err != nil {
				return nil, err
			}

			return &ContractCreator{
				From:        from,
				BlockNumber: receipt.BlockNumber.Uint64(),
				TxHash:      receipt.TxHash,
				TxTime:      block.Time(),
			}, nil
		}
	}
	return nil, errors.New("cannot get contract creator")
}

type ERC721Metadata struct {
	Name   string
	Symbol string
}

func (ethCli ethClients) GetERC721Metadata(ctx context.Context, contractAddr string) (*ERC721Metadata, error) {
	var info *ERC721Metadata
	var err error
	err = ethCli.WithClient(ctx, func(ctx context.Context, c *ethclient.Client) (bool, error) {
		info, err = ethCli.getERC721Metadata(ctx, c, contractAddr)
		return false, err
	})
	return info, err
}

func (ethCli ethClients) getERC721Metadata(ctx context.Context, ethClient *ethclient.Client, contractAddr string) (*ERC721Metadata, error) {
	contract := common.HexToAddress(contractAddr)
	instance, err := contracts.NewIERC721MetadataCaller(contract, ethClient)
	if err != nil {
		return nil, err
	}

	name, nameErr := instance.Name(&bind.CallOpts{
		Context: ctx,
	})
	if nameErr != nil {
		err = fmt.Errorf("name cannot be gained %v", nameErr)
	}
	symbol, symbolErr := instance.Symbol(&bind.CallOpts{
		Context: ctx,
	})
	if symbolErr != nil {
		err = fmt.Errorf("%v, %v", err, symbolErr)
	}

	return &ERC721Metadata{Name: name, Symbol: symbol}, err
}

// ReplaceID replaces the token's ID with the given ID
func (ethCli ethClients) ReplaceID(tokenURI, id string) string {
	_id := fmt.Sprintf("%064s", id)
	return strings.TrimSpace(strings.ReplaceAll(tokenURI, "{id}", _id))
}
