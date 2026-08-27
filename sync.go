package printfulsdk

// These are v1 API calls.  They do not yet exist in V2.
// Info is under the [Retired Resources](https://developers.printful.com/docs/v2-beta/#section/Errors/Retired-Resources),
//   but the note says they are "not available in version 2 of the API yet"

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/baldurstod/go-printful-sdk/model"
	"github.com/baldurstod/go-printful-sdk/model/responses"
)

func (c *PrintfulClient) GetSyncProducts(storeID string, opts ...RequestOption) ([]model.SyncProduct, error) {
	opt := getOptions(opts...)

	products := make([]model.SyncProduct, 0, 400)

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

	for {

		u, _ := buildURL(PRINTFUL_SYNC_PRODUCTS, opt)
		log.Println(u)
		resp, err := c.Get(u, headers, ctx)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to get printful response")
		}
		defer resp.Body.Close()

		response := &responses.SyncProductsResponse{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Println(err)
			return nil, errors.New("unable to decode printful response")
		}

		for _, syncProduct := range response.Result {
			products = append(products, syncProduct)
		}

		next := response.Paging.Offset + response.Paging.Limit
		if next >= response.Paging.Total {
			break
		}
		opt.offset = next
		opt.limit = response.Paging.Limit
	}

	return products, nil
}

func (c *PrintfulClient) GetSyncProduct(storeID string, productID int64, opts ...RequestOption) (model.SyncProductInfo, error) {

	return c.getSyncProduct(storeID, fmt.Sprintf("%v", productID), opts...)

}

func (c *PrintfulClient) GetSyncProductByExternalID(storeID string, productID string, opts ...RequestOption) (model.SyncProductInfo, error) {

	return c.getSyncProduct(storeID, productID, opts...)

}

// productID (if it's an external ID, is a string.  change to string); other alternative is to have to functions for each ID type, but they'd be the same except for the input
func (c *PrintfulClient) getSyncProduct(storeID string, productID string, opts ...RequestOption) (model.SyncProductInfo, error) {
	opt := getOptions(opts...)

	product := model.SyncProductInfo{}

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

	fmt.Printf("BuildURL: %s/%s\n", PRINTFUL_SYNC_PRODUCTS, productID)
	u, _ := buildURL(fmt.Sprintf("%s/%s", PRINTFUL_SYNC_PRODUCTS, productID), opt)
	log.Println(u)
	resp, err := c.Get(u, headers, ctx)
	if err != nil {
		log.Println(err)
		return product, errors.New("unable to get printful response")
	}
	defer resp.Body.Close()

	response := &responses.SyncProductResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return product, errors.New("unable to decode printful response")
	}

	product = response.Result

	return product, nil
}

func (c *PrintfulClient) GetSyncVariant(storeID string, syncVariantID string, opts ...RequestOption) (model.SyncVariantInfo, error) {
	opt := getOptions(opts...)

	syncVariantInfo := model.SyncVariantInfo{}

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

	u, _ := buildURL(fmt.Sprintf("%s/%v", PRINTFUL_SYNC_VARIANTS, syncVariantID), opt)
	log.Println(u)
	resp, err := c.Get(u, headers, ctx)
	if err != nil {
		log.Println(err)
		return syncVariantInfo, errors.New("unable to get printful response")
	}
	defer resp.Body.Close()

	response := &responses.SyncVariantResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return syncVariantInfo, errors.New("unable to decode printful response")
	}

	syncVariantInfo = response.Result

	return syncVariantInfo, nil
}

func (c *PrintfulClient) ModifySyncVariant(storeID string, syncVariantID int64, syncVariant model.SyncVariant, opts ...RequestOption) (*model.SyncVariantInfo, error) {
	opt := getOptions(opts...)

	var ctx context.Context
	var cancel context.CancelFunc
	if opt.timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), opt.timeout)
		defer cancel()
	}

	headers := make(map[string]string)
	headers["X-PF-Store-Id"] = storeID

	u, _ := buildURL(fmt.Sprintf("%s/%v", PRINTFUL_SYNC_VARIANT, syncVariantID), opt)


	// this likely isn't the best way, but I don't know what that is...
	body := map[string]interface{}{}
	body["variant_id"] = syncVariant.VariantID
	body["files"] = syncVariant.Files
	body["is_ignored"] = syncVariant.IsIgnored
	if len(syncVariant.SKU) > 0 {
		body["sku"] = syncVariant.SKU
	}
	// How to handle price if it isn't set (0 could be a valid value)
	/*
	if syncVariant.RetailPrice > 0 {
		body["retail_price"] = syncVariant.RetailPrice
	}
	*/


	resp, err := c.Put(u, headers, body, ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("put returned an error in ModifySyncVariant: %w", err)
	}

	response := &responses.SyncVariantResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to decode printful response in ModifySyncVariant")
	}

	return &response.Result, nil
}
