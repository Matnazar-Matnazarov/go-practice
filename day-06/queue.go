package main

import "fmt"

// Queue - queue ma'lumot tuzilmasi (slice asosida)
type Queue struct {
	items []int
}

// NewQueue - yangi queue yaratish
func NewQueue() *Queue {
	return &Queue{
		items: make([]int, 0),
	}
}

// Enqueue - element qo'shish (oxiriga)
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue - element olib tashlash (boshidan)
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Front - birinchi elementni ko'rish (olib tashlamasdan)
func (q *Queue) Front() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[0], true
}

// IsEmpty - queue bo'shligini tekshirish
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size - queue uzunligi
func (q *Queue) Size() int {
	return len(q.items)
}

// Display - queue ni ko'rsatish
func (q *Queue) Display() {
	fmt.Printf("Queue: %v\n", q.items)
}

// demonstrateQueueBasic - queue asoslari
func demonstrateQueueBasic() {
	queue := NewQueue()

	fmt.Println("Queue ga elementlar qo'shish:")
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Display()

	fmt.Printf("Queue uzunligi: %d\n", queue.Size())
	fmt.Printf("Queue bo'shmi: %t\n", queue.IsEmpty())
}

// demonstrateQueueOperations - queue operatsiyalari
func demonstrateQueueOperations() {
	queue := NewQueue()

	// Elementlar qo'shish
	fmt.Println("Elementlar qo'shish:")
	for i := 1; i <= 5; i++ {
		queue.Enqueue(i * 10)
		fmt.Printf("  %d qo'shildi\n", i*10)
	}
	queue.Display()

	// Elementlarni olib tashlash
	fmt.Println("\nElementlarni olib tashlash:")
	for !queue.IsEmpty() {
		item, ok := queue.Dequeue()
		if ok {
			fmt.Printf("  %d olib tashlandi\n", item)
		}
		queue.Display()
	}
}

// demonstrateQueueFront - front elementni ko'rish
func demonstrateQueueFront() {
	queue := NewQueue()
	queue.Enqueue(100)
	queue.Enqueue(200)
	queue.Enqueue(300)

	front, ok := queue.Front()
	if ok {
		fmt.Printf("Front element: %d\n", front)
		fmt.Printf("Queue o'zgarmadi: %v\n", queue.items)
	}
}

// demonstrateQueueString - string queue
type StringQueue struct {
	items []string
}

// NewStringQueue - yangi string queue
func NewStringQueue() *StringQueue {
	return &StringQueue{
		items: make([]string, 0),
	}
}

// Enqueue - string qo'shish
func (q *StringQueue) Enqueue(item string) {
	q.items = append(q.items, item)
}

// Dequeue - string olib tashlash
func (q *StringQueue) Dequeue() (string, bool) {
	if len(q.items) == 0 {
		return "", false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// IsEmpty - queue bo'shligini tekshirish
func (q *StringQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// demonstrateStringQueue - string queue misoli
func demonstrateStringQueue() {
	queue := NewStringQueue()

	tasks := []string{"Vazifa 1", "Vazifa 2", "Vazifa 3"}
	for _, task := range tasks {
		queue.Enqueue(task)
	}

	fmt.Println("String queue:")
	for len(queue.items) > 0 {
		task, ok := queue.Dequeue()
		if ok {
			fmt.Printf("  Bajarilmoqda: %s\n", task)
		}
	}
}

// demonstrateQueueGeneric - interface{} bilan generic queue
type GenericQueue struct {
	items []interface{}
}

// NewGenericQueue - yangi generic queue
func NewGenericQueue() *GenericQueue {
	return &GenericQueue{
		items: make([]interface{}, 0),
	}
}

// Enqueue - element qo'shish
func (q *GenericQueue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue - element olib tashlash
func (q *GenericQueue) Dequeue() (interface{}, bool) {
	if len(q.items) == 0 {
		return nil, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// demonstrateGenericQueue - generic queue misoli
func demonstrateGenericQueue() {
	queue := NewGenericQueue()

	queue.Enqueue(10)
	queue.Enqueue("Salom")
	queue.Enqueue(3.14)
	queue.Enqueue(true)

	fmt.Println("Generic queue:")
	for len(queue.items) > 0 {
		item, ok := queue.Dequeue()
		if ok {
			fmt.Printf("  %v (tip: %T)\n", item, item)
		}
	}
}

// demonstrateQueueCircular - circular queue (asosiy)
func demonstrateQueueCircular() {
	// Oddiy slice bilan circular queue simulatsiyasi
	queue := NewQueue()

	// Circular queue kabi ishlash
	for i := 1; i <= 5; i++ {
		queue.Enqueue(i)
	}

	fmt.Println("Circular queue simulatsiyasi:")
	for i := 0; i < 3; i++ {
		item, _ := queue.Dequeue()
		queue.Enqueue(item * 10) // Qayta qo'shish
		queue.Display()
	}
}

// demonstrateQueuePriority - priority queue (asosiy)
func demonstrateQueuePriority() {
	// Oddiy priority queue simulatsiyasi
	queue := NewQueue()

	// Priority bo'yicha qo'shish
	priorities := []int{3, 1, 4, 2, 5}
	for _, p := range priorities {
		queue.Enqueue(p)
	}

	fmt.Println("Priority queue (tartiblash):")
	// Oddiy sort qilish
	sorted := make([]int, 0)
	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		sorted = append(sorted, item)
	}

	// Bubble sort
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	// Priority bo'yicha qayta qo'shish
	for _, item := range sorted {
		queue.Enqueue(item)
	}
	queue.Display()
}

// demonstrateQueueTaskScheduler - amaliy misol: vazifa rejalashtirish
func demonstrateQueueTaskScheduler() {
	taskQueue := NewStringQueue()

	// Vazifalarni qo'shish
	tasks := []string{
		"Email tekshirish",
		"Kod yozish",
		"Test yozish",
		"Code review",
		"Deploy qilish",
	}

	fmt.Println("Vazifalar rejalashtirish:")
	for _, task := range tasks {
		taskQueue.Enqueue(task)
		fmt.Printf("  + %s qo'shildi\n", task)
	}

	fmt.Println("\nVazifalarni bajarish:")
	for !taskQueue.IsEmpty() {
		task, ok := taskQueue.Dequeue()
		if ok {
			fmt.Printf("  → %s bajarilmoqda...\n", task)
			fmt.Printf("  ✓ %s yakunlandi\n", task)
		}
	}
}

// demonstrateQueueBFS - amaliy misol: BFS simulatsiyasi
func demonstrateQueueBFS() {
	queue := NewQueue()

	// BFS: graph node larni qo'shish
	nodes := []int{1, 2, 3, 4, 5}
	for _, node := range nodes {
		queue.Enqueue(node)
	}

	fmt.Println("BFS (Breadth-First Search) simulatsiyasi:")
	level := 0
	for !queue.IsEmpty() {
		level++
		size := queue.Size()
		fmt.Printf("Level %d: ", level)

		for i := 0; i < size; i++ {
			node, _ := queue.Dequeue()
			fmt.Printf("%d ", node)

			// Keyingi level node lar (misol)
			if node < 5 {
				queue.Enqueue(node * 10)
			}
		}
		fmt.Println()
		if level >= 3 {
			break
		}
	}
}
