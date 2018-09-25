package gcs

import (
	"fmt"
	"strings"
)

type Restorer struct {
	bucketPairs       map[string]BucketPair
	executionStrategy Strategy
}

func NewRestorer(bucketPairs map[string]BucketPair, executionStrategy Strategy) Restorer {
	return Restorer{
		bucketPairs:       bucketPairs,
		executionStrategy: executionStrategy,
	}
}

func (r Restorer) Restore(backups map[string]BucketBackup) error {
	for bucketIdentifier := range backups {
		_, exists := r.bucketPairs[bucketIdentifier]
		if !exists {
			return fmt.Errorf("bucket identifier '%s' not found in bucketPairs configuration", bucketIdentifier)
		}
	}

	for bucketIdentifier, backup := range backups {
		bucket := r.bucketPairs[bucketIdentifier]

		errs := r.executionStrategy.Run(backup.Blobs, func(blob Blob) error {
			return bucket.liveBucket.CopyVersion(blob, backup.LiveBucketName)
		})

		if len(errs) != 0 {
			return formatErrors(fmt.Sprintf("failed to restore bucket '%s'", bucket.liveBucket.Name()), errs)
		}
	}

	return nil
}

func formatErrors(contextString string, errors []error) error {
	errorStrings := make([]string, len(errors))
	for i, err := range errors {
		errorStrings[i] = err.Error()
	}
	return fmt.Errorf("%s: %s", contextString, strings.Join(errorStrings, "\n"))
}
