package session

// Hooks constants
const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

type BeforeQueryInterface interface {
	BeforeQuery(s *Session) error
}

type AfterQueryInterface interface {
	AfterQuery(s *Session) error
}

type BeforeInsertInterface interface {
	BeforeInsert(s *Session) error
}

type AfterInsertInterface interface {
	AfterInsert(s *Session) error
}

type BerfoeUpdateInterface interface {
	BeforeUpdate(s *Session) error
}

type AfterUpdateInterface interface {
	AfterUpdate(s *Session) error
}

type BeforeDeleteInterface interface {
	BeforeDelete(s *Session) error
}

type AfterDeleteInterface interface {
	AfterDelete(s *Session) error
}

func (s *Session) CallMethod(method string, value interface{}) {
	switch method {
	case BeforeQuery:
		if i, ok := value.(BeforeQueryInterface); ok {
			i.BeforeQuery(s)
		}
	case AfterQuery:
		if i, ok := value.(AfterQueryInterface); ok {
			i.AfterQuery(s)
		}
	case BeforeInsert:
		if i, ok := value.(BeforeInsertInterface); ok {
			i.BeforeInsert(s)
		}
	case AfterInsert:
		if i, ok := value.(AfterInsertInterface); ok {
			i.AfterInsert(s)
		}
	case BeforeUpdate:
		if i, ok := value.(BerfoeUpdateInterface); ok {
			i.BeforeUpdate(s)
		}
	case AfterUpdate:
		if i, ok := value.(AfterUpdateInterface); ok {
			i.AfterUpdate(s)
		}
	case BeforeDelete:
		if i, ok := value.(BeforeDeleteInterface); ok {
			i.BeforeDelete(s)
		}
	case AfterDelete:
		if i, ok := value.(AfterDeleteInterface); ok {
			i.AfterDelete(s)
		}
	default:
		panic("unsupported callback function called " + method)
	}
}
