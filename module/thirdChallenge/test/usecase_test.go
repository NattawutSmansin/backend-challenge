package test

import (
	"backend-challenge/module/thirdChallenge/repository"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockRepository จำลองการเรียก repository สำหรับการทดสอบ
type MockRepository struct{}

func (m *MockRepository) FetchBeefData(ctx context.Context) (string, error) {
	return `Drumstick ut tenderloin meatball ham fugiat consectetur minim pancetta meatloaf.  Exercitation qui magna mollit, landjaeger dolore hamburger t-bone sirloin do cillum.  Chicken venison tongue, ham hock tri-tip voluptate beef ut quis.  Picanha shoulder aute short loin pork loin esse.
Bresaola drumstick biltong chislic dolore tenderloin prosciutto lorem chuck in quis pastrami velit.  Eiusmod culpa biltong, porchetta exercitation ex labore.  Nostrud salami laborum adipisicing voluptate bresaola.  Sausage chislic salami meatball chuck ut veniam consequat in culpa do nisi aute sed.  Elit in sausage brisket sint.  Boudin in qui, swine chicken excepteur lorem ut pariatur brisket tempor pastrami.  Pastrami enim duis cupim officia tempor ullamco prosciutto jowl.
Prosciutto mollit est, brisket dolor burgdoggen shankle fatback nostrud duis chuck beef ribs.  Reprehenderit consequat aliqua, dolor bresaola sint picanha deserunt.  Ipsum velit spare ribs, in kevin venison chicken tenderloin pancetta in aliquip anim culpa.  Tenderloin laboris elit pork loin.
Buffalo est chicken, incididunt meatball pork chop ad.  Ullamco quis jowl, pig pork loin buffalo dolore tempor ham hock occaecat turkey frankfurter eiusmod do incididunt.  Consequat capicola meatloaf ut drumstick esse.  Jerky est chislic picanha beef ribs venison spare ribs aliqua bresaola sunt cow ut tongue consequat.  Tail t-bone id, pork venison bresaola ex biltong.  Dolor ground round beef ribs pariatur, brisket quis shoulder do consectetur aute sed eiusmod proident nulla.  Picanha do deserunt incididunt qui dolore ex occaecat capicola.
Anim commodo laboris corned beef ball tip chislic quis pork belly beef ribs hamburger meatball.  Bacon deserunt quis pork, turducken hamburger nostrud.  Prosciutto tail pork chop rump labore.  In rump bresaola, culpa shank minim laboris enim meatball magna boudin landjaeger spare ribs nostrud.  Bresaola filet mignon drumstick, sausage leberkas in tenderloin deserunt fugiat dolore ea.  Turducken in elit, cow ribeye pig est non.  Qui nostrud nulla pariatur.
Tail kevin dolore ham hock, ribeye chislic in shoulder beef ribs sirloin ullamco flank duis fatback.  Officia shank meatball pork loin ball tip strip steak jowl culpa dolore deserunt venison capicola.  Esse fugiat sed veniam, cupidatat ham eu pork ex irure picanha nostrud prosciutto et rump.  Ea magna chislic elit commodo porchetta jowl.
Bacon enim deserunt beef ribs laborum, corned beef chislic tri-tip proident eu.  Cillum tail jowl fatback bacon.  Velit duis veniam ipsum sint strip steak in capicola chicken pork labore porchetta cupidatat pig voluptate.  Irure tempor filet mignon tenderloin, alcatra sunt do bacon nulla incididunt boudin consequat.  Nulla doner biltong porchetta labore landjaeger, commodo esse.  Ex qui pancetta tongue in meatloaf strip steak officia sirloin beef ribs aliquip venison anim non ham hock.
Pancetta strip steak qui reprehenderit laborum t-bone.`, nil
}

func TestFetchBeefDataEqual(t *testing.T) { // ตรวจสอบค่าว่าเท่ากันหรือไม่
	repo := repository.NewRepository()
	fetchBeefData, err := repo.FetchBeefData()

	assert.NoError(t, err)
	assert.NotEmpty(t, fetchBeefData)

	// ประเภทเนื้อที่ต้องตรวจสอบ
	meatType := []string{
		"t-bone",
		"fatback",
		"pastrami",
		"pork",
		"meatloaf",
		"jowl",
		"enim",
		"bresaola",
	}

	fetchBeefData = strings.ToLower(fetchBeefData)

	dataMeat := make(map[string]int32)
	for _, meat := range meatType {
		dataMeat[meat] = int32(strings.Count(fetchBeefData, meat))
	}

	expected := map[string]int32{
		"t-bone":   4,
		"fatback":  1,
		"pastrami": 1,
		"pork":     1,
		"meatloaf": 1,
		"jowl":     1,
		"enim":     1,
		"bresaola": 1,
	}

	assert.Equal(t, expected, dataMeat, "Meat count mismatch")
}

func TestFetchBeefDataContains(t *testing.T) { // ตรวจสออบแต่ละ key มีเท่่ากัน และเหมือนกันหรือไม่
	repo := repository.NewRepository()
	fetchBeefData, err := repo.FetchBeefData()

	assert.NoError(t, err)
	assert.NotEmpty(t, fetchBeefData)

	// ประเภทเนื้อที่ต้องตรวจสอบ
	meatType := []string{
		"t-bone",
		"fatback",
		"pastrami",
		"pork",
		"meatloaf",
		"jowl",
		"enim",
		"bresaola",
	}

	fetchBeefData = strings.ToLower(fetchBeefData)

	dataMeat := make(map[string]int32)
	for _, meat := range meatType {
		dataMeat[meat] = int32(strings.Count(fetchBeefData, meat))
	}

	// ประเภทเนื้อที่ต้องตรวจสอบ
	expected := map[string]int32{
		"t-bone":   4,
		"fatback":  1,
		"pastrami": 1,
		"pork":     1,
		"meatloaf": 1,
		"jowl":     1,
		"enim":     1,
		"bresaola": 1,
		"ribs":     1,
	}

	// เปรียบเทียบเฉพาะ meat types
	for key := range expected {
		assert.Contains(t, dataMeat, key, "Missing key: %s", key)
	}
}



func TestFetchBeefMockDataEqual(t *testing.T) { // ตรวจสอบค่าว่าเท่ากันหรือไม่
	mockRepo := &MockRepository{}
	ctx := context.Background()
	fetchBeefData, err := mockRepo.FetchBeefData(ctx)

	assert.NoError(t, err)
	assert.NotEmpty(t, fetchBeefData)

	// ประเภทเนื้อที่ต้องตรวจสอบ
	meatType := []string{
		"t-bone",
		"fatback",
		"pastrami",
		"pork",
		"meatloaf",
		"jowl",
		"enim",
		"bresaola",
	}

	fetchBeefData = strings.ToLower(fetchBeefData)

	dataMeat := make(map[string]int32)
	for _, meat := range meatType {
		dataMeat[meat] = int32(strings.Count(fetchBeefData, meat))
	}

	expected := map[string]int32{
		"t-bone":   3,
		"fatback":  3,
		"pastrami": 3,
		"pork":     11,
		"meatloaf": 3,
		"jowl":     5,
		"enim":     3,
		"bresaola": 7,
	}

	
	assert.Equal(t, expected, dataMeat, "Meat count mismatch")
}

func TestFetchBeefMockDataContains(t *testing.T) { // ตรวจสออบแต่ละ key มีเท่่ากัน และเหมือนกันหรือไม่
	mockRepo := &MockRepository{}
	ctx := context.Background()
	fetchBeefData, err := mockRepo.FetchBeefData(ctx)

	assert.NoError(t, err)
	assert.NotEmpty(t, fetchBeefData)

	// ประเภทเนื้อที่ต้องตรวจสอบ
	meatType := []string{
		"t-bone",
		"fatback",
		"pastrami",
		"pork",
		"meatloaf",
		"jowl",
		"enim",
		"bresaola",
	}

	fetchBeefData = strings.ToLower(fetchBeefData)

	dataMeat := make(map[string]int32)
	for _, meat := range meatType {
		dataMeat[meat] = int32(strings.Count(fetchBeefData, meat))
	}

	// ประเภทเนื้อที่ต้องตรวจสอบ
	expected := map[string]int32{
		"t-bone":   3,
		"fatback":  3,
		"pastrami": 3,
		"pork":     11,
		"meatloaf": 3,
		"jowl":     5,
		"enim":     3,
		"bresaola": 7,
	}

	// เปรียบเทียบเฉพาะ meat types
	for key := range expected {
		assert.Contains(t, dataMeat, key, "Missing key: %s", key)
	}
}