package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saiteja111997/chatGPT_service/pkg/utilities"
)

const (
	apiURL = "https://api.openai.com/v1/completions"
	model  = "text-davinci-003"
)

func GenerateFinalResultV2(c *gin.Context) {

	var request utilities.WebhookRequest
	err := c.ShouldBind(&request)

	if err != nil {
		log.Fatal(err.Error())
	}

	// prompt := "hi"
	prompt := request.IntentInfo.Parameters["any"].OriginalValue

	reqBody := map[string]interface{}{
		"model":             model,
		"prompt":            prompt,
		"temperature":       0.5,
		"max_tokens":        100,
		"top_p":             1,
		"frequency_penalty": 0,
		"presence_penalty":  0,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println(err)
		return
	}

	// openAiApiKey := os.Getenv("OPEN_AI_API_KEY")
	openAiApiKey := "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"
	fmt.Println(openAiApiKey)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+openAiApiKey)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response map[string]interface{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Print response : ", response)

	fmt.Println(response["choices"].([]interface{})[0].(map[string]interface{})["text"])

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "Convo summary!!",
	// 	"Tasks":   response["choices"].([]interface{})[0].(map[string]interface{})["text"],
	// })

	webhookResponse := utilities.WebhookResponse{
		FulfillmentResponse: utilities.Message{
			Messages: []utilities.MessageObject{{
				Text: utilities.TextObj{
					Text: []string{response["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)},
				},
			}},
		},
	}

	c.JSON(http.StatusOK, webhookResponse)

}

func TestWebhook(c *gin.Context) {

	var request utilities.WebhookRequest
	err := c.ShouldBind(&request)

	if err != nil {
		fmt.Println("The error is : ", err.Error())
	}

	// body := c.Request.Body

	// defer body.Close()

	// bytes, err := ioutil.ReadAll(body)
	// if err != nil {
	// 	fmt.Println("Error occured : ", err)
	// 	log.Fatal(err.Error())
	// }

	// err = json.Unmarshal(bytes, &request)
	// if err != nil {
	// 	fmt.Println("Error occured : ", err)
	// 	log.Fatal(err.Error())
	// }

	fmt.Println("Printing the request : ", request)
	fmt.Println("Printing the prompt : ", request.IntentInfo.Parameters["any"].OriginalValue)

	response := utilities.WebhookResponse{
		FulfillmentResponse: utilities.Message{
			Messages: []utilities.MessageObject{{
				Text: utilities.TextObj{
					Text: []string{"Hii there, this is a chatGPT service!!"},
				},
			}},
		},
	}

	c.JSON(http.StatusOK, response)

}
