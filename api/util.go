package api

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/btcsuite/btcutil/base58"
	"strconv"
	"strings"
)

type TransInfo struct {
	BlockNumber		int				// 区块number
	TxID			string			// 交易hash
	Amount			int				// 交易数量
	OwnerAddress	string			// OwnerAddress
	ToAddress		string			// ToAddress
	Timestamp		int64		// 交易时间
}

// GetTransInfo 获取交易记录
// [startNum,endNum)，左闭右开，包含startNum而不包含endNum
// 一次最多查询10个区块，一个区块大概包含1000条交易记录
func (obj *TRCNodeAPI)GetTransInfo(startNum int64, endNum int64) ([]TransInfo, error) {
	if endNum <= startNum || endNum-startNum > 10{
		return nil,errors.New("参数错误！")
	}
	r,err := obj.GetBlockByLimitNext(startNum, endNum)
	if err != nil {
		return nil,err
	}
	var data []TransInfo
	for _,v := range r.Block {
		for _,vv := range v.Transactions {
			for _,vvv := range vv.RawData.Contract {
				BlockNumber := v.BlockHeader.RawData.Number
				TxID := vv.TxID
				Amount := vvv.Parameter.Value.Amount
				OwnerAddress := HexStringTobase58check(vvv.Parameter.Value.OwnerAddress)
				ToAddress := HexStringTobase58check(vvv.Parameter.Value.ToAddress)
				//Timestamp := time.UnixMilli(vv.RawData.Timestamp)
				Timestamp := vv.RawData.Timestamp
				if Amount != 0 {
					data = append(data, TransInfo{
						BlockNumber:  BlockNumber,
						TxID:         TxID,
						Amount:       Amount,
						OwnerAddress: OwnerAddress,
						ToAddress:    ToAddress,
						Timestamp:    Timestamp,
					})
					//fmt.Println(BlockNumber, TxID, Amount, OwnerAddress,ToAddress, Timestamp)
				}
			}
		}
	}
	return data,nil
}

// GetNowBlockNum 获取最新区块号
func (obj *TRCNodeAPI)GetNowBlockNum() (int,error)  {
	r,err := obj.GetNowBlock()
	if err != nil {
		return 0,err
	}
	return r.BlockHeader.RawData.Number,nil
}

// GetValueByID 根据hash获取交易金额
func (obj *TRCNodeAPI)GetValueByID(id string) (int64,error) {
	r,err := obj.GetTransactionInfoById(id)
	if err != nil {
		return 0,err
	}
	intValue, err := strconv.ParseInt(r.Log[0].Data, 16, 64)
	if err != nil {
		return 0,err
	}
	return intValue,nil
}

func HexStringTobase58check(data string) string {
	if len(data) == 0 {
		return ""
	}
	if data[0:2] == "0x" {
		data = data[2:]
	}
	if len(data)%2 == 1 {
		data = "0" + data
	}

	inputBytes, _ := hex.DecodeString(data)
	hash0 := sha256.Sum256(inputBytes)
	hash1 := sha256.Sum256(hash0[:])
	inputCheck := append(inputBytes, hash1[:4]...)
	base58Check := base58.Encode(inputCheck)

	return base58Check
}

func ParseData(data string) (MethodID string,ToAddr string, Quant int64, err error) {
	//data := "a9059cbb00000000000000000000000062287f55a6b2c8c0f7bc922a55c5c6edd4df83d2000000000000000000000000000000000000000000000000000000047c77bb28"
	if len(data) != 136 {
		return "","",0,errors.New("参数长度错误！")
	}

	if data[30:32] == "00" {
		ToAddr = HexStringTobase58check("41" + data[32:72])
	}else{
		ToAddr = HexStringTobase58check(data[30:72])
	}

	MethodID =  strings.TrimRight(data[:30], "0")
	Quant, err = strconv.ParseInt(strings.TrimLeft(data[len(data)-64:len(data)], "0"), 16, 64)
	if err != nil {
		return "","",0,err
	}
	return
}
