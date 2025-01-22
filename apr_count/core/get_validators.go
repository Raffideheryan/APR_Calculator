package core

import (
	strc "apr_count/structs"
	tools "apr_count/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Getting Validators Data from head slot
func GetValidators(SlotNum string) (*strc.Response, error) {

	BaseUrl := "your_url_to_blockchain"
	ValidatorsUrl := BaseUrl + fmt.Sprintf("url_to_%s slot", SlotNum)

	resp, err := http.Get(ValidatorsUrl)
	if err != nil {
		logrus.WithError(err).Fatal("Can't get validators")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.WithError(err).Fatalf("Something went wrong %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Fatal("Error while reading data")
	}

	var rspRaw strc.Response
	var rspResult strc.Response
	if err := json.Unmarshal(body, &rspRaw); err != nil {
		logrus.WithError(err).Fatal("Error while decoding data")
	}
	for _, val := range rspRaw.Data {
		if val.Status == "active_ongoing" {
			rspResult.Data = append(rspResult.Data, val)
		}
	}

	tools.JsonValidatorWriter(rspResult, SlotNum)

	return &rspResult, nil
}

//Filtering ongoing validators now
func FilterMaps(mapBefore, mapAfter map[string]*strc.ValidatorData) map[string]*strc.ValidatorData {

	result := make(map[string]*strc.ValidatorData)

	for k, v := range mapBefore {
		if mapAfter[k] != nil {
			result[k] = v
		}
	}

	return result
}
