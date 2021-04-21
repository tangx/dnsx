package domain

import "gorm.io/gorm"

type Domain struct {
	gorm.Model
	Name string
}

func CreateDomain(name string) {

}
