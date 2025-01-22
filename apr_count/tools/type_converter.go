package tools

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/sirupsen/logrus"
)

func StrToBigInt(str string) (*big.Int, error) {
	balance, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return nil, fmt.Errorf("invalid number format")
	}
	return balance, nil
}

func IntToBigInt(num int64) *big.Int {
	return new(big.Int).SetInt64(num)
}

func HexToString(str string) int64 {

	HexValue, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		logrus.WithError(err).Fatal("Can't convert Invalid Type")
	}

	return HexValue
}
