package buffer

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"
)

type CycleBuffer struct {
	innerBuffer []string
	writer      io.StringWriter

	capacity int
	offset   int

	ctx    context.Context
	cancel context.CancelFunc
	mutex  sync.Mutex
	signal chan struct{}
	ticker time.Ticker
}

func New(writer io.StringWriter, capacity int) (*CycleBuffer, error) {
	if capacity < 1 {
		return nil, fmt.Errorf("capacity is lower than 1")
	}

	return &CycleBuffer{
		writer:      writer,
		capacity:    capacity,
		innerBuffer: make([]string, capacity),
		mutex:       sync.Mutex{},
		signal:      make(chan struct{}),
		ticker:      *time.NewTicker(time.Second * 5),
	}, nil
}

func (buffer *CycleBuffer) Start(ctx context.Context) {
	if buffer.cancel != nil {
		return
	}

	buffer.ctx, buffer.cancel = context.WithCancel(ctx)
	go buffer.syncDaemon()
}

func (buffer *CycleBuffer) Stop() {
	if buffer.cancel == nil {
		return
	}

	buffer.sync()

	buffer.cancel()
	buffer.cancel = nil
}

func (buffer *CycleBuffer) Sync() {
	buffer.signal <- struct{}{}
}

func (buffer *CycleBuffer) Add(items ...string) {
	for i := range items {
		if buffer.offset >= buffer.capacity {
			buffer.sync()
		}

		buffer.add(items[i] + "\n")
	}

	if buffer.offset >= buffer.capacity {
		buffer.sync()
	}
}

func (buffer *CycleBuffer) add(item string) {
	buffer.mutex.Lock()
	defer buffer.mutex.Unlock()

	buffer.innerBuffer[buffer.offset] = item
	buffer.offset++
}

func (buffer *CycleBuffer) syncDaemon() {
	var shouldWork = true

	for shouldWork {
		select {
		case <-buffer.ticker.C:
			buffer.sync()
		case <-buffer.signal:
			buffer.sync()
		case <-buffer.ctx.Done():
			buffer.sync()
			shouldWork = false
		}
	}
}

func (buffer *CycleBuffer) sync() {
	if buffer.offset == 0 {
		return
	}

	buffer.mutex.Lock()
	defer buffer.mutex.Unlock()

	for i := 0; i < buffer.offset; i++ {
		if _, err := buffer.writer.WriteString(buffer.innerBuffer[i]); err != nil {
			panic(err)
		}
	}

	buffer.offset = 0
}
