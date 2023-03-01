package builder

type Phone struct {
	model      string
	name       string
	os         string
	ramInGB    int8
	screenSize float32
}

type PhoneBuilder struct {
	phone Phone
}

func (b *PhoneBuilder) SetModel(model string) *PhoneBuilder {
	b.phone.model = model
	return b
}

func (b *PhoneBuilder) SetName(name string) *PhoneBuilder {
	b.phone.name = name
	return b
}

func (b *PhoneBuilder) SetOS(os string) *PhoneBuilder {
	b.phone.os = os
	return b
}

func (b *PhoneBuilder) SetRAM(ramInGB int8) *PhoneBuilder {
	b.phone.ramInGB = ramInGB
	return b
}

func (b *PhoneBuilder) SetScreenSize(screenSize float32) *PhoneBuilder {
	b.phone.screenSize = screenSize
	return b
}

func (b *PhoneBuilder) Build() *Phone {
	return &b.phone
}
