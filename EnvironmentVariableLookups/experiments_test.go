package testers

import (
	"os"
	"testing"
)

//gistsnip:start:EnvLookup
func BenchmarkEnvLookup(b *testing.B) {
	err := os.Setenv("GO_TEST_VAR", "ENV-Potatoes")
	if err != nil {
		b.Logf("Failed to set env var. Error: %s", err)
		b.Fail()
	}
	var output string
	for i := 0; i < b.N; i++ {
		output = os.Getenv("GO_TEST_VAR")
	}
	b.Logf("got output: %s", output)
}

//gistsnip:end:EnvLookup

//gistsnip:start:StructLookup
func BenchmarkStructLookup(b *testing.B) {
	s := struct {
		hostname string
	}{
		hostname: "STRUCT-Potatoes",
	}
	var output string
	for i := 0; i < b.N; i++ {
		output = s.hostname
	}
	b.Logf("got output: %s", output)
}

//gistsnip:end:StructLookup

//gistsnip:start:DoubleLookupBlankEnv
func BenchmarkDoubleLookupBlankEnv(b *testing.B) {
	s := struct {
		hostname string
	}{
		hostname: "STRUCT-Potatoes",
	}
	var output string
	for i := 0; i < b.N; i++ {
		env, ok := os.LookupEnv("Not_Available")
		if ok {
			output = env
		} else {
			output = s.hostname
		}
	}
	b.Logf("got output: %s", output)
}

//gistsnip:end:DoubleLookupBlankEnv

//gistsnip:start:DoubleLookupValidEnv
func BenchmarkDoubleLookupValidEnv(b *testing.B) {
	err := os.Setenv("GO_TEST_VAR", "ENV-Potatoes")
	if err != nil {
		b.Logf("Failed to set env var. Error: %s", err)
		b.Fail()
	}
	s := struct {
		hostname string
	}{
		hostname: "STRUCT-Potatoes",
	}
	var output string
	for i := 0; i < b.N; i++ {
		env, ok := os.LookupEnv("GO_TEST_VAR")
		if ok {
			output = env
		} else {
			output = s.hostname
		}
	}
	b.Logf("got output: %s", output)
}

//gistsnip:end:DoubleLookupValidEnv

//gistsnip:start:UseStruct
func BenchmarkUseStruct(b *testing.B) {
	err := os.Setenv("GO_TEST_VAR", "ENV-Potatoes")
	if err != nil {
		b.Logf("Failed to set env var. Error: %s", err)
		b.Fail()
	}
	s := struct {
		hostname string
	}{
		hostname: "STRUCT-Potatoes",
	}
	var output string
	for i := 0; i < b.N; i++ {
		if s.hostname != "" {
			output = s.hostname
			continue
		}

		// We don't reach this in this test
		env, ok := os.LookupEnv("Not_Available")
		if !ok {
			output = env
		} else {
			output = s.hostname
		}
	}
	b.Logf("got output: %s", output)
}

//gistsnip:end:UseStruct

//gistsnip:start:UseEnv
func BenchmarkUseEnv(b *testing.B) {
	err := os.Setenv("GO_TEST_VAR", "ENV-Potatoes")
	if err != nil {
		b.Logf("Failed to set env var. Error: %s", err)
		b.Fail()
	}
	s := struct {
		hostname string
	}{}
	var output string
	for i := 0; i < b.N; i++ {

		// This returns false in this test
		if s.hostname != "" {
			output = s.hostname
			continue
		}

		// We now reach this code
		env, ok := os.LookupEnv("GO_TEST_VAR")
		if !ok {
			output = env
		}
	}
	b.Logf("got output: %s", output)
}

//gistsnip:end:UseEnv

//gistsnip:start:LenLookup
func BenchmarkLenLookup(b *testing.B) {
	s := struct {
		hostname string
	}{
		hostname: "STRUCT-Potatoes",
	}

	var output int
	for i := 0; i < b.N; i++ {
		if len(s.hostname) > 1 {
			output = 1
		}
	}
	b.Logf("using output: %d", output)
}

//gistsnip:end:LenLookup

//gistsnip:start:StringLookup
func BenchmarkStringLookup(b *testing.B) {
	s := struct {
		hostname string
	}{
		hostname: "STRUCT-Potatoes",
	}

	var output int
	for i := 0; i < b.N; i++ {
		if s.hostname != "" {
			output = 1
		}
	}
	b.Logf("using output: %d", output)
}

//gistsnip:end:StringLookup
