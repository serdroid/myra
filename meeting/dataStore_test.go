package meeting

import (
    "testing"
    "os"
)

func TestMakeConnectionString(test *testing.T) {
    os.Setenv(userEnv, "postgres")
    os.Setenv(passwdEnv, "password")
    os.Setenv(dbEnv, "postgres")
    got := makeConnectionString()
    want := "user=postgres password=password dbname=postgres sslmode=disable"
    
    if got != want {
        test.Errorf("got %v, want %v", got, want)
    }
}

func TestGetDBEnv(test *testing.T) {
    os.Setenv(userEnv, "postgres")
    got, err := getDBEnv("DB_USER")
    clearEnv()
    if err != nil {
        test.Fatalf("Don't expect an error, but got one: %v", err)
    }
    want := "postgres"
    
    if got != want {
        test.Errorf("got %v, want %v", got, want)
    }
}

func TestCheckEnv(test *testing.T) {
    os.Setenv(userEnv, "postgres")
    errs := []string{}
    checkEnv(userEnv, &errs)
    checkEnv(passwdEnv, &errs)
    checkEnv(dbEnv, &errs)
    got := len ( errs )
    want := 2
    clearEnv()
    if got != want {
        test.Fatalf("Unexpected error count. got %d, want %d", got, want)
    }
    if errs[0] != passwdEnv || errs[1] != dbEnv {
        test.Errorf("Unexpected error list. got %v, want %v", errs, []string{passwdEnv, dbEnv})
    }
}

func TestMakeConnInfo(test *testing.T) {
    os.Setenv(userEnv, "postgres")
    os.Setenv(passwdEnv, "password")
    os.Setenv(dbEnv, "localhost:5432/...")
    got := makeConnInfo()
    want := connInfo {"localhost:5432/...", "postgres", "password"}
    clearEnv()
    if got != want {
        test.Errorf("got %v, want %v", got, want)
    }
}

func TestMakeConnInfoEmptyEnv(test *testing.T) {
    os.Setenv(userEnv, "postgres")
    got := makeConnInfo()
    want := connInfo {"", "postgres", ""}
    clearEnv()
    if got != want {
        test.Errorf("got %v, want %v", got, want)
    }
}

func clearEnv() {
    os.Unsetenv(userEnv)
    os.Unsetenv(passwdEnv)
    os.Unsetenv(dbEnv)
}

