package sqsconsumer_test

import (
	"testing"

	"github.com/ikarpovich/sqsconsumer"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
)

func TestSQSInterfaceImplementsSQSAPI(t *testing.T) {
	assert.Implements(t, (*sqsconsumer.SQSAPI)(nil), sqs.New(session.New()))
}
