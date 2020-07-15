package analyzer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyzer_Analyze(t *testing.T) {
	fileName := "test.json"
	keyword := "chicken"
	postcode := "10186"
	deliveryTime := "8AM - 9PM"

	anl := NewAnalyzer(&fileName, &keyword, &postcode, &deliveryTime)
	res := anl.Analyze()

	assert.Equal(t, int32(6), res.UniqueCount)
	assert.Equal(t, "10224", res.BusyPostcode.Code)
	assert.Equal(t, int32(3), res.BusyPostcode.Count)
	assert.Equal(t, 6, len(res.RecipeCount))
	assert.Equal(t, postcode, res.PostcodePerTime.Code)
	assert.Equal(t, int32(1), res.PostcodePerTime.Count)
	assert.Equal(t, "8AM", res.PostcodePerTime.From)
	assert.Equal(t, "9PM", res.PostcodePerTime.To)
	assert.Equal(t, 1, len(res.Matches))
	assert.Equal(t, "Creamy Dill Chicken", res.Matches[0])

}

func TestAnalyzer_Analyze_deliveryCountZeroAM(t *testing.T) {
	fileName := "test.json"
	keyword := "chicken"
	postcode := "10120"
	deliveryTime := "8AM - 10PM"

	anl := NewAnalyzer(&fileName, &keyword, &postcode, &deliveryTime)
	res := anl.Analyze()

	assert.Equal(t, int32(0), res.PostcodePerTime.Count)

}

func TestAnalyzer_Analyze_deliveryCountZeroPM(t *testing.T) {
	fileName := "test.json"
	keyword := "chicken"
	postcode := "10120"
	deliveryTime := "7AM - 3PM"

	anl := NewAnalyzer(&fileName, &keyword, &postcode, &deliveryTime)
	res := anl.Analyze()

	assert.Equal(t, int32(0), res.PostcodePerTime.Count)

}

func TestAnalyzer_Analyze_noMatches(t *testing.T) {
	fileName := "test.json"
	keyword := "wurst"
	postcode := "10186"
	deliveryTime := "8AM - 9PM"

	anl := NewAnalyzer(&fileName, &keyword, &postcode, &deliveryTime)
	res := anl.Analyze()

	assert.Equal(t, 0, len(res.Matches))

}

func TestAnalyzer_Analyze_postcodenotFound(t *testing.T) {
	fileName := "test.json"
	keyword := "chicken"
	postcode := "20190"
	deliveryTime := "8AM - 9PM"

	anl := NewAnalyzer(&fileName, &keyword, &postcode, &deliveryTime)
	res := anl.Analyze()

	assert.Equal(t, postcode, res.PostcodePerTime.Code)
	assert.Equal(t, int32(0), res.PostcodePerTime.Count)
	assert.Equal(t, "8AM", res.PostcodePerTime.From)
	assert.Equal(t, "9PM", res.PostcodePerTime.To)
}

func TestAnalyzer_Analyze_matches_alphabetical(t *testing.T) {
	fileName := "test.json"
	keyword := "chops"
	postcode := "10186"
	deliveryTime := "8AM - 9PM"

	anl := NewAnalyzer(&fileName, &keyword, &postcode, &deliveryTime)
	res := anl.Analyze()

	assert.Equal(t, 3, len(res.Matches))
	assert.Equal(t, "Balsamic Pork Chops", res.Matches[0])
	assert.Equal(t, "Cherry Balsamic Pork Chops", res.Matches[1])
	assert.Equal(t, "Huge Pork Chops", res.Matches[2])

}

func TestAnalyzer_Analyze_empty_flags(t *testing.T) {
	fileName := "test.json"
	keyword := ""
	postcode := ""
	deliveryTime := ""

	anl := NewAnalyzer(&fileName, &keyword, &postcode, &deliveryTime)
	res := anl.Analyze()

	assert.Equal(t, int32(6), res.UniqueCount)
	assert.Equal(t, "10224", res.BusyPostcode.Code)
	assert.Equal(t, int32(3), res.BusyPostcode.Count)
	assert.Equal(t, 6, len(res.RecipeCount))
	assert.Equal(t, "", res.PostcodePerTime.Code)
	assert.Equal(t, int32(0), res.PostcodePerTime.Count)
	assert.Equal(t, "", res.PostcodePerTime.From)
	assert.Equal(t, "", res.PostcodePerTime.To)
	assert.Equal(t, 0, len(res.Matches))

}
