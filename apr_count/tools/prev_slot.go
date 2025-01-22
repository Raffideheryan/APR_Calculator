package tools

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	strc "apr_count/structs"

	"github.com/sirupsen/logrus"
)

func GetHeadSlot(HeadUrl string) int64 {
	resp, err := http.Get(HeadUrl)
	if err != nil {
		logrus.WithError(err).Fatal("Can't get head slot")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.WithError(err).Fatalf("Something went wrong %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.WithError(err).Fatal("Can't read the response body")
	}

	var rsp strc.Slot
	if err := json.Unmarshal(body, &rsp); err != nil {
		logrus.WithError(err).Fatal("Error while decoding data")
	}

	SlotNumber := HexToString(rsp.BlockData.Msg.Slot)

	return SlotNumber

}

func GetPrevSlot(HeadSlot int64) string {
	const SecPerYear = 31536000

	PrevSlot := HeadSlot - (SecPerYear / 12)
	prevSlotStr := strconv.FormatInt(PrevSlot, 10)

	return prevSlotStr

}
