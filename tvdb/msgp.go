package tvdb

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *Actor) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "ID"
	o = append(o, 0x85, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, z.ID)
	// string "Image"
	o = append(o, 0xa5, 0x49, 0x6d, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Image)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Role"
	o = append(o, 0xa4, 0x52, 0x6f, 0x6c, 0x65)
	o = msgp.AppendString(o, z.Role)
	// string "SortOrder"
	o = append(o, 0xa9, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72)
	o = msgp.AppendInt(o, z.SortOrder)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Actor) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Image":
			z.Image, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Role":
			z.Role, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SortOrder":
			z.SortOrder, bts, err = msgp.ReadIntBytes(bts)
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
func (z *Actor) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.ID) + 6 + msgp.StringPrefixSize + len(z.Image) + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.StringPrefixSize + len(z.Role) + 10 + msgp.IntSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Banner) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 12
	// string "ID"
	o = append(o, 0x8c, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, z.ID)
	// string "BannerPath"
	o = append(o, 0xaa, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.BannerPath)
	// string "BannerType"
	o = append(o, 0xaa, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.BannerType)
	// string "BannerType2"
	o = append(o, 0xab, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x32)
	o = msgp.AppendString(o, z.BannerType2)
	// string "Colors"
	o = append(o, 0xa6, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x73)
	o = msgp.AppendString(o, z.Colors)
	// string "Language"
	o = append(o, 0xa8, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Language)
	// string "Rating"
	o = append(o, 0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendString(o, z.Rating)
	// string "RatingCount"
	o = append(o, 0xab, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendInt(o, z.RatingCount)
	// string "SeriesName"
	o = append(o, 0xaa, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.SeriesName)
	// string "ThumbnailPath"
	o = append(o, 0xad, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.ThumbnailPath)
	// string "VignettePath"
	o = append(o, 0xac, 0x56, 0x69, 0x67, 0x6e, 0x65, 0x74, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68)
	o = msgp.AppendString(o, z.VignettePath)
	// string "Season"
	o = append(o, 0xa6, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e)
	o = msgp.AppendInt(o, z.Season)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Banner) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "BannerPath":
			z.BannerPath, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "BannerType":
			z.BannerType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "BannerType2":
			z.BannerType2, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Colors":
			z.Colors, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Language":
			z.Language, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Rating":
			z.Rating, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RatingCount":
			z.RatingCount, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "SeriesName":
			z.SeriesName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ThumbnailPath":
			z.ThumbnailPath, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "VignettePath":
			z.VignettePath, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Season":
			z.Season, bts, err = msgp.ReadIntBytes(bts)
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
func (z *Banner) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.ID) + 11 + msgp.StringPrefixSize + len(z.BannerPath) + 11 + msgp.StringPrefixSize + len(z.BannerType) + 12 + msgp.StringPrefixSize + len(z.BannerType2) + 7 + msgp.StringPrefixSize + len(z.Colors) + 9 + msgp.StringPrefixSize + len(z.Language) + 7 + msgp.StringPrefixSize + len(z.Rating) + 12 + msgp.IntSize + 11 + msgp.StringPrefixSize + len(z.SeriesName) + 14 + msgp.StringPrefixSize + len(z.ThumbnailPath) + 13 + msgp.StringPrefixSize + len(z.VignettePath) + 7 + msgp.IntSize
	return
}

// MarshalMsg implements msgp.Marshaler
func (z BannersByRating) MarshalMsg(b []byte) (o []byte, err error) {
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
func (z *BannersByRating) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(BannersByRating, zb0002)
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
				(*z)[zb0001] = new(Banner)
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
func (z BannersByRating) Msgsize() (s int) {
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
func (z BySeasonAndEpisodeNumber) MarshalMsg(b []byte) (o []byte, err error) {
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
func (z *BySeasonAndEpisodeNumber) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(BySeasonAndEpisodeNumber, zb0002)
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
func (z BySeasonAndEpisodeNumber) Msgsize() (s int) {
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
func (z *Episode) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 21
	// string "ID"
	o = append(o, 0xde, 0x0, 0x15, 0xa2, 0x49, 0x44)
	o = msgp.AppendString(o, z.ID)
	// string "Director"
	o = append(o, 0xa8, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72)
	o = msgp.AppendString(o, z.Director)
	// string "EpisodeName"
	o = append(o, 0xab, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.EpisodeName)
	// string "EpisodeNumber"
	o = append(o, 0xad, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.EpisodeNumber)
	// string "FirstAired"
	o = append(o, 0xaa, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendString(o, z.FirstAired)
	// string "GuestStars"
	o = append(o, 0xaa, 0x47, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x72, 0x73)
	o = msgp.AppendString(o, z.GuestStars)
	// string "ImdbID"
	o = append(o, 0xa6, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44)
	o = msgp.AppendString(o, z.ImdbID)
	// string "Language"
	o = append(o, 0xa8, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Language)
	// string "Overview"
	o = append(o, 0xa8, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77)
	o = msgp.AppendString(o, z.Overview)
	// string "Rating"
	o = append(o, 0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendString(o, z.Rating)
	// string "RatingCount"
	o = append(o, 0xab, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendString(o, z.RatingCount)
	// string "SeasonNumber"
	o = append(o, 0xac, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.SeasonNumber)
	// string "Writer"
	o = append(o, 0xa6, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72)
	o = msgp.AppendString(o, z.Writer)
	// string "FileName"
	o = append(o, 0xa8, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.FileName)
	// string "LastUpdated"
	o = append(o, 0xab, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendString(o, z.LastUpdated)
	// string "SeasonID"
	o = append(o, 0xa8, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x44)
	o = msgp.AppendString(o, z.SeasonID)
	// string "SeriesID"
	o = append(o, 0xa8, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x44)
	o = msgp.AppendString(o, z.SeriesID)
	// string "ThumbHeight"
	o = append(o, 0xab, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74)
	o = msgp.AppendString(o, z.ThumbHeight)
	// string "ThumbWidth"
	o = append(o, 0xaa, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x57, 0x69, 0x64, 0x74, 0x68)
	o = msgp.AppendString(o, z.ThumbWidth)
	// string "AbsoluteNumber"
	o = append(o, 0xae, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendInt(o, z.AbsoluteNumber)
	// string "AbsoluteNumberString"
	o = append(o, 0xb4, 0x41, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67)
	o = msgp.AppendString(o, z.AbsoluteNumberString)
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
			z.ID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Director":
			z.Director, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "EpisodeName":
			z.EpisodeName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "EpisodeNumber":
			z.EpisodeNumber, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "FirstAired":
			z.FirstAired, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "GuestStars":
			z.GuestStars, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ImdbID":
			z.ImdbID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Language":
			z.Language, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Overview":
			z.Overview, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Rating":
			z.Rating, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RatingCount":
			z.RatingCount, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SeasonNumber":
			z.SeasonNumber, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Writer":
			z.Writer, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FileName":
			z.FileName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "LastUpdated":
			z.LastUpdated, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SeasonID":
			z.SeasonID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SeriesID":
			z.SeriesID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ThumbHeight":
			z.ThumbHeight, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ThumbWidth":
			z.ThumbWidth, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AbsoluteNumber":
			z.AbsoluteNumber, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "AbsoluteNumberString":
			z.AbsoluteNumberString, bts, err = msgp.ReadStringBytes(bts)
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
func (z *Episode) Msgsize() (s int) {
	s = 3 + 3 + msgp.StringPrefixSize + len(z.ID) + 9 + msgp.StringPrefixSize + len(z.Director) + 12 + msgp.StringPrefixSize + len(z.EpisodeName) + 14 + msgp.IntSize + 11 + msgp.StringPrefixSize + len(z.FirstAired) + 11 + msgp.StringPrefixSize + len(z.GuestStars) + 7 + msgp.StringPrefixSize + len(z.ImdbID) + 9 + msgp.StringPrefixSize + len(z.Language) + 9 + msgp.StringPrefixSize + len(z.Overview) + 7 + msgp.StringPrefixSize + len(z.Rating) + 12 + msgp.StringPrefixSize + len(z.RatingCount) + 13 + msgp.IntSize + 7 + msgp.StringPrefixSize + len(z.Writer) + 9 + msgp.StringPrefixSize + len(z.FileName) + 12 + msgp.StringPrefixSize + len(z.LastUpdated) + 9 + msgp.StringPrefixSize + len(z.SeasonID) + 9 + msgp.StringPrefixSize + len(z.SeriesID) + 12 + msgp.StringPrefixSize + len(z.ThumbHeight) + 11 + msgp.StringPrefixSize + len(z.ThumbWidth) + 15 + msgp.IntSize + 21 + msgp.StringPrefixSize + len(z.AbsoluteNumberString)
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
func (z *Season) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Season"
	o = append(o, 0x82, 0xa6, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e)
	o = msgp.AppendInt(o, z.Season)
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
		case "Season":
			z.Season, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
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
	s = 1 + 7 + msgp.IntSize + 9 + msgp.ArrayHeaderSize
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
	// map header, size 26
	// string "ID"
	o = append(o, 0xde, 0x0, 0x1a, 0xa2, 0x49, 0x44)
	o = msgp.AppendInt(o, z.ID)
	// string "ActorsSimple"
	o = append(o, 0xac, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65)
	o = msgp.AppendString(o, z.ActorsSimple)
	// string "AirsDayOfWeek"
	o = append(o, 0xad, 0x41, 0x69, 0x72, 0x73, 0x44, 0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b)
	o = msgp.AppendString(o, z.AirsDayOfWeek)
	// string "AirsTime"
	o = append(o, 0xa8, 0x41, 0x69, 0x72, 0x73, 0x54, 0x69, 0x6d, 0x65)
	o = msgp.AppendString(o, z.AirsTime)
	// string "ContentRating"
	o = append(o, 0xad, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendString(o, z.ContentRating)
	// string "FirstAired"
	o = append(o, 0xaa, 0x46, 0x69, 0x72, 0x73, 0x74, 0x41, 0x69, 0x72, 0x65, 0x64)
	o = msgp.AppendString(o, z.FirstAired)
	// string "Genre"
	o = append(o, 0xa5, 0x47, 0x65, 0x6e, 0x72, 0x65)
	o = msgp.AppendString(o, z.Genre)
	// string "ImdbID"
	o = append(o, 0xa6, 0x49, 0x6d, 0x64, 0x62, 0x49, 0x44)
	o = msgp.AppendString(o, z.ImdbID)
	// string "Language"
	o = append(o, 0xa8, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65)
	o = msgp.AppendString(o, z.Language)
	// string "Network"
	o = append(o, 0xa7, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b)
	o = msgp.AppendString(o, z.Network)
	// string "NetworkID"
	o = append(o, 0xa9, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44)
	o = msgp.AppendString(o, z.NetworkID)
	// string "Overview"
	o = append(o, 0xa8, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77)
	o = msgp.AppendString(o, z.Overview)
	// string "Rating"
	o = append(o, 0xa6, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67)
	o = msgp.AppendString(o, z.Rating)
	// string "RatingCount"
	o = append(o, 0xab, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendString(o, z.RatingCount)
	// string "RuntimeString"
	o = append(o, 0xad, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67)
	o = msgp.AppendString(o, z.RuntimeString)
	// string "SeriesID"
	o = append(o, 0xa8, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x44)
	o = msgp.AppendString(o, z.SeriesID)
	// string "SeriesName"
	o = append(o, 0xaa, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.SeriesName)
	// string "Status"
	o = append(o, 0xa6, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73)
	o = msgp.AppendString(o, z.Status)
	// string "Banner"
	o = append(o, 0xa6, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72)
	o = msgp.AppendString(o, z.Banner)
	// string "FanArt"
	o = append(o, 0xa6, 0x46, 0x61, 0x6e, 0x41, 0x72, 0x74)
	o = msgp.AppendString(o, z.FanArt)
	// string "LastUpdated"
	o = append(o, 0xab, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendInt(o, z.LastUpdated)
	// string "Poster"
	o = append(o, 0xa6, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72)
	o = msgp.AppendString(o, z.Poster)
	// string "Runtime"
	o = append(o, 0xa7, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt(o, z.Runtime)
	// string "Seasons"
	o = append(o, 0xa7, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Seasons)))
	for za0001 := range z.Seasons {
		if z.Seasons[za0001] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Seasons[za0001].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Banners"
	o = append(o, 0xa7, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Banners)))
	for za0002 := range z.Banners {
		if z.Banners[za0002] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Banners[za0002].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	// string "Actors"
	o = append(o, 0xa6, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Actors)))
	for za0003 := range z.Actors {
		if z.Actors[za0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Actors[za0003].MarshalMsg(o)
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
		case "ID":
			z.ID, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "ActorsSimple":
			z.ActorsSimple, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AirsDayOfWeek":
			z.AirsDayOfWeek, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AirsTime":
			z.AirsTime, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ContentRating":
			z.ContentRating, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FirstAired":
			z.FirstAired, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Genre":
			z.Genre, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "ImdbID":
			z.ImdbID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Language":
			z.Language, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Network":
			z.Network, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "NetworkID":
			z.NetworkID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Overview":
			z.Overview, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Rating":
			z.Rating, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RatingCount":
			z.RatingCount, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RuntimeString":
			z.RuntimeString, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SeriesID":
			z.SeriesID, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "SeriesName":
			z.SeriesName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Status":
			z.Status, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Banner":
			z.Banner, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "FanArt":
			z.FanArt, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "LastUpdated":
			z.LastUpdated, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Poster":
			z.Poster, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Runtime":
			z.Runtime, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "Seasons":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Seasons) >= int(zb0002) {
				z.Seasons = (z.Seasons)[:zb0002]
			} else {
				z.Seasons = make(SeasonList, zb0002)
			}
			for za0001 := range z.Seasons {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Seasons[za0001] = nil
				} else {
					if z.Seasons[za0001] == nil {
						z.Seasons[za0001] = new(Season)
					}
					bts, err = z.Seasons[za0001].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Banners":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Banners) >= int(zb0003) {
				z.Banners = (z.Banners)[:zb0003]
			} else {
				z.Banners = make([]*Banner, zb0003)
			}
			for za0002 := range z.Banners {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Banners[za0002] = nil
				} else {
					if z.Banners[za0002] == nil {
						z.Banners[za0002] = new(Banner)
					}
					bts, err = z.Banners[za0002].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		case "Actors":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Actors) >= int(zb0004) {
				z.Actors = (z.Actors)[:zb0004]
			} else {
				z.Actors = make([]*Actor, zb0004)
			}
			for za0003 := range z.Actors {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Actors[za0003] = nil
				} else {
					if z.Actors[za0003] == nil {
						z.Actors[za0003] = new(Actor)
					}
					bts, err = z.Actors[za0003].UnmarshalMsg(bts)
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
	s = 3 + 3 + msgp.IntSize + 13 + msgp.StringPrefixSize + len(z.ActorsSimple) + 14 + msgp.StringPrefixSize + len(z.AirsDayOfWeek) + 9 + msgp.StringPrefixSize + len(z.AirsTime) + 14 + msgp.StringPrefixSize + len(z.ContentRating) + 11 + msgp.StringPrefixSize + len(z.FirstAired) + 6 + msgp.StringPrefixSize + len(z.Genre) + 7 + msgp.StringPrefixSize + len(z.ImdbID) + 9 + msgp.StringPrefixSize + len(z.Language) + 8 + msgp.StringPrefixSize + len(z.Network) + 10 + msgp.StringPrefixSize + len(z.NetworkID) + 9 + msgp.StringPrefixSize + len(z.Overview) + 7 + msgp.StringPrefixSize + len(z.Rating) + 12 + msgp.StringPrefixSize + len(z.RatingCount) + 14 + msgp.StringPrefixSize + len(z.RuntimeString) + 9 + msgp.StringPrefixSize + len(z.SeriesID) + 11 + msgp.StringPrefixSize + len(z.SeriesName) + 7 + msgp.StringPrefixSize + len(z.Status) + 7 + msgp.StringPrefixSize + len(z.Banner) + 7 + msgp.StringPrefixSize + len(z.FanArt) + 12 + msgp.IntSize + 7 + msgp.StringPrefixSize + len(z.Poster) + 8 + msgp.IntSize + 8 + msgp.ArrayHeaderSize
	for za0001 := range z.Seasons {
		if z.Seasons[za0001] == nil {
			s += msgp.NilSize
		} else {
			s += z.Seasons[za0001].Msgsize()
		}
	}
	s += 8 + msgp.ArrayHeaderSize
	for za0002 := range z.Banners {
		if z.Banners[za0002] == nil {
			s += msgp.NilSize
		} else {
			s += z.Banners[za0002].Msgsize()
		}
	}
	s += 7 + msgp.ArrayHeaderSize
	for za0003 := range z.Actors {
		if z.Actors[za0003] == nil {
			s += msgp.NilSize
		} else {
			s += z.Actors[za0003].Msgsize()
		}
	}
	return
}
