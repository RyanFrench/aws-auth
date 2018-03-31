package awsrole

import (
	"fmt"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go/service/sts"
	log "github.com/sirupsen/logrus"
)

func cacheDirectoryPath(roleArn string) string {
	accountID, _, _ := getRoleDetails(roleArn)
	return os.Getenv("HOME") + "/.aws-role/" + accountID
}

func filePath(roleArn string) string {
	_, roleName, _ := getRoleDetails(roleArn)
	return cacheDirectoryPath(roleArn) + "/" + roleName
}

func createCacheDirectory(roleArn string) error {
	cacheDirectory := cacheDirectoryPath(roleArn)
	if _, err := os.Stat(cacheDirectory); err == nil {
		return nil
	}
	os.MkdirAll(cacheDirectory, os.ModePerm)
	return nil
}

func getRoleDetails(roleArn string) (string, string, error) {
	re := regexp.MustCompile(`arn:aws:iam::(?P<accountID>\d+):role/(?P<role>[a-zA-Z0-9\-]+)`)
	match := re.FindStringSubmatch(roleArn)
	fmt.Printf("%+v\n", match)
	return match[1], match[2], nil
}

// CacheCredentials will save the AWS credentials in a file in ~/.aws-role/<role-arn>
func CacheCredentials(roleArn string, settings *sts.Credentials) error {
	log.WithFields(log.Fields{
		"roleArn":  roleArn,
		"settings": settings,
	}).Debug("Saving credentials to home Directory")
	createCacheDirectory(roleArn)
	cacheFile := filePath(roleArn)
	// Delete the file if it already exists
	if _, err := os.Stat(cacheFile); err == nil {
		os.Remove(cacheFile)
	}
	f, err := os.Create(cacheFile)
	if err != nil {
		log.WithError(err).Errorln("Error while creating cache file")
		return err
	}
	defer f.Close()
	f.WriteString(settings.String())
	return nil
}
