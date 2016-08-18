package jusulsa

// define types
// QryData 는 수집한 data
type QryData struct {
	Price       int
	TotalVolume int
	Sell, Buy   int
}

// QryInfo 는 수집한 data를 가지고 추가로 정제한 정보
type QryInfo struct {
	TimeStr          string
	Data             *QryData
	Volume           int
	VolumeRatio      float32
	TotalVolumeRatio float32
	SBRatio          float32
	Curve            int
}

func NewQryInfo() *QryInfo {
	qrydata := &QryData{
		Price:       0,
		TotalVolume: 0,
		Sell:        0,
		Buy:         0,
	}

	qryinfo := &QryInfo{
		TimeStr:          getCurTimeString(),
		Data:             qrydata,
		Volume:           0,
		VolumeRatio:      0.0,
		TotalVolumeRatio: 0.0,
		SBRatio:          0.0,
		Curve:            0,
	}

	return qryinfo
}
