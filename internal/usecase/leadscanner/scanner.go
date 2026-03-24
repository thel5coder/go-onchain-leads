package leadscanner

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"go-onchain-leads/internal/domain"
	"log"
	"math/big"
	"time"
)

type LeadScanner struct {
	reader   BlockchainReader
	resolver IdentityResolver
	saver    LeadSaver
}

func NewLeadScanner(r BlockchainReader, id IdentityResolver, s LeadSaver) *LeadScanner {
	return &LeadScanner{
		reader:   r,
		resolver: id,
		saver:    s,
	}
}

func (s *LeadScanner) StartScanning() {
	chainID, err := s.reader.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get network ID: %v", err)
	}

	var lastScannedBlock uint64 = 0
	fmt.Println("🎧 Clean Arch Scanner is LIVE. Filtering and Unmasking Leads...")

	for {
		header, err := s.reader.HeaderByNumber(context.Background(), nil)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}

		latestBlock := header.Number.Uint64()

		if latestBlock > lastScannedBlock {
			startBlock := lastScannedBlock + 1
			if lastScannedBlock == 0 {
				startBlock = latestBlock
			}

			for i := startBlock; i <= latestBlock; i++ {
				blockNumber := new(big.Int).SetUint64(i)
				block, err := s.reader.BlockByNumber(context.Background(), blockNumber)
				if err != nil {
					continue
				}

				for _, tx := range block.Transactions() {
					if tx.To() == nil {
						receipt, err := s.reader.TransactionReceipt(context.Background(), tx.Hash())
						if err != nil || receipt.GasUsed < 300000 {
							continue
						}

						signer := types.LatestSignerForChainID(chainID)
						sender, err := types.Sender(signer, tx)
						if err != nil {
							continue
						}

						walletHex := sender.Hex()
						identity := s.resolver.ResolveIdentity(walletHex)

						lead := domain.Lead{
							DeveloperWallet: walletHex,
							ENSName:         identity.Name,
							Twitter:         identity.Twitter,
							Email:           identity.Email,
							GitHub:          identity.GitHub,
							ContractAddress: receipt.ContractAddress.Hex(),
							TransactionHash: tx.Hash().Hex(),
							GasUsed:         receipt.GasUsed,
						}

						fmt.Println("\n💎 HIGH-QUALITY LEAD UNMASKED!")
						if lead.ENSName != "Anonymous" {
							fmt.Printf("Identity:         %s\n", lead.ENSName)
							if lead.Twitter != "" {
								fmt.Printf("🐦 Twitter:       https://twitter.com/%s\n", lead.Twitter)
							}
							if lead.Email != "" {
								fmt.Printf("✉️ Email:         %s\n", lead.Email)
							}
						} else {
							fmt.Printf("Identity:         Anonymous Wallet\n")
						}
						fmt.Printf("Developer Wallet: %s\n", lead.DeveloperWallet)
						fmt.Printf("Gas Spent:        %d units\n", lead.GasUsed)
						fmt.Println("---------------------------------------------------")

						err = s.saver.SaveLead(lead)
						if err != nil {
							log.Printf("⚠️ Failed to save lead: %v\n", err)
						}
					}
				}
			}
			lastScannedBlock = latestBlock
		}
		time.Sleep(3 * time.Second)
	}
}
