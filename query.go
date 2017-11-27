package jusulsa

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func QueryData(code string) *QryData {

	qryData := NewQryData()
	/*
		qryinfo.Data.Price, qryinfo.Data.TotalVolume = getPV(code)
		qryinfo.Data.Sell, qryinfo.Data.Buy = getSB(code)
	*/
	getData(code, qryData)
	return qryData
}

func QueryInfo(code string) *QryInfo {
	qryinfo := NewQryInfo()

	getData(code, qryinfo.Data)
	return qryinfo
}

func getData(code string, data *QryData) {
	qry := fmt.Sprintf("http://finance.daum.net/item/quote.daum?code=%s", code)
	doc, e := goquery.NewDocument(qry)
	if e != nil {
		fmt.Println(e.Error())
	}
	objectPV := doc.Find(".leftDiv").Find("tbody > tr > td")

	// 현재가
	CPriceT := objectPV.Eq(1).Text()
	//fmt.Println(CPriceT)
	data.Price = removeChar(CPriceT, ",")

	// 거래량
	CVolumeT := objectPV.Eq(13).Text()
	// fmt.Println(CVolumeT)
	data.TotalVolume = removeChar(CVolumeT, ",")

	objectSB := doc.Find("#price10StepBody").Find("tfoot > tr").Eq(0).Find("td")

	// 팔자
	CSellT := objectSB.Eq(0).Text()
	//fmt.Println(CSellT)
	data.Sell = removeChar(CSellT, ",")

	// 사자
	CBuyT := objectSB.Eq(2).Text()
	//fmt.Println(CBuyT)
	data.Buy = removeChar(CBuyT, ",")
	return
}

func makeInfoStep1(oinfo, cinfo *QryInfo) {

	// calc SBRatio
	if cinfo.Data.Buy > 0 {
		cinfo.SBRatio = float32(cinfo.Data.Sell) / float32(cinfo.Data.Buy)
	}

	// Volume
	if nil == oinfo {
		cinfo.Volume = cinfo.Data.TotalVolume
		return
	}

	cinfo.Volume = cinfo.Data.TotalVolume - oinfo.Data.TotalVolume
	if cinfo.Data.Price > oinfo.Data.Price {
		if oinfo.Curve >= 0 {
			cinfo.Curve = oinfo.Curve + 1
		} else {
			cinfo.Curve = 0
		}
	} else if cinfo.Data.Price == oinfo.Data.Price {
		cinfo.Curve = oinfo.Curve
	} else { // <
		if oinfo.Curve <= 0 {
			cinfo.Curve = oinfo.Curve - 1
		} else {
			cinfo.Curve = 0
		}
	}
}
