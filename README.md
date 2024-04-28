# BigPay assessment - Autonomous Delivery System

## How to run
### Method 1: Read input from user
```bash
$ go run ./cmd/cli
Enter number of stations
3
Enter stations:
A
B
C

Enter number of edges:
2
Enter edges:
E1,A,B,30
E2,B,C,10

Enter number of parcels:
1
Enter 1 parcels:
K1,5,A,C

Enter number of trains:
1
Enter trains:
Q1,6,B

W=0, T=Q1, N1=B, P1=[], N2=A, P2=[]
W=30, T=Q1, N1=A, P1=[K1], N2=B, P2=[]
W=60, T=Q1, N1=B, P1=[], N2=C, P2=[]
W=70, T=Q1, N1=C, P1=[], N2=C, P2=[K1]
Total wait time: 70
```

### Method 2: Read input from pipe
```bash
$ cat examples/question_example.txt | go run ./cmd/cli                                         master
W=0, T=Q1, N1=B, P1=[], N2=A, P2=[]
W=30, T=Q1, N1=A, P1=[K1], N2=B, P2=[]
W=60, T=Q1, N1=B, P1=[], N2=C, P2=[]
W=70, T=Q1, N1=C, P1=[], N2=C, P2=[K1]
Total wait time: 70
```

## Potential issues with this implementation
1. It is a greedy algorithm, and does not take into account other trains within the network
   - Use a heuristic to make better informed decisions during graph traversal
2. There is currently no way to add new stations dynamically
   - The Deliver function in `internal/logic/trains/deliver.go` should be refactored into a method of a struct, and that struct should provide methods to add new nodes as well
   - Improve the network implementation to allow dynamically adding nodes
3. Missing some edge case error handling, for example if a train cannot reach a destination
   - Implement tests, and improve the error handling for the algorithm
