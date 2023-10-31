package food

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

var (
	recID       autoIncrement
	recepies    map[uint]*Recepie
	recepiesMux sync.Mutex
	aplID       autoIncrement
	produceID   autoIncrement

	produce = map[string]*Produce{
		"maize":         {uint(produceID.id()), "maize", nil},
		"sviests":       {uint(produceID.id()), "sviests", nil},
		"desa":          {uint(produceID.id()), "desa", nil},
		"siers":         {uint(produceID.id()), "siers", nil},
		"gurķis":        {uint(produceID.id()), "gurķis", nil},
		"sīpoli":        {uint(produceID.id()), "sīpoli", nil},
		"burkāni":       {uint(produceID.id()), "burkāni", nil},
		"sēnes":         {uint(produceID.id()), "sēnes", nil},
		"cukini":        {uint(produceID.id()), "cukini", nil},
		"pētersīļi":     {uint(produceID.id()), "pētersīļi", nil},
		"rozmarīns":     {uint(produceID.id()), "rozmarīns", nil},
		"sarkanvīns":    {uint(produceID.id()), "sarkanvīns", nil},
		"griķi":         {uint(produceID.id()), "griķi", nil},
		"paprika":       {uint(produceID.id()), "paprika", nil},
		"kārtainā":      {uint(produceID.id()), "kārtainā", nil},
		"mīkla":         {uint(produceID.id()), "mīkla", nil},
		"krēmsiers":     {uint(produceID.id()), "krēmsiers", nil},
		"persiki":       {uint(produceID.id()), "persiki", nil},
		"biezpiens":     {uint(produceID.id()), "biezpiens", nil},
		"cukurs":        {uint(produceID.id()), "cukurs", nil},
		"olas":          {uint(produceID.id()), "olas", nil},
		"krējums":       {uint(produceID.id()), "krējums", nil},
		"milti":         {uint(produceID.id()), "milti", nil},
		"kartupeļi":     {uint(produceID.id()), "kartupeļi", nil},
		"fileja vistas": {uint(produceID.id()), "fileja vistas", nil},
		"spināti":       {uint(produceID.id()), "spināti", nil},
		"sinepes":       {uint(produceID.id()), "sinepes", nil},
		"medus":         {uint(produceID.id()), "medus", nil},
		"ananass":       {uint(produceID.id()), "ananass", nil},
		"rīsi":          {uint(produceID.id()), "rīsi", nil},
	}

	appliances = map[string]*Appliance{
		"nazis":         {uint(aplID.id()), "nazis", nil},
		"dēlis":         {uint(aplID.id()), "dēlis", nil},
		"rīve":          {uint(aplID.id()), "rīve", nil},
		"plīts":         {uint(aplID.id()), "plīts", nil},
		"panna":         {uint(aplID.id()), "panna", nil},
		"cepamlāpstiņa": {uint(aplID.id()), "cepamlāpstiņa", nil},
		"katls":         {uint(aplID.id()), "katls", nil},
		"cepeškrāsns":   {uint(aplID.id()), "cepeškrāsns", nil},
		"bļoda":         {uint(aplID.id()), "bļoda", nil},
		"mīklas rullis": {uint(aplID.id()), "mīklas rullis", nil},
		"karote":        {uint(aplID.id()), "karote", nil},
		"krūze":         {uint(aplID.id()), "krūze", nil},
		"dakša":         {uint(aplID.id()), "dakša", nil},
		"cepešpanna":    {uint(aplID.id()), "cepešpanna", nil},
		"gaļas āmurs":   {uint(aplID.id()), "gaļas āmurs", nil},
	}
)

func init() {
	recepies = make(map[uint]*Recepie)
	for _, r := range rec {
		recepies[r.ID] = r
	}
}

func GetRecepies() map[uint]*Recepie {
	return recepies
}

func GetRecepie(id uint) (*Recepie, error) {
	recepiesMux.Lock()
	defer recepiesMux.Unlock()

	_, ok := recepies[id]
	if !ok {
		return nil, errors.New("incorrect recepie id")
	}

	return recepies[id], nil
}

func randInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
