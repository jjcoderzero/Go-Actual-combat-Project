package resource

import (
	"server/devops-tools/mysql/initconfig"
	"server/models"
)

func CreateProduct(code string, price int) {
	db := initconfig.InitConfig()
	db.Create(&models.Product{
		Code:  code,
		Price: price,
	})
}

func ReadProduct(code string) (product models.Product) {
	db := initconfig.InitConfig()
	db.First(&product, 1)
	db.First(&product, "code=?", code)
	return product
}

func UpdateProduct() {
	db := initconfig.InitConfig()
	db.Model(&models.Product{}).Update("")
}
