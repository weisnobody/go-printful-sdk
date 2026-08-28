package printfulsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/baldurstod/go-printful-sdk/model/responses"
)

func (c *PrintfulClient) AddMockupTask(storeID string, format string, mockup_width_px int, products []model.MockupTaskProduct, opts ...RequestOption) (*[]responses.MockupGeneratorTask, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.url = PRINTFUL_MOCKUP_ENDPOINT

	body := map[string]interface{}{}
	//body := BuildRequestBody(opt, FileRole, URL, Filename, FileVisible)
	if !slices.Contains([]string{"jpg", "png", }, format) {
		return nil, errors.New("Unknown Format")
	}
	body["format"] = format

	if mockup_width_px < 50 {
		return nil, errors.New("Invalid Mockup Width")
	}
	body["mockup_width_px"] = mockup_width_px

	body["products"] = products

	headers := make(map[string]string)
	headers["X-PF-Store-Id"] = storeID

	u := PRINTFUL_MOCKUP_ENDPOINT
	resp, err := c.Post(u, headers, body, ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("post returned an error in AddMockupTask: %w", err)
	}

	response := &responses.MockupTasksResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response in AddMockupTask")
	}

	return &response.Data, nil
}

func (c *PrintfulClient) GetMockupTasks(storeID string, taskIDs []int, opts ...RequestOption) ([]responses.MockupGeneratorTask, error) {
	opt := getOptions(opts...)

	mockupTasksResults := make([]responses.MockupGeneratorTask, 0, 400)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	opt.offset = 0
	opt.limit = 100
	headers := make(map[string]string)
	headers["X-PF-Store-Id"] = storeID

	baseURL := PRINTFUL_MOCKUP_ENDPOINT + "?id="
	idJoin := ""
	for i := range taskIDs {
		baseURL = fmt.Sprintf("%s%s%s", baseURL, idJoin, strconv.Itoa(taskIDs[i]))
		idJoin = ","
	}
	fmt.Printf("Base: %s\n", baseURL)

	for {

		u, _ := buildURL(baseURL, opt)
		log.Println(u)
		fmt.Printf("Get: %s\n", u)
		//return mockupTasksResults, errors.New("Bailing")
		resp, err := c.Get(u, headers, ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to get get mockup tasks response")
		}
		defer resp.Body.Close()

		response := &responses.MockupTasksResponse{}

		/* */
		//responseX := make(map[string]interface{})
		/*
			fmt.Println("# Response Body", resp.StatusCode, "#")
			respBody, _ := io.ReadAll(resp.Body)
			fmt.Println(string(respBody))
			json.Unmarshal(respBody, &response)

			return nil, errors.New("Check Out")
			//*/

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to decode get mockup tasks response")
		}

		for _, mockupTasksResult := range response.Data {
			mockupTasksResults = append(mockupTasksResults, mockupTasksResult)
		}

		next := response.Paging.Offset + response.Paging.Limit
		if next >= response.Paging.Total {
			break
		}
		opt.offset = next
		opt.limit = response.Paging.Limit
	}

	return mockupTasksResults, nil
}
