package store

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockExporter struct {
	writer io.Writer
	output string
}

func (m *mockExporter) ExportRun(_ string) {
	m.writer.Write([]byte(m.output))
}

func createMockExporter(writer io.Writer, output string) Exporter {
	exporter := mockExporter{writer: writer, output: output}
	return &exporter
}

func TestStateSystems(t *testing.T) {
	t.Run(
		"test state behaviour",
		func(t *testing.T) {
			strings := []string{
				"goodbye",
				"world",
				"what",
				"a",
				"tragedy",
				"it",
				"is",
			}
			state := createEmptyGlobalState[string]()
			for idx := range strings {
				state.AddObject(&strings[idx])
			}

			// Update basic obecs
			strings[0] = "hello"
			strings[4] = "miracle"

			retrieved := state.GetState()
			// Check lengths same
			assert.Len(t, retrieved, len(strings))

			// Ensure pointers maintain values...
			assert.Equal(t, strings, retrieved)

			// ...but are copies of the underlying object
			for idx := range strings {
				assert.NotSame(t, &strings[idx], &retrieved[idx])
			}

		},
	)
	t.Run(
		"test singleton forms properly",
		func(t *testing.T) {
			// Init constancs
			numInstances := 5000
			lock := sync.Mutex{}
			lockWg := sync.WaitGroup{}

			// Initialise test holders
			var (
				stringSingleton       GlobalState[string]
				firstState, lastState *GlobalState[string]
			)

			// Build a lot of backends in goroutines, add data as we do
			for idx := 0; idx < numInstances; idx++ {
				lockWg.Add(1)
				counter := idx
				go func() {
					state := CreateGlobalSingleton(&stringSingleton, &lock)
					go func() {
						defer lockWg.Done()
						s1 := fmt.Sprint(counter)
						s2 := "-" + s1
						state.AddObject(&s1)
						state.AddObject(&s2)
					}()
					if counter == 0 {
						firstState = state
					}
					if counter == numInstances-1 {
						lastState = state
					}
				}()

			}
			lockWg.Wait()

			// Check for singleton behaviour
			copyFirst := firstState.GetState()
			copyLast := lastState.GetState()
			baseState := stringSingleton.GetState()

			assert.Len(t, baseState, numInstances*2)
			assert.Equal(t, copyFirst, copyLast)
			assert.Equal(t, copyFirst, baseState)

		},
	)
	t.Run(
		"test state manager",
		func(t *testing.T) {
			out1, out2, out3 := "writing some output", "not to be seen", "i hope they find me"
			buf1 := &bytes.Buffer{}
			buf2 := &bytes.Buffer{}
			buf3 := &bytes.Buffer{}
			manager := CreateNewStateManager()
			manager.AddExporter("bob", createMockExporter(buf1, out1))
			manager.AddExporter("bob", createMockExporter(buf2, out2))
			manager.AddExporter("jane", createMockExporter(buf3, out3))

			manager.RunExport("/some/path")

			assert.Equal(t, buf1.String(), out1)
			assert.NotContains(t, buf2.String(), out2)
			assert.Equal(t, buf3.String(), out3)

		},
	)
}
