Problem statement:
We have a Zomato order placing service. 
We have a microservice to manage food order management, and a microservice for delivery management.
Implement a 2PC protocol, to place order only if both food is available and if delivery agent is available. If either is unavaible, cancel the order