package domain

type Devices []Device

func (ds *Devices) AddItem(d Device) {

}

type Device struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	AdminState     string    `json:"adminState"`
	OperatingState string    `json:"operatingState"`
	Labels         []string  `json:"labels"`
	Commands       []Command `json:"commands"`
}
