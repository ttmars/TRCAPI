package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TRCNodeAPI struct {
	NodeAddr string
}

var DTRCNodeAPI = TRCNodeAPI{
	NodeAddr: "http://34.87.43.172:8090",
}

func (obj *TRCNodeAPI)GetBlockByLimitNext(startNum int64, endNum int64) (*GetBlockByLimitNextStruct,error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"startNum": %v, "endNum": %v}`, startNum, endNum))
	req, err := http.NewRequest("POST", obj.NodeAddr + "/wallet/getblockbylimitnext", data)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var v GetBlockByLimitNextStruct
	err = json.Unmarshal(bodyText, &v)
	if err != nil {
		return nil,err
	}
	return &v,nil
}

func (obj *TRCNodeAPI)GetBlockByLatestNum(num int64) (*GetBlockByLatestNumStruct,error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"num" : %v}`, num))
	req, err := http.NewRequest("POST", obj.NodeAddr + "/wallet/getblockbylatestnum", data)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var v GetBlockByLatestNumStruct
	err = json.Unmarshal(bodyText, &v)
	if err != nil {
		return nil,err
	}
	return &v,nil
}

func (obj *TRCNodeAPI)GetNowBlock() (*GetNowBlockStruct,error)  {
	client := &http.Client{}
	req, err := http.NewRequest("POST", obj.NodeAddr + "/wallet/getnowblock", nil)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var v GetNowBlockStruct
	err = json.Unmarshal(bodyText, &v)
	if err != nil {
		return nil,err
	}
	return &v,nil
}

// GetTransactionInfoByBlockNum 查询区块上的交易列表
func (obj *TRCNodeAPI)GetTransactionInfoByBlockNum(num int64) (*GetTransactionInfoByBlockNumStruct,error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"num" : %v}`, num))
	req, err := http.NewRequest("POST", obj.NodeAddr + "/wallet/gettransactioninfobyblocknum", data)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var v GetTransactionInfoByBlockNumStruct
	err = json.Unmarshal(bodyText, &v)
	if err != nil {
		return nil,err
	}
	return &v,nil
}

// GetTransactionCountByBlockNum 查询区块上的交易数量
func (obj *TRCNodeAPI)GetTransactionCountByBlockNum(num int64) (*GetTransactionCountByBlockNumStruct,error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"num" : %v}`, num))
	req, err := http.NewRequest("POST", obj.NodeAddr + "/wallet/gettransactioncountbyblocknum", data)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var v GetTransactionCountByBlockNumStruct
	err = json.Unmarshal(bodyText, &v)
	if err != nil {
		return nil,err
	}
	return &v,nil
}


// GetTransactionInfoById 查询交易信息
func (obj *TRCNodeAPI)GetTransactionInfoById(id string) (*GetTransactionInfoByIdStruct,error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"value" : "%v"}`, id))
	req, err := http.NewRequest("POST", obj.NodeAddr + "/wallet/gettransactioninfobyid", data)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var v GetTransactionInfoByIdStruct
	err = json.Unmarshal(bodyText, &v)
	if err != nil {
		return nil,err
	}
	return &v,nil
}

// GetTransactionById 查询交易信息
func (obj *TRCNodeAPI)GetTransactionById(id string) (*GetTransactionByIdStruct,error) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"value" : "%v"}`, id))
	req, err := http.NewRequest("POST", obj.NodeAddr + "/wallet/gettransactionbyid", data)
	if err != nil {
		return nil,err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	var v GetTransactionByIdStruct
	err = json.Unmarshal(bodyText, &v)
	if err != nil {
		return nil,err
	}
	return &v,nil
}

