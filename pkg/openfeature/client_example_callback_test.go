package openfeature_test

import (
	"context"
	"fmt"
	"time"

	"github.com/open-feature/go-sdk/pkg/openfeature"
)

func ExampleClient_BooleanValueCallback() {
	provider := ErrorProvider{}
	openfeature.SetProvider(provider)
	client := openfeature.NewClient("example-client")
	called := false
	expensiveDefault := func() bool {
		// simulate an expensive call
		time.Sleep(1 * time.Second)
		called = true
		return true
	}
	value, _ := client.BooleanValueCallback(
		context.Background(), "test-flag", expensiveDefault, openfeature.EvaluationContext{},
	)
	fmt.Printf("callback called: %v - result: %v", called, value)
	// Output: callback called: true - result: true
}

// ErrorProvider implements the FeatureProvider interface and provides functions for evaluating flags
type ErrorProvider struct {
}

// Metadata returns the metadata of the provider
func (e ErrorProvider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{Name: "ErrorProvider"}
}

// BooleanEvaluation returns a boolean flag.
func (e ErrorProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	return openfeature.BoolResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewFlagNotFoundResolutionError(""),
			Reason:          openfeature.DefaultReason,
		},
	}
}

// StringEvaluation returns a string flag.
func (e ErrorProvider) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	return openfeature.StringResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewFlagNotFoundResolutionError(""),
			Reason:          openfeature.DefaultReason,
		},
	}
}

// FloatEvaluation returns a float flag.
func (e ErrorProvider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	return openfeature.FloatResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewFlagNotFoundResolutionError(""),
			Reason:          openfeature.DefaultReason,
		},
	}
}

// IntEvaluation returns an int flag.
func (e ErrorProvider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	return openfeature.IntResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewFlagNotFoundResolutionError(""),
			Reason:          openfeature.DefaultReason,
		},
	}
}

// ObjectEvaluation returns an object flag
func (e ErrorProvider) ObjectEvaluation(ctx context.Context, flag string, defaultValue interface{}, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	return openfeature.InterfaceResolutionDetail{
		Value: defaultValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			ResolutionError: openfeature.NewFlagNotFoundResolutionError(""),
			Reason:          openfeature.DefaultReason,
		},
	}
}

// Hooks returns hooks
func (e ErrorProvider) Hooks() []openfeature.Hook {
	return []openfeature.Hook{}
}
