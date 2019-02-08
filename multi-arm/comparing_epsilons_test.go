package main

import "testing"

func TestUpdate(t *testing.T) {
	b := Bandit{n: 1, mean: 1}
	b.Update(1.0)
	got := b.mean
	want := 1.0
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestFindMax(t *testing.T) {
	bandits := []*Bandit{&Bandit{mean: 1.0}, &Bandit{mean: 2.0}, &Bandit{mean: 3.0}}
	want := &Bandit{mean: 3.0}
	got := findMax(bandits)
	if got != want {
		t.Errorf("want %f, got %f", want, got)
	}
}
