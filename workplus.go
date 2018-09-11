package postoffice_plus

import "github.com/IvoryRaptor/postoffice-plus/message"

type IWorkPlus interface {
	Work(msg *MQMessage) ([]*message.PublishMessage, error)
}
