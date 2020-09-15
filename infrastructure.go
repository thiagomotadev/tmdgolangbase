package tmdgolangbase

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresBaseRepository ...
type PostgresBaseRepository struct {
	db *gorm.DB
}

func (PostgresBaseRepository) newErrBlankID() (err error) {
	err = errors.New("blank id")
	return
}

// GetID ...
func (postgresBaseRepository PostgresBaseRepository) GetID(model interface{}) (id uint, err error) {
	id, err = GetUintFieldValue(model, "ID")
	if id == 0 {
		err = postgresBaseRepository.newErrBlankID()
	}
	return
}

// Connect ...
func (postgresBaseRepository *PostgresBaseRepository) Connect(serverName string) (err error) {
	serverName = strings.ToUpper(serverName)

	host, err := GetEnvVar(fmt.Sprintf("%v_POSTGRES_HOST", serverName))
	if err != nil {
		return
	}

	port, err := GetEnvVarInt64(fmt.Sprintf("%v_POSTGRES_PORT", serverName))
	if err != nil {
		return
	}

	sslMode, err := GetEnvVar(fmt.Sprintf("%v_POSTGRES_SSL_MODE", serverName))
	if err != nil {
		return
	}

	database, err := GetEnvVar(fmt.Sprintf("%v_POSTGRES_DB", serverName))
	if err != nil {
		return
	}

	user, err := GetEnvVar(fmt.Sprintf("%v_POSTGRES_USER", serverName))
	if err != nil {
		return
	}

	password, err := GetEnvVar(fmt.Sprintf("%v_POSTGRES_PASSWORD", serverName))
	if err != nil {
		return
	}

	dsn := fmt.Sprintf("host=%v port=%v sslmode=%v dbname=%v user=%v password=%v",
		host, port, sslMode, database, user, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	err = sqlDB.Ping()
	if err != nil {
		return
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(15 * time.Minute)

	postgresBaseRepository.db = db

	return
}

// Add ...
func (postgresBaseRepository *PostgresBaseRepository) Add(entity interface{}) (err error) {
	err = postgresBaseRepository.db.Create(entity).Error
	if err != nil {
		return
	}
	return
}

// Update ...
func (postgresBaseRepository *PostgresBaseRepository) Update(entity interface{}) (err error) {
	_, err = postgresBaseRepository.GetID(entity)
	if err != nil {
		return
	}
	err = postgresBaseRepository.db.Save(entity).Error
	return
}

// DeleteByID ...
func (postgresBaseRepository *PostgresBaseRepository) DeleteByID(entity interface{}, id uint) (err error) {
	if id == 0 {
		err = postgresBaseRepository.newErrBlankID()
		return
	}

	err = postgresBaseRepository.db.Delete(entity, id).Error
	return
}

// GetAll ...
func (postgresBaseRepository *PostgresBaseRepository) GetAll(entities interface{}) (err error) {
	err = postgresBaseRepository.db.Find(entities).Error
	return
}

// GetByID ...
func (postgresBaseRepository *PostgresBaseRepository) GetByID(entity interface{}, id uint) (err error) {
	err = postgresBaseRepository.db.First(entity, id).Error
	return
}

// SearchOne ...
func (postgresBaseRepository *PostgresBaseRepository) SearchOne(searchEntity interface{}) (err error) {
	err = postgresBaseRepository.db.Where(&searchEntity).First(&searchEntity).Error
	return
}

// SearchOneAdvanced ...
func (postgresBaseRepository *PostgresBaseRepository) SearchOneAdvanced(entity interface{}, where string, args ...interface{}) (err error) {
	err = postgresBaseRepository.db.Where(`where`, args...).First(entity).Error
	return
}
