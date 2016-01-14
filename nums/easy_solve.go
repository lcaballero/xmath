package nums

/*
Find an algorithm to solve this equation: `2x = 100`.

Steps:
- Lex the equation: '2' 'x' '=' '100'
- Parse and analyze the equation to produce: ('=' (* '2' 'x') '100')
  - In this step the implicit multiplication is made explicit.

           =
          / \
         /   \
        *    100
       / \
      /   \
     2     x

     => (= (* 2 x) 100)
     => (= (/ (* 2 x ) 2) (/ 100 2))
     => (= (x) (50))
     => (= x 50)
     >> x = 50


- Create a cost function `J(theta)` that assigns a cost to the function
  `(= 100 (* 2 x))`.
  - In this case the equation or vector of symbols in the given order are `theta`.
- Iterate to minimize `J(theta)`.
  - Iteration here means trying different heuristics, or rebalancing, etc.
  - J(theta) might equal zero when the equation is reduced to `x = 50`.  This
    might suggest that assignment costs 0, and each side could have a cost of
    number of `terminals - 1`.



*/
