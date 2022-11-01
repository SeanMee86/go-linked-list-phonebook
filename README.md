# <center>Go Phonebook Console App</center>

A basic console application for storing, retrieving(1 or all), updating, and deleting contacts. 

The main purpose was to demonstrate an understanding of the Linked List data structure and utilize it in an application that would benefit more from a Linked List than a HashMap. The advantage of this can be seen in the print all contacts method where the contacts need to be displayed in Alphabetical order. If this was not a requirement it would be more efficient to utilize a map with the names as the keys and phone numbers as the values so that we could retrieve individual records at a constant time complexity instead of an O(n) time complexity.