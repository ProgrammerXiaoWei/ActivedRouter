package driver

//Memory Driven Based on Map Type
type MapContainer struct {
	data map[string]interface{}
}

//create new map container
func NewMapContainer() *MapContainer {
	return &MapContainer{data: make(map[string]interface{})}
}

func (this *MapContainer) PushKVPair(k, v interface{}) Containerer {
	if key, ok := k.(string); !ok {
		panic("key must be string type!")
	} else {
		this.data[key] = v
	}
	return this
}

func (this *MapContainer) Exist(k interface{}) bool {
	return true
}

func (this *MapContainer) EraseKVPair(k interface{}) Containerer {
	if key, ok := k.(string); !ok {
		panic("key must be string type!")
	} else {
		delete(this.data, key)
	}
	return this
}
func (this *MapContainer) PushKVMaps(maps ...map[string]interface{}) Containerer {
	for _, itemMap := range maps {
		for itemKey, itemValue := range itemMap {
			this.PushKVPair(itemKey, itemValue)
		}
	}
	return this
}
func (this *MapContainer) ResetKVPair(k string, v interface{}) Containerer {
	if _, ok := this.data[k]; ok {
		this.data[k] = v
	}
	return this
}
func (this *MapContainer) ResetOrAddKVPair(k string, v interface{}) Containerer {
	this.data[k] = v
	return this
}

func (this *MapContainer) ResetKVPairs(kvMaps map[string]interface{}) Containerer {
	for k, v := range kvMaps {
		if _, ok := this.data[k]; ok {
			this.data[k] = v
		}
	}
	return this
}

func (this *MapContainer) ResetOrAddKVPairs(kvMaps map[string]interface{}) Containerer {
	for k, v := range kvMaps {
		this.data[k] = v
	}
	return this
}

func (this *MapContainer) GetData() *map[string]interface{} {
	return &this.data
}
