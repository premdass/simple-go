package main

import (
	"math/big"
	"testing"
)

func TestFibonacciSumForZero(t *testing.T) {
	got := getEvenFibonacciSum(0)
	want := big.NewInt(0)
	assertEqual(got,want,t)
}


func TestFibonacciSumForOne(t *testing.T) {
	got := getEvenFibonacciSum(1)
	want := big.NewInt(2)
	assertEqual(got,want,t)
}

func TestFibonacciSumForTwenty(t *testing.T) {
	got := getEvenFibonacciSum(20)
	want := big.NewInt(2026369768940)
	assertEqual(got,want,t)
}


func TestFibonacciSumForHundred(t *testing.T) {
	got := getEvenFibonacciSum(100)
	want,_ := big.NewInt(0).SetString("290905784918002003245752779317049533129517076702883498623284700",0)
	assertEqual(got,want,t)
}


func assertEqual(got *big.Int, want *big.Int,t *testing.T){
	if (got.Cmp(want) != 0 ) {
		t.Errorf("got %q want %q", got, want)
	}
}
