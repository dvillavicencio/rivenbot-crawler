package dto 

type ManifestObject struct {
	Mode                      int               `json:"directActivityModeType"` // Using *int to handle nil values
	DisplayProperties         DisplayProperties `json:"displayProperties"`
	OriginalDisplayProperties DisplayProperties `json:"originalDisplayProperties"`
	ReleaseIcon               string            `json:"releaseIcon"`
	ReleaseTime               int               `json:"releaseTime"`
	ItemType                  int               `json:"itemType"`
	EquippingBlock            EquippingBlock    `json:"equippingBlock"`
	// Other fields omitted for brevity
}

type DisplayProperties struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	HasIcon     bool   `json:"hasIcon"`
}

type EquippingBlock struct {
	EquipmentSlotTypeHash int `json:"equipmentSlotTypeHash"`
	AmmoType              int `json:"ammoType"`
}
