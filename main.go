package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	store "./contracts"
)


func main() {
	client, err := ethclient.Dial("https://kovan.infura.io/v3/")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("PRIKEYPRIKEYPRIKEYPRIKEYPRIKEYPRIKEYPRIKEYPRIKEYPRIKEYPRIKEYPRIK")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("error casting public key to ECDSA")
		log.Fatal(err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddress := common.HexToAddress("0xf9a39a6443ae5dD43d093C0e5d29f08656fa81cD")

	f:= func(_s types.Signer,_addr common.Address,_tx *types.Transaction) (*types.Transaction, error){
		return types.SignTx(_tx, _s, privateKey)
	}

	instance,err:=store.NewStore(toAddress,client)
	param:=&bind.TransactOpts{
		From:fromAddress,
		Signer:f,
		Context:context.Background(),
	}
	signedTx,err:=instance.AFunc(param)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}