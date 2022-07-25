package buypokemonhandler

import (
	"context"
	"reflect"
	"testing"

	"pokemonstore/api/sdk"

	"github.com/golang/protobuf/ptypes/empty"
)

func TestBuyPokemonHandler_BuyPokemon(t *testing.T) {
	type fields struct {
		usecases buyPokemonUseCase
	}
	type args struct {
		ctx context.Context
		in  *sdk.BuyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *empty.Empty
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Handler{
				buyPokemonUseCase: tt.fields.usecases,
			}
			got, err := s.BuyPokemon(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuyPokemonHandler.BuyPokemon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuyPokemonHandler.BuyPokemon() = %v, want %v", got, tt.want)
			}
		})
	}
}
