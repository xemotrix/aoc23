package day05_test

import (
	"aoc23/day05"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Day05TestSuite struct {
	suite.Suite
}

func TestDay05TestSuite(t *testing.T) {
	suite.Run(t, new(Day05TestSuite))
}

func (suite *Day05TestSuite) SetupTest() {
}

func (suite *Day05TestSuite) TestMergeInnerFullOverlap() {
	r1 := day05.Range{From: 10, To: 20, Offset: 11}
	r2 := day05.Range{From: 5, To: 25, Offset: 22}
	res, overlap := r1.Merge(r2)

	expected := []day05.Range{{From: 10, To: 20, Offset: 33}}
	suite.Equal(expected, res)
	suite.True(overlap)
}

func (suite *Day05TestSuite) TestMergeOuterFullOverlap() {
	r1 := day05.Range{From: 5, To: 25, Offset: 11}
	r2 := day05.Range{From: 10, To: 20, Offset: 22}
	res, overlap := r1.Merge(r2)

	expected := []day05.Range{
		{From: 5, To: 9, Offset: 11},
		{From: 10, To: 20, Offset: 33},
		{From: 21, To: 25, Offset: 11},
	}
	suite.Equal(expected, res)
	suite.True(overlap)
}

func (suite *Day05TestSuite) TestMergeLeftOverlap() {
	r1 := day05.Range{From: 10, To: 20, Offset: 11}
	r2 := day05.Range{From: 5, To: 15, Offset: 22}
	res, overlap := r1.Merge(r2)

	expected := []day05.Range{
		{From: 10, To: 15, Offset: 33},
		{From: 16, To: 20, Offset: 11},
	}
	suite.Equal(expected, res)
	suite.True(overlap)
}

func (suite *Day05TestSuite) TestMergeRightOverlap() {
	r1 := day05.Range{From: 10, To: 20, Offset: 11}
	r2 := day05.Range{From: 15, To: 25, Offset: 22}
	res, overlap := r1.Merge(r2)

	expected := []day05.Range{
		{From: 10, To: 14, Offset: 11},
		{From: 15, To: 20, Offset: 33},
	}

	suite.Equal(expected, res)
	suite.True(overlap)
}

func (suite *Day05TestSuite) TestNoOverlap() {
	r1 := day05.Range{From: 10, To: 20, Offset: 11}
	r2 := day05.Range{From: 30, To: 40, Offset: 22}
	res, overlap := r1.Merge(r2)

	suite.Nil(res)
	suite.False(overlap)
}

func (suite *Day05TestSuite) TestMatchSeedWithRanges() {
	// func MatchSeedWithRanges(seed Range, ranges []Mapping) []Range {
	seed := day05.Range{From: 10, To: 20, Offset: 0}
	ranges := []day05.Mapping{
		{DestStart: 105, SourceStart: 5, Length: 10},
		{DestStart: 118, SourceStart: 18, Length: 10},
	}
	res := day05.MatchSeedWithRanges(seed, ranges)
	expected := []day05.Range{
		{From: 10, To: 14, Offset: 100},
		{From: 15, To: 17, Offset: 0},
		{From: 18, To: 20, Offset: 100},
	}
	suite.Equal(expected, res)
}

func (suite *Day05TestSuite) TestRunStep() {
	// func RunStep(seeds []Range, mappings []Mapping) []Range {

	seeds := []day05.Range{
		{From: 10, To: 20, Offset: 0},
		{From: 30, To: 40, Offset: 0},
	}
	ranges := []day05.Mapping{
		{DestStart: 105, SourceStart: 5, Length: 10},
		{DestStart: 118, SourceStart: 18, Length: 20},
	}
	res := day05.RunStep(seeds, ranges)
	expected := []day05.Range{
		{From: 10, To: 14, Offset: 100},
		{From: 15, To: 17, Offset: 0},
		{From: 18, To: 20, Offset: 100},
		{From: 30, To: 37, Offset: 100},
		{From: 38, To: 40, Offset: 0},
	}
	suite.Equal(expected, res)
}
