package events

import (
	"github.com/golang/protobuf/proto"
	bgcm "github.com/paralin/go-dota2/protocol/base_gcmessages"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// InvitationCreated confirms that an invitation has been created.
type InvitationCreated struct {
	bgcm.CMsgInvitationCreated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *InvitationCreated) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg(bgcm.EGCBaseMsg_k_EMsgGCInvitationCreated)
}

// GetEventBody returns the event body.
func (e *InvitationCreated) GetEventBody() proto.Message {
	return &e.CMsgInvitationCreated
}

// GetEventName returns the event name.
func (e *InvitationCreated) GetEventName() string {
	return "InvitationCreated"
}
