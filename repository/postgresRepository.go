package repository

import (
	"github.com/YAWAL/TodayInHistory/logging"
	"github.com/YAWAL/TodayInHistory/model"

	"github.com/jinzhu/gorm"
)

const (
	storeLinkError  = "error during storing Link: %v"
	storeEventError = "error during storing Event: %v"
	storeBirthError = "error during storing Birth: %v"
	storeDeathError = "error during storing Death: %v"
)

// PostgresHistoryRepository is a struct for Postgres repository which satisfy HistoryRepository
type PostgresHistoryRepository struct {
	conn *gorm.DB
}

// NewPostgresHistoryRepository creates new PostgresHistoryRepository
func NewPostgresHistoryRepository(conn *gorm.DB) PostgresHistoryRepository {
	return PostgresHistoryRepository{conn: conn}
}

// ProcessHistoryData performs Save/Update all data for HistoryData in PostgreSQL
func (pg PostgresHistoryRepository) ProcessHistoryData(h *model.HistoryData) error {
	events := h.Data.Events
	births := h.Data.Births
	deaths := h.Data.Deaths

	for _, event := range events {
		for _, link := range event.Links {
			if err := pg.conn.Save(&link).Error; err != nil {
				logging.Log.Errorf(storeLinkError, err)
			}
		}
		if err := pg.conn.Save(&event).Error; err != nil {
			logging.Log.Errorf(storeEventError, err)
		}
	}

	for _, birth := range births {
		for _, link := range birth.Links {
			if err := pg.conn.Save(&link).Error; err != nil {
				logging.Log.Errorf(storeLinkError, err)
			}
		}
		if err := pg.conn.Save(&birth).Error; err != nil {
			logging.Log.Errorf(storeBirthError, err)
		}
	}

	for _, death := range deaths {
		for _, link := range death.Links {
			if err := pg.conn.Save(&link).Error; err != nil {
				logging.Log.Errorf(storeLinkError, err)
			}
		}
		if err := pg.conn.Save(&death).Error; err != nil {
			logging.Log.Errorf(storeDeathError, err)
		}
	}
	return nil
}
