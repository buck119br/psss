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
	const maxFields0zaqg = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zaqg uint32
	totalEncodedFields0zaqg, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zaqg := totalEncodedFields0zaqg
	missingFieldsLeft0zaqg := maxFields0zaqg - totalEncodedFields0zaqg

	var nextMiss0zaqg int = -1
	var found0zaqg [maxFields0zaqg]bool
	var curField0zaqg int

doneWithStruct0zaqg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zaqg > 0 || missingFieldsLeft0zaqg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zaqg, missingFieldsLeft0zaqg, msgp.ShowFound(found0zaqg[:]), decodeMsgFieldOrder0zaqg)
		if encodedFieldsLeft0zaqg > 0 {
			encodedFieldsLeft0zaqg--
			curField0zaqg, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zaqg < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zaqg = 0
			}
			for nextMiss0zaqg < maxFields0zaqg && (found0zaqg[nextMiss0zaqg] || decodeMsgFieldSkip0zaqg[nextMiss0zaqg]) {
				nextMiss0zaqg++
			}
			if nextMiss0zaqg == maxFields0zaqg {
				// filled all the empty fields!
				break doneWithStruct0zaqg
			}
			missingFieldsLeft0zaqg--
			curField0zaqg = nextMiss0zaqg
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zaqg)
		switch curField0zaqg {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Host"
			found0zaqg[0] = true
			z.Host, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Port"
			found0zaqg[1] = true
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
	if nextMiss0zaqg != -1 {
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
var decodeMsgFieldOrder0zaqg = []string{"Host", "Port"}

var decodeMsgFieldSkip0zaqg = []bool{false, false}

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
	var empty_zevi [2]bool
	fieldsInUse_zomn := z.fieldsNotEmpty(empty_zevi[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zomn + 1)
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

	if !empty_zevi[0] {
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

	if !empty_zevi[1] {
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
	const maxFields1zgkm = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zgkm uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zgkm, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zgkm := totalEncodedFields1zgkm
	missingFieldsLeft1zgkm := maxFields1zgkm - totalEncodedFields1zgkm

	var nextMiss1zgkm int = -1
	var found1zgkm [maxFields1zgkm]bool
	var curField1zgkm int

doneWithStruct1zgkm:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zgkm > 0 || missingFieldsLeft1zgkm > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zgkm, missingFieldsLeft1zgkm, msgp.ShowFound(found1zgkm[:]), unmarshalMsgFieldOrder1zgkm)
		if encodedFieldsLeft1zgkm > 0 {
			encodedFieldsLeft1zgkm--
			curField1zgkm, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zgkm < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zgkm = 0
			}
			for nextMiss1zgkm < maxFields1zgkm && (found1zgkm[nextMiss1zgkm] || unmarshalMsgFieldSkip1zgkm[nextMiss1zgkm]) {
				nextMiss1zgkm++
			}
			if nextMiss1zgkm == maxFields1zgkm {
				// filled all the empty fields!
				break doneWithStruct1zgkm
			}
			missingFieldsLeft1zgkm--
			curField1zgkm = nextMiss1zgkm
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zgkm)
		switch curField1zgkm {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Host"
			found1zgkm[0] = true
			z.Host, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Port"
			found1zgkm[1] = true
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
	if nextMiss1zgkm != -1 {
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
var unmarshalMsgFieldOrder1zgkm = []string{"Host", "Port"}

var unmarshalMsgFieldSkip1zgkm = []bool{false, false}

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
	const maxFields2zkyw = 1

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2zkyw uint32
	totalEncodedFields2zkyw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zkyw := totalEncodedFields2zkyw
	missingFieldsLeft2zkyw := maxFields2zkyw - totalEncodedFields2zkyw

	var nextMiss2zkyw int = -1
	var found2zkyw [maxFields2zkyw]bool
	var curField2zkyw int

doneWithStruct2zkyw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zkyw > 0 || missingFieldsLeft2zkyw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zkyw, missingFieldsLeft2zkyw, msgp.ShowFound(found2zkyw[:]), decodeMsgFieldOrder2zkyw)
		if encodedFieldsLeft2zkyw > 0 {
			encodedFieldsLeft2zkyw--
			curField2zkyw, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2zkyw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zkyw = 0
			}
			for nextMiss2zkyw < maxFields2zkyw && (found2zkyw[nextMiss2zkyw] || decodeMsgFieldSkip2zkyw[nextMiss2zkyw]) {
				nextMiss2zkyw++
			}
			if nextMiss2zkyw == maxFields2zkyw {
				// filled all the empty fields!
				break doneWithStruct2zkyw
			}
			missingFieldsLeft2zkyw--
			curField2zkyw = nextMiss2zkyw
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zkyw)
		switch curField2zkyw {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Count"
			found2zkyw[0] = true
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
	if nextMiss2zkyw != -1 {
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
var decodeMsgFieldOrder2zkyw = []string{"Count"}

var decodeMsgFieldSkip2zkyw = []bool{false}

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
	var empty_znis [1]bool
	fieldsInUse_zfvt := z.fieldsNotEmpty(empty_znis[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zfvt + 1)
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

	if !empty_znis[0] {
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
	const maxFields3zjwt = 1

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zjwt uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zjwt, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zjwt := totalEncodedFields3zjwt
	missingFieldsLeft3zjwt := maxFields3zjwt - totalEncodedFields3zjwt

	var nextMiss3zjwt int = -1
	var found3zjwt [maxFields3zjwt]bool
	var curField3zjwt int

doneWithStruct3zjwt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zjwt > 0 || missingFieldsLeft3zjwt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zjwt, missingFieldsLeft3zjwt, msgp.ShowFound(found3zjwt[:]), unmarshalMsgFieldOrder3zjwt)
		if encodedFieldsLeft3zjwt > 0 {
			encodedFieldsLeft3zjwt--
			curField3zjwt, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zjwt < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zjwt = 0
			}
			for nextMiss3zjwt < maxFields3zjwt && (found3zjwt[nextMiss3zjwt] || unmarshalMsgFieldSkip3zjwt[nextMiss3zjwt]) {
				nextMiss3zjwt++
			}
			if nextMiss3zjwt == maxFields3zjwt {
				// filled all the empty fields!
				break doneWithStruct3zjwt
			}
			missingFieldsLeft3zjwt--
			curField3zjwt = nextMiss3zjwt
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zjwt)
		switch curField3zjwt {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Count"
			found3zjwt[0] = true
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
	if nextMiss3zjwt != -1 {
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
var unmarshalMsgFieldOrder3zjwt = []string{"Count"}

var unmarshalMsgFieldSkip3zjwt = []bool{false}

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
	const maxFields4zeuh = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zeuh uint32
	totalEncodedFields4zeuh, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zeuh := totalEncodedFields4zeuh
	missingFieldsLeft4zeuh := maxFields4zeuh - totalEncodedFields4zeuh

	var nextMiss4zeuh int = -1
	var found4zeuh [maxFields4zeuh]bool
	var curField4zeuh int

doneWithStruct4zeuh:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zeuh > 0 || missingFieldsLeft4zeuh > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zeuh, missingFieldsLeft4zeuh, msgp.ShowFound(found4zeuh[:]), decodeMsgFieldOrder4zeuh)
		if encodedFieldsLeft4zeuh > 0 {
			encodedFieldsLeft4zeuh--
			curField4zeuh, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zeuh < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zeuh = 0
			}
			for nextMiss4zeuh < maxFields4zeuh && (found4zeuh[nextMiss4zeuh] || decodeMsgFieldSkip4zeuh[nextMiss4zeuh]) {
				nextMiss4zeuh++
			}
			if nextMiss4zeuh == maxFields4zeuh {
				// filled all the empty fields!
				break doneWithStruct4zeuh
			}
			missingFieldsLeft4zeuh--
			curField4zeuh = nextMiss4zeuh
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zeuh)
		switch curField4zeuh {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "StartTime"
			found4zeuh[0] = true
			z.StartTime, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "State"
			found4zeuh[1] = true
			z.State, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "LoadAvg"
			found4zeuh[2] = true
			z.LoadAvg, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case 3:
			// zid 3 for "LoadInstant"
			found4zeuh[3] = true
			z.LoadInstant, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case 4:
			// zid 4 for "VmSize"
			found4zeuh[4] = true
			z.VmSize, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case 5:
			// zid 5 for "VmRSS"
			found4zeuh[5] = true
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
	if nextMiss4zeuh != -1 {
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
var decodeMsgFieldOrder4zeuh = []string{"StartTime", "State", "LoadAvg", "LoadInstant", "VmSize", "VmRSS"}

var decodeMsgFieldSkip4zeuh = []bool{false, false, false, false, false, false}

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
	var empty_zahw [6]bool
	fieldsInUse_zkvw := z.fieldsNotEmpty(empty_zahw[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zkvw + 1)
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

	if !empty_zahw[0] {
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

	if !empty_zahw[1] {
		// zid 2 for "State"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteString(z.State)
		if err != nil {
			return
		}
	}

	if !empty_zahw[2] {
		// zid 1 for "LoadAvg"
		err = en.Append(0x1)
		if err != nil {
			return err
		}
		err = en.WriteFloat64(z.LoadAvg)
		if err != nil {
			return
		}
	}

	if !empty_zahw[3] {
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

	if !empty_zahw[4] {
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

	if !empty_zahw[5] {
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
		// zid 2 for "State"
		o = append(o, 0x2)
		o = msgp.AppendString(o, z.State)
	}

	if !empty[2] {
		// zid 1 for "LoadAvg"
		o = append(o, 0x1)
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
	const maxFields5znbp = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5znbp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5znbp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5znbp := totalEncodedFields5znbp
	missingFieldsLeft5znbp := maxFields5znbp - totalEncodedFields5znbp

	var nextMiss5znbp int = -1
	var found5znbp [maxFields5znbp]bool
	var curField5znbp int

doneWithStruct5znbp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5znbp > 0 || missingFieldsLeft5znbp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5znbp, missingFieldsLeft5znbp, msgp.ShowFound(found5znbp[:]), unmarshalMsgFieldOrder5znbp)
		if encodedFieldsLeft5znbp > 0 {
			encodedFieldsLeft5znbp--
			curField5znbp, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5znbp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5znbp = 0
			}
			for nextMiss5znbp < maxFields5znbp && (found5znbp[nextMiss5znbp] || unmarshalMsgFieldSkip5znbp[nextMiss5znbp]) {
				nextMiss5znbp++
			}
			if nextMiss5znbp == maxFields5znbp {
				// filled all the empty fields!
				break doneWithStruct5znbp
			}
			missingFieldsLeft5znbp--
			curField5znbp = nextMiss5znbp
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5znbp)
		switch curField5znbp {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "StartTime"
			found5znbp[0] = true
			z.StartTime, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "State"
			found5znbp[1] = true
			z.State, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "LoadAvg"
			found5znbp[2] = true
			z.LoadAvg, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case 3:
			// zid 3 for "LoadInstant"
			found5znbp[3] = true
			z.LoadInstant, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case 4:
			// zid 4 for "VmSize"
			found5znbp[4] = true
			z.VmSize, bts, err = nbs.ReadUint64Bytes(bts)

			if err != nil {
				return
			}
		case 5:
			// zid 5 for "VmRSS"
			found5znbp[5] = true
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
	if nextMiss5znbp != -1 {
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
var unmarshalMsgFieldOrder5znbp = []string{"StartTime", "State", "LoadAvg", "LoadInstant", "VmSize", "VmRSS"}

var unmarshalMsgFieldSkip5znbp = []bool{false, false, false, false, false, false}

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
	const maxFields6zvsa = 5

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields6zvsa uint32
	totalEncodedFields6zvsa, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zvsa := totalEncodedFields6zvsa
	missingFieldsLeft6zvsa := maxFields6zvsa - totalEncodedFields6zvsa

	var nextMiss6zvsa int = -1
	var found6zvsa [maxFields6zvsa]bool
	var curField6zvsa int

doneWithStruct6zvsa:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zvsa > 0 || missingFieldsLeft6zvsa > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zvsa, missingFieldsLeft6zvsa, msgp.ShowFound(found6zvsa[:]), decodeMsgFieldOrder6zvsa)
		if encodedFieldsLeft6zvsa > 0 {
			encodedFieldsLeft6zvsa--
			curField6zvsa, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss6zvsa < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zvsa = 0
			}
			for nextMiss6zvsa < maxFields6zvsa && (found6zvsa[nextMiss6zvsa] || decodeMsgFieldSkip6zvsa[nextMiss6zvsa]) {
				nextMiss6zvsa++
			}
			if nextMiss6zvsa == maxFields6zvsa {
				// filled all the empty fields!
				break doneWithStruct6zvsa
			}
			missingFieldsLeft6zvsa--
			curField6zvsa = nextMiss6zvsa
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zvsa)
		switch curField6zvsa {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "ProcsStat"
			found6zvsa[0] = true
			var zkyu uint32
			zkyu, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.ProcsStat == nil && zkyu > 0 {
				z.ProcsStat = make(map[int]ProcStat, zkyu)
			} else if len(z.ProcsStat) > 0 {
				for key, _ := range z.ProcsStat {
					delete(z.ProcsStat, key)
				}
			}
			for zkyu > 0 {
				zkyu--
				var zbiy int
				var zzln ProcStat
				zbiy, err = dc.ReadInt()
				if err != nil {
					return
				}
				err = zzln.DecodeMsg(dc)
				if err != nil {
					return
				}
				z.ProcsStat[zbiy] = zzln
			}
		case 1:
			// zid 1 for "DoListen"
			found6zvsa[1] = true
			z.DoListen, err = dc.ReadBool()
			if err != nil {
				return
			}
		case 3:
			// zid 3 for "UpStream"
			found6zvsa[3] = true
			var zila uint32
			zila, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.UpStream == nil && zila > 0 {
				z.UpStream = make(map[string]AddrState, zila)
			} else if len(z.UpStream) > 0 {
				for key, _ := range z.UpStream {
					delete(z.UpStream, key)
				}
			}
			for zila > 0 {
				zila--
				var zvmt string
				var zjnz AddrState
				zvmt, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields7zyph = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields7zyph uint32
				totalEncodedFields7zyph, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft7zyph := totalEncodedFields7zyph
				missingFieldsLeft7zyph := maxFields7zyph - totalEncodedFields7zyph

				var nextMiss7zyph int = -1
				var found7zyph [maxFields7zyph]bool
				var curField7zyph int

			doneWithStruct7zyph:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft7zyph > 0 || missingFieldsLeft7zyph > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zyph, missingFieldsLeft7zyph, msgp.ShowFound(found7zyph[:]), decodeMsgFieldOrder7zyph)
					if encodedFieldsLeft7zyph > 0 {
						encodedFieldsLeft7zyph--
						curField7zyph, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss7zyph < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss7zyph = 0
						}
						for nextMiss7zyph < maxFields7zyph && (found7zyph[nextMiss7zyph] || decodeMsgFieldSkip7zyph[nextMiss7zyph]) {
							nextMiss7zyph++
						}
						if nextMiss7zyph == maxFields7zyph {
							// filled all the empty fields!
							break doneWithStruct7zyph
						}
						missingFieldsLeft7zyph--
						curField7zyph = nextMiss7zyph
					}
					//fmt.Printf("switching on curField: '%v'\n", curField7zyph)
					switch curField7zyph {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found7zyph[0] = true
						zjnz.Count, err = dc.ReadInt()
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
				if nextMiss7zyph != -1 {
					dc.PopAlwaysNil()
				}

				z.UpStream[zvmt] = zjnz
			}
		case 4:
			// zid 4 for "DownStream"
			found6zvsa[4] = true
			var zwvt uint32
			zwvt, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.DownStream == nil && zwvt > 0 {
				z.DownStream = make(map[string]AddrState, zwvt)
			} else if len(z.DownStream) > 0 {
				for key, _ := range z.DownStream {
					delete(z.DownStream, key)
				}
			}
			for zwvt > 0 {
				zwvt--
				var zfjx string
				var zywj AddrState
				zfjx, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields8zmpa = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields8zmpa uint32
				totalEncodedFields8zmpa, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft8zmpa := totalEncodedFields8zmpa
				missingFieldsLeft8zmpa := maxFields8zmpa - totalEncodedFields8zmpa

				var nextMiss8zmpa int = -1
				var found8zmpa [maxFields8zmpa]bool
				var curField8zmpa int

			doneWithStruct8zmpa:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft8zmpa > 0 || missingFieldsLeft8zmpa > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zmpa, missingFieldsLeft8zmpa, msgp.ShowFound(found8zmpa[:]), decodeMsgFieldOrder8zmpa)
					if encodedFieldsLeft8zmpa > 0 {
						encodedFieldsLeft8zmpa--
						curField8zmpa, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss8zmpa < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss8zmpa = 0
						}
						for nextMiss8zmpa < maxFields8zmpa && (found8zmpa[nextMiss8zmpa] || decodeMsgFieldSkip8zmpa[nextMiss8zmpa]) {
							nextMiss8zmpa++
						}
						if nextMiss8zmpa == maxFields8zmpa {
							// filled all the empty fields!
							break doneWithStruct8zmpa
						}
						missingFieldsLeft8zmpa--
						curField8zmpa = nextMiss8zmpa
					}
					//fmt.Printf("switching on curField: '%v'\n", curField8zmpa)
					switch curField8zmpa {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found8zmpa[0] = true
						zywj.Count, err = dc.ReadInt()
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
				if nextMiss8zmpa != -1 {
					dc.PopAlwaysNil()
				}

				z.DownStream[zfjx] = zywj
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss6zvsa != -1 {
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
var decodeMsgFieldOrder6zvsa = []string{"ProcsStat", "DoListen", "Addrs", "UpStream", "DownStream"}

var decodeMsgFieldSkip6zvsa = []bool{false, false, true, false, false}

// fields of AddrState
var decodeMsgFieldOrder7zyph = []string{"Count"}

var decodeMsgFieldSkip7zyph = []bool{false}

// fields of AddrState
var decodeMsgFieldOrder8zmpa = []string{"Count"}

var decodeMsgFieldSkip8zmpa = []bool{false}

// fieldsNotEmpty supports omitempty tags
func (z *ServiceInfo) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 4
	}
	var fieldsInUse uint32 = 4
	isempty[0] = (len(z.ProcsStat) == 0) // string, omitempty
	if isempty[0] {
		fieldsInUse--
	}
	isempty[1] = (!z.DoListen) // bool, omitempty
	if isempty[1] {
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
	var empty_znkn [5]bool
	fieldsInUse_zxjo := z.fieldsNotEmpty(empty_znkn[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zxjo + 1)
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

	if !empty_znkn[0] {
		// zid 0 for "ProcsStat"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.ProcsStat)))
		if err != nil {
			return
		}
		for zbiy, zzln := range z.ProcsStat {
			err = en.WriteInt(zbiy)
			if err != nil {
				return
			}
			err = zzln.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	if !empty_znkn[1] {
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

	if !empty_znkn[3] {
		// zid 3 for "UpStream"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.UpStream)))
		if err != nil {
			return
		}
		for zvmt, zjnz := range z.UpStream {
			err = en.WriteString(zvmt)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_zgkd [1]bool
			fieldsInUse_zcji := zjnz.fieldsNotEmpty(empty_zgkd[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zcji + 1)
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

			if !empty_zgkd[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zjnz.Count)
				if err != nil {
					return
				}
			}

		}
	}

	if !empty_znkn[4] {
		// zid 4 for "DownStream"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.DownStream)))
		if err != nil {
			return
		}
		for zfjx, zywj := range z.DownStream {
			err = en.WriteString(zfjx)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_ztia [1]bool
			fieldsInUse_zoak := zywj.fieldsNotEmpty(empty_ztia[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zoak + 1)
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

			if !empty_ztia[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zywj.Count)
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
		for zbiy, zzln := range z.ProcsStat {
			o = msgp.AppendInt(o, zbiy)
			o, err = zzln.MarshalMsg(o)
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

	if !empty[3] {
		// zid 3 for "UpStream"
		o = append(o, 0x3)
		o = msgp.AppendMapHeader(o, uint32(len(z.UpStream)))
		for zvmt, zjnz := range z.UpStream {
			o = msgp.AppendString(o, zvmt)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zjnz.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zjnz.Count)
			}

		}
	}

	if !empty[4] {
		// zid 4 for "DownStream"
		o = append(o, 0x4)
		o = msgp.AppendMapHeader(o, uint32(len(z.DownStream)))
		for zfjx, zywj := range z.DownStream {
			o = msgp.AppendString(o, zfjx)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zywj.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zywj.Count)
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
	const maxFields9zzzo = 5

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields9zzzo uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zzzo, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft9zzzo := totalEncodedFields9zzzo
	missingFieldsLeft9zzzo := maxFields9zzzo - totalEncodedFields9zzzo

	var nextMiss9zzzo int = -1
	var found9zzzo [maxFields9zzzo]bool
	var curField9zzzo int

doneWithStruct9zzzo:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zzzo > 0 || missingFieldsLeft9zzzo > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zzzo, missingFieldsLeft9zzzo, msgp.ShowFound(found9zzzo[:]), unmarshalMsgFieldOrder9zzzo)
		if encodedFieldsLeft9zzzo > 0 {
			encodedFieldsLeft9zzzo--
			curField9zzzo, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss9zzzo < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zzzo = 0
			}
			for nextMiss9zzzo < maxFields9zzzo && (found9zzzo[nextMiss9zzzo] || unmarshalMsgFieldSkip9zzzo[nextMiss9zzzo]) {
				nextMiss9zzzo++
			}
			if nextMiss9zzzo == maxFields9zzzo {
				// filled all the empty fields!
				break doneWithStruct9zzzo
			}
			missingFieldsLeft9zzzo--
			curField9zzzo = nextMiss9zzzo
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zzzo)
		switch curField9zzzo {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "ProcsStat"
			found9zzzo[0] = true
			if nbs.AlwaysNil {
				if len(z.ProcsStat) > 0 {
					for key, _ := range z.ProcsStat {
						delete(z.ProcsStat, key)
					}
				}

			} else {

				var zxot uint32
				zxot, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.ProcsStat == nil && zxot > 0 {
					z.ProcsStat = make(map[int]ProcStat, zxot)
				} else if len(z.ProcsStat) > 0 {
					for key, _ := range z.ProcsStat {
						delete(z.ProcsStat, key)
					}
				}
				for zxot > 0 {
					var zbiy int
					var zzln ProcStat
					zxot--
					zbiy, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					bts, err = zzln.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
					z.ProcsStat[zbiy] = zzln
				}
			}
		case 1:
			// zid 1 for "DoListen"
			found9zzzo[1] = true
			z.DoListen, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case 3:
			// zid 3 for "UpStream"
			found9zzzo[3] = true
			if nbs.AlwaysNil {
				if len(z.UpStream) > 0 {
					for key, _ := range z.UpStream {
						delete(z.UpStream, key)
					}
				}

			} else {

				var zxln uint32
				zxln, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.UpStream == nil && zxln > 0 {
					z.UpStream = make(map[string]AddrState, zxln)
				} else if len(z.UpStream) > 0 {
					for key, _ := range z.UpStream {
						delete(z.UpStream, key)
					}
				}
				for zxln > 0 {
					var zvmt string
					var zjnz AddrState
					zxln--
					zvmt, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields10zvfx = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields10zvfx uint32
					if !nbs.AlwaysNil {
						totalEncodedFields10zvfx, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft10zvfx := totalEncodedFields10zvfx
					missingFieldsLeft10zvfx := maxFields10zvfx - totalEncodedFields10zvfx

					var nextMiss10zvfx int = -1
					var found10zvfx [maxFields10zvfx]bool
					var curField10zvfx int

				doneWithStruct10zvfx:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft10zvfx > 0 || missingFieldsLeft10zvfx > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zvfx, missingFieldsLeft10zvfx, msgp.ShowFound(found10zvfx[:]), unmarshalMsgFieldOrder10zvfx)
						if encodedFieldsLeft10zvfx > 0 {
							encodedFieldsLeft10zvfx--
							curField10zvfx, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss10zvfx < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss10zvfx = 0
							}
							for nextMiss10zvfx < maxFields10zvfx && (found10zvfx[nextMiss10zvfx] || unmarshalMsgFieldSkip10zvfx[nextMiss10zvfx]) {
								nextMiss10zvfx++
							}
							if nextMiss10zvfx == maxFields10zvfx {
								// filled all the empty fields!
								break doneWithStruct10zvfx
							}
							missingFieldsLeft10zvfx--
							curField10zvfx = nextMiss10zvfx
						}
						//fmt.Printf("switching on curField: '%v'\n", curField10zvfx)
						switch curField10zvfx {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found10zvfx[0] = true
							zjnz.Count, bts, err = nbs.ReadIntBytes(bts)

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
					if nextMiss10zvfx != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.UpStream[zvmt] = zjnz
				}
			}
		case 4:
			// zid 4 for "DownStream"
			found9zzzo[4] = true
			if nbs.AlwaysNil {
				if len(z.DownStream) > 0 {
					for key, _ := range z.DownStream {
						delete(z.DownStream, key)
					}
				}

			} else {

				var zrcr uint32
				zrcr, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.DownStream == nil && zrcr > 0 {
					z.DownStream = make(map[string]AddrState, zrcr)
				} else if len(z.DownStream) > 0 {
					for key, _ := range z.DownStream {
						delete(z.DownStream, key)
					}
				}
				for zrcr > 0 {
					var zfjx string
					var zywj AddrState
					zrcr--
					zfjx, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields11zqdc = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields11zqdc uint32
					if !nbs.AlwaysNil {
						totalEncodedFields11zqdc, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft11zqdc := totalEncodedFields11zqdc
					missingFieldsLeft11zqdc := maxFields11zqdc - totalEncodedFields11zqdc

					var nextMiss11zqdc int = -1
					var found11zqdc [maxFields11zqdc]bool
					var curField11zqdc int

				doneWithStruct11zqdc:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft11zqdc > 0 || missingFieldsLeft11zqdc > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zqdc, missingFieldsLeft11zqdc, msgp.ShowFound(found11zqdc[:]), unmarshalMsgFieldOrder11zqdc)
						if encodedFieldsLeft11zqdc > 0 {
							encodedFieldsLeft11zqdc--
							curField11zqdc, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss11zqdc < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss11zqdc = 0
							}
							for nextMiss11zqdc < maxFields11zqdc && (found11zqdc[nextMiss11zqdc] || unmarshalMsgFieldSkip11zqdc[nextMiss11zqdc]) {
								nextMiss11zqdc++
							}
							if nextMiss11zqdc == maxFields11zqdc {
								// filled all the empty fields!
								break doneWithStruct11zqdc
							}
							missingFieldsLeft11zqdc--
							curField11zqdc = nextMiss11zqdc
						}
						//fmt.Printf("switching on curField: '%v'\n", curField11zqdc)
						switch curField11zqdc {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found11zqdc[0] = true
							zywj.Count, bts, err = nbs.ReadIntBytes(bts)

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
					if nextMiss11zqdc != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.DownStream[zfjx] = zywj
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss9zzzo != -1 {
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
var unmarshalMsgFieldOrder9zzzo = []string{"ProcsStat", "DoListen", "Addrs", "UpStream", "DownStream"}

var unmarshalMsgFieldSkip9zzzo = []bool{false, false, true, false, false}

// fields of AddrState
var unmarshalMsgFieldOrder10zvfx = []string{"Count"}

var unmarshalMsgFieldSkip10zvfx = []bool{false}

// fields of AddrState
var unmarshalMsgFieldOrder11zqdc = []string{"Count"}

var unmarshalMsgFieldSkip11zqdc = []bool{false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ServiceInfo) Msgsize() (s int) {
	s = 1 + 15 + msgp.MapHeaderSize
	if z.ProcsStat != nil {
		for zbiy, zzln := range z.ProcsStat {
			_ = zzln
			_ = zbiy
			s += msgp.IntSize + zzln.Msgsize()
		}
	}
	s += 15 + msgp.BoolSize + 15 + msgp.MapHeaderSize
	if z.UpStream != nil {
		for zvmt, zjnz := range z.UpStream {
			_ = zjnz
			_ = zvmt
			s += msgp.StringPrefixSize + len(zvmt) + 1 + 12 + msgp.IntSize
		}
	}
	s += 15 + msgp.MapHeaderSize
	if z.DownStream != nil {
		for zfjx, zywj := range z.DownStream {
			_ = zywj
			_ = zfjx
			s += msgp.StringPrefixSize + len(zfjx) + 1 + 12 + msgp.IntSize
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
	const maxFields12zxsf = 1

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields12zxsf uint32
	totalEncodedFields12zxsf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zxsf := totalEncodedFields12zxsf
	missingFieldsLeft12zxsf := maxFields12zxsf - totalEncodedFields12zxsf

	var nextMiss12zxsf int = -1
	var found12zxsf [maxFields12zxsf]bool
	var curField12zxsf int

doneWithStruct12zxsf:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zxsf > 0 || missingFieldsLeft12zxsf > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zxsf, missingFieldsLeft12zxsf, msgp.ShowFound(found12zxsf[:]), decodeMsgFieldOrder12zxsf)
		if encodedFieldsLeft12zxsf > 0 {
			encodedFieldsLeft12zxsf--
			curField12zxsf, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss12zxsf < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zxsf = 0
			}
			for nextMiss12zxsf < maxFields12zxsf && (found12zxsf[nextMiss12zxsf] || decodeMsgFieldSkip12zxsf[nextMiss12zxsf]) {
				nextMiss12zxsf++
			}
			if nextMiss12zxsf == maxFields12zxsf {
				// filled all the empty fields!
				break doneWithStruct12zxsf
			}
			missingFieldsLeft12zxsf--
			curField12zxsf = nextMiss12zxsf
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zxsf)
		switch curField12zxsf {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Services"
			found12zxsf[0] = true
			var zvkt uint32
			zvkt, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Services == nil && zvkt > 0 {
				z.Services = make(map[string]*ServiceInfo, zvkt)
			} else if len(z.Services) > 0 {
				for key, _ := range z.Services {
					delete(z.Services, key)
				}
			}
			for zvkt > 0 {
				zvkt--
				var zktf string
				var zlod *ServiceInfo
				zktf, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zlod != nil {
						dc.PushAlwaysNil()
						err = zlod.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zlod == nil {
						zlod = new(ServiceInfo)
					}
					err = zlod.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.Services[zktf] = zlod
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss12zxsf != -1 {
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
var decodeMsgFieldOrder12zxsf = []string{"Services"}

var decodeMsgFieldSkip12zxsf = []bool{false}

// fieldsNotEmpty supports omitempty tags
func (z *Topology) fieldsNotEmpty(isempty []bool) uint32 {
	if len(isempty) == 0 {
		return 1
	}
	var fieldsInUse uint32 = 1
	isempty[0] = (len(z.Services) == 0) // string, omitempty
	if isempty[0] {
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
	var empty_zolu [1]bool
	fieldsInUse_zxbc := z.fieldsNotEmpty(empty_zolu[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zxbc + 1)
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

	if !empty_zolu[0] {
		// zid 0 for "Services"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Services)))
		if err != nil {
			return
		}
		for zktf, zlod := range z.Services {
			err = en.WriteString(zktf)
			if err != nil {
				return
			}
			if zlod == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zlod.EncodeMsg(en)
				if err != nil {
					return
				}
			}
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
	var empty [1]bool
	fieldsInUse := z.fieldsNotEmpty(empty[:])
	o = msgp.AppendMapHeader(o, fieldsInUse+1)

	// runtime struct type identification for 'Topology'
	o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79})

	if !empty[0] {
		// zid 0 for "Services"
		o = append(o, 0x0)
		o = msgp.AppendMapHeader(o, uint32(len(z.Services)))
		for zktf, zlod := range z.Services {
			o = msgp.AppendString(o, zktf)
			if zlod == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zlod.MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
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
	const maxFields13zqpg = 1

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields13zqpg uint32
	if !nbs.AlwaysNil {
		totalEncodedFields13zqpg, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft13zqpg := totalEncodedFields13zqpg
	missingFieldsLeft13zqpg := maxFields13zqpg - totalEncodedFields13zqpg

	var nextMiss13zqpg int = -1
	var found13zqpg [maxFields13zqpg]bool
	var curField13zqpg int

doneWithStruct13zqpg:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft13zqpg > 0 || missingFieldsLeft13zqpg > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zqpg, missingFieldsLeft13zqpg, msgp.ShowFound(found13zqpg[:]), unmarshalMsgFieldOrder13zqpg)
		if encodedFieldsLeft13zqpg > 0 {
			encodedFieldsLeft13zqpg--
			curField13zqpg, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss13zqpg < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss13zqpg = 0
			}
			for nextMiss13zqpg < maxFields13zqpg && (found13zqpg[nextMiss13zqpg] || unmarshalMsgFieldSkip13zqpg[nextMiss13zqpg]) {
				nextMiss13zqpg++
			}
			if nextMiss13zqpg == maxFields13zqpg {
				// filled all the empty fields!
				break doneWithStruct13zqpg
			}
			missingFieldsLeft13zqpg--
			curField13zqpg = nextMiss13zqpg
		}
		//fmt.Printf("switching on curField: '%v'\n", curField13zqpg)
		switch curField13zqpg {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Services"
			found13zqpg[0] = true
			if nbs.AlwaysNil {
				if len(z.Services) > 0 {
					for key, _ := range z.Services {
						delete(z.Services, key)
					}
				}

			} else {

				var zoyq uint32
				zoyq, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Services == nil && zoyq > 0 {
					z.Services = make(map[string]*ServiceInfo, zoyq)
				} else if len(z.Services) > 0 {
					for key, _ := range z.Services {
						delete(z.Services, key)
					}
				}
				for zoyq > 0 {
					var zktf string
					var zlod *ServiceInfo
					zoyq--
					zktf, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zlod != nil {
							zlod.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zlod {
								zlod.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zlod == nil {
								zlod = new(ServiceInfo)
							}
							bts, err = zlod.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Services[zktf] = zlod
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss13zqpg != -1 {
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
var unmarshalMsgFieldOrder13zqpg = []string{"Services"}

var unmarshalMsgFieldSkip13zqpg = []bool{false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Topology) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.Services != nil {
		for zktf, zlod := range z.Services {
			_ = zlod
			_ = zktf
			s += msgp.StringPrefixSize + len(zktf)
			if zlod == nil {
				s += msgp.NilSize
			} else {
				s += zlod.Msgsize()
			}
		}
	}
	return
}

// FileSchema_go holds ZebraPack schema from file 'dataExt.go'
type FileSchema_go struct{}

// ZebraSchemaInMsgpack2Format provides the ZebraPack Schema in msgpack2 format, length 2341 bytes
func (FileSchema_go) ZebraSchemaInMsgpack2Format() []byte {
	return []byte{
		0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0xaa, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74,
		0x2e, 0x67, 0x6f, 0xad, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
		0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0xa4, 0x74, 0x6f,
		0x70, 0x6f, 0xad, 0x5a, 0x65, 0x62, 0x72, 0x61, 0x53, 0x63,
		0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0xd3, 0x00, 0x00, 0x4d,
		0xf6, 0x15, 0x1f, 0xb4, 0x97, 0xa7, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0x73, 0x85, 0xa8, 0x54, 0x6f, 0x70, 0x6f, 0x6c,
		0x6f, 0x67, 0x79, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63,
		0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x54, 0x6f, 0x70, 0x6f,
		0x6c, 0x6f, 0x67, 0x79, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x73, 0x91, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65,
		0xa8, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa8, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
		0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70,
		0x65, 0x53, 0x74, 0x72, 0xb7, 0x6d, 0x61, 0x70, 0x5b, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x2a, 0x53, 0x65, 0x72,
		0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46,
		0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d,
		0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61,
		0x6e, 0x67, 0x65, 0x83, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x1c,
		0xa3, 0x53, 0x74, 0x72, 0xa7, 0x50, 0x6f, 0x69, 0x6e, 0x74,
		0x65, 0x72, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72,
		0xab, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
		0x66, 0x6f, 0xa4, 0x41, 0x64, 0x64, 0x72, 0x82, 0xaa, 0x53,
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
		0xa8, 0x50, 0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x82,
		0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
		0x65, 0xa8, 0x50, 0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74,
		0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x96, 0x87, 0xa3,
		0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x53, 0x74, 0x61,
		0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa9,
		0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
		0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50,
		0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x11, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x11,
		0xa3, 0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34,
		0x87, 0xa3, 0x5a, 0x69, 0x64, 0x02, 0xab, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x53,
		0x74, 0x61, 0x74, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x53, 0x74,
		0x61, 0x74, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69,
		0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa7, 0x4c, 0x6f, 0x61, 0x64, 0x41,
		0x76, 0x67, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
		0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa7, 0x4c, 0x6f, 0x61, 0x64,
		0x41, 0x76, 0x67, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa7, 0x66, 0x6c, 0x6f,
		0x61, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69,
		0x74, 0x69, 0x76, 0x65, 0x04, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x04, 0xa3, 0x53, 0x74, 0x72,
		0xa7, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x87, 0xa3,
		0x5a, 0x69, 0x64, 0x03, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xab, 0x4c, 0x6f, 0x61,
		0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0xab, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x73, 0x74,
		0x61, 0x6e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa7, 0x66, 0x6c, 0x6f,
		0x61, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69,
		0x74, 0x69, 0x76, 0x65, 0x04, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x04, 0xa3, 0x53, 0x74, 0x72,
		0xa7, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x87, 0xa3,
		0x5a, 0x69, 0x64, 0x04, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x56, 0x6d, 0x53,
		0x69, 0x7a, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa6, 0x56, 0x6d, 0x53,
		0x69, 0x7a, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e,
		0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x0b, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x0b, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x87, 0xa3, 0x5a, 0x69,
		0x64, 0x05, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x56, 0x6d, 0x52, 0x53, 0x53,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa5, 0x56, 0x6d, 0x52, 0x53, 0x53, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
		0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67,
		0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x0b,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c,
		0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x0b, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74,
		0x36, 0x34, 0xab, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
		0x49, 0x6e, 0x66, 0x6f, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xab, 0x53, 0x65, 0x72,
		0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0xa6, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x73, 0x95, 0x86, 0xa3, 0x5a, 0x69,
		0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x50, 0x72, 0x6f, 0x63, 0x73,
		0x53, 0x74, 0x61, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x50, 0x72,
		0x6f, 0x63, 0x73, 0x53, 0x74, 0x61, 0x74, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0xb0, 0x6d, 0x61, 0x70, 0x5b, 0x69, 0x6e, 0x74, 0x5d, 0x50,
		0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61,
		0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x0d, 0xa3, 0x53, 0x74, 0x72, 0xa3,
		0x69, 0x6e, 0x74, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72,
		0xa8, 0x50, 0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x87,
		0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x44, 0x6f,
		0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa8,
		0x44, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
		0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53,
		0x74, 0x72, 0xa4, 0x62, 0x6f, 0x6f, 0x6c, 0x84, 0xa3, 0x5a,
		0x69, 0x64, 0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x41, 0x64, 0x64, 0x72,
		0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x41, 0x64, 0x64, 0x72, 0x73,
		0xa4, 0x53, 0x6b, 0x69, 0x70, 0xc3, 0x86, 0xa3, 0x5a, 0x69,
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
		0x63, 0x74, 0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73,
		0x90,
	}
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 2992 bytes
func (FileSchema_go) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath":"dataExt.go","SourcePackage":"topo","ZebraSchemaId":85719311692951,"Structs":{"Topology":{"StructName":"Topology","Fields":[{"Zid":0,"FieldGoName":"Services","FieldTagName":"Services","FieldTypeStr":"map[string]*ServiceInfo","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"ServiceInfo"}}}}]},"Addr":{"StructName":"Addr","Fields":[{"Zid":0,"FieldGoName":"Host","FieldTagName":"Host","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":1,"FieldGoName":"Port","FieldTagName":"Port","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}}]},"AddrState":{"StructName":"AddrState","Fields":[{"Zid":0,"FieldGoName":"Count","FieldTagName":"Count","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}}]},"ProcStat":{"StructName":"ProcStat","Fields":[{"Zid":0,"FieldGoName":"StartTime","FieldTagName":"StartTime","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}},{"Zid":2,"FieldGoName":"State","FieldTagName":"State","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":1,"FieldGoName":"LoadAvg","FieldTagName":"LoadAvg","FieldTypeStr":"float64","FieldCategory":23,"FieldPrimitive":4,"FieldFullType":{"Kind":4,"Str":"float64"}},{"Zid":3,"FieldGoName":"LoadInstant","FieldTagName":"LoadInstant","FieldTypeStr":"float64","FieldCategory":23,"FieldPrimitive":4,"FieldFullType":{"Kind":4,"Str":"float64"}},{"Zid":4,"FieldGoName":"VmSize","FieldTagName":"VmSize","FieldTypeStr":"uint64","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}},{"Zid":5,"FieldGoName":"VmRSS","FieldTagName":"VmRSS","FieldTypeStr":"uint64","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}}]},"ServiceInfo":{"StructName":"ServiceInfo","Fields":[{"Zid":0,"FieldGoName":"ProcsStat","FieldTagName":"ProcsStat","FieldTypeStr":"map[int]ProcStat","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":13,"Str":"int"},"Range":{"Kind":22,"Str":"ProcStat"}}},{"Zid":1,"FieldGoName":"DoListen","FieldTagName":"DoListen","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"}},{"Zid":2,"FieldGoName":"Addrs","FieldTagName":"Addrs","Skip":true},{"Zid":3,"FieldGoName":"UpStream","FieldTagName":"UpStream","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}},{"Zid":4,"FieldGoName":"DownStream","FieldTagName":"DownStream","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}}]}},"Imports":[]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 8031 bytes
func (FileSchema_go) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "dataExt.go",
    "SourcePackage": "topo",
    "ZebraSchemaId": 85719311692951,
    "Structs": {
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
        },
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
                    "Zid": 2,
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
                    "Zid": 1,
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
                    "Skip": true
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
        }
    },
    "Imports": []
}`)
}
