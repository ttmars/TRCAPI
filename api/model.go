package api

type GetAssetIssueListStruct struct {
	AssetIssue []struct {
		OwnerAddress string `json:"owner_address"`
		Name         string `json:"name"`
		Abbr         string `json:"abbr,omitempty"`
		TotalSupply  int64  `json:"total_supply"`
		TrxNum       int    `json:"trx_num"`
		Num          int    `json:"num"`
		StartTime    int64  `json:"start_time"`
		EndTime      int64  `json:"end_time"`
		Description  string `json:"description"`
		URL          string `json:"url"`
		ID           string `json:"id"`
		FrozenSupply []struct {
			FrozenAmount int `json:"frozen_amount"`
			FrozenDays   int `json:"frozen_days"`
		} `json:"frozen_supply,omitempty"`
		PublicLatestFreeNetTime int `json:"public_latest_free_net_time,omitempty"`
		VoteScore               int `json:"vote_score,omitempty"`
		FreeAssetNetLimit       int `json:"free_asset_net_limit,omitempty"`
		PublicFreeAssetNetLimit int `json:"public_free_asset_net_limit,omitempty"`
		Precision               int `json:"precision,omitempty"`
	} `json:"assetIssue"`
}

type GetAccountResourceStruct struct {
	FreeNetUsed  int `json:"freeNetUsed"`
	FreeNetLimit int `json:"freeNetLimit"`
	NetUsed      int `json:"NetUsed"`
	NetLimit     int `json:"NetLimit"`
	AssetNetUsed []struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	} `json:"assetNetUsed"`
	AssetNetLimit []struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	} `json:"assetNetLimit"`
	TotalNetLimit     int64 `json:"TotalNetLimit"`
	TotalNetWeight    int64 `json:"TotalNetWeight"`
	EnergyUsed        int   `json:"EnergyUsed"`
	EnergyLimit       int   `json:"EnergyLimit"`
	TotalEnergyLimit  int64 `json:"TotalEnergyLimit"`
	TotalEnergyWeight int64 `json:"TotalEnergyWeight"`
}

type GetAccountStruct struct {
	Address               string `json:"address"`
	Balance               int64  `json:"balance"`
	NetUsage              int    `json:"net_usage"`
	CreateTime            int64  `json:"create_time"`
	LatestOprationTime    int64  `json:"latest_opration_time"`
	FreeNetUsage          int    `json:"free_net_usage"`
	LatestConsumeTime     int64  `json:"latest_consume_time"`
	LatestConsumeFreeTime int64  `json:"latest_consume_free_time"`
	AccountResource       struct {
		EnergyUsage                             int   `json:"energy_usage"`
		LatestConsumeTimeForEnergy              int64 `json:"latest_consume_time_for_energy"`
		AcquiredDelegatedFrozenBalanceForEnergy int64 `json:"acquired_delegated_frozen_balance_for_energy"`
	} `json:"account_resource"`
	OwnerPermission struct {
		PermissionName string `json:"permission_name"`
		Threshold      int    `json:"threshold"`
		Keys           []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
	} `json:"owner_permission"`
	ActivePermission []struct {
		Type           string `json:"type"`
		ID             int    `json:"id"`
		PermissionName string `json:"permission_name"`
		Threshold      int    `json:"threshold"`
		Operations     string `json:"operations"`
		Keys           []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
	} `json:"active_permission"`
	FrozenV2 []struct {
		Type string `json:"type,omitempty"`
	} `json:"frozenV2"`
	AcquiredDelegatedFrozenBalanceForBandwidth int64 `json:"acquired_delegated_frozen_balance_for_bandwidth"`
	AssetV2                                    []struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	} `json:"assetV2"`
	FreeAssetNetUsageV2 []struct {
		Key   string `json:"key"`
		Value int    `json:"value"`
	} `json:"free_asset_net_usageV2"`
	AssetOptimized bool `json:"asset_optimized"`
}

type GetBlockByLimitNextStruct struct {
	Block []struct {
		BlockID     string `json:"blockID"`
		BlockHeader struct {
			RawData struct {
				Number         int    `json:"number"`
				TxTrieRoot     string `json:"txTrieRoot"`
				WitnessAddress string `json:"witness_address"`
				ParentHash     string `json:"parentHash"`
				Version        int    `json:"version"`
				Timestamp      int64  `json:"timestamp"`
			} `json:"raw_data"`
			WitnessSignature string `json:"witness_signature"`
		} `json:"block_header"`
		Transactions []struct {
			Ret []struct {
				ContractRet string `json:"contractRet"`
			} `json:"ret"`
			Signature []string `json:"signature"`
			TxID      string   `json:"txID"`
			RawData   struct {
				Contract []struct {
					Parameter struct {
						Value struct {
							Amount       int    `json:"amount"`
							AssetName    string `json:"asset_name"`
							OwnerAddress string `json:"owner_address"`
							ToAddress    string `json:"to_address"`
						} `json:"value"`
						TypeURL string `json:"type_url"`
					} `json:"parameter"`
					Type string `json:"type"`
				} `json:"contract"`
				RefBlockBytes string `json:"ref_block_bytes"`
				RefBlockHash  string `json:"ref_block_hash"`
				Expiration    int64  `json:"expiration"`
				Timestamp     int64  `json:"timestamp"`
			} `json:"raw_data"`
			RawDataHex string `json:"raw_data_hex"`
		} `json:"transactions"`
	} `json:"block"`
}

type GetBlockByLatestNumStruct struct {
	Block []struct {
		BlockID     string `json:"blockID"`
		BlockHeader struct {
			RawData struct {
				Number         int    `json:"number"`
				TxTrieRoot     string `json:"txTrieRoot"`
				WitnessAddress string `json:"witness_address"`
				ParentHash     string `json:"parentHash"`
				Version        int    `json:"version"`
				Timestamp      int64  `json:"timestamp"`
			} `json:"raw_data"`
			WitnessSignature string `json:"witness_signature"`
		} `json:"block_header"`
		Transactions []struct {
			Ret []struct {
				ContractRet string `json:"contractRet"`
			} `json:"ret"`
			Signature []string `json:"signature"`
			TxID      string   `json:"txID"`
			RawData   struct {
				Contract []struct {
					Parameter struct {
						Value struct {
							Amount       int    `json:"amount"`
							OwnerAddress string `json:"owner_address"`
							ToAddress    string `json:"to_address"`
						} `json:"value"`
						TypeURL string `json:"type_url"`
					} `json:"parameter"`
					Type string `json:"type"`
				} `json:"contract"`
				RefBlockBytes string `json:"ref_block_bytes"`
				RefBlockHash  string `json:"ref_block_hash"`
				Expiration    int64  `json:"expiration"`
				Timestamp     int64  `json:"timestamp"`
			} `json:"raw_data"`
			RawDataHex string `json:"raw_data_hex"`
		} `json:"transactions"`
	} `json:"block"`
}

type GetNowBlockStruct struct {
	BlockID     string `json:"blockID"`
	BlockHeader struct {
		RawData struct {
			Number         int    `json:"number"`
			TxTrieRoot     string `json:"txTrieRoot"`
			WitnessAddress string `json:"witness_address"`
			ParentHash     string `json:"parentHash"`
			Version        int    `json:"version"`
			Timestamp      int64  `json:"timestamp"`
		} `json:"raw_data"`
		WitnessSignature string `json:"witness_signature"`
	} `json:"block_header"`
	Transactions []struct {
		Ret []struct {
			ContractRet string `json:"contractRet"`
		} `json:"ret"`
		Signature []string `json:"signature"`
		TxID      string   `json:"txID"`
		RawData   struct {
			Contract []struct {
				Parameter struct {
					Value struct {
						Data            string `json:"data"`
						OwnerAddress    string `json:"owner_address"`
						ContractAddress string `json:"contract_address"`
					} `json:"value"`
					TypeURL string `json:"type_url"`
				} `json:"parameter"`
				Type string `json:"type"`
			} `json:"contract"`
			RefBlockBytes string `json:"ref_block_bytes"`
			RefBlockHash  string `json:"ref_block_hash"`
			Expiration    int64  `json:"expiration"`
			FeeLimit      int    `json:"fee_limit"`
			Timestamp     int64  `json:"timestamp"`
		} `json:"raw_data"`
		RawDataHex string `json:"raw_data_hex"`
	} `json:"transactions"`
}


type GetTransactionInfoByBlockNumStruct []struct {
	Log []struct {
		Address string   `json:"address"`
		Data    string   `json:"data"`
		Topics  []string `json:"topics"`
	} `json:"log,omitempty"`
	BlockNumber    int      `json:"blockNumber"`
	ContractResult []string `json:"contractResult"`
	BlockTimeStamp int64    `json:"blockTimeStamp"`
	Receipt        struct {
		Result             string `json:"result"`
		EnergyPenaltyTotal int    `json:"energy_penalty_total"`
		EnergyUsage        int    `json:"energy_usage"`
		EnergyUsageTotal   int    `json:"energy_usage_total"`
		OriginEnergyUsage  int    `json:"origin_energy_usage"`
		NetUsage           int    `json:"net_usage"`
	} `json:"receipt"`
	ID              string `json:"id"`
	ContractAddress string `json:"contract_address,omitempty"`
	Fee             int    `json:"fee,omitempty"`
	Result          string `json:"result,omitempty"`
	ResMessage      string `json:"resMessage,omitempty"`
	UnfreezeAmount  int    `json:"unfreeze_amount,omitempty"`
}


type GetTransactionCountByBlockNumStruct struct {
	Count int `json:"count"`
}


type GetTransactionInfoByIdStruct struct {
	ID              string   `json:"id"`
	Fee             int      `json:"fee"`
	BlockNumber     int      `json:"blockNumber"`
	BlockTimeStamp  int64    `json:"blockTimeStamp"`
	ContractResult  []string `json:"contractResult"`
	ContractAddress string   `json:"contract_address"`
	Receipt         struct {
		EnergyFee          int    `json:"energy_fee"`
		EnergyUsageTotal   int    `json:"energy_usage_total"`
		NetUsage           int    `json:"net_usage"`
		Result             string `json:"result"`
		EnergyPenaltyTotal int    `json:"energy_penalty_total"`
	} `json:"receipt"`
	Log []struct {
		Address string   `json:"address"`
		Topics  []string `json:"topics"`
		Data    string   `json:"data"`								// 交易金额
	} `json:"log"`
}

type GetTransactionByIdStruct struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	} `json:"ret"`
	Signature []string `json:"signature"`
	TxID      string   `json:"txID"`
	RawData   struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Data            string `json:"data"`
					OwnerAddress    string `json:"owner_address"`
					ContractAddress string `json:"contract_address"`
				} `json:"value"`
				TypeURL string `json:"type_url"`
			} `json:"parameter"`
			Type string `json:"type"`
		} `json:"contract"`
		RefBlockBytes string `json:"ref_block_bytes"`
		RefBlockHash  string `json:"ref_block_hash"`
		Expiration    int64  `json:"expiration"`
		FeeLimit      int    `json:"fee_limit"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"raw_data"`
	RawDataHex string `json:"raw_data_hex"`
}
