package openov

import "time"

// TimingPointPass is the information of a transport "passing" a timing point
type TimingPointPass struct {
	DestinationName     string    `mapstructure:"DestinationName50"`
	DataOwnerCode       string    `mapstructure:"DataOwnerCode"`
	OperatorCode        string    `mapstructure:"OperatorCode"`
	TransportType       string    `mapstructure:"TransportType"`
	Latitude            float64   `mapstructure:"Latitude"`
	Longitude           float64   `mapstructure:"Longitude"`
	JourneyNumber       int       `mapstructure:"JourneyNumber"`
	TimingPointCode     string    `mapstructure:"TimingPointCode"`
	TimingPointName     string    `mapstructure:"TimingPointName"`
	TimingPointTown     string    `mapstructure:"TimingPointTown"`
	LinePlanningNumber  string    `mapstructure:"LinePlanningNumber"`
	TargetArrivalTime   time.Time `mapstructure:"TargetArrivalTime"`
	TargetDepartureTime time.Time `mapstructure:"TargetDepartureTime"`
	ExpectedArrivalTime time.Time `mapstructure:"ExpectedArrivalTime"`
}
