package tcp

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "net/http/httptest"
  "net/http"
  "github.com/gorilla/mux"
  "bytes"
)

var _ = Describe("startOpHandler", func() {

  Context("ServeHTTP() method", func() {
    It("should return StatusCode of 400 if body of request is malformed", func() {

      /* arrange */
      objectUnderTest := startOpHandler{}
      recorder := httptest.NewRecorder()
      m := mux.NewRouter()
      m.Handle(startOpRelUrlTemplate, objectUnderTest)

      httpReq, err := http.NewRequest(http.MethodPost, startOpRelUrlTemplate, bytes.NewReader([]byte{}))
      if (nil != err) {
        Fail(err.Error())
      }

      /* act */
      m.ServeHTTP(recorder, httpReq)

      /* assert */
      Expect(recorder.Code).To(Equal(http.StatusBadRequest))

    })
  })
})
