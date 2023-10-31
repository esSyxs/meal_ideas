package food

var rec = []*Recepie{
	{
		uint(recID.id()),
		"Sviestmaize ar desu un sieru",
		`Uz maizes ar nazi uzsmērē sviestu.
		Nogriež sieru un desu, uzliek uz maizes.`,
		nil,
		[]*Produce{
			produce["maize"],
			produce["sviests"],
			produce["desa"],
			produce["siers"],
		},
		[]*Appliance{
			appliances["nazis"],
			appliances["dēlis"],
		},
	},
	{
		uint(recID.id()),
		"Sviestmaize ar desu, sieru un gurķi",
		`Uz maizes ar nazi uzsmērē sviestu. 
		Nogriež sieru un desu, uzliek uz maizes. 
		Nomazgā gurķi, sagriež, uzliek uz maizes.`,
		nil,
		[]*Produce{
			produce["maize"],
			produce["sviests"],
			produce["desa"],
			produce["siers"],
			produce["gurķis"],
		},
		[]*Appliance{
			appliances["nazis"],
			appliances["dēlis"],
		},
	},
	{
		uint(recID.id()),
		"Franču sīpolu zupa",
		`Nomazgā, nomizo un šķēlēs sagriež sīpolus un burkānus.
		Nomazgā un šķēlēs sagriež cukini un sēnes.
		Sagriež desu vienāda izmēra kubiņos.
		Smalki sagriež pētersīļus.
		Sarīvē sieru.
		Pannā uzkarsē sviestu un 1 minūti cep sīpolus ar rozmarīnu.
		Pievieno burkānus, sēnes, cukini, desu un cep 25 minūtes.
		Visu ievieto katlā.
		Pievieno sarkanvīnu, uzkarsē līdz vārīšanās temperatūrai, tad samazina karstumu un vāra līdz šķidrums ir iztvaikojis.
		Kad zupa gatava, virsū uzber sagrieztos pētersīļus un sarīvēto sieru.`,
		nil,
		[]*Produce{
			produce["sīpoli"],
			produce["burkāni"],
			produce["sēnes"],
			produce["cukini"],
			produce["desa"],
			produce["pētersīļi"],
			produce["rozmarīns"],
			produce["sarkanvīns"],
			produce["sviests"],
		},
		[]*Appliance{
			appliances["nazis"],
			appliances["dēlis"],
			appliances["rīve"],
			appliances["plīts"],
			appliances["panna"],
			appliances["cepamlāpstiņa"],
			appliances["katls"],
		},
	},
	{
		uint(recID.id()),
		"Griķi",
		`Noskalo griķus.
		Griķus ieber katlā ar ūdeni proporcijā 1:2.
		Kad ūdens sācis vārīties, ļauj tam burbuļot 10 minūtes, tad noņem no uguns un ļauj griķiem piebriest.
		Ja katlā palicis šķidrums, to nolej.`,
		nil,
		[]*Produce{
			produce["griķi"],
		},
		[]*Appliance{
			appliances["plīts"],
			appliances["katls"],
		},
	},
	{
		uint(recID.id()),
		"Griķu salāti",
		`Ievieto griķus katlā ar ūdeni un vāra līdz gatavi.
		Krāsnī uz restēm viegli apgrilē papriku. 
		Papriku vēl siltu sagriež mazos gabaliņos.
		Nomizo un sagriež sīpolus vidējos gabaliņos.
		Pannā uzkarsē sviestu un apcep sīpolus.
		Sarīvē sieru.
		Griķus ieber bļodā, pievieno silto papriku, sīpolus un sieru.`,
		nil,
		[]*Produce{
			produce["griķi"],
			produce["sīpoli"],
			produce["sviests"],
			produce["paprika"],
			produce["siers"],
		},
		[]*Appliance{
			appliances["plīts"],
			appliances["katls"],
			appliances["cepeškrāsns"],
			appliances["nazis"],
			appliances["dēlis"],
			appliances["panna"],
			appliances["cepamlāpstiņa"],
			appliances["rīve"],
			appliances["bļoda"],
		},
	},
	{
		uint(recID.id()),
		"Persiku biezpienmaizes",
		`Atlaidina kārtaino mīklu. 
		Ņem kārtainās mīklas plāksnīti, pārgriež uz pusēm, izrullē un iegriež divus pretējos stūrīšus taisnā leņķī.
		Sagriež persikus uz pusēm.
		Samaisa biezpienu ar krēmsieru un cukuru.
		Uz mīklas kvadrātiņa liek vienu karoti biezpiena masas, uz biezpiena liek persika pusīti un saloka pretējās mīklas maliņas pāri vienu otrai.
		Tāpat sagatavo arī pārējās mīklas plāksnes un izvieto tās uz cepešpannas.
		Krūzē ar dakšu sakuļ olas.
		Sagatavotās bulciņas pārsmērē ar sakultās olas un liek 175 grādos uzkarsētā cepeškrāsnī uz 20 minūtēm.`,
		nil,
		[]*Produce{
			produce["kārtainā mīkla"],
			produce["krēmsiers"],
			produce["persiki"],
			produce["biezpiens"],
			produce["cukurs"],
			produce["olas"],
		},
		[]*Appliance{
			appliances["nazis"],
			appliances["dēlis"],
			appliances["mīklas rullis"],
			appliances["karote"],
			appliances["bļoda"],
			appliances["krūze"],
			appliances["dakša"],
			appliances["cepešpanna"],
			appliances["cepeškrāsns"],
		},
	},
	{
		uint(recID.id()),
		"Persiku pīrāgs",
		`Izmīca mīkstu sviestu ar krējumu un miltiem.
		Mīklu ievieto cepešpannā.
		Liek cepeškrāsnī un cep līdz pamatne gatava, bet ne brūna.
		Persikus sagriež sķēlēs.
		Izņem pamatni no cepeškrāsns.
		Pamatnei liek virsū sagrieztos persikus.
		Samaisa krēmsieru ar olām un miltiem, ar šo mērci pārlej persikus.
		Liek cepties 180 grādos līdz virspuse ir viegli apbrūnināta.`,
		nil,
		[]*Produce{
			produce["sviests"],
			produce["krējums"],
			produce["krēmsiers"],
			produce["olas"],
			produce["milti"],
			produce["cukurs"],
			produce["persiki"],
		},
		[]*Appliance{
			appliances["bļoda"],
			appliances["cepešpanna"],
			appliances["cepeškrāsns"],
			appliances["nazis"],
			appliances["dēlis"],
			appliances["karote"],
		},
	},
	{
		uint(recID.id()),
		"Cepti kartupeļi",
		`Nomizo un sagriež kartupeļus.
		Ievieto pannā un cep līdz gatavi.`,
		nil,
		[]*Produce{
			produce["kartupeļi"],
		},
		[]*Appliance{
			appliances["nazis"],
			appliances["dēlis"],
			appliances["plīts"],
			appliances["panna"],
		},
	},
	{
		uint(recID.id()),
		"Vārīti kartupeļi",
		`Nomizo kartupeļus.
		Ievieto katlā ar ūdeni un vāra līdz gatavi.`,
		nil,
		[]*Produce{
			produce["kartupeļi"],
		},
		[]*Appliance{
			appliances["nazis"],
			appliances["plīts"],
			appliances["katls"],
		},
	},
	{
		uint(recID.id()),
		"Kartupeļu sacepums",
		`Nomizo kartupeļus un sagriež plānās sķēlītēs.
		Nomizo un sagriež sīpolus.
		Cepešpannā ieliek pāris piciņas sviesta un virsū nelielā kārtiņā izvieto kartupeļus.
		Virs kartupeļiem saliek sagrieztos sīpolus un pāris piciņas sviesta.
		Atkārto procesu un saliek šādas kārtas ar kartupeļiem, sīpoliem un sviestu līdz cepešpanna pilna.
		Pēdējo kārtu pārlej ar krējumu.
		Liek cepties uz 200 grādiem pusotru stundu.
		Sarīvē sieru un sagriež pētersīļus.
		Kad sacepums gandrīz gatavs, to pārkaisa ar sieru, pētersīļiem un liek cepeškrāsnī uz vēl 5 minūtēm.`,
		nil,
		[]*Produce{
			produce["kartupeļi"],
			produce["sīpoli"],
			produce["pētersīļi"],
			produce["krējums"],
			produce["siers"],
			produce["sviests"],
		},
		[]*Appliance{
			appliances["nazis"],
			appliances["dēlis"],
			appliances["cepešpanna"],
			appliances["cepeškrāsns"],
			appliances["rīve"],
		},
	},
	{
		uint(recID.id()),
		"Vistas fileja",
		`Vistas fileju nedaudz saplacina ar gaļas āmuru.
		Pannā izkausē piciņu sviesta.
		Pannā ievieto vistas fileju un apcep tai vienu pusi.
		Pēc brīža pievieno vēl vienu piciņu sviesta un vistas fileju apgriež uz otru pusi.
		Kad otra vistas filejas puse arī gatava, pannu noņem no uguns.
		Vistu atstāj uz pannas, lai pāris minūtes tā pasautējas savā sulā.`,
		nil,
		[]*Produce{
			produce["vistas fileja"],
			produce["sviests"],
		},
		[]*Appliance{
			appliances["plīts"],
			appliances["panna"],
			appliances["cepamlāpstiņa"],
			appliances["gaļas āmurs"],
		},
	},
	{
		uint(recID.id()),
		"Persiku un vistas salāti",
		`Vistas filejas nedaudz saplacina ar gaļas āmuru.
		Liek vistu uz uzkarsētas pannas un cep 5-7 minūtes no katras puses.
		Sagriež persikus uz pusēm un apcep uz pannas 3-4 minūtes.
		Šķēlēs sagriež sacepto vistu un persikus.
		Nomizo un sagriež sīpolus.
		Sarīvē sieru.
		Krūzē samaisa sinepes ar medu.
		Bļodā ievieto spinātus, sagriezto vistu, persikus, sīpolus un sieru.
		Salātus pārlej ar krūzē sagatvoto mērci.`,
		nil,
		[]*Produce{
			produce["persiki"],
			produce["vistas fileja"],
			produce["spināti"],
			produce["sīpoli"],
			produce["siers"],
			produce["sinepes"],
			produce["medus"],
		},
		[]*Appliance{
			appliances["gaļas āmurs"],
			appliances["plīts"],
			appliances["panna"],
			appliances["cepamlāpstiņa"],
			appliances["nazis"],
			appliances["dēlis"],
			appliances["krūze"],
			appliances["karote"],
			appliances["rīve"],
			appliances["bļoda"],
		},
	},
	{
		uint(recID.id()),
		"Cepti kartupeļi ar pētersīļiem",
		`Nomizo un sagriež kartupeļus.
		Ievieto pannā un cep līdz gatavi.
		Sagriež pētersīļus.
		Pārkaisa pētersīļus pāri kartupeļiem.`,
		nil,
		[]*Produce{
			produce["kartupeļi"],
			produce["pētersīļi"],
		},
		[]*Appliance{
			appliances["nazis"],
			appliances["dēlis"],
			appliances["plīts"],
			appliances["panna"],
		},
	},
	{
		uint(recID.id()),
		"Havajiešu karstmaize",
		`Uz maizes uzsmērē sviestu un nedaudz sinepes.
		Nogriež sieru un desu, uzliek uz maizes.
		Uzliek ananasa ripiņu uz maizes.
		Cep cepeškrāsnī līdz siers izkusis.`,
		nil,
		[]*Produce{
			produce["maize"],
			produce["sviests"],
			produce["siers"],
			produce["desa"],
			produce["ananass"],
			produce["sinepes"],
		},
		[]*Appliance{
			appliances["nazis"],
			appliances["dēlis"],
			appliances["cepeškrāsns"],
		},
	},
	{
		uint(recID.id()),
		"Rīsi",
		`Noskalo rīsus.
		Rīsus ieber katlā, aplejot tos ar karstu ūdeni.
		Liek katlu uz lielas uguns, lai ūdens nekavējoties uzvārītos. Tad uguni nogriež līdz minimālajai atzīmei un rīsus vāra 15 minūtes.`,
		nil,
		[]*Produce{
			produce["rīsi"],
		},
		[]*Appliance{
			appliances["plīts"],
			appliances["katls"],
		},
	},
}
