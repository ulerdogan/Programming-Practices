package nfts

import (
	"market_store_db/datasources/mysql/nfts_db"
	"market_store_db/utils/rest_errors"

	"errors"
)

const (
	queryInsertListing = "INSERT INTO nfts(owner_address, nft_address, nft_id, price_wei, deadline) VALUES(?, ?, ?, ?, ?);"
	queryGetListing = "SELECT id, owner_address, nft_address, nft_id, price_wei, deadline FROM nfts WHERE id=?;"
	queryDeleteListing = "DELETE FROM nfts WHERE id=?;"
	queryUpdateListing = "UPDATE nfts SET owner_address=?, price_wei=?, deadline=? WHERE id=?;"
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

func (listing *Listing) Delete() rest_errors.RestErr {
	stmt, err := nfts_db.Client.Prepare(queryDeleteListing)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to prepare delete listing statement", errors.New("database error"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(listing.Id); err != nil {
		return rest_errors.NewInternalServerError("error when trying to delete user", errors.New("database error"))
	}
	return nil
}

func (listing *Listing) Update() rest_errors.RestErr {
	stmt, err := nfts_db.Client.Prepare(queryUpdateListing)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to prepare update listing statement", errors.New("database error"))
	}
	defer stmt.Close()

	stmt.Exec(listing.OwnerAddress, listing.PriceWei, listing.Deadline, listing.Id)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to update listing", errors.New("database error"))
	}
	return nil
}