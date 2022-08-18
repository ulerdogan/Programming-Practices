package services

import (
	"market_store_db/domain/nfts"
	"market_store_db/utils/rest_errors"
)

var NftsService nftsServiceInterface = &nftsService{}

type nftsService struct{}

type nftsServiceInterface interface {
	CreateListing(nfts.Listing) (*nfts.Listing, rest_errors.RestErr)
	GetListing(id int64) (*nfts.Listing, rest_errors.RestErr)
	DeleteListing(id int64) (rest_errors.RestErr)
	UpdateListing(isPartial bool, listing nfts.Listing) (*nfts.Listing, rest_errors.RestErr)
}

func (l *nftsService) CreateListing(listing nfts.Listing) (*nfts.Listing, rest_errors.RestErr) {
	if err := listing.Validate(); err != nil {
		return nil, err
	}

	if err :=  listing.Save(); err != nil {
		return nil, err
	}
	
	return &listing, nil
}

func (l *nftsService) GetListing(id int64) (*nfts.Listing, rest_errors.RestErr) {
	nft := &nfts.Listing{Id: id}
	if err := nft.Get(); err != nil {
		return nil, err
	}

	return nft, nil
}

func (l *nftsService) DeleteListing(id int64) rest_errors.RestErr {
	nft := &nfts.Listing{Id: id}
	if err := nft.Delete(); err != nil {
		return err
	}
	return nil
}

func (l *nftsService) UpdateListing(isPartial bool, listing nfts.Listing) (*nfts.Listing, rest_errors.RestErr) {
	current := &nfts.Listing{Id: listing.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if isPartial {
		if listing.OwnerAddress != "" {
			current.OwnerAddress = listing.OwnerAddress
		}
		if listing.PriceWei != "" {
			current.PriceWei = listing.PriceWei
		}
		if listing.Deadline != "" {
			current.Deadline = listing.Deadline
		}
	} else {
		current.OwnerAddress = listing.OwnerAddress
		current.PriceWei = listing.PriceWei
		current.Deadline = listing.Deadline
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}