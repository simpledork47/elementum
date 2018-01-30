package tmdb

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z AlternativeTitle) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Iso3166_1"
	o = append(o, 0x82, 0xa9, 0x49, 0x73, 0x6f, 0x33, 0x31, 0x36, 0x36, 0x5f, 0x31)
	o = msgp.AppendString(o, z.Iso3166_1)
	// string "Title"
	o = append(o, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.Title)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AlternativeTitle) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Iso3166_1":
			z.Iso3166_1, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Title":
			z.Title, bts, err = msgp.ReadStringBytes(bts)
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
func (z AlternativeTitle) Msgsize() (s int) {
	s = 1 + 10 + msgp.StringPrefixSize + len(z.Iso3166_1) + 6 + msgp.StringPrefixSize + len(z.Title)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Cast) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "IDName"
	// map header, size 2
	// string "ID"
	o = append(o, 0x86, 0xa6, 0x49, 0x44, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa2, 0x49, 0x44)
	o = msgp.AppendInt(o, z.IDName.ID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.IDName.Name)
	// string "CastID"
	o = append(o, 0xa6, 0x43, 0x61, 0x73, 0x74, 0x49, 0x44)
	o = msgp.AppendInt(o, z.CastID)
	// string "Character"
	o = append(o, 0xa9, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72)
	o = msgp.AppendString(o, z.Character)
	// string "CreditID"
	o = append(o, 0xa8, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x49, 0x44)
	o = msgp.AppendString(o, z.CreditID)
	// string "Order"
	o = append(o, 0xa5, 0x4f, 0x72, 0x64, 0x65, 0x72)
	o = msgp.AppendInt(o, z.Order)
	// string "ProfilePath"
	o = append(o, 0xab, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.ProfilePath)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Cast) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "IDName":
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
				case "ID":
					z.IDName.ID, bts, err = msgp.ReadIntBytes(bts)
					if err != nil {
						return
					}
				case "Name":
					z.IDName.Name, bts, err = msgp.ReadStringBytes(bts)
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
		case "CastID":
			z.CastID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Character":
			z.Character, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "CreditID":
			z.CreditID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Order":
			z.Order, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "ProfilePath":
			z.ProfilePath, bts, err = msgp.ReadStringBytes(bts)
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
func (z *Cast) Msgsize() (s int) {
	s = 1 + 7 + 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.IDName.Name) + 7 + msgp.IntSize + 10 + msgp.StringPrefixSize + len(z.Character) + 9 + msgp.StringPrefixSize + len(z.CreditID) + 6 + msgp.IntSize + 12 + msgp.StringPrefixSize + len(z.ProfilePath)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Credits) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Cast"
	o = append(o, 0x82, 0xa4, 0x43, 0x61, 0x73, 0x74)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Cast)))
	for za0001 := range z.Cast {
		if z.Cast[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Cast[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Crew"
	o = append(o, 0xa4, 0x43, 0x72, 0x65, 0x77)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Crew)))
	for za0002 := range z.Crew {
		if z.Crew[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Crew[za0002].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Credits) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Cast":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Cast) >= int(zb0002) {
				z.Cast = (z.Cast)[:zb0002]
			} else {
				z.Cast = make([]*Cast, zb0002)
			}
			for za0001 := range z.Cast {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Cast[za0001] = nil
				} else {
					if z.Cast[za0001] == nil {
						z.Cast[za0001] = new(Cast)
					}
					bts, err = z.Cast[za0001].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Crew":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Crew) >= int(zb0003) {
				z.Crew = (z.Crew)[:zb0003]
			} else {
				z.Crew = make([]*Crew, zb0003)
			}
			for za0002 := range z.Crew {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Crew[za0002] = nil
				} else {
					if z.Crew[za0002] == nil {
						z.Crew[za0002] = new(Crew)
					}
					bts, err = z.Crew[za0002].UnmarshalMsg(bts)
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
func (z *Credits) Msgsize() (s int) {
	s = 1 + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Cast {
		if z.Cast[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Cast[za0001].Msgsize()
		}
	}
	s += 5 + msgp.ArrayHeaderSize
	for za0002 := range z.Crew {
		if z.Crew[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += z.Crew[za0002].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Crew) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "IDName"
	// map header, size 2
	// string "ID"
	o = append(o, 0x85, 0xa6, 0x49, 0x44, 0x4e, 0x61, 0x6d, 0x65, 0x82, 0xa2, 0x49, 0x44)
	o = msgp.AppendInt(o, z.IDName.ID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.IDName.Name)
	// string "CreditID"
	o = append(o, 0xa8, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x49, 0x44)
	o = msgp.AppendString(o, z.CreditID)
	// string "Department"
	o = append(o, 0xaa, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Department)
	// string "Job"
	o = append(o, 0xa3, 0x4a, 0x6f, 0x62)
	o = msgp.AppendString(o, z.Job)
	// string "ProfilePath"
	o = append(o, 0xab, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.ProfilePath)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Crew) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "IDName":
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
				case "ID":
					z.IDName.ID, bts, err = msgp.ReadIntBytes(bts)
					if err != nil {
						return
					}
				case "Name":
					z.IDName.Name, bts, err = msgp.ReadStringBytes(bts)
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
		case "CreditID":
			z.CreditID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Department":
			z.Department, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Job":
			z.Job, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ProfilePath":
			z.ProfilePath, bts, err = msgp.ReadStringBytes(bts)
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
func (z *Crew) Msgsize() (s int) {
	s = 1 + 7 + 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.IDName.Name) + 9 + msgp.StringPrefixSize + len(z.CreditID) + 11 + msgp.StringPrefixSize + len(z.Department) + 4 + msgp.StringPrefixSize + len(z.Job) + 12 + msgp.StringPrefixSize + len(z.ProfilePath)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Entity) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 14
	// string "IsAdult"
	o = append(o, 0x8e, 0xa7, 0x49, 0x73, 0x41, 0x64, 0x75, 0x6c, 0x74)
	o = msgp.AppendBool(o, z.IsAdult)
	// string "BackdropPath"
	o = append(o, 0xac, 0x42, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f, 0x70, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.BackdropPath)
	// string "ID"
	o = append(o, 0xa2, 0x49, 0x44)
	o = msgp.AppendInt(o, z.ID)
	// string "Genres"
	o = append(o, 0xa6, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Genres)))
	for za0001 := range z.Genres {
		if z.Genres[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "ID"
			o = append(o, 0x82, 0xa2, 0x49, 0x44)
			o = msgp.AppendInt(o, z.Genres[za0001].ID)
			// string "Name"
			o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.Genres[za0001].Name)
		}
	}
	// string "OriginalTitle"
	o = append(o, 0xad, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.OriginalTitle)
	// string "OriginalLanguage"
	o = append(o, 0xb0, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.OriginalLanguage)
	// string "ReleaseDate"
	o = append(o, 0xab, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.ReleaseDate)
	// string "FirstAirDate"
	o = append(o, 0xac, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x44, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.FirstAirDate)
	// string "PosterPath"
	o = append(o, 0xaa, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.PosterPath)
	// string "Title"
	o = append(o, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
	o = msgp.AppendString(o, z.Title)
	// string "VoteAverage"
	o = append(o, 0xab, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65)
	o = msgp.AppendFloat32(o, z.VoteAverage)
	// string "VoteCount"
	o = append(o, 0xa9, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.VoteCount)
	// string "OriginalName"
	o = append(o, 0xac, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.OriginalName)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Entity) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "IsAdult":
			z.IsAdult, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "BackdropPath":
			z.BackdropPath, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ID":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Genres":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Genres) >= int(zb0002) {
				z.Genres = (z.Genres)[:zb0002]
			} else {
				z.Genres = make([]*IDName, zb0002)
			}
			for za0001 := range z.Genres {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Genres[za0001] = nil
				} else {
					if z.Genres[za0001] == nil {
						z.Genres[za0001] = new(IDName)
					}
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
						case "ID":
							z.Genres[za0001].ID, bts, err = msgp.ReadIntBytes(bts)
							if err != nil {
								return
							}
						case "Name":
							z.Genres[za0001].Name, bts, err = msgp.ReadStringBytes(bts)
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
				}
			}
		case "OriginalTitle":
			z.OriginalTitle, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "OriginalLanguage":
			z.OriginalLanguage, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ReleaseDate":
			z.ReleaseDate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FirstAirDate":
			z.FirstAirDate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "PosterPath":
			z.PosterPath, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Title":
			z.Title, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "VoteAverage":
			z.VoteAverage, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				return
			}
		case "VoteCount":
			z.VoteCount, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "OriginalName":
			z.OriginalName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
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
func (z *Entity) Msgsize() (s int) {
	s = 1 + 8 + msgp.BoolSize + 13 + msgp.StringPrefixSize + len(z.BackdropPath) + 3 + msgp.IntSize + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Genres {
		if z.Genres[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Genres[za0001].Name)
		}
	}
	s += 14 + msgp.StringPrefixSize + len(z.OriginalTitle) + 17 + msgp.StringPrefixSize + len(z.OriginalLanguage) + 12 + msgp.StringPrefixSize + len(z.ReleaseDate) + 13 + msgp.StringPrefixSize + len(z.FirstAirDate) + 11 + msgp.StringPrefixSize + len(z.PosterPath) + 6 + msgp.StringPrefixSize + len(z.Title) + 12 + msgp.Float32Size + 10 + msgp.IntSize + 13 + msgp.StringPrefixSize + len(z.OriginalName) + 5 + msgp.StringPrefixSize + len(z.Name)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *EntityList) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Page"
	o = append(o, 0x84, 0xa4, 0x50, 0x61, 0x67, 0x65)
	o = msgp.AppendInt(o, z.Page)
	// string "Results"
	o = append(o, 0xa7, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Results)))
	for za0001 := range z.Results {
		if z.Results[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Results[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "TotalPages"
	o = append(o, 0xaa, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73)
	o = msgp.AppendInt(o, z.TotalPages)
	// string "TotalResults"
	o = append(o, 0xac, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	o = msgp.AppendInt(o, z.TotalResults)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EntityList) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Page":
			z.Page, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Results":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Results) >= int(zb0002) {
				z.Results = (z.Results)[:zb0002]
			} else {
				z.Results = make([]*Entity, zb0002)
			}
			for za0001 := range z.Results {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Results[za0001] = nil
				} else {
					if z.Results[za0001] == nil {
						z.Results[za0001] = new(Entity)
					}
					bts, err = z.Results[za0001].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "TotalPages":
			z.TotalPages, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "TotalResults":
			z.TotalResults, bts, err = msgp.ReadIntBytes(bts)
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
func (z *EntityList) Msgsize() (s int) {
	s = 1 + 5 + msgp.IntSize + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Results {
		if z.Results[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Results[za0001].Msgsize()
		}
	}
	s += 11 + msgp.IntSize + 13 + msgp.IntSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Episode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "ID"
	o = append(o, 0x89, 0xa2, 0x49, 0x44)
	o = msgp.AppendInt(o, z.ID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Overview"
	o = append(o, 0xa8, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77)
	o = msgp.AppendString(o, z.Overview)
	// string "AirDate"
	o = append(o, 0xa7, 0x41, 0x69, 0x72, 0x44, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.AirDate)
	// string "SeasonNumber"
	o = append(o, 0xac, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.SeasonNumber)
	// string "EpisodeNumber"
	o = append(o, 0xad, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.EpisodeNumber)
	// string "VoteAverage"
	o = append(o, 0xab, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65)
	o = msgp.AppendFloat32(o, z.VoteAverage)
	// string "StillPath"
	o = append(o, 0xa9, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.StillPath)
	// string "ExternalIDs"
	o = append(o, 0xab, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x73)
	if z.ExternalIDs == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.ExternalIDs.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Episode) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "ID":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Overview":
			z.Overview, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AirDate":
			z.AirDate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SeasonNumber":
			z.SeasonNumber, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "EpisodeNumber":
			z.EpisodeNumber, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "VoteAverage":
			z.VoteAverage, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				return
			}
		case "StillPath":
			z.StillPath, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ExternalIDs":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ExternalIDs = nil
			} else {
				if z.ExternalIDs == nil {
					z.ExternalIDs = new(ExternalIDs)
				}
				bts, err = z.ExternalIDs.UnmarshalMsg(bts)
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
func (z *Episode) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.StringPrefixSize + len(z.Overview) + 8 + msgp.StringPrefixSize + len(z.AirDate) + 13 + msgp.IntSize + 14 + msgp.IntSize + 12 + msgp.Float32Size + 10 + msgp.StringPrefixSize + len(z.StillPath) + 12
	if z.ExternalIDs == nil {
		s += msgp.NilSize
	} else {
		s += z.ExternalIDs.Msgsize()
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z EpisodeList) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		if z[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *EpisodeList) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(EpisodeList, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Episode)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z EpisodeList) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ExternalIDs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "IMDBId"
	o = append(o, 0x84, 0xa6, 0x49, 0x4d, 0x44, 0x42, 0x49, 0x64)
	o = msgp.AppendString(o, z.IMDBId)
	// string "FreeBaseID"
	o = append(o, 0xaa, 0x46, 0x72, 0x65, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x44)
	o = msgp.AppendString(o, z.FreeBaseID)
	// string "FreeBaseMID"
	o = append(o, 0xab, 0x46, 0x72, 0x65, 0x65, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x49, 0x44)
	o = msgp.AppendString(o, z.FreeBaseMID)
	// string "TVDBID"
	o = append(o, 0xa6, 0x54, 0x56, 0x44, 0x42, 0x49, 0x44)
	o, err = msgp.AppendIntf(o, z.TVDBID)
	if err != nil {
		return
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ExternalIDs) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "IMDBId":
			z.IMDBId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FreeBaseID":
			z.FreeBaseID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FreeBaseMID":
			z.FreeBaseMID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TVDBID":
			z.TVDBID, bts, err = msgp.ReadIntfBytes(bts)
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
func (z *ExternalIDs) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.IMDBId) + 11 + msgp.StringPrefixSize + len(z.FreeBaseID) + 12 + msgp.StringPrefixSize + len(z.FreeBaseMID) + 7 + msgp.GuessSize(z.TVDBID)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *FindResult) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "MovieResults"
	o = append(o, 0x85, 0xac, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MovieResults)))
	for za0001 := range z.MovieResults {
		if z.MovieResults[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.MovieResults[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "PersonResults"
	o = append(o, 0xad, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.PersonResults)))
	for za0002 := range z.PersonResults {
		if z.PersonResults[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.PersonResults[za0002].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "TVResults"
	o = append(o, 0xa9, 0x54, 0x56, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TVResults)))
	for za0003 := range z.TVResults {
		if z.TVResults[za0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.TVResults[za0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "TVEpisodeResults"
	o = append(o, 0xb0, 0x54, 0x56, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TVEpisodeResults)))
	for za0004 := range z.TVEpisodeResults {
		if z.TVEpisodeResults[za0004] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.TVEpisodeResults[za0004].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "TVSeasonResults"
	o = append(o, 0xaf, 0x54, 0x56, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.TVSeasonResults)))
	for za0005 := range z.TVSeasonResults {
		if z.TVSeasonResults[za0005] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.TVSeasonResults[za0005].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *FindResult) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "MovieResults":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.MovieResults) >= int(zb0002) {
				z.MovieResults = (z.MovieResults)[:zb0002]
			} else {
				z.MovieResults = make([]*Entity, zb0002)
			}
			for za0001 := range z.MovieResults {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.MovieResults[za0001] = nil
				} else {
					if z.MovieResults[za0001] == nil {
						z.MovieResults[za0001] = new(Entity)
					}
					bts, err = z.MovieResults[za0001].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "PersonResults":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.PersonResults) >= int(zb0003) {
				z.PersonResults = (z.PersonResults)[:zb0003]
			} else {
				z.PersonResults = make([]*Entity, zb0003)
			}
			for za0002 := range z.PersonResults {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.PersonResults[za0002] = nil
				} else {
					if z.PersonResults[za0002] == nil {
						z.PersonResults[za0002] = new(Entity)
					}
					bts, err = z.PersonResults[za0002].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "TVResults":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.TVResults) >= int(zb0004) {
				z.TVResults = (z.TVResults)[:zb0004]
			} else {
				z.TVResults = make([]*Entity, zb0004)
			}
			for za0003 := range z.TVResults {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.TVResults[za0003] = nil
				} else {
					if z.TVResults[za0003] == nil {
						z.TVResults[za0003] = new(Entity)
					}
					bts, err = z.TVResults[za0003].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "TVEpisodeResults":
			var zb0005 uint32
			zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.TVEpisodeResults) >= int(zb0005) {
				z.TVEpisodeResults = (z.TVEpisodeResults)[:zb0005]
			} else {
				z.TVEpisodeResults = make([]*Entity, zb0005)
			}
			for za0004 := range z.TVEpisodeResults {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.TVEpisodeResults[za0004] = nil
				} else {
					if z.TVEpisodeResults[za0004] == nil {
						z.TVEpisodeResults[za0004] = new(Entity)
					}
					bts, err = z.TVEpisodeResults[za0004].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "TVSeasonResults":
			var zb0006 uint32
			zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.TVSeasonResults) >= int(zb0006) {
				z.TVSeasonResults = (z.TVSeasonResults)[:zb0006]
			} else {
				z.TVSeasonResults = make([]*Entity, zb0006)
			}
			for za0005 := range z.TVSeasonResults {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.TVSeasonResults[za0005] = nil
				} else {
					if z.TVSeasonResults[za0005] == nil {
						z.TVSeasonResults[za0005] = new(Entity)
					}
					bts, err = z.TVSeasonResults[za0005].UnmarshalMsg(bts)
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
func (z *FindResult) Msgsize() (s int) {
	s = 1 + 13 + msgp.ArrayHeaderSize
	for za0001 := range z.MovieResults {
		if z.MovieResults[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.MovieResults[za0001].Msgsize()
		}
	}
	s += 14 + msgp.ArrayHeaderSize
	for za0002 := range z.PersonResults {
		if z.PersonResults[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += z.PersonResults[za0002].Msgsize()
		}
	}
	s += 10 + msgp.ArrayHeaderSize
	for za0003 := range z.TVResults {
		if z.TVResults[za0003] == nil {
			s += msgp.NilSize
		} else {
			s += z.TVResults[za0003].Msgsize()
		}
	}
	s += 17 + msgp.ArrayHeaderSize
	for za0004 := range z.TVEpisodeResults {
		if z.TVEpisodeResults[za0004] == nil {
			s += msgp.NilSize
		} else {
			s += z.TVEpisodeResults[za0004].Msgsize()
		}
	}
	s += 16 + msgp.ArrayHeaderSize
	for za0005 := range z.TVSeasonResults {
		if z.TVSeasonResults[za0005] == nil {
			s += msgp.NilSize
		} else {
			s += z.TVSeasonResults[za0005].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Genre) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "ID"
	o = append(o, 0x82, 0xa2, 0x49, 0x44)
	o = msgp.AppendInt(o, z.ID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Genre) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "ID":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
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
func (z Genre) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *GenreList) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Genres"
	o = append(o, 0x81, 0xa6, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Genres)))
	for za0001 := range z.Genres {
		if z.Genres[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "ID"
			o = append(o, 0x82, 0xa2, 0x49, 0x44)
			o = msgp.AppendInt(o, z.Genres[za0001].ID)
			// string "Name"
			o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.Genres[za0001].Name)
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *GenreList) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Genres":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Genres) >= int(zb0002) {
				z.Genres = (z.Genres)[:zb0002]
			} else {
				z.Genres = make([]*Genre, zb0002)
			}
			for za0001 := range z.Genres {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Genres[za0001] = nil
				} else {
					if z.Genres[za0001] == nil {
						z.Genres[za0001] = new(Genre)
					}
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
						case "ID":
							z.Genres[za0001].ID, bts, err = msgp.ReadIntBytes(bts)
							if err != nil {
								return
							}
						case "Name":
							z.Genres[za0001].Name, bts, err = msgp.ReadStringBytes(bts)
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
func (z *GenreList) Msgsize() (s int) {
	s = 1 + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Genres {
		if z.Genres[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Genres[za0001].Name)
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z IDName) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "ID"
	o = append(o, 0x82, 0xa2, 0x49, 0x44)
	o = msgp.AppendInt(o, z.ID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *IDName) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "ID":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
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
func (z IDName) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Image) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "FilePath"
	o = append(o, 0x84, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.FilePath)
	// string "Height"
	o = append(o, 0xa6, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendInt(o, z.Height)
	// string "Iso639_1"
	o = append(o, 0xa8, 0x49, 0x73, 0x6f, 0x36, 0x33, 0x39, 0x5f, 0x31)
	o = msgp.AppendString(o, z.Iso639_1)
	// string "Width"
	o = append(o, 0xa5, 0x57, 0x69, 0x64, 0x74, 0x68)
	o = msgp.AppendInt(o, z.Width)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Image) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "FilePath":
			z.FilePath, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Height":
			z.Height, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Iso639_1":
			z.Iso639_1, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Width":
			z.Width, bts, err = msgp.ReadIntBytes(bts)
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
func (z *Image) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.FilePath) + 7 + msgp.IntSize + 9 + msgp.StringPrefixSize + len(z.Iso639_1) + 6 + msgp.IntSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Images) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Backdrops"
	o = append(o, 0x83, 0xa9, 0x42, 0x61, 0x63, 0x6b, 0x64, 0x72, 0x6f, 0x70, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Backdrops)))
	for za0001 := range z.Backdrops {
		if z.Backdrops[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Backdrops[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Posters"
	o = append(o, 0xa7, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Posters)))
	for za0002 := range z.Posters {
		if z.Posters[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Posters[za0002].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Stills"
	o = append(o, 0xa6, 0x53, 0x74, 0x69, 0x6c, 0x6c, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Stills)))
	for za0003 := range z.Stills {
		if z.Stills[za0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Stills[za0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Images) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Backdrops":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Backdrops) >= int(zb0002) {
				z.Backdrops = (z.Backdrops)[:zb0002]
			} else {
				z.Backdrops = make([]*Image, zb0002)
			}
			for za0001 := range z.Backdrops {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Backdrops[za0001] = nil
				} else {
					if z.Backdrops[za0001] == nil {
						z.Backdrops[za0001] = new(Image)
					}
					bts, err = z.Backdrops[za0001].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Posters":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Posters) >= int(zb0003) {
				z.Posters = (z.Posters)[:zb0003]
			} else {
				z.Posters = make([]*Image, zb0003)
			}
			for za0002 := range z.Posters {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Posters[za0002] = nil
				} else {
					if z.Posters[za0002] == nil {
						z.Posters[za0002] = new(Image)
					}
					bts, err = z.Posters[za0002].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Stills":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Stills) >= int(zb0004) {
				z.Stills = (z.Stills)[:zb0004]
			} else {
				z.Stills = make([]*Image, zb0004)
			}
			for za0003 := range z.Stills {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Stills[za0003] = nil
				} else {
					if z.Stills[za0003] == nil {
						z.Stills[za0003] = new(Image)
					}
					bts, err = z.Stills[za0003].UnmarshalMsg(bts)
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
func (z *Images) Msgsize() (s int) {
	s = 1 + 10 + msgp.ArrayHeaderSize
	for za0001 := range z.Backdrops {
		if z.Backdrops[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Backdrops[za0001].Msgsize()
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for za0002 := range z.Posters {
		if z.Posters[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += z.Posters[za0002].Msgsize()
		}
	}
	s += 7 + msgp.ArrayHeaderSize
	for za0003 := range z.Stills {
		if z.Stills[za0003] == nil {
			s += msgp.NilSize
		} else {
			s += z.Stills[za0003].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Language) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "Iso639_1"
	o = append(o, 0x83, 0xa8, 0x49, 0x73, 0x6f, 0x36, 0x33, 0x39, 0x5f, 0x31)
	o = msgp.AppendString(o, z.Iso639_1)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "EnglishName"
	o = append(o, 0xab, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.EnglishName)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Language) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Iso639_1":
			z.Iso639_1, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "EnglishName":
			z.EnglishName, bts, err = msgp.ReadStringBytes(bts)
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
func (z Language) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Iso639_1) + 5 + msgp.StringPrefixSize + len(z.Name) + 12 + msgp.StringPrefixSize + len(z.EnglishName)
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *List) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "CreatedBy"
	o = append(o, 0x89, 0xa9, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79)
	o = msgp.AppendString(o, z.CreatedBy)
	// string "Description"
	o = append(o, 0xab, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Description)
	// string "FavoriteCount"
	o = append(o, 0xad, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.FavoriteCount)
	// string "ID"
	o = append(o, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, z.ID)
	// string "ItemCount"
	o = append(o, 0xa9, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.ItemCount)
	// string "Iso639_1"
	o = append(o, 0xa8, 0x49, 0x73, 0x6f, 0x36, 0x33, 0x39, 0x5f, 0x31)
	o = msgp.AppendString(o, z.Iso639_1)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "PosterPath"
	o = append(o, 0xaa, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.PosterPath)
	// string "Items"
	o = append(o, 0xa5, 0x49, 0x74, 0x65, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Items)))
	for za0001 := range z.Items {
		if z.Items[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Items[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *List) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "CreatedBy":
			z.CreatedBy, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Description":
			z.Description, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FavoriteCount":
			z.FavoriteCount, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "ID":
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ItemCount":
			z.ItemCount, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Iso639_1":
			z.Iso639_1, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "PosterPath":
			z.PosterPath, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Items":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Items) >= int(zb0002) {
				z.Items = (z.Items)[:zb0002]
			} else {
				z.Items = make([]*Entity, zb0002)
			}
			for za0001 := range z.Items {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Items[za0001] = nil
				} else {
					if z.Items[za0001] == nil {
						z.Items[za0001] = new(Entity)
					}
					bts, err = z.Items[za0001].UnmarshalMsg(bts)
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
func (z *List) Msgsize() (s int) {
	s = 1 + 10 + msgp.StringPrefixSize + len(z.CreatedBy) + 12 + msgp.StringPrefixSize + len(z.Description) + 14 + msgp.IntSize + 3 + msgp.StringPrefixSize + len(z.ID) + 10 + msgp.IntSize + 9 + msgp.StringPrefixSize + len(z.Iso639_1) + 5 + msgp.StringPrefixSize + len(z.Name) + 11 + msgp.StringPrefixSize + len(z.PosterPath) + 6 + msgp.ArrayHeaderSize
	for za0001 := range z.Items {
		if z.Items[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Items[za0001].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Movie) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 16
	// string "Entity"
	o = append(o, 0xde, 0x0, 0x10, 0xa6, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79)
	o, err = z.Entity.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "IMDBId"
	o = append(o, 0xa6, 0x49, 0x4d, 0x44, 0x42, 0x49, 0x64)
	o = msgp.AppendString(o, z.IMDBId)
	// string "Overview"
	o = append(o, 0xa8, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77)
	o = msgp.AppendString(o, z.Overview)
	// string "ProductionCompanies"
	o = append(o, 0xb3, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.ProductionCompanies)))
	for za0001 := range z.ProductionCompanies {
		if z.ProductionCompanies[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "ID"
			o = append(o, 0x82, 0xa2, 0x49, 0x44)
			o = msgp.AppendInt(o, z.ProductionCompanies[za0001].ID)
			// string "Name"
			o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.ProductionCompanies[za0001].Name)
		}
	}
	// string "Runtime"
	o = append(o, 0xa7, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt(o, z.Runtime)
	// string "TagLine"
	o = append(o, 0xa7, 0x54, 0x61, 0x67, 0x4c, 0x69, 0x6e, 0x65)
	o = msgp.AppendString(o, z.TagLine)
	// string "RawPopularity"
	o = append(o, 0xad, 0x52, 0x61, 0x77, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79)
	o, err = msgp.AppendIntf(o, z.RawPopularity)
	if err != nil {
		return
	}
	// string "Popularity"
	o = append(o, 0xaa, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79)
	o = msgp.AppendFloat64(o, z.Popularity)
	// string "SpokenLanguages"
	o = append(o, 0xaf, 0x53, 0x70, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.SpokenLanguages)))
	for za0002 := range z.SpokenLanguages {
		if z.SpokenLanguages[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 3
			// string "Iso639_1"
			o = append(o, 0x83, 0xa8, 0x49, 0x73, 0x6f, 0x36, 0x33, 0x39, 0x5f, 0x31)
			o = msgp.AppendString(o, z.SpokenLanguages[za0002].Iso639_1)
			// string "Name"
			o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.SpokenLanguages[za0002].Name)
			// string "EnglishName"
			o = append(o, 0xab, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.SpokenLanguages[za0002].EnglishName)
		}
	}
	// string "ExternalIDs"
	o = append(o, 0xab, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x73)
	if z.ExternalIDs == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.ExternalIDs.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "AlternativeTitles"
	o = append(o, 0xb1, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73)
	if z.AlternativeTitles == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 1
		// string "Titles"
		o = append(o, 0x81, 0xa6, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.AlternativeTitles.Titles)))
		for za0003 := range z.AlternativeTitles.Titles {
			if z.AlternativeTitles.Titles[za0003] == nil {
				o = msgp.AppendNil(o)
			} else {
				// map header, size 2
				// string "Iso3166_1"
				o = append(o, 0x82, 0xa9, 0x49, 0x73, 0x6f, 0x33, 0x31, 0x36, 0x36, 0x5f, 0x31)
				o = msgp.AppendString(o, z.AlternativeTitles.Titles[za0003].Iso3166_1)
				// string "Title"
				o = append(o, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
				o = msgp.AppendString(o, z.AlternativeTitles.Titles[za0003].Title)
			}
		}
	}
	// string "Translations"
	o = append(o, 0xac, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	if z.Translations == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 1
		// string "Translations"
		o = append(o, 0x81, 0xac, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Translations.Translations)))
		for za0004 := range z.Translations.Translations {
			if z.Translations.Translations[za0004] == nil {
				o = msgp.AppendNil(o)
			} else {
				// map header, size 3
				// string "Iso639_1"
				o = append(o, 0x83, 0xa8, 0x49, 0x73, 0x6f, 0x36, 0x33, 0x39, 0x5f, 0x31)
				o = msgp.AppendString(o, z.Translations.Translations[za0004].Iso639_1)
				// string "Name"
				o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
				o = msgp.AppendString(o, z.Translations.Translations[za0004].Name)
				// string "EnglishName"
				o = append(o, 0xab, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65)
				o = msgp.AppendString(o, z.Translations.Translations[za0004].EnglishName)
			}
		}
	}
	// string "Trailers"
	o = append(o, 0xa8, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73)
	if z.Trailers == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 1
		// string "Youtube"
		o = append(o, 0x81, 0xa7, 0x59, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Trailers.Youtube)))
		for za0005 := range z.Trailers.Youtube {
			if z.Trailers.Youtube[za0005] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.Trailers.Youtube[za0005].MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}
	// string "Credits"
	o = append(o, 0xa7, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73)
	if z.Credits == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Credits.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Images"
	o = append(o, 0xa6, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73)
	if z.Images == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Images.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "ReleaseDates"
	o = append(o, 0xac, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x73)
	if z.ReleaseDates == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 1
		// string "Results"
		o = append(o, 0x81, 0xa7, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.ReleaseDates.Results)))
		for za0006 := range z.ReleaseDates.Results {
			if z.ReleaseDates.Results[za0006] == nil {
				o = msgp.AppendNil(o)
			} else {
				o, err = z.ReleaseDates.Results[za0006].MarshalMsg(o)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Movie) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Entity":
			bts, err = z.Entity.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "IMDBId":
			z.IMDBId, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Overview":
			z.Overview, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ProductionCompanies":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.ProductionCompanies) >= int(zb0002) {
				z.ProductionCompanies = (z.ProductionCompanies)[:zb0002]
			} else {
				z.ProductionCompanies = make([]*IDName, zb0002)
			}
			for za0001 := range z.ProductionCompanies {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.ProductionCompanies[za0001] = nil
				} else {
					if z.ProductionCompanies[za0001] == nil {
						z.ProductionCompanies[za0001] = new(IDName)
					}
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
						case "ID":
							z.ProductionCompanies[za0001].ID, bts, err = msgp.ReadIntBytes(bts)
							if err != nil {
								return
							}
						case "Name":
							z.ProductionCompanies[za0001].Name, bts, err = msgp.ReadStringBytes(bts)
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
				}
			}
		case "Runtime":
			z.Runtime, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "TagLine":
			z.TagLine, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RawPopularity":
			z.RawPopularity, bts, err = msgp.ReadIntfBytes(bts)
			if err != nil {
				return
			}
		case "Popularity":
			z.Popularity, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "SpokenLanguages":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.SpokenLanguages) >= int(zb0004) {
				z.SpokenLanguages = (z.SpokenLanguages)[:zb0004]
			} else {
				z.SpokenLanguages = make([]*Language, zb0004)
			}
			for za0002 := range z.SpokenLanguages {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.SpokenLanguages[za0002] = nil
				} else {
					if z.SpokenLanguages[za0002] == nil {
						z.SpokenLanguages[za0002] = new(Language)
					}
					var zb0005 uint32
					zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zb0005 > 0 {
						zb0005--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "Iso639_1":
							z.SpokenLanguages[za0002].Iso639_1, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "Name":
							z.SpokenLanguages[za0002].Name, bts, err = msgp.ReadStringBytes(bts)
							if err != nil {
								return
							}
						case "EnglishName":
							z.SpokenLanguages[za0002].EnglishName, bts, err = msgp.ReadStringBytes(bts)
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
				}
			}
		case "ExternalIDs":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ExternalIDs = nil
			} else {
				if z.ExternalIDs == nil {
					z.ExternalIDs = new(ExternalIDs)
				}
				bts, err = z.ExternalIDs.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "AlternativeTitles":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.AlternativeTitles = nil
			} else {
				if z.AlternativeTitles == nil {
					z.AlternativeTitles = new(struct {
						Titles []*AlternativeTitle `json:"titles"`
					})
				}
				var zb0006 uint32
				zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0006 > 0 {
					zb0006--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Titles":
						var zb0007 uint32
						zb0007, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.AlternativeTitles.Titles) >= int(zb0007) {
							z.AlternativeTitles.Titles = (z.AlternativeTitles.Titles)[:zb0007]
						} else {
							z.AlternativeTitles.Titles = make([]*AlternativeTitle, zb0007)
						}
						for za0003 := range z.AlternativeTitles.Titles {
							if msgp.IsNil(bts) {
								bts, err = msgp.ReadNilBytes(bts)
								if err != nil {
									return
								}
								z.AlternativeTitles.Titles[za0003] = nil
							} else {
								if z.AlternativeTitles.Titles[za0003] == nil {
									z.AlternativeTitles.Titles[za0003] = new(AlternativeTitle)
								}
								var zb0008 uint32
								zb0008, bts, err = msgp.ReadMapHeaderBytes(bts)
								if err != nil {
									return
								}
								for zb0008 > 0 {
									zb0008--
									field, bts, err = msgp.ReadMapKeyZC(bts)
									if err != nil {
										return
									}
									switch msgp.UnsafeString(field) {
									case "Iso3166_1":
										z.AlternativeTitles.Titles[za0003].Iso3166_1, bts, err = msgp.ReadStringBytes(bts)
										if err != nil {
											return
										}
									case "Title":
										z.AlternativeTitles.Titles[za0003].Title, bts, err = msgp.ReadStringBytes(bts)
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
							}
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Translations":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Translations = nil
			} else {
				if z.Translations == nil {
					z.Translations = new(struct {
						Translations []*Language `json:"translations"`
					})
				}
				var zb0009 uint32
				zb0009, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0009 > 0 {
					zb0009--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Translations":
						var zb0010 uint32
						zb0010, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Translations.Translations) >= int(zb0010) {
							z.Translations.Translations = (z.Translations.Translations)[:zb0010]
						} else {
							z.Translations.Translations = make([]*Language, zb0010)
						}
						for za0004 := range z.Translations.Translations {
							if msgp.IsNil(bts) {
								bts, err = msgp.ReadNilBytes(bts)
								if err != nil {
									return
								}
								z.Translations.Translations[za0004] = nil
							} else {
								if z.Translations.Translations[za0004] == nil {
									z.Translations.Translations[za0004] = new(Language)
								}
								var zb0011 uint32
								zb0011, bts, err = msgp.ReadMapHeaderBytes(bts)
								if err != nil {
									return
								}
								for zb0011 > 0 {
									zb0011--
									field, bts, err = msgp.ReadMapKeyZC(bts)
									if err != nil {
										return
									}
									switch msgp.UnsafeString(field) {
									case "Iso639_1":
										z.Translations.Translations[za0004].Iso639_1, bts, err = msgp.ReadStringBytes(bts)
										if err != nil {
											return
										}
									case "Name":
										z.Translations.Translations[za0004].Name, bts, err = msgp.ReadStringBytes(bts)
										if err != nil {
											return
										}
									case "EnglishName":
										z.Translations.Translations[za0004].EnglishName, bts, err = msgp.ReadStringBytes(bts)
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
							}
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Trailers":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Trailers = nil
			} else {
				if z.Trailers == nil {
					z.Trailers = new(struct {
						Youtube []*Trailer `json:"youtube"`
					})
				}
				var zb0012 uint32
				zb0012, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0012 > 0 {
					zb0012--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Youtube":
						var zb0013 uint32
						zb0013, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Trailers.Youtube) >= int(zb0013) {
							z.Trailers.Youtube = (z.Trailers.Youtube)[:zb0013]
						} else {
							z.Trailers.Youtube = make([]*Trailer, zb0013)
						}
						for za0005 := range z.Trailers.Youtube {
							if msgp.IsNil(bts) {
								bts, err = msgp.ReadNilBytes(bts)
								if err != nil {
									return
								}
								z.Trailers.Youtube[za0005] = nil
							} else {
								if z.Trailers.Youtube[za0005] == nil {
									z.Trailers.Youtube[za0005] = new(Trailer)
								}
								bts, err = z.Trailers.Youtube[za0005].UnmarshalMsg(bts)
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
			}
		case "Credits":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Credits = nil
			} else {
				if z.Credits == nil {
					z.Credits = new(Credits)
				}
				bts, err = z.Credits.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Images":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Images = nil
			} else {
				if z.Images == nil {
					z.Images = new(Images)
				}
				bts, err = z.Images.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "ReleaseDates":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ReleaseDates = nil
			} else {
				if z.ReleaseDates == nil {
					z.ReleaseDates = new(ReleaseDatesResults)
				}
				var zb0014 uint32
				zb0014, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0014 > 0 {
					zb0014--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Results":
						var zb0015 uint32
						zb0015, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.ReleaseDates.Results) >= int(zb0015) {
							z.ReleaseDates.Results = (z.ReleaseDates.Results)[:zb0015]
						} else {
							z.ReleaseDates.Results = make([]*ReleaseDates, zb0015)
						}
						for za0006 := range z.ReleaseDates.Results {
							if msgp.IsNil(bts) {
								bts, err = msgp.ReadNilBytes(bts)
								if err != nil {
									return
								}
								z.ReleaseDates.Results[za0006] = nil
							} else {
								if z.ReleaseDates.Results[za0006] == nil {
									z.ReleaseDates.Results[za0006] = new(ReleaseDates)
								}
								bts, err = z.ReleaseDates.Results[za0006].UnmarshalMsg(bts)
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
func (z *Movie) Msgsize() (s int) {
	s = 3 + 7 + z.Entity.Msgsize() + 7 + msgp.StringPrefixSize + len(z.IMDBId) + 9 + msgp.StringPrefixSize + len(z.Overview) + 20 + msgp.ArrayHeaderSize
	for za0001 := range z.ProductionCompanies {
		if z.ProductionCompanies[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.ProductionCompanies[za0001].Name)
		}
	}
	s += 8 + msgp.IntSize + 8 + msgp.StringPrefixSize + len(z.TagLine) + 14 + msgp.GuessSize(z.RawPopularity) + 11 + msgp.Float64Size + 16 + msgp.ArrayHeaderSize
	for za0002 := range z.SpokenLanguages {
		if z.SpokenLanguages[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 9 + msgp.StringPrefixSize + len(z.SpokenLanguages[za0002].Iso639_1) + 5 + msgp.StringPrefixSize + len(z.SpokenLanguages[za0002].Name) + 12 + msgp.StringPrefixSize + len(z.SpokenLanguages[za0002].EnglishName)
		}
	}
	s += 12
	if z.ExternalIDs == nil {
		s += msgp.NilSize
	} else {
		s += z.ExternalIDs.Msgsize()
	}
	s += 18
	if z.AlternativeTitles == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 7 + msgp.ArrayHeaderSize
		for za0003 := range z.AlternativeTitles.Titles {
			if z.AlternativeTitles.Titles[za0003] == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 10 + msgp.StringPrefixSize + len(z.AlternativeTitles.Titles[za0003].Iso3166_1) + 6 + msgp.StringPrefixSize + len(z.AlternativeTitles.Titles[za0003].Title)
			}
		}
	}
	s += 13
	if z.Translations == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 13 + msgp.ArrayHeaderSize
		for za0004 := range z.Translations.Translations {
			if z.Translations.Translations[za0004] == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 9 + msgp.StringPrefixSize + len(z.Translations.Translations[za0004].Iso639_1) + 5 + msgp.StringPrefixSize + len(z.Translations.Translations[za0004].Name) + 12 + msgp.StringPrefixSize + len(z.Translations.Translations[za0004].EnglishName)
			}
		}
	}
	s += 9
	if z.Trailers == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 8 + msgp.ArrayHeaderSize
		for za0005 := range z.Trailers.Youtube {
			if z.Trailers.Youtube[za0005] == nil {
				s += msgp.NilSize
			} else {
				s += z.Trailers.Youtube[za0005].Msgsize()
			}
		}
	}
	s += 8
	if z.Credits == nil {
		s += msgp.NilSize
	} else {
		s += z.Credits.Msgsize()
	}
	s += 7
	if z.Images == nil {
		s += msgp.NilSize
	} else {
		s += z.Images.Msgsize()
	}
	s += 13
	if z.ReleaseDates == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 8 + msgp.ArrayHeaderSize
		for za0006 := range z.ReleaseDates.Results {
			if z.ReleaseDates.Results[za0006] == nil {
				s += msgp.NilSize
			} else {
				s += z.ReleaseDates.Results[za0006].Msgsize()
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Movies) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		if z[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Movies) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Movies, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Movie)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Movies) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ReleaseDate) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Certification"
	o = append(o, 0x85, 0xad, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendString(o, z.Certification)
	// string "Iso639_1"
	o = append(o, 0xa8, 0x49, 0x73, 0x6f, 0x36, 0x33, 0x39, 0x5f, 0x31)
	o = msgp.AppendString(o, z.Iso639_1)
	// string "Note"
	o = append(o, 0xa4, 0x4e, 0x6f, 0x74, 0x65)
	o = msgp.AppendString(o, z.Note)
	// string "ReleaseDate"
	o = append(o, 0xab, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.ReleaseDate)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt(o, z.Type)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ReleaseDate) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Certification":
			z.Certification, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Iso639_1":
			z.Iso639_1, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Note":
			z.Note, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ReleaseDate":
			z.ReleaseDate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadIntBytes(bts)
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
func (z *ReleaseDate) Msgsize() (s int) {
	s = 1 + 14 + msgp.StringPrefixSize + len(z.Certification) + 9 + msgp.StringPrefixSize + len(z.Iso639_1) + 5 + msgp.StringPrefixSize + len(z.Note) + 12 + msgp.StringPrefixSize + len(z.ReleaseDate) + 5 + msgp.IntSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ReleaseDates) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Iso3166_1"
	o = append(o, 0x82, 0xa9, 0x49, 0x73, 0x6f, 0x33, 0x31, 0x36, 0x36, 0x5f, 0x31)
	o = msgp.AppendString(o, z.Iso3166_1)
	// string "ReleaseDates"
	o = append(o, 0xac, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.ReleaseDates)))
	for za0001 := range z.ReleaseDates {
		if z.ReleaseDates[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.ReleaseDates[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ReleaseDates) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Iso3166_1":
			z.Iso3166_1, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ReleaseDates":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.ReleaseDates) >= int(zb0002) {
				z.ReleaseDates = (z.ReleaseDates)[:zb0002]
			} else {
				z.ReleaseDates = make([]*ReleaseDate, zb0002)
			}
			for za0001 := range z.ReleaseDates {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.ReleaseDates[za0001] = nil
				} else {
					if z.ReleaseDates[za0001] == nil {
						z.ReleaseDates[za0001] = new(ReleaseDate)
					}
					bts, err = z.ReleaseDates[za0001].UnmarshalMsg(bts)
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
func (z *ReleaseDates) Msgsize() (s int) {
	s = 1 + 10 + msgp.StringPrefixSize + len(z.Iso3166_1) + 13 + msgp.ArrayHeaderSize
	for za0001 := range z.ReleaseDates {
		if z.ReleaseDates[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.ReleaseDates[za0001].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ReleaseDatesResults) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "Results"
	o = append(o, 0x81, 0xa7, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Results)))
	for za0001 := range z.Results {
		if z.Results[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Results[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ReleaseDatesResults) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Results":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Results) >= int(zb0002) {
				z.Results = (z.Results)[:zb0002]
			} else {
				z.Results = make([]*ReleaseDates, zb0002)
			}
			for za0001 := range z.Results {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Results[za0001] = nil
				} else {
					if z.Results[za0001] == nil {
						z.Results[za0001] = new(ReleaseDates)
					}
					bts, err = z.Results[za0001].UnmarshalMsg(bts)
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
func (z *ReleaseDatesResults) Msgsize() (s int) {
	s = 1 + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Results {
		if z.Results[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Results[za0001].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Season) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "ID"
	o = append(o, 0x88, 0xa2, 0x49, 0x44)
	o = msgp.AppendInt(o, z.ID)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Season"
	o = append(o, 0xa6, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e)
	o = msgp.AppendInt(o, z.Season)
	// string "EpisodeCount"
	o = append(o, 0xac, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.EpisodeCount)
	// string "AirDate"
	o = append(o, 0xa7, 0x41, 0x69, 0x72, 0x44, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.AirDate)
	// string "Poster"
	o = append(o, 0xa6, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72)
	o = msgp.AppendString(o, z.Poster)
	// string "ExternalIDs"
	o = append(o, 0xab, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x73)
	if z.ExternalIDs == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.ExternalIDs.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Episodes"
	o = append(o, 0xa8, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Episodes)))
	for za0001 := range z.Episodes {
		if z.Episodes[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Episodes[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Season) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "ID":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Season":
			z.Season, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "EpisodeCount":
			z.EpisodeCount, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "AirDate":
			z.AirDate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Poster":
			z.Poster, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ExternalIDs":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ExternalIDs = nil
			} else {
				if z.ExternalIDs == nil {
					z.ExternalIDs = new(ExternalIDs)
				}
				bts, err = z.ExternalIDs.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Episodes":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Episodes) >= int(zb0002) {
				z.Episodes = (z.Episodes)[:zb0002]
			} else {
				z.Episodes = make(EpisodeList, zb0002)
			}
			for za0001 := range z.Episodes {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Episodes[za0001] = nil
				} else {
					if z.Episodes[za0001] == nil {
						z.Episodes[za0001] = new(Episode)
					}
					bts, err = z.Episodes[za0001].UnmarshalMsg(bts)
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
func (z *Season) Msgsize() (s int) {
	s = 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Name) + 7 + msgp.IntSize + 13 + msgp.IntSize + 8 + msgp.StringPrefixSize + len(z.AirDate) + 7 + msgp.StringPrefixSize + len(z.Poster) + 12
	if z.ExternalIDs == nil {
		s += msgp.NilSize
	} else {
		s += z.ExternalIDs.Msgsize()
	}
	s += 9 + msgp.ArrayHeaderSize
	for za0001 := range z.Episodes {
		if z.Episodes[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Episodes[za0001].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z SeasonList) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		if z[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SeasonList) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(SeasonList, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Season)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z SeasonList) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Show) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 23
	// string "Entity"
	o = append(o, 0xde, 0x0, 0x17, 0xa6, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79)
	o, err = z.Entity.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "EpisodeRunTime"
	o = append(o, 0xae, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x52, 0x75, 0x6e, 0x54, 0x69, 0x6d, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.EpisodeRunTime)))
	for za0001 := range z.EpisodeRunTime {
		o = msgp.AppendInt(o, z.EpisodeRunTime[za0001])
	}
	// string "Genres"
	o = append(o, 0xa6, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Genres)))
	for za0002 := range z.Genres {
		if z.Genres[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "ID"
			o = append(o, 0x82, 0xa2, 0x49, 0x44)
			o = msgp.AppendInt(o, z.Genres[za0002].ID)
			// string "Name"
			o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.Genres[za0002].Name)
		}
	}
	// string "Homepage"
	o = append(o, 0xa8, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Homepage)
	// string "InProduction"
	o = append(o, 0xac, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e)
	o = msgp.AppendBool(o, z.InProduction)
	// string "FirstAirDate"
	o = append(o, 0xac, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x44, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.FirstAirDate)
	// string "LastAirDate"
	o = append(o, 0xab, 0x4c, 0x61, 0x73, 0x74, 0x41, 0x69, 0x72, 0x44, 0x61, 0x74, 0x65)
	o = msgp.AppendString(o, z.LastAirDate)
	// string "Networks"
	o = append(o, 0xa8, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Networks)))
	for za0003 := range z.Networks {
		if z.Networks[za0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "ID"
			o = append(o, 0x82, 0xa2, 0x49, 0x44)
			o = msgp.AppendInt(o, z.Networks[za0003].ID)
			// string "Name"
			o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.Networks[za0003].Name)
		}
	}
	// string "NumberOfEpisodes"
	o = append(o, 0xb0, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73)
	o = msgp.AppendInt(o, z.NumberOfEpisodes)
	// string "NumberOfSeasons"
	o = append(o, 0xaf, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73)
	o = msgp.AppendInt(o, z.NumberOfSeasons)
	// string "OriginalName"
	o = append(o, 0xac, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.OriginalName)
	// string "OriginCountry"
	o = append(o, 0xad, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79)
	o = msgp.AppendArrayHeader(o, uint32(len(z.OriginCountry)))
	for za0004 := range z.OriginCountry {
		o = msgp.AppendString(o, z.OriginCountry[za0004])
	}
	// string "Overview"
	o = append(o, 0xa8, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77)
	o = msgp.AppendString(o, z.Overview)
	// string "RawPopularity"
	o = append(o, 0xad, 0x52, 0x61, 0x77, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79)
	o, err = msgp.AppendIntf(o, z.RawPopularity)
	if err != nil {
		return
	}
	// string "Popularity"
	o = append(o, 0xaa, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79)
	o = msgp.AppendFloat64(o, z.Popularity)
	// string "ProductionCompanies"
	o = append(o, 0xb3, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.ProductionCompanies)))
	for za0005 := range z.ProductionCompanies {
		if z.ProductionCompanies[za0005] == nil {
			o = msgp.AppendNil(o)
		} else {
			// map header, size 2
			// string "ID"
			o = append(o, 0x82, 0xa2, 0x49, 0x44)
			o = msgp.AppendInt(o, z.ProductionCompanies[za0005].ID)
			// string "Name"
			o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
			o = msgp.AppendString(o, z.ProductionCompanies[za0005].Name)
		}
	}
	// string "Status"
	o = append(o, 0xa6, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendString(o, z.Status)
	// string "ExternalIDs"
	o = append(o, 0xab, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x73)
	if z.ExternalIDs == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.ExternalIDs.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Translations"
	o = append(o, 0xac, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73)
	if z.Translations == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 1
		// string "Translations"
		o = append(o, 0x81, 0xac, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.Translations.Translations)))
		for za0006 := range z.Translations.Translations {
			if z.Translations.Translations[za0006] == nil {
				o = msgp.AppendNil(o)
			} else {
				// map header, size 3
				// string "Iso639_1"
				o = append(o, 0x83, 0xa8, 0x49, 0x73, 0x6f, 0x36, 0x33, 0x39, 0x5f, 0x31)
				o = msgp.AppendString(o, z.Translations.Translations[za0006].Iso639_1)
				// string "Name"
				o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
				o = msgp.AppendString(o, z.Translations.Translations[za0006].Name)
				// string "EnglishName"
				o = append(o, 0xab, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65)
				o = msgp.AppendString(o, z.Translations.Translations[za0006].EnglishName)
			}
		}
	}
	// string "AlternativeTitles"
	o = append(o, 0xb1, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73)
	if z.AlternativeTitles == nil {
		o = msgp.AppendNil(o)
	} else {
		// map header, size 1
		// string "Titles"
		o = append(o, 0x81, 0xa6, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x73)
		o = msgp.AppendArrayHeader(o, uint32(len(z.AlternativeTitles.Titles)))
		for za0007 := range z.AlternativeTitles.Titles {
			if z.AlternativeTitles.Titles[za0007] == nil {
				o = msgp.AppendNil(o)
			} else {
				// map header, size 2
				// string "Iso3166_1"
				o = append(o, 0x82, 0xa9, 0x49, 0x73, 0x6f, 0x33, 0x31, 0x36, 0x36, 0x5f, 0x31)
				o = msgp.AppendString(o, z.AlternativeTitles.Titles[za0007].Iso3166_1)
				// string "Title"
				o = append(o, 0xa5, 0x54, 0x69, 0x74, 0x6c, 0x65)
				o = msgp.AppendString(o, z.AlternativeTitles.Titles[za0007].Title)
			}
		}
	}
	// string "Credits"
	o = append(o, 0xa7, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73)
	if z.Credits == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Credits.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Images"
	o = append(o, 0xa6, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73)
	if z.Images == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Images.MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "Seasons"
	o = append(o, 0xa7, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Seasons)))
	for za0008 := range z.Seasons {
		if z.Seasons[za0008] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Seasons[za0008].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Show) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Entity":
			bts, err = z.Entity.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "EpisodeRunTime":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.EpisodeRunTime) >= int(zb0002) {
				z.EpisodeRunTime = (z.EpisodeRunTime)[:zb0002]
			} else {
				z.EpisodeRunTime = make([]int, zb0002)
			}
			for za0001 := range z.EpisodeRunTime {
				z.EpisodeRunTime[za0001], bts, err = msgp.ReadIntBytes(bts)
				if err != nil {
					return
				}
			}
		case "Genres":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Genres) >= int(zb0003) {
				z.Genres = (z.Genres)[:zb0003]
			} else {
				z.Genres = make([]*Genre, zb0003)
			}
			for za0002 := range z.Genres {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Genres[za0002] = nil
				} else {
					if z.Genres[za0002] == nil {
						z.Genres[za0002] = new(Genre)
					}
					var zb0004 uint32
					zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zb0004 > 0 {
						zb0004--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "ID":
							z.Genres[za0002].ID, bts, err = msgp.ReadIntBytes(bts)
							if err != nil {
								return
							}
						case "Name":
							z.Genres[za0002].Name, bts, err = msgp.ReadStringBytes(bts)
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
				}
			}
		case "Homepage":
			z.Homepage, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "InProduction":
			z.InProduction, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "FirstAirDate":
			z.FirstAirDate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "LastAirDate":
			z.LastAirDate, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Networks":
			var zb0005 uint32
			zb0005, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Networks) >= int(zb0005) {
				z.Networks = (z.Networks)[:zb0005]
			} else {
				z.Networks = make([]*IDName, zb0005)
			}
			for za0003 := range z.Networks {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Networks[za0003] = nil
				} else {
					if z.Networks[za0003] == nil {
						z.Networks[za0003] = new(IDName)
					}
					var zb0006 uint32
					zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zb0006 > 0 {
						zb0006--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "ID":
							z.Networks[za0003].ID, bts, err = msgp.ReadIntBytes(bts)
							if err != nil {
								return
							}
						case "Name":
							z.Networks[za0003].Name, bts, err = msgp.ReadStringBytes(bts)
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
				}
			}
		case "NumberOfEpisodes":
			z.NumberOfEpisodes, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "NumberOfSeasons":
			z.NumberOfSeasons, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "OriginalName":
			z.OriginalName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "OriginCountry":
			var zb0007 uint32
			zb0007, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.OriginCountry) >= int(zb0007) {
				z.OriginCountry = (z.OriginCountry)[:zb0007]
			} else {
				z.OriginCountry = make([]string, zb0007)
			}
			for za0004 := range z.OriginCountry {
				z.OriginCountry[za0004], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "Overview":
			z.Overview, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RawPopularity":
			z.RawPopularity, bts, err = msgp.ReadIntfBytes(bts)
			if err != nil {
				return
			}
		case "Popularity":
			z.Popularity, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "ProductionCompanies":
			var zb0008 uint32
			zb0008, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.ProductionCompanies) >= int(zb0008) {
				z.ProductionCompanies = (z.ProductionCompanies)[:zb0008]
			} else {
				z.ProductionCompanies = make([]*IDName, zb0008)
			}
			for za0005 := range z.ProductionCompanies {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.ProductionCompanies[za0005] = nil
				} else {
					if z.ProductionCompanies[za0005] == nil {
						z.ProductionCompanies[za0005] = new(IDName)
					}
					var zb0009 uint32
					zb0009, bts, err = msgp.ReadMapHeaderBytes(bts)
					if err != nil {
						return
					}
					for zb0009 > 0 {
						zb0009--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							return
						}
						switch msgp.UnsafeString(field) {
						case "ID":
							z.ProductionCompanies[za0005].ID, bts, err = msgp.ReadIntBytes(bts)
							if err != nil {
								return
							}
						case "Name":
							z.ProductionCompanies[za0005].Name, bts, err = msgp.ReadStringBytes(bts)
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
				}
			}
		case "Status":
			z.Status, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ExternalIDs":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.ExternalIDs = nil
			} else {
				if z.ExternalIDs == nil {
					z.ExternalIDs = new(ExternalIDs)
				}
				bts, err = z.ExternalIDs.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Translations":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Translations = nil
			} else {
				if z.Translations == nil {
					z.Translations = new(struct {
						Translations []*Language `json:"translations"`
					})
				}
				var zb0010 uint32
				zb0010, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0010 > 0 {
					zb0010--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Translations":
						var zb0011 uint32
						zb0011, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.Translations.Translations) >= int(zb0011) {
							z.Translations.Translations = (z.Translations.Translations)[:zb0011]
						} else {
							z.Translations.Translations = make([]*Language, zb0011)
						}
						for za0006 := range z.Translations.Translations {
							if msgp.IsNil(bts) {
								bts, err = msgp.ReadNilBytes(bts)
								if err != nil {
									return
								}
								z.Translations.Translations[za0006] = nil
							} else {
								if z.Translations.Translations[za0006] == nil {
									z.Translations.Translations[za0006] = new(Language)
								}
								var zb0012 uint32
								zb0012, bts, err = msgp.ReadMapHeaderBytes(bts)
								if err != nil {
									return
								}
								for zb0012 > 0 {
									zb0012--
									field, bts, err = msgp.ReadMapKeyZC(bts)
									if err != nil {
										return
									}
									switch msgp.UnsafeString(field) {
									case "Iso639_1":
										z.Translations.Translations[za0006].Iso639_1, bts, err = msgp.ReadStringBytes(bts)
										if err != nil {
											return
										}
									case "Name":
										z.Translations.Translations[za0006].Name, bts, err = msgp.ReadStringBytes(bts)
										if err != nil {
											return
										}
									case "EnglishName":
										z.Translations.Translations[za0006].EnglishName, bts, err = msgp.ReadStringBytes(bts)
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
							}
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "AlternativeTitles":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.AlternativeTitles = nil
			} else {
				if z.AlternativeTitles == nil {
					z.AlternativeTitles = new(struct {
						Titles []*AlternativeTitle `json:"results"`
					})
				}
				var zb0013 uint32
				zb0013, bts, err = msgp.ReadMapHeaderBytes(bts)
				if err != nil {
					return
				}
				for zb0013 > 0 {
					zb0013--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						return
					}
					switch msgp.UnsafeString(field) {
					case "Titles":
						var zb0014 uint32
						zb0014, bts, err = msgp.ReadArrayHeaderBytes(bts)
						if err != nil {
							return
						}
						if cap(z.AlternativeTitles.Titles) >= int(zb0014) {
							z.AlternativeTitles.Titles = (z.AlternativeTitles.Titles)[:zb0014]
						} else {
							z.AlternativeTitles.Titles = make([]*AlternativeTitle, zb0014)
						}
						for za0007 := range z.AlternativeTitles.Titles {
							if msgp.IsNil(bts) {
								bts, err = msgp.ReadNilBytes(bts)
								if err != nil {
									return
								}
								z.AlternativeTitles.Titles[za0007] = nil
							} else {
								if z.AlternativeTitles.Titles[za0007] == nil {
									z.AlternativeTitles.Titles[za0007] = new(AlternativeTitle)
								}
								var zb0015 uint32
								zb0015, bts, err = msgp.ReadMapHeaderBytes(bts)
								if err != nil {
									return
								}
								for zb0015 > 0 {
									zb0015--
									field, bts, err = msgp.ReadMapKeyZC(bts)
									if err != nil {
										return
									}
									switch msgp.UnsafeString(field) {
									case "Iso3166_1":
										z.AlternativeTitles.Titles[za0007].Iso3166_1, bts, err = msgp.ReadStringBytes(bts)
										if err != nil {
											return
										}
									case "Title":
										z.AlternativeTitles.Titles[za0007].Title, bts, err = msgp.ReadStringBytes(bts)
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
							}
						}
					default:
						bts, err = msgp.Skip(bts)
						if err != nil {
							return
						}
					}
				}
			}
		case "Credits":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Credits = nil
			} else {
				if z.Credits == nil {
					z.Credits = new(Credits)
				}
				bts, err = z.Credits.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Images":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Images = nil
			} else {
				if z.Images == nil {
					z.Images = new(Images)
				}
				bts, err = z.Images.UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "Seasons":
			var zb0016 uint32
			zb0016, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Seasons) >= int(zb0016) {
				z.Seasons = (z.Seasons)[:zb0016]
			} else {
				z.Seasons = make(SeasonList, zb0016)
			}
			for za0008 := range z.Seasons {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Seasons[za0008] = nil
				} else {
					if z.Seasons[za0008] == nil {
						z.Seasons[za0008] = new(Season)
					}
					bts, err = z.Seasons[za0008].UnmarshalMsg(bts)
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
func (z *Show) Msgsize() (s int) {
	s = 3 + 7 + z.Entity.Msgsize() + 15 + msgp.ArrayHeaderSize + (len(z.EpisodeRunTime) * (msgp.IntSize)) + 7 + msgp.ArrayHeaderSize
	for za0002 := range z.Genres {
		if z.Genres[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Genres[za0002].Name)
		}
	}
	s += 9 + msgp.StringPrefixSize + len(z.Homepage) + 13 + msgp.BoolSize + 13 + msgp.StringPrefixSize + len(z.FirstAirDate) + 12 + msgp.StringPrefixSize + len(z.LastAirDate) + 9 + msgp.ArrayHeaderSize
	for za0003 := range z.Networks {
		if z.Networks[za0003] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.Networks[za0003].Name)
		}
	}
	s += 17 + msgp.IntSize + 16 + msgp.IntSize + 13 + msgp.StringPrefixSize + len(z.OriginalName) + 14 + msgp.ArrayHeaderSize
	for za0004 := range z.OriginCountry {
		s += msgp.StringPrefixSize + len(z.OriginCountry[za0004])
	}
	s += 9 + msgp.StringPrefixSize + len(z.Overview) + 14 + msgp.GuessSize(z.RawPopularity) + 11 + msgp.Float64Size + 20 + msgp.ArrayHeaderSize
	for za0005 := range z.ProductionCompanies {
		if z.ProductionCompanies[za0005] == nil {
			s += msgp.NilSize
		} else {
			s += 1 + 3 + msgp.IntSize + 5 + msgp.StringPrefixSize + len(z.ProductionCompanies[za0005].Name)
		}
	}
	s += 7 + msgp.StringPrefixSize + len(z.Status) + 12
	if z.ExternalIDs == nil {
		s += msgp.NilSize
	} else {
		s += z.ExternalIDs.Msgsize()
	}
	s += 13
	if z.Translations == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 13 + msgp.ArrayHeaderSize
		for za0006 := range z.Translations.Translations {
			if z.Translations.Translations[za0006] == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 9 + msgp.StringPrefixSize + len(z.Translations.Translations[za0006].Iso639_1) + 5 + msgp.StringPrefixSize + len(z.Translations.Translations[za0006].Name) + 12 + msgp.StringPrefixSize + len(z.Translations.Translations[za0006].EnglishName)
			}
		}
	}
	s += 18
	if z.AlternativeTitles == nil {
		s += msgp.NilSize
	} else {
		s += 1 + 7 + msgp.ArrayHeaderSize
		for za0007 := range z.AlternativeTitles.Titles {
			if z.AlternativeTitles.Titles[za0007] == nil {
				s += msgp.NilSize
			} else {
				s += 1 + 10 + msgp.StringPrefixSize + len(z.AlternativeTitles.Titles[za0007].Iso3166_1) + 6 + msgp.StringPrefixSize + len(z.AlternativeTitles.Titles[za0007].Title)
			}
		}
	}
	s += 8
	if z.Credits == nil {
		s += msgp.NilSize
	} else {
		s += z.Credits.Msgsize()
	}
	s += 7
	if z.Images == nil {
		s += msgp.NilSize
	} else {
		s += z.Images.Msgsize()
	}
	s += 8 + msgp.ArrayHeaderSize
	for za0008 := range z.Seasons {
		if z.Seasons[za0008] == nil {
			s += msgp.NilSize
		} else {
			s += z.Seasons[za0008].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Shows) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for za0001 := range z {
		if z[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Shows) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Shows, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Show)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Shows) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Trailer) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "Name"
	o = append(o, 0x84, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Size"
	o = append(o, 0xa4, 0x53, 0x69, 0x7a, 0x65)
	o = msgp.AppendString(o, z.Size)
	// string "Source"
	o = append(o, 0xa6, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65)
	o = msgp.AppendString(o, z.Source)
	// string "Type"
	o = append(o, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.Type)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Trailer) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Size":
			z.Size, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Source":
			z.Source, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Type":
			z.Type, bts, err = msgp.ReadStringBytes(bts)
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
func (z *Trailer) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.StringPrefixSize + len(z.Size) + 7 + msgp.StringPrefixSize + len(z.Source) + 5 + msgp.StringPrefixSize + len(z.Type)
	return
}
