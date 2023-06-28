package token

import (
	"context"
	"errors"
	"sort"

	crud "github.com/web3eye-io/Web3Eye/nft-meta/pkg/crud/v1/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/imageconvert"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/milvusdb"
	val "github.com/web3eye-io/Web3Eye/proto/web3eye"
	nftmetanpool "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/token"
	rankernpool "github.com/web3eye-io/Web3Eye/proto/web3eye/ranker/v1/token"
	converter "github.com/web3eye-io/Web3Eye/ranker/pkg/converter/v1/token"
)

const (
	TopN           = 100
	PageLimit      = 10
	ShowSiblinsNum = 10
)

func (s *Server) Search(ctx context.Context, in *rankernpool.SearchTokenRequest) (*rankernpool.SearchTokenResponse, error) {
	// search from milvus
	scores, err := SerachFromMilvus(ctx, in.Vector)
	if err != nil {
		return nil, err
	}

	infos, err := QueryAndCollectTokens(ctx, scores)
	if err != nil {
		return nil, err
	}

	return &rankernpool.SearchTokenResponse{Infos: infos, Total: int32(len(infos)), Vector: in.Vector}, nil
}

func SerachFromMilvus(ctx context.Context, vector []float32) (map[int64]float32, error) {
	// search from milvus
	milvusmgr := milvusdb.NewNFTConllectionMGR()
	_vector := imageconvert.ToArrayVector(vector)

	_scores, err := milvusmgr.Search(ctx, [][milvusdb.VectorDim]float32{_vector}, TopN)
	if err != nil {
		return nil, err
	}

	if len(_scores) == 0 || len(_scores[0]) == 0 {
		return nil, errors.New("have no result")
	}

	return _scores[0], nil
}

func QueryAndCollectTokens(ctx context.Context, scores map[int64]float32) ([]*rankernpool.SearchToken, error) {
	vIDs := []int64{}
	for i := range scores {
		vIDs = append(vIDs, i)
	}

	conds := &nftmetanpool.Conds{
		VectorIDs: &val.Int64SliceVal{
			Op:    "in",
			Value: vIDs,
		},
	}

	// query from db
	rows, _, err := crud.Rows(ctx, conds, 0, len(vIDs))
	if err != nil {
		return nil, err
	}

	infos := []*rankernpool.SearchToken{}
	for i := 0; i < len(rows); i++ {
		info := converter.Ent2Grpc(rows[i])
		info.Distance = scores[info.VectorID]
		info.SiblingTokens = make([]*rankernpool.SiblingToken, 0)
		infos = append(infos, info)
	}

	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Distance < infos[j].Distance
	})

	// collection
	contractRecord := make(map[string]int)
	result := []*rankernpool.SearchToken{}
	for _, v := range infos {
		if _, ok := contractRecord[v.Contract]; ok {
			if len(result[contractRecord[v.Contract]].SiblingTokens) >= ShowSiblinsNum {
				continue
			}
			result[contractRecord[v.Contract]].SiblingTokens = append(
				result[contractRecord[v.Contract]].SiblingTokens, &rankernpool.SiblingToken{
					TokenID:      v.TokenID,
					ImageURL:     v.ImageURL,
					IPFSImageURL: v.IPFSImageURL,
				})
		} else {
			result = append(result, v)
			contractRecord[v.Contract] = len(result) - 1
		}
	}

	// count transfers and
	for _, v := range result {
		conds = &nftmetanpool.Conds{
			ChainType: &val.StringVal{
				Op:    "eq",
				Value: v.ChainType,
			},
			ChainID: &val.StringVal{
				Op:    "eq",
				Value: v.ChainID,
			},
			Contract: &val.StringVal{
				Op:    "eq",
				Value: v.Contract,
			},
		}

		limit := ShowSiblinsNum - len(v.SiblingTokens)
		if limit <= 0 {
			continue
		}

		tokens, num, err := crud.Rows(ctx, conds, 0, limit)
		if err != nil {
			return nil, err
		}
		v.SiblingsNum = uint32(num)

		for _, token := range tokens {
			v.SiblingTokens = append(v.SiblingTokens, &rankernpool.SiblingToken{
				TokenID:      token.TokenID,
				ImageURL:     token.ImageURL,
				IPFSImageURL: token.IpfsImageURL,
			})
		}

		v.SiblingTokens = SliceDeduplicate(v.SiblingTokens)
	}
	return result, nil
}

func SliceDeduplicate(s []*rankernpool.SiblingToken) []*rankernpool.SiblingToken {
	mapRecord := make(map[string]struct{})
	listRecord := []int{}
	for i, v := range s {
		if _, ok := mapRecord[v.TokenID]; ok {
			listRecord = append(listRecord, i)
		} else {
			mapRecord[v.TokenID] = struct{}{}
		}
	}
	recordLen := len(listRecord) - 1
	for i := range listRecord {
		s = append(s[0:listRecord[recordLen-i]], s[listRecord[recordLen-i]+1:]...)
	}
	return s
}
