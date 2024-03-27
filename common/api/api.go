package api

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Impact string

const (
	ImpactMinor       Impact = "minor"
	ImpactMajor       Impact = "major"
	ImpactCritical    Impact = "critical"
	ImpactMaintenance Impact = "maintenance"
	ImpactNone        Impact = "none"
)

type IncidentEventArray []IncidentEvent

func (sla *IncidentEventArray) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &sla)
}

func (sla IncidentEventArray) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

type IncidentEvent struct {
	Title       string
	Description string
	Time        time.Time
}

func NewIncidentEvent(title string, description string, time time.Time) IncidentEvent {
	return IncidentEvent{
		Title:       title,
		Description: description,
		Time:        time,
	}
}

type Incident struct {
	Title         string
	Components    []string           `gorm:"column:components;type:jsonb"`
	Events        IncidentEventArray `gorm:"column:events;type:jsonb"`
	StartTime     time.Time          `gorm:"secondarykey"`
	EndTime       *time.Time         `gorm:"secondarykey"`
	Description   *string
	DeepLink      string `gorm:"primarykey"`
	Impact        Impact `gorm:"secondarykey"`
	StatusPageUrl string `gorm:"secondarykey"`
}

func NewIncident(title string, components []string, events []IncidentEvent, startTime time.Time, endTime *time.Time, description *string, deepLink string, impact Impact, statusPageUrl string) Incident {
	return Incident{
		Title:         title,
		Components:    components,
		Events:        events,
		StartTime:     startTime,
		EndTime:       endTime,
		Description:   description,
		DeepLink:      deepLink,
		Impact:        impact,
		StatusPageUrl: statusPageUrl,
	}
}

type StatusPage struct {
	Name string `gorm:"secondarykey"`
	URL  string `gorm:"primarykey"`
	// Used to determine if we should run a scrape for this status page
	LastHistoricallyScraped time.Time
	LastCurrentlyScraped    time.Time
}

func NewStatusPage(name string, url string) StatusPage {
	return StatusPage{
		Name:                    name,
		URL:                     url,
		LastHistoricallyScraped: time.Time{},
		LastCurrentlyScraped:    time.Time{},
	}
}
