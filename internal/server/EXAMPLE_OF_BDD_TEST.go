package server

// import "net/http/httptest"

// package secure

// import (
// 	"crypto/sha512"
// 	"crypto/subtle"
// 	"encoding/base64"
// 	"encoding/hex"
// 	"fmt"
// 	"strings"
// 	"testing"

// 	"golang.org/x/crypto/argon2"
// 	. "gopkg.in/check.v1"
// )

// func TestSecure(t *testing.T) { TestingT(t) }

// var _ = Suite(&SecureTestSuite{})

// type SecureTestSuite struct {
// 	params *argonCrypt
// }

// // hashData Assumed that receiver first hashes the data before proceeding
// func hashData(data string) string {
// 	password := sha512.New()
// 	password.Sum(([]byte(data)))
// 	return hex.EncodeToString(password.Sum(nil))
// }

// func (s *SecureTestSuite) SetUpTest(c *C) {
// 	s.params = setup()
// }

// func (s *SecureTestSuite) TestHashValue(c *C) {
// 	// Given a sha512 hash
// 	data := hashData("123456")
// 	// When a value sha512 hash is rehashed with Argon2ID
// 	argonHash := HashValue(data)
// 	vals := strings.Split(argonHash, "$")
// 	var version int
// 	_, err := fmt.Sscanf(vals[2], "v=%d", &version)
// 	c.Check(err, IsNil)
// 	c.Assert(version, Equals, argon2.Version)
// 	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &s.params.memory, &s.params.iterations, &s.params.parallelism)
// 	c.Check(err, IsNil)
// 	salt, err := base64.RawStdEncoding.DecodeString(vals[4])
// 	c.Check(err, IsNil)
// 	s.params.saltLength = uint32(len(salt))
// 	hash, err := base64.RawStdEncoding.DecodeString(vals[5])
// 	c.Check(err, IsNil)
// 	s.params.keyLength = uint32(len(hash))
// 	otherHash := argon2.IDKey([]byte(data), salt, s.params.iterations, s.params.memory, s.params.parallelism, s.params.keyLength)
// 	c.Check(subtle.ConstantTimeCompare(hash, otherHash), Equals, 1)
// }

// func (s *SecureTestSuite) TestDecodeHash(c *C) {
// 	// Given a sha512 hash
// 	data := hashData("123456")
// 	// When a value sha512 hash is rehashed with Argon2ID
// 	argonHash := HashValue(data)
// 	vals := strings.Split(argonHash, "$")
// 	p, salt, hash, err := decodeHash(argonHash)
// 	c.Check(err, IsNil)
// 	c.Assert(p, DeepEquals, setup())
// 	compareSalt, err := base64.RawStdEncoding.DecodeString(vals[4])
// 	c.Check(err, IsNil)
// 	c.Assert(salt, DeepEquals, compareSalt)
// 	compareHash, err := base64.RawStdEncoding.DecodeString(vals[5])
// 	c.Check(err, IsNil)
// 	c.Assert(hash, DeepEquals, compareHash)
// }

// func (s *SecureTestSuite) TestVerifyValue(c *C) {
// 	// Given a sha512 hash
// 	data := hashData("123456")
// 	// When a value sha512 hash is rehashed with Argon2ID
// 	argonHash := HashValue(data)
// 	// And we verify our has values when decoded
// 	match, err := VerifyValue(data, argonHash)
// 	// Then expect there to be a match
// 	c.Check(err, IsNil)
// 	c.Assert(match, Equals, true)
// }

// https://speedscale.com/blog/testing-golang-with-httptest/


func TestSimpleHandlerPassingCacheTest(){
	// Given a Http Tests for calling GetWeatherForecastHandler
	// And when something is stored in the cache

	// When I call the handler
	// Then Expect the cache result to be returned
	
	// Then compare the result with my expected values


	
}
