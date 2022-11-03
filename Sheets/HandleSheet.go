// package HandleSheet
package main

import (
	CSVFile "awscostapi/CSVHandle"
	"fmt"
)

func calculate(rows [][]string) [][]string {

	// sum := 0
	// nb := 0

	
	for i := range rows{

		if i == 0 {
			rows[0] = append(rows[0], "Total")
			continue
		}

		servic := rows[i][1]
		fmt.Println(servic)

		// price, err := strconv.Atoi(strings.Replace(rows[i][0], ".", "", -1))
		// if err != nil {
		// 	log.Fatalf("Cannot retrieve price of %s: %s\n", service, err)
		// }

		// price, err := strconv.Atoi(rows[i][2])
		// if err != nil {
		// 	log.Fatalf("Cannot retrieve price of %s: %s\n", servic, err)
		// }

		// total := price
		// fmt.Println("HII")
		// fmt.Println(total)

		// rows[i] = append(rows[i], intToFloatString(price))

		// sum += price

	}

	// rows = append(rows, []string{"", "", "", "Sum", "", intToFloatString(sum)})
	// // rows = append(rows, []string{"", "", "", "Ball Pens", fmt.Sprint(nb), ""})
	// fmt.Println(rows)
	return rows
}

func intToFloatString(n int) string {
	intgr := n / 100
	frac := n - intgr*100
	return fmt.Sprintf("%d.%d", intgr, frac)
}

func main() {
	rows := CSVFile.ReadFile("../AWSBilling.csv")
	rows = calculate(rows)
	CSVFile.WriteOrders("ordersReport.csv", rows)
}

// func Valculate(rows [][]string) [][]string {
// 	for i := range rows {

// 		if i == 0 {
// 			rows[0] = append(rows[0], "Total")
// 			continue
// 		}
// 	}
// return rows
// }
