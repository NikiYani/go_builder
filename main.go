package main

// Builder

type ComputerBuilderI interface {
	CPU(val string) ComputerBuilderI
	GPU(val string) ComputerBuilderI
	RAM(val int) ComputerBuilderI
	MB(val string) ComputerBuilderI
	HDD(val int) ComputerBuilderI

	Build() Computer
}

type Computer struct {
	cpu string
	gpu string
	ram int
	mb  string
	hdd int
}

func NewComputerBuilder() ComputerBuilderI {
	return Computer{}
}

func (b Computer) CPU(val string) ComputerBuilderI {
	b.cpu = val
	return b
}

func (b Computer) GPU(val string) ComputerBuilderI {
	b.gpu = val
	return b
}

func (b Computer) RAM(val int) ComputerBuilderI {
	b.ram = val
	return b
}

func (b Computer) MB(val string) ComputerBuilderI {
	b.mb = val
	return b
}

func (b Computer) HDD(val int) ComputerBuilderI {
	b.hdd = val
	return b
}

func (b Computer) Build() Computer {
	return Computer{
		cpu: b.cpu,
		gpu: b.gpu,
		ram: b.ram,
		mb:  b.mb,
		hdd: b.hdd,
	}
}

func NewOfficeComputerBuilder() ComputerBuilderI {
	return officeComputerBuilder{}.CPU("Intel celeron").GPU("None").RAM(2).MB("Asrock").HDD(1)
}

type officeComputerBuilder struct {
	Computer
}

func (b *officeComputerBuilder) Build() Computer {
	return Computer{
		cpu: b.cpu,
		gpu: b.gpu,
		ram: b.ram,
		mb:  b.mb,
		hdd: b.hdd,
	}
}
