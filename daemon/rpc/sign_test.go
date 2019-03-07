package rpc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cybercongress/cyberd/app"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"log"
	"testing"
)

/*
7b226d736773223a5b7b22696e70757473223a5b7b2261646472657373223a22637962657231387661307872767935776a76343667386133776333663434346c6b7932763936636836637374222c22636f696e73223a5b7b2264656e6f6d223a22637962222c22616d6f756e74223a223130303030227d5d7d5d2c226f757470757473223a5b7b2261646472657373223a22637962657231746d303879326b7964766575687a75776d3065636a30706b6c3239656a743467307570786377222c22636f696e73223a5b7b2264656e6f6d223a22637962222c22616d6f756e74223a223130303030227d5d7d5d7d5d2c22666565223a7b22616d6f756e74223a5b7b2264656e6f6d223a22222c22616d6f756e74223a2230227d5d2c22676173223a3230303030307d2c227369676e617475726573223a5b7b227075625f6b6579223a5b322c32372c32342c302c3235352c39362c3134372c32312c36342c32392c3133322c3139322c3130382c3231392c35392c3133342c3230362c3230312c3132362c3232342c36332c3136302c32342c3233362c3137302c3132342c3136342c39352c34332c3138302c362c3234362c3235305d2c227369676e6174757265223a5b3136352c37362c3130392c36312c35332c3132392c3139302c3134372c35322c3232342c33342c3130362c3233352c3230382c3232342c33362c3139302c32352c3230342c33362c3232362c3132392c39372c3130392c33352c3133302c3231372c3232382c3134342c3130362c31302c3133342c31342c3138332c39352c3235322c3231392c3233352c32322c39322c33372c35332c332c38392c3131312c3137332c31322c3135382c3134362c37312c38322c3131332c3233362c3234312c3137302c3132312c3231372c32302c3233362c32332c3133312c33352c38302c32395d2c226163636f756e745f6e756d626572223a3133323332392c2273657175656e6365223a307d5d2c226d656d6f223a22227d
*/

/*
{
  "msgs": [
    {
      "inputs": [
        {
          "address": "cyber18va0xrvy5wjv46g8a3wc3f444lky2v96ch6cst",
          "coins": [
            {
              "denom": "cyb",
              "amount": "10000"
            }
          ]
        }
      ],
      "outputs": [
        {
          "address": "cyber1tm08y2kydveuhzuwm0ecj0pkl29ejt4g0upxcw",
          "coins": [
            {
              "denom": "cyb",
              "amount": "10000"
            }
          ]
        }
      ]
    }
  ],
  "fee": {
    "amount": [
      {
        "denom": "",
        "amount": "0"
      }
    ],
    "gas": 200000
  },
  "signatures": [
    {
      "pub_key": [2,27,24,0,255,96,147,21,64,29,132,192,108,219,59,134,206,201,126,224,63,160,24,236,170,124,164,95,43,180,6,246,250],
      "signature": [165,76,109,61,53,129,190,147,52,224,34,106,235,208,224,36,190,25,204,36,226,129,97,109,35,130,217,228,144,106,10,134,14,183,95,252,219,235,22,92,37,53,3,89,111,173,12,158,146,71,82,113,236,241,170,121,217,20,236,23,131,35,80,29],
      "account_number": 132329,
      "sequence": 0
    }
  ],
  "memo": ""
}
*/
func TestStdTxMarshaling(t *testing.T) {
	app.SetPrefix()

	pubKeyRaw := [33]byte{2, 27, 24, 0, 255, 96, 147, 21, 64, 29, 132, 192, 108, 219, 59, 134, 206, 201, 126, 224, 63, 160, 24, 236, 170, 124, 164, 95, 43, 180, 6, 246, 250}
	pubKey := secp256k1.PubKeySecp256k1(pubKeyRaw)

	signatureRaw := []byte{165, 76, 109, 61, 53, 129, 190, 147, 52, 224, 34, 106, 235, 208, 224, 36, 190, 25, 204, 36, 226, 129, 97, 109, 35, 130, 217, 228, 144, 106, 10, 134, 14, 183, 95, 252, 219, 235, 22, 92, 37, 53, 3, 89, 111, 173, 12, 158, 146, 71, 82, 113, 236, 241, 170, 121, 217, 20, 236, 23, 131, 35, 80, 29}
	signatures := []auth.StdSignature{{PubKey: pubKey, Signature: signatureRaw}}

	fee := auth.StdFee{Amount: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(0)}}, Gas: 200000}

	fromAddress, _ := sdk.AccAddressFromBech32("cyber18va0xrvy5wjv46g8a3wc3f444lky2v96ch6cst")
	toAddress, _ := sdk.AccAddressFromBech32("cyber1tm08y2kydveuhzuwm0ecj0pkl29ejt4g0upxcw")

	inputs := []bank.Input{{
		Address: fromAddress,
		Coins:   sdk.Coins{sdk.NewInt64Coin("cyb", 10000)},
	}}

	outputs := []bank.Output{{
		Address: toAddress,
		Coins:   sdk.Coins{sdk.NewInt64Coin("cyb", 10000)},
	}}

	sendMsg := bank.MsgMultiSend{
		Inputs:  inputs,
		Outputs: outputs,
	}

	stdTx := auth.StdTx{Msgs: []sdk.Msg{sendMsg}, Fee: fee, Signatures: signatures, Memo: ""}
	stdTxBytes, err := codec.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		panic(err)
	}

	log.Println(stdTxBytes)
	return
}
