package main

import (
	"testing"

	"github.com/runetui/runetui"
	rtest "github.com/runetui/runetui/testing"
)

func TestAsyncExample_RendersLoadingState(t *testing.T) {
	state := &asyncState{loading: true}
	rootFunc, _ := createAsyncApp(state)

	output := rtest.RenderToString(rootFunc, 50, 10)

	runetui.AssertContainsText(t, output, "Loading")
}

func TestAsyncExample_RendersLoadedState(t *testing.T) {
	state := &asyncState{
		loading: false,
		data:    "Hello from server!",
	}
	rootFunc, _ := createAsyncApp(state)

	output := rtest.RenderToString(rootFunc, 50, 10)

	runetui.AssertContainsText(t, output, "Hello from server!")
}

func TestAsyncExample_RendersErrorState(t *testing.T) {
	state := &asyncState{
		loading: false,
		err:     "connection failed",
	}
	rootFunc, _ := createAsyncApp(state)

	output := rtest.RenderToString(rootFunc, 50, 10)

	runetui.AssertContainsText(t, output, "Error")
	runetui.AssertContainsText(t, output, "connection failed")
}

func TestAsyncExample_Snapshot_Loading(t *testing.T) {
	state := &asyncState{loading: true, frame: 2}
	rootFunc, _ := createAsyncApp(state)

	output := rtest.RenderToString(rootFunc, 50, 10)

	rtest.AssertSnapshot(t, "async_loading", output)
}

func TestAsyncExample_Snapshot_Loaded(t *testing.T) {
	state := &asyncState{
		loading: false,
		data:    "Data loaded successfully",
	}
	rootFunc, _ := createAsyncApp(state)

	output := rtest.RenderToString(rootFunc, 50, 10)

	rtest.AssertSnapshot(t, "async_loaded", output)
}
