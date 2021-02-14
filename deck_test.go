package main

import (
	"reflect"
	"testing"
)

func Test_newDeck_shouldReturnDeckWithLength12(t *testing.T) {
	t.Run("deck length should be 12", func(t *testing.T) {
		if got := newDeck(); !reflect.DeepEqual(len(got), 12) {
			t.Errorf("newDeck() = %v, want %v", got, 12)
		}
	})
}

func Test_deal(t *testing.T) {
	type args struct {
		d             deck
		numberOfCards int
	}
	tests := []struct {
		name  string
		args  args
		want  deck
		want1 deck
	}{
		{name: "Should return hand with specified number and rest of the deck minus hand", args: struct {
			d             deck
			numberOfCards int
		}{d: deck{"1", "2"}, numberOfCards: 1}, want: deck{"1"}, want1: deck{"2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := deal(tt.args.d, tt.args.numberOfCards)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("deal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
