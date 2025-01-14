// generated by read_mkauto.lua, DO NOT EDIT
package convert

import (
	"github.com/anon55555/mt"
	"github.com/yuin/gopher-lua"
)

func readAOID(l *lua.LState, val lua.LValue, ptr *mt.AOID) {
	if val.Type() != lua.LTNumber {
		panic("invalid value for AOID: must be a number")
	}
	*ptr = mt.AOID(val.(lua.LNumber))
}

func readCompressionModes(l *lua.LState, val lua.LValue, ptr *mt.CompressionModes) {
	if val.Type() != lua.LTNumber {
		panic("invalid value for CompressionModes: must be a number")
	}
	*ptr = mt.CompressionModes(val.(lua.LNumber))
}

func readInt16(l *lua.LState, val lua.LValue, ptr *int16) {
	if val.Type() != lua.LTNumber {
		panic("invalid value for int16: must be a number")
	}
	*ptr = int16(val.(lua.LNumber))
}

func readInt32(l *lua.LState, val lua.LValue, ptr *int32) {
	if val.Type() != lua.LTNumber {
		panic("invalid value for int32: must be a number")
	}
	*ptr = int32(val.(lua.LNumber))
}

func readInteraction(l *lua.LState, val lua.LValue, ptr *mt.Interaction) {
	if val.Type() != lua.LTString {
		panic("invalid value for Interaction: must be a string")
	}
	str := string(val.(lua.LString))
	switch str {
	case "dig":
		*ptr = mt.Dig
	case "stop_digging":
		*ptr = mt.StopDigging
	case "dug":
		*ptr = mt.Dug
	case "place":
		*ptr = mt.Place
	case "use":
		*ptr = mt.Use
	case "activate":
		*ptr = mt.Activate
	default:
		panic("invalid value for interaction: " + str)
	}
}

func readKeys(l *lua.LState, val lua.LValue, ptr *mt.Keys) {
	if val.Type() != lua.LTTable {
		panic("invalid value for Keys: must be a table")
	}
	if l.GetField(val, "forward") == lua.LTrue {
		*ptr = *ptr | mt.ForwardKey
	}
	if l.GetField(val, "backward") == lua.LTrue {
		*ptr = *ptr | mt.BackwardKey
	}
	if l.GetField(val, "left") == lua.LTrue {
		*ptr = *ptr | mt.LeftKey
	}
	if l.GetField(val, "right") == lua.LTrue {
		*ptr = *ptr | mt.RightKey
	}
	if l.GetField(val, "jump") == lua.LTrue {
		*ptr = *ptr | mt.JumpKey
	}
	if l.GetField(val, "special") == lua.LTrue {
		*ptr = *ptr | mt.SpecialKey
	}
	if l.GetField(val, "sneak") == lua.LTrue {
		*ptr = *ptr | mt.SneakKey
	}
	if l.GetField(val, "dig") == lua.LTrue {
		*ptr = *ptr | mt.DigKey
	}
	if l.GetField(val, "place") == lua.LTrue {
		*ptr = *ptr | mt.PlaceKey
	}
	if l.GetField(val, "zoom") == lua.LTrue {
		*ptr = *ptr | mt.ZoomKey
	}
}

func readPlayerPos(l *lua.LState, val lua.LValue, ptr *mt.PlayerPos) {
	if val.Type() != lua.LTTable {
		panic("invalid value for PlayerPos: must be a table")
	}
	readUint8(l, l.GetField(val, "fov80"), &ptr.FOV80)
	readKeys(l, l.GetField(val, "keys"), &ptr.Keys)
	readInt32(l, l.GetField(val, "pitch100"), &ptr.Pitch100)
	readVec3Int32(l, l.GetField(val, "pos100"), &ptr.Pos100)
	readVec3Int32(l, l.GetField(val, "vel100"), &ptr.Vel100)
	readUint8(l, l.GetField(val, "wanted_range"), &ptr.WantedRange)
	readInt32(l, l.GetField(val, "yaw100"), &ptr.Yaw100)
}

func readSliceSoundID(l *lua.LState, val lua.LValue, ptr *[]mt.SoundID) {
	if val.Type() != lua.LTTable {
		panic("invalid value for []SoundID: must be a table")
	}
	tbl := val.(*lua.LTable)
	n := tbl.MaxN()
	*ptr = make([]mt.SoundID, n)
	for i := range *ptr {
		readSoundID(l, l.RawGetInt(tbl, i+1), &(*ptr)[i])
	}
}

func readSliceString(l *lua.LState, val lua.LValue, ptr *[]string) {
	if val.Type() != lua.LTTable {
		panic("invalid value for []string: must be a table")
	}
	tbl := val.(*lua.LTable)
	n := tbl.MaxN()
	*ptr = make([]string, n)
	for i := range *ptr {
		readString(l, l.RawGetInt(tbl, i+1), &(*ptr)[i])
	}
}

func readSliceVec3Int16(l *lua.LState, val lua.LValue, ptr *[][3]int16) {
	if val.Type() != lua.LTTable {
		panic("invalid value for [][3]int16: must be a table")
	}
	tbl := val.(*lua.LTable)
	n := tbl.MaxN()
	*ptr = make([][3]int16, n)
	for i := range *ptr {
		readVec3Int16(l, l.RawGetInt(tbl, i+1), &(*ptr)[i])
	}
}

func readSoundID(l *lua.LState, val lua.LValue, ptr *mt.SoundID) {
	if val.Type() != lua.LTNumber {
		panic("invalid value for SoundID: must be a number")
	}
	*ptr = mt.SoundID(val.(lua.LNumber))
}

func readUint16(l *lua.LState, val lua.LValue, ptr *uint16) {
	if val.Type() != lua.LTNumber {
		panic("invalid value for uint16: must be a number")
	}
	*ptr = uint16(val.(lua.LNumber))
}

func readUint8(l *lua.LState, val lua.LValue, ptr *uint8) {
	if val.Type() != lua.LTNumber {
		panic("invalid value for uint8: must be a number")
	}
	*ptr = uint8(val.(lua.LNumber))
}

func readVec3Int16(l *lua.LState, val lua.LValue, ptr *[3]int16) {
	if val.Type() != lua.LTTable {
		panic("invalid value for [3]int16: must be a table")
	}
	readInt16(l, l.GetField(val, "x"), &(*ptr)[0])
	readInt16(l, l.GetField(val, "y"), &(*ptr)[1])
	readInt16(l, l.GetField(val, "z"), &(*ptr)[2])
}

func readVec3Int32(l *lua.LState, val lua.LValue, ptr *[3]int32) {
	if val.Type() != lua.LTTable {
		panic("invalid value for [3]int32: must be a table")
	}
	readInt32(l, l.GetField(val, "x"), &(*ptr)[0])
	readInt32(l, l.GetField(val, "y"), &(*ptr)[1])
	readInt32(l, l.GetField(val, "z"), &(*ptr)[2])
}

func ReadCmd(l *lua.LState) mt.Cmd {
	str := l.CheckString(2)
	switch str {
	case "chat_msg":
		ptr := &mt.ToSrvChatMsg{}
		val := l.CheckTable(3)
		readString(l, l.GetField(val, "msg"), &ptr.Msg)
		return ptr
	case "clt_ready":
		ptr := &mt.ToSrvCltReady{}
		val := l.CheckTable(3)
		readUint16(l, l.GetField(val, "formspec"), &ptr.Formspec)
		readUint8(l, l.GetField(val, "major"), &ptr.Major)
		readUint8(l, l.GetField(val, "minor"), &ptr.Minor)
		readUint8(l, l.GetField(val, "patch"), &ptr.Patch)
		readString(l, l.GetField(val, "version"), &ptr.Version)
		return ptr
	case "deleted_blks":
		ptr := &mt.ToSrvDeletedBlks{}
		val := l.CheckTable(3)
		readSliceVec3Int16(l, l.GetField(val, "blks"), &ptr.Blks)
		return ptr
	case "fall_dmg":
		ptr := &mt.ToSrvFallDmg{}
		val := l.CheckTable(3)
		readUint16(l, l.GetField(val, "amount"), &ptr.Amount)
		return ptr
	case "first_srp":
		ptr := &mt.ToSrvFirstSRP{}
		val := l.CheckTable(3)
		readBool(l, l.GetField(val, "empty_passwd"), &ptr.EmptyPasswd)
		readSliceByte(l, l.GetField(val, "salt"), &ptr.Salt)
		readSliceByte(l, l.GetField(val, "verifier"), &ptr.Verifier)
		return ptr
	case "got_blks":
		ptr := &mt.ToSrvGotBlks{}
		val := l.CheckTable(3)
		readSliceVec3Int16(l, l.GetField(val, "blks"), &ptr.Blks)
		return ptr
	case "init":
		ptr := &mt.ToSrvInit{}
		val := l.CheckTable(3)
		readUint16(l, l.GetField(val, "max_proto_ver"), &ptr.MaxProtoVer)
		readUint16(l, l.GetField(val, "min_proto_ver"), &ptr.MinProtoVer)
		readString(l, l.GetField(val, "player_name"), &ptr.PlayerName)
		readBool(l, l.GetField(val, "send_full_item_meta"), &ptr.SendFullItemMeta)
		readUint8(l, l.GetField(val, "serialize_ver"), &ptr.SerializeVer)
		readCompressionModes(l, l.GetField(val, "supported_compression"), &ptr.SupportedCompression)
		return ptr
	case "init2":
		ptr := &mt.ToSrvInit2{}
		val := l.CheckTable(3)
		readString(l, l.GetField(val, "lang"), &ptr.Lang)
		return ptr
	case "interact":
		ptr := &mt.ToSrvInteract{}
		val := l.CheckTable(3)
		readInteraction(l, l.GetField(val, "action"), &ptr.Action)
		readUint16(l, l.GetField(val, "item_slot"), &ptr.ItemSlot)
		readPointedThing(l, l.GetField(val, "pointed"), &ptr.Pointed)
		readPlayerPos(l, l.GetField(val, "pos"), &ptr.Pos)
		return ptr
	case "inv_action":
		ptr := &mt.ToSrvInvAction{}
		val := l.CheckTable(3)
		readString(l, l.GetField(val, "action"), &ptr.Action)
		return ptr
	case "inv_fields":
		ptr := &mt.ToSrvInvFields{}
		val := l.CheckTable(3)
		readSliceField(l, l.GetField(val, "fields"), &ptr.Fields)
		readString(l, l.GetField(val, "formname"), &ptr.Formname)
		return ptr
	case "join_mod_chan":
		ptr := &mt.ToSrvJoinModChan{}
		val := l.CheckTable(3)
		readString(l, l.GetField(val, "channel"), &ptr.Channel)
		return ptr
	case "leave_mod_chan":
		ptr := &mt.ToSrvLeaveModChan{}
		val := l.CheckTable(3)
		readString(l, l.GetField(val, "channel"), &ptr.Channel)
		return ptr
	case "msg_mod_chan":
		ptr := &mt.ToSrvMsgModChan{}
		val := l.CheckTable(3)
		readString(l, l.GetField(val, "channel"), &ptr.Channel)
		readString(l, l.GetField(val, "msg"), &ptr.Msg)
		return ptr
	case "nil":
		ptr := &mt.ToSrvNil{}
		return ptr
	case "node_meta_fields":
		ptr := &mt.ToSrvNodeMetaFields{}
		val := l.CheckTable(3)
		readSliceField(l, l.GetField(val, "fields"), &ptr.Fields)
		readString(l, l.GetField(val, "formname"), &ptr.Formname)
		readVec3Int16(l, l.GetField(val, "pos"), &ptr.Pos)
		return ptr
	case "player_pos":
		ptr := &mt.ToSrvPlayerPos{}
		val := l.CheckTable(3)
		readPlayerPos(l, l.GetField(val, "pos"), &ptr.Pos)
		return ptr
	case "removed_sounds":
		ptr := &mt.ToSrvRemovedSounds{}
		val := l.CheckTable(3)
		readSliceSoundID(l, l.GetField(val, "ids"), &ptr.IDs)
		return ptr
	case "req_media":
		ptr := &mt.ToSrvReqMedia{}
		val := l.CheckTable(3)
		readSliceString(l, l.GetField(val, "filenames"), &ptr.Filenames)
		return ptr
	case "respawn":
		ptr := &mt.ToSrvRespawn{}
		return ptr
	case "select_item":
		ptr := &mt.ToSrvSelectItem{}
		val := l.CheckTable(3)
		readUint16(l, l.GetField(val, "slot"), &ptr.Slot)
		return ptr
	case "srp_bytes_a":
		ptr := &mt.ToSrvSRPBytesA{}
		val := l.CheckTable(3)
		readSliceByte(l, l.GetField(val, "a"), &ptr.A)
		readBool(l, l.GetField(val, "no_sha1"), &ptr.NoSHA1)
		return ptr
	case "srp_bytes_m":
		ptr := &mt.ToSrvSRPBytesM{}
		val := l.CheckTable(3)
		readSliceByte(l, l.GetField(val, "m"), &ptr.M)
		return ptr
	}

	panic("invalid packet type: " + str)
}
