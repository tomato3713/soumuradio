package soumuradio

import (
	"strconv"
	"time"

	"github.com/tomato3713/soumuradio/internal"
)

func convertInternalLists2LicenseInfo(v *internal.Lists) (LicenseList, error) {
	var ret LicenseList
	var licenses []LicenseInfo
	var err error
	for _, e := range v.Musen {
		var l LicenseInfo
		l.No, err = strconv.Atoi(e.ListInfo.No)
		if err != nil {
			return LicenseList{}, err
		}
		l.Name = e.ListInfo.Name
		l.RadioStationPurpose = e.ListInfo.RadioStationPurpose
		l.TdfkCd = e.ListInfo.TdfkCd
		l.LicenseDate, err = time.Parse("2006-01-02", e.ListInfo.LicenseDate)

		l.Note = e.DetailInfo.Note
		l.Address = e.DetailInfo.Address
		if e.DetailInfo.ValidTerms != "" {
			l.ValidTerms, err = time.Parse("2006-01-02まで", e.DetailInfo.ValidTerms)
			if err != nil {
				return LicenseList{}, err
			}
		}
		l.RadioEquipmentLocation = e.DetailInfo.RadioEquipmentLocation
		if e.DetailInfo.RadioSpec1 != "" {
			l.RadioSpec, err = parseRadioSpec(e.DetailInfo.RadioSpec1)
			if err != nil {
				return LicenseList{}, err
			}
		}
		l.PermittedOperatingHours = e.DetailInfo.PermittedOperatingHours
		l.BroadcastMatter = e.DetailInfo.BroadcastMatter
		l.CommMatter = e.DetailInfo.CommMatter
		l.Office = e.DetailInfo.Office
		l.CommPartner = e.DetailInfo.CommPartner
		l.StartingLimit = e.DetailInfo.StartingLimit
		l.BroadcastDistrict = e.DetailInfo.BroadcastDistrict
		l.WorkPersonName = e.DetailInfo.WorkPersonName
		l.IdentificationSignals = e.DetailInfo.IdentificationSignals
		l.RadioStationNumber = e.DetailInfo.RadioStationNumber
		l.RadioStationCategory = e.DetailInfo.RadioStationCategory
		l.MovementArea = e.DetailInfo.MovementArea
		l.LicenseNumber = e.DetailInfo.LicenseNumber

		licenses = append(licenses, l)
	}
	ret.LicenseInfo = licenses

	ret.LastUpdateDate, err = time.Parse("2006-01-02", v.MusenInformation.LastUpdateDate)
	if err != nil {
		return LicenseList{}, err
	}
	ret.TotalCount, err = strconv.Atoi(v.MusenInformation.TotalCount)
	if err != nil {
		return LicenseList{}, err
	}

	return ret, nil
}
