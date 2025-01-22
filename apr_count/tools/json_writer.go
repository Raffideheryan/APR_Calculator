package tools

import (
	strc "apr_count/structs"
	"encoding/json"
	"fmt"

	"os"

	"github.com/sirupsen/logrus"
)

// Writing Validators data from head slot into json file
func JsonValidatorWriter(response strc.Response, SlotNum string) {
	file, err := os.Create(fmt.Sprintf("json_data/validators_%s.json", SlotNum))
	if err != nil {
		logrus.WithError(err).Fatal("Can't create the file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	err = encoder.Encode(response)
	if err != nil {
		logrus.WithError(err).Fatal("Error while writing into file")
	}

	logrus.Info(fmt.Sprintf("Successfully saved validator data for slot %s", SlotNum))
}
