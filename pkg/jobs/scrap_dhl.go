package jobs

func (j *Jobs) ScrapDHL() {
	orders, err := j.services.GetAll()

	if err != nil {
		return
	}

	for _, order := range orders {
		// Get last location by tracking number
		lastLocation := j.scraper.GetLastLocation(order.TrackingNumber)

		if lastLocation != order.LastLocation {
			println("New location found: ", lastLocation)
			j.services.UpdateLastLocation(order.TrackingNumber, lastLocation)
		}
	}
}
