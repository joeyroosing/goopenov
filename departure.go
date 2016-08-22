package openov

// Departure ...
type Departure struct {
	DepartureFrom []TimingPointPass
	DepartureTo   []TimingPointPass
}

// NewDeparture creates a new Departure
func NewDeparture() *Departure {
	return &Departure{}
}

// AddTimingPointPass adds a timing point to departure from or departure to destinaton
func (d *Departure) AddTimingPointPass(timingPointPass TimingPointPass, from bool) {
	if from {
		d.DepartureFrom = append(d.DepartureFrom, timingPointPass)
	} else {
		d.DepartureTo = append(d.DepartureTo, timingPointPass)
	}
}
