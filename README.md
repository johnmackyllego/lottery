# Lottery Process Flow Diagram

Lotteries are an excellent use case for Ethereum. This repo will show the process flow diagram of lotteries in Ethereum and Hyperledger to show it's difference

## Ethereum
Following are the process diagrams of the lottery games.

* [Simple Lottery](https://github.com/johnmackyllego/lottery/blob/master/ethereum/simple-lottery/Simple-lottery.jpg) - This simple-lottery is nonrecurring, uses blockhashes for random numbers, and has only
one winner.

* [Recurring Lottery](https://github.com/johnmackyllego/lottery/blob/master/ethereum/recurring/RecurringLottery(eth).jpg) - The
recurring lottery will occur in rounds so that a new prize pool is started every
time the old one closes. It will also allow users to purchase multiple
tickets in one transaction instead of just one and add a couple of security
improvements

* [RNG Lottery](https://github.com/johnmackyllego/lottery/blob/master/ethereum/rng%20lottery/RNG%20Lotter.jpg) - The RNG Lottery is non-recurring, the numbers are randomly generated. Also after the round all users must reveal their tickets or else it will drop.


* [Powerball](https://github.com/johnmackyllego/lottery/blob/master/ethereum/powerball/Ethereum-PowerBall-Flow-Design.jpg) - In Powerball, the user picks six numbers per ticket. The first five
numbers are standard numbers from 1–69, and the sixth number is a
special Powerball number from 1–26 that offers extra rewards. Every three
or four days, a drawing is held, and a winning ticket consisting of five
standard numbers and a Powerball number is picked. Prizes are paid out
based on the number of winning numbers matched on your ticket.

# Hyperledger



* [Simple Lottery](https://github.com/johnmackyllego/lottery/blob/master/hyperledger/simple-lottery/hyperledger-simple-lottery.jpg) - The Simple Lottery in the Ethereum was translated into Hyperledger platform and every stake need to register into network to participate in the lottery hyperledger network.

* [Recurring Lottery](https://github.com/johnmackyllego/lottery/blob/master/hyperledger/Recurring/Recurring%20Hyper.jpg)  - The Recurring Lottery is a lottery that reccurs in rounds so that a new prize pool is started every time the old one closes.  The recurring in hyperledger the stake need to be a member in the hyperledger network to join in the lottery.

* [RNG Lottery](https://raw.githubusercontent.com/johnmackyllego/lottery/master/hyperledger/RNG-Lottery/RNG.jpg) - The RNG Lottery was a counter park of Recurring Lottery, because the RNG Lottery does not reccur every round and only has one winner and the stake need to be a member to participate in the network

* [Power Ball Lottery](https://github.com/johnmackyllego/lottery/tree/master/hyperledger/powerball) - The Power Ball Lottery is different from the other lotteries because the power ball draw the ticket every three or four days unlike in other lotteries that you can set the ticket round for the day, also the stakes need to be a member in the network to participate in the power ball lotttey.




## Contributors:

* Use cases is based from a book that is tutorial for building games using ethereum
* [John Machy Llego](https://github.com/johnmackyllego) 
* [Paul Kristian Erandio](https://github.com/tensaipaul) 
* [Mark Renz Manalo](https://github.com/mark-renz)
* [Justine Mervic Berango](https://github.com/hustino)
* [Kevin Perez](https://github.com/kvzprz)
