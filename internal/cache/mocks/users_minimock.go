// Code generated by http://github.com/gojuno/minimock (v3.3.14). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/andredubov/auth/internal/cache.Users -o users_minimock.go -n UsersMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/andredubov/auth/internal/service/model"
	"github.com/gojuno/minimock/v3"
)

// UsersMock implements cache.Users
type UsersMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, user model.User) (err error)
	inspectFuncCreate   func(ctx context.Context, user model.User)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mUsersMockCreate

	funcDelete          func(ctx context.Context, id int64) (err error)
	inspectFuncDelete   func(ctx context.Context, id int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mUsersMockDelete

	funcGet          func(ctx context.Context, id int64) (up1 *model.User, err error)
	inspectFuncGet   func(ctx context.Context, id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mUsersMockGet
}

// NewUsersMock returns a mock for cache.Users
func NewUsersMock(t minimock.Tester) *UsersMock {
	m := &UsersMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mUsersMockCreate{mock: m}
	m.CreateMock.callArgs = []*UsersMockCreateParams{}

	m.DeleteMock = mUsersMockDelete{mock: m}
	m.DeleteMock.callArgs = []*UsersMockDeleteParams{}

	m.GetMock = mUsersMockGet{mock: m}
	m.GetMock.callArgs = []*UsersMockGetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUsersMockCreate struct {
	optional           bool
	mock               *UsersMock
	defaultExpectation *UsersMockCreateExpectation
	expectations       []*UsersMockCreateExpectation

	callArgs []*UsersMockCreateParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UsersMockCreateExpectation specifies expectation struct of the Users.Create
type UsersMockCreateExpectation struct {
	mock      *UsersMock
	params    *UsersMockCreateParams
	paramPtrs *UsersMockCreateParamPtrs
	results   *UsersMockCreateResults
	Counter   uint64
}

// UsersMockCreateParams contains parameters of the Users.Create
type UsersMockCreateParams struct {
	ctx  context.Context
	user model.User
}

// UsersMockCreateParamPtrs contains pointers to parameters of the Users.Create
type UsersMockCreateParamPtrs struct {
	ctx  *context.Context
	user *model.User
}

// UsersMockCreateResults contains results of the Users.Create
type UsersMockCreateResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreate *mUsersMockCreate) Optional() *mUsersMockCreate {
	mmCreate.optional = true
	return mmCreate
}

// Expect sets up expected params for Users.Create
func (mmCreate *mUsersMockCreate) Expect(ctx context.Context, user model.User) *mUsersMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UsersMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UsersMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.paramPtrs != nil {
		mmCreate.mock.t.Fatalf("UsersMock.Create mock is already set by ExpectParams functions")
	}

	mmCreate.defaultExpectation.params = &UsersMockCreateParams{ctx, user}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// ExpectCtxParam1 sets up expected param ctx for Users.Create
func (mmCreate *mUsersMockCreate) ExpectCtxParam1(ctx context.Context) *mUsersMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UsersMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UsersMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("UsersMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &UsersMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.ctx = &ctx

	return mmCreate
}

// ExpectUserParam2 sets up expected param user for Users.Create
func (mmCreate *mUsersMockCreate) ExpectUserParam2(user model.User) *mUsersMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UsersMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UsersMockCreateExpectation{}
	}

	if mmCreate.defaultExpectation.params != nil {
		mmCreate.mock.t.Fatalf("UsersMock.Create mock is already set by Expect")
	}

	if mmCreate.defaultExpectation.paramPtrs == nil {
		mmCreate.defaultExpectation.paramPtrs = &UsersMockCreateParamPtrs{}
	}
	mmCreate.defaultExpectation.paramPtrs.user = &user

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the Users.Create
func (mmCreate *mUsersMockCreate) Inspect(f func(ctx context.Context, user model.User)) *mUsersMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for UsersMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by Users.Create
func (mmCreate *mUsersMockCreate) Return(err error) *UsersMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UsersMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UsersMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &UsersMockCreateResults{err}
	return mmCreate.mock
}

// Set uses given function f to mock the Users.Create method
func (mmCreate *mUsersMockCreate) Set(f func(ctx context.Context, user model.User) (err error)) *UsersMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the Users.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the Users.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the Users.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mUsersMockCreate) When(ctx context.Context, user model.User) *UsersMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UsersMock.Create mock is already set by Set")
	}

	expectation := &UsersMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &UsersMockCreateParams{ctx, user},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up Users.Create return parameters for the expectation previously defined by the When method
func (e *UsersMockCreateExpectation) Then(err error) *UsersMock {
	e.results = &UsersMockCreateResults{err}
	return e.mock
}

// Times sets number of times Users.Create should be invoked
func (mmCreate *mUsersMockCreate) Times(n uint64) *mUsersMockCreate {
	if n == 0 {
		mmCreate.mock.t.Fatalf("Times of UsersMock.Create mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreate.expectedInvocations, n)
	return mmCreate
}

func (mmCreate *mUsersMockCreate) invocationsDone() bool {
	if len(mmCreate.expectations) == 0 && mmCreate.defaultExpectation == nil && mmCreate.mock.funcCreate == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreate.mock.afterCreateCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreate.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Create implements cache.Users
func (mmCreate *UsersMock) Create(ctx context.Context, user model.User) (err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, user)
	}

	mm_params := UsersMockCreateParams{ctx, user}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_want_ptrs := mmCreate.CreateMock.defaultExpectation.paramPtrs

		mm_got := UsersMockCreateParams{ctx, user}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreate.t.Errorf("UsersMock.Create got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.user != nil && !minimock.Equal(*mm_want_ptrs.user, mm_got.user) {
				mmCreate.t.Errorf("UsersMock.Create got unexpected parameter user, want: %#v, got: %#v%s\n", *mm_want_ptrs.user, mm_got.user, minimock.Diff(*mm_want_ptrs.user, mm_got.user))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("UsersMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the UsersMock.Create")
		}
		return (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, user)
	}
	mmCreate.t.Fatalf("Unexpected call to UsersMock.Create. %v %v", ctx, user)
	return
}

// CreateAfterCounter returns a count of finished UsersMock.Create invocations
func (mmCreate *UsersMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of UsersMock.Create invocations
func (mmCreate *UsersMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to UsersMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mUsersMockCreate) Calls() []*UsersMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*UsersMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *UsersMock) MinimockCreateDone() bool {
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
func (m *UsersMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UsersMock.Create with params: %#v", *e.params)
		}
	}

	afterCreateCounter := mm_atomic.LoadUint64(&m.afterCreateCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && afterCreateCounter < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UsersMock.Create")
		} else {
			m.t.Errorf("Expected call to UsersMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && afterCreateCounter < 1 {
		m.t.Error("Expected call to UsersMock.Create")
	}

	if !m.CreateMock.invocationsDone() && afterCreateCounter > 0 {
		m.t.Errorf("Expected %d calls to UsersMock.Create but found %d calls",
			mm_atomic.LoadUint64(&m.CreateMock.expectedInvocations), afterCreateCounter)
	}
}

type mUsersMockDelete struct {
	optional           bool
	mock               *UsersMock
	defaultExpectation *UsersMockDeleteExpectation
	expectations       []*UsersMockDeleteExpectation

	callArgs []*UsersMockDeleteParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UsersMockDeleteExpectation specifies expectation struct of the Users.Delete
type UsersMockDeleteExpectation struct {
	mock      *UsersMock
	params    *UsersMockDeleteParams
	paramPtrs *UsersMockDeleteParamPtrs
	results   *UsersMockDeleteResults
	Counter   uint64
}

// UsersMockDeleteParams contains parameters of the Users.Delete
type UsersMockDeleteParams struct {
	ctx context.Context
	id  int64
}

// UsersMockDeleteParamPtrs contains pointers to parameters of the Users.Delete
type UsersMockDeleteParamPtrs struct {
	ctx *context.Context
	id  *int64
}

// UsersMockDeleteResults contains results of the Users.Delete
type UsersMockDeleteResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmDelete *mUsersMockDelete) Optional() *mUsersMockDelete {
	mmDelete.optional = true
	return mmDelete
}

// Expect sets up expected params for Users.Delete
func (mmDelete *mUsersMockDelete) Expect(ctx context.Context, id int64) *mUsersMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UsersMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UsersMockDeleteExpectation{}
	}

	if mmDelete.defaultExpectation.paramPtrs != nil {
		mmDelete.mock.t.Fatalf("UsersMock.Delete mock is already set by ExpectParams functions")
	}

	mmDelete.defaultExpectation.params = &UsersMockDeleteParams{ctx, id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// ExpectCtxParam1 sets up expected param ctx for Users.Delete
func (mmDelete *mUsersMockDelete) ExpectCtxParam1(ctx context.Context) *mUsersMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UsersMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UsersMockDeleteExpectation{}
	}

	if mmDelete.defaultExpectation.params != nil {
		mmDelete.mock.t.Fatalf("UsersMock.Delete mock is already set by Expect")
	}

	if mmDelete.defaultExpectation.paramPtrs == nil {
		mmDelete.defaultExpectation.paramPtrs = &UsersMockDeleteParamPtrs{}
	}
	mmDelete.defaultExpectation.paramPtrs.ctx = &ctx

	return mmDelete
}

// ExpectIdParam2 sets up expected param id for Users.Delete
func (mmDelete *mUsersMockDelete) ExpectIdParam2(id int64) *mUsersMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UsersMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UsersMockDeleteExpectation{}
	}

	if mmDelete.defaultExpectation.params != nil {
		mmDelete.mock.t.Fatalf("UsersMock.Delete mock is already set by Expect")
	}

	if mmDelete.defaultExpectation.paramPtrs == nil {
		mmDelete.defaultExpectation.paramPtrs = &UsersMockDeleteParamPtrs{}
	}
	mmDelete.defaultExpectation.paramPtrs.id = &id

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the Users.Delete
func (mmDelete *mUsersMockDelete) Inspect(f func(ctx context.Context, id int64)) *mUsersMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for UsersMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by Users.Delete
func (mmDelete *mUsersMockDelete) Return(err error) *UsersMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UsersMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &UsersMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &UsersMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the Users.Delete method
func (mmDelete *mUsersMockDelete) Set(f func(ctx context.Context, id int64) (err error)) *UsersMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the Users.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the Users.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the Users.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mUsersMockDelete) When(ctx context.Context, id int64) *UsersMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("UsersMock.Delete mock is already set by Set")
	}

	expectation := &UsersMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &UsersMockDeleteParams{ctx, id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up Users.Delete return parameters for the expectation previously defined by the When method
func (e *UsersMockDeleteExpectation) Then(err error) *UsersMock {
	e.results = &UsersMockDeleteResults{err}
	return e.mock
}

// Times sets number of times Users.Delete should be invoked
func (mmDelete *mUsersMockDelete) Times(n uint64) *mUsersMockDelete {
	if n == 0 {
		mmDelete.mock.t.Fatalf("Times of UsersMock.Delete mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmDelete.expectedInvocations, n)
	return mmDelete
}

func (mmDelete *mUsersMockDelete) invocationsDone() bool {
	if len(mmDelete.expectations) == 0 && mmDelete.defaultExpectation == nil && mmDelete.mock.funcDelete == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmDelete.mock.afterDeleteCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmDelete.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Delete implements cache.Users
func (mmDelete *UsersMock) Delete(ctx context.Context, id int64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, id)
	}

	mm_params := UsersMockDeleteParams{ctx, id}

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

		mm_got := UsersMockDeleteParams{ctx, id}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmDelete.t.Errorf("UsersMock.Delete got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmDelete.t.Errorf("UsersMock.Delete got unexpected parameter id, want: %#v, got: %#v%s\n", *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("UsersMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the UsersMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, id)
	}
	mmDelete.t.Fatalf("Unexpected call to UsersMock.Delete. %v %v", ctx, id)
	return
}

// DeleteAfterCounter returns a count of finished UsersMock.Delete invocations
func (mmDelete *UsersMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of UsersMock.Delete invocations
func (mmDelete *UsersMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to UsersMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mUsersMockDelete) Calls() []*UsersMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*UsersMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *UsersMock) MinimockDeleteDone() bool {
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
func (m *UsersMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UsersMock.Delete with params: %#v", *e.params)
		}
	}

	afterDeleteCounter := mm_atomic.LoadUint64(&m.afterDeleteCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && afterDeleteCounter < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UsersMock.Delete")
		} else {
			m.t.Errorf("Expected call to UsersMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && afterDeleteCounter < 1 {
		m.t.Error("Expected call to UsersMock.Delete")
	}

	if !m.DeleteMock.invocationsDone() && afterDeleteCounter > 0 {
		m.t.Errorf("Expected %d calls to UsersMock.Delete but found %d calls",
			mm_atomic.LoadUint64(&m.DeleteMock.expectedInvocations), afterDeleteCounter)
	}
}

type mUsersMockGet struct {
	optional           bool
	mock               *UsersMock
	defaultExpectation *UsersMockGetExpectation
	expectations       []*UsersMockGetExpectation

	callArgs []*UsersMockGetParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// UsersMockGetExpectation specifies expectation struct of the Users.Get
type UsersMockGetExpectation struct {
	mock      *UsersMock
	params    *UsersMockGetParams
	paramPtrs *UsersMockGetParamPtrs
	results   *UsersMockGetResults
	Counter   uint64
}

// UsersMockGetParams contains parameters of the Users.Get
type UsersMockGetParams struct {
	ctx context.Context
	id  int64
}

// UsersMockGetParamPtrs contains pointers to parameters of the Users.Get
type UsersMockGetParamPtrs struct {
	ctx *context.Context
	id  *int64
}

// UsersMockGetResults contains results of the Users.Get
type UsersMockGetResults struct {
	up1 *model.User
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmGet *mUsersMockGet) Optional() *mUsersMockGet {
	mmGet.optional = true
	return mmGet
}

// Expect sets up expected params for Users.Get
func (mmGet *mUsersMockGet) Expect(ctx context.Context, id int64) *mUsersMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UsersMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UsersMockGetExpectation{}
	}

	if mmGet.defaultExpectation.paramPtrs != nil {
		mmGet.mock.t.Fatalf("UsersMock.Get mock is already set by ExpectParams functions")
	}

	mmGet.defaultExpectation.params = &UsersMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// ExpectCtxParam1 sets up expected param ctx for Users.Get
func (mmGet *mUsersMockGet) ExpectCtxParam1(ctx context.Context) *mUsersMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UsersMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UsersMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("UsersMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &UsersMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.ctx = &ctx

	return mmGet
}

// ExpectIdParam2 sets up expected param id for Users.Get
func (mmGet *mUsersMockGet) ExpectIdParam2(id int64) *mUsersMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UsersMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UsersMockGetExpectation{}
	}

	if mmGet.defaultExpectation.params != nil {
		mmGet.mock.t.Fatalf("UsersMock.Get mock is already set by Expect")
	}

	if mmGet.defaultExpectation.paramPtrs == nil {
		mmGet.defaultExpectation.paramPtrs = &UsersMockGetParamPtrs{}
	}
	mmGet.defaultExpectation.paramPtrs.id = &id

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the Users.Get
func (mmGet *mUsersMockGet) Inspect(f func(ctx context.Context, id int64)) *mUsersMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for UsersMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by Users.Get
func (mmGet *mUsersMockGet) Return(up1 *model.User, err error) *UsersMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UsersMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &UsersMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &UsersMockGetResults{up1, err}
	return mmGet.mock
}

// Set uses given function f to mock the Users.Get method
func (mmGet *mUsersMockGet) Set(f func(ctx context.Context, id int64) (up1 *model.User, err error)) *UsersMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the Users.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the Users.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the Users.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mUsersMockGet) When(ctx context.Context, id int64) *UsersMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("UsersMock.Get mock is already set by Set")
	}

	expectation := &UsersMockGetExpectation{
		mock:   mmGet.mock,
		params: &UsersMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up Users.Get return parameters for the expectation previously defined by the When method
func (e *UsersMockGetExpectation) Then(up1 *model.User, err error) *UsersMock {
	e.results = &UsersMockGetResults{up1, err}
	return e.mock
}

// Times sets number of times Users.Get should be invoked
func (mmGet *mUsersMockGet) Times(n uint64) *mUsersMockGet {
	if n == 0 {
		mmGet.mock.t.Fatalf("Times of UsersMock.Get mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmGet.expectedInvocations, n)
	return mmGet
}

func (mmGet *mUsersMockGet) invocationsDone() bool {
	if len(mmGet.expectations) == 0 && mmGet.defaultExpectation == nil && mmGet.mock.funcGet == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmGet.mock.afterGetCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmGet.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Get implements cache.Users
func (mmGet *UsersMock) Get(ctx context.Context, id int64) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := UsersMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_want_ptrs := mmGet.GetMock.defaultExpectation.paramPtrs

		mm_got := UsersMockGetParams{ctx, id}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmGet.t.Errorf("UsersMock.Get got unexpected parameter ctx, want: %#v, got: %#v%s\n", *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmGet.t.Errorf("UsersMock.Get got unexpected parameter id, want: %#v, got: %#v%s\n", *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("UsersMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the UsersMock.Get")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to UsersMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished UsersMock.Get invocations
func (mmGet *UsersMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of UsersMock.Get invocations
func (mmGet *UsersMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to UsersMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mUsersMockGet) Calls() []*UsersMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*UsersMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *UsersMock) MinimockGetDone() bool {
	if m.GetMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.GetMock.invocationsDone()
}

// MinimockGetInspect logs each unmet expectation
func (m *UsersMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UsersMock.Get with params: %#v", *e.params)
		}
	}

	afterGetCounter := mm_atomic.LoadUint64(&m.afterGetCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && afterGetCounter < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UsersMock.Get")
		} else {
			m.t.Errorf("Expected call to UsersMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && afterGetCounter < 1 {
		m.t.Error("Expected call to UsersMock.Get")
	}

	if !m.GetMock.invocationsDone() && afterGetCounter > 0 {
		m.t.Errorf("Expected %d calls to UsersMock.Get but found %d calls",
			mm_atomic.LoadUint64(&m.GetMock.expectedInvocations), afterGetCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UsersMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockDeleteInspect()

			m.MinimockGetInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UsersMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *UsersMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone()
}
