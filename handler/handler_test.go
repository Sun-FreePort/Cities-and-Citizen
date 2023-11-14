package handler

import (
	"github.com/Sun-FreePort/Cities-and-Citizen/config"
	"github.com/Sun-FreePort/Cities-and-Citizen/model"
	"github.com/Sun-FreePort/Cities-and-Citizen/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"os"
	"testing"
)

var (
	d *gorm.DB
	r *router.Router
	e *fiber.App
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	configDict := config.GetConfig("")
	r = &router.Router{
		H: NewHandler(),
	}
	e = r.NewApp(configDict, nil)

	d = config.GetDB()
	config.AutoMigrate(d)
	config.GetRedis().FlushDB()
	r.RegisterF2E(e)
	err := loadFixtures()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func tearDown() {
	if err := config.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func loadFixtures() error {
	u1 := model.UserModel{
		ID:               1,
		Name:             "tester1",
		AvatarPath:       "",
		CityId:           1,
		Age:              1,
		Credit:           10,
		Money:            400,
		Energy:           0,
		Hungry:           0,
		Happy:            0,
		Health:           0,
		EquipHandId:      0,
		EquipBodyId:      0,
		EquipTrouserId:   0,
		SkillFarm:        0,
		SkillLumber:      0,
		SkillMine:        0,
		SkillHandwork:    0,
		SkillFight:       0,
		SkillFarmExp:     0,
		SkillLumberExp:   0,
		SkillMineExp:     0,
		SkillHandworkExp: 0,
		SkillFightExp:    0,
	}
	if err := d.Create(&u1).Error; err != nil {
		return err
	}

	return nil
}
