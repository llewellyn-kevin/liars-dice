// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: protos/game_state.proto

package game

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActionState int32

const (
	ActionState_BIDDING              ActionState = 0
	ActionState_BIDDING_AFTER_REROLL ActionState = 1
	ActionState_CHALLENGING          ActionState = 2
	ActionState_GAME_OVER            ActionState = 3
)

// Enum value maps for ActionState.
var (
	ActionState_name = map[int32]string{
		0: "BIDDING",
		1: "BIDDING_AFTER_REROLL",
		2: "CHALLENGING",
		3: "GAME_OVER",
	}
	ActionState_value = map[string]int32{
		"BIDDING":              0,
		"BIDDING_AFTER_REROLL": 1,
		"CHALLENGING":          2,
		"GAME_OVER":            3,
	}
)

func (x ActionState) Enum() *ActionState {
	p := new(ActionState)
	*p = x
	return p
}

func (x ActionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionState) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_game_state_proto_enumTypes[0].Descriptor()
}

func (ActionState) Type() protoreflect.EnumType {
	return &file_protos_game_state_proto_enumTypes[0]
}

func (x ActionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionState.Descriptor instead.
func (ActionState) EnumDescriptor() ([]byte, []int) {
	return file_protos_game_state_proto_rawDescGZIP(), []int{0}
}

type Hand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []int32 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *Hand) Reset() {
	*x = Hand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_game_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hand) ProtoMessage() {}

func (x *Hand) ProtoReflect() protoreflect.Message {
	mi := &file_protos_game_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hand.ProtoReflect.Descriptor instead.
func (*Hand) Descriptor() ([]byte, []int) {
	return file_protos_game_state_proto_rawDescGZIP(), []int{0}
}

func (x *Hand) GetValues() []int32 {
	if x != nil {
		return x.Values
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Score int32  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	// Types that are assignable to State:
	//
	//	*Player_Bidder
	//	*Player_Challenger
	State isPlayer_State `protobuf_oneof:"state"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_game_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_protos_game_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_protos_game_state_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (m *Player) GetState() isPlayer_State {
	if m != nil {
		return m.State
	}
	return nil
}

func (x *Player) GetBidder() *BidderState {
	if x, ok := x.GetState().(*Player_Bidder); ok {
		return x.Bidder
	}
	return nil
}

func (x *Player) GetChallenger() *ChallengerState {
	if x, ok := x.GetState().(*Player_Challenger); ok {
		return x.Challenger
	}
	return nil
}

type isPlayer_State interface {
	isPlayer_State()
}

type Player_Bidder struct {
	Bidder *BidderState `protobuf:"bytes,3,opt,name=bidder,proto3,oneof"`
}

type Player_Challenger struct {
	Challenger *ChallengerState `protobuf:"bytes,4,opt,name=challenger,proto3,oneof"`
}

func (*Player_Bidder) isPlayer_State() {}

func (*Player_Challenger) isPlayer_State() {}

type BidderState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Self         *Player        `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	Oppoonent    *Player        `protobuf:"bytes,2,opt,name=oppoonent,proto3" json:"oppoonent,omitempty"`
	CurrentHand  *Hand          `protobuf:"bytes,3,opt,name=current_hand,json=currentHand,proto3" json:"current_hand,omitempty"`
	CurrentBid   *BidData       `protobuf:"bytes,4,opt,name=current_bid,json=currentBid,proto3,oneof" json:"current_bid,omitempty"`
	LastBid      *BidData       `protobuf:"bytes,5,opt,name=last_bid,json=lastBid,proto3,oneof" json:"last_bid,omitempty"`
	ValidActions []PlayerAction `protobuf:"varint,6,rep,packed,name=valid_actions,json=validActions,proto3,enum=PlayerAction" json:"valid_actions,omitempty"`
}

func (x *BidderState) Reset() {
	*x = BidderState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_game_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidderState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidderState) ProtoMessage() {}

func (x *BidderState) ProtoReflect() protoreflect.Message {
	mi := &file_protos_game_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidderState.ProtoReflect.Descriptor instead.
func (*BidderState) Descriptor() ([]byte, []int) {
	return file_protos_game_state_proto_rawDescGZIP(), []int{2}
}

func (x *BidderState) GetSelf() *Player {
	if x != nil {
		return x.Self
	}
	return nil
}

func (x *BidderState) GetOppoonent() *Player {
	if x != nil {
		return x.Oppoonent
	}
	return nil
}

func (x *BidderState) GetCurrentHand() *Hand {
	if x != nil {
		return x.CurrentHand
	}
	return nil
}

func (x *BidderState) GetCurrentBid() *BidData {
	if x != nil {
		return x.CurrentBid
	}
	return nil
}

func (x *BidderState) GetLastBid() *BidData {
	if x != nil {
		return x.LastBid
	}
	return nil
}

func (x *BidderState) GetValidActions() []PlayerAction {
	if x != nil {
		return x.ValidActions
	}
	return nil
}

type ChallengerState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Self         *Player        `protobuf:"bytes,1,opt,name=self,proto3" json:"self,omitempty"`
	Oppoonent    *Player        `protobuf:"bytes,2,opt,name=oppoonent,proto3" json:"oppoonent,omitempty"`
	CurrentBid   *BidData       `protobuf:"bytes,3,opt,name=current_bid,json=currentBid,proto3,oneof" json:"current_bid,omitempty"`
	ValidActions []PlayerAction `protobuf:"varint,4,rep,packed,name=valid_actions,json=validActions,proto3,enum=PlayerAction" json:"valid_actions,omitempty"`
}

func (x *ChallengerState) Reset() {
	*x = ChallengerState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_game_state_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengerState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengerState) ProtoMessage() {}

func (x *ChallengerState) ProtoReflect() protoreflect.Message {
	mi := &file_protos_game_state_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengerState.ProtoReflect.Descriptor instead.
func (*ChallengerState) Descriptor() ([]byte, []int) {
	return file_protos_game_state_proto_rawDescGZIP(), []int{3}
}

func (x *ChallengerState) GetSelf() *Player {
	if x != nil {
		return x.Self
	}
	return nil
}

func (x *ChallengerState) GetOppoonent() *Player {
	if x != nil {
		return x.Oppoonent
	}
	return nil
}

func (x *ChallengerState) GetCurrentBid() *BidData {
	if x != nil {
		return x.CurrentBid
	}
	return nil
}

func (x *ChallengerState) GetValidActions() []PlayerAction {
	if x != nil {
		return x.ValidActions
	}
	return nil
}

type GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerOne          *Player     `protobuf:"bytes,1,opt,name=player_one,json=playerOne,proto3" json:"player_one,omitempty"`
	PlayerTwo          *Player     `protobuf:"bytes,2,opt,name=player_two,json=playerTwo,proto3" json:"player_two,omitempty"`
	BidderId           string      `protobuf:"bytes,3,opt,name=bidder_id,json=bidderId,proto3" json:"bidder_id,omitempty"`
	ChallengerId       string      `protobuf:"bytes,4,opt,name=challenger_id,json=challengerId,proto3" json:"challenger_id,omitempty"`
	CurrentHand        *Hand       `protobuf:"bytes,5,opt,name=current_hand,json=currentHand,proto3" json:"current_hand,omitempty"`
	CurrentBid         *BidData    `protobuf:"bytes,6,opt,name=current_bid,json=currentBid,proto3,oneof" json:"current_bid,omitempty"`
	LastBid            *BidData    `protobuf:"bytes,7,opt,name=last_bid,json=lastBid,proto3,oneof" json:"last_bid,omitempty"`
	CurrentActionState ActionState `protobuf:"varint,8,opt,name=current_action_state,json=currentActionState,proto3,enum=ActionState" json:"current_action_state,omitempty"`
	Winner             *string     `protobuf:"bytes,9,opt,name=winner,proto3,oneof" json:"winner,omitempty"`
}

func (x *GameState) Reset() {
	*x = GameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_game_state_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameState) ProtoMessage() {}

func (x *GameState) ProtoReflect() protoreflect.Message {
	mi := &file_protos_game_state_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameState.ProtoReflect.Descriptor instead.
func (*GameState) Descriptor() ([]byte, []int) {
	return file_protos_game_state_proto_rawDescGZIP(), []int{4}
}

func (x *GameState) GetPlayerOne() *Player {
	if x != nil {
		return x.PlayerOne
	}
	return nil
}

func (x *GameState) GetPlayerTwo() *Player {
	if x != nil {
		return x.PlayerTwo
	}
	return nil
}

func (x *GameState) GetBidderId() string {
	if x != nil {
		return x.BidderId
	}
	return ""
}

func (x *GameState) GetChallengerId() string {
	if x != nil {
		return x.ChallengerId
	}
	return ""
}

func (x *GameState) GetCurrentHand() *Hand {
	if x != nil {
		return x.CurrentHand
	}
	return nil
}

func (x *GameState) GetCurrentBid() *BidData {
	if x != nil {
		return x.CurrentBid
	}
	return nil
}

func (x *GameState) GetLastBid() *BidData {
	if x != nil {
		return x.LastBid
	}
	return nil
}

func (x *GameState) GetCurrentActionState() ActionState {
	if x != nil {
		return x.CurrentActionState
	}
	return ActionState_BIDDING
}

func (x *GameState) GetWinner() string {
	if x != nil && x.Winner != nil {
		return *x.Winner
	}
	return ""
}

var File_protos_game_state_proto protoreflect.FileDescriptor

var file_protos_game_state_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x04, 0x48, 0x61, 0x6e, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x62, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x12,
	0x32, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xa6, 0x02, 0x0a,
	0x0b, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x04,
	0x73, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x12, 0x25, 0x0a, 0x09, 0x6f, 0x70, 0x70,
	0x6f, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x09, 0x6f, 0x70, 0x70, 0x6f, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x6e, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x52, 0x0b, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x42, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x42, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42,
	0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x62, 0x69, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x73, 0x65, 0x6c,
	0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x04, 0x73, 0x65, 0x6c, 0x66, 0x12, 0x25, 0x0a, 0x09, 0x6f, 0x70, 0x70, 0x6f, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x09, 0x6f, 0x70, 0x70, 0x6f, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0a,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a,
	0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69,
	0x64, 0x22, 0xa6, 0x03, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x26, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x6e,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x52, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x42, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x42, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x42,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62,
	0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2a, 0x54, 0x0a, 0x0b, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x49, 0x44,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x49, 0x44, 0x44, 0x49, 0x4e,
	0x47, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x52, 0x4f, 0x4c, 0x4c, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x03,
	0x42, 0x07, 0x5a, 0x05, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_game_state_proto_rawDescOnce sync.Once
	file_protos_game_state_proto_rawDescData = file_protos_game_state_proto_rawDesc
)

func file_protos_game_state_proto_rawDescGZIP() []byte {
	file_protos_game_state_proto_rawDescOnce.Do(func() {
		file_protos_game_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_game_state_proto_rawDescData)
	})
	return file_protos_game_state_proto_rawDescData
}

var file_protos_game_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_game_state_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_game_state_proto_goTypes = []interface{}{
	(ActionState)(0),        // 0: ActionState
	(*Hand)(nil),            // 1: Hand
	(*Player)(nil),          // 2: Player
	(*BidderState)(nil),     // 3: BidderState
	(*ChallengerState)(nil), // 4: ChallengerState
	(*GameState)(nil),       // 5: GameState
	(*BidData)(nil),         // 6: BidData
	(PlayerAction)(0),       // 7: PlayerAction
}
var file_protos_game_state_proto_depIdxs = []int32{
	3,  // 0: Player.bidder:type_name -> BidderState
	4,  // 1: Player.challenger:type_name -> ChallengerState
	2,  // 2: BidderState.self:type_name -> Player
	2,  // 3: BidderState.oppoonent:type_name -> Player
	1,  // 4: BidderState.current_hand:type_name -> Hand
	6,  // 5: BidderState.current_bid:type_name -> BidData
	6,  // 6: BidderState.last_bid:type_name -> BidData
	7,  // 7: BidderState.valid_actions:type_name -> PlayerAction
	2,  // 8: ChallengerState.self:type_name -> Player
	2,  // 9: ChallengerState.oppoonent:type_name -> Player
	6,  // 10: ChallengerState.current_bid:type_name -> BidData
	7,  // 11: ChallengerState.valid_actions:type_name -> PlayerAction
	2,  // 12: GameState.player_one:type_name -> Player
	2,  // 13: GameState.player_two:type_name -> Player
	1,  // 14: GameState.current_hand:type_name -> Hand
	6,  // 15: GameState.current_bid:type_name -> BidData
	6,  // 16: GameState.last_bid:type_name -> BidData
	0,  // 17: GameState.current_action_state:type_name -> ActionState
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_protos_game_state_proto_init() }
func file_protos_game_state_proto_init() {
	if File_protos_game_state_proto != nil {
		return
	}
	file_protos_player_actions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_game_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hand); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_game_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_game_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidderState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_game_state_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengerState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_game_state_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_protos_game_state_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Player_Bidder)(nil),
		(*Player_Challenger)(nil),
	}
	file_protos_game_state_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_protos_game_state_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_protos_game_state_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_game_state_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_game_state_proto_goTypes,
		DependencyIndexes: file_protos_game_state_proto_depIdxs,
		EnumInfos:         file_protos_game_state_proto_enumTypes,
		MessageInfos:      file_protos_game_state_proto_msgTypes,
	}.Build()
	File_protos_game_state_proto = out.File
	file_protos_game_state_proto_rawDesc = nil
	file_protos_game_state_proto_goTypes = nil
	file_protos_game_state_proto_depIdxs = nil
}
