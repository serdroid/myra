package meeting

import (
    "os"
    "errors"
    "strings"
    "log"
    "fmt"
    "database/sql"
	_ "github.com/lib/pq"
)

const (
    userEnv = "DB_USER"
    passwdEnv = "DB_PASSWD"
    dbEnv = "DB_URL"
)

type postgresDataStore struct {
	DB	*sql.DB
}

type connInfo struct {
    dbURL string
    user string
    passwd string
}

func (p *postgresDataStore) findMeeting(host string, date string) Meeting {
//	rows, err := p.DB.Query("select id, host, guest, date, duration from meeting where host = $1 and date = $2", host, date )
	rows, err := p.DB.Query("select id, host, guest, date, duration from meeting where host = $1 ", host)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	mt := Meeting{}
	for rows.Next() {
		err := rows.Scan(&mt.ID, &mt.Host, &mt.Guest, &mt.Date, &mt.Duration)
		if err != nil {
			log.Fatal(err)
		}
		// break for demo purposes, shouldn't break normally
		break
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return mt
}

func (p *postgresDataStore) createMeeting(m *Meeting) error {
    m.ID = randomString(16)
    _, err := p.DB.Exec("INSERT INTO meeting(id, host, guest, date, duration) VALUES($1, $2, $3, $4, $5)", 
        m.ID, m.Host, m.Guest, m.Date, m.Duration)
    return err
}

func NewPostgresDataStore() *dataStore {
	db := connectToDB()
    var store dataStore = &postgresDataStore{db}
    return &store
}

func connectToDB() *sql.DB {
	connStr := makeConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func makeConnectionString() string {
	ci := makeConnInfo()
	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"

	// I abused dbURL to dbname fix this before real usage
    return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", ci.user, ci.passwd, ci.dbURL)
}

func getDBEnv(env string) (string, error) {
    val, ok := os.LookupEnv(env)
    if ! ok {
        return "", errors.New("no such env var")
    }
    return val, nil
}

func checkEnv(envVar string, errs *[]string) string {
    val, err := getDBEnv(envVar)
    if err != nil {
        *errs = append(*errs, envVar)
    }
    return val
}

func makeConnInfo() connInfo {
    errs := []string{}
    user := checkEnv(userEnv, &errs)
    passwd := checkEnv(passwdEnv, &errs)
    dbURL := checkEnv(dbEnv, &errs)
    if len ( errs ) > 0 {
		vn := strings.Join(errs, ",")
		log.Printf("Environment variable(s) %s should be defined", vn)
    }
    return connInfo{dbURL, user, passwd}
}

