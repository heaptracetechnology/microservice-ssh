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

var (
	host       = os.Getenv("SSH_HOST")
	port       = os.Getenv("SSH_PORT")
	privateKey = os.Getenv("SSH_PRIVATE_KEY")
)

var _ = Describe("SSH without private key and password", func() {

	ssh := SSHArguments{Command: "ifconfig", Username: "admin1", Host: host, Port: port}
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

var _ = Describe("SSH with password", func() {

	ssh := SSHArguments{Command: "ifconfig", Username: "admin1", Password: "admin", Host: host, Port: port}
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
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("SSH with private key and password", func() {

	os.Setenv("PRIVATE_KEY", privateKey)
	ssh := SSHArguments{Command: "ifconfig", Username: "admin1", Host: host, Password: "admin", Port: port}
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

var _ = Describe("SSH with private key", func() {

	os.Setenv("PRIVATE_KEY", privateKey)
	ssh := SSHArguments{Command: "ifconfig", Username: "admin1", Host: host, Port: port}
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
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
