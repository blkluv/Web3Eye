// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: web3eye/ranker/v1/contract/contract.proto

package contract

import (
	contract "github.com/web3eye-io/Web3Eye/proto/web3eye/nftmeta/v1/contract"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_web3eye_ranker_v1_contract_contract_proto protoreflect.FileDescriptor

var file_web3eye_ranker_v1_contract_contract_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x72, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a,
	0x2a, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb1, 0x03, 0x0a, 0x07,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x62, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x27, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x2b,
	0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6e, 0x66,
	0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4f, 0x6e, 0x6c,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x12, 0x28, 0x2e, 0x6e, 0x66,
	0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x6e, 0x66, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65,
	0x62, 0x33, 0x65, 0x79, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x57, 0x65, 0x62, 0x33, 0x45, 0x79, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x65, 0x79, 0x65, 0x2f, 0x72,
	0x61, 0x6e, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_web3eye_ranker_v1_contract_contract_proto_goTypes = []interface{}{
	(*contract.GetContractRequest)(nil),      // 0: nftmeta.v1.contract.GetContractRequest
	(*contract.GetContractOnlyRequest)(nil),  // 1: nftmeta.v1.contract.GetContractOnlyRequest
	(*contract.GetContractsRequest)(nil),     // 2: nftmeta.v1.contract.GetContractsRequest
	(*contract.CountContractsRequest)(nil),   // 3: nftmeta.v1.contract.CountContractsRequest
	(*contract.GetContractResponse)(nil),     // 4: nftmeta.v1.contract.GetContractResponse
	(*contract.GetContractOnlyResponse)(nil), // 5: nftmeta.v1.contract.GetContractOnlyResponse
	(*contract.GetContractsResponse)(nil),    // 6: nftmeta.v1.contract.GetContractsResponse
	(*contract.CountContractsResponse)(nil),  // 7: nftmeta.v1.contract.CountContractsResponse
}
var file_web3eye_ranker_v1_contract_contract_proto_depIdxs = []int32{
	0, // 0: ranker.v1.contract.Manager.GetContract:input_type -> nftmeta.v1.contract.GetContractRequest
	1, // 1: ranker.v1.contract.Manager.GetContractOnly:input_type -> nftmeta.v1.contract.GetContractOnlyRequest
	2, // 2: ranker.v1.contract.Manager.GetContracts:input_type -> nftmeta.v1.contract.GetContractsRequest
	3, // 3: ranker.v1.contract.Manager.CountContracts:input_type -> nftmeta.v1.contract.CountContractsRequest
	4, // 4: ranker.v1.contract.Manager.GetContract:output_type -> nftmeta.v1.contract.GetContractResponse
	5, // 5: ranker.v1.contract.Manager.GetContractOnly:output_type -> nftmeta.v1.contract.GetContractOnlyResponse
	6, // 6: ranker.v1.contract.Manager.GetContracts:output_type -> nftmeta.v1.contract.GetContractsResponse
	7, // 7: ranker.v1.contract.Manager.CountContracts:output_type -> nftmeta.v1.contract.CountContractsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_web3eye_ranker_v1_contract_contract_proto_init() }
func file_web3eye_ranker_v1_contract_contract_proto_init() {
	if File_web3eye_ranker_v1_contract_contract_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_web3eye_ranker_v1_contract_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web3eye_ranker_v1_contract_contract_proto_goTypes,
		DependencyIndexes: file_web3eye_ranker_v1_contract_contract_proto_depIdxs,
	}.Build()
	File_web3eye_ranker_v1_contract_contract_proto = out.File
	file_web3eye_ranker_v1_contract_contract_proto_rawDesc = nil
	file_web3eye_ranker_v1_contract_contract_proto_goTypes = nil
	file_web3eye_ranker_v1_contract_contract_proto_depIdxs = nil
}