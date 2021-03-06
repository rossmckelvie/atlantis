// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/webhooks (interfaces: Sender)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	webhooks "github.com/runatlantis/atlantis/server/events/webhooks"
	logging "github.com/runatlantis/atlantis/server/logging"
	"reflect"
	"time"
)

type MockSender struct {
	fail func(message string, callerSkip ...int)
}

func NewMockSender() *MockSender {
	return &MockSender{fail: pegomock.GlobalFailHandler}
}

func (mock *MockSender) Send(log *logging.SimpleLogger, applyResult webhooks.ApplyResult) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSender().")
	}
	params := []pegomock.Param{log, applyResult}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Send", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockSender) VerifyWasCalledOnce() *VerifierSender {
	return &VerifierSender{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockSender) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierSender {
	return &VerifierSender{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockSender) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierSender {
	return &VerifierSender{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockSender) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierSender {
	return &VerifierSender{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierSender struct {
	mock                   *MockSender
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierSender) Send(log *logging.SimpleLogger, applyResult webhooks.ApplyResult) *Sender_Send_OngoingVerification {
	params := []pegomock.Param{log, applyResult}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Send", params, verifier.timeout)
	return &Sender_Send_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Sender_Send_OngoingVerification struct {
	mock              *MockSender
	methodInvocations []pegomock.MethodInvocation
}

func (c *Sender_Send_OngoingVerification) GetCapturedArguments() (*logging.SimpleLogger, webhooks.ApplyResult) {
	log, applyResult := c.GetAllCapturedArguments()
	return log[len(log)-1], applyResult[len(applyResult)-1]
}

func (c *Sender_Send_OngoingVerification) GetAllCapturedArguments() (_param0 []*logging.SimpleLogger, _param1 []webhooks.ApplyResult) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*logging.SimpleLogger, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*logging.SimpleLogger)
		}
		_param1 = make([]webhooks.ApplyResult, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(webhooks.ApplyResult)
		}
	}
	return
}
