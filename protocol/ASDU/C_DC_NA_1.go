package ASDU

import (
	"github.com/vedrlabs/go-iec104/protocol/object"
	"github.com/vedrlabs/go-iec104/read_buf"
)

func init() {
	bindASDUStore(TypeCode_C_DC_NA_1, func() ASDUer {
		return New_C_DC_NA_1()
	})
}

func New_C_DC_NA_1() *C_DC_NA_1 {
	return &C_DC_NA_1{asduCap: &asduCap{}}
}

// C_DC_NA_1 双点命令
type C_DC_NA_1 struct {
	*asduCap
	dcoSlice []*object.DCO //双命令
}

func (c *C_DC_NA_1) BindItem(addr uint32, dcs, qu, se byte) {
	c.ioaSlice = append(c.ioaSlice, object.BuildIOA(c.ioaSize, c.ioaOrder, addr))
	c.dcoSlice = append(c.dcoSlice, object.BuildDCO(dcs, qu, se))
}

func (c *C_DC_NA_1) ASDUType() *TypeIdentification {
	return Type_C_DC_NA_1
}

func (c *C_DC_NA_1) Decode(sq byte, bf *read_buf.ReadBuf) error {
	itemSlice, err := c.unifyDecode(c.ASDUType(), sq, bf, c.model()...)
	if err == nil {
		for i := range itemSlice[0] {
			c.dcoSlice = append(c.dcoSlice, itemSlice[0][i].(*object.DCO))
		}
	}
	return err
}

func (c *C_DC_NA_1) Encode(sq byte) (frame []byte, err error) {
	return c.unifyEncode(c.ASDUType(), sq, object.ToObjectors(c.dcoSlice))
}

func (c *C_DC_NA_1) model() []object.Objector {
	return []object.Objector{
		object.NewDCO(),
	}
}

func (c *C_DC_NA_1) ObtainNext() (*object.IOA, *object.DCO) {
	ioa, dco, _ := c.ObtainItem(c.index())
	return ioa, dco
}

func (c *C_DC_NA_1) ObtainItem(index byte) (*object.IOA, *object.DCO, error) {
	if err := c.checkIndex(index); err != nil {
		return nil, nil, err
	}
	return c.ioaSlice[index], c.dcoSlice[index], nil
}
