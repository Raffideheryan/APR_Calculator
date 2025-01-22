package core

import (
	strc "apr_count/structs"
	tools "apr_count/tools"

	"github.com/sirupsen/logrus"
)

var (
	BaseUrl     = "your_url_to_base_blockchain"
	HeadSlotUrl = BaseUrl + "your_url_to_head_slot"
)


//Sending Response to get headslot and previous year slot numbers 
func ResponseSender() {
	headSlotNum := tools.GetHeadSlot(HeadSlotUrl)
	prevSlot := tools.GetPrevSlot(headSlotNum)

	ResponseNow, err := GetValidators("head")
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get validators")
	}

	ResponsePrev, err := GetValidators(prevSlot)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get validators")
	}

	DataMapping(ResponseNow, ResponsePrev)
}

//Saving Data after Response in maps to filter ongoing validators now
func DataMapping(ResponseNow, ResponsePrev *strc.Response) {

	mapBefore := make(map[string]*strc.ValidatorData)
	mapAfter := make(map[string]*strc.ValidatorData)

	for _, val := range ResponseNow.Data {
		mapAfter[val.Index] = &val
	}
	for _, val := range ResponsePrev.Data {
		mapBefore[val.Index] = &val
	}

	result := FilterMaps(mapBefore, mapAfter)

	CalculateBalances(result, mapAfter)
}
