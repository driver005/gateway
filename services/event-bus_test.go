package services

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Bus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventBus_doSubscribe(t *testing.T) {
	type args struct {
		topic   string
		fn      interface{}
		handler *eventHandler
	}
	tests := []struct {
		name    string
		bus     *EventBus
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bus.doSubscribe(tt.args.topic, tt.args.fn, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("EventBus.doSubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_Subscribe(t *testing.T) {
	type args struct {
		topic string
		fn    interface{}
	}
	tests := []struct {
		name    string
		bus     *EventBus
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bus.Subscribe(tt.args.topic, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("EventBus.Subscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_SubscribeAsync(t *testing.T) {
	type args struct {
		topic         string
		fn            interface{}
		transactional bool
	}
	tests := []struct {
		name    string
		bus     *EventBus
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bus.SubscribeAsync(tt.args.topic, tt.args.fn, tt.args.transactional); (err != nil) != tt.wantErr {
				t.Errorf("EventBus.SubscribeAsync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_SubscribeOnce(t *testing.T) {
	type args struct {
		topic string
		fn    interface{}
	}
	tests := []struct {
		name    string
		bus     *EventBus
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bus.SubscribeOnce(tt.args.topic, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("EventBus.SubscribeOnce() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_SubscribeOnceAsync(t *testing.T) {
	type args struct {
		topic string
		fn    interface{}
	}
	tests := []struct {
		name    string
		bus     *EventBus
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bus.SubscribeOnceAsync(tt.args.topic, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("EventBus.SubscribeOnceAsync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_HasCallback(t *testing.T) {
	type args struct {
		topic string
	}
	tests := []struct {
		name string
		bus  *EventBus
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bus.HasCallback(tt.args.topic); got != tt.want {
				t.Errorf("EventBus.HasCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventBus_Unsubscribe(t *testing.T) {
	type args struct {
		topic   string
		handler interface{}
	}
	tests := []struct {
		name    string
		bus     *EventBus
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bus.Unsubscribe(tt.args.topic, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("EventBus.Unsubscribe() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventBus_Publish(t *testing.T) {
	type args struct {
		topic string
		args  []interface{}
	}
	tests := []struct {
		name string
		bus  *EventBus
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bus.Publish(tt.args.topic, tt.args.args...)
		})
	}
}

func TestEventBus_doPublish(t *testing.T) {
	type args struct {
		handler *eventHandler
		topic   string
		args    []interface{}
	}
	tests := []struct {
		name string
		bus  *EventBus
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bus.doPublish(tt.args.handler, tt.args.topic, tt.args.args...)
		})
	}
}

func TestEventBus_doPublishAsync(t *testing.T) {
	type args struct {
		handler *eventHandler
		topic   string
		args    []interface{}
	}
	tests := []struct {
		name string
		bus  *EventBus
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bus.doPublishAsync(tt.args.handler, tt.args.topic, tt.args.args...)
		})
	}
}

func TestEventBus_removeHandler(t *testing.T) {
	type args struct {
		topic string
		idx   int
	}
	tests := []struct {
		name string
		bus  *EventBus
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bus.removeHandler(tt.args.topic, tt.args.idx)
		})
	}
}

func TestEventBus_findHandlerIdx(t *testing.T) {
	type args struct {
		topic    string
		callback reflect.Value
	}
	tests := []struct {
		name string
		bus  *EventBus
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bus.findHandlerIdx(tt.args.topic, tt.args.callback); got != tt.want {
				t.Errorf("EventBus.findHandlerIdx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventBus_setUpPublish(t *testing.T) {
	type args struct {
		callback *eventHandler
		args     []interface{}
	}
	tests := []struct {
		name string
		bus  *EventBus
		args args
		want []reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bus.setUpPublish(tt.args.callback, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EventBus.setUpPublish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventBus_WaitAsync(t *testing.T) {
	tests := []struct {
		name string
		bus  *EventBus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bus.WaitAsync()
		})
	}
}
