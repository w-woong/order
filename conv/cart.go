package conv

import (
	"github.com/w-woong/order/dto"
	"github.com/w-woong/order/entity"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&dto.Cart{}, &entity.Cart{})
	structmapper.StoreMapper(&entity.Cart{}, &dto.Cart{})

	structmapper.StoreMapper(&dto.CartProduct{}, &entity.CartProduct{})
	structmapper.StoreMapper(&entity.CartProduct{}, &dto.CartProduct{})
}

// ToUserEntity converts dto.User to entity.User
func ToCartEntity(input *dto.Cart) (output entity.Cart, err error) {
	err = structmapper.Map(input, &output)
	return
}

// ToUserDto converts entity.User to dto.User.
func ToCartDto(input *entity.Cart) (output dto.Cart, err error) {
	err = structmapper.Map(input, &output)
	return
}

func ToCartProductEntity(input *dto.CartProduct) (output entity.CartProduct, err error) {
	err = structmapper.Map(input, &output)
	return
}

func ToCartProductDto(input *entity.CartProduct) (output dto.CartProduct, err error) {
	err = structmapper.Map(input, &output)
	return
}

func ToCartProductListDto(input entity.CartProductList) (dto.CartProductList, error) {
	output := make(dto.CartProductList, len(input))
	for i := range input {
		c, err := ToCartProductDto(&input[i])
		if err != nil {
			return nil, err
		}
		output[i] = c
	}
	return output, nil
}

func ToCartProductListEntity(input dto.CartProductList) (entity.CartProductList, error) {
	output := make(entity.CartProductList, len(input))
	for i := range input {
		c, err := ToCartProductEntity(&input[i])
		if err != nil {
			return nil, err
		}
		output[i] = c
	}
	return output, nil
}
