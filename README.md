# go-printful-sdk

Wrapper for the [Printful API V2](https://developers.printful.com/docs/v2-beta)

# Usage #

With your Printful API token in the PF_TOKEN environment variable, retrieve a Product (from the Printful Catalog) and output the type of product, the product name, the number of variants, and the available sizes and colors.

```go
package main

import (
	"fmt"
	printfulsdk "github.com/baldurstod/go-printful-sdk"
	"os"
	"strings"
)

func printfulClient() *printfulsdk.PrintfulClient {

	token := os.Getenv("PF_TOKEN")
	if len(token) == 0 {
		fmt.Println("Need Printful API Token")
		return nil
	}

	return printfulsdk.NewPrintfulClient(token)
}

func main() {

	pfClient := printfulClient()
	if pfClient == nil {
		fmt.Printf("Unable to create Printful Client")
		return
	}

	product, err := pfClient.GetCatalogProduct(306)
	if err != nil {
		fmt.Printf("Get CatalogProduct Err: %v\n", err)
		return
	}

	var colors []string
	for _, color := range product.Colors {
		colors = append(colors, fmt.Sprintf("%s (%s)", color.Name, color.Value))
	}

	fmt.Printf(`# %s #

* Type: %s
* Name: %s 
* Brand: %s 
* Model: %s
* Variants: %d
* Sizes: %s
* Colors: %s
* [Image](%s)
`,
		product.Name,
		product.Type,
		product.Name,
		product.Brand,
		product.Model,
		product.VariantCount,
		strings.Join(product.Sizes, ", "),
		strings.Join(colors, ", "),
		product.Image,
	)

}
```

## Output ##

```markdown
# Toddler Staple Tee | Bella + Canvas 3001T #

* Type: T-SHIRT
* Name: Toddler Staple Tee | Bella + Canvas 3001T
* Brand: Bella + Canvas
* Model: 3001T
* Variants: 16
* Sizes: 2T, 3T, 4T, 5T
* Colors: Black (#060606), Heather Columbia Blue (#6495ff), Pink (#ffcfd9), White (#FFFFFF)
* [Image](https://files.cdn.printful.com/upload/product-catalog-img/70/7025b639b21097b95e85cd73b32bdd1e_l)
```
