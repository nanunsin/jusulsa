package jusulsa

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func QueryData(code string) *QryInfo {
	qryinfo := NewQryInfo()
	qryinfo.Data.Price, qryinfo.Data.TotalVolume = getPV(code)
	qryinfo.Data.Sell, qryinfo.Data.Buy = getSB(code)
	return qryinfo
}

func getPV(code string) (price int, volume int) {
	qry := fmt.Sprintf("http://paxnet.moneta.co.kr/stock/stockIntro/stockPrice/immedStockList.jsp?code=%s&wlog_pip=T_stockPrice", code)
	doc, e := goquery.NewDocument(qry)
	if e != nil {
		fmt.Println(e.Error())
	}
	object := doc.Find("#analysis").Find("tbody > tr > td")

	// 현재가
	CPriceT := object.Eq(4).Text()
	// fmt.Println(CPriceT)
	price = removeChar(CPriceT, ",")

	// 거래량
	CVolumeT := object.Eq(16).Text()
	// fmt.Println(CVolumeT)
	volume = removeChar(CVolumeT, ",")
	return
}

func getSB(code string) (sell, buy int) {
	return 0, 0
}

func makeInfoStep1(oinfo, cinfo *QryInfo) {
	if nil == oinfo {
		cinfo.Volume = cinfo.Data.TotalVolume
	} else {
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
}
