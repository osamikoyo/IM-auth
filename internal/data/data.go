package data

import (
	"github.com/golang-jwt/jwt"
	"github.com/osamikoyo/IM-auth/internal/config"
	"github.com/osamikoyo/IM-auth/internal/data/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Storage struct {
	db *gorm.DB
	key string
}

func New(cfg *config.Config) (*Storage, error) {
	g, err := gorm.Open(sqlite.Open(cfg.DSN))
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: g,
	}, nil
}

func (s *Storage) Register(user *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	user.Password = string(hash)

	return s.db.Create(user).Error
}

func generateJWT(id uint64, key string) (string, error) {
	claims := jwt.MapClaims{
		"id" : id,
		"exp" : time.Now().Add(180 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}


func (s *Storage) Login(email, password string) (string, error) {
	var user models.User

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}

	result := s.db.Where(&models.User{
		Email: email,
		Password: string(hash),
	}).Find(&user)
	if result.Error != nil{
		return "", err
	}

	return generateJWT(user.ID, s.key)
}