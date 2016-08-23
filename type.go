package jusulsa

import "time"

// define types
const (
	OPI_None = 0
	OPI_SS   = 1
	OPI_BB   = 2
)

// QryData 는 수집한 data
type QryData struct {
	Price       int
	TotalVolume int
	Sell, Buy   int
}

// QryInfo 는 수집한 data를 가지고 추가로 정제한 정보
type QryInfo struct {
	TimeStamp        time.Time
	Data             *QryData
	Volume           int
	VolumeRatio      float32
	TotalVolumeRatio float32
	SBRatio          float32
	Curve            int
	Opinion          int
}

// NewQryInfo : make new Query Infomation one.
func NewQryInfo() *QryInfo {
	qrydata := &QryData{
		Price:       0,
		TotalVolume: 0,
		Sell:        0,
		Buy:         0,
	}

	qryinfo := &QryInfo{
		TimeStamp:        time.Now(),
		Data:             qrydata,
		Volume:           0,
		VolumeRatio:      0.0,
		TotalVolumeRatio: 0.0,
		SBRatio:          0.0,
		Curve:            0,
		Opinion:          OPI_None,
	}

	return qryinfo
}
