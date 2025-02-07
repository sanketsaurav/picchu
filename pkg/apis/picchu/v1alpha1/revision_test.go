package v1alpha1

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestFailRevision(t *testing.T) {
	r := Revision{}
	r.Fail()
	failedAt, ok := r.Annotations[AnnotationFailedAt]
	assert.True(t, ok, "FailedAt not set")
	d := r.SinceFailed()
	assert.NotEqual(t, d, time.Duration(0), "SinceFailed reports 0")
	r.Fail()
	failedAtNow := r.Annotations[AnnotationFailedAt]
	assert.Equal(t, failedAt, failedAtNow, "FailedAt shouldn't update on subsequent calls")
}

func TestExternalTestPending(t *testing.T) {
	target := &RevisionTarget{}
	assert.False(t, target.IsExternalTestPending())

	target.ExternalTest.Enabled = true
	assert.True(t, target.IsExternalTestPending())

	target.ExternalTest.Completed = true
	assert.False(t, target.IsExternalTestPending())
}

func TestCanaryTestPending(t *testing.T) {
	target := &RevisionTarget{}
	dt := metav1.Time{}
	assert.False(t, target.IsCanaryPending(&dt))

	target.Canary.Percent = 1
	assert.False(t, target.IsCanaryPending(&dt))

	target.Canary.TTL = 1
	now := metav1.Now()
	assert.True(t, target.IsCanaryPending(&now))

	lastSecond := metav1.Time{time.Now().Add(-time.Second)}
	assert.False(t, target.IsCanaryPending(&lastSecond))
}
