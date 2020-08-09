package record

//New returns a new empty record instance
func New(layout string, data interface{}) Record {
	return Record{
		ID:            data.(map[string]interface{})["recordId"].(string),
		Layout:        layout,
		StagedChanges: make(map[string]interface{}),
		fieldData:     data.(map[string]interface{})["fieldData"].(map[string]interface{}),
	}
}