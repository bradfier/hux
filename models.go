package hux

import (
	"time"
)

// Go models for data returned via the Huxley endpoint proxy

type CallingPoint struct {
	AT           string `json:"at"`
	AdhocAlerts  string `json:"adhocAlerts"`
	CRS          string `json:"crs"`
	DetatchFront bool   `json:"detachFront"`
	ET           string `json:"et"`
	IsCancelled  bool   `json:"isCancelled"`
	Length       int    `json:"length"`
	LocationName string `json:"locationName"`
	ST           string `json:"st"`
}

type Station struct {
	AssocIsCancelled bool   `json:"assocIsCancelled"`
	CRS              string `json:"crs"`
	FutureChangeTo   string `json:"futureChangeTo"`
	LocationName     string `json:"locationName"`
	Via              string `json:"via"`
}

type TrainService struct {
	AdhocAlerts         string    `json:"adhocAlerts"`
	Origin              []Station `json:"origin"`
	Destination         []Station `json:"destination"`
	CurrentOrigins      []Station `json:"currentOrigins"`
	CurrentDestinations []Station `json:"currentDestinations"`
	ETA                 string    `json:"eta"`
	ETD                 string    `json:"eta"`
	IsCircularRoute     bool      `json:"isCircularRoute"`
	OperatorCode        string    `json:"operatorCode"`
	Operator            string    `json:"operator"`
	Platform            string    `json:"platform"`
	ServiceID           string    `json:"serviceID"`
	STA                 string    `json:"sta"`
	STD                 string    `json:"std"`
}

type ServiceDetails struct {
	GeneratedAt             time.Time                 `json:"generatedAt"`
	ATA                     string                    `json:"ata"`
	ATD                     string                    `json:"atd"`
	AdhocAlerts             string                    `json:"adhocAlerts"`
	CRS                     string                    `json:"crs"`
	CancelReason            string                    `json:"cancelReason"`
	DelayReason             string                    `json:"delayReason"`
	IsCancelled             bool                      `json:"isCancelled"`
	LocationName            string                    `json:"locationName"`
	Operator                string                    `json:"operator"`
	OperatorCode            string                    `json:"operatorCode"`
	ServiceType             int                       `json:"serviceType"`
	OverdueMessage          string                    `json:"overdueMessage"`
	Length                  int                       `json:"length"`
	DetachFront             bool                      `json:"detachFront"`
	IsReverseFormation      bool                      `json:"isReverseFormation"`
	Platform                string                    `json:"platform"`
	STA                     string                    `json:"sta"`
	ETA                     string                    `json:"eta"`
	STD                     string                    `json:"std"`
	ETD                     string                    `json:"etd"`
	PreviousCallingPoints   []PreviousCallingPoints   `json:"previousCallingPoints"`
	SubsequentCallingPoints []SubsequentCallingPoints `json:"subsequentCallingPoints"`
}

type PreviousCallingPoints struct {
	AssocIsCancelled      bool           `json:"serviceType"`
	CallingPoint          []CallingPoint `json:"callingPoint"`
	ServiceChangeRequired bool           `json:"serviceChangeRequired"`
	ServiceType           int            `json:"serviceType"`
}

type SubsequentCallingPoints struct {
	AssocIsCancelled      bool           `json:"serviceType"`
	CallingPoint          []CallingPoint `json:"callingPoint"`
	ServiceChangeRequired bool           `json:"serviceChangeRequired"`
	ServiceType           int            `json:"serviceType"`
}

type CRSStationCode struct {
	StationName string `json:"stationName"`
	CRSCode     string `json:"crsCode"`
}

type BoardResponse struct {
	TrainServices []TrainService `json:"trainServices"`
	busServices   bool           `json:"busServices"`
	CRS           string         `json:"crs"`
}
