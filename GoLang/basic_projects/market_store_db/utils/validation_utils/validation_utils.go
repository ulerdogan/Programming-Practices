package validation_utils

import (
	"strings"
	"fmt"
	"market_store_db/utils/rest_errors"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"time"
)

const dbTimeLayout = "2006-01-02 15:04:05"

func IsAddress(address string, t string) rest_errors.RestErr {
	address = strings.TrimSpace(address)

	if address == "" || !common.IsHexAddress(address) {
		return rest_errors.NewBadRequestError(fmt.Sprintf("invalid %v address", t))
	}

	return nil
}

func IsNumber(number string, t string) rest_errors.RestErr {
	number = strings.TrimSpace(number)

	_, err := strconv.Atoi(number)
	if number == "" || err != nil {
		return rest_errors.NewBadRequestError(fmt.Sprintf("invalid %v", t))
	}

	return nil
}

func ConvTimeToDB(tm string) (string, rest_errors.RestErr) {
	t, err := time.Parse(dbTimeLayout, tm)
	if err != nil {
		return "", rest_errors.NewBadRequestError("invalid timestamp")
	}

	return t.Format(dbTimeLayout), nil
}