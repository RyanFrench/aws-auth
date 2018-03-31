package awsrole

import (
	"github.com/stretchr/testify/mock"
)

type mockedCredentials struct {
	mock.Mock
}

// func TestCacheCredentials(t *testing.T) {
// 	testCases := []struct {
// 		roleArn     string
// 		credentials *mockedCredentials
// 		throws      error
// 	}{
// 		{
// 			"test-role", &mockedCredentials{}, nil,
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		err := CacheCredentials(tc.roleArn, tc.credentials)
//
// 	}
// }
