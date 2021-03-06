package uavtalk

import log "github.com/Sirupsen/logrus"

func CreateObjectRequest(name string, index int) *Packet {
	definition, err := AllDefinitions.GetDefinitionForName(name)
	if err != nil {
		log.Fatal(err)
	}
	packet := NewPacket(definition, ObjectRequest, uint16(index), map[string]interface{}{})
	return packet
}

func CreateObjectSetter(name string, index int, data map[string]interface{}) *Packet {
	definition, err := AllDefinitions.GetDefinitionForName(name)
	if err != nil {
		log.Fatal(err)
	}
	packet := NewPacket(definition, ObjectCmd, uint16(index), data)
	return packet
}

func CreateGCSTelemetryStatsObjectPacket(status string) Packet {
	definition, err := AllDefinitions.GetDefinitionForName("GCSTelemetryStats")
	if err != nil {
		log.Fatal(err)
	}
	packet := NewPacket(definition, ObjectCmd, 0, map[string]interface{}{
		"Status":     status,
		"TxDataRate": float64(0),
		"RxDataRate": float64(0),
		"TxFailures": float64(0),
		"RxFailures": float64(0),
		"TxRetries":  float64(0),
	})
	return *packet
}

func CreateSessionManagingRequest() Packet {
	definition, err := AllDefinitions.GetDefinitionForName("SessionManaging")
	if err != nil {
		log.Fatal(err)
	}
	packet := NewPacket(definition, ObjectRequest, 0, map[string]interface{}{})
	return *packet
}

func CreateSessionManagingPacket(sessionID uint16, objectOfInterestIndex uint8) Packet {
	definition, err := AllDefinitions.GetDefinitionForName("SessionManaging")
	if err != nil {
		log.Fatal(err)
	}
	packet := NewPacket(definition, ObjectCmd, 0, map[string]interface{}{
		"SessionID":             float64(sessionID),
		"ObjectID":              float64(0),
		"ObjectInstances":       float64(0),
		"NumberOfObjects":       float64(0),
		"ObjectOfInterestIndex": float64(objectOfInterestIndex),
	})
	return *packet
}

func CreatePersistObject(definition *Definition, instanceID uint16) Packet {
	objectPersistenceDefinition, err := AllDefinitions.GetDefinitionForName("ObjectPersistence")
	if err != nil {
		log.Fatal(err)
	}
	packet := NewPacket(objectPersistenceDefinition, ObjectCmdWithAck, instanceID, map[string]interface{}{
		"ObjectID":   float64(definition.ObjectID),
		"InstanceID": float64(instanceID),
		"Selection":  "SingleObject",
		"Operation":  "Save",
	})
	return *packet
}

func CreatePacketAck(definition *Definition) Packet {
	packet := NewPacket(definition, ObjectAck, 0, map[string]interface{}{})
	return *packet
}
