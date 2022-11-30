package models

type RewardsShop struct {
	IdRewardShop int `gorm:"primaryKey"`
	Name         string
	Price        int
	Url 		string
}

type RewardsUsers struct {
	Id           string `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null"`
	Id_User      string
	Name         string
	IdRewardShop string
	Url 		 string
}

type Coins struct {
	IdUser string
	Coins  int
}

