# Go-practice
## Singleton
### Use Case: 
When you want to ensure that a class has only one instance and
provide a global point of access to that instance.
### Related api
api/designPattern/getSingleton
api/designPattern/updateSingleton
### Benefit: 
Resource Sharing
### Example: 
Database Connection

## Factory pattern
### Use Case: 
When you want to delegate the responsibility of instantiating objects to a factory class,
providing an interface for creating instances of a class.
### Related api
api/designPattern/drawShape
### Benefit: 
The Factory Pattern allows you to introduce new types of objects or classes without changing
the client code. Clients only interact with the factory interface,
making it easy to extend or modify the system with minimal impact.
### Example: 
cryptoFactory that creates different types of crypto(TRX,BSC)

