package conv

import (
	"github.com/w-woong/order/dto"
	"github.com/w-woong/order/entity"
	"github.com/wonksing/structmapper"
)

func init() {
	structmapper.StoreMapper(&dto.Cart{}, &entity.Cart{})
	structmapper.StoreMapper(&entity.Cart{}, &dto.Cart{})
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
