package main

import (
	"Parse_EventLogSCC/code/contracts"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

type Event struct {
	_batchIndex        uint
	_batchRoot         byte
	_batchSize         uint
	_prevTotalElements uint
	_extraData         byte
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a39c40e31fb64ef8bdf7b0f6feccb277")
	if err != nil {
		log.Fatal(err)
	}

	//Assign contract address, and Block range
	contractAddress := common.HexToAddress("0x66b9f45E84A0aD7fE3983c97556798352a8E0a56")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(18609811),
		ToBlock:   big.NewInt(18654481),
		Addresses: []common.Address{
			contractAddress,
		},
		//Topics: the first value of Topics represent the event signature that can identify specific event.
		Topics: [][]common.Hash{
			{common.HexToHash("0x16be4c5129a4e03cf3350262e181dc02ddfb4a6008d925368c0899fcd97ca9c5")},
		},
	}

	//Among those log Filter the log that need.
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.ContractsMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
		log.Fatal(err)
	}

	for _, vLog := range logs {

		event, err := contractAbi.Unpack("StateBatchAppended", vLog.Data)
		if err != nil {
			log.Fatal(err)

		}

		fmt.Println("\nEvent Data : ", event, "")

		fmt.Println("\nBlockHash", vLog.BlockHash.Hex())
		fmt.Println("\nBlockNumber", vLog.BlockNumber)
		fmt.Println("\nTxHASH", vLog.TxHash.Hex())

		txHash := common.HexToHash(vLog.TxHash.Hex())
		tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(isPending, "\n")
		fmt.Println("data", tx.Data())

		fmt.Printf("\n----------------------------------------------------------------------------------------------------------------------------------")
		fmt.Printf("\n\nTx Input Data: %+v\n\n", common.Bytes2Hex(tx.Data()))
		a := strings.TrimLeft(common.Bytes2Hex(tx.Data()), "8ca5cbb9")
		n := 64
		result := []string{}
		for i := 0; i < len(a); i += n {
			end := i + n
			if end > len(a) {
				end = len(a)
			}
			result = append(result, a[i:end])
		}
		fmt.Printf("---------------------------------------------------------------------------------------------------------------------------------\n\n")
		for i, v := range result {
			fmt.Printf("%d Index: %s\n", i, v)
		}
		fmt.Printf("\n----------------------------------------------------------------------------------------------------------------------------------\n")

	}
}
