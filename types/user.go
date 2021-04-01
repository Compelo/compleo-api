package types

type User struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Cognome   string `json:"cognome"`
	CF        string `json:"cf"`
	Indirizzo string `json:"indirizzo"`
	//TODO: ELIMINARE I 3 CAMPI SEGUENTI E SOSTITUIRE CON UN SOLO CMPO CHE PUNTA A  ACTIVITY.CITTA
	Regione   string `json:"regione"`
	Citta     string `json:"citta"`
	Provincia string `json:"provincia"`

	Telefono string `json:"telefono"`
	EMail    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type POSTGotUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) CheckUser() bool {
	return u.Nome == "" || u.Cognome == "" || u.CF == "" || u.Indirizzo == "" || u.Regione == "" || u.Citta == "" || u.Provincia == "" || u.Telefono == "" || u.EMail == "" || u.Username == "" || u.Password == ""
}
