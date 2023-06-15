//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import "time"

// BatchClientCancelSparkBatchJobOptions contains the optional parameters for the BatchClient.CancelSparkBatchJob method.
type BatchClientCancelSparkBatchJobOptions struct {
	// placeholder for future optional parameters
}

// BatchClientCreateSparkBatchJobOptions contains the optional parameters for the BatchClient.CreateSparkBatchJob method.
type BatchClientCreateSparkBatchJobOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
}

// BatchClientGetSparkBatchJobOptions contains the optional parameters for the BatchClient.GetSparkBatchJob method.
type BatchClientGetSparkBatchJobOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
}

// BatchClientGetSparkBatchJobsOptions contains the optional parameters for the BatchClient.GetSparkBatchJobs method.
type BatchClientGetSparkBatchJobsOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool

	// Optional param specifying which index the list should begin from.
	From *int32

	// Optional param specifying the size of the returned list. By default it is 20 and that is the maximum.
	Size *int32
}

type BatchJob struct {
	// REQUIRED; The session Id.
	ID *int32

	// The application id of this session
	AppID *string

	// The detailed application info.
	AppInfo map[string]*string

	// The artifact identifier.
	ArtifactID *string

	// The error information.
	Errors []*ServiceError

	// The job type.
	JobType  *SparkJobType
	LivyInfo *BatchJobState

	// The log lines.
	LogLines []*string

	// The batch name.
	Name *string

	// The plugin information.
	Plugin *ServicePlugin

	// The Spark batch job result.
	Result *SparkBatchJobResultType

	// The scheduler information.
	Scheduler *Scheduler

	// The Spark pool name.
	SparkPoolName *string

	// The batch state
	State *LivyStates

	// The submitter identifier.
	SubmitterID *string

	// The submitter name.
	SubmitterName *string

	// The tags.
	Tags map[string]*string

	// The workspace name.
	WorkspaceName *string
}

// BatchJobCollection - Response for batch list operation.
type BatchJobCollection struct {
	// REQUIRED; The start index of fetched sessions.
	From *int32

	// REQUIRED; Number of sessions fetched.
	Total *int32

	// Batch list
	Sessions []*BatchJob
}

type BatchJobOptions struct {
	// REQUIRED
	File *string

	// REQUIRED
	Name       *string
	Archives   []*string
	Arguments  []*string
	ArtifactID *string
	ClassName  *string

	// Dictionary of
	Configuration  map[string]*string
	DriverCores    *int32
	DriverMemory   *string
	ExecutorCores  *int32
	ExecutorCount  *int32
	ExecutorMemory *string
	Files          []*string
	Jars           []*string
	PythonFiles    []*string

	// Dictionary of
	Tags map[string]*string
}

type BatchJobState struct {
	// the Spark job state.
	CurrentState *string

	// time that at which "dead" livy state was first seen.
	DeadAt             *time.Time
	JobCreationRequest *Request

	// the time that at which "not_started" livy state was first seen.
	NotStartedAt *time.Time

	// the time that at which "recovering" livy state was first seen.
	RecoveringAt *time.Time

	// the time that at which "running" livy state was first seen.
	RunningAt *time.Time

	// the time that at which "starting" livy state was first seen.
	StartingAt *time.Time

	// the time that at which "success" livy state was first seen.
	SuccessAt *time.Time

	// the time that at which "killed" livy state was first seen.
	TerminatedAt *time.Time
}

type Request struct {
	Archives  []*string
	Arguments []*string
	ClassName *string

	// Dictionary of
	Configuration  map[string]*string
	DriverCores    *int32
	DriverMemory   *string
	ExecutorCores  *int32
	ExecutorCount  *int32
	ExecutorMemory *string
	File           *string
	Files          []*string
	Jars           []*string
	Name           *string
	PythonFiles    []*string
}

type Scheduler struct {
	CancellationRequestedAt *time.Time
	CurrentState            *SchedulerCurrentState
	EndedAt                 *time.Time
	ScheduledAt             *time.Time
	SubmittedAt             *time.Time
}

type ServiceError struct {
	ErrorCode *string
	Message   *string
	Source    *SparkErrorSource
}

type ServicePlugin struct {
	CleanupStartedAt             *time.Time
	CurrentState                 *PluginCurrentState
	MonitoringStartedAt          *time.Time
	PreparationStartedAt         *time.Time
	ResourceAcquisitionStartedAt *time.Time
	SubmissionStartedAt          *time.Time
}

type Session struct {
	// REQUIRED
	ID    *int32
	AppID *string

	// Dictionary of
	AppInfo    map[string]*string
	ArtifactID *string
	Errors     []*ServiceError

	// The job type.
	JobType       *SparkJobType
	LivyInfo      *SessionState
	LogLines      []*string
	Name          *string
	Plugin        *ServicePlugin
	Result        *SparkSessionResultType
	Scheduler     *Scheduler
	SparkPoolName *string

	// The session state.
	State         *LivyStates
	SubmitterID   *string
	SubmitterName *string

	// Dictionary of
	Tags          map[string]*string
	WorkspaceName *string
}

// SessionClientCancelSparkSessionOptions contains the optional parameters for the SessionClient.CancelSparkSession method.
type SessionClientCancelSparkSessionOptions struct {
	// placeholder for future optional parameters
}

// SessionClientCancelSparkStatementOptions contains the optional parameters for the SessionClient.CancelSparkStatement method.
type SessionClientCancelSparkStatementOptions struct {
	// placeholder for future optional parameters
}

// SessionClientCreateSparkSessionOptions contains the optional parameters for the SessionClient.CreateSparkSession method.
type SessionClientCreateSparkSessionOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
}

// SessionClientCreateSparkStatementOptions contains the optional parameters for the SessionClient.CreateSparkStatement method.
type SessionClientCreateSparkStatementOptions struct {
	// placeholder for future optional parameters
}

// SessionClientGetSparkSessionOptions contains the optional parameters for the SessionClient.GetSparkSession method.
type SessionClientGetSparkSessionOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
}

// SessionClientGetSparkSessionsOptions contains the optional parameters for the SessionClient.GetSparkSessions method.
type SessionClientGetSparkSessionsOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool

	// Optional param specifying which index the list should begin from.
	From *int32

	// Optional param specifying the size of the returned list. By default it is 20 and that is the maximum.
	Size *int32
}

// SessionClientGetSparkStatementOptions contains the optional parameters for the SessionClient.GetSparkStatement method.
type SessionClientGetSparkStatementOptions struct {
	// placeholder for future optional parameters
}

// SessionClientGetSparkStatementsOptions contains the optional parameters for the SessionClient.GetSparkStatements method.
type SessionClientGetSparkStatementsOptions struct {
	// placeholder for future optional parameters
}

// SessionClientResetSparkSessionTimeoutOptions contains the optional parameters for the SessionClient.ResetSparkSessionTimeout
// method.
type SessionClientResetSparkSessionTimeoutOptions struct {
	// placeholder for future optional parameters
}

type SessionCollection struct {
	// REQUIRED
	From *int32

	// REQUIRED
	Total    *int32
	Sessions []*Session
}

type SessionOptions struct {
	// REQUIRED
	Name       *string
	Archives   []*string
	Arguments  []*string
	ArtifactID *string
	ClassName  *string

	// Dictionary of
	Configuration  map[string]*string
	DriverCores    *int32
	DriverMemory   *string
	ExecutorCores  *int32
	ExecutorCount  *int32
	ExecutorMemory *string
	File           *string
	Files          []*string
	Jars           []*string
	PythonFiles    []*string

	// Dictionary of
	Tags map[string]*string
}

type SessionState struct {
	BusyAt             *time.Time
	CurrentState       *string
	DeadAt             *time.Time
	ErrorAt            *time.Time
	IdleAt             *time.Time
	JobCreationRequest *Request
	NotStartedAt       *time.Time
	RecoveringAt       *time.Time
	ShuttingDownAt     *time.Time
	StartingAt         *time.Time
	TerminatedAt       *time.Time
}

type Statement struct {
	// REQUIRED
	ID     *int32
	Code   *string
	Output *StatementOutput
	State  *LivyStatementStates
}

type StatementCancellationResult struct {
	// The msg property from the Livy API. The value is always "canceled".
	Message *string
}

type StatementCollection struct {
	// REQUIRED
	Total      *int32
	Statements []*Statement
}

type StatementOptions struct {
	Code *string
	Kind *SparkStatementLanguageType
}

type StatementOutput struct {
	// REQUIRED
	ExecutionCount *int32

	// Anything
	Data       any
	ErrorName  *string
	ErrorValue *string
	Status     *string
	Traceback  []*string
}