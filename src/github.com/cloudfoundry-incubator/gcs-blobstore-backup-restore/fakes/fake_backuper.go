// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	gcs "github.com/cloudfoundry-incubator/gcs-blobstore-backup-restore"
)

type FakeBackuper struct {
	CreateLiveBucketSnapshotStub        func() (map[string]gcs.BackupBucketAddress, map[string][]gcs.Blob, error)
	createLiveBucketSnapshotMutex       sync.RWMutex
	createLiveBucketSnapshotArgsForCall []struct{}
	createLiveBucketSnapshotReturns     struct {
		result1 map[string]gcs.BackupBucketAddress
		result2 map[string][]gcs.Blob
		result3 error
	}
	createLiveBucketSnapshotReturnsOnCall map[int]struct {
		result1 map[string]gcs.BackupBucketAddress
		result2 map[string][]gcs.Blob
		result3 error
	}
	CopyBlobsWithinBackupBucketStub        func(map[string]gcs.BackupBucketAddress, map[string][]gcs.Blob) error
	copyBlobsWithinBackupBucketMutex       sync.RWMutex
	copyBlobsWithinBackupBucketArgsForCall []struct {
		arg1 map[string]gcs.BackupBucketAddress
		arg2 map[string][]gcs.Blob
	}
	copyBlobsWithinBackupBucketReturns struct {
		result1 error
	}
	copyBlobsWithinBackupBucketReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackuper) CreateLiveBucketSnapshot() (map[string]gcs.BackupBucketAddress, map[string][]gcs.Blob, error) {
	fake.createLiveBucketSnapshotMutex.Lock()
	ret, specificReturn := fake.createLiveBucketSnapshotReturnsOnCall[len(fake.createLiveBucketSnapshotArgsForCall)]
	fake.createLiveBucketSnapshotArgsForCall = append(fake.createLiveBucketSnapshotArgsForCall, struct{}{})
	fake.recordInvocation("CreateLiveBucketSnapshot", []interface{}{})
	fake.createLiveBucketSnapshotMutex.Unlock()
	if fake.CreateLiveBucketSnapshotStub != nil {
		return fake.CreateLiveBucketSnapshotStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.createLiveBucketSnapshotReturns.result1, fake.createLiveBucketSnapshotReturns.result2, fake.createLiveBucketSnapshotReturns.result3
}

func (fake *FakeBackuper) CreateLiveBucketSnapshotCallCount() int {
	fake.createLiveBucketSnapshotMutex.RLock()
	defer fake.createLiveBucketSnapshotMutex.RUnlock()
	return len(fake.createLiveBucketSnapshotArgsForCall)
}

func (fake *FakeBackuper) CreateLiveBucketSnapshotReturns(result1 map[string]gcs.BackupBucketAddress, result2 map[string][]gcs.Blob, result3 error) {
	fake.CreateLiveBucketSnapshotStub = nil
	fake.createLiveBucketSnapshotReturns = struct {
		result1 map[string]gcs.BackupBucketAddress
		result2 map[string][]gcs.Blob
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBackuper) CreateLiveBucketSnapshotReturnsOnCall(i int, result1 map[string]gcs.BackupBucketAddress, result2 map[string][]gcs.Blob, result3 error) {
	fake.CreateLiveBucketSnapshotStub = nil
	if fake.createLiveBucketSnapshotReturnsOnCall == nil {
		fake.createLiveBucketSnapshotReturnsOnCall = make(map[int]struct {
			result1 map[string]gcs.BackupBucketAddress
			result2 map[string][]gcs.Blob
			result3 error
		})
	}
	fake.createLiveBucketSnapshotReturnsOnCall[i] = struct {
		result1 map[string]gcs.BackupBucketAddress
		result2 map[string][]gcs.Blob
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBackuper) CopyBlobsWithinBackupBucket(arg1 map[string]gcs.BackupBucketAddress, arg2 map[string][]gcs.Blob) error {
	fake.copyBlobsWithinBackupBucketMutex.Lock()
	ret, specificReturn := fake.copyBlobsWithinBackupBucketReturnsOnCall[len(fake.copyBlobsWithinBackupBucketArgsForCall)]
	fake.copyBlobsWithinBackupBucketArgsForCall = append(fake.copyBlobsWithinBackupBucketArgsForCall, struct {
		arg1 map[string]gcs.BackupBucketAddress
		arg2 map[string][]gcs.Blob
	}{arg1, arg2})
	fake.recordInvocation("CopyBlobsWithinBackupBucket", []interface{}{arg1, arg2})
	fake.copyBlobsWithinBackupBucketMutex.Unlock()
	if fake.CopyBlobsWithinBackupBucketStub != nil {
		return fake.CopyBlobsWithinBackupBucketStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyBlobsWithinBackupBucketReturns.result1
}

func (fake *FakeBackuper) CopyBlobsWithinBackupBucketCallCount() int {
	fake.copyBlobsWithinBackupBucketMutex.RLock()
	defer fake.copyBlobsWithinBackupBucketMutex.RUnlock()
	return len(fake.copyBlobsWithinBackupBucketArgsForCall)
}

func (fake *FakeBackuper) CopyBlobsWithinBackupBucketArgsForCall(i int) (map[string]gcs.BackupBucketAddress, map[string][]gcs.Blob) {
	fake.copyBlobsWithinBackupBucketMutex.RLock()
	defer fake.copyBlobsWithinBackupBucketMutex.RUnlock()
	return fake.copyBlobsWithinBackupBucketArgsForCall[i].arg1, fake.copyBlobsWithinBackupBucketArgsForCall[i].arg2
}

func (fake *FakeBackuper) CopyBlobsWithinBackupBucketReturns(result1 error) {
	fake.CopyBlobsWithinBackupBucketStub = nil
	fake.copyBlobsWithinBackupBucketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackuper) CopyBlobsWithinBackupBucketReturnsOnCall(i int, result1 error) {
	fake.CopyBlobsWithinBackupBucketStub = nil
	if fake.copyBlobsWithinBackupBucketReturnsOnCall == nil {
		fake.copyBlobsWithinBackupBucketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyBlobsWithinBackupBucketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackuper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createLiveBucketSnapshotMutex.RLock()
	defer fake.createLiveBucketSnapshotMutex.RUnlock()
	fake.copyBlobsWithinBackupBucketMutex.RLock()
	defer fake.copyBlobsWithinBackupBucketMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBackuper) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ gcs.Backuper = new(FakeBackuper)