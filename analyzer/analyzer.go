package analyzer

import (
	"github.com/abondar24/RecipeCountAnalyzer/data"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Analyzer struct {
	file         []byte
	keyWord      string
	postcode     string
	deliveryTime string
}

const (
	RecipeField   = "recipe"
	PostcodeField = "postcode"
	DeliveryField = "delivery"
)

func NewAnalyzer(fileName *string, keyword *string, postcode *string, deliveryTime *string) *Analyzer {
	file, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	return &Analyzer{file, *keyword, *postcode, *deliveryTime}
}

func (an *Analyzer) Analyze() *data.CountResult {
	res := &data.CountResult{}

	recipes := an.getUniqueRecipes()

	res.UniqueCount = int32(len(*recipes))
	res.RecipeCount = *an.wrapRecipes(recipes)
	res.BusyPostcode = *an.getBusyAddress()
	res.PostcodePerTime = *an.getCountPerPostCode()
	res.Matches = *an.getRecipesByWord(recipes)

	return res
}

func (an *Analyzer) getUniqueRecipes() *map[string]int32 {
	recipes := map[string]int32{}
	_, err := jsonparser.ArrayEach(an.file, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		val, _, _, err := jsonparser.Get(value, RecipeField)
		recipes[string(val)] = recipes[string(val)] + 1
	})

	if err != nil {
		log.Fatal(err)
	}

	return &recipes
}

func (an *Analyzer) wrapRecipes(recipes *map[string]int32) *[]data.Recipe {
	var wrapped []data.Recipe

	for rec, count := range *recipes {
		wr := data.Recipe{RecipeName: rec, Count: count}
		wrapped = append(wrapped, wr)
	}

	return &wrapped
}

func (an *Analyzer) getBusyAddress() *data.PostcodeBusy {
	addrChan := make(chan *data.PostcodeBusy)

	go func() {
		postcodes := map[string]int32{}
		res := &data.PostcodeBusy{Code: "", Count: 0}

		_, err := jsonparser.ArrayEach(an.file, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			val, _, _, err := jsonparser.Get(value, PostcodeField)
			postcodes[string(val)] = postcodes[string(val)] + 1
		})

		if err != nil {
			log.Fatal(err)
		}

		for pcode, count := range postcodes {
			if count > res.Count {
				res.Count = count
				res.Code = pcode
			}
		}

		addrChan <- res
	}()

	return <-addrChan
}

func (an *Analyzer) getCountPerPostCode() *data.PostcodeTime {
	countCh := make(chan *data.PostcodeTime)
	go func() {
		var deliveryCount int32

		if len(an.deliveryTime) == 0 {
			countCh <- &data.PostcodeTime{}
			return
		}

		startTime, endTime := parseTime(&an.deliveryTime)

		res := &data.PostcodeTime{Code: an.postcode, From: *startTime, To: *endTime, Count: 0}

		_, err := jsonparser.ArrayEach(an.file, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			val, _, _, err := jsonparser.Get(value, PostcodeField)
			if string(val) == an.postcode {
				val, _, _, err = jsonparser.Get(value, DeliveryField)
				strVal := string(val)
				timeSlot := getTimeSlot(&strVal)
				if timeMatches(startTime, endTime, timeSlot) {
					deliveryCount++
				}
			}
		})

		if err != nil {
			log.Fatal(err)
		}

		res.Count = deliveryCount

		countCh <- res
	}()

	return <-countCh
}

func (an *Analyzer) getRecipesByWord(recipes *map[string]int32) *[]string {
	keyWord := strings.ToLower(an.keyWord)

	if len(keyWord) == 0 {
		return &[]string{}
	}

	var res []string

	for r := range *recipes {
		lr := strings.ToLower(r)
		if strings.Contains(lr, keyWord) {
			res = append(res, r)
		}
	}
	sort.Strings(res)
	return &res
}

func getTimeSlot(val *string) *string {
	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	timeSlot := ""
	for _, day := range days {
		if strings.Contains(*val, day) {
			timeSlot = strings.Split(*val, day+" ")[1]
		}
	}

	return &timeSlot
}

func timeMatches(startTime *string, endTime *string, checkSlot *string) bool {
	checkStartTime, checkEndTme := parseTime(checkSlot)

	amSuffix := "AM"
	pmSuffix := "PM"

	start := getTimeInt(startTime, &amSuffix)
	checkStart := getTimeInt(checkStartTime, &amSuffix)

	if start > checkStart {
		return false
	}

	end := getTimeInt(endTime, &pmSuffix)
	checkEnd := getTimeInt(checkEndTme, &pmSuffix)

	if end < checkEnd {
		return false
	}

	return true
}

func getTimeInt(time *string, suffix *string) int {
	timePure := strings.Split(*time, *suffix)
	intTime, err := strconv.Atoi(timePure[0])
	if err != nil {
		log.Fatal(err)
	}

	return intTime
}

func parseTime(deliveryTime *string) (*string, *string) {
	timeSplit := strings.Split(*deliveryTime, "-")

	startTime := strings.TrimSuffix(timeSplit[0], " ")
	endTime := strings.TrimPrefix(timeSplit[1], " ")

	return &startTime, &endTime
}
