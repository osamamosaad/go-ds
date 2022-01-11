package queue

import "testing"

// TestEnqueueStoreDataCorrectly Test Enqueue is pushed the data in the queue, and make sure its stored correctly
func TestEnqueueStoreDataCorrectly(t *testing.T) {
	t.Errorf("enqueue didn't store the data as expected")
}

// TestDequeueRemoveDataFromQueue Test Dequeue is impelemented and pop out the data from queue and make sure the data removed
func TestDequeueRemoveDataFromQueue(t *testing.T) {
	t.Errorf("dequeue didn't remove the data from queue as expected")
}

// TestPeekReturnTheTopOfTheQueue Test the peek returns the Item that in the top
func TestPeekReturnTheTopOfTheQueue(t *testing.T) {
	t.Errorf("peek return the wrong data in the top of the queue")
}

func TestIsEmptyInInitiationState(t *testing.T) {
	t.Errorf("queue isn't empty in initiation state")
}

func TestIsEmptyAfterDequeuingAllDataFromQueue(t *testing.T) {
	t.Errorf("queue isn't empty after dequeuing all values from the queue, by using dequeue method")
}
