package conf

import (
	"testing"
)

// AppConfig app conf
type AppConfig struct {
	App struct {
		Mode string
		Host string
		Port int
	}
}

func TestConfig(t *testing.T) {
	err := Load("../../config/")
	if err != nil {
		t.Error(err)
	}

	testConfig(t)
}

func testConfig(t *testing.T) {
	mode := Get("app.mode")
	t.Logf("mode: %s", mode)
	mode2 := File("sea").Get("app.mode")
	t.Logf("mode2: %s", mode2)

	addr := File("test").GetInt("http.addr")
	t.Logf("http: %d", addr)

	var appConfig AppConfig
	sea := File("sea")
	err := sea.Unmarshal(&appConfig)
	if err != nil {
		t.Errorf("error: %d", err)
	}

	t.Logf("AppConfig: %v", appConfig)
}
