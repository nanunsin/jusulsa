package jusulsa

import "time"

// define types
const (
	OpinionNone = 0
	OpinionSS   = 1
	OpinionBB   = 2
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
	Index            int
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
		Index:            0,
		Data:             qrydata,
		Volume:           0,
		VolumeRatio:      0.0,
		TotalVolumeRatio: 0.0,
		SBRatio:          0.0,
		Curve:            0,
		Opinion:          OpinionNone,
	}

	return qryinfo
}

// AnalyzeRule be used in analyze.go and analyzerule.go
type AnalyzeRule struct {
	Sensitivity int
	Price       int
	VolumeRatio int
	Curve       int
}
