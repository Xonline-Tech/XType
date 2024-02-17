package backend

const PREFIX_OBJECTTYPE string = "type:"

type ObjectType struct {
	typeId   string
	typeName string
}

func CreateObjectType(id string, typeName string) {
	objectType := ObjectType{typeId: id, typeName: typeName}
	SetStruct(PREFIX_OBJECTTYPE+id, objectType)
}

func QueryAll() {
	GetPrefix(PREFIX_OBJECTTYPE)
}
