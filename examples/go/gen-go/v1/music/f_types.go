// Autogenerated by Frugal Compiler (2.22.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package music

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
	frugal "github.com/Workiva/frugal/lib/go"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

func init() {
}

type Minutes float64
type PerfRightsOrg int64

const (
	PerfRightsOrg_ASCAP PerfRightsOrg = 1
	PerfRightsOrg_BMI   PerfRightsOrg = 2
	PerfRightsOrg_SESAC PerfRightsOrg = 3
	PerfRightsOrg_Other PerfRightsOrg = 4
)

func (p PerfRightsOrg) String() string {
	switch p {
	case PerfRightsOrg_ASCAP:
		return "ASCAP"
	case PerfRightsOrg_BMI:
		return "BMI"
	case PerfRightsOrg_SESAC:
		return "SESAC"
	case PerfRightsOrg_Other:
		return "Other"
	}
	return "<UNSET>"
}

func PerfRightsOrgFromString(s string) (PerfRightsOrg, error) {
	switch s {
	case "ASCAP":
		return PerfRightsOrg_ASCAP, nil
	case "BMI":
		return PerfRightsOrg_BMI, nil
	case "SESAC":
		return PerfRightsOrg_SESAC, nil
	case "Other":
		return PerfRightsOrg_Other, nil
	}
	return PerfRightsOrg(0), fmt.Errorf("not a valid PerfRightsOrg string")
}

func (p PerfRightsOrg) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *PerfRightsOrg) UnmarshalText(text []byte) error {
	q, err := PerfRightsOrgFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *PerfRightsOrg) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = PerfRightsOrg(v)
	return nil
}

func (p *PerfRightsOrg) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// Comments (with an @ symbol) will be added to generated code.
type Track struct {
	Title     string        `thrift:"title,1" db:"title" json:"title"`
	Artist    string        `thrift:"artist,2" db:"artist" json:"artist"`
	Publisher string        `thrift:"publisher,3" db:"publisher" json:"publisher"`
	Composer  string        `thrift:"composer,4" db:"composer" json:"composer"`
	Duration  Minutes       `thrift:"duration,5" db:"duration" json:"duration"`
	Pro       PerfRightsOrg `thrift:"pro,6" db:"pro" json:"pro"`
}

func NewTrack() *Track {
	return &Track{}
}

func (p *Track) GetTitle() string {
	return p.Title
}

func (p *Track) GetArtist() string {
	return p.Artist
}

func (p *Track) GetPublisher() string {
	return p.Publisher
}

func (p *Track) GetComposer() string {
	return p.Composer
}

func (p *Track) GetDuration() Minutes {
	return p.Duration
}

func (p *Track) GetPro() PerfRightsOrg {
	return p.Pro
}

func (p *Track) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Track) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Title = v
	}
	return nil
}

func (p *Track) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Artist = v
	}
	return nil
}

func (p *Track) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Publisher = v
	}
	return nil
}

func (p *Track) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.Composer = v
	}
	return nil
}

func (p *Track) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		temp := Minutes(v)
		p.Duration = temp
	}
	return nil
}

func (p *Track) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		temp := PerfRightsOrg(v)
		p.Pro = temp
	}
	return nil
}

func (p *Track) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Track"); err != nil {
		return thrift.PrependError("*music.Track write struct begin error: ", err)
	}
	if err := frugal.WriteString(oprot, p.Title, "title", 1); err != nil {
		return thrift.PrependError("*music.Track::title:1 ", err)
	}
	if err := frugal.WriteString(oprot, p.Artist, "artist", 2); err != nil {
		return thrift.PrependError("*music.Track::artist:2 ", err)
	}
	if err := frugal.WriteString(oprot, p.Publisher, "publisher", 3); err != nil {
		return thrift.PrependError("*music.Track::publisher:3 ", err)
	}
	if err := frugal.WriteString(oprot, p.Composer, "composer", 4); err != nil {
		return thrift.PrependError("*music.Track::composer:4 ", err)
	}
	if err := frugal.WriteDouble(oprot, float64(p.Duration), "duration", 5); err != nil {
		return thrift.PrependError("*music.Track::duration:5 ", err)
	}
	if err := frugal.WriteI32(oprot, int32(p.Pro), "pro", 6); err != nil {
		return thrift.PrependError("*music.Track::pro:6 ", err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Track) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Track(%+v)", *p)
}

// The IDL provides set, list, and map types for representing collections
// of data.  Our Album struct contains a list of Tracks.
type Album struct {
	Tracks   []*Track `thrift:"tracks,1" db:"tracks" json:"tracks"`
	Duration Minutes  `thrift:"duration,2" db:"duration" json:"duration"`
	ASIN     string   `thrift:"ASIN,3" db:"ASIN" json:"ASIN"`
}

func NewAlbum() *Album {
	return &Album{}
}

func (p *Album) GetTracks() []*Track {
	return p.Tracks
}

func (p *Album) GetDuration() Minutes {
	return p.Duration
}

func (p *Album) GetASIN() string {
	return p.ASIN
}

func (p *Album) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Album) ReadField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	p.Tracks = make([]*Track, 0, size)
	for i := 0; i < size; i++ {
		elem0 := NewTrack()
		if err := elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", elem0), err)
		}
		p.Tracks = append(p.Tracks, elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *Album) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		temp := Minutes(v)
		p.Duration = temp
	}
	return nil
}

func (p *Album) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.ASIN = v
	}
	return nil
}

func (p *Album) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Album"); err != nil {
		return thrift.PrependError("*music.Album write struct begin error: ", err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := frugal.WriteDouble(oprot, float64(p.Duration), "duration", 2); err != nil {
		return thrift.PrependError("*music.Album::duration:2 ", err)
	}
	if err := frugal.WriteString(oprot, p.ASIN, "ASIN", 3); err != nil {
		return thrift.PrependError("*music.Album::ASIN:3 ", err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Album) writeField1(oprot thrift.TProtocol) error {
	if err := oprot.WriteFieldBegin("tracks", thrift.LIST, 1); err != nil {
		return thrift.PrependError("*music.Album write field begin error 1:tracks: ", err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Tracks)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Tracks {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError("*music.Album write field end error 1:tracks: ", err)
	}
	return nil
}

func (p *Album) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Album(%+v)", *p)
}

// Exceptions are converted to the native format for each compiled
// language.
type PurchasingError struct {
	Message   string `thrift:"message,1" db:"message" json:"message"`
	ErrorCode int16  `thrift:"error_code,2" db:"error_code" json:"error_code"`
}

func NewPurchasingError() *PurchasingError {
	return &PurchasingError{}
}

func (p *PurchasingError) GetMessage() string {
	return p.Message
}

func (p *PurchasingError) GetErrorCode() int16 {
	return p.ErrorCode
}

func (p *PurchasingError) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PurchasingError) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *PurchasingError) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI16(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.ErrorCode = v
	}
	return nil
}

func (p *PurchasingError) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PurchasingError"); err != nil {
		return thrift.PrependError("*music.PurchasingError write struct begin error: ", err)
	}
	if err := frugal.WriteString(oprot, p.Message, "message", 1); err != nil {
		return thrift.PrependError("*music.PurchasingError::message:1 ", err)
	}
	if err := frugal.WriteI16(oprot, p.ErrorCode, "error_code", 2); err != nil {
		return thrift.PrependError("*music.PurchasingError::error_code:2 ", err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PurchasingError) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PurchasingError(%+v)", *p)
}

func (p *PurchasingError) Error() string {
	return p.String()
}
