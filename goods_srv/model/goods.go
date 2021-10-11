package model

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(20);not null"`
	Level            int32  `gorm:"type:int;not null;default:1"`
	IsTab            bool   `gorm:"default:false;not null"`
	ParentCategoryID int32  `gorm:"type:int;not null"`
	ParentCategory   *Category
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(20);not null;default:''"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;not null;default:1"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:"default:false;not null"`

	Name        string  `gorm:"type:varchar(50);not null"`
	GoodsSn     string  `gorm:"type:varchar(50);not null"`
	ClickNum    int32   `gorm:"type:int;not null;default:0"`
	SoldNum     int32   `gorm:"type:int;not null;default:0"`
	FavNum      int32   `gorm:"type:int;not null;default:0"`
	MarketPrice float32 `gorm:"not null"`
	ShopPrice   float32 `gorm:"not null"`
	GoodsBrief  string  `gorm:"type:varchar(100);not null"`

	Images          GormList `gorm:"type:varchar(1000);not null"`
	DescImages      GormList `gorm:"type:varchar(1000);not null"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null"`
}

type GoodsImages struct {
	GoodsID int
}