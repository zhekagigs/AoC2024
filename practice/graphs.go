package main

import (
	"fmt"
	"slices"
)

// Exercise: Build a friend network
// 1. Create structure for storing friends
// 2. Add/remove friendships
// 3. Find friends-of-friends

type FriendNetwork struct {
	friends map[string][]string
}

// TODO: Implement these functions
func NewFriendNetwork() *FriendNetwork {
	return &FriendNetwork{
		friends: make(map[string][]string),
	}
}

func (fn *FriendNetwork) AddFriendship(person1, person2 string) {
	// Add person2 to person1's friend list
	// Add person1 to person2's friend list (undirected graph)
	fn.friends[person1] = append(fn.friends[person1], person2)
	fn.friends[person2] = append(fn.friends[person2], person1)

}

func (fn *FriendNetwork) GetFriends(person string) []string {
	// Return direct friends of person
	return fn.friends[person]
}

func (fn *FriendNetwork) AreFriends(person1, person2 string) bool {
	// Check if two people are direct friends
	fri1 := fn.GetFriends(person1)
	return slices.Contains(fri1, person2)
}

func (fn *FriendNetwork) GetFriendsOfFriends(person string) []string {
	// Return friends of friends (excluding direct friends and person themselves)
	var fOfF []string
	seen := make(map[string]bool)
	seen[person] = true

	friends := fn.GetFriends(person)

	for _, v := range friends {
		seen[v] = true
	}
	fmt.Println(seen)
	for _, friend := range friends {
		friends2 := fn.GetFriends(friend)
		for _, friend2 := range friends2 {
			if !seen[friend2] {
				fOfF = append(fOfF, friend2)
				seen[friend2] = true
			}
		}

	}
	return fOfF
}

// Test data
func main() {
	network := NewFriendNetwork()

	network.AddFriendship("Alice", "Bob")
	network.AddFriendship("Alice", "Carol")
	network.AddFriendship("Bob", "Dave")
	network.AddFriendship("Carol", "Dave")

	// Add some friendships:
	// Alice -> Bob, Carol
	// Bob -> Dave
	// Carol -> Dave

	// Questions to answer:
	// 1. Who are Alice's direct friends?
	fmt.Println(network.GetFriends("Alice"))
	// 2. Who are Alice's friends-of-friends?
	fmt.Println(network.GetFriendsOfFriends("Alice"))
	// 3. Are Bob and Dave friends?
	fmt.Println(network.AreFriends("Bob", "Dave"))
}
