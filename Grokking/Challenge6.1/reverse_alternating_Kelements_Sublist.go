package main

import "fmt"

/*
Given the head of a LinkedList and a number ‘k’, reverse every alternating ‘k’ sized sub-list starting from the head.
If, in the end, you are left with a sub-list with less than ‘k’ elements, reverse it too.
*/

// Definizione della struttura ListNode per la LinkedList
type ListNode struct {
	Value int
	Next  *ListNode
}

// Funzione per invertire ogni sottolista alternata di dimensione 'k' a partire dalla testa della LinkedList.
// Se alla fine rimane una sottolista con meno di 'k' elementi, la inverte anche quella.
func reverseAlternatingKElementsSublist(head *ListNode, k int) *ListNode {
	// Gestire i casi speciali in cui k è minore o uguale a 1 o la testa è nulla
	if k <= 1 || head == nil {
		return head
	}

	var (
		current  = head    // Nodo attuale durante la scansione
		previous *ListNode // Nodo precedente durante la scansione
	)

	// Ciclo principale per scorrere la LinkedList
	for current != nil {
		lastNodeOfSublist := current // Ultimo nodo della sottolista da invertire
		lastNodeOfPrev := previous   // Ultimo nodo del gruppo precedente alla sottolista

		// Inverti la sottolista di dimensione 'k'
		for i := 0; i < k && current != nil; i++ {
			next := current.Next
			current.Next = previous
			previous = current
			current = next
		}

		// Aggiorna il collegamento tra il gruppo precedente e la sottolista invertita
		if lastNodeOfPrev == nil {
			head = previous
		} else {
			lastNodeOfPrev.Next = previous
		}

		// Ricollega l'ultimo nodo della sottolista invertita al primo nodo della prossima sottolista
		lastNodeOfSublist.Next = current

		// Salta k passi per raggiungere il prossimo gruppo da invertire
		for i := 0; i < k && current != nil; i++ {
			previous = current
			current = current.Next
		}
	}

	return head
}

// Funzione main per testare la funzione reverseAlternatingKElementsSublist
func main() {
	// Creazione di una LinkedList di esempio
	// In questo esempio, creeremo una LinkedList con i nodi 1, 2, 3, 4, 5, 6, 7, 8, 9
	// Dove il nodo 1 sarà la testa (head) della LinkedList
	node9 := &ListNode{Value: 9, Next: nil}
	node8 := &ListNode{Value: 8, Next: node9}
	node7 := &ListNode{Value: 7, Next: node8}
	node6 := &ListNode{Value: 6, Next: node7}
	node5 := &ListNode{Value: 5, Next: node6}
	node4 := &ListNode{Value: 4, Next: node5}
	node3 := &ListNode{Value: 3, Next: node4}
	node2 := &ListNode{Value: 2, Next: node3}
	node1 := &ListNode{Value: 1, Next: node2}

	// Stampa la LinkedList originale
	fmt.Println("LinkedList originale:")
	printLinkedList(node1)

	// Applica la funzione reverseAlternatingKElementsSublist sulla LinkedList di esempio con k=3
	k := 3
	newHead := reverseAlternatingKElementsSublist(node1, k)

	// Stampa la LinkedList risultante
	fmt.Printf("\nLinkedList dopo l'inversione alternata con k=%d:\n", k)
	printLinkedList(newHead)
}

// Funzione di utilità per stampare la LinkedList
func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
