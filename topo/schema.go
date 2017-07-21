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
	const maxFields0zkvt = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zkvt uint32
	totalEncodedFields0zkvt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zkvt := totalEncodedFields0zkvt
	missingFieldsLeft0zkvt := maxFields0zkvt - totalEncodedFields0zkvt

	var nextMiss0zkvt int = -1
	var found0zkvt [maxFields0zkvt]bool
	var curField0zkvt int

doneWithStruct0zkvt:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zkvt > 0 || missingFieldsLeft0zkvt > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zkvt, missingFieldsLeft0zkvt, msgp.ShowFound(found0zkvt[:]), decodeMsgFieldOrder0zkvt)
		if encodedFieldsLeft0zkvt > 0 {
			encodedFieldsLeft0zkvt--
			curField0zkvt, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zkvt < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zkvt = 0
			}
			for nextMiss0zkvt < maxFields0zkvt && (found0zkvt[nextMiss0zkvt] || decodeMsgFieldSkip0zkvt[nextMiss0zkvt]) {
				nextMiss0zkvt++
			}
			if nextMiss0zkvt == maxFields0zkvt {
				// filled all the empty fields!
				break doneWithStruct0zkvt
			}
			missingFieldsLeft0zkvt--
			curField0zkvt = nextMiss0zkvt
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zkvt)
		switch curField0zkvt {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Host"
			found0zkvt[0] = true
			z.Host, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Port"
			found0zkvt[1] = true
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
	if nextMiss0zkvt != -1 {
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
var decodeMsgFieldOrder0zkvt = []string{"Host", "Port"}

var decodeMsgFieldSkip0zkvt = []bool{false, false}

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
	var empty_zpsy [2]bool
	fieldsInUse_zdzj := z.fieldsNotEmpty(empty_zpsy[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zdzj + 1)
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

	if !empty_zpsy[0] {
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

	if !empty_zpsy[1] {
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
	const maxFields1zpzd = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zpzd uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zpzd, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zpzd := totalEncodedFields1zpzd
	missingFieldsLeft1zpzd := maxFields1zpzd - totalEncodedFields1zpzd

	var nextMiss1zpzd int = -1
	var found1zpzd [maxFields1zpzd]bool
	var curField1zpzd int

doneWithStruct1zpzd:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zpzd > 0 || missingFieldsLeft1zpzd > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zpzd, missingFieldsLeft1zpzd, msgp.ShowFound(found1zpzd[:]), unmarshalMsgFieldOrder1zpzd)
		if encodedFieldsLeft1zpzd > 0 {
			encodedFieldsLeft1zpzd--
			curField1zpzd, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zpzd < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zpzd = 0
			}
			for nextMiss1zpzd < maxFields1zpzd && (found1zpzd[nextMiss1zpzd] || unmarshalMsgFieldSkip1zpzd[nextMiss1zpzd]) {
				nextMiss1zpzd++
			}
			if nextMiss1zpzd == maxFields1zpzd {
				// filled all the empty fields!
				break doneWithStruct1zpzd
			}
			missingFieldsLeft1zpzd--
			curField1zpzd = nextMiss1zpzd
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zpzd)
		switch curField1zpzd {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Host"
			found1zpzd[0] = true
			z.Host, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Port"
			found1zpzd[1] = true
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
	if nextMiss1zpzd != -1 {
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
var unmarshalMsgFieldOrder1zpzd = []string{"Host", "Port"}

var unmarshalMsgFieldSkip1zpzd = []bool{false, false}

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
	const maxFields2zedu = 1

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2zedu uint32
	totalEncodedFields2zedu, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zedu := totalEncodedFields2zedu
	missingFieldsLeft2zedu := maxFields2zedu - totalEncodedFields2zedu

	var nextMiss2zedu int = -1
	var found2zedu [maxFields2zedu]bool
	var curField2zedu int

doneWithStruct2zedu:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zedu > 0 || missingFieldsLeft2zedu > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zedu, missingFieldsLeft2zedu, msgp.ShowFound(found2zedu[:]), decodeMsgFieldOrder2zedu)
		if encodedFieldsLeft2zedu > 0 {
			encodedFieldsLeft2zedu--
			curField2zedu, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2zedu < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zedu = 0
			}
			for nextMiss2zedu < maxFields2zedu && (found2zedu[nextMiss2zedu] || decodeMsgFieldSkip2zedu[nextMiss2zedu]) {
				nextMiss2zedu++
			}
			if nextMiss2zedu == maxFields2zedu {
				// filled all the empty fields!
				break doneWithStruct2zedu
			}
			missingFieldsLeft2zedu--
			curField2zedu = nextMiss2zedu
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zedu)
		switch curField2zedu {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Count"
			found2zedu[0] = true
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
	if nextMiss2zedu != -1 {
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
var decodeMsgFieldOrder2zedu = []string{"Count"}

var decodeMsgFieldSkip2zedu = []bool{false}

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
	var empty_zhvy [1]bool
	fieldsInUse_zpzw := z.fieldsNotEmpty(empty_zhvy[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zpzw + 1)
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

	if !empty_zhvy[0] {
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
	const maxFields3zrft = 1

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zrft uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zrft, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zrft := totalEncodedFields3zrft
	missingFieldsLeft3zrft := maxFields3zrft - totalEncodedFields3zrft

	var nextMiss3zrft int = -1
	var found3zrft [maxFields3zrft]bool
	var curField3zrft int

doneWithStruct3zrft:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zrft > 0 || missingFieldsLeft3zrft > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zrft, missingFieldsLeft3zrft, msgp.ShowFound(found3zrft[:]), unmarshalMsgFieldOrder3zrft)
		if encodedFieldsLeft3zrft > 0 {
			encodedFieldsLeft3zrft--
			curField3zrft, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zrft < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zrft = 0
			}
			for nextMiss3zrft < maxFields3zrft && (found3zrft[nextMiss3zrft] || unmarshalMsgFieldSkip3zrft[nextMiss3zrft]) {
				nextMiss3zrft++
			}
			if nextMiss3zrft == maxFields3zrft {
				// filled all the empty fields!
				break doneWithStruct3zrft
			}
			missingFieldsLeft3zrft--
			curField3zrft = nextMiss3zrft
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zrft)
		switch curField3zrft {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Count"
			found3zrft[0] = true
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
	if nextMiss3zrft != -1 {
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
var unmarshalMsgFieldOrder3zrft = []string{"Count"}

var unmarshalMsgFieldSkip3zrft = []bool{false}

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
	const maxFields4zaiz = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zaiz uint32
	totalEncodedFields4zaiz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zaiz := totalEncodedFields4zaiz
	missingFieldsLeft4zaiz := maxFields4zaiz - totalEncodedFields4zaiz

	var nextMiss4zaiz int = -1
	var found4zaiz [maxFields4zaiz]bool
	var curField4zaiz int

doneWithStruct4zaiz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zaiz > 0 || missingFieldsLeft4zaiz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zaiz, missingFieldsLeft4zaiz, msgp.ShowFound(found4zaiz[:]), decodeMsgFieldOrder4zaiz)
		if encodedFieldsLeft4zaiz > 0 {
			encodedFieldsLeft4zaiz--
			curField4zaiz, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zaiz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zaiz = 0
			}
			for nextMiss4zaiz < maxFields4zaiz && (found4zaiz[nextMiss4zaiz] || decodeMsgFieldSkip4zaiz[nextMiss4zaiz]) {
				nextMiss4zaiz++
			}
			if nextMiss4zaiz == maxFields4zaiz {
				// filled all the empty fields!
				break doneWithStruct4zaiz
			}
			missingFieldsLeft4zaiz--
			curField4zaiz = nextMiss4zaiz
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zaiz)
		switch curField4zaiz {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "StartTime"
			found4zaiz[0] = true
			z.StartTime, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "State"
			found4zaiz[1] = true
			z.State, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "LoadAvg"
			found4zaiz[2] = true
			z.LoadAvg, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case 3:
			// zid 3 for "LoadInstant"
			found4zaiz[3] = true
			z.LoadInstant, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case 4:
			// zid 4 for "VmSize"
			found4zaiz[4] = true
			z.VmSize, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case 5:
			// zid 5 for "VmRSS"
			found4zaiz[5] = true
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
	if nextMiss4zaiz != -1 {
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
var decodeMsgFieldOrder4zaiz = []string{"StartTime", "State", "LoadAvg", "LoadInstant", "VmSize", "VmRSS"}

var decodeMsgFieldSkip4zaiz = []bool{false, false, false, false, false, false}

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
	var empty_zdqe [6]bool
	fieldsInUse_zjzo := z.fieldsNotEmpty(empty_zdqe[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zjzo + 1)
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

	if !empty_zdqe[0] {
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

	if !empty_zdqe[1] {
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

	if !empty_zdqe[2] {
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

	if !empty_zdqe[3] {
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

	if !empty_zdqe[4] {
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

	if !empty_zdqe[5] {
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
	const maxFields5zcsq = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5zcsq uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zcsq, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zcsq := totalEncodedFields5zcsq
	missingFieldsLeft5zcsq := maxFields5zcsq - totalEncodedFields5zcsq

	var nextMiss5zcsq int = -1
	var found5zcsq [maxFields5zcsq]bool
	var curField5zcsq int

doneWithStruct5zcsq:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zcsq > 0 || missingFieldsLeft5zcsq > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zcsq, missingFieldsLeft5zcsq, msgp.ShowFound(found5zcsq[:]), unmarshalMsgFieldOrder5zcsq)
		if encodedFieldsLeft5zcsq > 0 {
			encodedFieldsLeft5zcsq--
			curField5zcsq, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5zcsq < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zcsq = 0
			}
			for nextMiss5zcsq < maxFields5zcsq && (found5zcsq[nextMiss5zcsq] || unmarshalMsgFieldSkip5zcsq[nextMiss5zcsq]) {
				nextMiss5zcsq++
			}
			if nextMiss5zcsq == maxFields5zcsq {
				// filled all the empty fields!
				break doneWithStruct5zcsq
			}
			missingFieldsLeft5zcsq--
			curField5zcsq = nextMiss5zcsq
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zcsq)
		switch curField5zcsq {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "StartTime"
			found5zcsq[0] = true
			z.StartTime, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "State"
			found5zcsq[1] = true
			z.State, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "LoadAvg"
			found5zcsq[2] = true
			z.LoadAvg, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case 3:
			// zid 3 for "LoadInstant"
			found5zcsq[3] = true
			z.LoadInstant, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case 4:
			// zid 4 for "VmSize"
			found5zcsq[4] = true
			z.VmSize, bts, err = nbs.ReadUint64Bytes(bts)

			if err != nil {
				return
			}
		case 5:
			// zid 5 for "VmRSS"
			found5zcsq[5] = true
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
	if nextMiss5zcsq != -1 {
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
var unmarshalMsgFieldOrder5zcsq = []string{"StartTime", "State", "LoadAvg", "LoadInstant", "VmSize", "VmRSS"}

var unmarshalMsgFieldSkip5zcsq = []bool{false, false, false, false, false, false}

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
	const maxFields6zwne = 5

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields6zwne uint32
	totalEncodedFields6zwne, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zwne := totalEncodedFields6zwne
	missingFieldsLeft6zwne := maxFields6zwne - totalEncodedFields6zwne

	var nextMiss6zwne int = -1
	var found6zwne [maxFields6zwne]bool
	var curField6zwne int

doneWithStruct6zwne:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zwne > 0 || missingFieldsLeft6zwne > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zwne, missingFieldsLeft6zwne, msgp.ShowFound(found6zwne[:]), decodeMsgFieldOrder6zwne)
		if encodedFieldsLeft6zwne > 0 {
			encodedFieldsLeft6zwne--
			curField6zwne, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss6zwne < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zwne = 0
			}
			for nextMiss6zwne < maxFields6zwne && (found6zwne[nextMiss6zwne] || decodeMsgFieldSkip6zwne[nextMiss6zwne]) {
				nextMiss6zwne++
			}
			if nextMiss6zwne == maxFields6zwne {
				// filled all the empty fields!
				break doneWithStruct6zwne
			}
			missingFieldsLeft6zwne--
			curField6zwne = nextMiss6zwne
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zwne)
		switch curField6zwne {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "ProcsStat"
			found6zwne[0] = true
			var zamf uint32
			zamf, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.ProcsStat == nil && zamf > 0 {
				z.ProcsStat = make(map[int]ProcStat, zamf)
			} else if len(z.ProcsStat) > 0 {
				for key, _ := range z.ProcsStat {
					delete(z.ProcsStat, key)
				}
			}
			for zamf > 0 {
				zamf--
				var zxmx int
				var zkyf ProcStat
				zxmx, err = dc.ReadInt()
				if err != nil {
					return
				}
				err = zkyf.DecodeMsg(dc)
				if err != nil {
					return
				}
				z.ProcsStat[zxmx] = zkyf
			}
		case 1:
			// zid 1 for "DoListen"
			found6zwne[1] = true
			z.DoListen, err = dc.ReadBool()
			if err != nil {
				return
			}
		case 3:
			// zid 3 for "UpStream"
			found6zwne[3] = true
			var zuxp uint32
			zuxp, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.UpStream == nil && zuxp > 0 {
				z.UpStream = make(map[string]AddrState, zuxp)
			} else if len(z.UpStream) > 0 {
				for key, _ := range z.UpStream {
					delete(z.UpStream, key)
				}
			}
			for zuxp > 0 {
				zuxp--
				var zsrz string
				var zctb AddrState
				zsrz, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields7zbig = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields7zbig uint32
				totalEncodedFields7zbig, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft7zbig := totalEncodedFields7zbig
				missingFieldsLeft7zbig := maxFields7zbig - totalEncodedFields7zbig

				var nextMiss7zbig int = -1
				var found7zbig [maxFields7zbig]bool
				var curField7zbig int

			doneWithStruct7zbig:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft7zbig > 0 || missingFieldsLeft7zbig > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zbig, missingFieldsLeft7zbig, msgp.ShowFound(found7zbig[:]), decodeMsgFieldOrder7zbig)
					if encodedFieldsLeft7zbig > 0 {
						encodedFieldsLeft7zbig--
						curField7zbig, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss7zbig < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss7zbig = 0
						}
						for nextMiss7zbig < maxFields7zbig && (found7zbig[nextMiss7zbig] || decodeMsgFieldSkip7zbig[nextMiss7zbig]) {
							nextMiss7zbig++
						}
						if nextMiss7zbig == maxFields7zbig {
							// filled all the empty fields!
							break doneWithStruct7zbig
						}
						missingFieldsLeft7zbig--
						curField7zbig = nextMiss7zbig
					}
					//fmt.Printf("switching on curField: '%v'\n", curField7zbig)
					switch curField7zbig {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found7zbig[0] = true
						zctb.Count, err = dc.ReadInt()
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
				if nextMiss7zbig != -1 {
					dc.PopAlwaysNil()
				}

				z.UpStream[zsrz] = zctb
			}
		case 4:
			// zid 4 for "DownStream"
			found6zwne[4] = true
			var zgca uint32
			zgca, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.DownStream == nil && zgca > 0 {
				z.DownStream = make(map[string]AddrState, zgca)
			} else if len(z.DownStream) > 0 {
				for key, _ := range z.DownStream {
					delete(z.DownStream, key)
				}
			}
			for zgca > 0 {
				zgca--
				var zyos string
				var zegv AddrState
				zyos, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields8zahd = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields8zahd uint32
				totalEncodedFields8zahd, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft8zahd := totalEncodedFields8zahd
				missingFieldsLeft8zahd := maxFields8zahd - totalEncodedFields8zahd

				var nextMiss8zahd int = -1
				var found8zahd [maxFields8zahd]bool
				var curField8zahd int

			doneWithStruct8zahd:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft8zahd > 0 || missingFieldsLeft8zahd > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zahd, missingFieldsLeft8zahd, msgp.ShowFound(found8zahd[:]), decodeMsgFieldOrder8zahd)
					if encodedFieldsLeft8zahd > 0 {
						encodedFieldsLeft8zahd--
						curField8zahd, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss8zahd < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss8zahd = 0
						}
						for nextMiss8zahd < maxFields8zahd && (found8zahd[nextMiss8zahd] || decodeMsgFieldSkip8zahd[nextMiss8zahd]) {
							nextMiss8zahd++
						}
						if nextMiss8zahd == maxFields8zahd {
							// filled all the empty fields!
							break doneWithStruct8zahd
						}
						missingFieldsLeft8zahd--
						curField8zahd = nextMiss8zahd
					}
					//fmt.Printf("switching on curField: '%v'\n", curField8zahd)
					switch curField8zahd {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found8zahd[0] = true
						zegv.Count, err = dc.ReadInt()
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
				if nextMiss8zahd != -1 {
					dc.PopAlwaysNil()
				}

				z.DownStream[zyos] = zegv
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss6zwne != -1 {
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
var decodeMsgFieldOrder6zwne = []string{"ProcsStat", "DoListen", "Addrs", "UpStream", "DownStream"}

var decodeMsgFieldSkip6zwne = []bool{false, false, true, false, false}

// fields of AddrState
var decodeMsgFieldOrder7zbig = []string{"Count"}

var decodeMsgFieldSkip7zbig = []bool{false}

// fields of AddrState
var decodeMsgFieldOrder8zahd = []string{"Count"}

var decodeMsgFieldSkip8zahd = []bool{false}

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
	var empty_zbkv [5]bool
	fieldsInUse_zqaj := z.fieldsNotEmpty(empty_zbkv[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqaj + 1)
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

	if !empty_zbkv[0] {
		// zid 0 for "ProcsStat"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.ProcsStat)))
		if err != nil {
			return
		}
		for zxmx, zkyf := range z.ProcsStat {
			err = en.WriteInt(zxmx)
			if err != nil {
				return
			}
			err = zkyf.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	if !empty_zbkv[1] {
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

	if !empty_zbkv[3] {
		// zid 3 for "UpStream"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.UpStream)))
		if err != nil {
			return
		}
		for zsrz, zctb := range z.UpStream {
			err = en.WriteString(zsrz)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_zaqx [1]bool
			fieldsInUse_zacc := zctb.fieldsNotEmpty(empty_zaqx[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zacc + 1)
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

			if !empty_zaqx[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zctb.Count)
				if err != nil {
					return
				}
			}

		}
	}

	if !empty_zbkv[4] {
		// zid 4 for "DownStream"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.DownStream)))
		if err != nil {
			return
		}
		for zyos, zegv := range z.DownStream {
			err = en.WriteString(zyos)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_znuf [1]bool
			fieldsInUse_zfxk := zegv.fieldsNotEmpty(empty_znuf[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zfxk + 1)
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

			if !empty_znuf[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zegv.Count)
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
		for zxmx, zkyf := range z.ProcsStat {
			o = msgp.AppendInt(o, zxmx)
			o, err = zkyf.MarshalMsg(o)
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
		for zsrz, zctb := range z.UpStream {
			o = msgp.AppendString(o, zsrz)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zctb.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zctb.Count)
			}

		}
	}

	if !empty[4] {
		// zid 4 for "DownStream"
		o = append(o, 0x4)
		o = msgp.AppendMapHeader(o, uint32(len(z.DownStream)))
		for zyos, zegv := range z.DownStream {
			o = msgp.AppendString(o, zyos)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zegv.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zegv.Count)
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
	const maxFields9zwsp = 5

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields9zwsp uint32
	if !nbs.AlwaysNil {
		totalEncodedFields9zwsp, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft9zwsp := totalEncodedFields9zwsp
	missingFieldsLeft9zwsp := maxFields9zwsp - totalEncodedFields9zwsp

	var nextMiss9zwsp int = -1
	var found9zwsp [maxFields9zwsp]bool
	var curField9zwsp int

doneWithStruct9zwsp:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft9zwsp > 0 || missingFieldsLeft9zwsp > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zwsp, missingFieldsLeft9zwsp, msgp.ShowFound(found9zwsp[:]), unmarshalMsgFieldOrder9zwsp)
		if encodedFieldsLeft9zwsp > 0 {
			encodedFieldsLeft9zwsp--
			curField9zwsp, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss9zwsp < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss9zwsp = 0
			}
			for nextMiss9zwsp < maxFields9zwsp && (found9zwsp[nextMiss9zwsp] || unmarshalMsgFieldSkip9zwsp[nextMiss9zwsp]) {
				nextMiss9zwsp++
			}
			if nextMiss9zwsp == maxFields9zwsp {
				// filled all the empty fields!
				break doneWithStruct9zwsp
			}
			missingFieldsLeft9zwsp--
			curField9zwsp = nextMiss9zwsp
		}
		//fmt.Printf("switching on curField: '%v'\n", curField9zwsp)
		switch curField9zwsp {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "ProcsStat"
			found9zwsp[0] = true
			if nbs.AlwaysNil {
				if len(z.ProcsStat) > 0 {
					for key, _ := range z.ProcsStat {
						delete(z.ProcsStat, key)
					}
				}

			} else {

				var zvhn uint32
				zvhn, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.ProcsStat == nil && zvhn > 0 {
					z.ProcsStat = make(map[int]ProcStat, zvhn)
				} else if len(z.ProcsStat) > 0 {
					for key, _ := range z.ProcsStat {
						delete(z.ProcsStat, key)
					}
				}
				for zvhn > 0 {
					var zxmx int
					var zkyf ProcStat
					zvhn--
					zxmx, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					bts, err = zkyf.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
					z.ProcsStat[zxmx] = zkyf
				}
			}
		case 1:
			// zid 1 for "DoListen"
			found9zwsp[1] = true
			z.DoListen, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case 3:
			// zid 3 for "UpStream"
			found9zwsp[3] = true
			if nbs.AlwaysNil {
				if len(z.UpStream) > 0 {
					for key, _ := range z.UpStream {
						delete(z.UpStream, key)
					}
				}

			} else {

				var zawl uint32
				zawl, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.UpStream == nil && zawl > 0 {
					z.UpStream = make(map[string]AddrState, zawl)
				} else if len(z.UpStream) > 0 {
					for key, _ := range z.UpStream {
						delete(z.UpStream, key)
					}
				}
				for zawl > 0 {
					var zsrz string
					var zctb AddrState
					zawl--
					zsrz, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields10zuco = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields10zuco uint32
					if !nbs.AlwaysNil {
						totalEncodedFields10zuco, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft10zuco := totalEncodedFields10zuco
					missingFieldsLeft10zuco := maxFields10zuco - totalEncodedFields10zuco

					var nextMiss10zuco int = -1
					var found10zuco [maxFields10zuco]bool
					var curField10zuco int

				doneWithStruct10zuco:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft10zuco > 0 || missingFieldsLeft10zuco > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zuco, missingFieldsLeft10zuco, msgp.ShowFound(found10zuco[:]), unmarshalMsgFieldOrder10zuco)
						if encodedFieldsLeft10zuco > 0 {
							encodedFieldsLeft10zuco--
							curField10zuco, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss10zuco < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss10zuco = 0
							}
							for nextMiss10zuco < maxFields10zuco && (found10zuco[nextMiss10zuco] || unmarshalMsgFieldSkip10zuco[nextMiss10zuco]) {
								nextMiss10zuco++
							}
							if nextMiss10zuco == maxFields10zuco {
								// filled all the empty fields!
								break doneWithStruct10zuco
							}
							missingFieldsLeft10zuco--
							curField10zuco = nextMiss10zuco
						}
						//fmt.Printf("switching on curField: '%v'\n", curField10zuco)
						switch curField10zuco {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found10zuco[0] = true
							zctb.Count, bts, err = nbs.ReadIntBytes(bts)

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
					if nextMiss10zuco != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.UpStream[zsrz] = zctb
				}
			}
		case 4:
			// zid 4 for "DownStream"
			found9zwsp[4] = true
			if nbs.AlwaysNil {
				if len(z.DownStream) > 0 {
					for key, _ := range z.DownStream {
						delete(z.DownStream, key)
					}
				}

			} else {

				var znvf uint32
				znvf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.DownStream == nil && znvf > 0 {
					z.DownStream = make(map[string]AddrState, znvf)
				} else if len(z.DownStream) > 0 {
					for key, _ := range z.DownStream {
						delete(z.DownStream, key)
					}
				}
				for znvf > 0 {
					var zyos string
					var zegv AddrState
					znvf--
					zyos, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields11zogd = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields11zogd uint32
					if !nbs.AlwaysNil {
						totalEncodedFields11zogd, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft11zogd := totalEncodedFields11zogd
					missingFieldsLeft11zogd := maxFields11zogd - totalEncodedFields11zogd

					var nextMiss11zogd int = -1
					var found11zogd [maxFields11zogd]bool
					var curField11zogd int

				doneWithStruct11zogd:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft11zogd > 0 || missingFieldsLeft11zogd > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zogd, missingFieldsLeft11zogd, msgp.ShowFound(found11zogd[:]), unmarshalMsgFieldOrder11zogd)
						if encodedFieldsLeft11zogd > 0 {
							encodedFieldsLeft11zogd--
							curField11zogd, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss11zogd < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss11zogd = 0
							}
							for nextMiss11zogd < maxFields11zogd && (found11zogd[nextMiss11zogd] || unmarshalMsgFieldSkip11zogd[nextMiss11zogd]) {
								nextMiss11zogd++
							}
							if nextMiss11zogd == maxFields11zogd {
								// filled all the empty fields!
								break doneWithStruct11zogd
							}
							missingFieldsLeft11zogd--
							curField11zogd = nextMiss11zogd
						}
						//fmt.Printf("switching on curField: '%v'\n", curField11zogd)
						switch curField11zogd {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found11zogd[0] = true
							zegv.Count, bts, err = nbs.ReadIntBytes(bts)

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
					if nextMiss11zogd != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.DownStream[zyos] = zegv
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss9zwsp != -1 {
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
var unmarshalMsgFieldOrder9zwsp = []string{"ProcsStat", "DoListen", "Addrs", "UpStream", "DownStream"}

var unmarshalMsgFieldSkip9zwsp = []bool{false, false, true, false, false}

// fields of AddrState
var unmarshalMsgFieldOrder10zuco = []string{"Count"}

var unmarshalMsgFieldSkip10zuco = []bool{false}

// fields of AddrState
var unmarshalMsgFieldOrder11zogd = []string{"Count"}

var unmarshalMsgFieldSkip11zogd = []bool{false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ServiceInfo) Msgsize() (s int) {
	s = 1 + 15 + msgp.MapHeaderSize
	if z.ProcsStat != nil {
		for zxmx, zkyf := range z.ProcsStat {
			_ = zkyf
			_ = zxmx
			s += msgp.IntSize + zkyf.Msgsize()
		}
	}
	s += 15 + msgp.BoolSize + 15 + msgp.MapHeaderSize
	if z.UpStream != nil {
		for zsrz, zctb := range z.UpStream {
			_ = zctb
			_ = zsrz
			s += msgp.StringPrefixSize + len(zsrz) + 1 + 12 + msgp.IntSize
		}
	}
	s += 15 + msgp.MapHeaderSize
	if z.DownStream != nil {
		for zyos, zegv := range z.DownStream {
			_ = zegv
			_ = zyos
			s += msgp.StringPrefixSize + len(zyos) + 1 + 12 + msgp.IntSize
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
	const maxFields12zodw = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields12zodw uint32
	totalEncodedFields12zodw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft12zodw := totalEncodedFields12zodw
	missingFieldsLeft12zodw := maxFields12zodw - totalEncodedFields12zodw

	var nextMiss12zodw int = -1
	var found12zodw [maxFields12zodw]bool
	var curField12zodw int

doneWithStruct12zodw:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft12zodw > 0 || missingFieldsLeft12zodw > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zodw, missingFieldsLeft12zodw, msgp.ShowFound(found12zodw[:]), decodeMsgFieldOrder12zodw)
		if encodedFieldsLeft12zodw > 0 {
			encodedFieldsLeft12zodw--
			curField12zodw, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss12zodw < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss12zodw = 0
			}
			for nextMiss12zodw < maxFields12zodw && (found12zodw[nextMiss12zodw] || decodeMsgFieldSkip12zodw[nextMiss12zodw]) {
				nextMiss12zodw++
			}
			if nextMiss12zodw == maxFields12zodw {
				// filled all the empty fields!
				break doneWithStruct12zodw
			}
			missingFieldsLeft12zodw--
			curField12zodw = nextMiss12zodw
		}
		//fmt.Printf("switching on curField: '%v'\n", curField12zodw)
		switch curField12zodw {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Services"
			found12zodw[0] = true
			var zsna uint32
			zsna, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Services == nil && zsna > 0 {
				z.Services = make(map[string]*ServiceInfo, zsna)
			} else if len(z.Services) > 0 {
				for key, _ := range z.Services {
					delete(z.Services, key)
				}
			}
			for zsna > 0 {
				zsna--
				var zfpf string
				var zrde *ServiceInfo
				zfpf, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if zrde != nil {
						dc.PushAlwaysNil()
						err = zrde.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if zrde == nil {
						zrde = new(ServiceInfo)
					}
					err = zrde.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.Services[zfpf] = zrde
			}
		case 1:
			// zid 1 for "Time"
			found12zodw[1] = true
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
	if nextMiss12zodw != -1 {
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
var decodeMsgFieldOrder12zodw = []string{"Services", "Time"}

var decodeMsgFieldSkip12zodw = []bool{false, false}

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
	var empty_zffj [2]bool
	fieldsInUse_zuuy := z.fieldsNotEmpty(empty_zffj[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zuuy + 1)
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

	if !empty_zffj[0] {
		// zid 0 for "Services"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Services)))
		if err != nil {
			return
		}
		for zfpf, zrde := range z.Services {
			err = en.WriteString(zfpf)
			if err != nil {
				return
			}
			if zrde == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = zrde.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zffj[1] {
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
		for zfpf, zrde := range z.Services {
			o = msgp.AppendString(o, zfpf)
			if zrde == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = zrde.MarshalMsg(o)
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
	const maxFields13zlfy = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields13zlfy uint32
	if !nbs.AlwaysNil {
		totalEncodedFields13zlfy, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft13zlfy := totalEncodedFields13zlfy
	missingFieldsLeft13zlfy := maxFields13zlfy - totalEncodedFields13zlfy

	var nextMiss13zlfy int = -1
	var found13zlfy [maxFields13zlfy]bool
	var curField13zlfy int

doneWithStruct13zlfy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft13zlfy > 0 || missingFieldsLeft13zlfy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zlfy, missingFieldsLeft13zlfy, msgp.ShowFound(found13zlfy[:]), unmarshalMsgFieldOrder13zlfy)
		if encodedFieldsLeft13zlfy > 0 {
			encodedFieldsLeft13zlfy--
			curField13zlfy, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss13zlfy < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss13zlfy = 0
			}
			for nextMiss13zlfy < maxFields13zlfy && (found13zlfy[nextMiss13zlfy] || unmarshalMsgFieldSkip13zlfy[nextMiss13zlfy]) {
				nextMiss13zlfy++
			}
			if nextMiss13zlfy == maxFields13zlfy {
				// filled all the empty fields!
				break doneWithStruct13zlfy
			}
			missingFieldsLeft13zlfy--
			curField13zlfy = nextMiss13zlfy
		}
		//fmt.Printf("switching on curField: '%v'\n", curField13zlfy)
		switch curField13zlfy {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Services"
			found13zlfy[0] = true
			if nbs.AlwaysNil {
				if len(z.Services) > 0 {
					for key, _ := range z.Services {
						delete(z.Services, key)
					}
				}

			} else {

				var zxwe uint32
				zxwe, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Services == nil && zxwe > 0 {
					z.Services = make(map[string]*ServiceInfo, zxwe)
				} else if len(z.Services) > 0 {
					for key, _ := range z.Services {
						delete(z.Services, key)
					}
				}
				for zxwe > 0 {
					var zfpf string
					var zrde *ServiceInfo
					zxwe--
					zfpf, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if zrde != nil {
							zrde.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != zrde {
								zrde.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if zrde == nil {
								zrde = new(ServiceInfo)
							}
							bts, err = zrde.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Services[zfpf] = zrde
				}
			}
		case 1:
			// zid 1 for "Time"
			found13zlfy[1] = true
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
	if nextMiss13zlfy != -1 {
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
var unmarshalMsgFieldOrder13zlfy = []string{"Services", "Time"}

var unmarshalMsgFieldSkip13zlfy = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Topology) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.Services != nil {
		for zfpf, zrde := range z.Services {
			_ = zrde
			_ = zfpf
			s += msgp.StringPrefixSize + len(zfpf)
			if zrde == nil {
				s += msgp.NilSize
			} else {
				s += zrde.Msgsize()
			}
		}
	}
	s += 12 + msgp.Int64Size
	return
}

// FileSchema_go holds ZebraPack schema from file 'topo/dataExt.go'
type FileSchema_go struct{}

// ZebraSchemaInMsgpack2Format provides the ZebraPack Schema in msgpack2 format, length 2468 bytes
func (FileSchema_go) ZebraSchemaInMsgpack2Format() []byte {
	return []byte{
		0x85, 0xaa, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
		0x74, 0x68, 0xaf, 0x74, 0x6f, 0x70, 0x6f, 0x2f, 0x64, 0x61,
		0x74, 0x61, 0x45, 0x78, 0x74, 0x2e, 0x67, 0x6f, 0xad, 0x53,
		0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61,
		0x67, 0x65, 0xa4, 0x74, 0x6f, 0x70, 0x6f, 0xad, 0x5a, 0x65,
		0x62, 0x72, 0x61, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49,
		0x64, 0xd3, 0x00, 0x00, 0x4d, 0xf6, 0x15, 0x1f, 0xb4, 0x97,
		0xa7, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x85, 0xab,
		0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66,
		0x6f, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e,
		0x61, 0x6d, 0x65, 0xab, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
		0x65, 0x49, 0x6e, 0x66, 0x6f, 0xa6, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x73, 0x95, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa9, 0x50, 0x72, 0x6f, 0x63, 0x73, 0x53, 0x74, 0x61,
		0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x50, 0x72, 0x6f, 0x63, 0x73,
		0x53, 0x74, 0x61, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xb0, 0x6d, 0x61,
		0x70, 0x5b, 0x69, 0x6e, 0x74, 0x5d, 0x50, 0x72, 0x6f, 0x63,
		0x53, 0x74, 0x61, 0x74, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54,
		0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18,
		0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44,
		0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x0d, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74,
		0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x16, 0xa3, 0x53, 0x74, 0x72, 0xa8, 0x50, 0x72,
		0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x87, 0xa3, 0x5a, 0x69,
		0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f,
		0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x44, 0x6f, 0x4c, 0x69, 0x73,
		0x74, 0x65, 0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x44, 0x6f, 0x4c,
		0x69, 0x73, 0x74, 0x65, 0x6e, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa4, 0x62,
		0x6f, 0x6f, 0x6c, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74,
		0x69, 0x76, 0x65, 0x12, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x12, 0xa3, 0x53, 0x74, 0x72, 0xa4,
		0x62, 0x6f, 0x6f, 0x6c, 0x84, 0xa3, 0x5a, 0x69, 0x64, 0x02,
		0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0xa5, 0x41, 0x64, 0x64, 0x72, 0x73, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0xa5, 0x41, 0x64, 0x64, 0x72, 0x73, 0xa4, 0x53, 0x6b,
		0x69, 0x70, 0xc3, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x03, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa8, 0x55, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa8, 0x55, 0x70, 0x53, 0x74, 0x72, 0x65,
		0x61, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xb4, 0x6d, 0x61, 0x70, 0x5b,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x41, 0x64, 0x64,
		0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67,
		0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x19, 0xa3, 0x53,
		0x74, 0x72, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x86,
		0xa3, 0x5a, 0x69, 0x64, 0x04, 0xab, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x44, 0x6f,
		0x77, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0xaa, 0x44, 0x6f, 0x77, 0x6e, 0x53, 0x74, 0x72, 0x65,
		0x61, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
		0x70, 0x65, 0x53, 0x74, 0x72, 0xb4, 0x6d, 0x61, 0x70, 0x5b,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x41, 0x64, 0x64,
		0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70,
		0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b,
		0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67,
		0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x19, 0xa3, 0x53,
		0x74, 0x72, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0xa8,
		0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x82, 0xaa,
		0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
		0xa8, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0xa6,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x92, 0x86, 0xa3, 0x5a,
		0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x53, 0x65, 0x72, 0x76,
		0x69, 0x63, 0x65, 0x73, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x53, 0x65,
		0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xb7,
		0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0x5d, 0x2a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
		0x6e, 0x66, 0x6f, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43,
		0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3,
		0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f,
		0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69,
		0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x83, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x1c, 0xa3, 0x53, 0x74, 0x72, 0xa7,
		0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0xa6, 0x44, 0x6f,
		0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64,
		0x16, 0xa3, 0x53, 0x74, 0x72, 0xab, 0x53, 0x65, 0x72, 0x76,
		0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x01, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x54, 0x69, 0x6d, 0x65,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa4, 0x54, 0x69, 0x6d, 0x65, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
		0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x11, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3,
		0x53, 0x74, 0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xa4,
		0x41, 0x64, 0x64, 0x72, 0x82, 0xaa, 0x53, 0x74, 0x72, 0x75,
		0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x41, 0x64, 0x64,
		0x72, 0xa6, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x92, 0x87,
		0xa3, 0x5a, 0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x48, 0x6f,
		0x73, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61,
		0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4, 0x48, 0x6f, 0x73, 0x74,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa4, 0x50, 0x6f, 0x72, 0x74, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa4,
		0x50, 0x6f, 0x72, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74,
		0x72, 0x69, 0x6e, 0x67, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64,
		0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69,
		0x74, 0x69, 0x76, 0x65, 0x02, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82,
		0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72,
		0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa9, 0x41, 0x64,
		0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x82, 0xaa, 0x53,
		0x74, 0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa9,
		0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0xa6,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x91, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x43, 0x6f, 0x75, 0x6e,
		0x74, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x43, 0x6f, 0x75, 0x6e, 0x74,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72,
		0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x0d, 0xad, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
		0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x0d, 0xa3,
		0x53, 0x74, 0x72, 0xa3, 0x69, 0x6e, 0x74, 0xa8, 0x50, 0x72,
		0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0x82, 0xaa, 0x53, 0x74,
		0x72, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x50,
		0x72, 0x6f, 0x63, 0x53, 0x74, 0x61, 0x74, 0xa6, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x73, 0x96, 0x87, 0xa3, 0x5a, 0x69, 0x64,
		0x00, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa9, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
		0x69, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xa9, 0x53, 0x74, 0x61,
		0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa5,
		0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x17,
		0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d,
		0x69, 0x74, 0x69, 0x76, 0x65, 0x11, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x11, 0xa3, 0x53, 0x74,
		0x72, 0xa5, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x87, 0xa3, 0x5a,
		0x69, 0x64, 0x02, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47,
		0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x53, 0x74, 0x61, 0x74,
		0x65, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa5, 0x53, 0x74, 0x61, 0x74, 0x65,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x02, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72,
		0x69, 0x6e, 0x67, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x01, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa7, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x67, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xa7, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x67,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa7, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36,
		0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
		0x65, 0x04, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x04, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x66, 0x6c,
		0x6f, 0x61, 0x74, 0x36, 0x34, 0x87, 0xa3, 0x5a, 0x69, 0x64,
		0x03, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xab, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x6e,
		0x73, 0x74, 0x61, 0x6e, 0x74, 0xac, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0xab, 0x4c,
		0x6f, 0x61, 0x64, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa7, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x36,
		0x34, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76,
		0x65, 0x04, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x04, 0xa3, 0x53, 0x74, 0x72, 0xa7, 0x66, 0x6c,
		0x6f, 0x61, 0x74, 0x36, 0x34, 0x87, 0xa3, 0x5a, 0x69, 0x64,
		0x04, 0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e,
		0x61, 0x6d, 0x65, 0xa6, 0x56, 0x6d, 0x53, 0x69, 0x7a, 0x65,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e,
		0x61, 0x6d, 0x65, 0xa6, 0x56, 0x6d, 0x53, 0x69, 0x7a, 0x65,
		0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
		0x53, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34,
		0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65,
		0x67, 0x6f, 0x72, 0x79, 0x17, 0xae, 0x46, 0x69, 0x65, 0x6c,
		0x64, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
		0x0b, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c,
		0x6c, 0x54, 0x79, 0x70, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e,
		0x64, 0x0b, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e,
		0x74, 0x36, 0x34, 0x87, 0xa3, 0x5a, 0x69, 0x64, 0x05, 0xab,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d,
		0x65, 0xa5, 0x56, 0x6d, 0x52, 0x53, 0x53, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65,
		0xa5, 0x56, 0x6d, 0x52, 0x53, 0x53, 0xac, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xa6,
		0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
		0x17, 0xae, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x69,
		0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x0b, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70,
		0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x0b, 0xa3, 0x53,
		0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xa7,
		0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x90,
	}
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 3154 bytes
func (FileSchema_go) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath":"topo/dataExt.go","SourcePackage":"topo","ZebraSchemaId":85719311692951,"Structs":{"ServiceInfo":{"StructName":"ServiceInfo","Fields":[{"Zid":0,"FieldGoName":"ProcsStat","FieldTagName":"ProcsStat","FieldTypeStr":"map[int]ProcStat","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":13,"Str":"int"},"Range":{"Kind":22,"Str":"ProcStat"}}},{"Zid":1,"FieldGoName":"DoListen","FieldTagName":"DoListen","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"}},{"Zid":2,"FieldGoName":"Addrs","FieldTagName":"Addrs","Skip":true},{"Zid":3,"FieldGoName":"UpStream","FieldTagName":"UpStream","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}},{"Zid":4,"FieldGoName":"DownStream","FieldTagName":"DownStream","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}}]},"Topology":{"StructName":"Topology","Fields":[{"Zid":0,"FieldGoName":"Services","FieldTagName":"Services","FieldTypeStr":"map[string]*ServiceInfo","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"ServiceInfo"}}}},{"Zid":1,"FieldGoName":"Time","FieldTagName":"Time","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"Addr":{"StructName":"Addr","Fields":[{"Zid":0,"FieldGoName":"Host","FieldTagName":"Host","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":1,"FieldGoName":"Port","FieldTagName":"Port","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}}]},"AddrState":{"StructName":"AddrState","Fields":[{"Zid":0,"FieldGoName":"Count","FieldTagName":"Count","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}}]},"ProcStat":{"StructName":"ProcStat","Fields":[{"Zid":0,"FieldGoName":"StartTime","FieldTagName":"StartTime","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}},{"Zid":2,"FieldGoName":"State","FieldTagName":"State","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":1,"FieldGoName":"LoadAvg","FieldTagName":"LoadAvg","FieldTypeStr":"float64","FieldCategory":23,"FieldPrimitive":4,"FieldFullType":{"Kind":4,"Str":"float64"}},{"Zid":3,"FieldGoName":"LoadInstant","FieldTagName":"LoadInstant","FieldTypeStr":"float64","FieldCategory":23,"FieldPrimitive":4,"FieldFullType":{"Kind":4,"Str":"float64"}},{"Zid":4,"FieldGoName":"VmSize","FieldTagName":"VmSize","FieldTypeStr":"uint64","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}},{"Zid":5,"FieldGoName":"VmRSS","FieldTagName":"VmRSS","FieldTypeStr":"uint64","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}}]}},"Imports":[]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 8454 bytes
func (FileSchema_go) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "topo/dataExt.go",
    "SourcePackage": "topo",
    "ZebraSchemaId": 85719311692951,
    "Structs": {
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
        }
    },
    "Imports": []
}`)
}
