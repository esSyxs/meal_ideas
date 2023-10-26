package server

import (
	"System/food"
	"reflect"
	"testing"
)

func Test_filterRecipes(t *testing.T) {

	type args struct {
		recepies map[uint]*food.Recepie
		pIDs     []uint
		aIDs     []uint
		pMatch   bool
		aMatch   bool
	}
	tests := []struct {
		name string
		args args
		want map[uint]*food.Recepie
	}{
		{
			"non_strict",
			args{
				map[uint]*food.Recepie{
					1: {
						ID: 1,
						Produces: []*food.Produce{
							{ID: 1},
							{ID: 2},
							{ID: 3},
						},
						Appliances: []*food.Appliance{
							{ID: 1},
							{ID: 2},
							{ID: 3},
						},
					},
				},
				[]uint{1},
				[]uint{2},
				false,
				false,
			},
			map[uint]*food.Recepie{
				1: {ID: 1,
					Produces: []*food.Produce{
						{ID: 1},
						{ID: 2},
						{ID: 3},
					},
					Appliances: []*food.Appliance{
						{ID: 1},
						{ID: 2},
						{ID: 3},
					},
				},
			},
		},
		{
			"non_strict_multiple",
			args{
				map[uint]*food.Recepie{
					1: {
						ID: 1,
						Produces: []*food.Produce{
							{ID: 1},
							{ID: 2},
							{ID: 3},
						},
						Appliances: []*food.Appliance{
							{ID: 1},
							{ID: 2},
							{ID: 3},
						},
					},
					2: {
						ID: 2,
						Produces: []*food.Produce{
							{ID: 4},
							{ID: 5},
							{ID: 6},
						},
						Appliances: []*food.Appliance{
							{ID: 4},
							{ID: 5},
							{ID: 6},
						},
					},
				},
				[]uint{1},
				[]uint{2},
				false,
				false,
			},
			map[uint]*food.Recepie{
				1: {ID: 1,
					Produces: []*food.Produce{
						{ID: 1},
						{ID: 2},
						{ID: 3},
					},
					Appliances: []*food.Appliance{
						{ID: 1},
						{ID: 2},
						{ID: 3},
					},
				},
			},
		},
		{
			"strict_produce",
			args{
				map[uint]*food.Recepie{
					1: {
						ID: 1,
						Produces: []*food.Produce{
							{ID: 1},
							{ID: 2},
						},
						Appliances: []*food.Appliance{
							{ID: 1},
							{ID: 2},
							{ID: 3},
						},
					},
					2: {
						ID: 2,
						Produces: []*food.Produce{
							{ID: 1},
							{ID: 2},
							{ID: 3},
						},
						Appliances: []*food.Appliance{
							{ID: 4},
							{ID: 5},
							{ID: 6},
						},
					},
				},
				[]uint{1, 2},
				[]uint{2},
				true,
				false,
			},
			map[uint]*food.Recepie{
				1: {ID: 1,
					Produces: []*food.Produce{
						{ID: 1},
						{ID: 2},
					},
					Appliances: []*food.Appliance{
						{ID: 1},
						{ID: 2},
						{ID: 3},
					},
				},
			},
		},
		{
			"strict_appliance",
			args{
				map[uint]*food.Recepie{
					1: {
						ID: 1,
						Produces: []*food.Produce{
							{ID: 1},
							{ID: 2},
						},
						Appliances: []*food.Appliance{
							{ID: 1},
							{ID: 2},
						},
					},
					2: {
						ID: 2,
						Produces: []*food.Produce{
							{ID: 1},
							{ID: 2},
							{ID: 3},
						},
						Appliances: []*food.Appliance{
							{ID: 4},
							{ID: 5},
							{ID: 6},
						},
					},
				},
				[]uint{2},
				[]uint{1, 2},
				false,
				true,
			},
			map[uint]*food.Recepie{
				1: {ID: 1,
					Produces: []*food.Produce{
						{ID: 1},
						{ID: 2},
					},
					Appliances: []*food.Appliance{
						{ID: 1},
						{ID: 2},
					},
				},
			},
		},
		{
			"strict_both",
			args{
				map[uint]*food.Recepie{
					1: {
						ID: 1,
						Produces: []*food.Produce{
							{ID: 1},
							{ID: 2},
						},
						Appliances: []*food.Appliance{
							{ID: 1},
							{ID: 2},
						},
					},
					2: {
						ID: 2,
						Produces: []*food.Produce{
							{ID: 1},
							{ID: 2},
							{ID: 3},
						},
						Appliances: []*food.Appliance{
							{ID: 4},
							{ID: 5},
							{ID: 6},
						},
					},
				},
				[]uint{1, 2},
				[]uint{1, 2},
				true,
				true,
			},
			map[uint]*food.Recepie{
				1: {ID: 1,
					Produces: []*food.Produce{
						{ID: 1},
						{ID: 2},
					},
					Appliances: []*food.Appliance{
						{ID: 1},
						{ID: 2},
					},
				},
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := filterRecipes(tt.args.recepies, tt.args.pIDs, tt.args.aIDs, tt.args.pMatch, tt.args.aMatch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterRecipes() = %v, want %v", got, tt.want)
			}
		})
	}
}
