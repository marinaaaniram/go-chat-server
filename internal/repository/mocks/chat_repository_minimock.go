// Code generated by http://github.com/gojuno/minimock (v3.4.2). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/marinaaaniram/go-chat-server/internal/repository.ChatRepository -o chat_repository_minimock.go -n ChatRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/marinaaaniram/go-chat-server/internal/model"
)

// ChatRepositoryMock implements repository.ChatRepository
type ChatRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, chat *model.Chat) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, chat *model.Chat)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mChatRepositoryMockCreate

	funcDelete          func(ctx context.Context, chat *model.Chat) (err error)
	inspectFuncDelete   func(ctx context.Context, chat *model.Chat)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mChatRepositoryMockDelete
}

// NewChatRepositoryMock returns a mock for repository.ChatRepository
func NewChatRepositoryMock(t minimock.Tester) *ChatRepositoryMock {
	m := &ChatRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mChatRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*ChatRepositoryMockCreateParams{}

	m.DeleteMock = mChatRepositoryMockDelete{mock: m}
	m.DeleteMock.callArgs = []*ChatRepositoryMockDeleteParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mChatRepositoryMockCreate struct {
	optional           bool
	mock               *ChatRepositoryMock
	defaultExpectation *ChatRepositoryMockCreateExpectation
	expectations       []*ChatRepositoryMockCreateExpectation

	callArgs []*ChatRepositoryMockCreateParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// ChatRepositoryMockCreateExpectation specifies expectation struct of the ChatRepository.Create
type ChatRepositoryMockCreateExpectation struct {
	mock      *ChatRepositoryMock
	params    *ChatRepositoryMockCreateParams
	paramPtrs *ChatRepositoryMockCreateParamPtrs
	results   *ChatRepositoryMockCreateResults
	Counter   uint64
}

// ChatRepositoryMockCreateParams contains parameters of the ChatRepository.Create
type ChatRepositoryMockCreateParams struct {
	ctx  context.Context
	chat *model.Chat
}

// ChatRepositoryMockCreateParamPtrs contains pointers to parameters of the ChatRepository.Create
type ChatRepositoryMockCreateParamPtrs struct {
	ctx  *context.Context
	chat **model.Chat
}

// ChatRepositoryMockCreateResults contains results of the ChatRepository.Create
type ChatRepositoryMockCreateResults struct {
	i1  int64
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreate *mChatRepositoryMockCreate) Optional() *mChatRepositoryMockCreate {
	mmCreate.optional = true
	return mmCreate
}

// Expect sets up expected params for ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) Expect(ctx context.Context, chat *model.Chat) *mChatRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatRepositoryMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.paramPtrs != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by ExpectParams functions")
	}

	mmCreate.defaultExpectation.params = &ChatRepositoryMockCreateParams{ctx, chat}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// ExpectCtxParam1 sets up expected param ctx for ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) ExpectCtxParam1(ctx context.Context) *mChatRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatRepositoryMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &ChatRepositoryMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.ctx = &ctx

	return mmCreate
}

// ExpectChatParam2 sets up expected param chat for ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) ExpectChatParam2(chat *model.Chat) *mChatRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatRepositoryMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &ChatRepositoryMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.chat = &chat

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) Inspect(f func(ctx context.Context, chat *model.Chat)) *mChatRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for ChatRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) Return(i1 int64, err error) *ChatRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &ChatRepositoryMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the ChatRepository.Create method
func (mmCreate *mChatRepositoryMockCreate) Set(f func(ctx context.Context, chat *model.Chat) (i1 int64, err error)) *ChatRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the ChatRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the ChatRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the ChatRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mChatRepositoryMockCreate) When(ctx context.Context, chat *model.Chat) *ChatRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	expectation := &ChatRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &ChatRepositoryMockCreateParams{ctx, chat},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up ChatRepository.Create return parameters for the expectation previously defined by the When method
func (e *ChatRepositoryMockCreateExpectation) Then(i1 int64, err error) *ChatRepositoryMock {
	e.results = &ChatRepositoryMockCreateResults{i1, err}
	return e.mock
}

// Times sets number of times ChatRepository.Create should be invoked
func (mmCreate *mChatRepositoryMockCreate) Times(n uint64) *mChatRepositoryMockCreate {
	if n == 0 {
		mmCreate.mock.t.Fatalf("Times of ChatRepositoryMock.Create mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreate.expectedInvocations, n)
	return mmCreate
}

func (mmCreate *mChatRepositoryMockCreate) invocationsDone() bool {
	if len(mmCreate.expectations) == 0 && mmCreate.defaultExpectation == nil && mmCreate.mock.funcCreate == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreate.mock.afterCreateCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreate.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Create implements repository.ChatRepository
func (mmCreate *ChatRepositoryMock) Create(ctx context.Context, chat *model.Chat) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, chat)
	}

	mm_params := ChatRepositoryMockCreateParams{ctx, chat}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_want_ptrs := mmCreate.CreateMock.defaultExpectation.paramPtrs

		mm_got := ChatRepositoryMockCreateParams{ctx, chat}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreate.t.Errorf("ChatRepositoryMock.Create got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.chat != nil && !minimock.Equal(*mm_want_ptrs.chat, mm_got.chat) {
				mmCreate.t.Errorf("ChatRepositoryMock.Create got unexpected parameter chat, want: %#v, got: %#v%s\n", *mm_want_ptrs.chat, mm_got.chat, minimock.Diff(*mm_want_ptrs.chat, mm_got.chat))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("ChatRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the ChatRepositoryMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, chat)
	}
	mmCreate.t.Fatalf("Unexpected call to ChatRepositoryMock.Create. %v %v", ctx, chat)
	return
}

// CreateAfterCounter returns a count of finished ChatRepositoryMock.Create invocations
func (mmCreate *ChatRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of ChatRepositoryMock.Create invocations
func (mmCreate *ChatRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to ChatRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mChatRepositoryMockCreate) Calls() []*ChatRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*ChatRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *ChatRepositoryMock) MinimockCreateDone() bool {
	if m.CreateMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateMock.invocationsDone()
}

// MinimockCreateInspect logs each unmet expectation
func (m *ChatRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	afterCreateCounter := mm_atomic.LoadUint64(&m.afterCreateCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && afterCreateCounter < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to ChatRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && afterCreateCounter < 1 {
		m.t.Error("Expected call to ChatRepositoryMock.Create")
	}

	if !m.CreateMock.invocationsDone() && afterCreateCounter > 0 {
		m.t.Errorf("Expected %d calls to ChatRepositoryMock.Create but found %d calls",
			mm_atomic.LoadUint64(&m.CreateMock.expectedInvocations), afterCreateCounter)
	}
}

type mChatRepositoryMockDelete struct {
	optional           bool
	mock               *ChatRepositoryMock
	defaultExpectation *ChatRepositoryMockDeleteExpectation
	expectations       []*ChatRepositoryMockDeleteExpectation

	callArgs []*ChatRepositoryMockDeleteParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// ChatRepositoryMockDeleteExpectation specifies expectation struct of the ChatRepository.Delete
type ChatRepositoryMockDeleteExpectation struct {
	mock      *ChatRepositoryMock
	params    *ChatRepositoryMockDeleteParams
	paramPtrs *ChatRepositoryMockDeleteParamPtrs
	results   *ChatRepositoryMockDeleteResults
	Counter   uint64
}

// ChatRepositoryMockDeleteParams contains parameters of the ChatRepository.Delete
type ChatRepositoryMockDeleteParams struct {
	ctx  context.Context
	chat *model.Chat
}

// ChatRepositoryMockDeleteParamPtrs contains pointers to parameters of the ChatRepository.Delete
type ChatRepositoryMockDeleteParamPtrs struct {
	ctx  *context.Context
	chat **model.Chat
}

// ChatRepositoryMockDeleteResults contains results of the ChatRepository.Delete
type ChatRepositoryMockDeleteResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmDelete *mChatRepositoryMockDelete) Optional() *mChatRepositoryMockDelete {
	mmDelete.optional = true
	return mmDelete
}

// Expect sets up expected params for ChatRepository.Delete
func (mmDelete *mChatRepositoryMockDelete) Expect(ctx context.Context, chat *model.Chat) *mChatRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatRepositoryMockDeleteExpectation{}
	}

	if mmDelete.defaultExpectation.paramPtrs != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by ExpectParams functions")
	}

	mmDelete.defaultExpectation.params = &ChatRepositoryMockDeleteParams{ctx, chat}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// ExpectCtxParam1 sets up expected param ctx for ChatRepository.Delete
func (mmDelete *mChatRepositoryMockDelete) ExpectCtxParam1(ctx context.Context) *mChatRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatRepositoryMockDeleteExpectation{}
	}

	if mmDelete.defaultExpectation.params != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Expect")
	}

	if mmDelete.defaultExpectation.paramPtrs == nil {
		mmDelete.defaultExpectation.paramPtrs = &ChatRepositoryMockDeleteParamPtrs{}
	}
	mmDelete.defaultExpectation.paramPtrs.ctx = &ctx

	return mmDelete
}

// ExpectChatParam2 sets up expected param chat for ChatRepository.Delete
func (mmDelete *mChatRepositoryMockDelete) ExpectChatParam2(chat *model.Chat) *mChatRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatRepositoryMockDeleteExpectation{}
	}

	if mmDelete.defaultExpectation.params != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Expect")
	}

	if mmDelete.defaultExpectation.paramPtrs == nil {
		mmDelete.defaultExpectation.paramPtrs = &ChatRepositoryMockDeleteParamPtrs{}
	}
	mmDelete.defaultExpectation.paramPtrs.chat = &chat

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the ChatRepository.Delete
func (mmDelete *mChatRepositoryMockDelete) Inspect(f func(ctx context.Context, chat *model.Chat)) *mChatRepositoryMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for ChatRepositoryMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by ChatRepository.Delete
func (mmDelete *mChatRepositoryMockDelete) Return(err error) *ChatRepositoryMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatRepositoryMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &ChatRepositoryMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the ChatRepository.Delete method
func (mmDelete *mChatRepositoryMockDelete) Set(f func(ctx context.Context, chat *model.Chat) (err error)) *ChatRepositoryMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the ChatRepository.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the ChatRepository.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the ChatRepository.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mChatRepositoryMockDelete) When(ctx context.Context, chat *model.Chat) *ChatRepositoryMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Set")
	}

	expectation := &ChatRepositoryMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &ChatRepositoryMockDeleteParams{ctx, chat},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up ChatRepository.Delete return parameters for the expectation previously defined by the When method
func (e *ChatRepositoryMockDeleteExpectation) Then(err error) *ChatRepositoryMock {
	e.results = &ChatRepositoryMockDeleteResults{err}
	return e.mock
}

// Times sets number of times ChatRepository.Delete should be invoked
func (mmDelete *mChatRepositoryMockDelete) Times(n uint64) *mChatRepositoryMockDelete {
	if n == 0 {
		mmDelete.mock.t.Fatalf("Times of ChatRepositoryMock.Delete mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmDelete.expectedInvocations, n)
	return mmDelete
}

func (mmDelete *mChatRepositoryMockDelete) invocationsDone() bool {
	if len(mmDelete.expectations) == 0 && mmDelete.defaultExpectation == nil && mmDelete.mock.funcDelete == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmDelete.mock.afterDeleteCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmDelete.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Delete implements repository.ChatRepository
func (mmDelete *ChatRepositoryMock) Delete(ctx context.Context, chat *model.Chat) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, chat)
	}

	mm_params := ChatRepositoryMockDeleteParams{ctx, chat}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, &mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_want_ptrs := mmDelete.DeleteMock.defaultExpectation.paramPtrs

		mm_got := ChatRepositoryMockDeleteParams{ctx, chat}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmDelete.t.Errorf("ChatRepositoryMock.Delete got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.chat != nil && !minimock.Equal(*mm_want_ptrs.chat, mm_got.chat) {
				mmDelete.t.Errorf("ChatRepositoryMock.Delete got unexpected parameter chat, want: %#v, got: %#v%s\n", *mm_want_ptrs.chat, mm_got.chat, minimock.Diff(*mm_want_ptrs.chat, mm_got.chat))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("ChatRepositoryMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the ChatRepositoryMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, chat)
	}
	mmDelete.t.Fatalf("Unexpected call to ChatRepositoryMock.Delete. %v %v", ctx, chat)
	return
}

// DeleteAfterCounter returns a count of finished ChatRepositoryMock.Delete invocations
func (mmDelete *ChatRepositoryMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of ChatRepositoryMock.Delete invocations
func (mmDelete *ChatRepositoryMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to ChatRepositoryMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mChatRepositoryMockDelete) Calls() []*ChatRepositoryMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*ChatRepositoryMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *ChatRepositoryMock) MinimockDeleteDone() bool {
	if m.DeleteMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.DeleteMock.invocationsDone()
}

// MinimockDeleteInspect logs each unmet expectation
func (m *ChatRepositoryMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatRepositoryMock.Delete with params: %#v", *e.params)
		}
	}

	afterDeleteCounter := mm_atomic.LoadUint64(&m.afterDeleteCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && afterDeleteCounter < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatRepositoryMock.Delete")
		} else {
			m.t.Errorf("Expected call to ChatRepositoryMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && afterDeleteCounter < 1 {
		m.t.Error("Expected call to ChatRepositoryMock.Delete")
	}

	if !m.DeleteMock.invocationsDone() && afterDeleteCounter > 0 {
		m.t.Errorf("Expected %d calls to ChatRepositoryMock.Delete but found %d calls",
			mm_atomic.LoadUint64(&m.DeleteMock.expectedInvocations), afterDeleteCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ChatRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockDeleteInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ChatRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ChatRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone()
}
