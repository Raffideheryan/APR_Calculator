package core

import (
	strc "apr_count/structs"
	tools "apr_count/tools"
	"fmt"

	"math/big"

	"github.com/sirupsen/logrus"
)

//Calculating balances for all validators
func CalculateBalances(mapBefore, mapAfter map[string]*strc.ValidatorData){
	OverAllBalancesBefore := big.NewInt(0)
	for _, v := range mapBefore {
		num, err := tools.StrToBigInt(v.Balance)
		if err != nil {
			logrus.WithError(err).Fatal("Can't convert the type")
		}
		OverAllBalancesBefore = new(big.Int).Add(OverAllBalancesBefore, num)
	}

	OverAllBalancesNow := big.NewInt(0)
	for k := range mapBefore {
		num, err := tools.StrToBigInt(mapAfter[k].Balance)
		if err != nil {
			logrus.WithError(err).Fatal("Can't convert the type")
		}
		OverAllBalancesNow = new(big.Int).Add(OverAllBalancesNow, num)
	}

	res, _ := AprCalculation(OverAllBalancesBefore, OverAllBalancesNow).Float64()
	fmt.Printf("APR = %.3f%%\n", (1-res)*100)


}

func AprCalculation(BalanceNow, BalancePrev *big.Int) *big.Float {
	before, _ := BalancePrev.Float64()
	after, _ := BalanceNow.Float64()
	beforeFloat := big.NewFloat(before)
	afterFloat := big.NewFloat(after)
	APR := new(big.Float).Quo(afterFloat, beforeFloat)

	return APR
}
