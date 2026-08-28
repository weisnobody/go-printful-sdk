package printfulsdk

// This is a v1 API call.  They do not yet exist in V2 that I could find / figure out

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/baldurstod/go-printful-sdk/model/responses"
)


func (c *PrintfulClient) GetPrintFiles(storeID string, productId int, technique string, opts ...RequestOption) (model.PrintFileInfo, error) {
	opt := getOptions(opts...)

	printfile := model.PrintFileInfo{}

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	headers := make(map[string]string)
	headers["X-PF-Store-Id"] = storeID

	u, _ := buildURL(PRINTFUL_PRINTFILES+"/"+strconv.Itoa(productId)+"?technique="+technique, opt)
	log.Println(u)
	resp, err := c.Get(u, headers, ctx)
	if err != nil {
		log.Println(err)
		return printfile, errors.New("unable to get printfile response")
	}
	defer resp.Body.Close()

	response := &responses.Printfile{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return printfile, errors.New("unable to decode printfile response")
	}

	printfile = response.Result

	return printfile, nil
}
