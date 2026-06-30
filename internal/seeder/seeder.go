package seeder

import "gorm.io/gorm"

func SeedData(db *gorm.DB) error {
	if err := InitOfficeSeed(db); err != nil {
		return err
	}

	if err := InitCasbinSeed(db); err != nil {
		return err
	}

	if err := InitUserAdmin(db); err != nil {
		return err
	}
	return nil
}
