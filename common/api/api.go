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
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

func NewIncidentEvent(title string, description string, time time.Time) IncidentEvent {
	return IncidentEvent{
		Title:       title,
		Description: description,
		Time:        time,
	}
}

type Incident struct {
	Title                   string             `json:"title"`
	Components              []string           `gorm:"column:components;type:jsonb" json:"components"`
	Events                  IncidentEventArray `gorm:"column:events;type:jsonb" json:"events"`
	StartTime               time.Time          `gorm:"column:start_time;secondarykey" json:"startTime"`
	EndTime                 *time.Time         `gorm:"column:end_time;secondarykey" json:"endTime"`
	Description             *string            `json:"column:description;description"`
	DeepLink                string             `gorm:"column:deep_link;primarykey" json:"deepLink"`
	Impact                  Impact             `gorm:"column:impact;secondarykey" json:"impact"`
	StatusPageUrl           string             `gorm:"column:status_page_url;secondarykey" json:"statusPageUrl"`
	NotificationJobsStarted bool               `gorm:"column:notification_jobs_started;secondarykey" json:"notificationJobsStarted"`
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
	Name string `gorm:"secondarykey" json:"name"`
	URL  string `gorm:"primarykey" json:"url"`
	// Used to determine if we should run a scrape for this status page
	LastHistoricallyScraped time.Time `json:"lastHistoricallyScraped"`
	LastCurrentlyScraped    time.Time `json:"lastCurrentlyScraped"`
	// IsIndexed is used to determine if the status page has ever been indexed in the search engine successfully
	IsIndexed bool `json:"isIndexed"`
}

func NewStatusPage(name string, url string) StatusPage {
	return StatusPage{
		Name:                    name,
		URL:                     url,
		LastHistoricallyScraped: time.Time{},
		LastCurrentlyScraped:    time.Time{},
	}
}
