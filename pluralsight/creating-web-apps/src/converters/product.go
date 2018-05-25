package converters

import (
"github.com/GOCODE/pluralsight/creating-web-apps/src/models"
"github.com/GOCODE/pluralsight/creating-web-apps/src/viewmodels"
)

func ConvertProductToViewModel(product models.Product) viewmodels.Product {
	result := viewmodels.Product{
		Name: product.Name(),
		DescriptionShort: product.DescriptionShort(),
		DescriptionLong: product.DescriptionLong(),
		PricePerLiter: product.PricePerLiter(),
		PricePer10Liter: product.PricePer10Liter(),
		Origin: product.Origin(),
		IsOrganic: product.IsOrganic(),
		ImageUrl: product.ImageUrl(),
		Id: product.Id(),
	}
	
	return result
}