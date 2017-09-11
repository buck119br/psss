package topo

// NOTE: THIS FILE WAS PRODUCED BY THE
// ZEBRAPACK CODE GENERATION TOOL (github.com/glycerine/zebrapack)
// DO NOT EDIT

import (
	"github.com/glycerine/zebrapack/msgp"
)

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Addr) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields0zaom = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zaom uint32
	totalEncodedFields0zaom, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zaom := totalEncodedFields0zaom
	missingFieldsLeft0zaom := maxFields0zaom - totalEncodedFields0zaom

	var nextMiss0zaom int = -1
	var found0zaom [maxFields0zaom]bool
	var curField0zaom int

doneWithStruct0zaom:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zaom > 0 || missingFieldsLeft0zaom > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zaom, missingFieldsLeft0zaom, msgp.ShowFound(found0zaom[:]), decodeMsgFieldOrder0zaom)
		if encodedFieldsLeft0zaom > 0 {
			encodedFieldsLeft0zaom--
			curField0zaom, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zaom < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zaom = 0
			}
			for nextMiss0zaom < maxFields0zaom && (found0zaom[nextMiss0zaom] || decodeMsgFieldSkip0zaom[nextMiss0zaom]) {
				nextMiss0zaom++
			}
			if nextMiss0zaom == maxFields0zaom {
				// filled all the empty fields!
				break doneWithStruct0zaom
			}
			missingFieldsLeft0zaom--
			curField0zaom = nextMiss0zaom
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zaom)
		switch curField0zaom {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Host"
			found0zaom[0] = true
			z.Host, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Port"
			found0zaom[1] = true
			z.Port, err = dc.ReadString()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss0zaom != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Addr
var decodeMsgFieldOrder0zaom = []string{"Host", "Port"}

var decodeMsgFieldSkip0zaom = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z Addr) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 2
	}
	var fieldsInUse uint32 = 2
	isempty[0] = (len(z.Host) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.Port) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z Addr) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zynl [2]bool
	fieldsInUse_zuxr := z.fieldsNotEmpty(empty_zynl[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zuxr + 1)
	if err != nil {
		return err
	}

	// runtime struct type identification for 'Addr'
	err = en.Append(0xff)
	if err != nil {
		return err
	}
	err = en.WriteStringFromBytes([]byte{0x41, 0x64, 0x64, 0x72})
	if err != nil {
		return err
	}

	if !empty_zynl[0] {
		// zid 0 for "Host"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Host)
		if err != nil {
			return
		}
	}

	if !empty_zynl[1] {
		// zid 1 for "Port"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteString(z.Port)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z Addr) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [2]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse+1)

	// runtime struct type identification for 'Addr'
	o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72})

	if !empty[0] {
		// zid 0 for "Host"
		o = append(o, 0x0)
		o = msgp.AppendString(o, z.Host)
	}

	if !empty[1] {
		// zid 1 for "Port"
		o = append(o, 0x1)
		o = msgp.AppendString(o, z.Port)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Addr) UnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.UnmarshalMsgWithCfg(bts, cfg)
}
func (z *Addr) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields1zhkk = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zhkk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zhkk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zhkk := totalEncodedFields1zhkk
	missingFieldsLeft1zhkk := maxFields1zhkk - totalEncodedFields1zhkk

	var nextMiss1zhkk int = -1
	var found1zhkk [maxFields1zhkk]bool
	var curField1zhkk int

doneWithStruct1zhkk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zhkk > 0 || missingFieldsLeft1zhkk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zhkk, missingFieldsLeft1zhkk, msgp.ShowFound(found1zhkk[:]), unmarshalMsgFieldOrder1zhkk)
		if encodedFieldsLeft1zhkk > 0 {
			encodedFieldsLeft1zhkk--
			curField1zhkk, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zhkk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zhkk = 0
			}
			for nextMiss1zhkk < maxFields1zhkk && (found1zhkk[nextMiss1zhkk] || unmarshalMsgFieldSkip1zhkk[nextMiss1zhkk]) {
				nextMiss1zhkk++
			}
			if nextMiss1zhkk == maxFields1zhkk {
				// filled all the empty fields!
				break doneWithStruct1zhkk
			}
			missingFieldsLeft1zhkk--
			curField1zhkk = nextMiss1zhkk
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zhkk)
		switch curField1zhkk {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Host"
			found1zhkk[0] = true
			z.Host, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Port"
			found1zhkk[1] = true
			z.Port, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss1zhkk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Addr
var unmarshalMsgFieldOrder1zhkk = []string{"Host", "Port"}

var unmarshalMsgFieldSkip1zhkk = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Addr) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.Host) + 7 + msgp.StringPrefixSize + len(z.Port)
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *AddrState) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields2zofc = 1

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2zofc uint32
	totalEncodedFields2zofc, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zofc := totalEncodedFields2zofc
	missingFieldsLeft2zofc := maxFields2zofc - totalEncodedFields2zofc

	var nextMiss2zofc int = -1
	var found2zofc [maxFields2zofc]bool
	var curField2zofc int

doneWithStruct2zofc:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zofc > 0 || missingFieldsLeft2zofc > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zofc, missingFieldsLeft2zofc, msgp.ShowFound(found2zofc[:]), decodeMsgFieldOrder2zofc)
		if encodedFieldsLeft2zofc > 0 {
			encodedFieldsLeft2zofc--
			curField2zofc, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2zofc < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zofc = 0
			}
			for nextMiss2zofc < maxFields2zofc && (found2zofc[nextMiss2zofc] || decodeMsgFieldSkip2zofc[nextMiss2zofc]) {
				nextMiss2zofc++
			}
			if nextMiss2zofc == maxFields2zofc {
				// filled all the empty fields!
				break doneWithStruct2zofc
			}
			missingFieldsLeft2zofc--
			curField2zofc = nextMiss2zofc
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zofc)
		switch curField2zofc {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Count"
			found2zofc[0] = true
			z.Count, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss2zofc != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of AddrState
var decodeMsgFieldOrder2zofc = []string{"Count"}

var decodeMsgFieldSkip2zofc = []bool{false}

// fieldsNotEmpty supports omitempty tags
func (z AddrState) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 1
	}
	var fieldsInUse uint32 = 1
	isempty[0] = (z.Count == 0) // number, omitempty
	if isempty[0] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z AddrState) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zepf [1]bool
	fieldsInUse_zncq := z.fieldsNotEmpty(empty_zepf[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zncq + 1)
	if err != nil {
		return err
	}

	// runtime struct type identification for 'AddrState'
	err = en.Append(0xff)
	if err != nil {
		return err
	}
	err = en.WriteStringFromBytes([]byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})
	if err != nil {
		return err
	}

	if !empty_zepf[0] {
		// zid 0 for "Count"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteInt(z.Count)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z AddrState) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [1]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse+1)

	// runtime struct type identification for 'AddrState'
	o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

	if !empty[0] {
		// zid 0 for "Count"
		o = append(o, 0x0)
		o = msgp.AppendInt(o, z.Count)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AddrState) UnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.UnmarshalMsgWithCfg(bts, cfg)
}
func (z *AddrState) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields3zckh = 1

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zckh uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zckh, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zckh := totalEncodedFields3zckh
	missingFieldsLeft3zckh := maxFields3zckh - totalEncodedFields3zckh

	var nextMiss3zckh int = -1
	var found3zckh [maxFields3zckh]bool
	var curField3zckh int

doneWithStruct3zckh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zckh > 0 || missingFieldsLeft3zckh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zckh, missingFieldsLeft3zckh, msgp.ShowFound(found3zckh[:]), unmarshalMsgFieldOrder3zckh)
		if encodedFieldsLeft3zckh > 0 {
			encodedFieldsLeft3zckh--
			curField3zckh, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zckh < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zckh = 0
			}
			for nextMiss3zckh < maxFields3zckh && (found3zckh[nextMiss3zckh] || unmarshalMsgFieldSkip3zckh[nextMiss3zckh]) {
				nextMiss3zckh++
			}
			if nextMiss3zckh == maxFields3zckh {
				// filled all the empty fields!
				break doneWithStruct3zckh
			}
			missingFieldsLeft3zckh--
			curField3zckh = nextMiss3zckh
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zckh)
		switch curField3zckh {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Count"
			found3zckh[0] = true
			z.Count, bts, err = nbs.ReadIntBytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss3zckh != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of AddrState
var unmarshalMsgFieldOrder3zckh = []string{"Count"}

var unmarshalMsgFieldSkip3zckh = []bool{false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z AddrState) Msgsize() (s int) {
	s = 1 + 12 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *ProcStat) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields4zuvz = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zuvz uint32
	totalEncodedFields4zuvz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zuvz := totalEncodedFields4zuvz
	missingFieldsLeft4zuvz := maxFields4zuvz - totalEncodedFields4zuvz

	var nextMiss4zuvz int = -1
	var found4zuvz [maxFields4zuvz]bool
	var curField4zuvz int

doneWithStruct4zuvz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zuvz > 0 || missingFieldsLeft4zuvz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zuvz, missingFieldsLeft4zuvz, msgp.ShowFound(found4zuvz[:]), decodeMsgFieldOrder4zuvz)
		if encodedFieldsLeft4zuvz > 0 {
			encodedFieldsLeft4zuvz--
			curField4zuvz, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zuvz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zuvz = 0
			}
			for nextMiss4zuvz < maxFields4zuvz && (found4zuvz[nextMiss4zuvz] || decodeMsgFieldSkip4zuvz[nextMiss4zuvz]) {
				nextMiss4zuvz++
			}
			if nextMiss4zuvz == maxFields4zuvz {
				// filled all the empty fields!
				break doneWithStruct4zuvz
			}
			missingFieldsLeft4zuvz--
			curField4zuvz = nextMiss4zuvz
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zuvz)
		switch curField4zuvz {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "StartTime"
			found4zuvz[0] = true
			z.StartTime, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "State"
			found4zuvz[1] = true
			z.State, err = dc.ReadString()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "LoadAvg"
			found4zuvz[2] = true
			z.LoadAvg, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case 3:
			// zid 3 for "LoadInstant"
			found4zuvz[3] = true
			z.LoadInstant, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case 4:
			// zid 4 for "VmSize"
			found4zuvz[4] = true
			z.VmSize, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case 5:
			// zid 5 for "VmRSS"
			found4zuvz[5] = true
			z.VmRSS, err = dc.ReadUint64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss4zuvz != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of ProcStat
var decodeMsgFieldOrder4zuvz = []string{"StartTime", "State", "LoadAvg", "LoadInstant", "VmSize", "VmRSS"}

var decodeMsgFieldSkip4zuvz = []bool{false, false, false, false, false, false}

// fieldsNotEmpty supports omitempty tags
func (z *ProcStat) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 6
	}
	var fieldsInUse uint32 = 6
	isempty[0] = (z.StartTime == 0) // number, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (len(z.State) == 0) // string, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (z.LoadAvg == 0) // number, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (z.LoadInstant == 0) // number, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (z.VmSize == 0) // number, omitempty
	if isempty[4] {
		fieldsInUse--
	}
	isempty[5] = (z.VmRSS == 0) // number, omitempty
	if isempty[5] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *ProcStat) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zxoi [6]bool
	fieldsInUse_zwdx := z.fieldsNotEmpty(empty_zxoi[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zwdx + 1)
	if err != nil {
		return err
	}

	// runtime struct type identification for 'ProcStat'
	err = en.Append(0xff)
	if err != nil {
		return err
	}
	err = en.WriteStringFromBytes([]byte{0x50, 0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74})
	if err != nil {
		return err
	}

	if !empty_zxoi[0] {
		// zid 0 for "StartTime"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.StartTime)
		if err != nil {
			return
		}
	}

	if !empty_zxoi[1] {
		// zid 1 for "State"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteString(z.State)
		if err != nil {
			return
		}
	}

	if !empty_zxoi[2] {
		// zid 2 for "LoadAvg"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.LoadAvg)
		if err != nil {
			return
		}
	}

	if !empty_zxoi[3] {
		// zid 3 for "LoadInstant"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.LoadInstant)
		if err != nil {
			return
		}
	}

	if !empty_zxoi[4] {
		// zid 4 for "VmSize"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteUint64(z.VmSize)
		if err != nil {
			return
		}
	}

	if !empty_zxoi[5] {
		// zid 5 for "VmRSS"
		err = en.Append(0x5)
		if err != nil {
			return err
		}
		err = en.WriteUint64(z.VmRSS)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ProcStat) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [6]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse+1)

	// runtime struct type identification for 'ProcStat'
	o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x50, 0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74})

	if !empty[0] {
		// zid 0 for "StartTime"
		o = append(o, 0x0)
		o = msgp.AppendInt64(o, z.StartTime)
	}

	if !empty[1] {
		// zid 1 for "State"
		o = append(o, 0x1)
		o = msgp.AppendString(o, z.State)
	}

	if !empty[2] {
		// zid 2 for "LoadAvg"
		o = append(o, 0x2)
		o = msgp.AppendFloat64(o, z.LoadAvg)
	}

	if !empty[3] {
		// zid 3 for "LoadInstant"
		o = append(o, 0x3)
		o = msgp.AppendFloat64(o, z.LoadInstant)
	}

	if !empty[4] {
		// zid 4 for "VmSize"
		o = append(o, 0x4)
		o = msgp.AppendUint64(o, z.VmSize)
	}

	if !empty[5] {
		// zid 5 for "VmRSS"
		o = append(o, 0x5)
		o = msgp.AppendUint64(o, z.VmRSS)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ProcStat) UnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.UnmarshalMsgWithCfg(bts, cfg)
}
func (z *ProcStat) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields5zdxn = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5zdxn uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zdxn, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zdxn := totalEncodedFields5zdxn
	missingFieldsLeft5zdxn := maxFields5zdxn - totalEncodedFields5zdxn

	var nextMiss5zdxn int = -1
	var found5zdxn [maxFields5zdxn]bool
	var curField5zdxn int

doneWithStruct5zdxn:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zdxn > 0 || missingFieldsLeft5zdxn > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zdxn, missingFieldsLeft5zdxn, msgp.ShowFound(found5zdxn[:]), unmarshalMsgFieldOrder5zdxn)
		if encodedFieldsLeft5zdxn > 0 {
			encodedFieldsLeft5zdxn--
			curField5zdxn, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5zdxn < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zdxn = 0
			}
			for nextMiss5zdxn < maxFields5zdxn && (found5zdxn[nextMiss5zdxn] || unmarshalMsgFieldSkip5zdxn[nextMiss5zdxn]) {
				nextMiss5zdxn++
			}
			if nextMiss5zdxn == maxFields5zdxn {
				// filled all the empty fields!
				break doneWithStruct5zdxn
			}
			missingFieldsLeft5zdxn--
			curField5zdxn = nextMiss5zdxn
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zdxn)
		switch curField5zdxn {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "StartTime"
			found5zdxn[0] = true
			z.StartTime, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "State"
			found5zdxn[1] = true
			z.State, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "LoadAvg"
			found5zdxn[2] = true
			z.LoadAvg, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case 3:
			// zid 3 for "LoadInstant"
			found5zdxn[3] = true
			z.LoadInstant, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case 4:
			// zid 4 for "VmSize"
			found5zdxn[4] = true
			z.VmSize, bts, err = nbs.ReadUint64Bytes(bts)

			if err != nil {
				return
			}
		case 5:
			// zid 5 for "VmRSS"
			found5zdxn[5] = true
			z.VmRSS, bts, err = nbs.ReadUint64Bytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss5zdxn != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of ProcStat
var unmarshalMsgFieldOrder5zdxn = []string{"StartTime", "State", "LoadAvg", "LoadInstant", "VmSize", "VmRSS"}

var unmarshalMsgFieldSkip5zdxn = []bool{false, false, false, false, false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ProcStat) Msgsize() (s int) {
	s = 1 + 12 + msgp.Int64Size + 12 + msgp.StringPrefixSize + len(z.State) + 12 + msgp.Float64Size + 12 + msgp.Float64Size + 12 + msgp.Uint64Size + 12 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *ServiceInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields6zcxa = 5

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields6zcxa uint32
	totalEncodedFields6zcxa, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zcxa := totalEncodedFields6zcxa
	missingFieldsLeft6zcxa := maxFields6zcxa - totalEncodedFields6zcxa

	var nextMiss6zcxa int = -1
	var found6zcxa [maxFields6zcxa]bool
	var curField6zcxa int

doneWithStruct6zcxa:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zcxa > 0 || missingFieldsLeft6zcxa > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zcxa, missingFieldsLeft6zcxa, msgp.ShowFound(found6zcxa[:]), decodeMsgFieldOrder6zcxa)
		if encodedFieldsLeft6zcxa > 0 {
			encodedFieldsLeft6zcxa--
			curField6zcxa, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss6zcxa < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zcxa = 0
			}
			for nextMiss6zcxa < maxFields6zcxa && (found6zcxa[nextMiss6zcxa] || decodeMsgFieldSkip6zcxa[nextMiss6zcxa]) {
				nextMiss6zcxa++
			}
			if nextMiss6zcxa == maxFields6zcxa {
				// filled all the empty fields!
				break doneWithStruct6zcxa
			}
			missingFieldsLeft6zcxa--
			curField6zcxa = nextMiss6zcxa
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zcxa)
		switch curField6zcxa {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "ProcsStat"
			found6zcxa[0] = true
			var zogj uint32
			zogj, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.ProcsStat == nil && zogj > 0 {
				z.ProcsStat = make(map[int]ProcStat, zogj)
			} else if len(z.ProcsStat) > 0 {
				for key, _ := range z.ProcsStat {
					delete(z.ProcsStat, key)
				}
			}
			for zogj > 0 {
				zogj--
				var ztfj int
				var zlbr ProcStat
				ztfj, err = dc.ReadInt()
				if err != nil {
					return
				}
				err = zlbr.DecodeMsg(dc)
				if err != nil {
					return
				}
				z.ProcsStat[ztfj] = zlbr
			}
		case 1:
			// zid 1 for "DoListen"
			found6zcxa[1] = true
			z.DoListen, err = dc.ReadBool()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Addrs"
			found6zcxa[2] = true
			var zkuu uint32
			zkuu, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Addrs == nil && zkuu > 0 {
				z.Addrs = make(map[string]AddrState, zkuu)
			} else if len(z.Addrs) > 0 {
				for key, _ := range z.Addrs {
					delete(z.Addrs, key)
				}
			}
			for zkuu > 0 {
				zkuu--
				var zvgr string
				var zjfp AddrState
				zvgr, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields7zkab = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields7zkab uint32
				totalEncodedFields7zkab, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft7zkab := totalEncodedFields7zkab
				missingFieldsLeft7zkab := maxFields7zkab - totalEncodedFields7zkab

				var nextMiss7zkab int = -1
				var found7zkab [maxFields7zkab]bool
				var curField7zkab int

			doneWithStruct7zkab:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft7zkab > 0 || missingFieldsLeft7zkab > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zkab, missingFieldsLeft7zkab, msgp.ShowFound(found7zkab[:]), decodeMsgFieldOrder7zkab)
					if encodedFieldsLeft7zkab > 0 {
						encodedFieldsLeft7zkab--
						curField7zkab, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss7zkab < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss7zkab = 0
						}
						for nextMiss7zkab < maxFields7zkab && (found7zkab[nextMiss7zkab] || decodeMsgFieldSkip7zkab[nextMiss7zkab]) {
							nextMiss7zkab++
						}
						if nextMiss7zkab == maxFields7zkab {
							// filled all the empty fields!
							break doneWithStruct7zkab
						}
						missingFieldsLeft7zkab--
						curField7zkab = nextMiss7zkab
					}
					//fmt.Printf("switching on curField: '%v'\n", curField7zkab)
					switch curField7zkab {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found7zkab[0] = true
						zjfp.Count, err = dc.ReadInt()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
				if nextMiss7zkab != -1 {
					dc.PopAlwaysNil()
				}

				z.Addrs[zvgr] = zjfp
			}
		case 3:
			// zid 3 for "UpStream"
			found6zcxa[3] = true
			var zpyb uint32
			zpyb, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.UpStream == nil && zpyb > 0 {
				z.UpStream = make(map[string]AddrState, zpyb)
			} else if len(z.UpStream) > 0 {
				for key, _ := range z.UpStream {
					delete(z.UpStream, key)
				}
			}
			for zpyb > 0 {
				zpyb--
				var znzi string
				var zysk AddrState
				znzi, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields8zsnn = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields8zsnn uint32
				totalEncodedFields8zsnn, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft8zsnn := totalEncodedFields8zsnn
				missingFieldsLeft8zsnn := maxFields8zsnn - totalEncodedFields8zsnn

				var nextMiss8zsnn int = -1
				var found8zsnn [maxFields8zsnn]bool
				var curField8zsnn int

			doneWithStruct8zsnn:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft8zsnn > 0 || missingFieldsLeft8zsnn > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zsnn, missingFieldsLeft8zsnn, msgp.ShowFound(found8zsnn[:]), decodeMsgFieldOrder8zsnn)
					if encodedFieldsLeft8zsnn > 0 {
						encodedFieldsLeft8zsnn--
						curField8zsnn, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss8zsnn < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss8zsnn = 0
						}
						for nextMiss8zsnn < maxFields8zsnn && (found8zsnn[nextMiss8zsnn] || decodeMsgFieldSkip8zsnn[nextMiss8zsnn]) {
							nextMiss8zsnn++
						}
						if nextMiss8zsnn == maxFields8zsnn {
							// filled all the empty fields!
							break doneWithStruct8zsnn
						}
						missingFieldsLeft8zsnn--
						curField8zsnn = nextMiss8zsnn
					}
					//fmt.Printf("switching on curField: '%v'\n", curField8zsnn)
					switch curField8zsnn {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found8zsnn[0] = true
						zysk.Count, err = dc.ReadInt()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
				if nextMiss8zsnn != -1 {
					dc.PopAlwaysNil()
				}

				z.UpStream[znzi] = zysk
			}
		case 4:
			// zid 4 for "DownStream"
			found6zcxa[4] = true
			var zeby uint32
			zeby, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.DownStream == nil && zeby > 0 {
				z.DownStream = make(map[string]AddrState, zeby)
			} else if len(z.DownStream) > 0 {
				for key, _ := range z.DownStream {
					delete(z.DownStream, key)
				}
			}
			for zeby > 0 {
				zeby--
				var zypj string
				var zkdn AddrState
				zypj, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields9zadc = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields9zadc uint32
				totalEncodedFields9zadc, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft9zadc := totalEncodedFields9zadc
				missingFieldsLeft9zadc := maxFields9zadc - totalEncodedFields9zadc

				var nextMiss9zadc int = -1
				var found9zadc [maxFields9zadc]bool
				var curField9zadc int

			doneWithStruct9zadc:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft9zadc > 0 || missingFieldsLeft9zadc > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zadc, missingFieldsLeft9zadc, msgp.ShowFound(found9zadc[:]), decodeMsgFieldOrder9zadc)
					if encodedFieldsLeft9zadc > 0 {
						encodedFieldsLeft9zadc--
						curField9zadc, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss9zadc < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss9zadc = 0
						}
						for nextMiss9zadc < maxFields9zadc && (found9zadc[nextMiss9zadc] || decodeMsgFieldSkip9zadc[nextMiss9zadc]) {
							nextMiss9zadc++
						}
						if nextMiss9zadc == maxFields9zadc {
							// filled all the empty fields!
							break doneWithStruct9zadc
						}
						missingFieldsLeft9zadc--
						curField9zadc = nextMiss9zadc
					}
					//fmt.Printf("switching on curField: '%v'\n", curField9zadc)
					switch curField9zadc {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found9zadc[0] = true
						zkdn.Count, err = dc.ReadInt()
						if err != nil {
							return
						}
					default:
						err = dc.Skip()
						if err != nil {
							return
						}
					}
				}
				if nextMiss9zadc != -1 {
					dc.PopAlwaysNil()
				}

				z.DownStream[zypj] = zkdn
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss6zcxa != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of ServiceInfo
var decodeMsgFieldOrder6zcxa = []string{"ProcsStat", "DoListen", "Addrs", "UpStream", "DownStream"}

var decodeMsgFieldSkip6zcxa = []bool{false, false, false, false, false}

// fields of AddrState
var decodeMsgFieldOrder7zkab = []string{"Count"}

var decodeMsgFieldSkip7zkab = []bool{false}

// fields of AddrState
var decodeMsgFieldOrder8zsnn = []string{"Count"}

var decodeMsgFieldSkip8zsnn = []bool{false}

// fields of AddrState
var decodeMsgFieldOrder9zadc = []string{"Count"}

var decodeMsgFieldSkip9zadc = []bool{false}

// fieldsNotEmpty supports omitempty tags
func (z *ServiceInfo) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 5
	}
	var fieldsInUse uint32 = 5
	isempty[0] = (len(z.ProcsStat) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (!z.DoListen) // bool, omitempty
	if isempty[1] {
		fieldsInUse--
	}
	isempty[2] = (len(z.Addrs) == 0) // string, omitempty
	if isempty[2] {
		fieldsInUse--
	}
	isempty[3] = (len(z.UpStream) == 0) // string, omitempty
	if isempty[3] {
		fieldsInUse--
	}
	isempty[4] = (len(z.DownStream) == 0) // string, omitempty
	if isempty[4] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *ServiceInfo) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zrzp [5]bool
	fieldsInUse_zzcs := z.fieldsNotEmpty(empty_zrzp[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zzcs + 1)
	if err != nil {
		return err
	}

	// runtime struct type identification for 'ServiceInfo'
	err = en.Append(0xff)
	if err != nil {
		return err
	}
	err = en.WriteStringFromBytes([]byte{0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f})
	if err != nil {
		return err
	}

	if !empty_zrzp[0] {
		// zid 0 for "ProcsStat"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.ProcsStat)))
		if err != nil {
			return
		}
		for ztfj, zlbr := range z.ProcsStat {
			err = en.WriteInt(ztfj)
			if err != nil {
				return
			}
			err = zlbr.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	if !empty_zrzp[1] {
		// zid 1 for "DoListen"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteBool(z.DoListen)
		if err != nil {
			return
		}
	}

	if !empty_zrzp[2] {
		// zid 2 for "Addrs"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Addrs)))
		if err != nil {
			return
		}
		for zvgr, zjfp := range z.Addrs {
			err = en.WriteString(zvgr)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_zqny [1]bool
			fieldsInUse_zkul := zjfp.fieldsNotEmpty(empty_zqny[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zkul + 1)
			if err != nil {
				return err
			}

			// runtime struct type identification for 'AddrState'
			err = en.Append(0xff)
			if err != nil {
				return err
			}
			err = en.WriteStringFromBytes([]byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})
			if err != nil {
				return err
			}

			if !empty_zqny[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zjfp.Count)
				if err != nil {
					return
				}
			}

		}
	}

	if !empty_zrzp[3] {
		// zid 3 for "UpStream"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.UpStream)))
		if err != nil {
			return
		}
		for znzi, zysk := range z.UpStream {
			err = en.WriteString(znzi)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_zoew [1]bool
			fieldsInUse_zieu := zysk.fieldsNotEmpty(empty_zoew[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zieu + 1)
			if err != nil {
				return err
			}

			// runtime struct type identification for 'AddrState'
			err = en.Append(0xff)
			if err != nil {
				return err
			}
			err = en.WriteStringFromBytes([]byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})
			if err != nil {
				return err
			}

			if !empty_zoew[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zysk.Count)
				if err != nil {
					return
				}
			}

		}
	}

	if !empty_zrzp[4] {
		// zid 4 for "DownStream"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.DownStream)))
		if err != nil {
			return
		}
		for zypj, zkdn := range z.DownStream {
			err = en.WriteString(zypj)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_zrrc [1]bool
			fieldsInUse_zxxj := zkdn.fieldsNotEmpty(empty_zrrc[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zxxj + 1)
			if err != nil {
				return err
			}

			// runtime struct type identification for 'AddrState'
			err = en.Append(0xff)
			if err != nil {
				return err
			}
			err = en.WriteStringFromBytes([]byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})
			if err != nil {
				return err
			}

			if !empty_zrrc[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zkdn.Count)
				if err != nil {
					return
				}
			}

		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ServiceInfo) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [5]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse+1)

	// runtime struct type identification for 'ServiceInfo'
	o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f})

	if !empty[0] {
		// zid 0 for "ProcsStat"
		o = append(o, 0x0)
		o = msgp.AppendMapHeader(o, uint32(len(z.ProcsStat)))
		for ztfj, zlbr := range z.ProcsStat {
			o = msgp.AppendInt(o, ztfj)
			o, err = zlbr.MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}

	if !empty[1] {
		// zid 1 for "DoListen"
		o = append(o, 0x1)
		o = msgp.AppendBool(o, z.DoListen)
	}

	if !empty[2] {
		// zid 2 for "Addrs"
		o = append(o, 0x2)
		o = msgp.AppendMapHeader(o, uint32(len(z.Addrs)))
		for zvgr, zjfp := range z.Addrs {
			o = msgp.AppendString(o, zvgr)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zjfp.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zjfp.Count)
			}

		}
	}

	if !empty[3] {
		// zid 3 for "UpStream"
		o = append(o, 0x3)
		o = msgp.AppendMapHeader(o, uint32(len(z.UpStream)))
		for znzi, zysk := range z.UpStream {
			o = msgp.AppendString(o, znzi)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zysk.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zysk.Count)
			}

		}
	}

	if !empty[4] {
		// zid 4 for "DownStream"
		o = append(o, 0x4)
		o = msgp.AppendMapHeader(o, uint32(len(z.DownStream)))
		for zypj, zkdn := range z.DownStream {
			o = msgp.AppendString(o, zypj)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zkdn.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zkdn.Count)
			}

		}
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ServiceInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.UnmarshalMsgWithCfg(bts, cfg)
}
func (z *ServiceInfo) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields10zjfk = 5

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields10zjfk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields10zjfk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft10zjfk := totalEncodedFields10zjfk
	missingFieldsLeft10zjfk := maxFields10zjfk - totalEncodedFields10zjfk

	var nextMiss10zjfk int = -1
	var found10zjfk [maxFields10zjfk]bool
	var curField10zjfk int

doneWithStruct10zjfk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zjfk > 0 || missingFieldsLeft10zjfk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zjfk, missingFieldsLeft10zjfk, msgp.ShowFound(found10zjfk[:]), unmarshalMsgFieldOrder10zjfk)
		if encodedFieldsLeft10zjfk > 0 {
			encodedFieldsLeft10zjfk--
			curField10zjfk, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss10zjfk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss10zjfk = 0
			}
			for nextMiss10zjfk < maxFields10zjfk && (found10zjfk[nextMiss10zjfk] || unmarshalMsgFieldSkip10zjfk[nextMiss10zjfk]) {
				nextMiss10zjfk++
			}
			if nextMiss10zjfk == maxFields10zjfk {
				// filled all the empty fields!
				break doneWithStruct10zjfk
			}
			missingFieldsLeft10zjfk--
			curField10zjfk = nextMiss10zjfk
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zjfk)
		switch curField10zjfk {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "ProcsStat"
			found10zjfk[0] = true
			if nbs.AlwaysNil {
				if len(z.ProcsStat) > 0 {
					for key, _ := range z.ProcsStat {
						delete(z.ProcsStat, key)
					}
				}

			} else {

				var zzip uint32
				zzip, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.ProcsStat == nil && zzip > 0 {
					z.ProcsStat = make(map[int]ProcStat, zzip)
				} else if len(z.ProcsStat) > 0 {
					for key, _ := range z.ProcsStat {
						delete(z.ProcsStat, key)
					}
				}
				for zzip > 0 {
					var ztfj int
					var zlbr ProcStat
					zzip--
					ztfj, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					bts, err = zlbr.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
					z.ProcsStat[ztfj] = zlbr
				}
			}
		case 1:
			// zid 1 for "DoListen"
			found10zjfk[1] = true
			z.DoListen, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Addrs"
			found10zjfk[2] = true
			if nbs.AlwaysNil {
				if len(z.Addrs) > 0 {
					for key, _ := range z.Addrs {
						delete(z.Addrs, key)
					}
				}

			} else {

				var zpfw uint32
				zpfw, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Addrs == nil && zpfw > 0 {
					z.Addrs = make(map[string]AddrState, zpfw)
				} else if len(z.Addrs) > 0 {
					for key, _ := range z.Addrs {
						delete(z.Addrs, key)
					}
				}
				for zpfw > 0 {
					var zvgr string
					var zjfp AddrState
					zpfw--
					zvgr, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields11zssv = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields11zssv uint32
					if !nbs.AlwaysNil {
						totalEncodedFields11zssv, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft11zssv := totalEncodedFields11zssv
					missingFieldsLeft11zssv := maxFields11zssv - totalEncodedFields11zssv

					var nextMiss11zssv int = -1
					var found11zssv [maxFields11zssv]bool
					var curField11zssv int

				doneWithStruct11zssv:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft11zssv > 0 || missingFieldsLeft11zssv > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zssv, missingFieldsLeft11zssv, msgp.ShowFound(found11zssv[:]), unmarshalMsgFieldOrder11zssv)
						if encodedFieldsLeft11zssv > 0 {
							encodedFieldsLeft11zssv--
							curField11zssv, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss11zssv < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss11zssv = 0
							}
							for nextMiss11zssv < maxFields11zssv && (found11zssv[nextMiss11zssv] || unmarshalMsgFieldSkip11zssv[nextMiss11zssv]) {
								nextMiss11zssv++
							}
							if nextMiss11zssv == maxFields11zssv {
								// filled all the empty fields!
								break doneWithStruct11zssv
							}
							missingFieldsLeft11zssv--
							curField11zssv = nextMiss11zssv
						}
						//fmt.Printf("switching on curField: '%v'\n", curField11zssv)
						switch curField11zssv {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found11zssv[0] = true
							zjfp.Count, bts, err = nbs.ReadIntBytes(bts)

							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
					if nextMiss11zssv != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.Addrs[zvgr] = zjfp
				}
			}
		case 3:
			// zid 3 for "UpStream"
			found10zjfk[3] = true
			if nbs.AlwaysNil {
				if len(z.UpStream) > 0 {
					for key, _ := range z.UpStream {
						delete(z.UpStream, key)
					}
				}

			} else {

				var zvvc uint32
				zvvc, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.UpStream == nil && zvvc > 0 {
					z.UpStream = make(map[string]AddrState, zvvc)
				} else if len(z.UpStream) > 0 {
					for key, _ := range z.UpStream {
						delete(z.UpStream, key)
					}
				}
				for zvvc > 0 {
					var znzi string
					var zysk AddrState
					zvvc--
					znzi, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields12zwnt = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields12zwnt uint32
					if !nbs.AlwaysNil {
						totalEncodedFields12zwnt, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft12zwnt := totalEncodedFields12zwnt
					missingFieldsLeft12zwnt := maxFields12zwnt - totalEncodedFields12zwnt

					var nextMiss12zwnt int = -1
					var found12zwnt [maxFields12zwnt]bool
					var curField12zwnt int

				doneWithStruct12zwnt:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft12zwnt > 0 || missingFieldsLeft12zwnt > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zwnt, missingFieldsLeft12zwnt, msgp.ShowFound(found12zwnt[:]), unmarshalMsgFieldOrder12zwnt)
						if encodedFieldsLeft12zwnt > 0 {
							encodedFieldsLeft12zwnt--
							curField12zwnt, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss12zwnt < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss12zwnt = 0
							}
							for nextMiss12zwnt < maxFields12zwnt && (found12zwnt[nextMiss12zwnt] || unmarshalMsgFieldSkip12zwnt[nextMiss12zwnt]) {
								nextMiss12zwnt++
							}
							if nextMiss12zwnt == maxFields12zwnt {
								// filled all the empty fields!
								break doneWithStruct12zwnt
							}
							missingFieldsLeft12zwnt--
							curField12zwnt = nextMiss12zwnt
						}
						//fmt.Printf("switching on curField: '%v'\n", curField12zwnt)
						switch curField12zwnt {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found12zwnt[0] = true
							zysk.Count, bts, err = nbs.ReadIntBytes(bts)

							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
					if nextMiss12zwnt != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.UpStream[znzi] = zysk
				}
			}
		case 4:
			// zid 4 for "DownStream"
			found10zjfk[4] = true
			if nbs.AlwaysNil {
				if len(z.DownStream) > 0 {
					for key, _ := range z.DownStream {
						delete(z.DownStream, key)
					}
				}

			} else {

				var zetl uint32
				zetl, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.DownStream == nil && zetl > 0 {
					z.DownStream = make(map[string]AddrState, zetl)
				} else if len(z.DownStream) > 0 {
					for key, _ := range z.DownStream {
						delete(z.DownStream, key)
					}
				}
				for zetl > 0 {
					var zypj string
					var zkdn AddrState
					zetl--
					zypj, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields13zdkh = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields13zdkh uint32
					if !nbs.AlwaysNil {
						totalEncodedFields13zdkh, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft13zdkh := totalEncodedFields13zdkh
					missingFieldsLeft13zdkh := maxFields13zdkh - totalEncodedFields13zdkh

					var nextMiss13zdkh int = -1
					var found13zdkh [maxFields13zdkh]bool
					var curField13zdkh int

				doneWithStruct13zdkh:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft13zdkh > 0 || missingFieldsLeft13zdkh > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zdkh, missingFieldsLeft13zdkh, msgp.ShowFound(found13zdkh[:]), unmarshalMsgFieldOrder13zdkh)
						if encodedFieldsLeft13zdkh > 0 {
							encodedFieldsLeft13zdkh--
							curField13zdkh, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss13zdkh < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss13zdkh = 0
							}
							for nextMiss13zdkh < maxFields13zdkh && (found13zdkh[nextMiss13zdkh] || unmarshalMsgFieldSkip13zdkh[nextMiss13zdkh]) {
								nextMiss13zdkh++
							}
							if nextMiss13zdkh == maxFields13zdkh {
								// filled all the empty fields!
								break doneWithStruct13zdkh
							}
							missingFieldsLeft13zdkh--
							curField13zdkh = nextMiss13zdkh
						}
						//fmt.Printf("switching on curField: '%v'\n", curField13zdkh)
						switch curField13zdkh {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found13zdkh[0] = true
							zkdn.Count, bts, err = nbs.ReadIntBytes(bts)

							if err != nil {
								return
							}
						default:
							bts, err = msgp.Skip(bts)
							if err != nil {
								return
							}
						}
					}
					if nextMiss13zdkh != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.DownStream[zypj] = zkdn
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss10zjfk != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of ServiceInfo
var unmarshalMsgFieldOrder10zjfk = []string{"ProcsStat", "DoListen", "Addrs", "UpStream", "DownStream"}

var unmarshalMsgFieldSkip10zjfk = []bool{false, false, false, false, false}

// fields of AddrState
var unmarshalMsgFieldOrder11zssv = []string{"Count"}

var unmarshalMsgFieldSkip11zssv = []bool{false}

// fields of AddrState
var unmarshalMsgFieldOrder12zwnt = []string{"Count"}

var unmarshalMsgFieldSkip12zwnt = []bool{false}

// fields of AddrState
var unmarshalMsgFieldOrder13zdkh = []string{"Count"}

var unmarshalMsgFieldSkip13zdkh = []bool{false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ServiceInfo) Msgsize() (s int) {
	s = 1 + 15 + msgp.MapHeaderSize
	if z.ProcsStat != nil {
		for ztfj, zlbr := range z.ProcsStat {
			_ = zlbr
			_ = ztfj
			s += msgp.IntSize + zlbr.Msgsize()
		}
	}
	s += 15 + msgp.BoolSize + 15 + msgp.MapHeaderSize
	if z.Addrs != nil {
		for zvgr, zjfp := range z.Addrs {
			_ = zjfp
			_ = zvgr
			s += msgp.StringPrefixSize + len(zvgr) + 1 + 12 + msgp.IntSize
		}
	}
	s += 15 + msgp.MapHeaderSize
	if z.UpStream != nil {
		for znzi, zysk := range z.UpStream {
			_ = zysk
			_ = znzi
			s += msgp.StringPrefixSize + len(znzi) + 1 + 12 + msgp.IntSize
		}
	}
	s += 15 + msgp.MapHeaderSize
	if z.DownStream != nil {
		for zypj, zkdn := range z.DownStream {
			_ = zkdn
			_ = zypj
			s += msgp.StringPrefixSize + len(zypj) + 1 + 12 + msgp.IntSize
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
// We treat empty fields as if we read a Nil from the wire.
func (z *Topology) DecodeMsg(dc *msgp.Reader) (err error) {
	var sawTopNil bool
	if dc.IsNil() {
		sawTopNil = true
		err = dc.ReadNil()
		if err != nil {
			return
		}
		dc.PushAlwaysNil()
	}

	var field []byte
	_ = field
	const maxFields14zldp = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields14zldp uint32
	totalEncodedFields14zldp, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft14zldp := totalEncodedFields14zldp
	missingFieldsLeft14zldp := maxFields14zldp - totalEncodedFields14zldp

	var nextMiss14zldp int = -1
	var found14zldp [maxFields14zldp]bool
	var curField14zldp int

doneWithStruct14zldp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zldp > 0 || missingFieldsLeft14zldp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zldp, missingFieldsLeft14zldp, msgp.ShowFound(found14zldp[:]), decodeMsgFieldOrder14zldp)
		if encodedFieldsLeft14zldp > 0 {
			encodedFieldsLeft14zldp--
			curField14zldp, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss14zldp < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss14zldp = 0
			}
			for nextMiss14zldp < maxFields14zldp && (found14zldp[nextMiss14zldp] || decodeMsgFieldSkip14zldp[nextMiss14zldp]) {
				nextMiss14zldp++
			}
			if nextMiss14zldp == maxFields14zldp {
				// filled all the empty fields!
				break doneWithStruct14zldp
			}
			missingFieldsLeft14zldp--
			curField14zldp = nextMiss14zldp
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zldp)
		switch curField14zldp {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Services"
			found14zldp[0] = true
			var zqgg uint32
			zqgg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Services == nil && zqgg > 0 {
				z.Services = make(map[string]*ServiceInfo, zqgg)
			} else if len(z.Services) > 0 {
				for key, _ := range z.Services {
					delete(z.Services, key)
				}
			}
			for zqgg > 0 {
				zqgg--
				var ztjf string
				var zmrj *ServiceInfo
				ztjf, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zmrj != nil {
						dc.PushAlwaysNil()
						err = zmrj.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zmrj == nil {
						zmrj = new(ServiceInfo)
					}
					err = zmrj.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.Services[ztjf] = zmrj
			}
		case 1:
			// zid 1 for "Time"
			found14zldp[1] = true
			z.Time, err = dc.ReadInt64()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss14zldp != -1 {
		dc.PopAlwaysNil()
	}

	if sawTopNil {
		dc.PopAlwaysNil()
	}

	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Topology
var decodeMsgFieldOrder14zldp = []string{"Services", "Time"}

var decodeMsgFieldSkip14zldp = []bool{false, false}

// fieldsNotEmpty supports omitempty tags
func (z *Topology) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 2
	}
	var fieldsInUse uint32 = 2
	isempty[0] = (len(z.Services) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (z.Time == 0) // number, omitempty
	if isempty[1] {
		fieldsInUse--
	}

	return fieldsInUse
}

// EncodeMsg implements msgp.Encodable
func (z *Topology) EncodeMsg(en *msgp.Writer) (err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	// honor the omitempty tags
	var empty_zpds [2]bool
	fieldsInUse_zgek := z.fieldsNotEmpty(empty_zpds[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zgek + 1)
	if err != nil {
		return err
	}

	// runtime struct type identification for 'Topology'
	err = en.Append(0xff)
	if err != nil {
		return err
	}
	err = en.WriteStringFromBytes([]byte{0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79})
	if err != nil {
		return err
	}

	if !empty_zpds[0] {
		// zid 0 for "Services"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Services)))
		if err != nil {
			return
		}
		for ztjf, zmrj := range z.Services {
			err = en.WriteString(ztjf)
			if err != nil {
				return
			}
			if zmrj == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zmrj.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zpds[1] {
		// zid 1 for "Time"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteInt64(z.Time)
		if err != nil {
			return
		}
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Topology) MarshalMsg(b []byte) (o []byte, err error) {
	if p, ok := interface{}(z).(msgp.PreSave); ok {
		p.PreSaveHook()
	}

	o = msgp.Require(b, z.Msgsize())

	// honor the omitempty tags
	var empty [2]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse+1)

	// runtime struct type identification for 'Topology'
	o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79})

	if !empty[0] {
		// zid 0 for "Services"
		o = append(o, 0x0)
		o = msgp.AppendMapHeader(o, uint32(len(z.Services)))
		for ztjf, zmrj := range z.Services {
			o = msgp.AppendString(o, ztjf)
			if zmrj == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zmrj.MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty[1] {
		// zid 1 for "Time"
		o = append(o, 0x1)
		o = msgp.AppendInt64(o, z.Time)
	}

	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Topology) UnmarshalMsg(bts []byte) (o []byte, err error) {
	cfg := &msgp.RuntimeConfig{UnsafeZeroCopy: true}
	return z.UnmarshalMsgWithCfg(bts, cfg)
}
func (z *Topology) UnmarshalMsgWithCfg(bts []byte, cfg *msgp.RuntimeConfig) (o []byte, err error) {
	var nbs msgp.NilBitsStack
	nbs.Init(cfg)
	var sawTopNil bool
	if msgp.IsNil(bts) {
		sawTopNil = true
		bts = nbs.PushAlwaysNil(bts[1:])
	}

	var field []byte
	_ = field
	const maxFields15zprm = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields15zprm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields15zprm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft15zprm := totalEncodedFields15zprm
	missingFieldsLeft15zprm := maxFields15zprm - totalEncodedFields15zprm

	var nextMiss15zprm int = -1
	var found15zprm [maxFields15zprm]bool
	var curField15zprm int

doneWithStruct15zprm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft15zprm > 0 || missingFieldsLeft15zprm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zprm, missingFieldsLeft15zprm, msgp.ShowFound(found15zprm[:]), unmarshalMsgFieldOrder15zprm)
		if encodedFieldsLeft15zprm > 0 {
			encodedFieldsLeft15zprm--
			curField15zprm, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss15zprm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss15zprm = 0
			}
			for nextMiss15zprm < maxFields15zprm && (found15zprm[nextMiss15zprm] || unmarshalMsgFieldSkip15zprm[nextMiss15zprm]) {
				nextMiss15zprm++
			}
			if nextMiss15zprm == maxFields15zprm {
				// filled all the empty fields!
				break doneWithStruct15zprm
			}
			missingFieldsLeft15zprm--
			curField15zprm = nextMiss15zprm
		}
		//fmt.Printf("switching on curField: '%v'\n", curField15zprm)
		switch curField15zprm {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Services"
			found15zprm[0] = true
			if nbs.AlwaysNil {
				if len(z.Services) > 0 {
					for key, _ := range z.Services {
						delete(z.Services, key)
					}
				}

			} else {

				var zfnk uint32
				zfnk, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Services == nil && zfnk > 0 {
					z.Services = make(map[string]*ServiceInfo, zfnk)
				} else if len(z.Services) > 0 {
					for key, _ := range z.Services {
						delete(z.Services, key)
					}
				}
				for zfnk > 0 {
					var ztjf string
					var zmrj *ServiceInfo
					zfnk--
					ztjf, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zmrj != nil {
							zmrj.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zmrj {
								zmrj.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zmrj == nil {
								zmrj = new(ServiceInfo)
							}
							bts, err = zmrj.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Services[ztjf] = zmrj
				}
			}
		case 1:
			// zid 1 for "Time"
			found15zprm[1] = true
			z.Time, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss15zprm != -1 {
		bts = nbs.PopAlwaysNil()
	}

	if sawTopNil {
		bts = nbs.PopAlwaysNil()
	}
	o = bts
	if p, ok := interface{}(z).(msgp.PostLoad); ok {
		p.PostLoadHook()
	}

	return
}

// fields of Topology
var unmarshalMsgFieldOrder15zprm = []string{"Services", "Time"}

var unmarshalMsgFieldSkip15zprm = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Topology) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.Services != nil {
		for ztjf, zmrj := range z.Services {
			_ = zmrj
			_ = ztjf
			s += msgp.StringPrefixSize + len(ztjf)
			if zmrj == nil {
				s += msgp.NilSize
			} else {
				s += zmrj.Msgsize()
			}
		}
	}
	s += 12 + msgp.Int64Size
	return
}

// FileSchema_go holds ZebraPack schema from file 'topo/dataExt.go'
type FileSchema_go struct{}

// ZebraSchemaInMsgpack2Format provides the ZebraPack Schema in msgpack2 format, length 2589 bytes
func (FileSchema_go) ZebraSchemaInMsgpack2Format() []byte {
	return []byte{
		0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0xaf, 0x74, 0x6f, 0x70, 0x6f, 0x2f, 0x64, 0x61,
		0x74, 0x61, 0x45, 0x78, 0x74, 0x2e, 0x67, 0x6f, 0xad, 0x53,
		0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
		0x67, 0x65, 0xa4, 0x74, 0x6f, 0x70, 0x6f, 0xad, 0x5a, 0x65,
		0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49,
		0x64, 0xd3, 0x00, 0x00, 0x4d, 0xf6, 0x15, 0x1f, 0xb4, 0x97,
		0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x85, 0xa8,
		0x50, 0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x82, 0xaa,
		0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
		0xa8, 0x50, 0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0xa6,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x96, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x53, 0x74, 0x61, 0x72,
		0x74, 0x54, 0x69, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x53,
		0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x11, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3,
		0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x87,
		0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x53, 0x74,
		0x61, 0x74, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x53, 0x74, 0x61,
		0x74, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64,
		0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa7, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76,
		0x67, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa7, 0x4c, 0x6f, 0x61, 0x64, 0x41,
		0x76, 0x67, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xa7, 0x66, 0x6c, 0x6f, 0x61,
		0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x04, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x04, 0xa3, 0x53, 0x74, 0x72, 0xa7,
		0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x03, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xab, 0x4c, 0x6f, 0x61, 0x64,
		0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xab, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61,
		0x6e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xa7, 0x66, 0x6c, 0x6f, 0x61,
		0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x04, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x04, 0xa3, 0x53, 0x74, 0x72, 0xa7,
		0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x04, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x56, 0x6d, 0x53, 0x69,
		0x7a, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
		0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x56, 0x6d, 0x53, 0x69,
		0x7a, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74,
		0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61,
		0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
		0x76, 0x65, 0x0b, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x0b, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x75,
		0x69, 0x6e, 0x74, 0x36, 0x34, 0x87, 0xa3, 0x5a, 0x69, 0x64,
		0x05, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa5, 0x56, 0x6d, 0x52, 0x53, 0x53, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa5, 0x56, 0x6d, 0x52, 0x53, 0x53, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
		0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x0b, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x0b,
		0xa3, 0x53, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36,
		0x34, 0xab, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
		0x6e, 0x66, 0x6f, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63,
		0x74, 0x4e, 0x61, 0x6d, 0x65, 0xab, 0x53, 0x65, 0x72, 0x76,
		0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0xa6, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x73, 0x95, 0x86, 0xa3, 0x5a, 0x69, 0x64,
		0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa9, 0x50, 0x72, 0x6f, 0x63, 0x73, 0x53,
		0x74, 0x61, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x50, 0x72, 0x6f,
		0x63, 0x73, 0x53, 0x74, 0x61, 0x74, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xb0,
		0x6d, 0x61, 0x70, 0x5b, 0x69, 0x6e, 0x74, 0x5d, 0x50, 0x72,
		0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x0d, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x69,
		0x6e, 0x74, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xa8,
		0x50, 0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x87, 0xa3,
		0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x44, 0x6f, 0x4c,
		0x69, 0x73, 0x74, 0x65, 0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x44,
		0x6f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17,
		0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
		0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74,
		0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa3, 0x5a, 0x69,
		0x64, 0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x41, 0x64, 0x64, 0x72, 0x73,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa5, 0x41, 0x64, 0x64, 0x72, 0x73, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0xb4, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0x5d, 0x41, 0x64, 0x64, 0x72, 0x53, 0x74,
		0x61, 0x74, 0x65, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3,
		0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f,
		0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x19, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x86, 0xa3, 0x5a, 0x69,
		0x64, 0x03, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x55, 0x70, 0x53, 0x74, 0x72,
		0x65, 0x61, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x55, 0x70, 0x53,
		0x74, 0x72, 0x65, 0x61, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xb4, 0x6d,
		0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d,
		0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3,
		0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
		0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74,
		0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52,
		0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x19, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x04, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0xaa, 0x44, 0x6f, 0x77, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61,
		0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x44, 0x6f, 0x77, 0x6e, 0x53,
		0x74, 0x72, 0x65, 0x61, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xb4, 0x6d,
		0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d,
		0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3,
		0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
		0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74,
		0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52,
		0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x19, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0xa8, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
		0x79, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e,
		0x61, 0x6d, 0x65, 0xa8, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
		0x67, 0x79, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x92,
		0x86, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x53,
		0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xa8, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0xb7, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0x5d, 0x2a, 0x53, 0x65, 0x72, 0x76, 0x69,
		0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67,
		0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53,
		0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xab, 0x53,
		0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
		0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x54,
		0x69, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x54, 0x69, 0x6d,
		0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x11, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x11, 0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74,
		0x36, 0x34, 0xa4, 0x41, 0x64, 0x64, 0x72, 0x82, 0xaa, 0x53,
		0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa4,
		0x41, 0x64, 0x64, 0x72, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x73, 0x92, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0xa4, 0x48, 0x6f, 0x73, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x48,
		0x6f, 0x73, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69,
		0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x50, 0x6f, 0x72, 0x74, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa4, 0x50, 0x6f, 0x72, 0x74, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x02, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3,
		0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xa9, 0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
		0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61,
		0x6d, 0x65, 0xa9, 0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61,
		0x74, 0x65, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x91,
		0x87, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x43,
		0x6f, 0x75, 0x6e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x43, 0x6f,
		0x75, 0x6e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x0d, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x0d, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74,
		0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x90,
	}
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 3309 bytes
func (FileSchema_go) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath":"topo/dataExt.go","SourcePackage":"topo","ZebraSchemaId":85719311692951,"Structs":{"ProcStat":{"StructName":"ProcStat","Fields":[{"Zid":0,"FieldGoName":"StartTime","FieldTagName":"StartTime","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}},{"Zid":1,"FieldGoName":"State","FieldTagName":"State","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":2,"FieldGoName":"LoadAvg","FieldTagName":"LoadAvg","FieldTypeStr":"float64","FieldCategory":23,"FieldPrimitive":4,"FieldFullType":{"Kind":4,"Str":"float64"}},{"Zid":3,"FieldGoName":"LoadInstant","FieldTagName":"LoadInstant","FieldTypeStr":"float64","FieldCategory":23,"FieldPrimitive":4,"FieldFullType":{"Kind":4,"Str":"float64"}},{"Zid":4,"FieldGoName":"VmSize","FieldTagName":"VmSize","FieldTypeStr":"uint64","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}},{"Zid":5,"FieldGoName":"VmRSS","FieldTagName":"VmRSS","FieldTypeStr":"uint64","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}}]},"ServiceInfo":{"StructName":"ServiceInfo","Fields":[{"Zid":0,"FieldGoName":"ProcsStat","FieldTagName":"ProcsStat","FieldTypeStr":"map[int]ProcStat","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":13,"Str":"int"},"Range":{"Kind":22,"Str":"ProcStat"}}},{"Zid":1,"FieldGoName":"DoListen","FieldTagName":"DoListen","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"}},{"Zid":2,"FieldGoName":"Addrs","FieldTagName":"Addrs","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}},{"Zid":3,"FieldGoName":"UpStream","FieldTagName":"UpStream","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}},{"Zid":4,"FieldGoName":"DownStream","FieldTagName":"DownStream","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}}]},"Topology":{"StructName":"Topology","Fields":[{"Zid":0,"FieldGoName":"Services","FieldTagName":"Services","FieldTypeStr":"map[string]*ServiceInfo","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"ServiceInfo"}}}},{"Zid":1,"FieldGoName":"Time","FieldTagName":"Time","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"Addr":{"StructName":"Addr","Fields":[{"Zid":0,"FieldGoName":"Host","FieldTagName":"Host","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":1,"FieldGoName":"Port","FieldTagName":"Port","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}}]},"AddrState":{"StructName":"AddrState","Fields":[{"Zid":0,"FieldGoName":"Count","FieldTagName":"Count","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}}]}},"Imports":[]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 8948 bytes
func (FileSchema_go) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "topo/dataExt.go",
    "SourcePackage": "topo",
    "ZebraSchemaId": 85719311692951,
    "Structs": {
        "ProcStat": {
            "StructName": "ProcStat",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "StartTime",
                    "FieldTagName": "StartTime",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "State",
                    "FieldTagName": "State",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "LoadAvg",
                    "FieldTagName": "LoadAvg",
                    "FieldTypeStr": "float64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 4,
                    "FieldFullType": {
                        "Kind": 4,
                        "Str": "float64"
                    }
                },
                {
                    "Zid": 3,
                    "FieldGoName": "LoadInstant",
                    "FieldTagName": "LoadInstant",
                    "FieldTypeStr": "float64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 4,
                    "FieldFullType": {
                        "Kind": 4,
                        "Str": "float64"
                    }
                },
                {
                    "Zid": 4,
                    "FieldGoName": "VmSize",
                    "FieldTagName": "VmSize",
                    "FieldTypeStr": "uint64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 11,
                    "FieldFullType": {
                        "Kind": 11,
                        "Str": "uint64"
                    }
                },
                {
                    "Zid": 5,
                    "FieldGoName": "VmRSS",
                    "FieldTagName": "VmRSS",
                    "FieldTypeStr": "uint64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 11,
                    "FieldFullType": {
                        "Kind": 11,
                        "Str": "uint64"
                    }
                }
            ]
        },
        "ServiceInfo": {
            "StructName": "ServiceInfo",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "ProcsStat",
                    "FieldTagName": "ProcsStat",
                    "FieldTypeStr": "map[int]ProcStat",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 13,
                            "Str": "int"
                        },
                        "Range": {
                            "Kind": 22,
                            "Str": "ProcStat"
                        }
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "DoListen",
                    "FieldTagName": "DoListen",
                    "FieldTypeStr": "bool",
                    "FieldCategory": 23,
                    "FieldPrimitive": 18,
                    "FieldFullType": {
                        "Kind": 18,
                        "Str": "bool"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "Addrs",
                    "FieldTagName": "Addrs",
                    "FieldTypeStr": "map[string]AddrState",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 25,
                            "Str": "Struct"
                        }
                    }
                },
                {
                    "Zid": 3,
                    "FieldGoName": "UpStream",
                    "FieldTagName": "UpStream",
                    "FieldTypeStr": "map[string]AddrState",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 25,
                            "Str": "Struct"
                        }
                    }
                },
                {
                    "Zid": 4,
                    "FieldGoName": "DownStream",
                    "FieldTagName": "DownStream",
                    "FieldTypeStr": "map[string]AddrState",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 25,
                            "Str": "Struct"
                        }
                    }
                }
            ]
        },
        "Topology": {
            "StructName": "Topology",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "Services",
                    "FieldTagName": "Services",
                    "FieldTypeStr": "map[string]*ServiceInfo",
                    "FieldCategory": 24,
                    "FieldFullType": {
                        "Kind": 24,
                        "Str": "Map",
                        "Domain": {
                            "Kind": 2,
                            "Str": "string"
                        },
                        "Range": {
                            "Kind": 28,
                            "Str": "Pointer",
                            "Domain": {
                                "Kind": 22,
                                "Str": "ServiceInfo"
                            }
                        }
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "Time",
                    "FieldTagName": "Time",
                    "FieldTypeStr": "int64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 17,
                    "FieldFullType": {
                        "Kind": 17,
                        "Str": "int64"
                    }
                }
            ]
        },
        "Addr": {
            "StructName": "Addr",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "Host",
                    "FieldTagName": "Host",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "Port",
                    "FieldTagName": "Port",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                }
            ]
        },
        "AddrState": {
            "StructName": "AddrState",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "Count",
                    "FieldTagName": "Count",
                    "FieldTypeStr": "int",
                    "FieldCategory": 23,
                    "FieldPrimitive": 13,
                    "FieldFullType": {
                        "Kind": 13,
                        "Str": "int"
                    }
                }
            ]
        }
    },
    "Imports": []
}`)
}
