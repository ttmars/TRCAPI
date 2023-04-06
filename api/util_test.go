package api

import (
	"fmt"
	"testing"
)

func TestTRCNodeAPI_GetNowBlockNum(t *testing.T) {
	r,err := DTRCNodeAPI.GetNowBlockNum()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("最新区块号：%v\n", r)
}

func TestTRCNodeAPI_GetValueByID(t *testing.T) {
	hash := "55bd9353eb66a72878612a381ba1ed7e48a1a7c5ca0c2112af53ecbee85bdc12"
	r,err := DTRCNodeAPI.GetValueByID(hash)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("交易金额：", r)
}

func TestTRCNodeAPI_GetTransInfo(t *testing.T) {
	r,err := DTRCNodeAPI.GetTransInfo(49958100, 49958101)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("交易数量：", len(r))
}

func TestHexStringToBase58check(t *testing.T) {
	//HexString := "41E552F6487585C2B58BC2C9BB4492BC1F17132CD0"
	HexString := "41e552f6487585c2b58bc2c9bb4492bc1f17132cd0"
	Base58check := HexStringToBase58check(HexString)
	HS := Base58checkToHexString(Base58check)
	if HexString != HS {
		t.Error("test fail!")
	}
	fmt.Println(HS)
}
