package food

import (
	mysql "System/db"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

var (
	recID    autoIncrement
	recepies map[uint]*Recepie

	produce = map[string]*Produce{
		"maize":         {Name: "maize"},
		"sviests":       {Name: "sviests"},
		"desa":          {Name: "desa"},
		"siers":         {Name: "siers"},
		"gurķis":        {Name: "gurķis"},
		"sīpoli":        {Name: "sīpoli"},
		"burkāni":       {Name: "burkāni"},
		"sēnes":         {Name: "sēnes"},
		"cukini":        {Name: "cukini"},
		"pētersīļi":     {Name: "pētersīļi"},
		"rozmarīns":     {Name: "rozmarīns"},
		"sarkanvīns":    {Name: "sarkanvīns"},
		"griķi":         {Name: "griķi"},
		"paprika":       {Name: "paprika"},
		"kārtainā":      {Name: "kārtainā"},
		"mīkla":         {Name: "mīkla"},
		"krēmsiers":     {Name: "krēmsiers"},
		"persiki":       {Name: "persiki"},
		"biezpiens":     {Name: "biezpiens"},
		"cukurs":        {Name: "cukurs"},
		"olas":          {Name: "olas"},
		"krējums":       {Name: "krējums"},
		"milti":         {Name: "milti"},
		"kartupeļi":     {Name: "kartupeļi"},
		"vistas fileja": {Name: "vistas fileja"},
		"spināti":       {Name: "spināti"},
		"sinepes":       {Name: "sinepes"},
		"medus":         {Name: "medus"},
		"ananass":       {Name: "ananass"},
		"rīsi":          {Name: "rīsi"},
	}

	appliances = map[string]*Appliance{
		"nazis":         {Name: "nazis"},
		"dēlis":         {Name: "dēlis"},
		"rīve":          {Name: "rīve"},
		"plīts":         {Name: "plīts"},
		"panna":         {Name: "panna"},
		"cepamlāpstiņa": {Name: "cepamlāpstiņa"},
		"katls":         {Name: "katls"},
		"cepeškrāsns":   {Name: "cepeškrāsns"},
		"bļoda":         {Name: "bļoda"},
		"mīklas rullis": {Name: "mīklas rullis"},
		"karote":        {Name: "karote"},
		"krūze":         {Name: "krūze"},
		"dakša":         {Name: "dakša"},
		"cepešpanna":    {Name: "cepešpanna"},
		"gaļas āmurs":   {Name: "gaļas āmurs"},
	}
)

func InitDB(user, pass, host, port, db string) error {
	conn, err := mysql.Connect(user, pass, host, port, db)
	if err != nil {
		return err
	}
	defer conn.Close()

	for _, p := range produce {
		_, err := conn.Exec(fmt.Sprintf(mysql.CreateProduceQuery, p.Name))
		if err != nil {
			return err
		}
	}

	for _, apl := range appliances {
		_, err := conn.Exec(fmt.Sprintf(mysql.CreateApplienceQuery, apl.Name))
		if err != nil {
			return err
		}
	}

	var prod []Produce
	var apl []Appliance

	for _, p := range produce {
		var tmp Produce
		err := conn.QueryRow(fmt.Sprintf(mysql.GetProduceByNameQuery, p.Name)).Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			panic(err)
		}

		prod = append(prod, tmp)
	}

	for _, a := range appliances {
		var tmp Appliance
		err := conn.QueryRow(fmt.Sprintf(mysql.GetApplienceByNameQuery, a.Name)).Scan(&tmp.ID, &tmp.Name)
		if err != nil {
			panic(err)
		}

		apl = append(apl, tmp)
	}

	for _, r := range rec {
		_, err = conn.Exec(fmt.Sprintf(mysql.InsertRecepieQuery, r.Name, r.Desciption, ""))
		if err != nil {
			panic(err)
		}

		t := struct {
			ID uint
		}{}

		err := conn.QueryRow(fmt.Sprintf(mysql.GetRecipesByNameQuery, r.Name)).Scan(&t.ID)
		if err != nil {
			panic(err)
		}

	produce_loop:
		for _, p := range r.Produces {
			if p == nil {
				continue produce_loop
			}

			for _, tmp := range prod {
				if p.Name == tmp.Name {
					_, err = conn.Exec(fmt.Sprintf(mysql.InsertRecepieQuery2, fmt.Sprintf("(%d, %d)", t.ID, tmp.ID)))
					if err != nil {
						panic(err)
					}

					continue produce_loop
				}
			}
		}

	apl_loop:
		for _, a := range r.Appliances {
			if a == nil {
				continue apl_loop
			}

			for _, tmp := range apl {
				if a.Name == tmp.Name {
					_, err = conn.Exec(fmt.Sprintf(mysql.InsertRecepieQuery3, fmt.Sprintf("(%d, %d)", t.ID, tmp.ID)))
					if err != nil {
						panic(err)
					}

					continue apl_loop
				}
			}
		}
	}

	return nil
}

func GetRecepies(conn *sql.DB) map[uint]*Recepie {
	rows, err := conn.Query(fmt.Sprintf(mysql.GetRecipes))
	if err != nil {
		panic(err)
	}

	res, err := todata(rows)
	if err != nil {
		panic(err)
	}

	out := make(map[uint]*Recepie)

	for _, r := range res {
		var tmpApliences []*Appliance
		var tmpProduce []*Produce

		i, err := strconv.Atoi(r["ID"])
		if err != nil {
			panic(err)
		}

		aplRows, err := conn.Query(fmt.Sprintf(mysql.GetRecipesApl))
		if err != nil {
			panic(err)
		}

		aplData, err := todata(aplRows)
		if err != nil {
			panic(err)
		}

		for _, apl := range aplData {
			if apl["recipe_id"] == r["ID"] {
				ai, err := strconv.Atoi(apl["appliance_id"])
				if err != nil {
					panic(err)
				}

				var a Appliance
				err = conn.QueryRow(fmt.Sprintf(mysql.GetApplienceByIDQuery, ai)).Scan(&a.ID, &a.Name)
				if err != nil {
					panic(err)
				}
				tmpApliences = append(tmpApliences, &a)
			}
		}

		prodRows, err := conn.Query(fmt.Sprintf(mysql.GetRecipesProd))
		if err != nil {
			panic(err)
		}

		prodData, err := todata(prodRows)
		if err != nil {
			panic(err)
		}

		for _, p := range prodData {
			if p["recipe_id"] == r["ID"] {
				pr, err := strconv.Atoi(p["produce_id"])
				if err != nil {
					panic(err)
				}

				var p Produce
				err = conn.QueryRow(fmt.Sprintf(mysql.GetProduceByIDQuery, pr)).Scan(&p.ID, &p.Name)
				if err != nil {
					panic(err)
				}
				tmpProduce = append(tmpProduce, &p)
			}
		}

		out[uint(i)] = &Recepie{
			ID:         uint(i),
			Name:       r["Name"],
			Desciption: r["Description"],
			Users:      nil,
			Appliances: tmpApliences,
			Produces:   tmpProduce,
		}
	}

	return out
}

func GetRecepie(conn *sql.DB, id uint) (*Recepie, error) {
	rows, err := conn.Query(fmt.Sprintf(mysql.GetRecipes))
	if err != nil {
		return nil, err
	}

	res, err := todata(rows)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		var tmpApliences []*Appliance
		var tmpProduce []*Produce

		i, err := strconv.Atoi(r["ID"])
		if err != nil {
			return nil, err
		}

		if uint(i) != id {
			continue
		}

		aplRows, err := conn.Query(fmt.Sprintf(mysql.GetRecipesApl))
		if err != nil {
			return nil, err
		}

		aplData, err := todata(aplRows)
		if err != nil {
			return nil, err
		}

		for _, apl := range aplData {
			if apl["recipe_id"] == r["ID"] {
				ai, err := strconv.Atoi(apl["appliance_id"])
				if err != nil {
					return nil, err
				}

				var a Appliance
				err = conn.QueryRow(fmt.Sprintf(mysql.GetApplienceByIDQuery, ai)).Scan(&a.ID, &a.Name)
				if err != nil {
					return nil, err
				}
				tmpApliences = append(tmpApliences, &a)
			}
		}

		prodRows, err := conn.Query(fmt.Sprintf(mysql.GetRecipesProd))
		if err != nil {
			return nil, err
		}

		prodData, err := todata(prodRows)
		if err != nil {
			return nil, err
		}

		for _, p := range prodData {
			if p["recipe_id"] == r["ID"] {
				pr, err := strconv.Atoi(p["produce_id"])
				if err != nil {
					return nil, err
				}

				var p Produce
				err = conn.QueryRow(fmt.Sprintf(mysql.GetProduceByIDQuery, pr)).Scan(&p.ID, &p.Name)
				if err != nil {
					return nil, err
				}
				tmpProduce = append(tmpProduce, &p)
			}
		}

		return &Recepie{
			ID:         uint(i),
			Name:       r["Name"],
			Desciption: r["Description"],
			Users:      nil,
			Appliances: tmpApliences,
			Produces:   tmpProduce,
		}, nil
	}

	return nil, errors.New("not found")
}

func todata(rows *sql.Rows) (result []map[string]string, err error) {
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		entry := make(map[string]string)

		for i, col := range values {
			if col == nil {
				entry[columns[i]] = ""
			} else {
				entry[columns[i]] = string(col)
			}
		}

		result = append(result, entry)
	}

	return result, nil
}
