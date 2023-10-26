package server

// import (
// 	"System/food"
// 	"reflect"
// 	"testing"
// )

// func Test_filterRecipes(t *testing.T) {

// 	type args struct {
// 		pIDs   []uint
// 		aIDs   []uint
// 		pMatch bool
// 		aMatch bool
// 	}
// 	tests := []struct {
// 		name     string
// 		args     args
// 		recepies []food.Recepie
// 		want     map[uint]*food.Recepie
// 	}{
// 		{
// 			"non_strict",
// 			args{},
// 			[]food.Recepie{
// 				{
// 					ID: 1,
// 					Produces: map[uint]food.Produce{
// 					1:food.Produce{ID: 1},
// 					2:food.Produce{ID: 2},
// 					3:food.Produce{ID: 3},
// 					},
// 					Appliances: make([]*food.Appliance, ),

// 			},

// 			make(map[uint]*food.Recepie),
// 		},
// 	}
// 	for _, tt := range tests {

// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := filterRecipes(tt.args.pIDs, tt.args.aIDs, tt.args.pMatch, tt.args.aMatch); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("filterRecipes() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
