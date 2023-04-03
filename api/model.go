package api

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
