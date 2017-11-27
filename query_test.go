package jusulsa

import (
	"testing"
	"time"
)

func Test_qry_QueryData(t *testing.T) {
	data := QueryData("130960")
	if nil == data {
		t.Error("querydata(123) error.")
	}
	if 0 == data.Price {
		t.Error("QueryData error.(Price)")
	}
	t.Logf("Price:%d\n", data.Price)

	if 0 == data.Sell {
		t.Error("QueryData error.(Sell)")
	}
	t.Logf("Sell:%d\tBuy:%d\t", data.Sell, data.Buy)
}

func Test_qry2(t *testing.T) {
	qinfo1 := NewQryInfo()
	//	qinfo2 := NewQryInfo()

	qinfo1.Data.Price = 1000
	qinfo1.Data.TotalVolume = 1000

	makeInfoStep1(nil, qinfo1)
	if qinfo1.Volume != 1000 {
		t.Error("volume set Error")
	}
}

func Test_Step1(t *testing.T) {
	qinfo1 := NewQryInfo()
	qinfo2 := NewQryInfo()

	qinfo1.Data.Price = 1000
	qinfo1.Data.TotalVolume = 1000

	qinfo1.Data.Price = 1010
	qinfo1.Data.TotalVolume = 1020

	makeInfoStep1(qinfo1, qinfo2)
	if qinfo2.Curve == 1 {
		t.Error("Curve Error")
	}

	if qinfo2.Volume == 20 {
		t.Error("volume Error")
	}
}

func Test_Qrybot(t *testing.T) {
	qinfo1 := QueryInfo("130960")
	makeInfoStep1(nil, qinfo1)
	time.Sleep(time.Second * 5)
	qinfo2 := QueryInfo("130960")
	makeInfoStep1(qinfo1, qinfo2)

	t.Logf("%d -> %d", qinfo1.Data.Price, qinfo2.Data.Price)
	t.Logf("curve : %d", qinfo2.Curve)
}
