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
	const maxFields0zwwa = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields0zwwa uint32
	totalEncodedFields0zwwa, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft0zwwa := totalEncodedFields0zwwa
	missingFieldsLeft0zwwa := maxFields0zwwa - totalEncodedFields0zwwa

	var nextMiss0zwwa int = -1
	var found0zwwa [maxFields0zwwa]bool
	var curField0zwwa int

doneWithStruct0zwwa:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft0zwwa > 0 || missingFieldsLeft0zwwa > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft0zwwa, missingFieldsLeft0zwwa, msgp.ShowFound(found0zwwa[:]), decodeMsgFieldOrder0zwwa)
		if encodedFieldsLeft0zwwa > 0 {
			encodedFieldsLeft0zwwa--
			curField0zwwa, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss0zwwa < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss0zwwa = 0
			}
			for nextMiss0zwwa < maxFields0zwwa && (found0zwwa[nextMiss0zwwa] || decodeMsgFieldSkip0zwwa[nextMiss0zwwa]) {
				nextMiss0zwwa++
			}
			if nextMiss0zwwa == maxFields0zwwa {
				// filled all the empty fields!
				break doneWithStruct0zwwa
			}
			missingFieldsLeft0zwwa--
			curField0zwwa = nextMiss0zwwa
		}
		//fmt.Printf("switching on curField: '%v'\n", curField0zwwa)
		switch curField0zwwa {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Host"
			found0zwwa[0] = true
			z.Host, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Port"
			found0zwwa[1] = true
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
	if nextMiss0zwwa != -1 {
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
var decodeMsgFieldOrder0zwwa = []string{"Host", "Port"}

var decodeMsgFieldSkip0zwwa = []bool{false, false}

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
	var empty_ztnb [2]bool
	fieldsInUse_zqis := z.fieldsNotEmpty(empty_ztnb[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqis + 1)
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

	if !empty_ztnb[0] {
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

	if !empty_ztnb[1] {
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
	const maxFields1zlmy = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields1zlmy uint32
	if !nbs.AlwaysNil {
		totalEncodedFields1zlmy, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft1zlmy := totalEncodedFields1zlmy
	missingFieldsLeft1zlmy := maxFields1zlmy - totalEncodedFields1zlmy

	var nextMiss1zlmy int = -1
	var found1zlmy [maxFields1zlmy]bool
	var curField1zlmy int

doneWithStruct1zlmy:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft1zlmy > 0 || missingFieldsLeft1zlmy > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft1zlmy, missingFieldsLeft1zlmy, msgp.ShowFound(found1zlmy[:]), unmarshalMsgFieldOrder1zlmy)
		if encodedFieldsLeft1zlmy > 0 {
			encodedFieldsLeft1zlmy--
			curField1zlmy, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss1zlmy < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss1zlmy = 0
			}
			for nextMiss1zlmy < maxFields1zlmy && (found1zlmy[nextMiss1zlmy] || unmarshalMsgFieldSkip1zlmy[nextMiss1zlmy]) {
				nextMiss1zlmy++
			}
			if nextMiss1zlmy == maxFields1zlmy {
				// filled all the empty fields!
				break doneWithStruct1zlmy
			}
			missingFieldsLeft1zlmy--
			curField1zlmy = nextMiss1zlmy
		}
		//fmt.Printf("switching on curField: '%v'\n", curField1zlmy)
		switch curField1zlmy {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Host"
			found1zlmy[0] = true
			z.Host, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "Port"
			found1zlmy[1] = true
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
	if nextMiss1zlmy != -1 {
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
var unmarshalMsgFieldOrder1zlmy = []string{"Host", "Port"}

var unmarshalMsgFieldSkip1zlmy = []bool{false, false}

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
	const maxFields2zpqi = 1

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields2zpqi uint32
	totalEncodedFields2zpqi, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft2zpqi := totalEncodedFields2zpqi
	missingFieldsLeft2zpqi := maxFields2zpqi - totalEncodedFields2zpqi

	var nextMiss2zpqi int = -1
	var found2zpqi [maxFields2zpqi]bool
	var curField2zpqi int

doneWithStruct2zpqi:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft2zpqi > 0 || missingFieldsLeft2zpqi > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft2zpqi, missingFieldsLeft2zpqi, msgp.ShowFound(found2zpqi[:]), decodeMsgFieldOrder2zpqi)
		if encodedFieldsLeft2zpqi > 0 {
			encodedFieldsLeft2zpqi--
			curField2zpqi, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss2zpqi < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss2zpqi = 0
			}
			for nextMiss2zpqi < maxFields2zpqi && (found2zpqi[nextMiss2zpqi] || decodeMsgFieldSkip2zpqi[nextMiss2zpqi]) {
				nextMiss2zpqi++
			}
			if nextMiss2zpqi == maxFields2zpqi {
				// filled all the empty fields!
				break doneWithStruct2zpqi
			}
			missingFieldsLeft2zpqi--
			curField2zpqi = nextMiss2zpqi
		}
		//fmt.Printf("switching on curField: '%v'\n", curField2zpqi)
		switch curField2zpqi {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Count"
			found2zpqi[0] = true
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
	if nextMiss2zpqi != -1 {
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
var decodeMsgFieldOrder2zpqi = []string{"Count"}

var decodeMsgFieldSkip2zpqi = []bool{false}

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
	var empty_zspv [1]bool
	fieldsInUse_zcpf := z.fieldsNotEmpty(empty_zspv[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zcpf + 1)
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

	if !empty_zspv[0] {
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
	const maxFields3zyov = 1

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields3zyov uint32
	if !nbs.AlwaysNil {
		totalEncodedFields3zyov, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft3zyov := totalEncodedFields3zyov
	missingFieldsLeft3zyov := maxFields3zyov - totalEncodedFields3zyov

	var nextMiss3zyov int = -1
	var found3zyov [maxFields3zyov]bool
	var curField3zyov int

doneWithStruct3zyov:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft3zyov > 0 || missingFieldsLeft3zyov > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft3zyov, missingFieldsLeft3zyov, msgp.ShowFound(found3zyov[:]), unmarshalMsgFieldOrder3zyov)
		if encodedFieldsLeft3zyov > 0 {
			encodedFieldsLeft3zyov--
			curField3zyov, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss3zyov < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss3zyov = 0
			}
			for nextMiss3zyov < maxFields3zyov && (found3zyov[nextMiss3zyov] || unmarshalMsgFieldSkip3zyov[nextMiss3zyov]) {
				nextMiss3zyov++
			}
			if nextMiss3zyov == maxFields3zyov {
				// filled all the empty fields!
				break doneWithStruct3zyov
			}
			missingFieldsLeft3zyov--
			curField3zyov = nextMiss3zyov
		}
		//fmt.Printf("switching on curField: '%v'\n", curField3zyov)
		switch curField3zyov {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Count"
			found3zyov[0] = true
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
	if nextMiss3zyov != -1 {
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
var unmarshalMsgFieldOrder3zyov = []string{"Count"}

var unmarshalMsgFieldSkip3zyov = []bool{false}

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
	const maxFields4zwih = 6

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields4zwih uint32
	totalEncodedFields4zwih, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft4zwih := totalEncodedFields4zwih
	missingFieldsLeft4zwih := maxFields4zwih - totalEncodedFields4zwih

	var nextMiss4zwih int = -1
	var found4zwih [maxFields4zwih]bool
	var curField4zwih int

doneWithStruct4zwih:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft4zwih > 0 || missingFieldsLeft4zwih > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft4zwih, missingFieldsLeft4zwih, msgp.ShowFound(found4zwih[:]), decodeMsgFieldOrder4zwih)
		if encodedFieldsLeft4zwih > 0 {
			encodedFieldsLeft4zwih--
			curField4zwih, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss4zwih < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss4zwih = 0
			}
			for nextMiss4zwih < maxFields4zwih && (found4zwih[nextMiss4zwih] || decodeMsgFieldSkip4zwih[nextMiss4zwih]) {
				nextMiss4zwih++
			}
			if nextMiss4zwih == maxFields4zwih {
				// filled all the empty fields!
				break doneWithStruct4zwih
			}
			missingFieldsLeft4zwih--
			curField4zwih = nextMiss4zwih
		}
		//fmt.Printf("switching on curField: '%v'\n", curField4zwih)
		switch curField4zwih {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "StartTime"
			found4zwih[0] = true
			z.StartTime, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "State"
			found4zwih[1] = true
			z.State, err = dc.ReadString()
			if err != nil {
				return
			}
		case 1:
			// zid 1 for "LoadAvg"
			found4zwih[2] = true
			z.LoadAvg, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case 3:
			// zid 3 for "LoadInstant"
			found4zwih[3] = true
			z.LoadInstant, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case 4:
			// zid 4 for "VmSize"
			found4zwih[4] = true
			z.VmSize, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case 5:
			// zid 5 for "VmRSS"
			found4zwih[5] = true
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
	if nextMiss4zwih != -1 {
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
var decodeMsgFieldOrder4zwih = []string{"StartTime", "State", "LoadAvg", "LoadInstant", "VmSize", "VmRSS"}

var decodeMsgFieldSkip4zwih = []bool{false, false, false, false, false, false}

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
	var empty_zsze [6]bool
	fieldsInUse_zbxa := z.fieldsNotEmpty(empty_zsze[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zbxa + 1)
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

	if !empty_zsze[0] {
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

	if !empty_zsze[1] {
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

	if !empty_zsze[2] {
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

	if !empty_zsze[3] {
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

	if !empty_zsze[4] {
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

	if !empty_zsze[5] {
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
	const maxFields5zyoz = 6

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields5zyoz uint32
	if !nbs.AlwaysNil {
		totalEncodedFields5zyoz, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft5zyoz := totalEncodedFields5zyoz
	missingFieldsLeft5zyoz := maxFields5zyoz - totalEncodedFields5zyoz

	var nextMiss5zyoz int = -1
	var found5zyoz [maxFields5zyoz]bool
	var curField5zyoz int

doneWithStruct5zyoz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft5zyoz > 0 || missingFieldsLeft5zyoz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft5zyoz, missingFieldsLeft5zyoz, msgp.ShowFound(found5zyoz[:]), unmarshalMsgFieldOrder5zyoz)
		if encodedFieldsLeft5zyoz > 0 {
			encodedFieldsLeft5zyoz--
			curField5zyoz, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss5zyoz < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss5zyoz = 0
			}
			for nextMiss5zyoz < maxFields5zyoz && (found5zyoz[nextMiss5zyoz] || unmarshalMsgFieldSkip5zyoz[nextMiss5zyoz]) {
				nextMiss5zyoz++
			}
			if nextMiss5zyoz == maxFields5zyoz {
				// filled all the empty fields!
				break doneWithStruct5zyoz
			}
			missingFieldsLeft5zyoz--
			curField5zyoz = nextMiss5zyoz
		}
		//fmt.Printf("switching on curField: '%v'\n", curField5zyoz)
		switch curField5zyoz {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "StartTime"
			found5zyoz[0] = true
			z.StartTime, bts, err = nbs.ReadInt64Bytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "State"
			found5zyoz[1] = true
			z.State, bts, err = nbs.ReadStringBytes(bts)

			if err != nil {
				return
			}
		case 1:
			// zid 1 for "LoadAvg"
			found5zyoz[2] = true
			z.LoadAvg, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case 3:
			// zid 3 for "LoadInstant"
			found5zyoz[3] = true
			z.LoadInstant, bts, err = nbs.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}
		case 4:
			// zid 4 for "VmSize"
			found5zyoz[4] = true
			z.VmSize, bts, err = nbs.ReadUint64Bytes(bts)

			if err != nil {
				return
			}
		case 5:
			// zid 5 for "VmRSS"
			found5zyoz[5] = true
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
	if nextMiss5zyoz != -1 {
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
var unmarshalMsgFieldOrder5zyoz = []string{"StartTime", "State", "LoadAvg", "LoadInstant", "VmSize", "VmRSS"}

var unmarshalMsgFieldSkip5zyoz = []bool{false, false, false, false, false, false}

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
	const maxFields6zkrz = 5

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields6zkrz uint32
	totalEncodedFields6zkrz, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft6zkrz := totalEncodedFields6zkrz
	missingFieldsLeft6zkrz := maxFields6zkrz - totalEncodedFields6zkrz

	var nextMiss6zkrz int = -1
	var found6zkrz [maxFields6zkrz]bool
	var curField6zkrz int

doneWithStruct6zkrz:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft6zkrz > 0 || missingFieldsLeft6zkrz > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft6zkrz, missingFieldsLeft6zkrz, msgp.ShowFound(found6zkrz[:]), decodeMsgFieldOrder6zkrz)
		if encodedFieldsLeft6zkrz > 0 {
			encodedFieldsLeft6zkrz--
			curField6zkrz, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss6zkrz < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss6zkrz = 0
			}
			for nextMiss6zkrz < maxFields6zkrz && (found6zkrz[nextMiss6zkrz] || decodeMsgFieldSkip6zkrz[nextMiss6zkrz]) {
				nextMiss6zkrz++
			}
			if nextMiss6zkrz == maxFields6zkrz {
				// filled all the empty fields!
				break doneWithStruct6zkrz
			}
			missingFieldsLeft6zkrz--
			curField6zkrz = nextMiss6zkrz
		}
		//fmt.Printf("switching on curField: '%v'\n", curField6zkrz)
		switch curField6zkrz {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "ProcsStat"
			found6zkrz[0] = true
			var ztif uint32
			ztif, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.ProcsStat == nil && ztif > 0 {
				z.ProcsStat = make(map[int]ProcStat, ztif)
			} else if len(z.ProcsStat) > 0 {
				for key, _ := range z.ProcsStat {
					delete(z.ProcsStat, key)
				}
			}
			for ztif > 0 {
				ztif--
				var zrkb int
				var zbsw ProcStat
				zrkb, err = dc.ReadInt()
				if err != nil {
					return
				}
				err = zbsw.DecodeMsg(dc)
				if err != nil {
					return
				}
				z.ProcsStat[zrkb] = zbsw
			}
		case 1:
			// zid 1 for "DoListen"
			found6zkrz[1] = true
			z.DoListen, err = dc.ReadBool()
			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Addrs"
			found6zkrz[2] = true
			var zjna uint32
			zjna, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Addrs == nil && zjna > 0 {
				z.Addrs = make(map[string]AddrState, zjna)
			} else if len(z.Addrs) > 0 {
				for key, _ := range z.Addrs {
					delete(z.Addrs, key)
				}
			}
			for zjna > 0 {
				zjna--
				var zrmu string
				var zhwu AddrState
				zrmu, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields7zswa = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields7zswa uint32
				totalEncodedFields7zswa, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft7zswa := totalEncodedFields7zswa
				missingFieldsLeft7zswa := maxFields7zswa - totalEncodedFields7zswa

				var nextMiss7zswa int = -1
				var found7zswa [maxFields7zswa]bool
				var curField7zswa int

			doneWithStruct7zswa:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft7zswa > 0 || missingFieldsLeft7zswa > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft7zswa, missingFieldsLeft7zswa, msgp.ShowFound(found7zswa[:]), decodeMsgFieldOrder7zswa)
					if encodedFieldsLeft7zswa > 0 {
						encodedFieldsLeft7zswa--
						curField7zswa, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss7zswa < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss7zswa = 0
						}
						for nextMiss7zswa < maxFields7zswa && (found7zswa[nextMiss7zswa] || decodeMsgFieldSkip7zswa[nextMiss7zswa]) {
							nextMiss7zswa++
						}
						if nextMiss7zswa == maxFields7zswa {
							// filled all the empty fields!
							break doneWithStruct7zswa
						}
						missingFieldsLeft7zswa--
						curField7zswa = nextMiss7zswa
					}
					//fmt.Printf("switching on curField: '%v'\n", curField7zswa)
					switch curField7zswa {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found7zswa[0] = true
						zhwu.Count, err = dc.ReadInt()
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
				if nextMiss7zswa != -1 {
					dc.PopAlwaysNil()
				}

				z.Addrs[zrmu] = zhwu
			}
		case 3:
			// zid 3 for "UpStream"
			found6zkrz[3] = true
			var zsrb uint32
			zsrb, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.UpStream == nil && zsrb > 0 {
				z.UpStream = make(map[string]AddrState, zsrb)
			} else if len(z.UpStream) > 0 {
				for key, _ := range z.UpStream {
					delete(z.UpStream, key)
				}
			}
			for zsrb > 0 {
				zsrb--
				var zfzf string
				var zmre AddrState
				zfzf, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields8zyqr = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields8zyqr uint32
				totalEncodedFields8zyqr, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft8zyqr := totalEncodedFields8zyqr
				missingFieldsLeft8zyqr := maxFields8zyqr - totalEncodedFields8zyqr

				var nextMiss8zyqr int = -1
				var found8zyqr [maxFields8zyqr]bool
				var curField8zyqr int

			doneWithStruct8zyqr:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft8zyqr > 0 || missingFieldsLeft8zyqr > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft8zyqr, missingFieldsLeft8zyqr, msgp.ShowFound(found8zyqr[:]), decodeMsgFieldOrder8zyqr)
					if encodedFieldsLeft8zyqr > 0 {
						encodedFieldsLeft8zyqr--
						curField8zyqr, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss8zyqr < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss8zyqr = 0
						}
						for nextMiss8zyqr < maxFields8zyqr && (found8zyqr[nextMiss8zyqr] || decodeMsgFieldSkip8zyqr[nextMiss8zyqr]) {
							nextMiss8zyqr++
						}
						if nextMiss8zyqr == maxFields8zyqr {
							// filled all the empty fields!
							break doneWithStruct8zyqr
						}
						missingFieldsLeft8zyqr--
						curField8zyqr = nextMiss8zyqr
					}
					//fmt.Printf("switching on curField: '%v'\n", curField8zyqr)
					switch curField8zyqr {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found8zyqr[0] = true
						zmre.Count, err = dc.ReadInt()
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
				if nextMiss8zyqr != -1 {
					dc.PopAlwaysNil()
				}

				z.UpStream[zfzf] = zmre
			}
		case 4:
			// zid 4 for "DownStream"
			found6zkrz[4] = true
			var zicg uint32
			zicg, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.DownStream == nil && zicg > 0 {
				z.DownStream = make(map[string]AddrState, zicg)
			} else if len(z.DownStream) > 0 {
				for key, _ := range z.DownStream {
					delete(z.DownStream, key)
				}
			}
			for zicg > 0 {
				zicg--
				var ziip string
				var zpaq AddrState
				ziip, err = dc.ReadString()
				if err != nil {
					return
				}
				const maxFields9zjpo = 1

				// -- templateDecodeMsgZid starts here--
				var totalEncodedFields9zjpo uint32
				totalEncodedFields9zjpo, err = dc.ReadMapHeader()
				if err != nil {
					return
				}
				encodedFieldsLeft9zjpo := totalEncodedFields9zjpo
				missingFieldsLeft9zjpo := maxFields9zjpo - totalEncodedFields9zjpo

				var nextMiss9zjpo int = -1
				var found9zjpo [maxFields9zjpo]bool
				var curField9zjpo int

			doneWithStruct9zjpo:
				// First fill all the encoded fields, then
				// treat the remaining, missing fields, as Nil.
				for encodedFieldsLeft9zjpo > 0 || missingFieldsLeft9zjpo > 0 {
					//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft9zjpo, missingFieldsLeft9zjpo, msgp.ShowFound(found9zjpo[:]), decodeMsgFieldOrder9zjpo)
					if encodedFieldsLeft9zjpo > 0 {
						encodedFieldsLeft9zjpo--
						curField9zjpo, err = dc.ReadInt()
						if err != nil {
							return
						}
					} else {
						//missing fields need handling
						if nextMiss9zjpo < 0 {
							// tell the reader to only give us Nils
							// until further notice.
							dc.PushAlwaysNil()
							nextMiss9zjpo = 0
						}
						for nextMiss9zjpo < maxFields9zjpo && (found9zjpo[nextMiss9zjpo] || decodeMsgFieldSkip9zjpo[nextMiss9zjpo]) {
							nextMiss9zjpo++
						}
						if nextMiss9zjpo == maxFields9zjpo {
							// filled all the empty fields!
							break doneWithStruct9zjpo
						}
						missingFieldsLeft9zjpo--
						curField9zjpo = nextMiss9zjpo
					}
					//fmt.Printf("switching on curField: '%v'\n", curField9zjpo)
					switch curField9zjpo {
					// -- templateDecodeMsgZid ends here --

					case 0:
						// zid 0 for "Count"
						found9zjpo[0] = true
						zpaq.Count, err = dc.ReadInt()
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
				if nextMiss9zjpo != -1 {
					dc.PopAlwaysNil()
				}

				z.DownStream[ziip] = zpaq
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	if nextMiss6zkrz != -1 {
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
var decodeMsgFieldOrder6zkrz = []string{"ProcsStat", "DoListen", "Addrs", "UpStream", "DownStream"}

var decodeMsgFieldSkip6zkrz = []bool{false, false, false, false, false}

// fields of AddrState
var decodeMsgFieldOrder7zswa = []string{"Count"}

var decodeMsgFieldSkip7zswa = []bool{false}

// fields of AddrState
var decodeMsgFieldOrder8zyqr = []string{"Count"}

var decodeMsgFieldSkip8zyqr = []bool{false}

// fields of AddrState
var decodeMsgFieldOrder9zjpo = []string{"Count"}

var decodeMsgFieldSkip9zjpo = []bool{false}

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
	var empty_zbwx [5]bool
	fieldsInUse_zumy := z.fieldsNotEmpty(empty_zbwx[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zumy + 1)
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

	if !empty_zbwx[0] {
		// zid 0 for "ProcsStat"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.ProcsStat)))
		if err != nil {
			return
		}
		for zrkb, zbsw := range z.ProcsStat {
			err = en.WriteInt(zrkb)
			if err != nil {
				return
			}
			err = zbsw.EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}

	if !empty_zbwx[1] {
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

	if !empty_zbwx[2] {
		// zid 2 for "Addrs"
		err = en.Append(0x2)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Addrs)))
		if err != nil {
			return
		}
		for zrmu, zhwu := range z.Addrs {
			err = en.WriteString(zrmu)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_zusu [1]bool
			fieldsInUse_zumx := zhwu.fieldsNotEmpty(empty_zusu[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zumx + 1)
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

			if !empty_zusu[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zhwu.Count)
				if err != nil {
					return
				}
			}

		}
	}

	if !empty_zbwx[3] {
		// zid 3 for "UpStream"
		err = en.Append(0x3)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.UpStream)))
		if err != nil {
			return
		}
		for zfzf, zmre := range z.UpStream {
			err = en.WriteString(zfzf)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_zsya [1]bool
			fieldsInUse_zmlx := zmre.fieldsNotEmpty(empty_zsya[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zmlx + 1)
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

			if !empty_zsya[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zmre.Count)
				if err != nil {
					return
				}
			}

		}
	}

	if !empty_zbwx[4] {
		// zid 4 for "DownStream"
		err = en.Append(0x4)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.DownStream)))
		if err != nil {
			return
		}
		for ziip, zpaq := range z.DownStream {
			err = en.WriteString(ziip)
			if err != nil {
				return
			}

			// honor the omitempty tags
			var empty_zpkt [1]bool
			fieldsInUse_zcnh := zpaq.fieldsNotEmpty(empty_zpkt[:])

			// map header
			err = en.WriteMapHeader(fieldsInUse_zcnh + 1)
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

			if !empty_zpkt[0] {
				// zid 0 for "Count"
				err = en.Append(0x0)
				if err != nil {
					return err
				}
				err = en.WriteInt(zpaq.Count)
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
		for zrkb, zbsw := range z.ProcsStat {
			o = msgp.AppendInt(o, zrkb)
			o, err = zbsw.MarshalMsg(o)
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
		for zrmu, zhwu := range z.Addrs {
			o = msgp.AppendString(o, zrmu)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zhwu.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zhwu.Count)
			}

		}
	}

	if !empty[3] {
		// zid 3 for "UpStream"
		o = append(o, 0x3)
		o = msgp.AppendMapHeader(o, uint32(len(z.UpStream)))
		for zfzf, zmre := range z.UpStream {
			o = msgp.AppendString(o, zfzf)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zmre.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zmre.Count)
			}

		}
	}

	if !empty[4] {
		// zid 4 for "DownStream"
		o = append(o, 0x4)
		o = msgp.AppendMapHeader(o, uint32(len(z.DownStream)))
		for ziip, zpaq := range z.DownStream {
			o = msgp.AppendString(o, ziip)

			// honor the omitempty tags
			var empty [1]bool
			fieldsInUse := zpaq.fieldsNotEmpty(empty[:])
			o = msgp.AppendMapHeader(o, fieldsInUse+1)

			// runtime struct type identification for 'AddrState'
			o = msgp.AppendNegativeOneAndStringAsBytes(o, []byte{0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65})

			if !empty[0] {
				// zid 0 for "Count"
				o = append(o, 0x0)
				o = msgp.AppendInt(o, zpaq.Count)
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
	const maxFields10zjzk = 5

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields10zjzk uint32
	if !nbs.AlwaysNil {
		totalEncodedFields10zjzk, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft10zjzk := totalEncodedFields10zjzk
	missingFieldsLeft10zjzk := maxFields10zjzk - totalEncodedFields10zjzk

	var nextMiss10zjzk int = -1
	var found10zjzk [maxFields10zjzk]bool
	var curField10zjzk int

doneWithStruct10zjzk:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft10zjzk > 0 || missingFieldsLeft10zjzk > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft10zjzk, missingFieldsLeft10zjzk, msgp.ShowFound(found10zjzk[:]), unmarshalMsgFieldOrder10zjzk)
		if encodedFieldsLeft10zjzk > 0 {
			encodedFieldsLeft10zjzk--
			curField10zjzk, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss10zjzk < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss10zjzk = 0
			}
			for nextMiss10zjzk < maxFields10zjzk && (found10zjzk[nextMiss10zjzk] || unmarshalMsgFieldSkip10zjzk[nextMiss10zjzk]) {
				nextMiss10zjzk++
			}
			if nextMiss10zjzk == maxFields10zjzk {
				// filled all the empty fields!
				break doneWithStruct10zjzk
			}
			missingFieldsLeft10zjzk--
			curField10zjzk = nextMiss10zjzk
		}
		//fmt.Printf("switching on curField: '%v'\n", curField10zjzk)
		switch curField10zjzk {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "ProcsStat"
			found10zjzk[0] = true
			if nbs.AlwaysNil {
				if len(z.ProcsStat) > 0 {
					for key, _ := range z.ProcsStat {
						delete(z.ProcsStat, key)
					}
				}

			} else {

				var zcmq uint32
				zcmq, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.ProcsStat == nil && zcmq > 0 {
					z.ProcsStat = make(map[int]ProcStat, zcmq)
				} else if len(z.ProcsStat) > 0 {
					for key, _ := range z.ProcsStat {
						delete(z.ProcsStat, key)
					}
				}
				for zcmq > 0 {
					var zrkb int
					var zbsw ProcStat
					zcmq--
					zrkb, bts, err = nbs.ReadIntBytes(bts)
					if err != nil {
						return
					}
					bts, err = zbsw.UnmarshalMsg(bts)
					if err != nil {
						return
					}
					if err != nil {
						return
					}
					z.ProcsStat[zrkb] = zbsw
				}
			}
		case 1:
			// zid 1 for "DoListen"
			found10zjzk[1] = true
			z.DoListen, bts, err = nbs.ReadBoolBytes(bts)

			if err != nil {
				return
			}
		case 2:
			// zid 2 for "Addrs"
			found10zjzk[2] = true
			if nbs.AlwaysNil {
				if len(z.Addrs) > 0 {
					for key, _ := range z.Addrs {
						delete(z.Addrs, key)
					}
				}

			} else {

				var zloa uint32
				zloa, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Addrs == nil && zloa > 0 {
					z.Addrs = make(map[string]AddrState, zloa)
				} else if len(z.Addrs) > 0 {
					for key, _ := range z.Addrs {
						delete(z.Addrs, key)
					}
				}
				for zloa > 0 {
					var zrmu string
					var zhwu AddrState
					zloa--
					zrmu, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields11zryh = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields11zryh uint32
					if !nbs.AlwaysNil {
						totalEncodedFields11zryh, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft11zryh := totalEncodedFields11zryh
					missingFieldsLeft11zryh := maxFields11zryh - totalEncodedFields11zryh

					var nextMiss11zryh int = -1
					var found11zryh [maxFields11zryh]bool
					var curField11zryh int

				doneWithStruct11zryh:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft11zryh > 0 || missingFieldsLeft11zryh > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft11zryh, missingFieldsLeft11zryh, msgp.ShowFound(found11zryh[:]), unmarshalMsgFieldOrder11zryh)
						if encodedFieldsLeft11zryh > 0 {
							encodedFieldsLeft11zryh--
							curField11zryh, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss11zryh < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss11zryh = 0
							}
							for nextMiss11zryh < maxFields11zryh && (found11zryh[nextMiss11zryh] || unmarshalMsgFieldSkip11zryh[nextMiss11zryh]) {
								nextMiss11zryh++
							}
							if nextMiss11zryh == maxFields11zryh {
								// filled all the empty fields!
								break doneWithStruct11zryh
							}
							missingFieldsLeft11zryh--
							curField11zryh = nextMiss11zryh
						}
						//fmt.Printf("switching on curField: '%v'\n", curField11zryh)
						switch curField11zryh {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found11zryh[0] = true
							zhwu.Count, bts, err = nbs.ReadIntBytes(bts)

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
					if nextMiss11zryh != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.Addrs[zrmu] = zhwu
				}
			}
		case 3:
			// zid 3 for "UpStream"
			found10zjzk[3] = true
			if nbs.AlwaysNil {
				if len(z.UpStream) > 0 {
					for key, _ := range z.UpStream {
						delete(z.UpStream, key)
					}
				}

			} else {

				var zmhf uint32
				zmhf, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.UpStream == nil && zmhf > 0 {
					z.UpStream = make(map[string]AddrState, zmhf)
				} else if len(z.UpStream) > 0 {
					for key, _ := range z.UpStream {
						delete(z.UpStream, key)
					}
				}
				for zmhf > 0 {
					var zfzf string
					var zmre AddrState
					zmhf--
					zfzf, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields12zlsu = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields12zlsu uint32
					if !nbs.AlwaysNil {
						totalEncodedFields12zlsu, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft12zlsu := totalEncodedFields12zlsu
					missingFieldsLeft12zlsu := maxFields12zlsu - totalEncodedFields12zlsu

					var nextMiss12zlsu int = -1
					var found12zlsu [maxFields12zlsu]bool
					var curField12zlsu int

				doneWithStruct12zlsu:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft12zlsu > 0 || missingFieldsLeft12zlsu > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft12zlsu, missingFieldsLeft12zlsu, msgp.ShowFound(found12zlsu[:]), unmarshalMsgFieldOrder12zlsu)
						if encodedFieldsLeft12zlsu > 0 {
							encodedFieldsLeft12zlsu--
							curField12zlsu, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss12zlsu < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss12zlsu = 0
							}
							for nextMiss12zlsu < maxFields12zlsu && (found12zlsu[nextMiss12zlsu] || unmarshalMsgFieldSkip12zlsu[nextMiss12zlsu]) {
								nextMiss12zlsu++
							}
							if nextMiss12zlsu == maxFields12zlsu {
								// filled all the empty fields!
								break doneWithStruct12zlsu
							}
							missingFieldsLeft12zlsu--
							curField12zlsu = nextMiss12zlsu
						}
						//fmt.Printf("switching on curField: '%v'\n", curField12zlsu)
						switch curField12zlsu {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found12zlsu[0] = true
							zmre.Count, bts, err = nbs.ReadIntBytes(bts)

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
					if nextMiss12zlsu != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.UpStream[zfzf] = zmre
				}
			}
		case 4:
			// zid 4 for "DownStream"
			found10zjzk[4] = true
			if nbs.AlwaysNil {
				if len(z.DownStream) > 0 {
					for key, _ := range z.DownStream {
						delete(z.DownStream, key)
					}
				}

			} else {

				var zuze uint32
				zuze, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.DownStream == nil && zuze > 0 {
					z.DownStream = make(map[string]AddrState, zuze)
				} else if len(z.DownStream) > 0 {
					for key, _ := range z.DownStream {
						delete(z.DownStream, key)
					}
				}
				for zuze > 0 {
					var ziip string
					var zpaq AddrState
					zuze--
					ziip, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					const maxFields13zvjr = 1

					// -- templateUnmarshalMsgZid starts here--
					var totalEncodedFields13zvjr uint32
					if !nbs.AlwaysNil {
						totalEncodedFields13zvjr, bts, err = nbs.ReadMapHeaderBytes(bts)
						if err != nil {
							return
						}
					}
					encodedFieldsLeft13zvjr := totalEncodedFields13zvjr
					missingFieldsLeft13zvjr := maxFields13zvjr - totalEncodedFields13zvjr

					var nextMiss13zvjr int = -1
					var found13zvjr [maxFields13zvjr]bool
					var curField13zvjr int

				doneWithStruct13zvjr:
					// First fill all the encoded fields, then
					// treat the remaining, missing fields, as Nil.
					for encodedFieldsLeft13zvjr > 0 || missingFieldsLeft13zvjr > 0 {
						//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft13zvjr, missingFieldsLeft13zvjr, msgp.ShowFound(found13zvjr[:]), unmarshalMsgFieldOrder13zvjr)
						if encodedFieldsLeft13zvjr > 0 {
							encodedFieldsLeft13zvjr--
							curField13zvjr, bts, err = nbs.ReadIntBytes(bts)
							if err != nil {
								return
							}
						} else {
							//missing fields need handling
							if nextMiss13zvjr < 0 {
								// set bts to contain just mnil (0xc0)
								bts = nbs.PushAlwaysNil(bts)
								nextMiss13zvjr = 0
							}
							for nextMiss13zvjr < maxFields13zvjr && (found13zvjr[nextMiss13zvjr] || unmarshalMsgFieldSkip13zvjr[nextMiss13zvjr]) {
								nextMiss13zvjr++
							}
							if nextMiss13zvjr == maxFields13zvjr {
								// filled all the empty fields!
								break doneWithStruct13zvjr
							}
							missingFieldsLeft13zvjr--
							curField13zvjr = nextMiss13zvjr
						}
						//fmt.Printf("switching on curField: '%v'\n", curField13zvjr)
						switch curField13zvjr {
						// -- templateUnmarshalMsgZid ends here --

						case 0:
							// zid 0 for "Count"
							found13zvjr[0] = true
							zpaq.Count, bts, err = nbs.ReadIntBytes(bts)

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
					if nextMiss13zvjr != -1 {
						bts = nbs.PopAlwaysNil()
					}

					z.DownStream[ziip] = zpaq
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	if nextMiss10zjzk != -1 {
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
var unmarshalMsgFieldOrder10zjzk = []string{"ProcsStat", "DoListen", "Addrs", "UpStream", "DownStream"}

var unmarshalMsgFieldSkip10zjzk = []bool{false, false, false, false, false}

// fields of AddrState
var unmarshalMsgFieldOrder11zryh = []string{"Count"}

var unmarshalMsgFieldSkip11zryh = []bool{false}

// fields of AddrState
var unmarshalMsgFieldOrder12zlsu = []string{"Count"}

var unmarshalMsgFieldSkip12zlsu = []bool{false}

// fields of AddrState
var unmarshalMsgFieldOrder13zvjr = []string{"Count"}

var unmarshalMsgFieldSkip13zvjr = []bool{false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ServiceInfo) Msgsize() (s int) {
	s = 1 + 15 + msgp.MapHeaderSize
	if z.ProcsStat != nil {
		for zrkb, zbsw := range z.ProcsStat {
			_ = zbsw
			_ = zrkb
			s += msgp.IntSize + zbsw.Msgsize()
		}
	}
	s += 15 + msgp.BoolSize + 15 + msgp.MapHeaderSize
	if z.Addrs != nil {
		for zrmu, zhwu := range z.Addrs {
			_ = zhwu
			_ = zrmu
			s += msgp.StringPrefixSize + len(zrmu) + 1 + 12 + msgp.IntSize
		}
	}
	s += 15 + msgp.MapHeaderSize
	if z.UpStream != nil {
		for zfzf, zmre := range z.UpStream {
			_ = zmre
			_ = zfzf
			s += msgp.StringPrefixSize + len(zfzf) + 1 + 12 + msgp.IntSize
		}
	}
	s += 15 + msgp.MapHeaderSize
	if z.DownStream != nil {
		for ziip, zpaq := range z.DownStream {
			_ = zpaq
			_ = ziip
			s += msgp.StringPrefixSize + len(ziip) + 1 + 12 + msgp.IntSize
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
	const maxFields14zkgl = 2

	// -- templateDecodeMsgZid starts here--
	var totalEncodedFields14zkgl uint32
	totalEncodedFields14zkgl, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	encodedFieldsLeft14zkgl := totalEncodedFields14zkgl
	missingFieldsLeft14zkgl := maxFields14zkgl - totalEncodedFields14zkgl

	var nextMiss14zkgl int = -1
	var found14zkgl [maxFields14zkgl]bool
	var curField14zkgl int

doneWithStruct14zkgl:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft14zkgl > 0 || missingFieldsLeft14zkgl > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft14zkgl, missingFieldsLeft14zkgl, msgp.ShowFound(found14zkgl[:]), decodeMsgFieldOrder14zkgl)
		if encodedFieldsLeft14zkgl > 0 {
			encodedFieldsLeft14zkgl--
			curField14zkgl, err = dc.ReadInt()
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss14zkgl < 0 {
				// tell the reader to only give us Nils
				// until further notice.
				dc.PushAlwaysNil()
				nextMiss14zkgl = 0
			}
			for nextMiss14zkgl < maxFields14zkgl && (found14zkgl[nextMiss14zkgl] || decodeMsgFieldSkip14zkgl[nextMiss14zkgl]) {
				nextMiss14zkgl++
			}
			if nextMiss14zkgl == maxFields14zkgl {
				// filled all the empty fields!
				break doneWithStruct14zkgl
			}
			missingFieldsLeft14zkgl--
			curField14zkgl = nextMiss14zkgl
		}
		//fmt.Printf("switching on curField: '%v'\n", curField14zkgl)
		switch curField14zkgl {
		// -- templateDecodeMsgZid ends here --

		case 0:
			// zid 0 for "Services"
			found14zkgl[0] = true
			var zutl uint32
			zutl, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.Services == nil && zutl > 0 {
				z.Services = make(map[string]*ServiceInfo, zutl)
			} else if len(z.Services) > 0 {
				for key, _ := range z.Services {
					delete(z.Services, key)
				}
			}
			for zutl > 0 {
				zutl--
				var zgxc string
				var ziuo *ServiceInfo
				zgxc, err = dc.ReadString()
				if err != nil {
					return
				}
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}

					if ziuo != nil {
						dc.PushAlwaysNil()
						err = ziuo.DecodeMsg(dc)
						if err != nil {
							return
						}
						dc.PopAlwaysNil()
					}
				} else {
					// not Nil, we have something to read

					if ziuo == nil {
						ziuo = new(ServiceInfo)
					}
					err = ziuo.DecodeMsg(dc)
					if err != nil {
						return
					}
				}
				z.Services[zgxc] = ziuo
			}
		case 1:
			// zid 1 for "Time"
			found14zkgl[1] = true
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
	if nextMiss14zkgl != -1 {
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
var decodeMsgFieldOrder14zkgl = []string{"Services", "Time"}

var decodeMsgFieldSkip14zkgl = []bool{false, false}

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
	var empty_zujc [2]bool
	fieldsInUse_zqyr := z.fieldsNotEmpty(empty_zujc[:])

	// map header
	err = en.WriteMapHeader(fieldsInUse_zqyr + 1)
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

	if !empty_zujc[0] {
		// zid 0 for "Services"
		err = en.Append(0x0)
		if err != nil {
			return err
		}
		err = en.WriteMapHeader(uint32(len(z.Services)))
		if err != nil {
			return
		}
		for zgxc, ziuo := range z.Services {
			err = en.WriteString(zgxc)
			if err != nil {
				return
			}
			if ziuo == nil {
				err = en.WriteNil()
				if err != nil {
					return
				}
			} else {
				err = ziuo.EncodeMsg(en)
				if err != nil {
					return
				}
			}
		}
	}

	if !empty_zujc[1] {
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
		for zgxc, ziuo := range z.Services {
			o = msgp.AppendString(o, zgxc)
			if ziuo == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = ziuo.MarshalMsg(o)
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
	const maxFields15zzms = 2

	// -- templateUnmarshalMsgZid starts here--
	var totalEncodedFields15zzms uint32
	if !nbs.AlwaysNil {
		totalEncodedFields15zzms, bts, err = nbs.ReadMapHeaderBytes(bts)
		if err != nil {
			return
		}
	}
	encodedFieldsLeft15zzms := totalEncodedFields15zzms
	missingFieldsLeft15zzms := maxFields15zzms - totalEncodedFields15zzms

	var nextMiss15zzms int = -1
	var found15zzms [maxFields15zzms]bool
	var curField15zzms int

doneWithStruct15zzms:
	// First fill all the encoded fields, then
	// treat the remaining, missing fields, as Nil.
	for encodedFieldsLeft15zzms > 0 || missingFieldsLeft15zzms > 0 {
		//fmt.Printf("encodedFieldsLeft: %v, missingFieldsLeft: %v, found: '%v', fields: '%#v'\n", encodedFieldsLeft15zzms, missingFieldsLeft15zzms, msgp.ShowFound(found15zzms[:]), unmarshalMsgFieldOrder15zzms)
		if encodedFieldsLeft15zzms > 0 {
			encodedFieldsLeft15zzms--
			curField15zzms, bts, err = nbs.ReadIntBytes(bts)
			if err != nil {
				return
			}
		} else {
			//missing fields need handling
			if nextMiss15zzms < 0 {
				// set bts to contain just mnil (0xc0)
				bts = nbs.PushAlwaysNil(bts)
				nextMiss15zzms = 0
			}
			for nextMiss15zzms < maxFields15zzms && (found15zzms[nextMiss15zzms] || unmarshalMsgFieldSkip15zzms[nextMiss15zzms]) {
				nextMiss15zzms++
			}
			if nextMiss15zzms == maxFields15zzms {
				// filled all the empty fields!
				break doneWithStruct15zzms
			}
			missingFieldsLeft15zzms--
			curField15zzms = nextMiss15zzms
		}
		//fmt.Printf("switching on curField: '%v'\n", curField15zzms)
		switch curField15zzms {
		// -- templateUnmarshalMsgZid ends here --

		case 0:
			// zid 0 for "Services"
			found15zzms[0] = true
			if nbs.AlwaysNil {
				if len(z.Services) > 0 {
					for key, _ := range z.Services {
						delete(z.Services, key)
					}
				}

			} else {

				var zcps uint32
				zcps, bts, err = nbs.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				if z.Services == nil && zcps > 0 {
					z.Services = make(map[string]*ServiceInfo, zcps)
				} else if len(z.Services) > 0 {
					for key, _ := range z.Services {
						delete(z.Services, key)
					}
				}
				for zcps > 0 {
					var zgxc string
					var ziuo *ServiceInfo
					zcps--
					zgxc, bts, err = nbs.ReadStringBytes(bts)
					if err != nil {
						return
					}
					if nbs.AlwaysNil {
						if ziuo != nil {
							ziuo.UnmarshalMsg(msgp.OnlyNilSlice)
						}
					} else {
						// not nbs.AlwaysNil
						if msgp.IsNil(bts) {
							bts = bts[1:]
							if nil != ziuo {
								ziuo.UnmarshalMsg(msgp.OnlyNilSlice)
							}
						} else {
							// not nbs.AlwaysNil and not IsNil(bts): have something to read

							if ziuo == nil {
								ziuo = new(ServiceInfo)
							}
							bts, err = ziuo.UnmarshalMsg(bts)
							if err != nil {
								return
							}
							if err != nil {
								return
							}
						}
					}
					z.Services[zgxc] = ziuo
				}
			}
		case 1:
			// zid 1 for "Time"
			found15zzms[1] = true
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
	if nextMiss15zzms != -1 {
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
var unmarshalMsgFieldOrder15zzms = []string{"Services", "Time"}

var unmarshalMsgFieldSkip15zzms = []bool{false, false}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Topology) Msgsize() (s int) {
	s = 1 + 12 + msgp.MapHeaderSize
	if z.Services != nil {
		for zgxc, ziuo := range z.Services {
			_ = ziuo
			_ = zgxc
			s += msgp.StringPrefixSize + len(zgxc)
			if ziuo == nil {
				s += msgp.NilSize
			} else {
				s += ziuo.Msgsize()
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
		0x74, 0x72, 0xa6, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0xab,
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
		0x62, 0x6f, 0x6f, 0x6c, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x02,
		0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0xa5, 0x41, 0x64, 0x64, 0x72, 0x73, 0xac, 0x46,
		0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61, 0x6d,
		0x65, 0xa5, 0x41, 0x64, 0x64, 0x72, 0x73, 0xac, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
		0xb4, 0x6d, 0x61, 0x70, 0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e,
		0x67, 0x5d, 0x41, 0x64, 0x64, 0x72, 0x53, 0x74, 0x61, 0x74,
		0x65, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x61, 0x74,
		0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0xad, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
		0x84, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74,
		0x72, 0xa3, 0x4d, 0x61, 0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61,
		0x69, 0x6e, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3,
		0x53, 0x74, 0x72, 0xa6, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
		0xa5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x19, 0xa3, 0x53, 0x74, 0x72, 0xa6, 0x53, 0x74,
		0x72, 0x75, 0x63, 0x74, 0x86, 0xa3, 0x5a, 0x69, 0x64, 0x03,
		0xab, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61,
		0x6d, 0x65, 0xa8, 0x55, 0x70, 0x53, 0x74, 0x72, 0x65, 0x61,
		0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67,
		0x4e, 0x61, 0x6d, 0x65, 0xa8, 0x55, 0x70, 0x53, 0x74, 0x72,
		0x65, 0x61, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xb4, 0x6d, 0x61, 0x70,
		0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x41, 0x64,
		0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61,
		0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e,
		0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x19, 0xa3,
		0x53, 0x74, 0x72, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
		0x86, 0xa3, 0x5a, 0x69, 0x64, 0x04, 0xab, 0x46, 0x69, 0x65,
		0x6c, 0x64, 0x47, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0xaa, 0x44,
		0x6f, 0x77, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xac,
		0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x61, 0x67, 0x4e, 0x61,
		0x6d, 0x65, 0xaa, 0x44, 0x6f, 0x77, 0x6e, 0x53, 0x74, 0x72,
		0x65, 0x61, 0x6d, 0xac, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54,
		0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0xb4, 0x6d, 0x61, 0x70,
		0x5b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5d, 0x41, 0x64,
		0x64, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0xad, 0x46, 0x69,
		0x65, 0x6c, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
		0x79, 0x18, 0xad, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x75,
		0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x84, 0xa4, 0x4b, 0x69,
		0x6e, 0x64, 0x18, 0xa3, 0x53, 0x74, 0x72, 0xa3, 0x4d, 0x61,
		0x70, 0xa6, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x82, 0xa4,
		0x4b, 0x69, 0x6e, 0x64, 0x02, 0xa3, 0x53, 0x74, 0x72, 0xa6,
		0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0xa5, 0x52, 0x61, 0x6e,
		0x67, 0x65, 0x82, 0xa4, 0x4b, 0x69, 0x6e, 0x64, 0x19, 0xa3,
		0x53, 0x74, 0x72, 0xa6, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
		0xa7, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x90,
	}
}

// ZebraSchemaInJsonCompact provides the ZebraPack Schema in compact JSON format, length 3309 bytes
func (FileSchema_go) ZebraSchemaInJsonCompact() []byte {
	return []byte(`{"SourcePath":"topo/dataExt.go","SourcePackage":"topo","ZebraSchemaId":85719311692951,"Structs":{"Topology":{"StructName":"Topology","Fields":[{"Zid":0,"FieldGoName":"Services","FieldTagName":"Services","FieldTypeStr":"map[string]*ServiceInfo","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":28,"Str":"Pointer","Domain":{"Kind":22,"Str":"ServiceInfo"}}}},{"Zid":1,"FieldGoName":"Time","FieldTagName":"Time","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}}]},"Addr":{"StructName":"Addr","Fields":[{"Zid":0,"FieldGoName":"Host","FieldTagName":"Host","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":1,"FieldGoName":"Port","FieldTagName":"Port","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}}]},"AddrState":{"StructName":"AddrState","Fields":[{"Zid":0,"FieldGoName":"Count","FieldTagName":"Count","FieldTypeStr":"int","FieldCategory":23,"FieldPrimitive":13,"FieldFullType":{"Kind":13,"Str":"int"}}]},"ProcStat":{"StructName":"ProcStat","Fields":[{"Zid":0,"FieldGoName":"StartTime","FieldTagName":"StartTime","FieldTypeStr":"int64","FieldCategory":23,"FieldPrimitive":17,"FieldFullType":{"Kind":17,"Str":"int64"}},{"Zid":2,"FieldGoName":"State","FieldTagName":"State","FieldTypeStr":"string","FieldCategory":23,"FieldPrimitive":2,"FieldFullType":{"Kind":2,"Str":"string"}},{"Zid":1,"FieldGoName":"LoadAvg","FieldTagName":"LoadAvg","FieldTypeStr":"float64","FieldCategory":23,"FieldPrimitive":4,"FieldFullType":{"Kind":4,"Str":"float64"}},{"Zid":3,"FieldGoName":"LoadInstant","FieldTagName":"LoadInstant","FieldTypeStr":"float64","FieldCategory":23,"FieldPrimitive":4,"FieldFullType":{"Kind":4,"Str":"float64"}},{"Zid":4,"FieldGoName":"VmSize","FieldTagName":"VmSize","FieldTypeStr":"uint64","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}},{"Zid":5,"FieldGoName":"VmRSS","FieldTagName":"VmRSS","FieldTypeStr":"uint64","FieldCategory":23,"FieldPrimitive":11,"FieldFullType":{"Kind":11,"Str":"uint64"}}]},"ServiceInfo":{"StructName":"ServiceInfo","Fields":[{"Zid":0,"FieldGoName":"ProcsStat","FieldTagName":"ProcsStat","FieldTypeStr":"map[int]ProcStat","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":13,"Str":"int"},"Range":{"Kind":22,"Str":"ProcStat"}}},{"Zid":1,"FieldGoName":"DoListen","FieldTagName":"DoListen","FieldTypeStr":"bool","FieldCategory":23,"FieldPrimitive":18,"FieldFullType":{"Kind":18,"Str":"bool"}},{"Zid":2,"FieldGoName":"Addrs","FieldTagName":"Addrs","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}},{"Zid":3,"FieldGoName":"UpStream","FieldTagName":"UpStream","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}},{"Zid":4,"FieldGoName":"DownStream","FieldTagName":"DownStream","FieldTypeStr":"map[string]AddrState","FieldCategory":24,"FieldFullType":{"Kind":24,"Str":"Map","Domain":{"Kind":2,"Str":"string"},"Range":{"Kind":25,"Str":"Struct"}}}]}},"Imports":[]}`)
}

// ZebraSchemaInJsonPretty provides the ZebraPack Schema in pretty JSON format, length 8948 bytes
func (FileSchema_go) ZebraSchemaInJsonPretty() []byte {
	return []byte(`{
    "SourcePath": "topo/dataExt.go",
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
        }
    },
    "Imports": []
}`)
}
