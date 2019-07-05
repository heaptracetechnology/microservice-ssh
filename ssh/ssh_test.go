package ssh

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("SSH without private key and password", func() {

	ssh := SSHArguments{Command: "ifconfig", Username: "mockUsername"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(ssh)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/ssh", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SSH)
	handler.ServeHTTP(recorder, request)

	Describe("Run command on SSH server", func() {
		Context("SSH server", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("SSH with private key and password", func() {

	os.Setenv("PRIVATE_KEY", "mockPrivateKey")
	ssh := SSHArguments{Command: "ifconfig", Username: "mockUsername", Password: "mockPassword"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(ssh)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/ssh", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SSH)
	handler.ServeHTTP(recorder, request)

	Describe("Run command on SSH server", func() {
		Context("SSH server", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("SSH with private key and password", func() {

	os.Setenv("PRIVATE_KEY", "")
	ssh := SSHArguments{Command: "ifconfig", Username: "mockUsername", Password: "mockPassword"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(ssh)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/ssh", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SSH)
	handler.ServeHTTP(recorder, request)

	Describe("Run command on SSH server", func() {
		Context("SSH server", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})
