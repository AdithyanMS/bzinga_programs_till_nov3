package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func schedule(db *sql.DB) {
	endSpot := time.Now()
	endHourly := time.Now()
	endTwh := time.Now()
	endTfh := time.Now()
	startSpot := time.Now()
	var durationSpot int
	updateSpot := time.Now()
	var intervalSpot int
	startHourly := time.Now()
	var durationHourly int
	updateHourly := time.Now()
	var intervalHourly int
	startTwh := time.Now()
	var durationTwh int
	startTfh := time.Now()
	var durationTfh int
	updateTwh := time.Now()
	updateTfh := time.Now()
	var intervalTwh int
	var intervalTfh int

	var count int
	var id int
	var ref string
	var category string
	var prod_name string
	var desc string
	var mrp int
	out1 := "insert into spottable values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
	out2 := "insert into hourlytable values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
	out3 := "insert into twelvehourtable values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
	out4 := "insert into twentyfourhourtable values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"

	rowLine := db.QueryRow("select * from logic1 where id=$1", 1)
	err1 := rowLine.Scan(&id, &count, &startSpot, &durationSpot, &intervalSpot, &endSpot)
	if err1 != nil {
		log.Fatal(err1)
	}

	rowLine2 := db.QueryRow("select * from logic1 where id=$1", 2)
	err2 := rowLine2.Scan(&id, &count, &startHourly, &durationHourly, &intervalHourly, &endHourly)
	if err2 != nil {
		log.Fatal(err2)
	}

	rowLine3 := db.QueryRow("select * from logic1 where id=$1", 3)
	err3 := rowLine3.Scan(&id, &count, &startTwh, &durationTwh, &intervalTwh, &endTwh)
	if err3 != nil {
		log.Fatal(err3)
	}

	rowLine4 := db.QueryRow("select * from logic1 where id=$1", 4)
	err4 := rowLine4.Scan(&id, &count, &startTfh, &durationTfh, &intervalTfh, &endTfh)
	if err4 != nil {
		log.Fatal(err4)
	}

	rows, err5 := db.Query("select * from inputtable order by random()")
	if err5 != nil {
		log.Fatal(err5)
	}

	spotIndex := 1
	spotAuction := 1
	hourIndex := 1
	hourAuction := 1
	twelveIndex := 1
	twelveAuction := 1
	twentyFourIndex := 1
	twentyFourAuction := 1
	sf := false
	hf := false
	twf := true
	tff := true

	for rows.Next() {
		err := rows.Scan(&ref, &category, &prod_name, &desc, &mrp)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("spot: ", spotIndex, " hour: ", hourIndex, "twelve: ", twelveIndex, "twentyFour: ", twentyFourIndex)

		if mrp >= 1 && mrp <= 999 {
			// if int(startSpot.Month()) >= int(endSpot.Month()) && startSpot.Day() >= endSpot.Day() {
			// 	if startSpot.Year() >= endSpot.Year() && startSpot.Hour() >= endSpot.Hour() {
			// 		sf = false
			// 		continue
			// 	}
			// }

			if startSpot.After(endSpot) {
				sf = false
				continue
			}

			if startSpot.Hour() >= 0 && startSpot.Hour() < 9 {
				nt := startSpot.String()
				ntl := strings.Split(nt, " ")
				ntl[1] = "09:00:00"
				ntl = ntl[:2]
				nt = ntl[0] + "T" + ntl[1] + "+00:00"
				start2, _ := time.Parse("2006-01-02T15:04:05Z07:00", nt)
				startSpot = start2

			}

			updateSpot = startSpot.Add(time.Duration(durationSpot) * time.Minute)
			base := 0.99 * float64(mrp)
			err, _ := db.Exec(out1, spotIndex, spotAuction, startSpot, updateSpot, 1, 1, ref, category, prod_name, desc, mrp, 1, int(base))

			if err != nil {
				fmt.Println(err)
			}
			spotIndex++
			if spotIndex%count == 0 {
				spotAuction++
				startSpot = startSpot.Add(time.Duration(intervalSpot) * time.Minute)
			}

		} else if mrp >= 1000 && mrp <= 2999 {
			// if int(startHourly.Month()) >= int(endHourly.Month()) && startHourly.Day() >= endHourly.Day() {
			// 	if startHourly.Year() >= endHourly.Year() && startHourly.Hour() >= endHourly.Hour() {
			// 		hf = false
			// 		continue
			// 	}
			// }

			if startHourly.After(endHourly) {
				hf = false
				continue
			}

			if startHourly.Hour() >= 0 && startHourly.Hour() < 9 {
				nt := startHourly.String()
				ntl := strings.Split(nt, " ")
				ntl[1] = "09:00:00"
				ntl = ntl[:2]
				nt = ntl[0] + "T" + ntl[1] + "+00:00"
				start2, _ := time.Parse("2006-01-02T15:04:05Z07:00", nt)
				startHourly = start2

			}
			updateHourly = startHourly.Add(time.Duration(durationHourly) * time.Minute)
			base := 0.99 * float64(mrp)
			err, _ := db.Exec(out2, hourIndex, hourAuction, startHourly, updateHourly, 1, 1, ref, category, prod_name, desc, mrp, 1, int(base))

			if err != nil {
				fmt.Println(err)
			}
			hourIndex++
			if hourIndex%count == 0 {
				hourAuction++
				startHourly = startHourly.Add(time.Duration(intervalHourly) * time.Minute)
			}

		} else if mrp >= 3000 && mrp <= 4999 {
			// if int(startTwh.Month()) >= int(endTwh.Month()) && startTwh.Day() >= endTwh.Day() {
			// 	if startTwh.Year() >= endTwh.Year() && startTwh.Hour() >= endTwh.Hour() {
			// 		twf = false
			// 		continue
			// 	}
			// }

			if startTwh.After(endTwh) {
				twf = false
				continue
			}

			updateTwh = startTwh.Add(time.Duration(durationTwh) * time.Minute)
			base := 0.99 * float64(mrp)
			err, _ := db.Exec(out3, twelveIndex, twelveAuction, startTwh, updateTwh, 1, 1, ref, category, prod_name, desc, mrp, 1, int(base))

			if err != nil {
				fmt.Println(err)
			}
			twelveIndex++
			if twelveIndex%count == 0 {
				twelveAuction++
				startTwh = startTwh.Add(time.Duration(intervalTwh) * time.Minute)
			}

		} else if mrp >= 5000 && mrp <= 9999 {
			// if int(startTfh.Month()) >= int(endTfh.Month()) && startTfh.Day() >= endTfh.Day() {
			// 	if startTfh.Year() >= endTfh.Year() && startTfh.Hour() >= endTfh.Hour() {
			// 		tff = false
			// 		continue
			// 	}
			// }
			if startTfh.After(endTfh) {
				tff = false
				continue
			}

			updateTfh = startTfh.Add(time.Duration(durationTfh) * time.Minute)
			base := 0.99 * float64(mrp)
			err, _ := db.Exec(out4, twentyFourIndex, twentyFourAuction, startTfh, updateTfh, 1, 1, ref, category, prod_name, desc, mrp, 1, int(base))
			if err != nil {
				fmt.Println(err)
			}
			twentyFourIndex++
			if twentyFourIndex%count == 0 {
				twentyFourAuction++
				startTfh = startTfh.Add(time.Duration(intervalTfh) * time.Minute)
			}

		}

	}
	var tbase int
	for sf {
		spotrows, spoterr := db.Query("select ref,category,product,description,mrp,base from spottable order by random();")
		if spoterr != nil {
			log.Fatal("error selecting from spottable", spoterr)
		}
		for spotrows.Next() {
			fmt.Println("spot: ", spotIndex)
			serr := spotrows.Scan(&ref, &category, &prod_name, &desc, &mrp, &tbase)
			if serr != nil {
				log.Fatal("error scanning spot", serr)
			}
			// if int(startSpot.Month()) >= int(endSpot.Month()) && startSpot.Day() >= endSpot.Day() {
			// 	if startSpot.Year() >= endSpot.Year() && startSpot.Hour() >= endSpot.Hour() {
			// 		sf = false
			// 		break
			// 	}
			// }
			if startSpot.After(endSpot) {
				sf = false
				break
			}

			if startSpot.Hour() >= 0 && startSpot.Hour() < 9 {
				nt := startSpot.String()
				ntl := strings.Split(nt, " ")
				ntl[1] = "09:00:00"
				ntl = ntl[:2]
				nt = ntl[0] + "T" + ntl[1] + "+00:00"
				start2, _ := time.Parse("2006-01-02T15:04:05Z07:00", nt)
				startSpot = start2

			}
			updateSpot = startSpot.Add(time.Duration(durationSpot) * time.Minute)
			err, _ := db.Exec(out1, spotIndex, spotAuction, startSpot, updateSpot, 1, 1, ref, category, prod_name, desc, mrp, 1, tbase)

			if err != nil {
				fmt.Println(err)
			}
			spotIndex++
			if spotIndex%count == 0 {
				spotAuction++
				startSpot = startSpot.Add(time.Duration(intervalSpot) * time.Minute)
			}
		}

	}

	for hf {
		hourrows, hourerr := db.Query("select ref,category,product,description,mrp,base from hourlytable order by random();")
		if hourerr != nil {
			log.Fatal("error selecting from hourlytable", hourerr)
		}
		for hourrows.Next() {
			fmt.Println("hour: ", hourIndex)
			herr := hourrows.Scan(&ref, &category, &prod_name, &desc, &mrp, &tbase)
			if herr != nil {
				log.Fatal("error scanning hour", herr)
			}
			// if int(startHourly.Month()) >= int(endHourly.Month()) && startHourly.Day() >= endHourly.Day() {
			// 	if startHourly.Year() >= endHourly.Year() && startHourly.Hour() >= endHourly.Hour() {
			// 		hf = false
			// 		break
			// 	}
			// }
			if startHourly.After(endHourly) {
				hf = false
				break
			}
			if startHourly.Hour() >= 0 && startHourly.Hour() < 9 {
				nt := startHourly.String()
				ntl := strings.Split(nt, " ")
				ntl[1] = "09:00:00"
				ntl = ntl[:2]
				nt = ntl[0] + "T" + ntl[1] + "+00:00"
				start2, _ := time.Parse("2006-01-02T15:04:05Z07:00", nt)
				startHourly = start2

			}
			updateHourly = startHourly.Add(time.Duration(durationHourly) * time.Minute)
			err, _ := db.Exec(out2, hourIndex, hourAuction, startHourly, updateHourly, 1, 1, ref, category, prod_name, desc, mrp, 1, tbase)

			if err != nil {
				fmt.Println(err)
			}
			hourIndex++
			if hourIndex%count == 0 {
				hourAuction++
				startHourly = startHourly.Add(time.Duration(intervalHourly) * time.Minute)
			}
		}

	}

	for twf {
		twelverows, twelveerr := db.Query("select ref,category,product,description,mrp,base from twelvehourtable order by random();")
		if twelveerr != nil {
			log.Fatal("error selecting from twelvetable", twelveerr)
		}
		for twelverows.Next() {
			fmt.Println("12 hour: ", twelveIndex)
			twerr := twelverows.Scan(&ref, &category, &prod_name, &desc, &mrp, &tbase)
			if twerr != nil {
				log.Fatal("error scanning twelve", twerr)
			}
			// if int(startTwh.Month()) >= int(endTwh.Month()) && startTwh.Day() >= endTwh.Day() {
			// 	if startTwh.Year() >= endTwh.Year() && startTwh.Hour() >= endTwh.Hour() {
			// 		twf = false
			// 		break
			// 	}
			// }
			if startTwh.After(endTwh) {
				twf = false
				break
			}

			updateTwh = startTwh.Add(time.Duration(durationTwh) * time.Minute)
			err, _ := db.Exec(out3, twelveIndex, twelveAuction, startTwh, updateTwh, 1, 1, ref, category, prod_name, desc, mrp, 1, tbase)

			if err != nil {
				fmt.Println(err)
			}
			twelveIndex++
			if twelveIndex%count == 0 {
				twelveAuction++
				startTwh = startTwh.Add(time.Duration(intervalTwh) * time.Minute)
			}
		}

	}

	for tff {
		twentyrows, twentyerr := db.Query("select ref,category,product,description,mrp,base from twentyfourhourtable order by random();")
		if twentyerr != nil {
			log.Fatal("error selecting from twentytable", twentyerr)
		}
		for twentyrows.Next() {
			fmt.Println("24 hour: ", twentyFourIndex)
			tferr := twentyrows.Scan(&ref, &category, &prod_name, &desc, &mrp, &tbase)
			if tferr != nil {
				log.Fatal("error scanning twenty", tferr)
			}
			// if int(startTfh.Month()) >= int(endTfh.Month()) && startTfh.Day() >= endTfh.Day() {
			// 	if startTfh.Year() >= endTfh.Year() && startTfh.Hour() >= endTfh.Hour() {
			// 		tff = false
			// 		break
			// 	}
			// }

			if startTfh.After(endTfh) {
				tff = false
				break
			}

			updateTfh = startTfh.Add(time.Duration(durationTfh) * time.Minute)
			err, _ := db.Exec(out4, twentyFourIndex, twentyFourAuction, startTfh, updateTfh, 1, 1, ref, category, prod_name, desc, mrp, 1, tbase)
			if err != nil {
				fmt.Println(err)
			}
			twentyFourIndex++
			if twentyFourIndex%count == 0 {
				twentyFourAuction++
				startTfh = startTfh.Add(time.Duration(intervalTfh) * time.Minute)
			}
		}

	}

}
func main() {
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatal(err1)
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	fmt.Println("Database student opened and ready.")
	schedule(db)
}
