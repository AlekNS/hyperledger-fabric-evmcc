package ethvm

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/JincorTech/hyperledger-fabric-evmcc/burrow/word256"
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *MsgPackAccountPermissions) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "base":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "perms":
					z.Base.Perms, err = dc.ReadUint64()
					if err != nil {
						return
					}
				case "setbit":
					z.Base.SetBit, err = dc.ReadUint64()
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
		case "roles":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Roles) >= int(zb0003) {
				z.Roles = (z.Roles)[:zb0003]
			} else {
				z.Roles = make([]string, zb0003)
			}
			for za0001 := range z.Roles {
				z.Roles[za0001], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MsgPackAccountPermissions) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "base"
	// map header, size 2
	// write "perms"
	err = en.Append(0x82, 0xa4, 0x62, 0x61, 0x73, 0x65, 0x82, 0xa5, 0x70, 0x65, 0x72, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.Base.Perms)
	if err != nil {
		return
	}
	// write "setbit"
	err = en.Append(0xa6, 0x73, 0x65, 0x74, 0x62, 0x69, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.Base.SetBit)
	if err != nil {
		return
	}
	// write "roles"
	err = en.Append(0xa5, 0x72, 0x6f, 0x6c, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Roles)))
	if err != nil {
		return
	}
	for za0001 := range z.Roles {
		err = en.WriteString(z.Roles[za0001])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MsgPackAccountPermissions) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "base"
	// map header, size 2
	// string "perms"
	o = append(o, 0x82, 0xa4, 0x62, 0x61, 0x73, 0x65, 0x82, 0xa5, 0x70, 0x65, 0x72, 0x6d, 0x73)
	o = msgp.AppendUint64(o, z.Base.Perms)
	// string "setbit"
	o = append(o, 0xa6, 0x73, 0x65, 0x74, 0x62, 0x69, 0x74)
	o = msgp.AppendUint64(o, z.Base.SetBit)
	// string "roles"
	o = append(o, 0xa5, 0x72, 0x6f, 0x6c, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Roles)))
	for za0001 := range z.Roles {
		o = msgp.AppendString(o, z.Roles[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MsgPackAccountPermissions) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "base":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "perms":
					z.Base.Perms, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						return
					}
				case "setbit":
					z.Base.SetBit, bts, err = msgp.ReadUint64Bytes(bts)
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
		case "roles":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Roles) >= int(zb0003) {
				z.Roles = (z.Roles)[:zb0003]
			} else {
				z.Roles = make([]string, zb0003)
			}
			for za0001 := range z.Roles {
				z.Roles[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MsgPackAccountPermissions) Msgsize() (s int) {
	s = 1 + 5 + 1 + 6 + msgp.Uint64Size + 7 + msgp.Uint64Size + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Roles {
		s += msgp.StringPrefixSize + len(z.Roles[za0001])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MsgPackBasePermissions) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "perms":
			z.Perms, err = dc.ReadUint64()
			if err != nil {
				return
			}
		case "setbit":
			z.SetBit, err = dc.ReadUint64()
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
	return
}

// EncodeMsg implements msgp.Encodable
func (z MsgPackBasePermissions) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "perms"
	err = en.Append(0x82, 0xa5, 0x70, 0x65, 0x72, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.Perms)
	if err != nil {
		return
	}
	// write "setbit"
	err = en.Append(0xa6, 0x73, 0x65, 0x74, 0x62, 0x69, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.SetBit)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z MsgPackBasePermissions) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "perms"
	o = append(o, 0x82, 0xa5, 0x70, 0x65, 0x72, 0x6d, 0x73)
	o = msgp.AppendUint64(o, z.Perms)
	// string "setbit"
	o = append(o, 0xa6, 0x73, 0x65, 0x74, 0x62, 0x69, 0x74)
	o = msgp.AppendUint64(o, z.SetBit)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MsgPackBasePermissions) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "perms":
			z.Perms, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				return
			}
		case "setbit":
			z.SetBit, bts, err = msgp.ReadUint64Bytes(bts)
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
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MsgPackBasePermissions) Msgsize() (s int) {
	s = 1 + 6 + msgp.Uint64Size + 7 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *MsgPackEvmAccount) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "address":
			err = dc.ReadExactBytes((z.Address)[:])
			if err != nil {
				return
			}
		case "balance":
			z.Balance, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "code":
			z.Code, err = dc.ReadBytes(z.Code)
			if err != nil {
				return
			}
		case "nonce":
			z.Nonce, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "other":
			z.Other, err = dc.ReadIntf()
			if err != nil {
				return
			}
		case "account_perms":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "base":
					var zb0003 uint32
					zb0003, err = dc.ReadMapHeader()
					if err != nil {
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, err = dc.ReadMapKeyPtr()
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "perms":
							z.Permissions.Base.Perms, err = dc.ReadUint64()
							if err != nil {
								return
							}
						case "setbit":
							z.Permissions.Base.SetBit, err = dc.ReadUint64()
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
				case "roles":
					var zb0004 uint32
					zb0004, err = dc.ReadArrayHeader()
					if err != nil {
						return
					}
					if cap(z.Permissions.Roles) >= int(zb0004) {
						z.Permissions.Roles = (z.Permissions.Roles)[:zb0004]
					} else {
						z.Permissions.Roles = make([]string, zb0004)
					}
					for za0002 := range z.Permissions.Roles {
						z.Permissions.Roles[za0002], err = dc.ReadString()
						if err != nil {
							return
						}
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *MsgPackEvmAccount) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "address"
	err = en.Append(0x86, 0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteBytes((z.Address)[:])
	if err != nil {
		return
	}
	// write "balance"
	err = en.Append(0xa7, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Balance)
	if err != nil {
		return
	}
	// write "code"
	err = en.Append(0xa4, 0x63, 0x6f, 0x64, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Code)
	if err != nil {
		return
	}
	// write "nonce"
	err = en.Append(0xa5, 0x6e, 0x6f, 0x6e, 0x63, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteInt64(z.Nonce)
	if err != nil {
		return
	}
	// write "other"
	err = en.Append(0xa5, 0x6f, 0x74, 0x68, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteIntf(z.Other)
	if err != nil {
		return
	}
	// write "account_perms"
	// map header, size 2
	// write "base"
	// map header, size 2
	// write "perms"
	err = en.Append(0xad, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x82, 0xa4, 0x62, 0x61, 0x73, 0x65, 0x82, 0xa5, 0x70, 0x65, 0x72, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.Permissions.Base.Perms)
	if err != nil {
		return
	}
	// write "setbit"
	err = en.Append(0xa6, 0x73, 0x65, 0x74, 0x62, 0x69, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteUint64(z.Permissions.Base.SetBit)
	if err != nil {
		return
	}
	// write "roles"
	err = en.Append(0xa5, 0x72, 0x6f, 0x6c, 0x65, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Permissions.Roles)))
	if err != nil {
		return
	}
	for za0002 := range z.Permissions.Roles {
		err = en.WriteString(z.Permissions.Roles[za0002])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MsgPackEvmAccount) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "address"
	o = append(o, 0x86, 0xa7, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73)
	o = msgp.AppendBytes(o, (z.Address)[:])
	// string "balance"
	o = append(o, 0xa7, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65)
	o = msgp.AppendInt64(o, z.Balance)
	// string "code"
	o = append(o, 0xa4, 0x63, 0x6f, 0x64, 0x65)
	o = msgp.AppendBytes(o, z.Code)
	// string "nonce"
	o = append(o, 0xa5, 0x6e, 0x6f, 0x6e, 0x63, 0x65)
	o = msgp.AppendInt64(o, z.Nonce)
	// string "other"
	o = append(o, 0xa5, 0x6f, 0x74, 0x68, 0x65, 0x72)
	o, err = msgp.AppendIntf(o, z.Other)
	if err != nil {
		return
	}
	// string "account_perms"
	// map header, size 2
	// string "base"
	// map header, size 2
	// string "perms"
	o = append(o, 0xad, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x73, 0x82, 0xa4, 0x62, 0x61, 0x73, 0x65, 0x82, 0xa5, 0x70, 0x65, 0x72, 0x6d, 0x73)
	o = msgp.AppendUint64(o, z.Permissions.Base.Perms)
	// string "setbit"
	o = append(o, 0xa6, 0x73, 0x65, 0x74, 0x62, 0x69, 0x74)
	o = msgp.AppendUint64(o, z.Permissions.Base.SetBit)
	// string "roles"
	o = append(o, 0xa5, 0x72, 0x6f, 0x6c, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Permissions.Roles)))
	for za0002 := range z.Permissions.Roles {
		o = msgp.AppendString(o, z.Permissions.Roles[za0002])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MsgPackEvmAccount) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "address":
			bts, err = msgp.ReadExactBytes(bts, (z.Address)[:])
			if err != nil {
				return
			}
		case "balance":
			z.Balance, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "code":
			z.Code, bts, err = msgp.ReadBytesBytes(bts, z.Code)
			if err != nil {
				return
			}
		case "nonce":
			z.Nonce, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "other":
			z.Other, bts, err = msgp.ReadIntfBytes(bts)
			if err != nil {
				return
			}
		case "account_perms":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "base":
					var zb0003 uint32
					zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zb0003 > 0 {
						zb0003--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "perms":
							z.Permissions.Base.Perms, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								return
							}
						case "setbit":
							z.Permissions.Base.SetBit, bts, err = msgp.ReadUint64Bytes(bts)
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
				case "roles":
					var zb0004 uint32
					zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						return
					}
					if cap(z.Permissions.Roles) >= int(zb0004) {
						z.Permissions.Roles = (z.Permissions.Roles)[:zb0004]
					} else {
						z.Permissions.Roles = make([]string, zb0004)
					}
					for za0002 := range z.Permissions.Roles {
						z.Permissions.Roles[za0002], bts, err = msgp.ReadStringBytes(bts)
						if err != nil {
							return
						}
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MsgPackEvmAccount) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize + (word256.Word256Length * (msgp.ByteSize)) + 8 + msgp.Int64Size + 5 + msgp.BytesPrefixSize + len(z.Code) + 6 + msgp.Int64Size + 6 + msgp.GuessSize(z.Other) + 14 + 1 + 5 + 1 + 6 + msgp.Uint64Size + 7 + msgp.Uint64Size + 6 + msgp.ArrayHeaderSize
	for za0002 := range z.Permissions.Roles {
		s += msgp.StringPrefixSize + len(z.Permissions.Roles[za0002])
	}
	return
}
