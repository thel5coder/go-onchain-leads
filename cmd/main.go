package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v3"
	"go-onchain-leads/internal/domain"
	"go-onchain-leads/internal/usecase/leadscanner"
	"log"
)

type EthereumNameService struct {
	client *ethclient.Client
}

func (e *EthereumNameService) ResolveIdentity(walletAddress string) domain.Identity {
	address := common.HexToAddress(walletAddress)

	identity := domain.Identity{
		Name: "Anonymous",
	}

	name, err := ens.ReverseResolve(e.client, address)
	if err != nil || name == "" {
		return identity
	}
	identity.Name = name

	resolver, err := ens.NewResolver(e.client, name)
	if err != nil {
		return identity
	}

	twitter, _ := resolver.Text("com.twitter")
	email, _ := resolver.Text("email")
	github, _ := resolver.Text("com.github")

	identity.Twitter = twitter
	identity.Email = email
	identity.GitHub = github

	return identity
}

func main() {
	rpcURL := "https://eth-mainnet.g.alchemy.com/v2/r4x-MwmV9gOcnbn7Ec3IZ"

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to Alchemy: %v", err)
	}

	ensResolver := &EthereumNameService{client: client}
	scannerService := leadscanner.NewLeadScanner(client, ensResolver)
	scannerService.StartScanning()
}
