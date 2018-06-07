package message

/*Location Location*/
type Location struct {
	Message
	LocationX string `xml:"location_x"`
	LocationY string `xml:"location_y"`
	Scale     string `xml:"scale"`
	Label     string `xml:"label"`
}
