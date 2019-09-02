package database

import (
	"github.com/YAWAL/TodayInHistory/model"
)

// HistoryRepository is an interface for different DB storages to proceed with history data
type HistoryRepository interface {
	ProcessHistoryData(hd *model.HistoryData) error
}
