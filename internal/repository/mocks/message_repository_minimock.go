// Code generated by http://github.com/gojuno/minimock (v3.4.2). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/marinaaaniram/go-chat-server/internal/repository.MessageRepository -o message_repository_minimock.go -n MessageRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/marinaaaniram/go-chat-server/internal/model"
)

// MessageRepositoryMock implements repository.MessageRepository
type MessageRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcSend          func(ctx context.Context, message *model.Message) (err error)
	inspectFuncSend   func(ctx context.Context, message *model.Message)
	afterSendCounter  uint64
	beforeSendCounter uint64
	SendMock          mMessageRepositoryMockSend
}

// NewMessageRepositoryMock returns a mock for repository.MessageRepository
func NewMessageRepositoryMock(t minimock.Tester) *MessageRepositoryMock {
	m := &MessageRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SendMock = mMessageRepositoryMockSend{mock: m}
	m.SendMock.callArgs = []*MessageRepositoryMockSendParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mMessageRepositoryMockSend struct {
	optional           bool
	mock               *MessageRepositoryMock
	defaultExpectation *MessageRepositoryMockSendExpectation
	expectations       []*MessageRepositoryMockSendExpectation

	callArgs []*MessageRepositoryMockSendParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// MessageRepositoryMockSendExpectation specifies expectation struct of the MessageRepository.Send
type MessageRepositoryMockSendExpectation struct {
	mock      *MessageRepositoryMock
	params    *MessageRepositoryMockSendParams
	paramPtrs *MessageRepositoryMockSendParamPtrs
	results   *MessageRepositoryMockSendResults
	Counter   uint64
}

// MessageRepositoryMockSendParams contains parameters of the MessageRepository.Send
type MessageRepositoryMockSendParams struct {
	ctx     context.Context
	message *model.Message
}

// MessageRepositoryMockSendParamPtrs contains pointers to parameters of the MessageRepository.Send
type MessageRepositoryMockSendParamPtrs struct {
	ctx     *context.Context
	message **model.Message
}

// MessageRepositoryMockSendResults contains results of the MessageRepository.Send
type MessageRepositoryMockSendResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmSend *mMessageRepositoryMockSend) Optional() *mMessageRepositoryMockSend {
	mmSend.optional = true
	return mmSend
}

// Expect sets up expected params for MessageRepository.Send
func (mmSend *mMessageRepositoryMockSend) Expect(ctx context.Context, message *model.Message) *mMessageRepositoryMockSend {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("MessageRepositoryMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &MessageRepositoryMockSendExpectation{}
	}

	if mmSend.defaultExpectation.paramPtrs != nil {
		mmSend.mock.t.Fatalf("MessageRepositoryMock.Send mock is already set by ExpectParams functions")
	}

	mmSend.defaultExpectation.params = &MessageRepositoryMockSendParams{ctx, message}
	for _, e := range mmSend.expectations {
		if minimock.Equal(e.params, mmSend.defaultExpectation.params) {
			mmSend.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSend.defaultExpectation.params)
		}
	}

	return mmSend
}

// ExpectCtxParam1 sets up expected param ctx for MessageRepository.Send
func (mmSend *mMessageRepositoryMockSend) ExpectCtxParam1(ctx context.Context) *mMessageRepositoryMockSend {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("MessageRepositoryMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &MessageRepositoryMockSendExpectation{}
	}

	if mmSend.defaultExpectation.params != nil {
		mmSend.mock.t.Fatalf("MessageRepositoryMock.Send mock is already set by Expect")
	}

	if mmSend.defaultExpectation.paramPtrs == nil {
		mmSend.defaultExpectation.paramPtrs = &MessageRepositoryMockSendParamPtrs{}
	}
	mmSend.defaultExpectation.paramPtrs.ctx = &ctx

	return mmSend
}

// ExpectMessageParam2 sets up expected param message for MessageRepository.Send
func (mmSend *mMessageRepositoryMockSend) ExpectMessageParam2(message *model.Message) *mMessageRepositoryMockSend {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("MessageRepositoryMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &MessageRepositoryMockSendExpectation{}
	}

	if mmSend.defaultExpectation.params != nil {
		mmSend.mock.t.Fatalf("MessageRepositoryMock.Send mock is already set by Expect")
	}

	if mmSend.defaultExpectation.paramPtrs == nil {
		mmSend.defaultExpectation.paramPtrs = &MessageRepositoryMockSendParamPtrs{}
	}
	mmSend.defaultExpectation.paramPtrs.message = &message

	return mmSend
}

// Inspect accepts an inspector function that has same arguments as the MessageRepository.Send
func (mmSend *mMessageRepositoryMockSend) Inspect(f func(ctx context.Context, message *model.Message)) *mMessageRepositoryMockSend {
	if mmSend.mock.inspectFuncSend != nil {
		mmSend.mock.t.Fatalf("Inspect function is already set for MessageRepositoryMock.Send")
	}

	mmSend.mock.inspectFuncSend = f

	return mmSend
}

// Return sets up results that will be returned by MessageRepository.Send
func (mmSend *mMessageRepositoryMockSend) Return(err error) *MessageRepositoryMock {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("MessageRepositoryMock.Send mock is already set by Set")
	}

	if mmSend.defaultExpectation == nil {
		mmSend.defaultExpectation = &MessageRepositoryMockSendExpectation{mock: mmSend.mock}
	}
	mmSend.defaultExpectation.results = &MessageRepositoryMockSendResults{err}
	return mmSend.mock
}

// Set uses given function f to mock the MessageRepository.Send method
func (mmSend *mMessageRepositoryMockSend) Set(f func(ctx context.Context, message *model.Message) (err error)) *MessageRepositoryMock {
	if mmSend.defaultExpectation != nil {
		mmSend.mock.t.Fatalf("Default expectation is already set for the MessageRepository.Send method")
	}

	if len(mmSend.expectations) > 0 {
		mmSend.mock.t.Fatalf("Some expectations are already set for the MessageRepository.Send method")
	}

	mmSend.mock.funcSend = f
	return mmSend.mock
}

// When sets expectation for the MessageRepository.Send which will trigger the result defined by the following
// Then helper
func (mmSend *mMessageRepositoryMockSend) When(ctx context.Context, message *model.Message) *MessageRepositoryMockSendExpectation {
	if mmSend.mock.funcSend != nil {
		mmSend.mock.t.Fatalf("MessageRepositoryMock.Send mock is already set by Set")
	}

	expectation := &MessageRepositoryMockSendExpectation{
		mock:   mmSend.mock,
		params: &MessageRepositoryMockSendParams{ctx, message},
	}
	mmSend.expectations = append(mmSend.expectations, expectation)
	return expectation
}

// Then sets up MessageRepository.Send return parameters for the expectation previously defined by the When method
func (e *MessageRepositoryMockSendExpectation) Then(err error) *MessageRepositoryMock {
	e.results = &MessageRepositoryMockSendResults{err}
	return e.mock
}

// Times sets number of times MessageRepository.Send should be invoked
func (mmSend *mMessageRepositoryMockSend) Times(n uint64) *mMessageRepositoryMockSend {
	if n == 0 {
		mmSend.mock.t.Fatalf("Times of MessageRepositoryMock.Send mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmSend.expectedInvocations, n)
	return mmSend
}

func (mmSend *mMessageRepositoryMockSend) invocationsDone() bool {
	if len(mmSend.expectations) == 0 && mmSend.defaultExpectation == nil && mmSend.mock.funcSend == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmSend.mock.afterSendCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmSend.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Send implements repository.MessageRepository
func (mmSend *MessageRepositoryMock) Send(ctx context.Context, message *model.Message) (err error) {
	mm_atomic.AddUint64(&mmSend.beforeSendCounter, 1)
	defer mm_atomic.AddUint64(&mmSend.afterSendCounter, 1)

	if mmSend.inspectFuncSend != nil {
		mmSend.inspectFuncSend(ctx, message)
	}

	mm_params := MessageRepositoryMockSendParams{ctx, message}

	// Record call args
	mmSend.SendMock.mutex.Lock()
	mmSend.SendMock.callArgs = append(mmSend.SendMock.callArgs, &mm_params)
	mmSend.SendMock.mutex.Unlock()

	for _, e := range mmSend.SendMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSend.SendMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSend.SendMock.defaultExpectation.Counter, 1)
		mm_want := mmSend.SendMock.defaultExpectation.params
		mm_want_ptrs := mmSend.SendMock.defaultExpectation.paramPtrs

		mm_got := MessageRepositoryMockSendParams{ctx, message}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmSend.t.Errorf("MessageRepositoryMock.Send got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.message != nil && !minimock.Equal(*mm_want_ptrs.message, mm_got.message) {
				mmSend.t.Errorf("MessageRepositoryMock.Send got unexpected parameter message, want: %#v, got: %#v%s\n", *mm_want_ptrs.message, mm_got.message, minimock.Diff(*mm_want_ptrs.message, mm_got.message))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSend.t.Errorf("MessageRepositoryMock.Send got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSend.SendMock.defaultExpectation.results
		if mm_results == nil {
			mmSend.t.Fatal("No results are set for the MessageRepositoryMock.Send")
		}
		return (*mm_results).err
	}
	if mmSend.funcSend != nil {
		return mmSend.funcSend(ctx, message)
	}
	mmSend.t.Fatalf("Unexpected call to MessageRepositoryMock.Send. %v %v", ctx, message)
	return
}

// SendAfterCounter returns a count of finished MessageRepositoryMock.Send invocations
func (mmSend *MessageRepositoryMock) SendAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.afterSendCounter)
}

// SendBeforeCounter returns a count of MessageRepositoryMock.Send invocations
func (mmSend *MessageRepositoryMock) SendBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSend.beforeSendCounter)
}

// Calls returns a list of arguments used in each call to MessageRepositoryMock.Send.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSend *mMessageRepositoryMockSend) Calls() []*MessageRepositoryMockSendParams {
	mmSend.mutex.RLock()

	argCopy := make([]*MessageRepositoryMockSendParams, len(mmSend.callArgs))
	copy(argCopy, mmSend.callArgs)

	mmSend.mutex.RUnlock()

	return argCopy
}

// MinimockSendDone returns true if the count of the Send invocations corresponds
// the number of defined expectations
func (m *MessageRepositoryMock) MinimockSendDone() bool {
	if m.SendMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.SendMock.invocationsDone()
}

// MinimockSendInspect logs each unmet expectation
func (m *MessageRepositoryMock) MinimockSendInspect() {
	for _, e := range m.SendMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageRepositoryMock.Send with params: %#v", *e.params)
		}
	}

	afterSendCounter := mm_atomic.LoadUint64(&m.afterSendCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.SendMock.defaultExpectation != nil && afterSendCounter < 1 {
		if m.SendMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageRepositoryMock.Send")
		} else {
			m.t.Errorf("Expected call to MessageRepositoryMock.Send with params: %#v", *m.SendMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSend != nil && afterSendCounter < 1 {
		m.t.Error("Expected call to MessageRepositoryMock.Send")
	}

	if !m.SendMock.invocationsDone() && afterSendCounter > 0 {
		m.t.Errorf("Expected %d calls to MessageRepositoryMock.Send but found %d calls",
			mm_atomic.LoadUint64(&m.SendMock.expectedInvocations), afterSendCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockSendInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MessageRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSendDone()
}
