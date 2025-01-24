package main

import "fmt"

type State interface {
	Publish(context *Document)
	GetStateName() string
}

type Document struct {
	state State
}

func NewDocument() *Document {
	return &Document{
		state: &DraftState{},
	}
}

func (d *Document) SetState(state State) {
	d.state = state
}

func (d *Document) Publish() {
	d.state.Publish(d)
}

func (d *Document) GetStateName() string {
	return d.state.GetStateName()
}

type DraftState struct{}

func (d *DraftState) Publish(context *Document) {
	fmt.Println("Chuyển tài liệu từ Draft sang Moderation.")
	context.SetState(&ModerationState{})
}

func (d *DraftState) GetStateName() string {
	return "Draft"
}

type ModerationState struct{}

func (m *ModerationState) Publish(context *Document) {
	fmt.Println("Chuyển tài liệu từ Moderation sang Published.")
	context.SetState(&PublishedState{})
}

func (m *ModerationState) GetStateName() string {
	return "Moderation"
}

type PublishedState struct{}

func (p *PublishedState) Publish(context *Document) {
	fmt.Println("Tài liệu đã xuất bản. Không có hành động nào được thực hiện.")
}

func (p *PublishedState) GetStateName() string {
	return "Published"
}

func main() {
	doc := NewDocument()
	fmt.Printf("Trạng thái hiện tại: %s\n", doc.GetStateName())
	doc.Publish()
	fmt.Printf("Trạng thái hiện tại: %s\n", doc.GetStateName())
	doc.Publish()
	fmt.Printf("Trạng thái hiện tại: %s\n", doc.GetStateName())
	doc.Publish()
}
