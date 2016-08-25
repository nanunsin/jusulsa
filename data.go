package jusulsa

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// WriteCSV make file
func WriteCSV(bot *Mark1, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		// Error
		log.Fatal("Create file Error:", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	// Make CSV []String
	for i, value := range bot.ObjInfo {
		wdata := make([]string, 10)

		wdata[0] = fmt.Sprint(i)
		wdata[1] = getTimeStamp(value.TimeStamp, 10)
		wdata[2] = fmt.Sprint(value.Data.Price)
		wdata[3] = fmt.Sprint(value.Volume)
		wdata[4] = fmt.Sprint(value.Data.TotalVolume)
		wdata[5] = fmt.Sprint(value.Curve)
		wdata[6] = fmt.Sprint(value.Data.Sell)
		wdata[7] = fmt.Sprint(value.Data.Buy)
		wdata[8] = fmt.Sprint(value.SBRatio)

		writer.Write(wdata)
	}
	defer writer.Flush()
}
