package services

import (
	"context"
	"fmt"

	"github.com/paycrest/paycrest-protocol/ent"
	"github.com/paycrest/paycrest-protocol/utils/crypto"
)

// ReceiveAddressService provides functionality releated to managing receive addresses
type ReceiveAddressService struct {
	db *ent.Client
}

// NewReceiveAddressService creates a new instance of NewReceiveAddress.
func NewReceiveAddressService(db *ent.Client) *ReceiveAddressService {
	return &ReceiveAddressService{
		db: db,
	}
}

// GenerateAndSaveAddress function generates a new address for a user
func (s *ReceiveAddressService) GenerateAndSaveAddress(ctx context.Context) (string, error) {
	count, err := s.db.ReceiveAddress.
		Query().
		Count(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to query receive addresses: %w", err)
	}

	// accountIndex = number of receive addresses in DB + 1
	accountIndex := count + 1
	address, _, err := crypto.GenerateReceiveAddress(accountIndex)

	_, err = s.db.ReceiveAddress.
		Create().
		SetAddress(address).
		SetAccountIndex(accountIndex).
		SetStatus("unused").
		Save(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to save address: %w", err)
	}

	return address, nil
}
