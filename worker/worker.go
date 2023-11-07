package worker

import (
	"CustomerLabsTest/model"
	"regexp"
)

func Worker(input chan map[string]interface{}, output chan *model.OuputData) {
	defer close(output)
	inputData := <-input
	// Convert the input data to the desired output structure
	outputData := &model.OuputData{
		Event:            inputData["ev"].(string),
		Event_Type:       inputData["et"].(string),
		App_Id:           inputData["id"].(string),
		User_Id:          inputData["uid"].(string),
		Message_Id:       inputData["mid"].(string),
		Page_Title:       inputData["t"].(string),
		Page_Url:         inputData["p"].(string),
		Browser_Language: inputData["l"].(string),
		Screen_Size:      inputData["sc"].(string),
		Attributes:       make(map[string]model.Attribute),
		Traits:           make(map[string]model.Trait),
	}

	attributeRegex := regexp.MustCompile(`^atrk(\d+)$`)
	traitRegex := regexp.MustCompile(`^uatrk(\d+)$`)

	// Iterate over input data to extract attributes and traits
	for key, value := range inputData {
		if matches := attributeRegex.FindStringSubmatch(key); len(matches) == 2 {
			outputData.Attributes[value.(string)] = model.Attribute{
				Value: inputData["atrv"+matches[1]].(string),
				Type:  inputData["atrt"+matches[1]].(string),
			}

		} else if matches := traitRegex.FindStringSubmatch(key); len(matches) == 2 {
			outputData.Traits[value.(string)] = model.Trait{
				Value: inputData["uatrv"+matches[1]].(string),
				Type:  inputData["uatrt"+matches[1]].(string),
			}
		}
	}

	output <- outputData

}
