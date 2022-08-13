package nfts

import (
	"market_store_db/datasources/mysql/nfts_db"
	"market_store_db/utils/rest_errors"

	"errors"
)

const (
	queryInsertListing = "INSERT INTO nfts(owner_address, nft_address, nft_id, price_wei, deadline) VALUES(?, ?, ?, ?, ?);"
	queryGetListing = "SELECT id, owner_address, nft_address, nft_id, price_wei, deadline FROM nfts WHERE id=?;"
)

func (listing *Listing) Save() rest_errors.RestErr {
	stmt, err := nfts_db.Client.Prepare(queryInsertListing)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to prepare save listing statement", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(listing.OwnerAddress, listing.NftAddress, listing.NftId, listing.PriceWei, listing.Deadline)
	if saveErr != nil {
		return rest_errors.NewInternalServerError("error when trying to save listing", errors.New("database error"))
	}
	
	listingId, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get last insert id after creating a new listing", errors.New("database error"))
	}
	listing.Id = listingId

	return nil
}

func(listing *Listing) Get() rest_errors.RestErr {
	stmt, err := nfts_db.Client.Prepare(queryGetListing)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to prepare get listing statement", errors.New("database error"))
	}
	defer stmt.Close()

	result := stmt.QueryRow(listing.Id)

	if getErr := result.Scan(&listing.Id, &listing.OwnerAddress, &listing.NftAddress, &listing.NftId, &listing.PriceWei, &listing.Deadline); getErr != nil {
		return rest_errors.NewInternalServerError("error when trying to get listing by id", errors.New("database error"))
	}

	return nil
}