module system-tests

require (
	github.com/onsi/ginkgo v1.15.2
	github.com/onsi/gomega v1.11.0
	s3-blobstore-backup-restore v0.0.0
)

replace system-tests => ./

replace s3-blobstore-backup-restore => ../s3-blobstore-backup-restore
