package nfts

import (
	"market_store_db/utils/rest_errors"
	"market_store_db/utils/validation_utils"
	"strings"
)

type Listing struct {
	Id           int64  `json:"id"`
	OwnerAddress string `json:"owner_address"`
	NftAddress   string `json:"nft_address"`
	NftId        string `json:"nft_id"`
	PriceWei     string `json:"price_wei"`
	Deadline     string `json:"deadline"`
}

func (listing *Listing) Validate() rest_errors.RestErr {

	if err := validation_utils.IsAddress(listing.OwnerAddress, "owner"); err != nil {
		return err
	}

	if err := validation_utils.IsAddress(listing.NftAddress, "nft"); err != nil {
		return err
	}

	if err := validation_utils.IsNumber(listing.NftId, "nft id"); err != nil {
		return err
	}

	if err := validation_utils.IsNumber(listing.PriceWei, "amount"); err != nil {
		return err
	}

	if listing.Deadline != ""{
		if t, err := validation_utils.ConvTimeToDB(strings.TrimSpace(listing.Deadline)); err != nil {
			return err
		} else {
			listing.Deadline = t
		}
	}

	return nil
}
