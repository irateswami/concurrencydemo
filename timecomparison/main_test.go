package main

import (
	"testing"
)

func BenchmarkNonCon(b *testing.B) {
	NonCon()
}

func BenchmarkCon(b *testing.B) {
	Con()
}

func BenchmarkConGooglePkg(b *testing.B) {
	ConGooglePkg()
}
