package metric

import (
	"github.com/dgryski/go-topk"
)

// Counter is the struct to calculate the topk elements in the stream.
// Using 'Filtered Space-Saving TopK streaming algorithm'
type Counter struct {
	stream *topk.Stream
}

// NewCounter initialize counter to calculate topk elements
func NewCounter(topK int) Counter {
	return Counter{
		stream: topk.New(topK),
	}
}

// Inc to increment the value of key by 1.
func (c *Counter) Inc(key string) {
	if c != nil && c.stream != nil {
		c.stream.Insert(key, 1)
	}
}

// GetTop returns the top elements in the stream
func (c *Counter) GetTop() (output []PairElement) {
	if c != nil && c.stream != nil {
		output = make([]PairElement, 0)
		for _, element := range c.stream.Keys() {
			output = append(output, PairElement{Value: element.Key, Count: element.Count})
		}
	}

	return output
}
