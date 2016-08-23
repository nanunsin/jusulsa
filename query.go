package jusulsa

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func QueryData(code string) *QryInfo {
	qryinfo := NewQryInfo()
	/*
		qryinfo.Data.Price, qryinfo.Data.TotalVolume = getPV(code)
		qryinfo.Data.Sell, qryinfo.Data.Buy = getSB(code)
	*/
	getData(code, qryinfo)
	return qryinfo
}

func getData(code string, qryinfo *QryInfo) {
	qry := fmt.Sprintf("http://paxnet.moneta.co.kr/stock/stockIntro/stockPrice/immedStockList.jsp?code=%s&wlog_pip=T_stockPrice", code)
	doc, e := goquery.NewDocument(qry)
	if e != nil {
		fmt.Println(e.Error())
	}
	objectPV := doc.Find("#analysis").Find("tbody > tr > td")

	// 현재가
	CPriceT := objectPV.Eq(4).Text()
	// fmt.Println(CPriceT)
	qryinfo.Data.Price = removeChar(CPriceT, ",")

	// 거래량
	CVolumeT := objectPV.Eq(16).Text()
	// fmt.Println(CVolumeT)
	qryinfo.Data.TotalVolume = removeChar(CVolumeT, ",")

	objectSB := doc.Find("#10hoga").Find("tbody > tr").Eq(13).Find("td")

	// 현재가
	CSellT := objectSB.Eq(0).Text()
	qryinfo.Data.Sell = removeChar(CSellT, ",")

	// 거래량
	CBuyT := objectSB.Eq(2).Text()
	qryinfo.Data.Buy = removeChar(CBuyT, ",")
	return
}

func makeInfoStep1(oinfo, cinfo *QryInfo) {
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
