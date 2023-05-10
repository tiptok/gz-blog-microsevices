package types

import "github.com/tiptok/gz-blog-microsevices/app/skateboard/interanl/pkg/domain"

var GlobalShopConvertor = ShopConvertor{}
var GlobalAddressConvertor = AddressConvertor{}

type ShopConvertor struct {
}

func (c ShopConvertor) EntityToDomain(item ShopItem) *domain.Shop {
	return &domain.Shop{
		Name:         item.Name,
		Introduction: item.Introduction,
		Rank:         item.Rank,
		Address:      GlobalAddressConvertor.EntityToDomain(item.Address),
	}
}

func (c ShopConvertor) DomainToEntity(item *domain.Shop) ShopItem {
	return ShopItem{
		ShopId:       int(item.Id),
		Name:         item.Name,
		Introduction: item.Introduction,
		Rank:         item.Rank,
		Address:      GlobalAddressConvertor.DomainToEntity(item.Address),
	}
}

type AddressConvertor struct {
}

func (c AddressConvertor) EntityToDomain(address *Address) *domain.Address {
	if address == nil {
		return &domain.Address{}
	}
	return &domain.Address{
		Address: address.Address,
		Lat:     address.Lat,
		Lon:     address.Lon,
	}
}
func (c AddressConvertor) DomainToEntity(address *domain.Address) *Address {
	if address == nil {
		return &Address{}
	}
	return &Address{
		Address: address.Address,
		Lat:     address.Lat,
		Lon:     address.Lon,
	}
}
